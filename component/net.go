package component

import (
	"math/rand"
)

type Net struct {
	InputNeurons  int
	HiddenNeurons int
	OutputNeurons int

	InputValues  []float64
	HiddenValues [][]float64
	//HiddenValues []float64
	OutputValues []float64

	HiddenBiases []float64
	OutputBiases []float64

	//ItoHweights []float64
	HtoHweights [][]float64
	HtoOweights []float64

	HiddenLayers [][]float64
}

func NewNet(inputs, hiddens, outputs, layers int, source rand.Source) Net {
	rand := rand.New(source)

	var iv []float64
	for i := 0; i < inputs; i++ {
		iv = append(iv, 0.0)
	}

	var hb []float64
	var hv [][]float64
	for i := 0; i < layers; i++ {
		for j := 0; j < hiddens; j++ {
			r := float64(rand.Intn(10000))
			r -= 5000
			r /= 1000
			hb = append(hb, r)
			hv[i] = append(hv[i], r)
		}
	}

	var hw [][]float64
	for i := 0; i < layers; i++ {
		hw = append(hv, []float64{})
		switch i {
		case 0:
			for j := 0; j < inputs*hiddens; j++ {
				r := float64(rand.Intn(10000))
				r -= 5000
				r /= 1000
				hw[i] = append(hw[i], r)
			}
		case layers - 1:
			continue
		default:
			for j := 0; j < hiddens*hiddens; j++ {
				r := float64(rand.Intn(10000))
				r -= 5000
				r /= 1000
				hw[i] = append(hw[i], r)
			}
		}
	}

	// var ihw []float64
	// for i := 0; i < inputs*hiddens; i++ {
	// 	r := float64(rand.Intn(10000))
	// 	r -= 5000
	// 	r /= 1000
	// 	ihw = append(ihw, r)
	// }

	// var hb []float64
	// var hv []float64
	// for i := 0; i < hiddens; i++ {
	// 	r := float64(rand.Intn(10000))
	// 	r -= 5000
	// 	r /= 1000
	// 	hb = append(hb, r)
	// 	hv = append(hv, 0.0)
	// }

	// var ihw []float64
	// for i := 0; i < inputs*hiddens; i++ {
	// 	r := float64(rand.Intn(10000))
	// 	r -= 5000
	// 	r /= 1000
	// 	ihw = append(ihw, r)
	// }

	var ob []float64
	var ov []float64
	for i := 0; i < outputs; i++ {
		r := float64(rand.Intn(10000))
		r -= 5000
		r /= 1000
		ob = append(ob, r)
		ov = append(ov, 0.0)
	}

	var how []float64
	for i := 0; i < hiddens*outputs; i++ {
		r := float64(rand.Intn(10000))
		r -= 5000
		r /= 1000
		how = append(how, r)
	}

	n := Net{
		InputNeurons:  inputs,
		HiddenNeurons: hiddens,
		OutputNeurons: outputs,
		InputValues:   iv,
		HiddenValues:  hv,
		OutputValues:  ov,
		HiddenBiases:  hb,
		OutputBiases:  ob,
		HtoOweights:   how,
		//HiddenValues:  hv,
		//ItoHweights: ihw,
	}

	return n
}
