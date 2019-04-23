package scenes

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/fralonra/engo-utils"
)

type GameScene struct {
	utils.Scene
}

func (*GameScene) Preload() {}

func (*GameScene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})
	w.AddSystem(&GameSystem{})
	w.AddSystem(&utils.ClickableSystem{})

	label1 := utils.Label{
		World: w,
		Font:  lgFont,
		Text:  "Game",
		Position: engo.Point{
			X: 300,
			Y: 240,
		},
	}
	label1.Init()

	button := utils.Button{
		World: w,
		Font:  smFont,
		Text:  "Go back",
		Position: engo.Point{
			X: 100,
			Y: 500,
		},
	}
	button.Init()
	button.OnClick(func() {
		engo.SetSceneByName(sceneMainMenu, false)
	})
}

type GameSystem struct {
	utils.System
}

func (*GameSystem) Update(dt float32) {}
