package utils

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Label struct {
	Font     *common.Font
	Position engo.Point
	Text     string
	World    *ecs.World

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (l *Label) Init() {
	l.BasicEntity = ecs.NewBasic()

	width, height, _ := l.Font.TextDimensions(l.Text)

	l.SpaceComponent = common.SpaceComponent{
		Width:  float32(width),
		Height: float32(height),
	}
	l.SpaceComponent.Position = l.Position
	l.RenderComponent.Drawable = common.Text{
		Font: l.Font,
		Text: l.Text,
	}

	l.SetShader(common.TextHUDShader)

	for _, system := range l.World.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&l.BasicEntity, &l.RenderComponent, &l.SpaceComponent)
		}
	}
}
