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

func TestFunnelChartRender(t *testing.T) {
	assert := assert.New(t)

	d, err := NewDraw(DrawOption{
		Width:  250,
		Height: 150,
	})
	assert.Nil(err)
	f, _ := chart.GetDefaultFont()
	err = funnelChartRender(funnelChartOption{
		Font: f,
		SeriesList: []Series{
			{
				Type: ChartTypeFunnel,
				Name: "Visit",
				Data: NewSeriesDataFromValues([]float64{
					60,
				}),
			},
			{
				Type: ChartTypeFunnel,
				Name: "Inquiry",
				Data: NewSeriesDataFromValues([]float64{
					40,
				}),
				index: 1,
			},
			{
				Type: ChartTypeFunnel,
				Name: "Order",
				Data: NewSeriesDataFromValues([]float64{
					20,
				}),
				index: 2,
			},
			{
				Type: ChartTypeFunnel,
				Name: "Click",
				Data: NewSeriesDataFromValues([]float64{
					80,
				}),
				index: 3,
			},
			{
				Type: ChartTypeFunnel,
				Name: "Show",
				Data: NewSeriesDataFromValues([]float64{
					100,
				}),
				index: 4,
			},
		},
	}, &basicRenderResult{
		d: d,
	})
	assert.Nil(err)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"250\" height=\"150\">\\n<path  d=\"M 0 0\nL 250 0\nL 225 28\nL 25 28\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(115,192,222,1.0)\"/><text x=\"89\" y=\"14\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Show(100%)</text><path  d=\"M 25 30\nL 225 30\nL 200 58\nL 50 58\nL 25 30\" style=\"stroke-width:0;stroke:none;fill:rgba(238,102,102,1.0)\"/><text x=\"94\" y=\"44\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Click(80%)</text><path  d=\"M 50 60\nL 200 60\nL 175 88\nL 75 88\nL 50 60\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><text x=\"96\" y=\"74\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Visit(60%)</text><path  d=\"M 75 90\nL 175 90\nL 150 118\nL 100 118\nL 75 90\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><text x=\"89\" y=\"104\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Inquiry(40%)</text><path  d=\"M 100 120\nL 150 120\nL 125 148\nL 125 148\nL 100 120\" style=\"stroke-width:0;stroke:none;fill:rgba(250,200,88,1.0)\"/><text x=\"93\" y=\"134\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Order(20%)</text></svg>", string(data))
}
