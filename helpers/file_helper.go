package helpers

import (
	"io"
	"os"
)

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || !os.IsNotExist(err)
}

func GetFileContent(file *os.File) string {
	file.Seek(0, 0)
	bytes, _ := io.ReadAll(file)
	return string(bytes)
}
