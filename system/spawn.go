package system

import (
	"math/rand"
	"skharv/ecslife/component"
	"skharv/ecslife/entity"
	"skharv/ecslife/helper/enum"
	"skharv/ecslife/helper/globals"

	"github.com/sedyh/mizu/pkg/engine"
)

type Spawn struct {
}

func NewSpawn() *Spawn {
	return &Spawn{}
}

func (s *Spawn) Update(w engine.World) {
	ent, found := w.View(
		component.Spawn{},
		component.Type{},
	).Get()
	if !found {
		return
	}

	var spa *component.Spawn
	var typ *component.Type

	ent.Get(&spa, &typ)

	for i := 0; i < spa.Quantity; i++ {
		if typ.T == enum.TypeFood {
			w.AddEntities(&entity.Food{
				Position: component.NewPosition(float64(rand.Intn(globals.ScreenWidth)), float64(rand.Intn(globals.ScreenHeight))),
				Radius:   component.NewRadius(5),
				Color:    component.NewColor(0, 128, 0, 255),
				Type:     component.NewType(enum.TypeFood),
			})
		}
	}

	w.RemoveEntity(ent)
}
