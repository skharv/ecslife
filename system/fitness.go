package system

import (
	"skharv/ecslife/component"

	"github.com/sedyh/mizu/pkg/engine"
)

type Fitness struct {
	*component.Fitness
}

func NewFitness() *Fitness {
	return &Fitness{}
}

func (f *Fitness) Update(w engine.World) {
	f.Fitness.Lifetime++
	f.Fitness.TimeSinceEat++
	f.Fitness.Score = f.Fitness.Lifetime - f.Fitness.TimeSinceEat
}
