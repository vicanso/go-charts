// MIT License

// Copyright (c) 2022 Tree Xie

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package charts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestConvertToArray(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]byte(`[1]`), convertToArray([]byte("1")))
	assert.Equal([]byte(`[1]`), convertToArray([]byte("[1]")))
}

func TestEChartsPosition(t *testing.T) {
	assert := assert.New(t)
	var p EChartsPosition
	err := p.UnmarshalJSON([]byte("1"))
	assert.Nil(err)
	assert.Equal(EChartsPosition("1"), p)
	err = p.UnmarshalJSON([]byte(`"left"`))
	assert.Nil(err)
	assert.Equal(EChartsPosition("left"), p)
}

func TestEChartsSeriesDataValue(t *testing.T) {
	assert := assert.New(t)

	es := EChartsSeriesDataValue{}
	err := es.UnmarshalJSON([]byte(`[1, 2]`))
	assert.Nil(err)
	assert.Equal(EChartsSeriesDataValue{
		values: []float64{
			1,
			2,
		},
	}, es)
	assert.Equal(NewEChartsSeriesDataValue(1, 2), es)
	assert.Equal(1.0, es.First())
}

func TestEChartsSeriesData(t *testing.T) {
	assert := assert.New(t)
	es := EChartsSeriesData{}
	err := es.UnmarshalJSON([]byte("1.1"))
	assert.Nil(err)
	assert.Equal(EChartsSeriesDataValue{
		values: []float64{
			1.1,
		},
	}, es.Value)

	err = es.UnmarshalJSON([]byte(`{"value":200,"itemStyle":{"color":"#a90000"}}`))
	assert.Nil(err)
	assert.Nil(err)
	assert.Equal(EChartsSeriesData{
		Value: EChartsSeriesDataValue{
			values: []float64{
				200.0,
			},
		},
		ItemStyle: EChartStyle{
			Color: "#a90000",
		},
	}, es)
}

func TestEChartsXAxis(t *testing.T) {
	assert := assert.New(t)
	ex := EChartsXAxis{}
	err := ex.UnmarshalJSON([]byte(`{"boundaryGap": true, "splitNumber": 5, "data": ["a", "b"], "type": "value"}`))
	assert.Nil(err)

	assert.Equal(EChartsXAxis{
		Data: []EChartsXAxisData{
			{
				BoundaryGap: TrueFlag(),
				SplitNumber: 5,
				Data: []string{
					"a",
					"b",
				},
				Type: "value",
			},
		},
	}, ex)
}

func TestEChartStyle(t *testing.T) {
	assert := assert.New(t)

	es := EChartStyle{
		Color: "#999",
	}
	color := drawing.Color{
		R: 153,
		G: 153,
		B: 153,
		A: 255,
	}
	assert.Equal(Style{
		FillColor:   color,
		FontColor:   color,
		StrokeColor: color,
	}, es.ToStyle())
}

func TestEChartsPadding(t *testing.T) {
	assert := assert.New(t)

	eb := EChartsPadding{}

	err := eb.UnmarshalJSON([]byte(`1`))
	assert.Nil(err)
	assert.Equal(Box{
		Left:   1,
		Top:    1,
		Right:  1,
		Bottom: 1,
	}, eb.Box)

	err = eb.UnmarshalJSON([]byte(`[2, 3]`))
	assert.Nil(err)
	assert.Equal(Box{
		Left:   3,
		Top:    2,
		Right:  3,
		Bottom: 2,
	}, eb.Box)

	err = eb.UnmarshalJSON([]byte(`[4, 5, 6]`))
	assert.Nil(err)
	assert.Equal(Box{
		Left:   5,
		Top:    4,
		Right:  5,
		Bottom: 6,
	}, eb.Box)

	err = eb.UnmarshalJSON([]byte(`[4, 5, 6, 7]`))
	assert.Nil(err)
	assert.Equal(Box{
		Left:   7,
		Top:    4,
		Right:  5,
		Bottom: 6,
	}, eb.Box)
}

func TestEChartsMarkPoint(t *testing.T) {
	assert := assert.New(t)

	emp := EChartsMarkPoint{
		SymbolSize: 30,
		Data: []EChartsMarkData{
			{
				Type: "test",
			},
		},
	}
	assert.Equal(SeriesMarkPoint{
		SymbolSize: 30,
		Data: []SeriesMarkData{
			{
				Type: "test",
			},
		},
	}, emp.ToSeriesMarkPoint())
}

func TestEChartsMarkLine(t *testing.T) {
	assert := assert.New(t)

	eml := EChartsMarkLine{
		Data: []EChartsMarkData{
			{
				Type: "min",
			},
			{
				Type: "max",
			},
		},
	}
	assert.Equal(SeriesMarkLine{
		Data: []SeriesMarkData{
			{
				Type: "min",
			},
			{
				Type: "max",
			},
		},
	}, eml.ToSeriesMarkLine())
}

