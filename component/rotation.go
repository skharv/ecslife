package component

type Rotation struct {
	R float64
}

func NewRotation(r float64) Rotation {
	return Rotation{r}
}
