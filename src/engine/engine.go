package engine

import(
	"github.com/gin-gonic/gin"
	"pitaka-web/src/routes"
	"pitaka-web/src/utils"
	"pitaka-web/src/config"
)

type EngineWrapper struct{
	engine *gin.Engine
}

func New() *EngineWrapper{
	engineWrapper := EngineWrapper{}
	ginEngine := gin.New()
	engineWrapper.engine = ginEngine
	return &engineWrapper
}

func (engineWrapper EngineWrapper) AddMiddlewares(){
	engineWrapper.engine.Use(gin.Logger())
	engineWrapper.engine.Use(gin.Recovery())
}

func (engineWrapper EngineWrapper) LoadStaticFiles(directoryPath string){
	var htmlTemplatePaths []string = utils.HTMLTemplatePathsUnder(directoryPath)
	engineWrapper.engine.LoadHTMLFiles(htmlTemplatePaths...)
	
	//todo loadTemplate과 중복된다. CSS만 뽑아내지 못한다.
	engineWrapper.engine.Static("/static/templates", directoryPath)
}

func (engineWrapper EngineWrapper) SetRoutes(configurations *config.Configuration){
	routes.SetRouter(engineWrapper.engine, configurations)
}

func (engineWrapper EngineWrapper) Run(){
	err := engineWrapper.engine.Run()
	if err!= nil{
		panic(err)
	}
}