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
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestBarSeries(t *testing.T) {
	assert := assert.New(t)

	customStyle := chart.Style{
		StrokeColor: drawing.ColorBlue,
	}
	bs := BarSeries{
		CustomStyles: []BarSeriesCustomStyle{
			{
				PointIndex: 1,
				Style:      customStyle,
			},
		},
	}

	assert.Equal(customStyle, bs.GetBarStyle(0, 1))

	assert.True(bs.GetBarStyle(1, 0).IsZero())
}

func TestBarSeriesGetWidthValues(t *testing.T) {
	assert := assert.New(t)

	bs := BarSeries{
		Count: 1,
		BaseSeries: BaseSeries{
			XValues: []float64{
				1,
				2,
				3,
			},
		},
	}
	widthValues := bs.getWidthValues(300)
	assert.Equal(barSeriesWidthValues{
		columnWidth:  100,
		columnMargin: 10,
		margin:       10,
		barWidth:     80,
	}, widthValues)

	// 指定margin
	bs.Margin = 5
	widthValues = bs.getWidthValues(300)
	assert.Equal(barSeriesWidthValues{
		columnWidth:  100,
		columnMargin: 10,
		margin:       5,
		barWidth:     80,
	}, widthValues)

	// 指定bar的宽度
	bs.BarWidth = 60
	widthValues = bs.getWidthValues(300)
	assert.Equal(barSeriesWidthValues{
		columnWidth:  100,
		columnMargin: 20,
		margin:       5,
		barWidth:     60,
	}, widthValues)
}

func TestBarSeriesRender(t *testing.T) {
	assert := assert.New(t)

	width := 800
	height := 400

	r, err := chart.SVG(width, height)
	assert.Nil(err)

	bs := BarSeries{
		Count: 1,
		CustomStyles: []BarSeriesCustomStyle{
			{
				Index:      0,
				PointIndex: 1,
				Style: chart.Style{
					StrokeColor: SeriesColorsLight[1],
				},
			},
		},
		BaseSeries: BaseSeries{
			TickPosition: chart.TickPositionBetweenTicks,
			Style: chart.Style{
				StrokeColor: SeriesColorsLight[0],
				StrokeWidth: 1,
			},
			XValues: []float64{
				0,
				1,
				2,
				3,
				4,
				5,
				6,
				7,
			},
			YValues: []float64{
				// 第一个点为占位点
				0,
				120,
				200,
				150,
				80,
				70,
				110,
				130,
			},
		},
	}
	xrange := &chart.ContinuousRange{
		Min:    0,
		Max:    7,
		Domain: 753,
	}
	yrange := &chart.ContinuousRange{
		Min:    70,
		Max:    200,
		Domain: 362,
	}
	bs.Render(r, chart.Box{
		Top:    11,
		Left:   42,
		Right:  795,
		Bottom: 373,
	}, xrange, yrange, chart.Style{})

	buffer := bytes.Buffer{}
	err = r.Save(&buffer)
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"800\" height=\"400\">\\n<path  d=\"M 53 233\nL 140 233\nL 140 372\nL 53 372\nL 53 233\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 161 11\nL 248 11\nL 248 372\nL 161 372\nL 161 11\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:none\"/><path  d=\"M 268 150\nL 355 150\nL 355 372\nL 268 372\nL 268 150\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 376 345\nL 463 345\nL 463 372\nL 376 372\nL 376 345\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 483 373\nL 570 373\nL 570 372\nL 483 372\nL 483 373\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 591 261\nL 678 261\nL 678 372\nL 591 372\nL 591 261\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 698 205\nL 785 205\nL 785 372\nL 698 372\nL 698 205\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/></svg>", buffer.String())
}
