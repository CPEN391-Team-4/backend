package videostore

import (
	"bytes"
	"fmt"
	uuid "github.com/google/uuid"
	"io/fs"
	"os"
	"strconv"
)

// rwx rwx r-x
const DIR_PERM = 0775

type FileWriter struct {
	Directory string
}

func (fw *FileWriter) Save(dir string, frameNum int, data bytes.Buffer) (string, error) {
	subPath := dir + "/" + strconv.Itoa(frameNum) + ".raw"
	path := fw.Directory + "/" + subPath
	fmt.Println("Saving file to", path)

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
func (fw *FileWriter) RemoveSubdir(id string) error {
	err := os.RemoveAll(fw.Directory + "/" + id)
	if err != nil {
		return fmt.Errorf("cannot remove directory: %w", err)
	}
	return nil
}

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