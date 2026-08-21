package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwtware "github.com/gofiber/contrib/v3/jwt" // community library
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	response "kaifin_clone_api/pkg/http"
	types "kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
)

var publicPaths = []string{
	"/api/v1/front/auth/login-user",
	"/api/v1/front/register/create",
	"/api/v1/admin/auth/login",
	"/api/v1/mobile/auth/login-mobile",
}

func isPublicPath(path string) bool {
	for _, p := range publicPaths {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}

// NewJwtMinddleWare Registers JWT Middleware into Fiber App Pipeline
func NewJwtMinddleWare(app *fiber.App, db_pool *sqlx.DB, redis *redis.Client) {
	// 1. Read .env file
	errs := godotenv.Load()
	if errs != nil {
		log.Fatalf("Error loading .env file")
	}

	// 2. Bind Secret Key from .env
	secret_key := os.Getenv("JWT_SECRET_KEY")

	//  Middleware Step 1: Verification Gate
	app.Use(func(c fiber.Ctx) error {

		if isPublicPath(c.Path()) {
			return c.Next()
		}

		// Check WebSocket upgrade request
		if websocketUpgrade := c.Get("Upgrade"); websocketUpgrade == "websocket" {
			webSocketProtocol := c.Get("Sec-webSocket-Protocol")

			if webSocketProtocol == "" {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error": "Missing WebSocket protocol for authentication",
				})
			}

			parts := strings.Split(webSocketProtocol, ",")
			if len(parts) != 2 || strings.TrimSpace(parts[0]) != "Bearer" {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid webSocket protocol authentication format",
				})
			}

			tokenString := strings.TrimSpace(parts[1])

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret_key), nil
			})
			if err != nil || !token.Valid {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid or expired JWT token",
				})
			}

			c.Locals("jwt_data", token)
			c.Set("Sec-WebSocket-Protocol", "Bearer")
			return c.Next()
		}

		return jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte(secret_key)},
			SuccessHandler: func(ctx fiber.Ctx) error {
				if token := jwtware.FromContext(ctx); token != nil {
					ctx.Locals("jwt_data", token)
				}
				return ctx.Next()
			},
			ErrorHandler: func(ctx fiber.Ctx, err error) error {
				return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error":   "Invalid or expired JWT",
					"details": err.Error(),
				})
			},
		})(c)
	})

	// Middleware Step 2: Context Binding Gate
	app.Use(func(c fiber.Ctx) error {
		// ✅ Skip ដដែលសម្រាប់ public path (ព្រោះ Step 1 didn't set jwt_data)
		if isPublicPath(c.Path()) {
			return c.Next()
		}

		rawToken := c.Locals("jwt_data")
		if rawToken == nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: No token data found in context",
			})
		}

		user_token, ok := rawToken.(*jwt.Token)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Token assertion failed",
			})
		}

		uclaim, ok := user_token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Claims assertion failed",
			})
		}

		return handleUserContext(c, uclaim, db_pool, redis)
	})
}

func handleUserContext(c fiber.Ctx, uclaim jwt.MapClaims, db *sqlx.DB, redis *redis.Client) error {
	login_session, ok := uclaim["login_session"].(string)
	if !ok || login_session == "" {
		smg_error := response.NewResponseError(
			translate.Translate(c, "loginSessionMissing"),
			-500,
			fmt.Errorf("%s", translate.Translate(c, "loginSessionMissing")),
		)
		return c.Status(http.StatusUnprocessableEntity).JSON(smg_error)
	}

	username, _ := uclaim["user_name"].(string)
	user_id := uclaim["user_id"].(float64)
	uCtx := types.UserContext{
		UserID:       int64(user_id),
		UserName:     username,
		RoleID:       int(uclaim["role_id"].(float64)),
		LoginSession: login_session,
		Exp:          time.Unix(int64(uclaim["exp"].(float64)), 0),
		UserAgent:    c.Get("User-Agent", "unknown"),
		Ip:           c.IP(),
	}
	c.Locals("UserContext", uCtx)

	return c.Next()
}
