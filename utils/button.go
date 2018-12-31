package utils

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	eventClick     = "click"
	eventMouseover = "mouseover"
)

type Button struct {
	eventHandlers map[string]func()
	Font          *common.Font
	label         *Label
	Position      engo.Point
	Text          string
	World         *ecs.World

	ecs.BasicEntity
	common.MouseComponent
	common.RenderComponent
	common.SpaceComponent
}

func (b *Button) Init() {
	b.BasicEntity = ecs.NewBasic()

	b.eventHandlers = make(map[string]func())

	b.label = &Label{
		World:    b.World,
		Font:     b.Font,
		Text:     b.Text,
		Position: b.Position,
	}
	b.label.Init()

	b.MouseComponent = common.MouseComponent{}
	b.RenderComponent = b.label.RenderComponent
	b.SpaceComponent = b.label.SpaceComponent

	b.SetShader(common.HUDShader)

	for _, system := range b.World.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&b.BasicEntity, &b.MouseComponent, &b.SpaceComponent, &b.RenderComponent)
		case *ButtonSystem:
			sys.Add(b)
		}
	}
}

func (b *Button) On(event string, f func()) {
	b.eventHandlers[event] = f
}

func (b *Button) OnClick(f func()) {
	b.On(eventClick, f)
}

func (b *Button) OnMouseOver(f func()) {
	b.On(eventMouseover, f)
}

type buttonEntity struct {
	*Button
}

type ButtonSystem struct {
	entities []buttonEntity
}

func (s *ButtonSystem) Add(b *Button) {
	s.entities = append(s.entities, buttonEntity{b})
}

func (s *ButtonSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range s.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete+1:]...)
	}
}

func (s *ButtonSystem) Update(dt float32) {
	for _, e := range s.entities {
		if e.Clicked {
			pos := engo.Point{
				X: e.MouseX,
				Y: e.MouseY,
			}
			if e.Contains(pos) {
				if f, ok := e.eventHandlers[eventClick]; ok {
					f()
				}
			}
		}
	}
}