func TestEChartsOption(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		option string
	}{
		{
			option: `{
				"xAxis": {
					"type": "category",
					"data": [
						"Mon",
						"Tue",
						"Wed",
						"Thu",
						"Fri",
						"Sat",
						"Sun"
					]
				},
				"yAxis": {
					"type": "value"
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
			option: `{
				"title": {
					"text": "Referer of a Website",
					"subtext": "Fake Data",
					"left": "center"
				},
				"tooltip": {
					"trigger": "item"
				},
				"legend": {
					"orient": "vertical",
					"left": "left"
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
						],
						"emphasis": {
							"itemStyle": {
								"shadowBlur": 10,
								"shadowOffsetX": 0,
								"shadowColor": "rgba(0, 0, 0, 0.5)"
							}
						}
					}
				]
			}`,
		},
		{
			option: `{
				"title": {
					"text": "Rainfall vs Evaporation",
					"subtext": "Fake Data"
				},
				"tooltip": {
					"trigger": "axis"
				},
				"legend": {
					"data": [
						"Rainfall",
						"Evaporation"
					]
				},
				"toolbox": {
					"show": true,
					"feature": {
						"dataView": {
							"show": true,
							"readOnly": false
						},
						"magicType": {
							"show": true,
							"type": [
								"line",
								"bar"
							]
						},
						"restore": {
							"show": true
						},
						"saveAsImage": {
							"show": true
						}
					}
				},
				"calculable": true,
				"xAxis": [
					{
						"type": "category",
						"data": [
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
							"Dec"
						]
					}
				],
				"yAxis": [
					{
						"type": "value"
					}
				],
				"series": [
					{
						"name": "Rainfall",
						"type": "bar",
						"data": [
							2,
							4.9,
							7,
							23.2,
							25.6,
							76.7,
							135.6,
							162.2,
							32.6,
							20,
							6.4,
							3.3
						],
						"markPoint": {
							"data": [
								{
									"type": "max",
									"name": "Max"
								},
								{
									"type": "min",
									"name": "Min"
								}
							]
						},
						"markLine": {
							"data": [
								{
									"type": "average",
									"name": "Avg"
								}
							]
						}
					},
					{
						"name": "Evaporation",
						"type": "bar",
						"data": [
							2.6,
							5.9,
							9,
							26.4,
							28.7,
							70.7,
							175.6,
							182.2,
							48.7,
							18.8,
							6,
							2.3
						],
						"markPoint": {
							"data": [
								{
									"name": "Max",
									"value": 182.2,
									"xAxis": 7,
									"yAxis": 183
								},
								{
									"name": "Min",
									"value": 2.3,
									"xAxis": 11,
									"yAxis": 3
								}
							]
						},
						"markLine": {
							"data": [
								{
									"type": "average",
									"name": "Avg"
								}
							]
						}
					}
				]
			}`,
		},
	}
	for _, tt := range tests {
		opt := EChartsOption{}
		err := json.Unmarshal([]byte(tt.option), &opt)
		assert.Nil(err)
		assert.NotEmpty(opt.Series)
		assert.NotEmpty(opt.ToOption().SeriesList)
	}
}

func TestRenderEChartsToSVG(t *testing.T) {
	assert := assert.New(t)

	data, err := RenderEChartsToSVG(`{
		"title": {
			"text": "Rainfall vs Evaporation",
			"subtext": "Fake Data"
		},
		"legend": {
			"data": [
				"Rainfall",
				"Evaporation"
			]
		},
		"padding": [10, 30, 10, 10],
		"xAxis": [
			{
				"type": "category",
				"data": [
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
					"Dec"
				]
			}
		],
		"series": [
			{
				"name": "Rainfall",
				"type": "bar",
				"data": [
					2,
					4.9,
					7,
					23.2,
					25.6,
					76.7,
					135.6,
					162.2,
					32.6,
					20,
					6.4,
					3.3
				],
				"markPoint": {
					"data": [
						{
							"type": "max"
						},
						{
							"type": "min"
						}
					]
				},
				"markLine": {
					"data": [
						{
							"type": "average"
						}
					]
				}
			},
			{
				"name": "Evaporation",
				"type": "bar",
				"data": [
					2.6,
					5.9,
					9,
					26.4,
					28.7,
					70.7,
					175.6,
					182.2,
					48.7,
					18.8,
					6,
					2.3
				],
				"markPoint": {
					"data": [
						{
							"type": "max"
						},
						{
							"type": "min"
						}
					]
				},
				"markLine": {
					"data": [
						{
							"type": "average"
						}
					]
				}
			}
		]
	}`)
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 400\nL 0 400\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 182 19\nL 212 19\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><circle cx=\"197\" cy=\"19\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"214\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Rainfall</text><path  d=\"M 286 19\nL 316 19\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><circle cx=\"301\" cy=\"19\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"318\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Evaporation</text><text x=\"10\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Rainfall vs Evaporation</text><text x=\"54\" y=\"40\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Fake Data</text><text x=\"10\" y=\"67\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">240</text><text x=\"10\" y=\"117\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">200</text><text x=\"10\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">160</text><text x=\"10\" y=\"217\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">120</text><text x=\"19\" y=\"267\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">80</text><text x=\"19\" y=\"317\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">40</text><text x=\"28\" y=\"367\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">0</text><path  d=\"M 47 60\nL 570 60\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 110\nL 570 110\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 160\nL 570 160\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 210\nL 570 210\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 260\nL 570 260\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 310\nL 570 310\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 365\nL 47 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 177 365\nL 177 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 308 365\nL 308 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 439 365\nL 439 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 570 365\nL 570 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 47 360\nL 570 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><text x=\"99\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Feb</text><text x=\"227\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">May</text><text x=\"359\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Aug</text><text x=\"490\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Nov</text><path  d=\"M 52 358\nL 67 358\nL 67 359\nL 52 359\nL 52 358\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 95 354\nL 110 354\nL 110 359\nL 95 359\nL 95 354\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 139 352\nL 154 352\nL 154 359\nL 139 359\nL 139 352\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 182 331\nL 197 331\nL 197 359\nL 182 359\nL 182 331\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 226 328\nL 241 328\nL 241 359\nL 226 359\nL 226 328\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 269 265\nL 284 265\nL 284 359\nL 269 359\nL 269 265\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 313 191\nL 328 191\nL 328 359\nL 313 359\nL 313 191\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 357 158\nL 372 158\nL 372 359\nL 357 359\nL 357 158\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 400 320\nL 415 320\nL 415 359\nL 400 359\nL 400 320\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 444 335\nL 459 335\nL 459 359\nL 444 359\nL 444 335\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 487 352\nL 502 352\nL 502 359\nL 487 359\nL 487 352\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 531 356\nL 546 356\nL 546 359\nL 531 359\nL 531 356\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 70 357\nL 85 357\nL 85 359\nL 70 359\nL 70 357\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 113 353\nL 128 353\nL 128 359\nL 113 359\nL 113 353\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 157 349\nL 172 349\nL 172 359\nL 157 359\nL 157 349\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 200 327\nL 215 327\nL 215 359\nL 200 359\nL 200 327\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 244 325\nL 259 325\nL 259 359\nL 244 359\nL 244 325\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 287 272\nL 302 272\nL 302 359\nL 287 359\nL 287 272\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 331 141\nL 346 141\nL 346 359\nL 331 359\nL 331 141\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 375 133\nL 390 133\nL 390 359\nL 375 359\nL 375 133\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 418 300\nL 433 300\nL 433 359\nL 418 359\nL 418 300\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 462 337\nL 477 337\nL 477 359\nL 462 359\nL 462 337\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 505 353\nL 520 353\nL 520 359\nL 505 359\nL 505 353\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 549 358\nL 564 358\nL 564 359\nL 549 359\nL 549 358\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 361 150\nA 15 15 330.00 1 1 367 150\nL 364 136\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 349 136\nQ364,173 379,136\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><text x=\"351\" y=\"141\" style=\"stroke-width:0;stroke:none;fill:rgba(238,238,238,1.0);font-size:10.2px;font-family:'Roboto Medium',sans-serif\">162.2</text><path  d=\"M 56 350\nA 15 15 330.00 1 1 62 350\nL 59 336\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 44 336\nQ59,373 74,336\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><text x=\"55\" y=\"341\" style=\"stroke-width:0;stroke:none;fill:rgba(238,238,238,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2</text><path  d=\"M 379 125\nA 15 15 330.00 1 1 385 125\nL 382 111\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 367 111\nQ382,148 397,111\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><text x=\"369\" y=\"116\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:10.2px;font-family:'Roboto Medium',sans-serif\">182.2</text><path  d=\"M 553 350\nA 15 15 330.00 1 1 559 350\nL 556 336\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 541 336\nQ556,373 571,336\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><text x=\"547\" y=\"341\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2.3</text><circle cx=\"50\" cy=\"308\" r=\"3\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 56 308\nL 552 308\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 552 303\nL 568 308\nL 552 313\nL 557 308\nL 552 303\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"570\" y=\"312\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">41.62</text><circle cx=\"50\" cy=\"300\" r=\"3\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 56 300\nL 552 300\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 552 295\nL 568 300\nL 552 305\nL 557 300\nL 552 295\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"570\" y=\"304\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">48.07</text></svg>", string(data))
}
