package system

import (
	"skharv/ecslife/component"
	"skharv/ecslife/helper/enum"
	"skharv/ecslife/helper/num"

	"github.com/sedyh/mizu/pkg/engine"
)

type Eat struct {
}

func NewEat() *Eat {
	return &Eat{}
}

func (e *Eat) Update(w engine.World) {
	ents := w.View(
		component.Position{},
		component.Radius{},
		component.Facing{},
		component.Health{},
	).Filter()

	for _, n := range ents {
		var pos *component.Position
		var rad *component.Radius
		var hea *component.Health
		var fac *component.Facing

		n.Get(&pos, &rad, &hea, &fac)

		if fac.Target != nil {
			var tPos *component.Position
			var tRad *component.Radius

			fac.Target.Get(&tPos, &tRad)

			if num.Dist(pos.X, pos.Y, tPos.X, tPos.Y) < rad.R+tRad.R {
				w.RemoveEntity(fac.Target)
				fac.Target = nil
				fac.Type = enum.TypeNone
				hea.H = hea.Max
			}
		}
	}
}
