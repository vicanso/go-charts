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
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestParentOption(t *testing.T) {
	assert := assert.New(t)
	p, err := NewDraw(DrawOption{
		Width:  400,
		Height: 300,
	})
	assert.Nil(err)

	d, err := NewDraw(DrawOption{
		Parent: p,
	})
	assert.Nil(err)
	assert.Equal(p, d.parent)
}

func TestWidthHeightOption(t *testing.T) {
	assert := assert.New(t)

	// no parent
	width := 300
	height := 200
	d, err := NewDraw(DrawOption{
		Width:  width,
		Height: height,
	})
	assert.Nil(err)
	assert.Equal(chart.Box{
		Top:    0,
		Left:   0,
		Right:  width,
		Bottom: height,
	}, d.Box)

	width = 500
	height = 600
	// with parent
	p, err := NewDraw(
		DrawOption{
			Width:  width,
			Height: height,
		},
		PaddingOption(chart.NewBox(5, 5, 5, 5)),
	)
	assert.Nil(err)
	d, err = NewDraw(
		DrawOption{
			Parent: p,
		},
		PaddingOption(chart.NewBox(1, 2, 3, 4)),
	)
	assert.Nil(err)
	assert.Equal(chart.Box{
		Top:    6,
		Left:   7,
		Right:  492,
		Bottom: 591,
	}, d.Box)
}

func TestPaddingOption(t *testing.T) {
	assert := assert.New(t)

	d, err := NewDraw(DrawOption{
		Width:  400,
		Height: 300,
	})
	assert.Nil(err)

	// 默认的box
	assert.Equal(chart.Box{
		Right:  400,
		Bottom: 300,
	}, d.Box)

	// 设置padding之后的box
	d, err = NewDraw(DrawOption{
		Width:  400,
		Height: 300,
	}, PaddingOption(chart.Box{
		Left:   1,
		Top:    2,
		Right:  3,
		Bottom: 4,
	}))
	assert.Nil(err)
	assert.Equal(chart.Box{
		Top:    2,
		Left:   1,
		Right:  397,
		Bottom: 296,
	}, d.Box)

	p := d
	// 设置父元素之后的box
	d, err = NewDraw(
		DrawOption{
			Parent: p,
		},
		PaddingOption(chart.Box{
			Left:   1,
			Top:    2,
			Right:  3,
			Bottom: 4,
		}),
	)
	assert.Nil(err)
	assert.Equal(chart.Box{
		Top:    4,
		Left:   2,
		Right:  394,
		Bottom: 292,
	}, d.Box)
}

func TestParentTop(t *testing.T) {
	assert := assert.New(t)
	d1, err := NewDraw(DrawOption{
		Width:  400,
		Height: 300,
	})
	assert.Nil(err)

	d2, err := NewDraw(DrawOption{
		Parent: d1,
	})
	assert.Nil(err)

	d3, err := NewDraw(DrawOption{
		Parent: d2,
	})
	assert.Nil(err)

	assert.Equal(d2, d3.Parent())
	assert.Equal(d1, d2.Parent())
	assert.Equal(d1, d3.Top())
	assert.Equal(d1, d2.Top())
}

func TestDraw(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		fn     func(d *Draw)
		result string
	}{
		// moveTo, lineTo
		{
			fn: func(d *Draw) {
				d.moveTo(1, 1)
				d.lineTo(2, 2)
				d.Render.Stroke()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 6 11\nL 7 12\" style=\"stroke-width:0;stroke:none;fill:none\"/></svg>",
		},
		// circle
		{
			fn: func(d *Draw) {
				d.circle(5, 2, 3)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<circle cx=\"7\" cy=\"13\" r=\"5\" style=\"stroke-width:0;stroke:none;fill:none\"/></svg>",
		},
		// text
		{
			fn: func(d *Draw) {
				d.text("hello world!", 3, 6)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<text x=\"8\" y=\"16\" style=\"stroke-width:0;stroke:none;fill:none\">hello world!</text></svg>",
		},
		// line stroke
		{
			fn: func(d *Draw) {
				d.lineStroke([]Point{
					{
						X: 1,
						Y: 2,
					},
					{
						X: 3,
						Y: 4,
					},
				}, LineStyle{
					StrokeColor: drawing.ColorBlack,
					StrokeWidth: 1,
				})
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 6 12\nL 8 14\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/></svg>",
		},
		// set background
		{
			fn: func(d *Draw) {
				d.setBackground(400, 300, chart.ColorWhite)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 0 0\nL 400 0\nL 400 300\nL 0 300\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/></svg>",
		},
		// arcTo
		{
			fn: func(d *Draw) {
				chart.Style{
					StrokeWidth: 1,
					StrokeColor: drawing.ColorBlack,
					FillColor:   drawing.ColorBlue,
				}.WriteToRenderer(d.Render)
				d.arcTo(100, 100, 100, 100, 0, math.Pi/2)
				d.Render.Close()
				d.Render.FillStroke()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 205 110\nA 100 100 90.00 0 1 105 210\nZ\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,255,1.0)\"/></svg>",
		},
		{
			fn: func(d *Draw) {
				chart.Style{
					StrokeWidth: 1,
					StrokeColor: drawing.Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
					FillColor: drawing.Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
				}.WriteToRenderer(d.Render)
				d.pin(30, 30, 30)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 32 47\nA 15 15 330.00 1 1 38 47\nL 35 33\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 20 33\nQ35,70 50,33\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/></svg>",
		},
	}
	for _, tt := range tests {
		d, err := NewDraw(DrawOption{
			Width:  400,
			Height: 300,
		}, PaddingOption(chart.Box{
			Left: 5,
			Top:  10,
		}))
		assert.Nil(err)
		tt.fn(d)
		data, err := d.Bytes()
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}
