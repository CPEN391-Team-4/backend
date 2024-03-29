package imagestore

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"os"
)

type FileWriter struct {
	Directory string
}

// Save: Save a file under <FileWriter.Directory>/<UUID>.<ext>
// returns <UUID>.<ext>
func (fw *FileWriter) Save(ext string, data bytes.Buffer) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}
	idStr := id.String()

	path := fw.Directory + "/" + idStr + ext
	fmt.Println("Saving file to", path)

	file, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %w", err)
	}

	_, err = data.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	return idStr + ext, nil
}

// Remove the file at <FileWriter.Directory>/<id>
func (fw *FileWriter) Remove(id string) error {
	err := os.Remove(fw.Directory + "/" + id)
	if err != nil {
		return fmt.Errorf("cannot remove file: %w", err)
	}
	return nil
}