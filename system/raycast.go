package system

import (
	"image/color"
	"math"
	"skharv/ecslife/component"
	"skharv/ecslife/helper/enum"
	"skharv/ecslife/helper/globals"
	"skharv/ecslife/helper/num"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
)

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
		component.Type{},
	).Filter()

	var target engine.Entity
	var targetType enum.Type
	var distance = r.Vision.V

	for _, e := range targets {
		var pos *component.Position
		var rad *component.Radius
		var col *component.Color
		var typ *component.Type
		e.Get(&pos, &rad, &col, &typ)

		if r.Position == pos {
			continue
		}

		if num.Dist(r.Position.X, r.Position.Y, pos.X, pos.Y)+rad.R > r.Vision.V {
			continue
		}

		for i := 0; i < int(r.Vision.V); i++ {
			vpX := r.Position.X + (math.Sin(r.Rotation.R*math.Pi/180) * float64(i))
			vpY := r.Position.Y + (math.Cos(r.Rotation.R*math.Pi/180) * float64(i))

			d := num.Dist(vpX, vpY, pos.X, pos.Y)

			if vpX < 0 || vpX > globals.ScreenWidth {
				distance = d
				target = nil
				targetType = enum.TypeWall
				break
			}
			if vpY < 0 || vpY > globals.ScreenHeight {
				distance = d
				target = nil
				targetType = enum.TypeWall
				break
			}

			if d > rad.R {
				continue
			} else {
				if d < distance {
					distance = d
					target = e
					targetType = typ.T
				}
			}
		}
	}

	r.Facing.Target = target
	r.Facing.Type = targetType
	r.Facing.Distance = distance
}

func (r *Raycast) Draw(w engine.World, screen *ebiten.Image) {
	if globals.Debug {
		vpX := r.Position.X + (math.Sin(r.Rotation.R*math.Pi/180) * r.Vision.V)
		vpY := r.Position.Y + (math.Cos(r.Rotation.R*math.Pi/180) * r.Vision.V)

		ebitenutil.DrawLine(screen, r.Position.X, r.Position.Y, vpX, vpY, color.White)
	}
}
