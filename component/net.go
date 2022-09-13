package component

import "math/rand"

type Net struct {
	InputQuantity        int
	LayerQuantity        int
	InternalLayerDensity int
	OutputQuantity       int

	Inputs  []float64
	Layers  [][]float64
	Outputs []float64
}

func NewNet(inputs, layers, layerDensity, outputs int) Net {
	n := Net{
		InputQuantity:        inputs,
		LayerQuantity:        layers,
		InternalLayerDensity: layerDensity,
		OutputQuantity:       outputs,
	}

	n.Inputs = n.Inputs[:inputs]

	for i := 0; i < layers; i++ {
		for j := 0; j < layerDensity; j++ {
			v := float64(rand.Intn(800))
			v -= 400
			v /= 100
			n.Layers[i] = append(n.Layers[i], v)
		}
	}

	n.Outputs = n.Outputs[:outputs]

	return n
}
