// MIT License

// Copyright (c) 2021 Tree Xie

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
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestConvertToArray(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(convertToArray([]byte(" ")))

	assert.Equal([]byte("[{}]"), convertToArray([]byte("{}")))
	assert.Equal([]byte("[{}]"), convertToArray([]byte("[{}]")))
}

func TestECharsSeriesData(t *testing.T) {
	assert := assert.New(t)

	es := ECharsSeriesData{}
	err := es.UnmarshalJSON([]byte(" "))
	assert.Nil(err)
	assert.Equal(ECharsSeriesData{}, es)

	es = ECharsSeriesData{}
	err = es.UnmarshalJSON([]byte("12.1"))
	assert.Nil(err)
	assert.Equal(ECharsSeriesData{
		Value: 12.1,
	}, es)

	es = ECharsSeriesData{}
	err = es.UnmarshalJSON([]byte(`{
		"value": 12.1,
		"name": "test",
		"itemStyle": {
			"color": "#333"
		}
	}`))
	assert.Nil(err)
	assert.Equal(ECharsSeriesData{
		Value: 12.1,
		Name:  "test",
		ItemStyle: EChartStyle{
			Color: "#333",
		},
	}, es)
}

func TestEChartsPadding(t *testing.T) {
	assert := assert.New(t)

	ep := EChartsPadding{}
	err := ep.UnmarshalJSON([]byte(" "))
	assert.Nil(err)
	assert.Equal(EChartsPadding{}, ep)

	ep = EChartsPadding{}
	err = ep.UnmarshalJSON([]byte("1"))
	assert.Nil(err)
	assert.Equal(EChartsPadding{
		box: chart.Box{
			Top:    1,
			Left:   1,
			Right:  1,
			Bottom: 1,
		},
	}, ep)

	ep = EChartsPadding{}
	err = ep.UnmarshalJSON([]byte("[1, 2]"))
	assert.Nil(err)
	assert.Equal(EChartsPadding{
		box: chart.Box{
			Top:    1,
			Left:   2,
			Right:  2,
			Bottom: 1,
		},
	}, ep)

	ep = EChartsPadding{}
	err = ep.UnmarshalJSON([]byte("[1, 2, 3]"))
	assert.Nil(err)
	assert.Equal(EChartsPadding{
		box: chart.Box{
			Top:    1,
			Right:  2,
			Bottom: 3,
			Left:   2,
		},
	}, ep)

	ep = EChartsPadding{}
	err = ep.UnmarshalJSON([]byte("[1, 2, 3, 4]"))
	assert.Nil(err)
	assert.Equal(EChartsPadding{
		box: chart.Box{
			Top:    1,
			Right:  2,
			Bottom: 3,
			Left:   4,
		},
	}, ep)
}

func TestConvertEChartsSeries(t *testing.T) {
	assert := assert.New(t)

	seriesList, tickPosition := convertEChartsSeries(&ECharsOptions{})
	assert.Empty(seriesList)
	assert.Equal(chart.TickPositionUnset, tickPosition)

	e := ECharsOptions{}
	err := json.Unmarshal([]byte(`{
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
					}
				]
			}
		]
	}`), &e)
	assert.Nil(err)
	seriesList, tickPosition = convertEChartsSeries(&e)
	assert.Equal(chart.TickPositionUnset, tickPosition)
	assert.Equal([]Series{
		{
			Data: []SeriesData{
				{
					Value: 1048,
				},
			},
			Type: SeriesPie,
			Name: "Search Engine",
		},
		{
			Data: []SeriesData{
				{
					Value: 735,
				},
			},
			Type: SeriesPie,
			Name: "Direct",
		},
	}, seriesList)

	err = json.Unmarshal([]byte(`{
		"series": [
			{
				"name": "Evaporation",
				"type": "bar",
				"data": [2, {
					"value": 4.9,
					"itemStyle": {
						"color": "#a90000"
					}
				}, 7, 23.2, 25.6, 76.7, 135.6]
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
	}`), &e)
	assert.Nil(err)
	bar1Data := NewSeriesDataListFromFloat([]float64{
		2, 4.9, 7, 23.2, 25.6, 76.7, 135.6,
	})
	bar1Data[1].Style.FillColor = parseColor("#a90000")
	bar1Data[1].Style.StrokeColor = bar1Data[1].Style.FillColor

	seriesList, tickPosition = convertEChartsSeries(&e)
	assert.Equal(chart.TickPositionBetweenTicks, tickPosition)
	assert.Equal([]Series{
		{
			Data: bar1Data,
			Type: SeriesBar,
		},
		{
			Data: NewSeriesDataListFromFloat([]float64{
				2.6, 5.9, 9, 26.4, 28.7, 70.7, 175.6,
			}),
			Type: SeriesBar,
		},
		{
			Data: NewSeriesDataListFromFloat([]float64{
				2, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3,
			}),
			Type:       SeriesLine,
			YAxisIndex: 1,
		},
	}, seriesList)

}

