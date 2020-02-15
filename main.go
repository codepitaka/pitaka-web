package main

import (
	"src/config"
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
