package GameObjects

type Bullet struct {
	GameObject

	Damage float32
}

func NewMachineGunBulletType() Bullet {
	return Bullet{
		Damage: 5,
	}
}