func TestParseECharsOptions(t *testing.T) {

	assert := assert.New(t)
	options, err := ParseECharsOptions(`{
		"theme": "dark",
		"padding": [5, 10],
		"title": {
			"text": "Multi Line",
			"textAlign": "left",
			"textStyle": {
				"color": "#333",
				"fontSize": 24,
				"height": 40
			}
		},
		"legend": {
			"align": "left",
			"padding": [5, 0, 0, 50],
			"data": ["Email", "Union Ads", "Video Ads", "Direct", "Search Engine"]
		},
		"xAxis": {
			"type": "category",
			"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
			"splitNumber": 10
		},
		"yAxis": [
			{
				"min": 0,
				"max": 250
			},
			{
				"min": 0,
				"max": 25
			}
		],
		"series": [
			{
				"name": "Evaporation",
				"type": "bar",
				"data": [2, {
					"value": 4.9,
					"itemStyle": {
						"color": "#a90000"
					}
				}, 7, 23.2, 25.6, 76.7, 135.6]
			},
			{
				"name": "Precipitation",
				"type": "bar",
				"itemStyle": {
					"color": "#0052d9"
				},
				"data": [2.6, 5.9, 9, 26.4, 28.7, 70.7, 175.6]
			},
			{
				"name": "Temperature",
				"type": "line",
				"yAxisIndex": 1,
				"data": [2, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3]
			}
		]
	}`)

	assert.Nil(err)

	min1 := float64(0)
	max1 := float64(250)
	min2 := float64(0)
	max2 := float64(25)

	bar1Data := NewSeriesDataListFromFloat([]float64{
		2, 4.9, 7, 23.2, 25.6, 76.7, 135.6,
	})
	bar1Data[1].Style.FillColor = parseColor("#a90000")
	bar1Data[1].Style.StrokeColor = bar1Data[1].Style.FillColor

	assert.Equal(Options{
		Theme: ThemeDark,
		Padding: chart.Box{
			Top:    5,
			Bottom: 5,
			Left:   10,
			Right:  10,
		},
		Title: Title{
			Text: "Multi Line",
			Style: chart.Style{
				FontColor: parseColor("#333"),
				FontSize:  24,
				Padding: chart.Box{
					Top:    8,
					Bottom: 8,
				},
			},
		},
		Legend: Legend{
			Data: []string{
				"Email", "Union Ads", "Video Ads", "Direct", "Search Engine",
			},
			Align: "left",
			Padding: chart.Box{
				Top:    5,
				Right:  0,
				Bottom: 0,
				Left:   50,
			},
		},
		XAxis: XAxis{
			Data:        []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
			SplitNumber: 10,
		},
		TickPosition: chart.TickPositionBetweenTicks,
		YAxisOptions: []*YAxisOption{
			{
				Min: &min1,
				Max: &max1,
			},
			{
				Min: &min2,
				Max: &max2,
			},
		},
		Series: []Series{
			{
				Data: bar1Data,
				Type: SeriesBar,
			},
			{
				Data: NewSeriesDataListFromFloat([]float64{
					2.6, 5.9, 9, 26.4, 28.7, 70.7, 175.6,
				}),
				Type: SeriesBar,
				Style: chart.Style{
					StrokeColor: drawing.Color{
						R: 0,
						G: 82,
						B: 217,
						A: 255,
					},
					FillColor: drawing.Color{
						R: 0,
						G: 82,
						B: 217,
						A: 255,
					},
				},
			},
			{
				Data: NewSeriesDataListFromFloat([]float64{
					2, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3,
				}),
				Type:       SeriesLine,
				YAxisIndex: 1,
			},
		},
	}, options)
}
