package main

import (
	"graph"
	"lab/funcs"
)

func main() {
	x := graph.LinearArray(-10, 10, 201)
	freqX := graph.LinearArray(0, 5, 201/2)

	sincY := funcs.Sinc(x)
	triangY := funcs.Triang(x)
	rectY := funcs.Rect(x)
	combY := funcs.Comb(x)
	diracY := funcs.Dirac(x)
	sgnY := funcs.Sgn(x)
	sinPiY := funcs.SinPi(x)
	cosPiY := funcs.CosPi(x)

	arrs := [][]float64{diracY, combY, rectY, triangY, sgnY, sincY, sinPiY, cosPiY}
	toRun := []bool{true, true, true, true, true, true, true, true}
	names := []string{"dirac", "comb", "rect", "triang", "sgn", "sinc", "sinPi", "cosPi"}

	ls := graph.NewLS()
	ls.Solid()

	ftls := graph.NewLS()
	ftls.Pillars()

	for i, arr := range arrs {
		if toRun[i] {
			g := graph.NewGraph()
			g.SetLineStyle(ls)

			g.Plot(x, arr)

			if err := g.Draw(); err != nil {
				panic(err)
			}

			if err := g.SavePNG("images/"+names[i]+".png", true); err != nil {
				panic(err)
			}

			g.Clear()

			// FT

			g.SetLineStyle(ftls)

			res := funcs.FourierTransform(arr)

			g.Plot(freqX, res)

			if err := g.Draw(); err != nil {
				panic(err)
			}

			if err := g.SavePNG("images/ft/"+names[i]+".png", true); err != nil {
				panic(err)
			}
			g.Clear()
		}
	}
}
