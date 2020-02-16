package main

import (
	"pitaka-web/src/config"
	"pitaka-web/src/engine"
)

func main() {
	configurations := config.New()
	engine := engine.New()
	
	engine.AddMiddlewares()
	engine.LoadStaticFiles("src/static/templates")
	engine.SetRoutes(configurations)
	
	engine.Run()
}