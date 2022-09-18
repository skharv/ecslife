package component

import (
	"skharv/ecslife/helper/enum"

	"github.com/sedyh/mizu/pkg/engine"
)

type Facing struct {
	Target   engine.Entity
	Type     enum.Type
	Distance float64
}

func NewFacing(target engine.Entity, targetType enum.Type) Facing {
	return Facing{target, targetType, -1}
}
