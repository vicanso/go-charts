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

func TestNewLegendCustomize(t *testing.T) {
	assert := assert.New(t)

	series := GetSeries([]Series{
		{
			Name: "chrome",
		},
		{
			Name: "edge",
		},
	}, chart.TickPositionBetweenTicks, "")

	tests := []struct {
		align string
		svg   string
	}{
		{
			align: LegendAlignLeft,
			svg:   "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"800\" height=\"600\">\\n<path  d=\"M 449 111\nL 582 111\nL 582 121\nL 449 121\nL 449 111\" style=\"stroke-width:0;stroke:rgba(51,51,51,1.0);fill:none\"/><path  d=\"M 449 118\nL 474 118\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:none\"/><circle cx=\"461\" cy=\"118\" r=\"5\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"479\" y=\"121\" style=\"stroke-width:0;stroke:none;fill:rgba(51,51,51,1.0);font-size:10.2px;font-family:'Roboto Medium',sans-serif\">chrome</text><path  d=\"M 534 118\nL 559 118\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"546\" cy=\"118\" r=\"5\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"564\" y=\"121\" style=\"stroke-width:0;stroke:none;fill:rgba(51,51,51,1.0);font-size:10.2px;font-family:'Roboto Medium',sans-serif\">edge</text></svg>",
		},
		{
			align: LegendAlignRight,
			svg:   "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"800\" height=\"600\">\\n<path  d=\"M 449 111\nL 582 111\nL 582 121\nL 449 121\nL 449 111\" style=\"stroke-width:0;stroke:rgba(51,51,51,1.0);fill:none\"/><text x=\"449\" y=\"121\" style=\"stroke-width:0;stroke:none;fill:rgba(51,51,51,1.0);font-size:10.2px;font-family:'Roboto Medium',sans-serif\">chrome</text><path  d=\"M 489 118\nL 514 118\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:none\"/><circle cx=\"501\" cy=\"118\" r=\"5\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"539\" y=\"121\" style=\"stroke-width:0;stroke:none;fill:rgba(51,51,51,1.0);font-size:10.2px;font-family:'Roboto Medium',sans-serif\">edge</text><path  d=\"M 567 118\nL 592 118\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"579\" cy=\"118\" r=\"5\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/></svg>",
		},
	}

	for _, tt := range tests {
		r, err := chart.SVG(800, 600)
		assert.Nil(err)
		fn := NewLegendCustomize(series, LegendOption{
			Align:    tt.align,
			IconDraw: DefaultLegendIconDraw,
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
		assert.Equal(tt.svg, buf.String())
	}
}

func TestConvertPercent(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(-1.0, convertPercent("12"))

	assert.Equal(0.12, convertPercent("12%"))
}

func TestGetLegendLeft(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(150, getLegendLeft(500, 200, LegendOption{}))

	assert.Equal(0, getLegendLeft(500, 200, LegendOption{
		Left: "left",
	}))
	assert.Equal(100, getLegendLeft(500, 200, LegendOption{
		Left: "20%",
	}))
	assert.Equal(20, getLegendLeft(500, 200, LegendOption{
		Left: "20",
	}))

	assert.Equal(300, getLegendLeft(500, 200, LegendOption{
		Right: "right",
	}))
	assert.Equal(200, getLegendLeft(500, 200, LegendOption{
		Right: "20%",
	}))
	assert.Equal(280, getLegendLeft(500, 200, LegendOption{
		Right: "20",
	}))

}
