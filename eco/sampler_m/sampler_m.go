// Coenoses sampler: Multinomial sampling model.
// Performs multinomial sampling of species populations given population densities and sample size.
package main

import (
	"flag"
	"fmt"
	"os"
	"code.google.com/p/go-eco/eco/aux"
	. "code.google.com/p/probab/dst"
)

func usage() {
        fmt.Fprintf(os.Stderr, "usage: sampler_m [-n sample_size]  [datafile.csv]")
        os.Exit(2)
}

func main() {
	var (
		inFile *os.File
		err error
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
		// recalculate species to proportions
		θ := make([]float64, nSpec)
		sum := 0.0
		for j := 0; j < nSpec; j++ {
			sum += mtx.Get(j, i)
		}
		for j := 0; j < nSpec; j++ {
			θ[j] =  mtx.Get(j, i) / sum
		}
		y := MultinomialNext(θ, *sSize)
		// write it out
		for j := 0; j < nSpec; j++ {
			out.Set(j, i, float64(y[j]))
		}
	}
	for i := 0; i < nSpec; i++ {
		for j := 0; j < nSamp; j++ {
			if j == 0 {
				fmt.Print(out.Get(i,j))
			} else {
				fmt.Print(",", out.Get(i,j))
			}
		}
		fmt.Println()
	}
}
