package utils

import (
	"os"
	"log"
	"path/filepath"
)

func FilesUnder(path string) []string{
	var templates []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if err != nil {
				return err
			}
			templates = append(templates, path)
			return nil
	})
	if err != nil {
		log.Println(err)
	}
	return templates
}