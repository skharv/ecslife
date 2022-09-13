package component

import "github.com/sedyh/mizu/pkg/engine"

type Facing struct {
	Target engine.Entity
}

func NewFacing(target engine.Entity) Facing {
	return Facing{target}
}
