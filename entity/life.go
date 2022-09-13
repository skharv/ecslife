package entity

import "skharv/ecslife/component"

type Life struct {
	component.Position
	component.Rotation
	component.Radius
	component.Speed
	component.Vision
	component.Color
	component.Facing
	component.Net
}
