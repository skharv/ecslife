package system

import (
	"skharv/ecslife/component"
	"skharv/ecslife/helper/enum"
	"skharv/ecslife/helper/num"

	"github.com/sedyh/mizu/pkg/engine"
)

type Eat struct {
	*component.Position
	*component.Radius
	*component.Facing
	*component.Health
}

func NewEat() *Eat {
	return &Eat{}
}

func (e *Eat) Update(w engine.World) {
	if e.Facing.Target != nil {
		var pos *component.Position
		var rad *component.Radius

		e.Facing.Target.Get(&pos, &rad)

		if num.Dist(e.Position.X, e.Position.Y, pos.X, pos.Y) < e.Radius.R+rad.R {
			w.RemoveEntity(e.Facing.Target)
			e.Facing.Target = nil
			e.Facing.Type = enum.TypeNone
			e.Health.H = e.Health.Max
		}
	}
}
