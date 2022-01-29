package main

import (
	"bytes"
	"net/http"

	charts "github.com/vicanso/go-charts"
	"github.com/wcharczuk/go-chart/v2"
)

var html = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<link rel="icon" href="/favicon.ico" />
		<link type="text/css" rel="styleSheet" href="https://unpkg.com/normalize.css@8.0.1/normalize.css" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<style>
			body {
				background-color: #e0e0e0;
			}
			* {
				margin: 0;
				padding: 0;
			}
			.charts {
				width: 810px;
				margin: 10px auto;
				overflow: hidden;
			}
			.grid {
				float: left;
				margin-right: 10px;
			}
			.grid:nth-child(even) {
				margin-right: 0px;
			}
			.grid svg {
				margin-bottom: 10px;
			}
			h1 {
				text-align: center;
			}
			pre {
				width: 100%;
				margin: auto auto 20px auto;
				max-height: 300px;
				overflow: auto;
				display: block;
			}
			svg{
				margin: auto auto 50px auto;
				display: block;
			}
		</style>
		<title>go-charts</title>
	</head>
	<body>
		<div class="charts">{{body}}</div>
	</body>
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	d, err := charts.NewLineChart(charts.LineChartOption{
		ChartOption: charts.ChartOption{
			Theme: "dark",
			Padding: chart.Box{
				Left:   5,
				Top:    15,
				Bottom: 5,
				Right:  10,
			},
			XAxis: charts.XAxisOption{
				Data: []string{
					"Mon",
					"Tue",
					"Wed",
					"Thu",
					"Fri",
					"Sat",
					"Sun",
				},
				// BoundaryGap: charts.FalseFlag(),
			},
			SeriesList: []charts.Series{
				{
					Data: []charts.SeriesData{
						{
							Value: 150,
						},
						{
							Value: 230,
						},
						{
							Value: 224,
						},
						{
							Value: 218,
						},
						{
							Value: 135,
						},
						{
							Value: 147,
						},
						{
							Value: 260,
						},
					},
				},
				{
					Data: []charts.SeriesData{
						{
							Value: 220,
						},
						{
							Value: 182,
						},
						{
							Value: 191,
						},
						{
							Value: 234,
						},
						{
							Value: 290,
						},
						{
							Value: 330,
						},
						{
							Value: 310,
						},
					},
				},
				{
					Data: []charts.SeriesData{
						{
							Value: 150,
						},
						{
							Value: 232,
						},
						{
							Value: 201,
						},
						{
							Value: 154,
						},
						{
							Value: 190,
						},
						{
							Value: 330,
						},
						{
							Value: 410,
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	buf, _ := d.Bytes()

	data := bytes.ReplaceAll([]byte(html), []byte("{{body}}"), buf)
	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3012", nil)
}
