// MIT License

// Copyright (c) 2021 Tree Xie

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
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
)

func TestLegendCustomize(t *testing.T) {
	assert := assert.New(t)

	series := GetSeries([]Series{
		{
			Name: "chrome",
		},
		{
			Name: "edge",
		},
	}, chart.TickPositionBetweenTicks, "")
	r, err := chart.SVG(800, 600)
	assert.Nil(err)
	fn := LegendCustomize(series, LegendOption{
		TextPosition: LegendTextPositionRight,
		IconDraw:     DefaultLegendIconDraw,
		Align:        LegendAlignLeft,
		Padding: chart.Box{
			Left: 100,
			Top:  100,
		},
	})
	fn(r, chart.NewBox(11, 47, 784, 373), chart.Style{
		Font: chart.StyleTextDefaults().Font,
	})
	buf := bytes.Buffer{}
	err = r.Save(&buf)
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"800\" height=\"600\">\\n<path  d=\"M 100 100\nL 208 100\nL 208 110\nL 100 110\nL 100 100\" style=\"stroke-width:0;stroke:rgba(51,51,51,1.0);fill:none\"/><path  d=\"M 100 107\nL 125 107\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:none\"/><circle cx=\"112\" cy=\"107\" r=\"5\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"130\" y=\"110\" style=\"stroke-width:0;stroke:none;fill:rgba(51,51,51,1.0);font-size:10.2px;font-family:'Roboto Medium',sans-serif\">chrome</text><path  d=\"M 185 107\nL 210 107\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"197\" cy=\"107\" r=\"5\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"215\" y=\"110\" style=\"stroke-width:0;stroke:none;fill:rgba(51,51,51,1.0);font-size:10.2px;font-family:'Roboto Medium',sans-serif\">edge</text></svg>", buf.String())
}
