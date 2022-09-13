package system

import (
	"image/color"
	"math"
	"skharv/ecslife/component"
	"skharv/ecslife/helper/globals"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
)

func sq(x float64) float64 { return x * x }

func dist(X1, Y1, X2, Y2 float64) float64 {
	return math.Sqrt(sq(X1-X2) + sq(Y1-Y2))
}

type Raycast struct {
	*component.Position
	*component.Vision
	*component.Rotation
	*component.Facing
}

func NewRaycast() *Raycast {
	return &Raycast{}
}

func (r *Raycast) Update(w engine.World) {
	targets := w.View(
		component.Position{},
		component.Radius{},
		component.Color{},
	).Filter()

	smallestDist := r.Vision.V
	var target engine.Entity

	for _, e := range targets {
		var pos *component.Position
		var rad *component.Radius
		var col *component.Color
		e.Get(&pos, &rad, &col)

		if r.Position == pos {
			continue
		}

		if dist(r.Position.X, r.Position.Y, pos.X, pos.Y)+rad.R > r.Vision.V {
			continue
		}

		for i := 0; i < int(r.Vision.V); i++ {
			vpX := r.Position.X + (math.Sin(r.Rotation.R*math.Pi/180) * float64(i))
			vpY := r.Position.Y + (math.Cos(r.Rotation.R*math.Pi/180) * float64(i))

			d := dist(vpX, vpY, pos.X, pos.Y)

			if d > rad.R {
				continue
			} else {
				if d < smallestDist {
					smallestDist = d
					target = e
				}
			}
		}
	}

	r.Facing.Target = target
}

func (r *Raycast) Draw(w engine.World, screen *ebiten.Image) {
	if globals.Debug {
		vpX := r.Position.X + (math.Sin(r.Rotation.R*math.Pi/180) * r.Vision.V)
		vpY := r.Position.Y + (math.Cos(r.Rotation.R*math.Pi/180) * r.Vision.V)

		ebitenutil.DrawLine(screen, r.Position.X, r.Position.Y, vpX, vpY, color.White)
	}
}
