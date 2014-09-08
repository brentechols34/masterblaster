package GameObjects

import (
	"github.com/brentechols34/masterblaster/physics"
)

type Wall struct {
	physics.Base
	Length uint16
	Width  uint16
}
