package util

import "mime/multipart"

type WebContext interface {
	SaveFile(fileheader *multipart.FileHeader, path string) error
}
