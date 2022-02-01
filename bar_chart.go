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
)

func BarChartRender(opt ChartOption) (*Draw, error) {
	result, err := chartBasicRender(&opt)
	if err != nil {
		return nil, err
	}
	d := result.d

	bd, err := NewDraw(DrawOption{
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
	x0, x1 := xRange.GetRange(0)
	width := int(x1 - x0)
	// 每一块之间的margin
	margin := 10
	// 每一个bar之间的margin
	barMargin := 5

	seriesCount := len(opt.SeriesList)
	// 总的宽度-两个margin-(总数-1)的barMargin
	barWidth := (width - 2*margin - barMargin*(seriesCount-1)) / len(opt.SeriesList)

	barMaxHeight := yRange.Size
	theme := NewTheme(opt.Theme)

	for i, series := range opt.SeriesList {
		for j, item := range series.Data {
			x0, _ := xRange.GetRange(j)
			x := int(x0)
			x += margin
			if i != 0 {
				x += i * (barWidth + barMargin)
			}

			h := int(yRange.getHeight(item.Value))

			bd.Bar(chart.Box{
				Top:    barMaxHeight - h,
				Left:   x,
				Right:  x + barWidth,
				Bottom: barMaxHeight - 1,
			}, BarStyle{
				FillColor: theme.GetSeriesColor(i),
			})
		}
	}

	return d, nil
}
