package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vicanso/go-charts"
)

func writeFile(buf []byte) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file := filepath.Join(tmpPath, "pie-chart.png")
	err = ioutil.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	p, err := charts.NewPainter(charts.PainterOptions{
		Width:  800,
		Height: 600,
		Type:   charts.ChartOutputPNG,
	})
	if err != nil {
		panic(err)
	}
	_, err = charts.NewPieChart(p, charts.PieChartOption{
		Title: charts.TitleOption{
			Text:    "Rainfall vs Evaporation",
			Subtext: "Fake Data",
			Left:    charts.PositionCenter,
		},
		Padding: charts.Box{
			Top:    20,
			Right:  20,
			Bottom: 20,
			Left:   20,
		},
		Legend: charts.LegendOption{
			Orient: charts.OrientVertical,
			Data: []string{
				"Search Engine",
				"Direct",
				"Email",
				"Union Ads",
				"Video Ads",
			},
			Left: charts.PositionLeft,
		},
		SeriesList: charts.NewPieSeriesList([]float64{
			1048,
			735,
			580,
			484,
			300,
		}, charts.PieSeriesOption{
			Label: charts.SeriesLabel{
				Show: true,
			},
			Radius: "35%",
		}),
	}).Render()
	if err != nil {
		panic(err)
	}

	buf, err := p.Bytes()
	if err != nil {
		panic(err)
	}
	err = writeFile(buf)
	if err != nil {
		panic(err)
	}
}
