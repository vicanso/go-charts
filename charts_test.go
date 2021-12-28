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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/roboto"
)

func TestFont(t *testing.T) {
	assert := assert.New(t)

	fontFamily := "roboto"
	err := InstallFont(fontFamily, roboto.Roboto)
	assert.Nil(err)

	font, err := GetFont(fontFamily)
	assert.Nil(err)
	assert.NotNil(font)
}

func TestChartsOptions(t *testing.T) {
	assert := assert.New(t)

	o := Options{}

	assert.Equal(errors.New("series can not be empty"), o.validate())

	o.Series = []Series{
		{
			Data: []SeriesData{
				{
					Value: 1,
				},
			},
		},
	}
	assert.Equal(errors.New("series and xAxis is not matched"), o.validate())
	o.XAxis.Data = []string{
		"1",
	}
	assert.Nil(o.validate())

	assert.Equal(DefaultChartWidth, o.getWidth())
	o.Width = 10
	assert.Equal(10, o.getWidth())

	assert.Equal(DefaultChartHeight, o.getHeight())
	o.Height = 10
	assert.Equal(10, o.getHeight())

	padding := chart.NewBox(10, 10, 10, 10)
	o.Padding = padding
	assert.Equal(padding, o.getBackground().Padding)
}

func TestNewPieChart(t *testing.T) {
	assert := assert.New(t)

	data := []Series{
		{
			Data: []SeriesData{
				{
					Value: 10,
				},
			},
			Name: "chrome",
		},
		{
			Data: []SeriesData{
				{
					Value: 2,
				},
			},
			Name: "edge",
		},
	}
	pie := newPieChart(Options{
		Series: data,
	})
	for index, item := range pie.Values {
		assert.Equal(data[index].Name, item.Label)
		assert.Equal(data[index].Data[0].Value, item.Value)
	}
}

func TestNewChart(t *testing.T) {
	assert := assert.New(t)

	data := []Series{
		{
			Data: []SeriesData{
				{
					Value: 10,
				},
				{
					Value: 20,
				},
			},
			Name: "chrome",
		},
		{
			Data: []SeriesData{
				{
					Value: 2,
				},
				{
					Value: 3,
				},
			},
			Name: "edge",
		},
	}

	c := newChart(Options{
		Series: data,
	})
	assert.Empty(c.Elements)
	for index, series := range c.Series {
		assert.Equal(data[index].Name, series.GetName())
	}

	c = newChart(Options{
		Legend: Legend{
			Data: []string{
				"chrome",
				"edge",
			},
		},
	})
	assert.Equal(1, len(c.Elements))
}
