package system

import (
	"skharv/ecslife/component"

	"github.com/sedyh/mizu/pkg/engine"
)

type Decay struct {
}

func NewDecay() *Decay {
	return &Decay{}
}

func (d *Decay) Update(w engine.World) {
	ents := w.View(
		component.Health{},
	).Filter()

	for _, e := range ents {
		var hea *component.Health
		e.Get(&hea)

		if hea.H <= 0 {
			w.RemoveEntity(e)
		} else {
			hea.H -= 0.05
		}
	}
}
