package GameObjects

type Bullet struct {
	Base
	Damage float32
}

func NewMachineGunBulletType() Bullet {
	return Bullet{
		Damage: 5,
	}
}
