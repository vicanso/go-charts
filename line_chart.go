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
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func lineChartRender(opt ChartOption, result *basicRenderResult) (*Draw, error) {

	d := result.d
	theme := NewTheme(opt.Theme)

	sd, err := NewDraw(DrawOption{
		Parent: d,
	}, PaddingOption(chart.Box{
		Top:  result.titleBox.Height(),
		Left: YAxisWidth,
	}))
	if err != nil {
		return nil, err
	}
	yRange := result.yRange
	xRange := result.xRange
	for i, series := range opt.SeriesList {
		points := make([]Point, 0)
		for j, item := range series.Data {
			y := yRange.getRestHeight(item.Value)
			points = append(points, Point{
				Y: y,
				X: xRange.getWidth(float64(j)),
			})
			index := series.index
			if index == 0 {
				index = i
			}
			seriesColor := theme.GetSeriesColor(index)
			dotFillColor := drawing.ColorWhite
			if theme.IsDark() {
				dotFillColor = seriesColor
			}
			sd.Line(points, LineStyle{
				StrokeColor:  seriesColor,
				StrokeWidth:  2,
				DotColor:     seriesColor,
				DotWidth:     2,
				DotFillColor: dotFillColor,
			})
		}
	}

	return d, nil
}
