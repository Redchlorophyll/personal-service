package function

import (
	"context"
	"mime/multipart"
)

type ValidateFileRequest struct {
	File            *multipart.FileHeader
	AllowedFileType []string
	AllowedFileSize int
}

func ValidateFile(context context.Context, request ValidateFileRequest) error {
	return nil
}
