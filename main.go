package main

import (
	"github.com/gin-gonic/gin"
	"github.com/codepitaka/pitaka-web/src/routes"
	"os"
	"log"
	"path/filepath"
)

func main() {
	engine := gin.New()
	engine = addMiddlewares(engine)
	engine = loadTemplates(engine)
	engine = routes.SetRouter(engine)
	
	err := engine.Run()
	if err != nil {
		panic(err.Error())
    }
}

func addMiddlewares(engine *gin.Engine) *gin.Engine{
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	return engine
}

func loadTemplates(engine *gin.Engine) *gin.Engine{
	var templatePaths []string
	err := filepath.Walk("src/static/templates",
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if err != nil {
				return err
			}
			templatePaths = append(templatePaths, path)
			return nil
	})
	if err != nil {
		log.Println(err)
	}
	engine.LoadHTMLFiles(templatePaths...)
	return engine
}