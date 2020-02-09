package main

import (
	"github.com/codepitaka/pitaka-web/src/config"
	"github.com/codepitaka/pitaka-web/src/engine"
)

func main() {
	configurations := config.New()
	engine := engine.New()
	
	engine.AddMiddlewares()
	engine.LoadStaticFiles("src/static/templates")
	engine.SetRoutes(configurations)
	
	engine.Run()
}