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
)

func TestDrawYAxis(t *testing.T) {
	assert := assert.New(t)
	newDraw := func() *Draw {
		d, _ := NewDraw(DrawOption{
			Width:  400,
			Height: 300,
		})
		return d
	}

	tests := []struct {
		newDraw     func() *Draw
		newOption   func() *ChartOption
		xAxisHeight int
		result      string
	}{
		{
			newDraw: newDraw,
			newOption: func() *ChartOption {
				return &ChartOption{
					YAxis: YAxisOption{
						Max: toFloatPoint(20),
					},
					SeriesList: []Series{
						{
							Data: []SeriesData{
								{
									Value: 1,
								},
								{
									Value: 2,
								},
							},
						},
					},
				}
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 50 10\nL 50 290\" style=\"stroke-width:1;stroke:none;fill:none\"/><path  d=\"M 50 10\nL 390 10\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 57\nL 390 57\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 104\nL 390 104\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 151\nL 390 151\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 198\nL 390 198\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 244\nL 390 244\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><text x=\"36\" y=\"294\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">0</text><text x=\"18\" y=\"248\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">3.33</text><text x=\"18\" y=\"202\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">6.66</text><text x=\"29\" y=\"155\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">10</text><text x=\"11\" y=\"108\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">13.33</text><text x=\"11\" y=\"61\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">16.66</text><text x=\"29\" y=\"14\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">20</text></svg>",
		},
	}

	for _, tt := range tests {
		d := tt.newDraw()
		r, err := drawYAxis(d, tt.newOption(), tt.xAxisHeight, chart.NewBox(10, 10, 10, 10))
		assert.Nil(err)
		assert.Equal(&Range{
			divideCount: 6,
			Max:         20,
			Size:        280,
		}, r)

		data, err := d.Bytes()
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}
