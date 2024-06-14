package firebase

import (
	"context"
	"mime/multipart"

	firebase "firebase.google.com/go"
	envVariable "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
)

type FilesRequest struct {
	File       *multipart.FileHeader
	FolderPath *string
}

type UploadFileToFirebaseRequest struct {
	FirebaseProject string
	BucketName      string
	Files           []FilesRequest
}

type FirebaseService struct {
	Env  envVariable.Config
	apps map[string]*firebase.App
}

type FirebaseServiceConfig struct {
	Env envVariable.Config
}

type FirebaseServiceProvider interface {
	initializeFirebase()
	UploadFile(ctx context.Context, request UploadFileToFirebaseRequest) ([]string, error)
}
