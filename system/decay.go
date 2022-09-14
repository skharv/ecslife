package system

import (
	"skharv/ecslife/component"

	"github.com/sedyh/mizu/pkg/engine"
)

type Decay struct {
	*component.Position
	*component.Radius
	*component.Facing
	*component.Health
}

func NewDecay() *Decay {
	return &Decay{}
}

func (d *Decay) Update(w engine.World) {
	if d.Health.H <= 0 {
		// me := w.View(d.Health).Filter()
		// for _, e := range me {
		// 	w.RemoveEntity(e)
		// 	break
		// }
	} else {
		d.Health.H -= 0.5
	}
}
