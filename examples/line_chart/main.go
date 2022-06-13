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

	file := filepath.Join(tmpPath, "line-chart.png")
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
	_, err = charts.NewLineChart(p, charts.LineChartOption{
		Padding: charts.Box{
			Left:   10,
			Top:    10,
			Right:  10,
			Bottom: 10,
		},
		TitleOption: charts.TitleOption{
			Text: "Line",
		},
		LegendOption: charts.LegendOption{
			Data: []string{
				"Email",
				"Union Ads",
				"Video Ads",
				"Direct",
				"Search Engine",
			},
			Left: charts.PositionCenter,
		},
		XAxis: charts.NewXAxisOption([]string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		}),
		SeriesList: charts.SeriesList{
			charts.NewSeriesFromValues([]float64{
				120,
				132,
				101,
				134,
				90,
				230,
				210,
			}),
			charts.NewSeriesFromValues([]float64{
				220,
				182,
				191,
				234,
				290,
				330,
				310,
			}),
			charts.NewSeriesFromValues([]float64{
				150,
				232,
				201,
				154,
				190,
				330,
				410,
			}),
			charts.NewSeriesFromValues([]float64{
				320,
				332,
				301,
				334,
				390,
				330,
				320,
			}),
			charts.NewSeriesFromValues([]float64{
				820,
				932,
				901,
				934,
				1290,
				1330,
				1320,
			}),
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
