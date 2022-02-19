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

func TestBarChartRender(t *testing.T) {
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

	markPointOptions, err := barChartRender(barChartOption{
		Font: f,
		SeriesList: SeriesList{
			{
				Label: SeriesLabel{
					Show:  true,
					Color: drawing.ColorBlue,
				},
				MarkLine: NewMarkLine(
					SeriesMarkDataTypeMin,
				),
				Data: []SeriesData{
					{
						Value: 20,
					},
					{
						Value: 60,
						Style: chart.Style{
							FillColor: drawing.ColorRed,
						},
					},
					{
						Value: 90,
					},
				},
			},
			NewSeriesFromValues([]float64{
				80,
				30,
				70,
			}),
		},
	}, &result)
	assert.Nil(err)
	assert.Equal(2, len(markPointOptions))
	assert.Equal([]Point{
		{
			X: 28,
			Y: 240,
		},
		{
			X: 128,
			Y: 120,
		},
		{
			X: 228,
			Y: 30,
		},
	}, markPointOptions[0].Points)
	assert.Equal([]Point{
		{
			X: 70,
			Y: 60,
		},
		{
			X: 170,
			Y: 210,
		},
		{
			X: 270,
			Y: 90,
		},
	}, markPointOptions[1].Points)

	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<circle cx=\"40\" cy=\"240\" r=\"3\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path stroke-dasharray=\"4.0, 2.0\" d=\"M 45 240\nL 382 240\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 382 235\nL 398 240\nL 382 245\nL 387 240\nL 382 235\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"400\" y=\"244\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">20</text><path  d=\"M 50 240\nL 87 240\nL 87 299\nL 50 299\nL 50 240\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"61\" y=\"235\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">20</text><path  d=\"M 150 120\nL 187 120\nL 187 299\nL 150 299\nL 150 120\" style=\"stroke-width:1;stroke:rgba(255,0,0,1.0);fill:rgba(255,0,0,1.0)\"/><text x=\"161\" y=\"115\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">60</text><path  d=\"M 250 30\nL 287 30\nL 287 299\nL 250 299\nL 250 30\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"261\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">90</text><path  d=\"M 92 60\nL 129 60\nL 129 299\nL 92 299\nL 92 60\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 192 210\nL 229 210\nL 229 299\nL 192 299\nL 192 210\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 292 90\nL 329 90\nL 329 299\nL 292 299\nL 292 90\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/></svg>", string(data))
}
