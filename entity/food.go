package entity

import "skharv/ecslife/component"

type Food struct {
	component.Position
	component.Radius
	component.Color
}
