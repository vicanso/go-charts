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

func TestGetPieRadius(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(50.0, getPieRadius(100, "50%"))
	assert.Equal(30.0, getPieRadius(100, "30"))
	assert.Equal(40.0, getPieRadius(100, ""))
}

func TestPieChartRender(t *testing.T) {
	assert := assert.New(t)

	d, err := NewDraw(DrawOption{
		Width:  250,
		Height: 150,
	})
	assert.Nil(err)

	f, _ := chart.GetDefaultFont()

	err = pieChartRender(pieChartOption{
		Font: f,
		SeriesList: NewPieSeriesList([]float64{
			5,
			10,
		}, PieSeriesOption{
			Names: []string{
				"a",
				"b",
			},
			Label: SeriesLabel{
				Show:  true,
				Color: drawing.ColorRed,
			},
			Radius: "20%",
		}),
	}, &basicRenderResult{
		d: d,
	})
	assert.Nil(err)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"250\" height=\"150\">\\n<path  d=\"M 125 75\nL 125 45\nA 30 30 120.00 0 1 150 89\nL 125 75\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 150 60\nL 159 55\nM 159 55\nL 169 55\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"172\" y=\"60\" style=\"stroke-width:0;stroke:none;fill:rgba(255,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">a: 33.33%</text><path  d=\"M 125 75\nL 150 89\nA 30 30 240.00 1 1 125 45\nL 125 75\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 100 90\nL 91 95\nM 91 95\nL 81 95\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"22\" y=\"100\" style=\"stroke-width:0;stroke:none;fill:rgba(255,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">b: 66.66%</text></svg>", string(data))
}
