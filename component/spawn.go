package component

type Spawn struct {
	Quantity int
}

func NewSpawn(quantity int) Spawn {
	return Spawn{quantity}
}
