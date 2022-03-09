package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	charts "github.com/vicanso/go-charts"
)

func writeFile(file string, buf []byte) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file = filepath.Join(tmpPath, file)
	err = ioutil.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func chartsRender() ([]byte, error) {
	d, err := charts.LineRender([][]float64{
		{
			150,
			230,
			224,
			218,
			135,
			147,
			260,
		},
	},
		// output type
		charts.PNGTypeOption(),
		// title
		charts.TitleOptionFunc(charts.TitleOption{
			Text: "Line",
		}),
		// x axis
		charts.XAxisOptionFunc(charts.NewXAxisOption([]string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		})),
	)
	if err != nil {
		return nil, err
	}
	return d.Bytes()
}

func echartsRender() ([]byte, error) {
	return charts.RenderEChartsToPNG(`{
		"title": {
			"text": "Line"
		},
		"xAxis": {
			"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
		},
		"series": [
			{
				"data": [150, 230, 224, 218, 135, 147, 260]
			}
		]
	}`)
}

type Render func() ([]byte, error)

func main() {
	m := map[string]Render{
		"charts-line.png":  chartsRender,
		"echarts-line.png": echartsRender,
	}
	for name, fn := range m {
		buf, err := fn()
		if err != nil {
			panic(err)
		}
		err = writeFile(name, buf)
		if err != nil {
			panic(err)
		}
	}
}
