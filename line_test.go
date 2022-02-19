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

func TestLineStyle(t *testing.T) {
	assert := assert.New(t)

	ls := LineStyle{
		ClassName: "test",
		StrokeDashArray: []float64{
			1.0,
		},
		StrokeColor:  drawing.ColorBlack,
		StrokeWidth:  1,
		FillColor:    drawing.ColorBlack.WithAlpha(60),
		DotWidth:     2,
		DotColor:     drawing.ColorBlack,
		DotFillColor: drawing.ColorWhite,
	}

	assert.Equal(chart.Style{
		ClassName: "test",
		StrokeDashArray: []float64{
			1.0,
		},
		StrokeColor: drawing.ColorBlack,
		StrokeWidth: 1,
		FillColor:   drawing.ColorBlack.WithAlpha(60),
		DotWidth:    2,
		DotColor:    drawing.ColorBlack,
	}, ls.Style())
}

func TestDrawLineFill(t *testing.T) {
	assert := assert.New(t)

	ls := LineStyle{
		StrokeColor:  drawing.ColorBlack,
		StrokeWidth:  1,
		FillColor:    drawing.ColorBlack.WithAlpha(60),
		DotWidth:     2,
		DotColor:     drawing.ColorBlack,
		DotFillColor: drawing.ColorWhite,
	}
	d, err := NewDraw(DrawOption{
		Width:  400,
		Height: 300,
	})
	assert.Nil(err)
	d.lineFill([]Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 10,
			Y: 20,
		},
		{
			X: 50,
			Y: 60,
		},
	}, ls)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 0 0\nL 10 20\nL 50 60\nL 50 300\nL 0 300\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,0.2)\"/></svg>", string(data))
}

func TestDrawLineDot(t *testing.T) {
	assert := assert.New(t)

	ls := LineStyle{
		StrokeColor:  drawing.ColorBlack,
		StrokeWidth:  1,
		FillColor:    drawing.ColorBlack.WithAlpha(60),
		DotWidth:     2,
		DotColor:     drawing.ColorBlack,
		DotFillColor: drawing.ColorWhite,
	}
	d, err := NewDraw(DrawOption{
		Width:  400,
		Height: 300,
	})
	assert.Nil(err)
	d.lineDot([]Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 10,
			Y: 20,
		},
		{
			X: 50,
			Y: 60,
		},
	}, ls)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<circle cx=\"0\" cy=\"0\" r=\"2\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"10\" cy=\"20\" r=\"2\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"50\" cy=\"60\" r=\"2\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/></svg>", string(data))
}

func TestDrawLine(t *testing.T) {
	assert := assert.New(t)

	ls := LineStyle{
		StrokeColor:  drawing.ColorBlack,
		StrokeWidth:  1,
		FillColor:    drawing.ColorBlack.WithAlpha(60),
		DotWidth:     2,
		DotColor:     drawing.ColorBlack,
		DotFillColor: drawing.ColorWhite,
	}
	d, err := NewDraw(DrawOption{
		Width:  400,
		Height: 300,
	})
	assert.Nil(err)
	d.Line([]Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 10,
			Y: 20,
		},
		{
			X: 50,
			Y: 60,
		},
	}, ls)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 0 0\nL 10 20\nL 50 60\nL 50 300\nL 0 300\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,0.2)\"/><path  d=\"M 0 0\nL 10 20\nL 50 60\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><circle cx=\"0\" cy=\"0\" r=\"2\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"10\" cy=\"20\" r=\"2\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"50\" cy=\"60\" r=\"2\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(255,255,255,1.0)\"/></svg>", string(data))
}
