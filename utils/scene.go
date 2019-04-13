package utils

import (
	"github.com/EngoEngine/engo"
	"log"
)

type Scene struct {
	Name string
}

func (s *Scene) Type() string {
	return s.Name
}

func (*Scene) Preload() {}

func (*Scene) Setup(engo.Updater) {}

func (s *Scene) Hide() {
	log.Println(s.Name + " is now hidden")
}

func (s *Scene) Show() {
	log.Println(s.Name + " is now shown")
}
