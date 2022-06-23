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
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestGrid(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewGridPainter(p, GridPainterOption{
					StrokeColor:       drawing.ColorBlack,
					Column:            6,
					Row:               6,
					IgnoreFirstRow:    true,
					IgnoreLastRow:     true,
					IgnoreFirstColumn: true,
					IgnoreLastColumn:  true,
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 100 0\nL 100 400\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 200 0\nL 200 400\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 300 0\nL 300 400\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 400 0\nL 400 400\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 500 0\nL 500 400\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 66\nL 600 66\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 133\nL 600 133\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 200\nL 600 200\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 266\nL 600 266\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 333\nL 600 333\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/></svg>",
		},
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewGridPainter(p, GridPainterOption{
					StrokeColor: drawing.ColorBlack,
					ColumnSpans: []int{
						2,
						5,
						3,
					},
					Row: 6,
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 0 400\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 120 0\nL 120 400\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 420 0\nL 420 400\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 600 0\nL 600 400\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 0\nL 600 0\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 66\nL 600 66\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 133\nL 600 133\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 200\nL 600 200\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 266\nL 600 266\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 333\nL 600 333\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 400\nL 600 400\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/></svg>",
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
