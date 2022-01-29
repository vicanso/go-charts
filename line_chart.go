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

const YAxisWidth = 50

func drawXAxis(d *Draw, opt *XAxisOption, theme *Theme) (int, *Range, error) {
	dXAxis, err := NewDraw(
		DrawOption{
			Parent: d,
		},
		PaddingOption(chart.Box{
			Left: YAxisWidth,
		}),
	)
	if err != nil {
		return 0, nil, err
	}
	data := NewAxisDataListFromStringList(opt.Data)
	style := AxisStyle{
		BoundaryGap: opt.BoundaryGap,
		StrokeColor: theme.GetAxisStrokeColor(),
		FontColor:   theme.GetAxisStrokeColor(),
		StrokeWidth: 1,
	}

	boundary := true
	max := float64(len(opt.Data))
	if opt.BoundaryGap != nil && !*opt.BoundaryGap {
		boundary = false
		max--
	}

	dXAxis.Axis(data, style)
	return d.measureAxis(data, style), &Range{
		divideCount: len(opt.Data),
		Min:         0,
		Max:         max,
		Size:        dXAxis.Box.Width(),
		Boundary:    boundary,
	}, nil
}

func drawYAxis(d *Draw, opt *ChartOption, theme *Theme, xAxisHeight int) (*Range, error) {
	yRange := opt.getYRange(0)
	data := NewAxisDataListFromStringList(yRange.Values())
	style := AxisStyle{
		Position:    PositionLeft,
		BoundaryGap: FalseFlag(),
		// StrokeColor:    theme.GetAxisStrokeColor(),
		FontColor:      theme.GetAxisStrokeColor(),
		StrokeWidth:    1,
		SplitLineColor: theme.GetAxisSplitLineColor(),
		SplitLineShow:  true,
	}
	width := d.measureAxis(data, style)

	dYAxis, err := NewDraw(
		DrawOption{
			Parent: d,
			Width:  d.Box.Width(),
			// 减去x轴的高
			Height: d.Box.Height() - xAxisHeight,
		},
		PaddingOption(chart.Box{
			Left: YAxisWidth - width,
		}),
	)
	if err != nil {
		return nil, err
	}
	dYAxis.Axis(data, style)
	yRange.Size = dYAxis.Box.Height()
	return &yRange, nil
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
	// 设置背景色
	bg := opt.BackgroundColor
	if bg.IsZero() {
		bg = theme.GetBackgroundColor()
	}
	if opt.Parent == nil {
		d.setBackground(opt.getWidth(), opt.getHeight(), bg)
	}

	xAxisHeight, xRange, err := drawXAxis(d, &opt.XAxis, &theme)
	if err != nil {
		return nil, err
	}
	// 暂时仅支持单一yaxis
	yRange, err := drawYAxis(d, &opt.ChartOption, &theme, xAxisHeight)
	if err != nil {
		return nil, err
	}
	sd, err := NewDraw(DrawOption{
		Parent: d,
	}, PaddingOption(chart.Box{
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
	// fmt.Println(yRange)

	return d, nil
}
