package scene

import (
	"math/rand"
	"skharv/ecslife/component"
	"skharv/ecslife/entity"
	"skharv/ecslife/helper/globals"
	"skharv/ecslife/system"
	"time"

	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	source := rand.NewSource(time.Now().UnixNano())

	w.AddComponents(
		component.Position{},
		component.Rotation{},
		component.Facing{},
		component.Speed{},
		component.Vision{},
		component.Radius{},
		component.Color{},
		component.Net{},
	)

	w.AddSystems(
		system.NewRender(),
		system.NewRaycast(),
		system.NewThink(),
	)

	w.AddEntities(
		&entity.Life{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Rotation: component.NewRotation(45),
			Radius:   component.NewRadius(10),
			Speed:    component.NewSpeed(2, 0),
			Vision:   component.NewVision(150),
			Color:    component.NewColor(128, 0, 128, 255),
			Facing:   component.NewFacing(nil),
			Net:      component.NewNet(5, 10, 2, source),
		},
		&entity.Life{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Rotation: component.NewRotation(45),
			Radius:   component.NewRadius(10),
			Speed:    component.NewSpeed(2, 0),
			Vision:   component.NewVision(150),
			Color:    component.NewColor(255, 0, 0, 255),
			Facing:   component.NewFacing(nil),
			Net:      component.NewNet(5, 5, 2, source),
		},
		&entity.Life{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Rotation: component.NewRotation(45),
			Radius:   component.NewRadius(10),
			Speed:    component.NewSpeed(2, 0),
			Vision:   component.NewVision(150),
			Color:    component.NewColor(128, 255, 128, 255),
			Facing:   component.NewFacing(nil),
			Net:      component.NewNet(5, 32, 2, source),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
		&entity.Food{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Radius:   component.NewRadius(5),
			Color:    component.NewColor(0, 128, 0, 255),
		},
	)
}
