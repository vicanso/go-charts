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

type LineChartOption struct {
	ChartOption
}

func NewLineChart(opt LineChartOption) (*Draw, error) {
	d, err := NewDraw(
		DrawOption{
			Parent: opt.Parent,
			Width:  opt.getWidth(),
			Height: opt.getHeight(),
		},
		PaddingOption(opt.Padding),
	)
	if err != nil {
		return nil, err
	}

	theme := Theme{
		mode: opt.Theme,
	}
	opt.FillDefault(&theme)
	if opt.Parent == nil {
		d.setBackground(opt.getWidth(), opt.getHeight(), opt.BackgroundColor)
	}

	// 标题
	titleBox, err := drawTitle(d, &opt.Title)
	if err != nil {
		return nil, err
	}

	opt.Legend.Style.Padding.Left += titleBox.Right
	_, err = drawLegend(d, &opt.Legend, &theme)
	if err != nil {
		return nil, err
	}

	// xAxis
	xAxisHeight, xRange, err := drawXAxis(d, &opt.XAxis, &theme)
	if err != nil {
		return nil, err
	}

	// 暂时仅支持单一yaxis
	yRange, err := drawYAxis(d, &opt.ChartOption, &theme, xAxisHeight, chart.Box{
		Top: titleBox.Height(),
	})
	if err != nil {
		return nil, err
	}

	sd, err := NewDraw(DrawOption{
		Parent: d,
	}, PaddingOption(chart.Box{
		Top:  titleBox.Height(),
		Left: YAxisWidth,
	}))
	if err != nil {
		return nil, err
	}
	for i, series := range opt.SeriesList {
		points := make([]Point, 0)
		for j, item := range series.Data {
			y := yRange.getHeight(item.Value)
			points = append(points, Point{
				Y: y,
				X: xRange.getWidth(float64(j)),
			})
			seriesColor := theme.GetSeriesColor(i)
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
