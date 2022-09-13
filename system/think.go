package system

import (
	"math"
	"skharv/ecslife/component"

	"github.com/sedyh/mizu/pkg/engine"
)

type Think struct {
	*component.Position
	*component.Vision
	*component.Rotation
	*component.Facing
	*component.Speed
}

func NewThink() *Think {
	return &Think{}
}

func (t *Think) Update(w engine.World) {
	if t.Facing.Target == nil {
		t.Rotation.R += 0.5
		t.Speed.S = 0.0
	} else {
		t.Speed.S = 1.0
	}

	t.Position.Y += math.Cos(t.Rotation.R*math.Pi/180) * t.Speed.S
	t.Position.X += math.Sin(t.Rotation.R*math.Pi/180) * t.Speed.S
}
