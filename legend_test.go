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

func TestNewLegendOption(t *testing.T) {
	assert := assert.New(t)

	opt := NewLegendOption([]string{
		"a",
		"b",
	}, PositionRight)
	assert.Equal(LegendOption{
		Data: []string{
			"a",
			"b",
		},
		Left: PositionRight,
	}, opt)
}

func TestLegendRender(t *testing.T) {
	assert := assert.New(t)

	newDraw := func() *Draw {
		d, _ := NewDraw(DrawOption{
			Width:  400,
			Height: 300,
		})
		return d
	}
	style := chart.Style{
		FontSize:  10,
		FontColor: drawing.ColorBlack,
	}
	style.Font, _ = chart.GetDefaultFont()

	tests := []struct {
		newDraw   func() *Draw
		newLegend func(*Draw) *legend
		box       chart.Box
		result    string
	}{
		{
			newDraw: newDraw,
			newLegend: func(d *Draw) *legend {
				return NewLegend(d, LegendOption{
					Top: "10",
					Data: []string{
						"Mon",
						"Tue",
						"Wed",
					},
					Style: style,
				})
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 0 20\nL 30 20\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"15\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"35\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><path  d=\"M 76 20\nL 106 20\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"91\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"111\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><path  d=\"M 148 20\nL 178 20\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"163\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"183\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text></svg>",
			box: chart.Box{
				Right:  214,
				Bottom: 25,
			},
		},
		{
			newDraw: newDraw,
			newLegend: func(d *Draw) *legend {
				return NewLegend(d, LegendOption{
					Top:   "10",
					Left:  PositionRight,
					Align: PositionRight,
					Data: []string{
						"Mon",
						"Tue",
						"Wed",
					},
					Style: style,
				})
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<text x=\"191\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><path  d=\"M 222 20\nL 252 20\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"237\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"267\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><path  d=\"M 294 20\nL 324 20\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"309\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"339\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text><path  d=\"M 370 20\nL 400 20\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"385\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/></svg>",
			box: chart.Box{
				Right:  400,
				Bottom: 25,
			},
		},
		{
			newDraw: newDraw,
			newLegend: func(d *Draw) *legend {
				return NewLegend(d, LegendOption{
					Top:  "10",
					Left: PositionCenter,
					Data: []string{
						"Mon",
						"Tue",
						"Wed",
					},
					Style: style,
				})
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 93 20\nL 123 20\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"108\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"128\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><path  d=\"M 169 20\nL 199 20\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"184\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"204\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><path  d=\"M 241 20\nL 271 20\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"256\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"276\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text></svg>",
			box: chart.Box{
				Right:  307,
				Bottom: 25,
			},
		},
		{
			newDraw: newDraw,
			newLegend: func(d *Draw) *legend {
				return NewLegend(d, LegendOption{
					Top:  "10",
					Left: PositionLeft,
					Data: []string{
						"Mon",
						"Tue",
						"Wed",
					},
					Style:  style,
					Orient: OrientVertical,
				})
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 0 20\nL 30 20\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"15\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"35\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><path  d=\"M 0 40\nL 30 40\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"15\" cy=\"40\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"35\" y=\"45\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><path  d=\"M 0 60\nL 30 60\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"15\" cy=\"60\" r=\"5\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"35\" y=\"65\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text></svg>",
			box: chart.Box{
				Right:  61,
				Bottom: 80,
			},
		},
		{
			newDraw: newDraw,
			newLegend: func(d *Draw) *legend {
				return NewLegend(d, LegendOption{
					Top:  "10",
					Left: "10%",
					Data: []string{
						"Mon",
						"Tue",
						"Wed",
					},
					Style:  style,
					Orient: OrientVertical,
				})
			},
			box: chart.Box{
				Right:  101,
				Bottom: 80,
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 40 20\nL 70 20\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"55\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"75\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><path  d=\"M 40 40\nL 70 40\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"55\" cy=\"40\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"75\" y=\"45\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><path  d=\"M 40 60\nL 70 60\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"55\" cy=\"60\" r=\"5\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"75\" y=\"65\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text></svg>",
		},
	}

	for _, tt := range tests {
		d := tt.newDraw()
		b, err := tt.newLegend(d).Render()
		assert.Nil(err)
		assert.Equal(tt.box, b)
		data, err := d.Bytes()
		fmt.Println(string(data))
		assert.Nil(err)
		assert.NotEmpty(data)
		assert.Equal(tt.result, string(data))
	}
}
