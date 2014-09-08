package GameObjects

type Player struct {
	Base
	Name   string
	Guns   []Gun
	Health float32
}
