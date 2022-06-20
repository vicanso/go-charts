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

func TestEChartsOption(t *testing.T) {
	assert := assert.New(t)

	opt := EChartsOption{}
	err := json.Unmarshal([]byte(`{
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
	}`), &opt)

	assert.Nil(err)
	assert.NotEmpty(opt.Series)
}
