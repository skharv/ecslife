package component

type Fitness struct {
	Lifetime     float64
	TimeSinceEat float64
	Score        float64
}

func NewFitness() Fitness {
	return Fitness{0, 0, 0}
}
