package post

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/websocket"
	// "kaifin_clone_api/internal/front/web/post"
	constants "kaifin_clone_api/pkg/constants"
	response "kaifin_clone_api/pkg/http"
	"kaifin_clone_api/pkg/share"
	"kaifin_clone_api/pkg/translate"
	"kaifin_clone_api/pkg/utls"
)

type PostHandlerImpl struct {
	ps PostService
	ws *websocket.WebSocketManager
}

func NewPostHandlerImpl(dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *PostHandlerImpl {
	return &PostHandlerImpl{
		ps: NewPostServiceImpl(dbpool, rdb),
		ws: ws,
	}
}

func (h *PostHandlerImpl) Create(c fiber.Ctx) error {
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

	req := &CreatePostRequest{}
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

	tagNames := form.Value["tag_name"]
	tagStickerIDs := form.Value["tag_sticker_ids[]"]

	var hashtags []HashtagInput
	for i, name := range tagNames {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		var stickerID *int64
		if i < len(tagStickerIDs) && tagStickerIDs[i] != "" {
			if id, err := strconv.ParseInt(tagStickerIDs[i], 10, 64); err == nil {
				stickerID = &id
			}
		}
		hashtags = append(hashtags, HashtagInput{Name: name, StickerID: stickerID})
	}
	req.Hashtags = hashtags

	req.StickerIDs = nil
	stickerIDStrs := form.Value["sticker_ids[]"]
	for _, idStr := range stickerIDStrs {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}
		if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
			req.StickerIDs = append(req.StickerIDs, id)
		} else {
			fmt.Println("INVALID STICKER ID:", idStr, err)
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

func (h *PostHandlerImpl) Show(c fiber.Ctx) error {
	if uCtx, ok := c.Locals("UserContext").(share.UserContext); ok {
		b, _ := json.MarshalIndent(uCtx, "", "  ")
		h.ps.SetUserCtx(&uCtx) //
		fmt.Println("jwt_data:", string(b))
	}
	var postShowRequest ShowPostRequest
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

func (h *PostHandlerImpl) Delete(c fiber.Ctx) error {
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

func (h *PostHandlerImpl) View(c fiber.Ctx) error {
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

func (h *PostHandlerImpl) CreateShare(c fiber.Ctx) error {
	uCtx, ok := c.Locals("UserContext").(share.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			response.NewResponseError("Unauthorized", constants.Generic_invalid, errors.New("user context not found")),
		)
	}

	req := &CreateShareRequest{}
	v := utls.NewValidator()
	if err := req.bind(c, v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid request body", constants.Generic_invalid, err),
		)
	}

	// 👇 ត្រឹមត្រូវតាម Interface គឺ Return តែ ១ តម្លៃ (e) ប៉ុណ្ណោះ
	e := h.ps.CreateShare(req.PostID, &uCtx)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, e.MessageID)
		if e_msg != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.NewResponseError(msg, constants.Generic_error, e.Err))
	}

	msg, e_msg := translate.TranslateWithError(c, "post_shared")
	if e_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(e_msg.Err.Error(), constants.Translate_Failed, e_msg.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(response.NewResponse(msg, constants.Generic_success, fiber.Map{"shared": true}))
}

func (h *PostHandlerImpl) ShowShares(c fiber.Ctx) error {
	postID, err := strconv.ParseInt(c.Query("post_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid post id", constants.Generic_invalid, err),
		)
	}

	count, e := h.ps.GetShareCount(postID)
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("database_error", constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse("shares_retrieved", constants.Generic_success, fiber.Map{
			"post_id":     postID,
			"share_count": count,
		}),
	)
}

func (h *PostHandlerImpl) ShowProfilePosts(c fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError("Invalid user id", constants.Generic_invalid, err),
		)
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("perpage", "10"))

	postsResp, e := h.ps.ShowWithReposts(userID, page, perPage)
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError("database_error", constants.Generic_error, e.Err),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.NewResponseWithPaing("Profile posts retrieved", constants.Generic_success, postsResp.Posts, page, perPage, postsResp.Total),
	)
}
