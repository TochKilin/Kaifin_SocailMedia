package post_mobile

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type PostMobileHandlerImpl struct {
	ps PostService
	ws *websocket.WebSocketManager
}

func NewPostMobileHandlerImpl(dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *PostMobileHandlerImpl {
	return &PostMobileHandlerImpl{
		ps: NewPostServiceImpl(dbpool, rdb),
		ws: ws,
	}
}

func (h *PostMobileHandlerImpl) Create(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	fmt.Println("Has Context:", ok)
	fmt.Printf("Context: %+v\n", uCtx)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError(
				"Unauthorized",
				constants.Generic_invalid,
				errors.New("user context not found"),
			),
		)
	}

	req := &CreatePostMobileRequest{}
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request body", constants.Generic_invalid, err),
		)
	}
	if err := v.Validate(req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			fe := ve[0]
			msg, _ := translate.TranslateWithError(c, "validation_"+fe.Tag(),
				map[string]any{
					"Field": fe.Field(),
					"Param": fe.Param(),
				})
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(msg, constants.Generic_invalid, err),
			)
		}
		return err
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	req.Hashtags = form.Value["tag_name"]
	req.StickerIDs = nil
	for _, s := range form.Value["sticker_ids[]"] {
		if id, convErr := strconv.ParseInt(s, 10, 64); convErr == nil {
			req.StickerIDs = append(req.StickerIDs, id)
		}
	}

	files := form.File["image"]
	fmt.Printf("%+v\n", req)
	fmt.Println("hashtags =", req.Hashtags)
	fmt.Println("Manually parsed StickerIDs:", req.StickerIDs)

	videoFile, err := c.FormFile("video")
	if err == nil {
		filePath := fmt.Sprintf("uploads/%s", videoFile.Filename)

		if err := c.SaveFile(videoFile, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(
					"upload_failed",
					constants.Generic_error,
					err,
				),
			)
		}

		videoURL := "/" + filePath
		req.VideoURL = &videoURL
	}

	for _, file := range files {
		filePath := fmt.Sprintf("uploads/%s", file.Filename)

		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(500).JSON(
				response.NewResponseError(
					"upload_failed",
					constants.Generic_error,
					err,
				),
			)
		}

		fmt.Println(filePath)
		req.ImageURLs = append(req.ImageURLs, "/"+filePath)
	}

	e := h.ps.Create(req, &uCtx)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}
	msg, e_msg := translate.TranslateWithError(c, "user_created")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		response.NewResponse(msg, constants.Generic_success, true),
	)
}

func (h *PostMobileHandlerImpl) Show(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		b, _ := json.MarshalIndent(uCtx, "", "  ")
		h.ps.SetUserCtx(&uCtx) //
		fmt.Println("jwt_data:", string(b))
	}
	var postShowRequest ShowPostMobileRequest
	v := utls.NewValidator()
	if err := postShowRequest.bind(c, v); err != nil {
		msg, err_msg := translate.TranslateWithError(c, "invalid_request")
		if err_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					err_msg.ErrorString(),
					constants.Translate_Failed,
					err_msg.Err,
				),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(
				msg,
				constants.Invalid_request,
				err,
			),
		)
	}
	posts, e := h.ps.Show(postShowRequest)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}
	msg, e_msg := translate.TranslateWithError(c, "posts_retrieved")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponseWithPaing(msg, constants.Generic_success, posts, postShowRequest.PageOption.Page, postShowRequest.PageOption.Perpage, posts.Total),
	)
}

func (h *PostMobileHandlerImpl) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		msg, _ := translate.TranslateWithError(c, "invalid_user_id")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.Generic_invalid, err),
		)
	}

	if e := h.ps.Delete(id); e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}
	msg, e_msg := translate.TranslateWithError(c, "user_deleted")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(msg, constants.Generic_success, nil),
	)
}

func (h *PostMobileHandlerImpl) View(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid post id", constants.Generic_invalid, err),
		)
	}

	newCount, e := h.ps.IncrementView(id)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(msg, constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("view_recorded", constants.Generic_success, fiber.Map{
			"post_id":     id,
			"views_count": newCount,
		}),
	)
}
