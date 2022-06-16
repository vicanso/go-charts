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

	file := filepath.Join(tmpPath, "radar-chart.png")
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
	_, err = charts.NewRadarChart(p, charts.RadarChartOption{
		Padding: charts.Box{
			Left:   10,
			Top:    10,
			Right:  10,
			Bottom: 10,
		},
		Title: charts.TitleOption{
			Text: "Basic Radar Chart",
		},
		Legend: charts.NewLegendOption([]string{
			"Allocated Budget",
			"Actual Spending",
		}),
		RadarIndicators: []charts.RadarIndicator{
			{
				Name: "Sales",
				Max:  6500,
			},
			{
				Name: "Administration",
				Max:  16000,
			},
			{
				Name: "Information Technology",
				Max:  30000,
			},
			{
				Name: "Customer Support",
				Max:  38000,
			},
			{
				Name: "Development",
				Max:  52000,
			},
			{
				Name: "Marketing",
				Max:  25000,
			},
		},
		SeriesList: charts.SeriesList{
			{
				Type: charts.ChartTypeRadar,
				Data: charts.NewSeriesDataFromValues([]float64{
					4200,
					3000,
					20000,
					35000,
					50000,
					18000,
				}),
			},
			{
				Type: charts.ChartTypeRadar,
				Data: charts.NewSeriesDataFromValues([]float64{
					5000,
					14000,
					28000,
					26000,
					42000,
					21000,
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
