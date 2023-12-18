package usecase

import (
	"fiber-upload-file/internal/util"
	"mime/multipart"
)

type Usecase struct{}

// go:generate mockgen -package=mock_usecase -source=internal/usecase/usecase.go -destination=mock/usecase/usecase_mock.go
type UsecaseInterface interface {
	SaveFile(util.WebContext, *multipart.FileHeader) error
	HitAPIsConcurrently() bool
}

func NewUsecase() UsecaseInterface {
	return &Usecase{}
}
