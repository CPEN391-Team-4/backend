package videostore

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"io/fs"
	"log"
	"os"
	"strconv"
)

// Directory permissions: rwx rwx r-x
const DIR_PERM = 0775

type FileWriter struct {
	Directory string
}

// Save: Save a file under <FileWriter.Directory>/<frameNum>.jpg
func (fw *FileWriter) Save(dir string, frameNum int, data bytes.Buffer) (string, error) {
	subPath := dir + "/" + strconv.Itoa(frameNum) + ".jpg"
	path := fw.Directory + "/" + subPath
	log.Println("Saving file to", path)

	file, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("cannot create frame file: %w", err)
	}

	_, err = data.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	return subPath, nil
}

// RemoveSubdir Remove a subdirectory recursively under FileWriter.Directory
func (fw *FileWriter) RemoveSubdir(id string) error {
	err := os.RemoveAll(fw.Directory + "/" + id)
	if err != nil {
		return fmt.Errorf("cannot remove directory: %w", err)
	}
	return nil
}

// CreateSubdir Create a subdirectory under FileWriter.Directory
func (fw *FileWriter) CreateSubdir() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate directory id: %w", err)
	}
	idStr := id.String()

	err = os.Mkdir(fw.Directory + "/" + idStr, fs.FileMode(DIR_PERM))
	if err != nil {
		return "", fmt.Errorf("cannot create directory: %w", err)
	}

	return idStr, nil
}