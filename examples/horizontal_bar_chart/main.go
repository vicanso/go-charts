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

	file := filepath.Join(tmpPath, "horizontal-bar-chart.png")
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
	_, err = charts.NewHorizontalBarChart(p, charts.HorizontalBarChartOption{
		Title: charts.TitleOption{
			Text: "World Population",
		},
		Padding: charts.Box{
			Top:    20,
			Right:  40,
			Bottom: 20,
			Left:   20,
		},
		Legend: charts.NewLegendOption([]string{
			"2011",
			"2012",
		}),
		YAxisOptions: charts.NewYAxisOptions([]string{
			"Brazil",
			"Indonesia",
			"USA",
			"India",
			"China",
			"World",
		}),
		SeriesList: []charts.Series{
			{
				Type: charts.ChartTypeHorizontalBar,
				Data: charts.NewSeriesDataFromValues([]float64{
					18203,
					23489,
					29034,
					104970,
					131744,
					630230,
				}),
			},
			{
				Type: charts.ChartTypeHorizontalBar,
				Data: charts.NewSeriesDataFromValues([]float64{
					19325,
					23438,
					31000,
					121594,
					134141,
					681807,
				}),
			},
		},
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
