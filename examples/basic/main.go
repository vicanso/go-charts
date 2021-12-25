package main

import (
	"os"

	charts "github.com/vicanso/go-echarts"
)

func main() {
	buf, err := charts.RenderEChartsToPNG(`{
		"title": {
			"text": "Line"
		},
		"xAxis": {
			"type": "category",
			"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
		},
		"series": [
			{
				"data": [150, 230, 224, 218, 135, 147, 260]
			}
		]
	}`)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(buf)
}
