// Coenocline modeller: Gaussian SRF (species population density response function). 
// Models species population densities along (environmental) gradient sampled at distinct points. 
// Density model is Gaussian (Normal) PDF, where, for every species,  μ is its optimum on the gradient, 
// σ is a measure of tolerance,  99.7% of population lies within μ±3*σ. 
// μ and σ themselves, and total population size have meta-models: uniform | Gaussian | Beta. 
// Uniform and Beta distributions are reparametrized to use mean and standard deviation as parameters, for consistency.
// Influences external to the model are modelled as Gaussian "noise".
// Sample spacing model is uniform | exponential (Poisson process). 
// Gradient spans from 0 to 1. 
package main


// To do: -p 2 does not work: panic: σ too big, α, β out of range

import (
	. "code.google.com/p/go-eco/eco/cc"
	"flag"
	"fmt"
	"math/rand"
)

func main() {

	var mod Models

	help := flag.Bool("h", false, "print usage")
	srfModel := flag.String("srf", "gauss", "type of Species Response Function: gauss | beta")
	nSpec := flag.Int("n", 20, "number of species")
	nSamp := flag.Int("m", 200, "number of samples")
	sampModel := flag.Int("s", 0, "sampling model: 0 - regular, 1 - uniform, 2 - Poisson")
	optModel := flag.Int("o", 0, "optima model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	μOpt := flag.Float64("om", 0.5, "optimum mean")
	σOpt := flag.Float64("os", 0.35, "optimum standard deviation")
	popModel := flag.Int("p", 0, "population size model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	μPop := flag.Float64("pm", 10.0, "mean population size")
	σPop := flag.Float64("ps", 0.2, "standard deviation of population size, relative to population size = 1")
	tolModel := flag.Int("t", 1, "tolerance model: 0 - flat, 1 - Gaussian, 2 - Beta...")
	μTol := flag.Float64("tm", 0.2, "mean tolerance")
	σTol := flag.Float64("ts", 0.4, "standard deviation of tolerance, relative to tolerance = 1")
	ρPopTol := flag.Float64("r", 0.0, "correlation between ")
	σNoise := flag.Float64("e", 0.0, "outer noise as standard deviation")
	seed := flag.Int64("z", 1, "seed of random number generator")

	flag.Parse()
	rand.Seed(*seed)

	if *help {
		flag.PrintDefaults()
		return
	}

	// set up models
	mp:= &mod	// pointer to models
	mp.SetUpModels(*srfModel, *sampModel, *optModel, *popModel, *tolModel, *μOpt, *σOpt, *μPop, *σPop, *μTol, *σTol, *ρPopTol, *σNoise)

	// compute the coenocline matrix
	mtx := Coenocline(*nSpec, *nSamp, mod)

	// and write it out as CSV, transposed so that rows are species, columns are sampling points (to be reimplemented using csv.WriteAll)
	for i := 0; i < *nSpec; i++ {
		for j := 0; j < *nSamp; j++ {
			if j == 0 {
				fmt.Print(mtx.Get(j, i))
			} else {
				fmt.Print(",", mtx.Get(j, i))
			}
		}
		fmt.Println()
	}
}
