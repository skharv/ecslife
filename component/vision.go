package component

type Vision struct {
	V float64
}

func NewVision(v float64) Vision {
	return Vision{v}
}
