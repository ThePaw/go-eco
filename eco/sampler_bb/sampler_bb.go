// Coenoses sampler: Beta-Binomial sampling model.
// Performs Beta-Binomial sampling of species populations given population densities, sample size and an array of overdispersion parameters for every species.
package main

// Thanks to Jari Oksanen: betasimu.c
// To do: Overdispersion parameter generator to be implemented. 

import (
	"code.google.com/p/go-eco/eco/aux"
	"code.google.com/p/probab/dst"
	"flag"
	"fmt"
	"os"
)

// Estimates the parameters a,b of beta distribution from expected proportion (pi), binomial denominator (m), and shape parameter (tau2). 
func betapara(pi, m, tau2 float64) (a, b float64) {
	// Solution (hopefully correct) of Exercise 4.17 of McCullagh & Nelder 1989, helped by Moore, Appl Stat 36, 8-14; 1987.
	t1 := tau2 * m
	t2 := t1 - m - tau2 + 1
	t3 := 1 / (1 + t1 - tau2)
	t4 := t2 * t3
	a = -t4 * pi
	b = t4 * (pi - 1)
	return
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: sampler_bb [-n sample_size]  [datafile.csv]")
	os.Exit(2)
}

func main() {
	var (
		inFile *os.File
		err    error
	)

	help := flag.Bool("h", false, "show usage message")
	sSize := flag.Int64("n", 100, "number of individuals per sample")
	flag.Usage = usage
	flag.Parse()

	// from where to input
	switch flag.NArg() {
	case 0:
		inFile = os.Stdin
	case 1:
		inFile, err = os.Open(flag.Arg(0))
	default:
		flag.Usage()
		os.Exit(1)
	}

	if *help {
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		flag.Usage()
		os.Exit(1)
	}

	// read data, species are rows, samples are cols! 
	mtx := aux.ReadCsvMatrix(inFile)
	nSamp := mtx.C
	nSpec := mtx.R
	out := aux.NewMatrix(nSpec, nSamp)

	// for every sample
	for i := 0; i < nSamp; i++ {
		// recalculate species counts to proportions
		θ := make([]float64, nSpec)
		sum := 0.0
		for j := 0; j < nSpec; j++ {
			sum += mtx.Get(j, i)
		}
		for j := 0; j < nSpec; j++ {
			θ[j] = mtx.Get(j, i) / sum
		}
		y := dst.MultinomialNext(θ, *sSize)
		for j := 0; j < nSpec; j++ {
			shape := 0.01 // Overdispersion parameter: constant for now; needs to be properly modelled for every species separately as shape[j]
			mu := float64(y[j])
			resp := 0.0
			if mu > 0.0 {
				a, b := betapara(mu, float64(*sSize), shape)
				mu = dst.BetaNext(a, b)
				resp = float64(dst.BinomialNext(*sSize, mu))
			}
			out.Set(j, i, resp)
		}
	}
	for i := 0; i < nSpec; i++ {
		for j := 0; j < nSamp; j++ {
			if j == 0 {
				fmt.Print(out.Get(i, j))
			} else {
				fmt.Print(",", out.Get(i, j))
			}
		}
		fmt.Println()
	}
}
