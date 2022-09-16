package system

import (
	"math"
	"skharv/ecslife/component"
	"skharv/ecslife/helper/enum"
	"skharv/ecslife/helper/globals"
	"skharv/ecslife/helper/num"
	"time"

	"github.com/sedyh/mizu/pkg/engine"
)

type Think struct {
	*component.Position
	*component.Vision
	*component.Rotation
	*component.Facing
	*component.Speed
	*component.Health
	*component.Net
}

func NewThink() *Think {
	return &Think{}
}

func (t *Think) Update(w engine.World) {
	//Set inputs
	if t.Facing.Target == nil && t.Facing.Type != enum.TypeWall {
		t.Net.InputValues[0] = -1.0
	} else {
		t.Net.InputValues[0] = 1.0
	}

	t.Net.InputValues[1] = t.Speed.S

	t.Net.InputValues[2] = math.Sin(float64(time.Now().UnixMilli()))

	t.Net.InputValues[3] = (t.Position.X/globals.ScreenWidth)*2 - 1
	t.Net.InputValues[4] = (t.Position.Y/globals.ScreenHeight)*2 - 1

	t.Net.InputValues[5] = (t.Health.H/t.Health.Max)*2 - 1

	//Think
	itohcount := 0
	for i := 0; i < len(t.Net.HiddenValues[0]); i++ {
		var value float64
		for j := 0; j < len(t.Net.InputValues); j++ {
			value += (t.Net.InputValues[j] * t.Net.ItoHweights[itohcount])
			itohcount++
		}
		value += t.Net.HiddenBiases[i]
		t.Net.HiddenValues[i] = num.Sigmoid(value)
	}

	htohcount := 0
	for i := 0; i < len(t.Net.HiddenValues); i++ {
		var value float64
		for j := 0; j < len(t.Net.InputValues); j++ {
			value += (t.Net.InputValues[j] * t.Net.ItoHweights[itohcount])
			itohcount++
		}
		value += t.Net.HiddenBiases[i]
		t.Net.HiddenValues[i] = num.Sigmoid(value)
	}

	htoocount := 0
	for i := 0; i < len(t.Net.OutputValues); i++ {
		var value float64
		for j := 0; j < len(t.Net.HiddenValues); j++ {
			value += (t.Net.HiddenValues[j] * t.Net.HtoOweights[htoocount])
			htoocount++
		}
		value += t.Net.OutputBiases[i]
		t.Net.OutputValues[i] = num.Sigmoid(value)
	}

	//Act
	t.Rotation.R += t.Net.OutputValues[0]*2 - 1
	t.Speed.S = t.Net.OutputValues[1] * t.Speed.B

	t.Position.Y += math.Cos(t.Rotation.R*math.Pi/180) * t.Speed.S
	t.Position.Y = num.Clamp(t.Position.Y, 0, globals.ScreenHeight)
	t.Position.X += math.Sin(t.Rotation.R*math.Pi/180) * t.Speed.S
	t.Position.X = num.Clamp(t.Position.X, 0, globals.ScreenWidth)
}
