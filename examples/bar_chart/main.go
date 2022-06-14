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

	file := filepath.Join(tmpPath, "bar-chart.png")
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
	_, err = charts.NewBarChart(p, charts.BarChartOption{
		Title: charts.TitleOption{
			Text:    "Rainfall vs Evaporation",
			Subtext: "Fake Data",
		},
		Padding: charts.Box{
			Top:    20,
			Right:  20,
			Bottom: 20,
			Left:   20,
		},
		XAxis: charts.NewXAxisOption([]string{
			"Jan",
			"Feb",
			"Mar",
			"Apr",
			"May",
			"Jun",
			"Jul",
			"Aug",
			"Sep",
			"Oct",
			"Nov",
			"Dec",
		}),
		Legend: charts.NewLegendOption([]string{
			"Rainfall",
			"Evaporation",
		}, charts.PositionRight),
		SeriesList: []charts.Series{
			{
				Type: charts.ChartTypeBar,
				Data: charts.NewSeriesDataFromValues([]float64{
					2.0,
					4.9,
					7.0,
					23.2,
					25.6,
					76.7,
					135.6,
					162.2,
					32.6,
					20.0,
					6.4,
					3.3,
				}),
				MarkPoint: charts.NewMarkPoint(
					charts.SeriesMarkDataTypeMax,
					charts.SeriesMarkDataTypeMin,
				),
				MarkLine: charts.NewMarkLine(
					charts.SeriesMarkDataTypeAverage,
				),
			},
			{
				Type: charts.ChartTypeBar,
				Data: charts.NewSeriesDataFromValues([]float64{
					2.6,
					5.9,
					9.0,
					26.4,
					28.7,
					70.7,
					175.6,
					182.2,
					48.7,
					18.8,
					6.0,
					2.3,
				}),
				MarkPoint: charts.NewMarkPoint(
					charts.SeriesMarkDataTypeMax,
					charts.SeriesMarkDataTypeMin,
				),
				MarkLine: charts.NewMarkLine(
					charts.SeriesMarkDataTypeAverage,
				),
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
