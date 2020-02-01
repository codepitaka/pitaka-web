package main

import (
	"github.com/codepitaka/pitaka-web/src/routes"
	"github.com/codepitaka/pitaka-web/src/utils"
	"github.com/gin-gonic/gin"
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

func addMiddlewares(engine *gin.Engine) *gin.Engine {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	return engine
}

func loadTemplates(engine *gin.Engine) *gin.Engine {
	var templatePaths []string = utils.FilePathsUnder("src/static/templates")
	engine.LoadHTMLFiles(templatePaths...)
	return engine
}
