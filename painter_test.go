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

	"github.com/golang/freetype/truetype"
	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestPainterOption(t *testing.T) {
	assert := assert.New(t)

	font := &truetype.Font{}
	d, err := NewPainter(PainterOptions{
		Width:  800,
		Height: 600,
		Type:   ChartOutputSVG,
	},
		PainterBoxOption(Box{
			Right:  400,
			Bottom: 300,
		}),
		PainterPaddingOption(Box{
			Left:   1,
			Top:    2,
			Right:  3,
			Bottom: 4,
		}),
		PainterFontOption(font),
		PainterStyleOption(Style{
			ClassName: "test",
		}),
	)
	assert.Nil(err)
	assert.Equal(Box{
		Left:   1,
		Top:    2,
		Right:  397,
		Bottom: 296,
	}, d.box)
	assert.Equal(font, d.font)
	assert.Equal("test", d.style.ClassName)
}

func TestPainter(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		fn     func(*Painter)
		result string
	}{
		// moveTo, lineTo
		{
			fn: func(p *Painter) {
				p.MoveTo(1, 1)
				p.LineTo(2, 2)
				p.Stroke()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 6 11\nL 7 12\" style=\"stroke-width:0;stroke:none;fill:none\"/></svg>",
		},
		// circle
		{
			fn: func(p *Painter) {
				p.Circle(5, 2, 3)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<circle cx=\"7\" cy=\"13\" r=\"5\" style=\"stroke-width:0;stroke:none;fill:none\"/></svg>",
		},
		// text
		{
			fn: func(p *Painter) {
				p.Text("hello world!", 3, 6)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<text x=\"8\" y=\"16\" style=\"stroke-width:0;stroke:none;fill:none;font-family:'Roboto Medium',sans-serif\">hello world!</text></svg>",
		},
		// line stroke
		{
			fn: func(p *Painter) {
				p.SetDrawingStyle(Style{
					StrokeColor: drawing.ColorBlack,
					StrokeWidth: 1,
				})
				p.LineStroke([]Point{
					{
						X: 1,
						Y: 2,
					},
					{
						X: 3,
						Y: 4,
					},
				})
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 6 12\nL 8 14\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/></svg>",
		},
		// set background
		{
			fn: func(p *Painter) {
				p.SetBackground(400, 300, chart.ColorWhite)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 0 0\nL 400 0\nL 400 300\nL 0 300\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/></svg>",
		},
		// arcTo
		{
			fn: func(p *Painter) {
				p.SetStyle(Style{
					StrokeWidth: 1,
					StrokeColor: drawing.ColorBlack,
					FillColor:   drawing.ColorBlue,
				})
				p.ArcTo(100, 100, 100, 100, 0, math.Pi/2)
				p.Close()
				p.FillStroke()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 205 110\nA 100 100 90.00 0 1 105 210\nZ\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,255,1.0)\"/></svg>",
		},
		// pin
		{
			fn: func(p *Painter) {
				p.SetStyle(Style{
					StrokeWidth: 1,
					StrokeColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
					FillColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
				})
				p.Pin(30, 30, 30)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 32 47\nA 15 15 330.00 1 1 38 47\nL 35 33\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 20 33\nQ35,70 50,33\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/></svg>",
		},
		// arrow left
		{
			fn: func(p *Painter) {
				p.SetStyle(Style{
					StrokeWidth: 1,
					StrokeColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
					FillColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
				})
				p.ArrowLeft(30, 30, 16, 10)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 51 35\nL 35 40\nL 51 45\nL 46 40\nL 51 35\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/></svg>",
		},
		// arrow right
		{
			fn: func(p *Painter) {
				p.SetStyle(Style{
					StrokeWidth: 1,
					StrokeColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
					FillColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
				})
				p.ArrowRight(30, 30, 16, 10)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 19 35\nL 35 40\nL 19 45\nL 24 40\nL 19 35\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/></svg>",
		},
		// arrow top
		{
			fn: func(p *Painter) {
				p.SetStyle(Style{
					StrokeWidth: 1,
					StrokeColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
					FillColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
				})
				p.ArrowTop(30, 30, 10, 16)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 30 40\nL 35 24\nL 40 40\nL 35 35\nL 30 40\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/></svg>",
		},
		// arrow bottom
		{
			fn: func(p *Painter) {
				p.SetStyle(Style{
					StrokeWidth: 1,
					StrokeColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
					FillColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
				})
				p.ArrowBottom(30, 30, 10, 16)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 30 24\nL 35 40\nL 40 24\nL 35 30\nL 30 24\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/></svg>",
		},
		// mark line
		{
			fn: func(p *Painter) {
				p.SetStyle(Style{
					StrokeWidth: 1,
					StrokeColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
					FillColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
					StrokeDashArray: []float64{
						4,
						2,
					},
				})
				p.MarkLine(0, 20, 300)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<circle cx=\"8\" cy=\"30\" r=\"3\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 14 30\nL 289 30\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 289 25\nL 305 30\nL 289 35\nL 294 30\nL 289 25\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/></svg>",
		},
		// polygon
		{
			fn: func(p *Painter) {
				p.SetStyle(Style{
					StrokeWidth: 1,
					StrokeColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
				})
				p.Polygon(Point{
					X: 100,
					Y: 100,
				}, 50, 6)
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 105 60\nL 148 85\nL 148 134\nL 105 160\nL 62 135\nL 62 86\nL 105 60\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:none\"/></svg>",
		},
		// FillArea
		{
			fn: func(p *Painter) {
				p.SetDrawingStyle(Style{
					FillColor: Color{
						R: 84,
						G: 112,
						B: 198,
						A: 255,
					},
				})
				p.FillArea([]Point{
					{
						X: 0,
						Y: 0,
					},
					{
						X: 0,
						Y: 100,
					},
					{
						X: 100,
						Y: 100,
					},
					{
						X: 0,
						Y: 0,
					},
				})
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 5 10\nL 5 110\nL 105 110\nL 5 10\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/></svg>",
		},
	}
	for _, tt := range tests {
		d, err := NewPainter(PainterOptions{
			Width:  400,
			Height: 300,
			Type:   ChartOutputSVG,
		}, PainterPaddingOption(chart.Box{
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

func TestRoundedRect(t *testing.T) {
	assert := assert.New(t)
	p, err := NewPainter(PainterOptions{
		Width:  400,
		Height: 300,
		Type:   ChartOutputSVG,
	})
	assert.Nil(err)
	p.OverrideDrawingStyle(Style{
		FillColor:   drawing.ColorWhite,
		StrokeWidth: 1,
		StrokeColor: drawing.ColorWhite,
	}).RoundedRect(Box{
		Left:   10,
		Right:  30,
		Bottom: 150,
		Top:    10,
	}, 5)
	buf, err := p.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 15 10\nL 25 10\nL 25 10\nA 5 5 90.00 0 1 30 15\nL 30 145\nL 30 145\nA 5 5 90.00 0 1 25 150\nL 15 150\nL 15 150\nA 5 5 90.00 0 1 10 145\nL 10 15\nL 10 15\nA 5 5 90.00 0 1 15 10\nZ\" style=\"stroke-width:1;stroke:rgba(255,255,255,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(255,255,255,1.0);fill:rgba(255,255,255,1.0)\"/></svg>", string(buf))
}

func TestPainterTextFit(t *testing.T) {
	assert := assert.New(t)
	p, err := NewPainter(PainterOptions{
		Width:  400,
		Height: 300,
		Type:   ChartOutputSVG,
	})
	assert.Nil(err)
	f, _ := GetDefaultFont()
	style := Style{
		FontSize:  12,
		FontColor: chart.ColorBlack,
		Font:      f,
	}
	p.SetStyle(style)
	box := p.TextFit("Hello World!", 0, 20, 80)
	assert.Equal(chart.Box{
		Right:  45,
		Bottom: 35,
	}, box)

	box = p.TextFit("Hello World!", 0, 100, 200)
	assert.Equal(chart.Box{
		Right:  84,
		Bottom: 15,
	}, box)

	buf, err := p.Bytes()
	assert.Nil(err)
	assert.Equal(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="400" height="300">\n<text x="0" y="20" style="stroke-width:0;stroke:none;fill:rgba(51,51,51,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif">Hello</text><text x="0" y="40" style="stroke-width:0;stroke:none;fill:rgba(51,51,51,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif">World!</text><text x="0" y="100" style="stroke-width:0;stroke:none;fill:rgba(51,51,51,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif">Hello World!</text></svg>`, string(buf))
}
