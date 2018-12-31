package utils

import (
	"engo.io/ecs"
)

type System struct{}

func (*System) Remove(ecs.BasicEntity) {}

func (*System) Update(dt float32) {}
