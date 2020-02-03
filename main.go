package main

import (
	"github.com/gin-gonic/gin"
	"github.com/codepitaka/pitaka-web/src/config"
	"github.com/codepitaka/pitaka-web/src/routes"
	"github.com/codepitaka/pitaka-web/src/utils"
)

func main() {
	configurations := config.New()
	
	engine := NewEngine()
	
	engine.addMiddlewares()
	engine.loadStaticFiles()
	engine.setRoutes(configurations)
	
	engine.Run()
}

type EngineWrapper struct{
	engine *gin.Engine
}

func NewEngine() *EngineWrapper{
	engineWrapper := EngineWrapper{}
	ginEngine := gin.New()
	engineWrapper.engine = ginEngine
	return &engineWrapper
}

func (engineWrapper *EngineWrapper) addMiddlewares(){
	engineWrapper.engine.Use(gin.Logger())
	engineWrapper.engine.Use(gin.Recovery())
}

func (engineWrapper *EngineWrapper) loadStaticFiles(){
	//todo loadCSS와 중복된다. templates만 뽑아내지 못한다.
	var templatePaths []string = utils.FilePathsUnder("src/static/templates")
	engineWrapper.engine.LoadHTMLFiles(templatePaths...)
	
	//todo loadTemplate과 중복된다. CSS만 뽑아내지 못한다.
 	//loadTemplate없이 loadCSS에 있는 내부함수만으로 html까지 제공할 수 있으려나?
	engineWrapper.engine.Static("/css", "src/static/templates")
}

func (engineWrapper *EngineWrapper) setRoutes(configurations *config.Configuration){
	routes.SetRouter(engineWrapper.engine, configurations)
}

func (engineWrapper *EngineWrapper) Run(){
	engineWrapper.engine.Run()
}