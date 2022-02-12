package main

import (
	"bytes"
	"net/http"

	charts "github.com/vicanso/go-charts"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

var html = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<link rel="icon" href="/favicon.ico" />
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
				margin: auto auto 10px auto;
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

func indexHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		return
	}
	theme := req.URL.Query().Get("theme")
	chartOptions := []charts.ChartOption{
		// 普通折线图
		{
			Theme: theme,
			Title: charts.TitleOption{
				Text: "Line",
			},
			Legend: charts.NewLegendOption([]string{
				"Email",
				"Union Ads",
				"Video Ads",
				"Direct",
				"Search Engine",
			}),
			XAxis: charts.NewXAxisOption([]string{
				"Mon",
				"Tue",
				"Wed",
				"Thu",
				"Fri",
				"Sat",
				"Sun",
			}),
			SeriesList: []charts.Series{
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
		},
		// 温度折线图
		{
			Theme: theme,
			Title: charts.TitleOption{
				Text: "Temperature Change in the Coming Week",
			},
			Padding: chart.Box{
				Top:    20,
				Left:   20,
				Right:  30,
				Bottom: 20,
			},
			Legend: charts.NewLegendOption([]string{
				"Highest",
				"Lowest",
			}, charts.PositionRight),
			XAxis: charts.NewXAxisOption([]string{
				"Mon",
				"Tue",
				"Wed",
				"Thu",
				"Fri",
				"Sat",
				"Sun",
			}, charts.FalseFlag()),
			SeriesList: []charts.Series{
				{
					Data: charts.NewSeriesDataFromValues([]float64{
						14,
						11,
						13,
						11,
						12,
						12,
						7,
					}),
					MarkPoint: charts.NewMarkPoint(charts.SeriesMarkDataTypeMax, charts.SeriesMarkDataTypeMin),
					MarkLine:  charts.NewMarkLine(charts.SeriesMarkDataTypeAverage),
				},
				{
					Data: charts.NewSeriesDataFromValues([]float64{
						1,
						-2,
						2,
						5,
						3,
						2,
						0,
					}),
					MarkLine: charts.NewMarkLine(charts.SeriesMarkDataTypeAverage),
				},
			},
		},
		// 柱状图
		{
			Theme: theme,
			Title: charts.TitleOption{
				Text: "Bar",
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
			Legend: charts.NewLegendOption([]string{
				"Rainfall",
				"Evaporation",
			}),
			SeriesList: []charts.Series{
				charts.NewSeriesFromValues([]float64{
					120,
					200,
					150,
					80,
					70,
					110,
					130,
				}, charts.ChartTypeBar),
				{
					Type: charts.ChartTypeBar,
					Data: []charts.SeriesData{
						{
							Value: 100,
						},
						{
							Value: 190,
							Style: chart.Style{
								FillColor: drawing.Color{
									R: 169,
									G: 0,
									B: 0,
									A: 255,
								},
							},
						},
						{
							Value: 230,
						},
						{
							Value: 140,
						},
						{
							Value: 100,
						},
						{
							Value: 200,
						},
						{
							Value: 180,
						},
					},
				},
			},
		},
		// 柱状图+mark
		{
			Theme: theme,
			Title: charts.TitleOption{
				Text:    "Rainfall vs Evaporation",
				Subtext: "Fake Data",
			},
			Padding: chart.Box{
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
		},
		// 双Y轴示例
		{
			Theme: theme,
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
				"Evaporation",
				"Precipitation",
				"Temperature",
			}),
			YAxisList: []charts.YAxisOption{
				{
					Formatter: "{value}°C",
					Color: drawing.Color{
						R: 250,
						G: 200,
						B: 88,
						A: 255,
					},
				},
				{
					Formatter: "{value}ml",
					Color: drawing.Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
				},
			},
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
					YAxisIndex: 1,
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
					YAxisIndex: 1,
				},
				{
					Data: charts.NewSeriesDataFromValues([]float64{
						2.0,
						2.2,
						3.3,
						4.5,
						6.3,
						10.2,
						20.3,
						23.4,
						23.0,
						16.5,
						12.0,
						6.2,
					}),
				},
			},
		},
		// 饼图
		{
			Theme: theme,
			Title: charts.TitleOption{
				Text:    "Referer of a Website",
				Subtext: "Fake Data",
				Left:    charts.PositionCenter,
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
				LabelShow: true,
				Radius:    "35%",
			}),
		},
		// 多图展示
		{
			Theme: theme,
			Legend: charts.LegendOption{
				Top: "-90",
				Data: []string{
					"Milk Tea",
					"Matcha Latte",
					"Cheese Cocoa",
					"Walnut Brownie",
				},
			},
			Padding: chart.Box{
				Top: 100,
			},
			XAxis: charts.NewXAxisOption([]string{
				"2012",
				"2013",
				"2014",
				"2015",
				"2016",
				"2017",
			}),
			YAxisList: []charts.YAxisOption{
				{

					Min: charts.NewFloatPoint(0),
					Max: charts.NewFloatPoint(90),
				},
			},
			SeriesList: []charts.Series{
				charts.NewSeriesFromValues([]float64{
					56.5,
					82.1,
					88.7,
					70.1,
					53.4,
					85.1,
				}),
				charts.NewSeriesFromValues([]float64{
					51.1,
					51.4,
					55.1,
					53.3,
					73.8,
					68.7,
				}),
				charts.NewSeriesFromValues([]float64{
					40.1,
					62.2,
					69.5,
					36.4,
					45.2,
					32.5,
				}, charts.ChartTypeBar),
				charts.NewSeriesFromValues([]float64{
					25.2,
					37.1,
					41.2,
					18,
					33.9,
					49.1,
				}, charts.ChartTypeBar),
			},
			Children: []charts.ChartOption{
				{
					Legend: charts.LegendOption{
						Show: charts.FalseFlag(),
						Data: []string{
							"Milk Tea",
							"Matcha Latte",
							"Cheese Cocoa",
							"Walnut Brownie",
						},
					},
					Height: 20,
					Padding: chart.Box{
						Left: 250,
						Top:  -80,
					},
					SeriesList: charts.NewPieSeriesList([]float64{
						435.9,
						354.3,
						285.9,
						204.5,
					}, charts.PieSeriesOption{
						LabelShow: true,
						Radius:    "35%",
					}),
				},
			},
		},
	}
	bytesList := make([][]byte, 0)
	for _, opt := range chartOptions {
		d, err := charts.Render(opt)
		if err != nil {
			panic(err)
		}
		buf, err := d.Bytes()
		if err != nil {
			panic(err)
		}
		bytesList = append(bytesList, buf)
	}
	data := bytes.ReplaceAll([]byte(html), []byte("{{body}}"), bytes.Join(bytesList, []byte("")))
	w.Header().Set("Content-Type", "text/html")
	w.Write(data)

}

