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

func TestMarkLine(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				markLine := NewMarkLinePainter(p)
				series := NewSeriesFromValues([]float64{
					1,
					2,
					3,
				})
				series.MarkLine = NewMarkLine(
					SeriesMarkDataTypeMax,
					SeriesMarkDataTypeAverage,
					SeriesMarkDataTypeMin,
				)
				markLine.Add(markLineRenderOption{
					FillColor:   drawing.ColorBlack,
					FontColor:   drawing.ColorBlack,
					StrokeColor: drawing.ColorBlack,
					Series:      series,
					Range: NewRange(AxisRangeOption{
						Min:         0,
						Max:         5,
						Size:        p.Height(),
						DivideCount: 6,
					}),
				})
				_, err := markLine.Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<circle cx=\"23\" cy=\"290\" r=\"3\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 29 290\nL 562 290\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 562 285\nL 578 290\nL 562 295\nL 567 290\nL 562 285\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><text x=\"580\" y=\"294\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">3</text><circle cx=\"23\" cy=\"320\" r=\"3\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 29 320\nL 562 320\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 562 315\nL 578 320\nL 562 325\nL 567 320\nL 562 315\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><text x=\"580\" y=\"324\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2</text><circle cx=\"23\" cy=\"350\" r=\"3\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 29 350\nL 562 350\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 562 345\nL 578 350\nL 562 355\nL 567 350\nL 562 345\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><text x=\"580\" y=\"354\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">1</text></svg>",
		},
	}
	for _, tt := range tests {
		p, err := NewPainter(PainterOptions{
			Type:   ChartOutputSVG,
			Width:  600,
			Height: 400,
		}, PainterThemeOption(defaultTheme))
		assert.Nil(err)
		data, err := tt.render(p.Child(PainterPaddingOption(Box{
			Left:   20,
			Top:    20,
			Right:  20,
			Bottom: 20,
		})))
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}
