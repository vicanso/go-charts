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

func TestRadarChart(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				values := [][]float64{
					{
						4200,
						3000,
						20000,
						35000,
						50000,
						18000,
					},
					{
						5000,
						14000,
						28000,
						26000,
						42000,
						21000,
					},
				}
				_, err := NewRadarChart(p, RadarChartOption{
					SeriesList: NewSeriesListDataFromValues(values, ChartTypeRadar),
					Title: TitleOption{
						Text: "Basic Radar Chart",
					},
					Legend: NewLegendOption([]string{
						"Allocated Budget",
						"Actual Spending",
					}),
					RadarIndicators: NewRadarIndicators([]string{
						"Sales",
						"Administration",
						"Information Technology",
						"Customer Support",
						"Development",
						"Marketing",
					}, []float64{
						6500,
						16000,
						30000,
						38000,
						52000,
						25000,
					}),
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 560 0\nL 560 360\nL 0 360\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 143 29\nL 173 29\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><circle cx=\"158\" cy=\"29\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"175\" y=\"35\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Allocated Budget</text><path  d=\"M 313 29\nL 343 29\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><circle cx=\"328\" cy=\"29\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"345\" y=\"35\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Actual Spending</text><text x=\"20\" y=\"35\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Basic Radar Chart</text><path  d=\"M 300 179\nL 319 191\nL 319 213\nL 300 225\nL 281 213\nL 281 191\nL 300 179\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 156\nL 339 179\nL 339 224\nL 300 248\nL 261 225\nL 261 180\nL 300 156\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 133\nL 359 168\nL 359 236\nL 300 271\nL 241 236\nL 241 168\nL 300 133\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 110\nL 379 156\nL 379 247\nL 300 294\nL 221 248\nL 221 157\nL 300 110\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 87\nL 399 145\nL 399 259\nL 300 317\nL 201 259\nL 201 145\nL 300 87\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 202\nL 300 87\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 202\nL 399 145\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 202\nL 399 259\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 202\nL 300 317\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 202\nL 201 259\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 202\nL 201 145\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><text x=\"284\" y=\"80\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sales</text><text x=\"404\" y=\"150\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Administration</text><text x=\"404\" y=\"264\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Information Technology</text><text x=\"248\" y=\"334\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Customer Support</text><text x=\"120\" y=\"264\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Development</text><text x=\"137\" y=\"150\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Marketing</text><path  d=\"M 300 128\nL 318 192\nL 366 240\nL 300 307\nL 205 257\nL 229 161\nL 300 128\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,0.1)\"/><path  d=\"M 300 128\nL 318 192\nL 366 240\nL 300 307\nL 205 257\nL 229 161\nL 300 128\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,0.1)\"/><circle cx=\"300\" cy=\"128\" r=\"2\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"318\" cy=\"192\" r=\"2\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"366\" cy=\"240\" r=\"2\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"300\" cy=\"307\" r=\"2\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"205\" cy=\"257\" r=\"2\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"229\" cy=\"161\" r=\"2\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"300\" cy=\"128\" r=\"2\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"M 300 114\nL 387 152\nL 392 255\nL 300 280\nL 220 248\nL 217 154\nL 300 114\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,0.1)\"/><path  d=\"M 300 114\nL 387 152\nL 392 255\nL 300 280\nL 220 248\nL 217 154\nL 300 114\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,0.1)\"/><circle cx=\"300\" cy=\"114\" r=\"2\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"387\" cy=\"152\" r=\"2\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"392\" cy=\"255\" r=\"2\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"300\" cy=\"280\" r=\"2\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"220\" cy=\"248\" r=\"2\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"217\" cy=\"154\" r=\"2\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"300\" cy=\"114\" r=\"2\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/></svg>",
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
