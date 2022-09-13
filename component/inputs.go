package component

type Inputs struct {
	Neurons []float64
}

func NewInputs(quantity int) Inputs {
	var neurons []float64
	neurons = neurons[:quantity]
	return Inputs{neurons}
}
