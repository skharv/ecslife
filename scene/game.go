package scene

import (
	"skharv/ecslife/component"
	"skharv/ecslife/entity"
	"skharv/ecslife/helper/globals"
	"skharv/ecslife/system"

	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Position{},
		component.Rotation{},
		component.Facing{},
		component.Speed{},
		component.Vision{},
		component.Radius{},
		component.Color{},
	)

	w.AddSystems(
		system.NewRender(),
		system.NewRaycast(),
		system.NewThink(),
	)

	w.AddEntities(
		&entity.Life{
			Position: component.NewPosition(globals.ScreenWidth/2, globals.ScreenHeight/2),
			Rotation: component.NewRotation(45),
			Radius:   component.NewRadius(10),
			Speed:    component.NewSpeed(0),
			Vision:   component.NewVision(150),
			Color:    component.NewColor(128, 0, 128, 255),
			Facing:   component.NewFacing(nil),
		},
		&entity.Food{
			Position: component.NewPosition(globals.ScreenWidth/2-50, globals.ScreenHeight/2),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
	)
}
