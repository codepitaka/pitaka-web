package engine

import(
	"github.com/gin-gonic/gin"
	"github.com/codepitaka/pitaka-web/src/routes"
	"github.com/codepitaka/pitaka-web/src/utils"
	"github.com/codepitaka/pitaka-web/src/config"
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

func (engineWrapper EngineWrapper) LoadStaticFiles(path string){
	//todo loadCSS와 중복된다. templates만 뽑아내지 못한다.
	var templatePaths []string = utils.FilePathsUnder(path)
	engineWrapper.engine.LoadHTMLFiles(templatePaths...)
	
	//todo loadTemplate과 중복된다. CSS만 뽑아내지 못한다.
 	//loadTemplate없이 loadCSS에 있는 내부함수만으로 html까지 제공할 수 있으려나?
	engineWrapper.engine.Static("/css", path)
	
	engineWrapper.engine.Static("/js", "src/static/templates/vecty/scripts")
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