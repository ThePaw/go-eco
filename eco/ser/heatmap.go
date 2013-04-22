// Heat Map

package ser

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
)

func HeatMap(mat Matrix64) {
	heatData := getData(mat)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	bs, err := plotter.NewHeatMap(heatData)
	if err != nil {
		panic(err)
	}
	bs.Colorer = plotter.GoldRamp3
	p.Add(bs)

	if err := p.Save(4, 4, "heatmap.svg"); err != nil {
		panic(err)
	}
}

// getData fetches a data matrix.
func getData(a Matrix64) plotter.XYZs {
	n := len(a)
	m := len(a[0])
	data := make(plotter.XYZs, n*m)
	k := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			data[k].X = float64(i)
			data[k].Y = float64(j)
			data[k].Z = a[i][j]
			k++
		}
	}
	return data
}
