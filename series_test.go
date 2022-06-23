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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSeriesListDataFromValues(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(SeriesList{
		{
			Type: ChartTypeBar,
			Data: []SeriesData{
				{
					Value: 1.0,
				},
			},
		},
	}, NewSeriesListDataFromValues([][]float64{
		{
			1,
		},
	}, ChartTypeBar))
}

func TestSeriesLists(t *testing.T) {
	assert := assert.New(t)
	seriesList := NewSeriesListDataFromValues([][]float64{
		{
			1,
			2,
		},
		{
			10,
		},
	}, ChartTypeBar)

	assert.Equal(2, len(seriesList.Filter(ChartTypeBar)))
	assert.Equal(0, len(seriesList.Filter(ChartTypeLine)))

	max, min := seriesList.GetMaxMin(0)
	assert.Equal(float64(10), max)
	assert.Equal(float64(1), min)

	assert.Equal(seriesSummary{
		MaxIndex:     1,
		MaxValue:     2,
		MinIndex:     0,
		MinValue:     1,
		AverageValue: 1.5,
	}, seriesList[0].Summary())
}

func TestFormatter(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("a: 12%", NewPieLabelFormatter([]string{
		"a",
		"b",
	}, "")(0, 10, 0.12))

	assert.Equal("10", NewValueLabelFormatter([]string{
		"a",
		"b",
	}, "")(0, 10, 0.12))
}
