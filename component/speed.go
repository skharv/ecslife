package component

type Speed struct {
	B, S float64
}

func NewSpeed(b, s float64) Speed {
	return Speed{b, s}
}
