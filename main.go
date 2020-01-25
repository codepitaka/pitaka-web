package main

import (
	"github.com/codepitaka/pitaka-web/src/routes"
)

func main() {
	router := routes.SetupRouter()
	if err := router.Run(); err != nil {
		panic(err.Error())
    }
}