package handler

import (
	"fiber-upload-file/internal/model"
	"mime/multipart"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (h *Handler) UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		log.Errorf("error uploading image", err)
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	validateRes := h.validateFile(file)
	if !validateRes.Success {
		return c.Status(fiber.StatusBadRequest).JSON(validateRes)
	}

	err = h.Usecase.SaveFile(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	resBool := h.Usecase.HitAPIsConcurrently()

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Success: resBool,
	})
}

func (h *Handler) validateFile(file *multipart.FileHeader) model.Response {
	fileExtension := filepath.Ext(file.Filename)
	if fileExtension != ".jpg" && fileExtension != ".jpeg" && fileExtension != ".png" {
		return model.Response{
			Success: false,
			Message: "file extension must be .jpg, .jpeg, or .png",
		}
	}

	return model.Response{
		Success: true,
	}
}
