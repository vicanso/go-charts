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

func TestNewSeriesFromValues(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Series{
		Data: []SeriesData{
			{
				Value: 1,
			},
			{
				Value: 2,
			},
		},
		Type: ChartTypeBar,
	}, NewSeriesFromValues([]float64{
		1,
		2,
	}, ChartTypeBar))
}

func TestNewSeriesDataFromValues(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]SeriesData{
		{
			Value: 1,
		},
		{
			Value: 2,
		},
	}, NewSeriesDataFromValues([]float64{
		1,
		2,
	}))
}

func TestNewPieSeriesList(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]Series{
		{
			Type: ChartTypePie,
			Label: SeriesLabel{
				Show: true,
			},
			Radius: "30%",
			Data: []SeriesData{
				{
					Value: 1,
				},
			},
		},
		{
			Type: ChartTypePie,
			Label: SeriesLabel{
				Show: true,
			},
			Radius: "30%",
			Data: []SeriesData{
				{
					Value: 2,
				},
			},
		},
	}, NewPieSeriesList([]float64{
		1,
		2,
	}, PieSeriesOption{
		Radius:    "30%",
		LabelShow: true,
	}))
}

func TestSeriesSummary(t *testing.T) {
	assert := assert.New(t)

	s := Series{
		Data: NewSeriesDataFromValues([]float64{
			1,
			3,
			5,
			7,
			9,
		}),
	}
	assert.Equal(seriesSummary{
		MaxIndex:     4,
		MaxValue:     9,
		MinIndex:     0,
		MinValue:     1,
		AverageValue: 5,
	}, s.Summary())
}

func TestGetSeriesNames(t *testing.T) {
	assert := assert.New(t)

	sl := SeriesList{
		{
			Name: "a",
		},
		{
			Name: "b",
		},
	}
	assert.Equal([]string{
		"a",
		"b",
	}, sl.Names())
}

func TestNewPieLabelFormatter(t *testing.T) {
	assert := assert.New(t)

	fn := NewPieLabelFormatter([]string{
		"a",
		"b",
	}, "")
	assert.Equal("a: 35%", fn(0, 1.2, 0.35))
}

func TestNewValueLabelFormater(t *testing.T) {
	assert := assert.New(t)
	fn := NewValueLabelFormater([]string{
		"a",
		"b",
	}, "")
	assert.Equal("1.2", fn(0, 1.2, 0.35))
}
