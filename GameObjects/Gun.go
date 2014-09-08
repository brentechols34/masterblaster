package GameObjects

type Gun struct {
	Id uint32

	AmmoType Bullet 
	RateOfFire float32
	ClipSize uint16
	CurrentAmmo uint16
	MaxAmmo uint16
}

func NewMachineGun (id uint32) Gun {
	return Gun {
		Id: id,
		AmmoType: NewMachineGunBulletType(),
		RateOfFire: 15,
		ClipSize: 30,
		CurrentAmmo: 30,
		MaxAmmo: 150,
	}
}