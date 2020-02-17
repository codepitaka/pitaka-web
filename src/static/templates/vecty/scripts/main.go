package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	// "honnef.co/go/js/dom"
	// "github.com/gopherjs/gopherjs/js"
)

func main() {
	vecty.SetTitle("Hello Vecty!")
	vecty.RenderInto("#vecty-content",&PageView{})
}

// PageView is our main page component.
type PageView struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Style("float", "left"),
			vecty.Style("background-color", "red"),
		),
		vecty.Text("Hello Vecty!"),
	)
}