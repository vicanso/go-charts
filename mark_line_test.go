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

func TestNewMarkLine(t *testing.T) {
	assert := assert.New(t)

	markLine := NewMarkLine(
		SeriesMarkDataTypeMax,
		SeriesMarkDataTypeMin,
		SeriesMarkDataTypeAverage,
	)

	assert.Equal(SeriesMarkLine{
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
	}, markLine)
}

func TestMarkLineRender(t *testing.T) {
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

	markLineRender(&markLineRenderOption{
		Draw:        d,
		FillColor:   drawing.ColorBlack,
		FontColor:   drawing.ColorBlack,
		StrokeColor: drawing.ColorBlack,
		Font:        f,
		Series: &Series{
			MarkLine: NewMarkLine(
				SeriesMarkDataTypeMax,
				SeriesMarkDataTypeMin,
				SeriesMarkDataTypeAverage,
			),
			Data: NewSeriesDataFromValues([]float64{
				1,
				3,
				5,
				7,
				9,
			}),
		},
		Range: &Range{
			Min:  0,
			Max:  10,
			Size: 200,
		},
	})
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<circle cx=\"20\" cy=\"20\" r=\"3\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 25 20\nL 362 20\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path  d=\"M 362 15\nL 378 20\nL 362 25\nL 367 20\nL 362 15\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><text x=\"380\" y=\"24\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">9</text><circle cx=\"20\" cy=\"180\" r=\"3\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 25 180\nL 362 180\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path  d=\"M 362 175\nL 378 180\nL 362 185\nL 367 180\nL 362 175\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><text x=\"380\" y=\"184\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">1</text><circle cx=\"20\" cy=\"100\" r=\"3\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 25 100\nL 362 100\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><path  d=\"M 362 95\nL 378 100\nL 362 105\nL 367 100\nL 362 95\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:rgba(0,0,0,1.0)\"/><text x=\"380\" y=\"104\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">5</text></svg>", string(data))
}
