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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
)

func TestNewSeriesDataListFromFloat(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]SeriesData{
		{
			Value: 1,
		},
		{
			Value: 2,
		},
	}, NewSeriesDataListFromFloat([]float64{
		1,
		2,
	}))
}

func TestGetSeries(t *testing.T) {
	assert := assert.New(t)

	xValues := []float64{
		1,
		2,
		3,
		4,
		5,
	}

	barData := NewSeriesDataListFromFloat([]float64{
		10,
		20,
		30,
		40,
		50,
	})
	barData[1].Style = chart.Style{
		FillColor: AxisColorDark,
	}
	seriesList := GetSeries([]Series{
		{
			Type:       SeriesBar,
			Data:       barData,
			XValues:    xValues,
			YAxisIndex: 1,
		},
		{
			Data: NewSeriesDataListFromFloat([]float64{
				11,
				21,
				31,
				41,
				51,
			}),
			XValues: xValues,
		},
	}, chart.TickPositionBetweenTicks, "")

	assert.Equal(seriesList[0].GetYAxis(), chart.YAxisPrimary)
	assert.Equal(seriesList[1].GetYAxis(), chart.YAxisSecondary)

	barSeries, ok := seriesList[0].(BarSeries)
	assert.True(ok)
	// 居中前置多插入一个点
	assert.Equal([]float64{
		0,
		10,
		20,
		30,
		40,
		50,
	}, barSeries.YValues)
	assert.Equal(xValues, barSeries.XValues)
	assert.Equal(1, barSeries.Count)
	assert.Equal(0, barSeries.Index)
	assert.Equal([]BarSeriesCustomStyle{
		{
			PointIndex: 1,
			Index:      0,
			Style:      barData[1].Style,
		},
	}, barSeries.CustomStyles)

	lineSeries, ok := seriesList[1].(LineSeries)
	assert.True(ok)
	// 居中前置多插入一个点
	assert.Equal([]float64{
		0,
		11,
		21,
		31,
		41,
		51,
	}, lineSeries.YValues)
	assert.Equal(xValues, lineSeries.XValues)
}
