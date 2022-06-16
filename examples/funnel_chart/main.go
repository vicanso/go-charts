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

	file := filepath.Join(tmpPath, "funnel-chart.png")
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
	_, err = charts.NewFunnelChart(p, charts.FunnelChartOption{
		Title: charts.TitleOption{
			Text: "Funnel",
		},
		Legend: charts.NewLegendOption([]string{
			"Show",
			"Click",
			"Visit",
			"Inquiry",
			"Order",
		}),
		SeriesList: []charts.Series{

			{
				Type: charts.ChartTypeFunnel,
				Name: "Show",
				Data: charts.NewSeriesDataFromValues([]float64{
					100,
				}),
			},
			{
				Type: charts.ChartTypeFunnel,
				Name: "Click",
				Data: charts.NewSeriesDataFromValues([]float64{
					80,
				}),
			},
			{
				Type: charts.ChartTypeFunnel,
				Name: "Visit",
				Data: charts.NewSeriesDataFromValues([]float64{
					60,
				}),
			},
			{
				Type: charts.ChartTypeFunnel,
				Name: "Inquiry",
				Data: charts.NewSeriesDataFromValues([]float64{
					40,
				}),
			},
			{
				Type: charts.ChartTypeFunnel,
				Name: "Order",
				Data: charts.NewSeriesDataFromValues([]float64{
					20,
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
