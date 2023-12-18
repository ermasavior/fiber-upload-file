package handler

import (
	"fiber-upload-file/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Usecase usecase.UsecaseInterface
}

type HandlerInterface interface {
	UploadImage(c *fiber.Ctx) error
}

func NewHandler(h Handler) HandlerInterface {
	return &h
}
