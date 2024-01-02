package usecase

import (
	"fiber-upload-file/internal/util"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2/log"
)

func (u *Usecase) SaveFile(file *multipart.FileHeader) error {
	destination := fmt.Sprintf(util.ImagesPath, file.Filename)

	fileData, err := file.Open()
	defer fileData.Close()
	if err != nil {
		return err
	}

	f, err := os.Create(destination)
	if err != nil {
		log.Errorf("error saving image", err)
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, fileData)
	if err != nil {
		log.Errorf("error copying image content", err)
		return err
	}

	return nil
}
