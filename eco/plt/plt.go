package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"
	"code.google.com/p/go-eco/eco/aux"
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/vg"
)

func usage() {
        fmt.Fprintf(os.Stderr, "usage: plt [-t title] -x [x-label] [-y y-label] [-p palette] [-o outfile]  [datafile]")
        os.Exit(2)
}


// getPoints returns x, y points from the specified row of the data matrix.
func getPoints(m *aux.Matrix, row int) plotter.XYs {
	pts := make(plotter.XYs, m.C)
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = m.Get(row, i)
	}
	return pts
}

func main() {
	var (
		inFile *os.File
		err error
		lineData plotter.XYs
		l        *plotter.Line
		palette color.Palette
	)

	help := flag.Bool("h", false, "show usage message")
	outFile := flag.String("o", "plot.svg", "output file")
	title := flag.String("t", "", "plot title")
	xLabel := flag.String("x", "", "label of the X axis")
	yLabel := flag.String("y", "", "label of the Y ayis")
	pal := flag.String("p", "W3CColors", "color palette to be used")

	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

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

	// read data
	mtx := aux.ReadCsvMatrix(inFile)

	// select palette
	switch *pal {
		case "W3C":
			palette = W3C
		case "Gold":
			palette = Gold
		case "Hilite":
			palette = Hilite
		case "Inkscape":
			palette = Inkscape
		default :
			palette = W3C
	}
	numColors := len(palette)

	// Create a new plot, set its title and axis labels.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = *title
	p.X.Label.Text = *xLabel
	p.Y.Label.Text = *yLabel

	for row := 0; row < mtx.R; row++ {
		n := row%numColors

		// Make a line plotter and set its style.
		lineData = getPoints(mtx, n)
		l = plotter.NewLine(lineData)
		l.LineStyle.Width = vg.Points(1)
		l.LineStyle.Color = palette[n]

		// Add the plotter to the plot
		p.Add(l)

	}

	// Save the plot to a file.
	if err := p.Save(4, 4, *outFile); err != nil {
		panic(err)
	}
}
