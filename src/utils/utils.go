package utils

import (
	"os"
	"log"
	"path/filepath"
)

func FilePathsUnder(path string) []string{
	var targetPaths []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if err != nil {
				return err
			}
			targetPaths = append(targetPaths, path)
			return nil
	})
	if err != nil {
		log.Println(err)
	}
	return targetPaths
}