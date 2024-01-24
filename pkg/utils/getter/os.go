package getter

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetCurrentDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, nil
}

func GetRootPath() string {
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	return filepath.Join(filepath.Dir(b), "../../../")
}
