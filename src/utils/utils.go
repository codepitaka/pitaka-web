package utils

import (
	"os"
	"log"
	"path/filepath"
	"strings"
)

func HTMLTemplatePathsUnder(path string) []string{
	var targetPaths []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if !isTemplate(path) {
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

func isTemplate(path string) bool{
	return strings.HasSuffix(path, ".tmpl") || strings.HasSuffix(path, ".html")
}