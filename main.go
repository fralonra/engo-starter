package main

import (
	"engo.io/engo"
	"fralonra/engo-starter/scenes"
)

const (
	Title  = "My Game"
	Width  = 800
	Height = 600
)

func main() {
	engo.RegisterScene(scenes.MainMenu)
	engo.RegisterScene(scenes.Game)

	opts := engo.RunOptions{
		Title:          Title,
		Width:          Width,
		Height:         Height,
		StandardInputs: true,
	}

	engo.Run(opts, scenes.Splash)
}
