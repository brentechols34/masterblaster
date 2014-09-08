package physics

type Base struct {
	Id           uint32
	Position     Vect2
	Velocity     Vect2
	Acceleration Vect2
	Size         int16
}

type Vect2 struct {
	X float32
	Y float32
}

// Modifies Vector in place
func (v Vect2) Add(v2 Vect2) {
	v.X += v2.X
	v.Y += v2.Y
}

func MultVect2(v1 Vect2, val float32) Vect2 {
	return Vect2{X: v1.X * val, Y: v1.Y * val}
}
