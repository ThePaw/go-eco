// Coenoses sampler: Poisson
// Performs Poisson sampling of species populations given population densities and sample size.
package main

import (
	"flag"
	"fmt"
	"os"
	"code.google.com/p/go-eco/eco/aux"
	. "code.google.com/p/probab/dst"
)

func usage() {
        fmt.Fprintf(os.Stderr, "usage: sampler_m [-s sample_size]  [datafile.csv]")
        os.Exit(2)
}

func main() {
	var (
		inFile *os.File
		err error
	)

	help := flag.Bool("h", false, "show usage message")
	sSize := flag.Float64("s", 100, "sample size")
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

	for i := 0; i < nSamp; i++ {
		for j := 0; j < nSpec; j++ {
			y := *sSize * mtx.Get(j, i) 
			y = float64(PoissonNext(y))
			out.Set(j, i, y)
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