func indexHandlerBak(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	zero := float64(0)
	outputType := r.URL.Query().Get("type")
	chartOption := charts.ChartOption{
		// Theme: "dark",
		Type: outputType,
		Title: charts.TitleOption{
			Left: charts.PositionCenter,
			Style: chart.Style{
				FontColor: chart.ColorAlternateBlue,
			},
			SubtextStyle: chart.Style{
				FontColor: chart.ColorRed,
			},
			Text:    "Stacked Line",
			Subtext: "Hello World!",
		},
		Padding: chart.Box{
			Left:   5,
			Top:    15,
			Bottom: 5,
			Right:  10,
		},
		YAxisList: []charts.YAxisOption{
			{
				Min: &zero,
			},
			{
				Formatter: "{value} °C",
				// Max:       charts.NewFloatPoint(24),
			},
		},
		Legend: charts.LegendOption{
			Data: []string{
				"Email",
				"Union Ads",
				// "Video Ads",
			},
			Left: charts.PositionLeft,
			// Orient: charts.OrientVertical,
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
			// SplitNumber: 4,
			// BoundaryGap: charts.FalseFlag(),
		},
		SeriesList: []charts.Series{
			{
				// Type: charts.ChartTypeBar,
				MarkPoint: charts.SeriesMarkPoint{
					Data: []charts.SeriesMarkPointData{
						{
							Type: "max",
						},
						{
							Type: "min",
						},
					},
				},
				MarkLine: charts.SeriesMarkLine{
					Data: []charts.SeriesMarkLineData{
						// {
						// Type: "max",
						// },
						{
							Type: "average",
						},
					},
				},
				Data: charts.NewSeriesDataFromValues([]float64{
					2.0,
					4.9,
					7.0,
					23.2,
					25.6,
					76.7,
					135.6,
				}),
			},
			// {
			// 	// Type: charts.ChartTypeBar,
			// 	Data: charts.NewSeriesDataFromValues([]float64{
			// 		2.6,
			// 		5.9,
			// 		9.0,
			// 		26.4,
			// 		28.7,
			// 		70.7,
			// 		175.6,
			// 	}),
			// },
			{
				Data: charts.NewSeriesDataFromValues([]float64{
					2.0,
					2.2,
					3.3,
					4.5,
					6.3,
					10.2,
					20.3,
				}),
				YAxisIndex: 1,
			},
			// {
			// 	Data: []charts.SeriesData{
			// 		{
			// 			Value: 220,
			// 		},
			// 		{
			// 			Value: 182,
			// 		},
			// 		{
			// 			Value: 191,
			// 		},
			// 		{
			// 			Value: 234,
			// 		},
			// 		{
			// 			Value: 290,
			// 		},
			// 		{
			// 			Value: 330,
			// 		},
			// 		{
			// 			Value: 310,
			// 		},
			// 	},
			// },
			// {
			// 	Data: []charts.SeriesData{
			// 		{
			// 			Value: 150,
			// 		},
			// 		{
			// 			Value: 232,
			// 		},
			// 		{
			// 			Value: 201,
			// 		},
			// 		{
			// 			Value: 154,
			// 		},
			// 		{
			// 			Value: 190,
			// 		},
			// 		{
			// 			Value: 330,
			// 		},
			// 		{
			// 			Value: 410,
			// 		},
			// 	},
			// },
		},
		// Children: []charts.ChartOption{
		// 	{
		// 		Padding: chart.Box{
		// 			Left: 350,
		// 			Top:  0,
		// 		},
		// 		Legend: charts.LegendOption{
		// 			Show: charts.FalseFlag(),
		// 		},
		// 		Width:  150,
		// 		Height: 150,
		// 		SeriesList: []charts.Series{
		// 			charts.NewSeriesFromValues([]float64{
		// 				1048,
		// 			}, charts.ChartTypePie),
		// 			{
		// 				Data: charts.NewSeriesDataFromValues([]float64{
		// 					735,
		// 				}),
		// 				Radius: "50%",
		// 				Name:   "test",
		// 			},
		// 			charts.NewSeriesFromValues([]float64{
		// 				580,
		// 			}),
		// 			charts.NewSeriesFromValues([]float64{
		// 				484,
		// 			}),
		// 		},
		// 	},
		// },
	}
	d, err := charts.Render(chartOption)
	if err != nil {
		panic(err)
	}

	buf, _ := d.Bytes()

	if outputType == "png" {
		w.Header().Set("Content-Type", "image/png")
		w.Write(buf)
	} else {
		data := bytes.ReplaceAll([]byte(html), []byte("{{body}}"), buf)
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3012", nil)
}
