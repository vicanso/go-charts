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

func TestNewLegend(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewLegendPainter(p, LegendOption{
					Data: []string{
						"One",
						"Two",
						"Three",
					},
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 184 9\nL 214 9\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><circle cx=\"199\" cy=\"9\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"216\" y=\"15\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">One</text><path  d=\"M 264 9\nL 294 9\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><circle cx=\"279\" cy=\"9\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"296\" y=\"15\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Two</text><path  d=\"M 346 9\nL 376 9\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><circle cx=\"361\" cy=\"9\" r=\"5\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"378\" y=\"15\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Three</text></svg>",
		},
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewLegendPainter(p, LegendOption{
					Data: []string{
						"One",
						"Two",
						"Three",
					},
					Left: PositionLeft,
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 9\nL 30 9\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><circle cx=\"15\" cy=\"9\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"32\" y=\"15\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">One</text><path  d=\"M 80 9\nL 110 9\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><circle cx=\"95\" cy=\"9\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"112\" y=\"15\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Two</text><path  d=\"M 162 9\nL 192 9\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><circle cx=\"177\" cy=\"9\" r=\"5\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"194\" y=\"15\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Three</text></svg>",
		},
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewLegendPainter(p, LegendOption{
					Data: []string{
						"One",
						"Two",
						"Three",
					},
					Orient: OrientVertical,
					Icon:   IconRect,
					Left:   "10%",
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 60 3\nL 90 3\nL 90 16\nL 60 16\nL 60 3\" style=\"stroke-width:0;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"92\" y=\"15\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">One</text><path  d=\"M 60 23\nL 90 23\nL 90 36\nL 60 36\nL 60 23\" style=\"stroke-width:0;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"92\" y=\"35\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Two</text><path  d=\"M 60 43\nL 90 43\nL 90 56\nL 60 56\nL 60 43\" style=\"stroke-width:0;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"92\" y=\"55\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Three</text></svg>",
		},
	}
	for _, tt := range tests {
		p, err := NewPainter(PainterOptions{
			Type:   ChartOutputSVG,
			Width:  600,
			Height: 400,
		}, PainterThemeOption(defaultTheme))
		assert.Nil(err)
		data, err := tt.render(p)
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}
