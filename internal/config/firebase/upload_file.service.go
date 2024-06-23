package firebase

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"

	"github.com/Redchlorophyll/personal-service/internal/utils/function"
	"github.com/gofiber/fiber/v2/log"
)

func (firebaseService FirebaseService) UploadFile(ctx context.Context, request UploadFileToFirebaseRequest) ([]string, error) {
	results := make([]string, len(request.Files))

	app := firebaseService.apps[request.FirebaseProject]
	client, err := app.Storage(ctx)
	if err != nil {
		log.Error("[config][firebase][UploadFile] error when execute app.Storage(). ", err, ctx, request)
		return []string{}, err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Error("[config][firebase][UploadFile] error when execute client.DefaultBucket(). ", err, ctx, request)
		return []string{}, err
	}

	for index, file := range request.Files {
		multipartFile, err := file.File.Open()
		if err != nil {
			log.Error("[config][firebase][UploadFile] error when execute file.File.Open(). ", err, ctx, request)

			results[index] = "[ERROR]"
		}

		byteContainer, err := ioutil.ReadAll(multipartFile)
		if err != nil {
			log.Error("[config][firebase][UploadFile] error when execute  ioutil.ReadAll(). ", err, ctx, request)

			results[index] = "[ERROR] failed to read file"
		}

		fileName := function.RandomizeFilename(file.File.Filename)
		if file.FolderPath != nil {
			fileName = firebaseService.Env.Env + "/" + *file.FolderPath + fileName
		}

		object := bucket.Object(fileName)
		writer := object.NewWriter(ctx)
		defer writer.Close()

		_, err = io.Copy(writer, bytes.NewReader(byteContainer))
		if err != nil {
			log.Error("[config][firebase][UploadFile] error when execute io.Copy(). ", err, ctx, request)

			results[index] = "[ERROR]: failed to upload file"
			continue
		}

		fileURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", request.BucketName, url.PathEscape(fileName))

		results[index] = fileURL
	}

	return results, nil
}
