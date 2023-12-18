package usecase

import (
	"fiber-upload-file/internal/util"
	"fmt"
	"mime/multipart"

	"github.com/gofiber/fiber/v2/log"
)

func (u *Usecase) SaveFile(c util.WebContext, file *multipart.FileHeader) error {
	destination := fmt.Sprintf(util.ImagesPath, file.Filename)
	if err := c.SaveFile(file, destination); err != nil {
		log.Errorf("error saving image", err)
		return err
	}
	return nil
}
