package GameObjects

import (
	"github.com/brentechols34/masterblaster/physics"
)

type Player struct {
	physics.Base
	Name   string
	Guns   []Gun
	Health float32
}
