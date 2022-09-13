package system

import (
	"math"
	"skharv/ecslife/component"
	"time"

	"github.com/sedyh/mizu/pkg/engine"
)

type Think struct {
	*component.Position
	*component.Vision
	*component.Rotation
	*component.Facing
	*component.Speed
	*component.Net
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func NewThink() *Think {
	return &Think{}
}

func (t *Think) Update(w engine.World) {
	//Set inputs
	if t.Facing.Target == nil {
		t.Net.InputValues[0] = -1.0
	} else {
		t.Net.InputValues[0] = 1.0
	}

	t.Net.InputValues[1] = t.Speed.S

	t.Net.InputValues[2] = math.Sin(float64(time.Now().UnixMilli()))

	//Think
	for i := 0; i < len(t.Net.InputValues); i++ {
		value := sigmoid(t.Net.InputBiases[i] + t.Net.InputValues[i])
		t.Net.InputValues[i] = value
	}

	itohcount := 0
	for i := 0; i < len(t.Net.HiddenValues); i++ {
		var value float64
		for j := 0; j < len(t.Net.InputValues); j++ {
			value += (t.Net.InputValues[j] * t.Net.ItoHweights[itohcount])
			itohcount++
		}
		value += t.Net.HiddenBiases[i]
		t.Net.HiddenValues[i] = sigmoid(value)
	}

	htoocount := 0
	for i := 0; i < len(t.Net.OutputValues); i++ {
		var value float64
		for j := 0; j < len(t.Net.HiddenValues); j++ {
			value += (t.Net.HiddenValues[j] * t.Net.HtoOweights[htoocount])
			htoocount++
		}
		value += t.Net.OutputBiases[i]
		t.Net.OutputValues[i] = sigmoid(value)
	}

	//Act
	t.Rotation.R += t.Net.OutputValues[0]
	t.Speed.S = t.Net.OutputValues[1] * t.Speed.B

	t.Position.Y += math.Cos(t.Rotation.R*math.Pi/180) * t.Speed.S
	t.Position.X += math.Sin(t.Rotation.R*math.Pi/180) * t.Speed.S
}
