package utils

import (
	// "fmt"
	// "image/color"

	"engo.io/ecs"
	// "engo.io/engo"
	"engo.io/engo/common"
)

type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (hud *HUD) Init() {

}
