package system

import (
	"image/color"
	"skharv/ecslife/component"
	"skharv/ecslife/helper/globals"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct {
	offscreen *ebiten.Image
}

func NewRender() *Render {
	r := Render{}

	r.offscreen = ebiten.NewImage(globals.ScreenWidth, globals.ScreenHeight)

	return &r
}

func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	screen.Fill(color.Black)

	renders := w.View(
		component.Position{},
		component.Radius{},
		component.Color{},
	).Filter()

	for _, e := range renders {
		var pos *component.Position
		var rad *component.Radius
		var col *component.Color
		e.Get(&pos, &rad, &col)

		ebitenutil.DrawCircle(r.offscreen, pos.X, pos.Y, rad.R, col.C)
	}

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(r.offscreen, op)
	r.offscreen.Clear()
}
