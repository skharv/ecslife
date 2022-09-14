package component

type Health struct {
	Max float64
	H   float64
}

func NewHealth(h float64) Health {
	return Health{h, h}
}
