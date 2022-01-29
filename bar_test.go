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
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestBarStyle(t *testing.T) {
	assert := assert.New(t)

	bs := BarStyle{
		ClassName: "test",
		StrokeDashArray: []float64{
			1.0,
		},
		FillColor: drawing.ColorBlack,
	}

	assert.Equal(chart.Style{
		ClassName: "test",
		StrokeDashArray: []float64{
			1.0,
		},
		StrokeWidth: 1,
		FillColor:   drawing.ColorBlack,
		StrokeColor: drawing.ColorBlack,
	}, bs.Style())
}

func TestDrawBar(t *testing.T) {
	assert := assert.New(t)
	d, err := NewDraw(DrawOption{
		Width:  400,
		Height: 300,
	}, PaddingOption(chart.Box{
		Left:   10,
		Top:    20,
		Right:  30,
		Bottom: 40,
	}))
	assert.Nil(err)
	d.Bar(chart.Box{
		Left:   0,
		Top:    0,
		Right:  20,
		Bottom: 200,
	}, BarStyle{
		FillColor: drawing.ColorBlack,
	})
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 10 20\nL 30 20\nL 30 220\nL 10 220\nL 10 20\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/></svg>", string(data))
}
