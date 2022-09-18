package component

import (
	"math/rand"
)

type neuron struct {
	Value   float64
	Bias    float64
	Weights []float64
}

type Net struct {
	InputNeurons  int
	HiddenNeurons int
	OutputNeurons int
	HiddenLayers  int

	Layers [][]neuron
}

func NewNet(inputNeurons, hiddenNeurons, outputNeurons, hiddenLayers int, source rand.Source) Net {
	rand := rand.New(source)

	var layers [][]neuron

	for i := 0; i < hiddenLayers+2; i++ {
		var layer []neuron

		switch i {
		case 0:
			for j := 0; j < inputNeurons; j++ {
				layer = append(layer, neuron{
					Value:   0,
					Bias:    0,
					Weights: nil,
				})
			}
		case hiddenLayers + 1:
			for j := 0; j < outputNeurons; j++ {
				r := float64(rand.Intn(10000))
				r -= 5000
				r /= 1000

				var weights []float64
				for j := 0; j < len(layers[i-1]); j++ {
					w := float64(rand.Intn(10000))
					w -= 5000
					w /= 1000
					weights = append(weights, w)
				}
				layer = append(layer, neuron{
					Value:   0,
					Bias:    r,
					Weights: weights,
				})
			}
		default:
			for j := 0; j < hiddenNeurons; j++ {
				r := float64(rand.Intn(10000))
				r -= 5000
				r /= 1000

				var weights []float64
				for j := 0; j < len(layers[i-1]); j++ {
					w := float64(rand.Intn(10000))
					w -= 5000
					w /= 1000
					weights = append(weights, w)
				}
				layer = append(layer, neuron{
					Value:   0,
					Bias:    r,
					Weights: weights,
				})
			}
		}

		layers = append(layers, layer)
	}

	n := Net{
		InputNeurons:  inputNeurons,
		HiddenNeurons: hiddenNeurons,
		OutputNeurons: outputNeurons,
		HiddenLayers:  hiddenLayers,
		Layers:        layers,
	}

	return n
}
