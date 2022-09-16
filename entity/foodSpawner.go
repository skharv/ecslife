package entity

import "skharv/ecslife/component"

type FoodSpawner struct {
	component.Spawn
	component.Type
}
