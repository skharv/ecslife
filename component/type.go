package component

import (
	"skharv/ecslife/helper/enum"
)

type Type struct {
	T enum.Type
}

func NewType(t enum.Type) Type {
	return Type{t}
}
