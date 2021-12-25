package main

import (
	"bytes"
	"net/http"

	charts "github.com/vicanso/go-charts"
)

var html = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<link rel="icon" href="/favicon.ico" />
		<link type="text/css" rel="styleSheet" href="https://unpkg.com/normalize.css@8.0.1/normalize.css" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<style>
			.charts {
				width: 830px;
				margin: 10px auto;
				overflow: hidden;
			}
			.grid {
				float: left;
				margin-right: 10px;
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

var chartOptions = []map[string]string{
	{
		"title": "折线图",
		"option": `{
	"title": {
		"text": "Line",
		"textAlign": "left",
		"textStyle": {
			"fontSize": 24,
			"height": 40
		}
	},
	"yAxis": {
		"min": 0,
		"max": 300
	},
	"xAxis": {
		"type": "category",
		"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
	},
	"series": [
		{
			"data": [150, 230, 224, 218, 135, 147, 260],
			"type": "line"
		}
	]
}`,
	},
	{
		"title": "多折线图",
		"option": `{
	"title": {
		"text": "Multi Line"
	},
	"legend": {
		"align": "left",
		"right": 0,
		"data": ["Email", "Union Ads", "Video Ads", "Direct", "Search Engine"]
	},
	"xAxis": {
		"type": "category",
		"boundaryGap": false,
		"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
	},
	"series": [
		{
			"type": "line",
			"data": [120, 132, 101, 134, 90, 230, 210]
		},
		{
			"data": [220, 182, 191, 234, 290, 330, 310]
		},
		{
			"data": [150, 232, 201, 154, 190, 330, 410]
		},
		{
			"data": [320, 332, 301, 334, 390, 330, 320]
		},
		{
			"data": [820, 932, 901, 934, 1290, 1330, 1320]
		}
	]
}`,
	},
	{
		"title": "柱状图",
		"option": `{
	"xAxis": {
		"type": "category",
		"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
	},
	"series": [
		{
			"data": [120, 200, 150, 80, 70, 110, 130],
			"type": "bar"
		}
	]
}`,
	},
	{
		"title": "柱状图（自定义颜色)",
		"option": `{
	"xAxis": {
		"type": "category",
		"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
	},
	"series": [
		{
			"data": [
				120,
				{
					"value": 200,
					"itemStyle": {
						"color": "#a90000"
					}
				},
				150,
				80,
				70,
				110,
				130
			],
			"type": "bar"
		}
	]
}`,
	},
	{
		"title": "多柱状图",
		"option": `{
	"title": {
		"text": "Rainfall vs Evaporation"
	},
	"legend": {
		"data": ["Rainfall", "Evaporation"]
	},
	"xAxis": {
		"type": "category",
		"splitNumber": 12,
		"data": ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
	},
	"series": [
		{
			"name": "Rainfall",
			"type": "bar",
			"data": [2, 4.9, 7, 23.2, 25.6, 76.7, 135.6, 162.2, 32.6, 20, 6.4, 3.3]
		},
		{
			"name": "Evaporation",
			"type": "bar",
			"data": [2.6, 5.9, 9, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6, 2.3]
		}
	]
}`,
	},
	{
		"title": "折柱混合",
		"option": `{
	"legend": {
		"data": [
			"Evaporation",
			"Precipitation",
			"Temperature"
		]
	},
	"xAxis": {
		"type": "category",
		"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
	},
	"yAxis": [
		{
			"type": "value",
			"name": "Precipitation",
			"min": 0,
			"max": 250,
			"interval": 50,
			"axisLabel": {
				"formatter": "{value} ml"
			}
		},
		{
			"type": "value",
			"name": "Temperature",
			"min": 0,
			"max": 25,
			"interval": 5,
			"axisLabel": {
				"formatter": "{value} °C"
			}
		}
	],
	"series": [
		{
			"name": "Evaporation",
			"type": "bar",
			"itemStyle": {
				"color": "#0052d9"
			},
			"data": [2, 4.9, 7, 23.2, 25.6, 76.7, 135.6]
		},
		{
			"name": "Precipitation",
			"type": "bar",
			"data": [2.6, 5.9, 9, 26.4, 28.7, 70.7, 175.6]
		},
		{
			"name": "Temperature",
			"type": "line",
			"yAxisIndex": 1,
			"data": [2, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3]
		}
	]
}`,
	},
	{
		"title": "降雨量",
		"option": `{
	"title": {
		"text": "降雨量"
	},
	"legend": {
		"data": ["GZ", "SH"]
	},
	"xAxis": {
		"type": "category",
		"splitNumber": 6,
		"data": ["01-01","01-02","01-03","01-04","01-05","01-06","01-07","01-08","01-09","01-10","01-11","01-12","01-13","01-14","01-15","01-16","01-17","01-18","01-19","01-20","01-21","01-22","01-23","01-24","01-25","01-26","01-27","01-28","01-29","01-30","01-31"]
	},
	"yAxis": {
		"axisLabel": {
			"formatter": "{value} mm"
		}
	},
	"series": [
		{
			"type": "bar",
			"data": [928,821,889,600,547,783,197,853,430,346,63,465,309,334,141,538,792,58,922,807,298,243,744,885,812,231,330,220,984,221,429]
		},
		{
			"type": "bar",
			"data": [749,201,296,579,255,159,902,246,149,158,507,776,186,79,390,222,601,367,221,411,714,620,966,73,203,631,833,610,487,677,596]
		}
	]
}`,
	},
	{
		"title": "饼图",
		"option": `{
	"title": {
		"text": "Referer of a Website"
	},
	"series": [
		{
			"name": "Access From",
			"type": "pie",
			"radius": "50%",
			"data": [
				{
					"value": 1048,
					"name": "Search Engine"
				},
				{
					"value": 735,
					"name": "Direct"
				},
				{
					"value": 580,
					"name": "Email"
				},
				{
					"value": 484,
					"name": "Union Ads"
				},
				{
					"value": 300,
					"name": "Video Ads"
				}
			]
		}
	]
}`,
	},
}

type renderOptions struct {
	theme      string
	width      int
	height     int
	onlyCharts bool
}

func render(opts renderOptions) ([]byte, error) {
	data := bytes.Buffer{}
	for _, m := range chartOptions {
		chartHTML := []byte(`<div>
			<h1>{{title}}</h1>
			<pre>{{option}}</pre>
			{{svg}}
		</div>`)
		if opts.onlyCharts {
			chartHTML = []byte(`<div class="grid">
				{{svg}}
			</div>`)
		}
		o, err := charts.ParseECharsOptions(m["option"])
		if opts.width > 0 {
			o.Width = opts.width
		}
		if opts.height > 0 {
			o.Height = opts.height
		}

		o.Theme = opts.theme
		if err != nil {
			return nil, err
		}
		g, err := charts.New(o)
		if err != nil {
			return nil, err
		}
		buf, err := charts.ToSVG(g)
		if err != nil {
			return nil, err
		}
		buf = bytes.ReplaceAll(chartHTML, []byte("{{svg}}"), buf)
		buf = bytes.ReplaceAll(buf, []byte("{{title}}"), []byte(m["title"]))
		buf = bytes.ReplaceAll(buf, []byte("{{option}}"), []byte(m["option"]))
		data.Write(buf)
	}
	return data.Bytes(), nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	opts := renderOptions{
		theme: query.Get("theme"),
	}
	if query.Get("view") == "grid" {
		opts.width = 400
		opts.height = 200
		opts.onlyCharts = true
	}

	buf, err := render(opts)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	data := bytes.ReplaceAll([]byte(html), []byte("{{body}}"), buf)
	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3012", nil)
}
