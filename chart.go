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
	"errors"
	"math"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const (
	ChartTypeLine = "line"
	ChartTypeBar  = "bar"
	ChartTypePie  = "pie"
)

type Point struct {
	X int
	Y int
}

type ChartOption struct {
	Font            *truetype.Font
	Theme           string
	Title           TitleOption
	Legend          LegendOption
	XAxis           XAxisOption
	YAxis           YAxisOption
	Width           int
	Height          int
	Parent          *Draw
	Padding         chart.Box
	SeriesList      []Series
	BackgroundColor drawing.Color
	Children        []ChartOption
}

func (o *ChartOption) FillDefault(theme string) {
	t := NewTheme(theme)
	if o.Font == nil {
		o.Font, _ = chart.GetDefaultFont()
	}
	if o.BackgroundColor.IsZero() {
		o.BackgroundColor = t.GetBackgroundColor()
	}

	// 标题的默认值
	if o.Title.Style.FontColor.IsZero() {
		o.Title.Style.FontColor = t.GetTextColor()
	}
	if o.Title.Style.FontSize == 0 {
		o.Title.Style.FontSize = 14
	}
	if o.Title.Style.Font == nil {
		o.Title.Style.Font = o.Font
	}
	if o.Title.Style.Padding.IsZero() {
		o.Title.Style.Padding = chart.Box{
			Left:   5,
			Top:    5,
			Right:  5,
			Bottom: 5,
		}
	}
	// 副标题
	if o.Title.SubtextStyle.FontColor.IsZero() {
		o.Title.SubtextStyle.FontColor = o.Title.Style.FontColor.WithAlpha(180)
	}
	if o.Title.SubtextStyle.FontSize == 0 {
		o.Title.SubtextStyle.FontSize = 10
	}
	if o.Title.SubtextStyle.Font == nil {
		o.Title.SubtextStyle.Font = o.Font
	}

	o.Legend.Theme = theme
	if o.Legend.Style.FontSize == 0 {
		o.Legend.Style.FontSize = 10
	}
	if o.Legend.Left == "" {
		o.Legend.Left = PositionCenter
	}
	if len(o.Legend.Data) == 0 {
		names := make([]string, len(o.SeriesList))
		for index, item := range o.SeriesList {
			names[index] = item.Name
		}
		o.Legend.Data = names
	}
	if o.Legend.Style.Font == nil {
		o.Legend.Style.Font = o.Font
	}
	if o.Legend.Style.FontColor.IsZero() {
		o.Legend.Style.FontColor = t.GetTextColor()
	}
	if o.XAxis.Theme == "" {
		o.XAxis.Theme = theme
	}
}

func (o *ChartOption) getWidth() int {
	if o.Width == 0 {
		return 600
	}
	return o.Width
}

func (o *ChartOption) getHeight() int {
	if o.Height == 0 {
		return 400
	}
	return o.Height
}

func (o *ChartOption) getYRange(axisIndex int) Range {
	min := math.MaxFloat64
	max := -math.MaxFloat64

	for _, series := range o.SeriesList {
		if series.YAxisIndex != axisIndex {
			continue
		}
		for _, item := range series.Data {
			if item.Value > max {
				max = item.Value
			}
			if item.Value < min {
				min = item.Value
			}
		}
	}
	min = min * 0.9
	max = max * 1.1
	if o.YAxis.Min != nil {
		min = *o.YAxis.Min
	}
	if o.YAxis.Max != nil {
		max = *o.YAxis.Max
	}
	divideCount := 6
	// y轴分设置默认划分为6块
	r := NewRange(min, max, divideCount)

	// 由于NewRange会重新计算min max
	if o.YAxis.Min != nil {
		r.Min = min
	}
	if o.YAxis.Max != nil {
		r.Max = max
	}

	return r
}

type basicRenderResult struct {
	xRange   *Range
	yRange   *Range
	d        *Draw
	titleBox chart.Box
}

func Render(opt ChartOption) (*Draw, error) {
	if len(opt.SeriesList) == 0 {
		return nil, errors.New("series can not be nil")
	}

	lineSeries := make([]Series, 0)
	barSeries := make([]Series, 0)
	isPieChart := false
	for index, item := range opt.SeriesList {
		item.index = index
		switch item.Type {
		case ChartTypePie:
			isPieChart = true
		case ChartTypeBar:
			barSeries = append(barSeries, item)
		default:
			lineSeries = append(lineSeries, item)
		}
	}
	// 如果指定了pie，则以pie的形式处理，pie不支持多类型图表
	// pie不需要axis
	if isPieChart {
		opt.XAxis.Hidden = true
		opt.YAxis.Hidden = true
	}
	result, err := chartBasicRender(&opt)
	if err != nil {
		return nil, err
	}
	fns := []func() error{
		// pie render
		func() error {
			if !isPieChart {
				return nil
			}
			_, err := pieChartRender(opt, result)
			return err
		},
		// bar render
		func() error {
			// 如果是pie或者无bar类型的series
			if isPieChart || len(barSeries) == 0 {
				return nil
			}
			o := opt
			o.SeriesList = barSeries
			_, err := barChartRender(o, result)
			return err
		},
		// line render
		func() error {
			// 如果是pie或者无line类型的series
			if isPieChart || len(lineSeries) == 0 {
				return nil
			}
			o := opt
			o.SeriesList = lineSeries
			_, err := lineChartRender(o, result)
			return err
		},
		// legend需要在顶层，因此最后render
		func() error {
			_, err := NewLegend(result.d, opt.Legend).Render()
			return err
		},
	}

	for _, fn := range fns {
		err = fn()
		if err != nil {
			return nil, err
		}
	}
	for _, child := range opt.Children {
		child.Parent = result.d
		if len(child.Theme) == 0 {
			child.Theme = opt.Theme
		}
		_, err = Render(child)
		if err != nil {
			return nil, err
		}
	}
	return result.d, nil
}

func chartBasicRender(opt *ChartOption) (*basicRenderResult, error) {
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

	opt.FillDefault(opt.Theme)
	if opt.Parent == nil {
		d.setBackground(opt.getWidth(), opt.getHeight(), opt.BackgroundColor)
	}

	// 标题
	titleBox, err := drawTitle(d, &opt.Title)
	if err != nil {
		return nil, err
	}

	xAxisHeight := 0
	var xRange *Range

	if !opt.XAxis.Hidden {
		// xAxis
		xAxisHeight, xRange, err = drawXAxis(d, &opt.XAxis)
		if err != nil {
			return nil, err
		}
	}

	// 暂时仅支持单一yaxis
	var yRange *Range
	if !opt.YAxis.Hidden {
		yRange, err = drawYAxis(d, opt, xAxisHeight, chart.Box{
			Top: titleBox.Height(),
		})
		if err != nil {
			return nil, err
		}
	}
	return &basicRenderResult{
		xRange:   xRange,
		yRange:   yRange,
		d:        d,
		titleBox: titleBox,
	}, nil
}
