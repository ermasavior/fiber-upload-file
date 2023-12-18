package main

import (
	"log"

	"fiber-upload-file/internal/handler"
	"fiber-upload-file/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	uc := usecase.NewUsecase()
	h := handler.NewHandler(handler.Handler{
		Usecase: uc,
	})

	config := fiber.Config{
		BodyLimit: 1 * 1024 * 1024, // this is the default limit of 1MB
	}
	app := fiber.New(config)
	app.Post("/upload", h.UploadImage)

	if err := app.Listen(":5000"); err != nil {
		log.Fatal(err)
	}
}
