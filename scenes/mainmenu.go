package scenes

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fralonra/engo-demo/utils"
)

type MainMenuScene struct {
	utils.Scene
}

func (*MainMenuScene) Preload() {}

func (*MainMenuScene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})
	w.AddSystem(&MainMenuSystem{})
	w.AddSystem(&utils.ButtonSystem{})

	entriesText := []string{"New", "Quit"}
	entriesClicker := []func(){
		func() {
			engo.SetSceneByName(sceneGame, false)
		}, func() {
			engo.Exit()
		},
	}

	for index, text := range entriesText {
		entry := utils.Button{
			World: w,
			Font:  mdFont,
			Text:  text,
			Position: engo.Point{
				X: 300,
				Y: 240 + float32(index*100),
			},
		}
		entry.Init()
		entry.OnClick(entriesClicker[index])
	}
}

type MainMenuSystem struct {
	utils.System
}

func (s *MainMenuSystem) Update(dt float32) {}
