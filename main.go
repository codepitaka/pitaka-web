package main

import (
	"github.com/gin-gonic/gin"
	"github.com/codepitaka/pitaka-web/src/routes"
	"github.com/codepitaka/pitaka-web/src/utils"
)

func main() {
	engine := gin.New()
	engine = addMiddlewares(engine)
	engine = loadTemplates(engine)
	engine = loadCSS(engine)
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
	//todo loadCSS와 중복된다. templates만 뽑아내지 못한다.
	var templatePaths []string = utils.FilePathsUnder("src/static/templates")
	engine.LoadHTMLFiles(templatePaths...)
	return engine
}

func loadCSS(engine *gin.Engine) *gin.Engine{
	//todo loadTemplate과 중복된다. CSS만 뽑아내지 못한다.
	//loadTemplate없이 loadCSS에 있는 내부함수만으로 html까지 제공할 수 있으려나?
	engine.Static("/css", "src/static/templates")
	return engine
}