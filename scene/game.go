package scene

import (
	"math/rand"
	"skharv/ecslife/component"
	"skharv/ecslife/entity"
	"skharv/ecslife/helper/enum"
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
		component.Health{},
		component.Type{},
		component.Spawn{},
	)

	w.AddSystems(
		system.NewRender(),
		system.NewRaycast(),
		system.NewThink(),
		system.NewDecay(),
		system.NewEat(),
		system.NewSpawn(),
	)

	w.AddEntities(
		&entity.Life{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Rotation: component.NewRotation(45),
			Radius:   component.NewRadius(10),
			Speed:    component.NewSpeed(2, 0),
			Vision:   component.NewVision(150),
			Color:    component.NewColor(128, 0, 128, 255),
			Facing:   component.NewFacing(nil, 0),
			Net:      component.NewNet(6, 128, 2, 2, source),
			Health:   component.NewHealth(100),
			Type:     component.NewType(enum.TypeLife),
		},
		&entity.Life{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Rotation: component.NewRotation(45),
			Radius:   component.NewRadius(10),
			Speed:    component.NewSpeed(2, 0),
			Vision:   component.NewVision(150),
			Color:    component.NewColor(255, 0, 0, 255),
			Facing:   component.NewFacing(nil, 0),
			Net:      component.NewNet(6, 64, 2, 4, source),
			Health:   component.NewHealth(100),
			Type:     component.NewType(enum.TypeLife),
		},
		&entity.Life{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Rotation: component.NewRotation(45),
			Radius:   component.NewRadius(10),
			Speed:    component.NewSpeed(2, 0),
			Vision:   component.NewVision(150),
			Color:    component.NewColor(255, 128, 128, 255),
			Facing:   component.NewFacing(nil, 0),
			Net:      component.NewNet(6, 4112, 2, 4, source),
			Health:   component.NewHealth(100),
			Type:     component.NewType(enum.TypeLife),
		},
		&entity.Life{
			Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
			Rotation: component.NewRotation(45),
			Radius:   component.NewRadius(10),
			Speed:    component.NewSpeed(2, 0),
			Vision:   component.NewVision(150),
			Color:    component.NewColor(128, 128, 128, 255),
			Facing:   component.NewFacing(nil, 0),
			Net:      component.NewNet(6, 256, 2, 4, source),
			Health:   component.NewHealth(100),
			Type:     component.NewType(enum.TypeLife),
		},
		&entity.FoodSpawner{
			Spawn: component.NewSpawn(75),
			Type:  component.NewType(enum.TypeFood),
		},
	)
}
