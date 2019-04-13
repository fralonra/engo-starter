package utils

import (
	"github.com/EngoEngine/ecs"
)

type System struct{}

func (*System) Remove(ecs.BasicEntity) {}

func (*System) Update(dt float32) {}
