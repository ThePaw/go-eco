// Main

package main

import (
	"flag"
)

func main() {
	help := flag.Bool("h", false, "Coenocline modeller\nUsage: coenocline -m [0|1|2] [-rsxya]")
	srfModel := flag.Int("m", 0, "Response function: 0 = Gaussian, 1 = Beta, 2 = Triangular")
	spacing := flag.Int("r", false, "true if spacing of samples along the gradient is random")
	nSpec := flag.Int("x", 20, "number of species")
	nSamp := flag.Int("y", 30, "number of samples")
	abuModel := flag.Int("abuModel", 0, "model of distribution of abundances")
	aa := flag.Float64("aa", 0, "scale param of distribution of abundances")
	ba := flag.Float64("ba", 0, "shape param of distribution of abundances")
	tolModel := flag.Int("tolModel", 0, "model of distribution of tolerances")
	at := flag.Float64("at", 0, "scale param of distribution of tolerances")
	bt := flag.Float64("bt", 0, "shape param of distribution of tolerances")
	//	seed := flag.Int64("s", 0, "random number seed")
	abumax := flag.Float64("a", 100.0, "maximum abundance")
	tmax := flag.Float64("tmax", 0.3, "maximum t")
	alphamax := flag.Float64("alpha", 5.5, "maximum alpha")
	gammamax := flag.Float64("gamma", 5.5, "maximum gamma")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
	} else {
		mtx := Coenocline(*nSpec, *nSamp, *srfModel, *abuModel, *aa, *ba, *tolModel, *at, *bt, *spacing, *abumax, *tmax, *alphamax, *gammamax)
		for i := 0; i < *nSamp; i++ {
			for j := 0; j < *nSpec; j++ {
				fmt.Print(mtx.Get(i, j), ",")
			}
			fmt.Println()
		}
	}
}
