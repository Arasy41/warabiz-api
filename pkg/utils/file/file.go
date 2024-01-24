package file

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// Functions used to get file extension {sample: .png .jpg .zip}
func GetFileExtension(filename string) string {
	return filepath.Ext(filename)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func SaveFile(file *multipart.FileHeader, rootPath string, path string) error {

	//* Check if exist
	if ok := IsExist(rootPath); !ok {
		return errors.New("can't find root path")
	}
	parts := strings.Split(path, "/")

	if err := os.MkdirAll(strings.Replace(rootPath+path, parts[len(parts)-1], "", 1), 0755); err != nil {
		return err
	}

	dst, err := os.Create(rootPath + path)
	if err != nil {
		return err
	}
	defer dst.Close()

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFile(fullpath string) error {
	return os.Remove(fullpath)
}
