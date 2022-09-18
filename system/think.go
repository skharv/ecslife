package system

import (
	"math"
	"skharv/ecslife/component"
	"skharv/ecslife/helper/enum"
	"skharv/ecslife/helper/globals"
	"skharv/ecslife/helper/num"

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
	*component.Fitness
}

func NewThink() *Think {
	return &Think{}
}

func (t *Think) Update(w engine.World) {
	//Feel
	if t.Facing.Target == nil && t.Facing.Type != enum.TypeWall {
		t.Net.Layers[0][0].Value = -1.0
	} else {
		t.Net.Layers[0][0].Value = 1.0
	}
	t.Net.Layers[0][1].Value = t.Facing.Distance / t.Vision.V
	t.Net.Layers[0][2].Value = float64(t.Facing.Type - 2)
	t.Net.Layers[0][3].Value = t.Speed.S
	t.Net.Layers[0][4].Value = (t.Position.X/globals.ScreenWidth)*2 - 1
	t.Net.Layers[0][5].Value = (t.Position.Y/globals.ScreenHeight)*2 - 1
	t.Net.Layers[0][6].Value = (t.Health.H/t.Health.Max)*2 - 1
	t.Net.Layers[0][7].Value = t.Fitness.Lifetime
	t.Net.Layers[0][8].Value = t.Fitness.TimeSinceEat
	t.Net.Layers[0][9].Value = t.Fitness.Score

	//Think
	// For each neuron (j) in layer i
	// We need to get the weights (k) of j * i-1 value
	// then add j bias
	for i := 1; i < len(t.Net.Layers); i++ {
		for j := 0; j < len(t.Net.Layers[i]); j++ {
			var value float64
			for k := 0; k < len(t.Net.Layers[i-1]); k++ {
				value += (t.Net.Layers[i][j].Weights[k] * t.Net.Layers[i-1][k].Value)
			}
			value += t.Net.Layers[i][j].Bias
			t.Net.Layers[i][j].Value = num.Sigmoid(value)
		}
	}

	//Act
	if t.Net.Layers[len(t.Net.Layers)-1][0].Value > t.Net.Layers[len(t.Net.Layers)-1][1].Value {
		t.Rotation.R += t.Net.Layers[len(t.Net.Layers)-1][3].Value
	} else {
		t.Rotation.R -= t.Net.Layers[len(t.Net.Layers)-1][3].Value
	}
	t.Speed.S = t.Net.Layers[len(t.Net.Layers)-1][3].Value

	t.Position.Y += math.Cos(t.Rotation.R*math.Pi/180) * t.Speed.S
	t.Position.Y = num.Clamp(t.Position.Y, 0, globals.ScreenHeight)
	t.Position.X += math.Sin(t.Rotation.R*math.Pi/180) * t.Speed.S
	t.Position.X = num.Clamp(t.Position.X, 0, globals.ScreenWidth)
}
