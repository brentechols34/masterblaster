package GameObjects

type Player struct {
	GameObject

	Name string
	Guns []Gun
	Health float32
}