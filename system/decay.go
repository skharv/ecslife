package system

import (
	"fmt"
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
		component.Fitness{},
		component.Color{},
	).Filter()

	for _, e := range ents {
		var hea *component.Health
		var fit *component.Fitness
		var col *component.Color
		e.Get(&hea, &fit, &col)

		if hea.H <= 0 {
			fmt.Println(col.C.RGBA())
			fmt.Println(fit)
			w.RemoveEntity(e)
		} else {
			hea.H -= 0.05
		}
	}
}
