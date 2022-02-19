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

func TestLineChartRender(t *testing.T) {
	assert := assert.New(t)

	width := 400
	height := 300
	d, err := NewDraw(DrawOption{
		Width:  width,
		Height: height,
	})
	assert.Nil(err)

	result := basicRenderResult{
		xRange: &Range{
			Min:         0,
			Max:         4,
			divideCount: 4,
			Size:        width,
			Boundary:    true,
		},
		yRangeList: []*Range{
			{
				divideCount: 6,
				Max:         100,
				Min:         0,
				Size:        height,
			},
		},
		d: d,
	}
	f, _ := chart.GetDefaultFont()
	_, err = lineChartRender(lineChartOption{
		Font: f,
		SeriesList: SeriesList{
			{
				Label: SeriesLabel{
					Show:  true,
					Color: drawing.ColorBlue,
				},
				MarkLine: NewMarkLine(
					SeriesMarkDataTypeAverage,
				),
				Data: []SeriesData{
					{
						Value: 20,
					},
					{
						Value: 60,
					},
					{
						Value: 90,
					},
				},
			},
			NewSeriesFromValues([]float64{
				40,
				60,
				70,
			}),
		},
	}, &result)
	assert.Nil(err)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<circle cx=\"40\" cy=\"130\" r=\"3\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 45 130\nL 382 130\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 382 125\nL 398 130\nL 382 135\nL 387 130\nL 382 125\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"400\" y=\"134\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">56.66</text><text x=\"83\" y=\"235\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">20</text><text x=\"183\" y=\"115\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">60</text><text x=\"283\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">90</text><path  d=\"M 90 240\nL 190 120\nL 290 30\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:none\"/><circle cx=\"90\" cy=\"240\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"190\" cy=\"120\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"290\" cy=\"30\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"M 90 180\nL 190 120\nL 290 90\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:none\"/><circle cx=\"90\" cy=\"180\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"190\" cy=\"120\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"290\" cy=\"90\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/></svg>", string(data))
}
