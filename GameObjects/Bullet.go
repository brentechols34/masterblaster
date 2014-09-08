package GameObjects

import (
	"github.com/brentechols34/masterblaster/physics"
)

type Bullet struct {
	physics.Base
	Damage float32
}

func NewMachineGunBulletType() Bullet {
	return Bullet{
		Damage: 5,
	}
}
