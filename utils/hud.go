package utils

import (
	// "fmt"
	// "image/color"

	"github.com/EngoEngine/ecs"
	// "github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (hud *HUD) Init() {

}
