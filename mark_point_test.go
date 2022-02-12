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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestNewMarkPoint(t *testing.T) {
	assert := assert.New(t)

	markPoint := NewMarkPoint(
		SeriesMarkDataTypeMax,
		SeriesMarkDataTypeMin,
		SeriesMarkDataTypeAverage,
	)

	assert.Equal(SeriesMarkPoint{
		Data: []SeriesMarkData{
			{
				Type: SeriesMarkDataTypeMax,
			},
			{
				Type: SeriesMarkDataTypeMin,
			},
			{
				Type: SeriesMarkDataTypeAverage,
			},
		},
	}, markPoint)
}

func TestMarkPointRender(t *testing.T) {
	assert := assert.New(t)

	d, err := NewDraw(DrawOption{
		Width:  400,
		Height: 300,
	}, PaddingOption(chart.Box{
		Left:  20,
		Right: 20,
	}))
	assert.Nil(err)
	f, _ := chart.GetDefaultFont()

	markPointRender(markPointRenderOption{
		Draw:      d,
		FillColor: drawing.ColorBlack,
		Font:      f,
		Series: &Series{
			MarkPoint: NewMarkPoint(
				SeriesMarkDataTypeMax,
				SeriesMarkDataTypeMin,
			),
			Data: NewSeriesDataFromValues([]float64{
				1,
				3,
				5,
			}),
		},
		Points: []Point{
			{
				X: 1,
				Y: 50,
			},
			{
				X: 100,
				Y: 100,
			},
			{
				X: 200,
				Y: 200,
			},
		},
	})
	data, err := d.Bytes()
	assert.Nil(err)
	fmt.Println(string(data))
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 217 192\nA 15 15 330.00 1 1 223 192\nL 220 178\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0)\"/><path  d=\"M 205 178\nQ220,215 235,178\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0)\"/><text x=\"216\" y=\"183\" style=\"stroke-width:0;stroke:none;fill:rgba(238,238,238,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">5</text><path  d=\"M 18 42\nA 15 15 330.00 1 1 24 42\nL 21 28\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0)\"/><path  d=\"M 6 28\nQ21,65 36,28\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0)\"/><text x=\"17\" y=\"33\" style=\"stroke-width:0;stroke:none;fill:rgba(238,238,238,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">1</text></svg>", string(data))
}
