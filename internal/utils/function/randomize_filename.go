package function

import (
	"path/filepath"

	"github.com/google/uuid"
)

func RandomizeFilename(filename string) string {
	newFilename := uuid.New()

	extension := filepath.Ext(filename)

	return newFilename.String() + extension
}
