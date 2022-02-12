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

const labelFontSize = 10

type ChartOption struct {
	Type            string
	Font            *truetype.Font
	Theme           string
	Title           TitleOption
	Legend          LegendOption
	XAxis           XAxisOption
	YAxisList       []YAxisOption
	Width           int
	Height          int
	Parent          *Draw
	Padding         chart.Box
	Box             chart.Box
	SeriesList      []Series
	BackgroundColor drawing.Color
	Children        []ChartOption
}

func (o *ChartOption) FillDefault(theme string) {
	t := NewTheme(theme)
	// 如果为空，初始化
	yAxisCount := 1
	for _, series := range o.SeriesList {
		if series.YAxisIndex >= yAxisCount {
			yAxisCount++
		}
	}
	yAxisList := make([]YAxisOption, yAxisCount)
	copy(yAxisList, o.YAxisList)
	o.YAxisList = yAxisList

	if o.Font == nil {
		o.Font, _ = chart.GetDefaultFont()
	}
	if o.BackgroundColor.IsZero() {
		o.BackgroundColor = t.GetBackgroundColor()
	}
	if o.Padding.IsZero() {
		o.Padding = chart.Box{
			Top:    20,
			Right:  10,
			Bottom: 10,
			Left:   10,
		}
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

	o.Legend.theme = theme
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
	if o.Width != 0 {
		return o.Width
	}
	if o.Parent != nil {
		return o.Parent.Box.Width()
	}
	return 600
}

func (o *ChartOption) getHeight() int {

	if o.Height != 0 {
		return o.Height
	}
	if o.Parent != nil {
		return o.Parent.Box.Height()
	}
	return 400
}

func (o *ChartOption) newYRange(axisIndex int) Range {
	min := math.MaxFloat64
	max := -math.MaxFloat64
	if axisIndex >= len(o.YAxisList) {
		axisIndex = 0
	}
	yAxis := o.YAxisList[axisIndex]

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
	if yAxis.Min != nil {
		min = *yAxis.Min
	}
	if yAxis.Max != nil {
		max = *yAxis.Max
	}
	divideCount := 6
	// y轴分设置默认划分为6块
	r := NewRange(min, max, divideCount)

	// 由于NewRange会重新计算min max
	if yAxis.Min != nil {
		r.Min = min
	}
	if yAxis.Max != nil {
		r.Max = max
	}

	return r
}

type basicRenderResult struct {
	xRange     *Range
	yRangeList []*Range
	d          *Draw
	titleBox   chart.Box
}

func (r *basicRenderResult) getYRange(index int) *Range {
	if index >= len(r.yRangeList) {
		index = 0
	}
	return r.yRangeList[index]
}

func Render(opt ChartOption) (*Draw, error) {
	if len(opt.SeriesList) == 0 {
		return nil, errors.New("series can not be nil")
	}
	opt.FillDefault(opt.Theme)

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
		for index := range opt.YAxisList {
			opt.YAxisList[index].Hidden = true
		}
	}
	result, err := chartBasicRender(&opt)
	if err != nil {
		return nil, err
	}
	markPointRenderOptions := make([]*markPointRenderOption, 0)
	fns := []func() error{
		// pie render
		func() error {
			if !isPieChart {
				return nil
			}
			err := pieChartRender(opt, result)
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
			options, err := barChartRender(o, result)
			if err != nil {
				return err
			}
			markPointRenderOptions = append(markPointRenderOptions, options...)
			return nil
		},
		// line render
		func() error {
			// 如果是pie或者无line类型的series
			if isPieChart || len(lineSeries) == 0 {
				return nil
			}
			o := opt
			o.SeriesList = lineSeries
			options, err := lineChartRender(o, result)
			if err != nil {
				return err
			}
			markPointRenderOptions = append(markPointRenderOptions, options...)
			return nil
		},
		// legend需要在顶层，因此此处render
		func() error {
			_, err := NewLegend(result.d, opt.Legend).Render()
			return err
		},
		// mark point最后render
		func() error {
			// mark point render不会出错
			for _, opt := range markPointRenderOptions {
				markPointRender(opt)
			}
			return nil
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
			Type:   opt.Type,
			Parent: opt.Parent,
			Width:  opt.getWidth(),
			Height: opt.getHeight(),
		},
		PaddingOption(opt.Padding),
		BoxOption(opt.Box),
	)
	if err != nil {
		return nil, err
	}

	if len(opt.YAxisList) > 2 {
		return nil, errors.New("y axis should not be gt 2")
	}
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
		xAxisHeight, xRange, err = drawXAxis(d, &opt.XAxis, len(opt.YAxisList))
		if err != nil {
			return nil, err
		}
	}

	yRangeList := make([]*Range, len(opt.YAxisList))

	for index, yAxis := range opt.YAxisList {
		var yRange *Range
		if !yAxis.Hidden {
			yRange, err = drawYAxis(d, opt, index, xAxisHeight, chart.Box{
				Top: titleBox.Height(),
			})
			if err != nil {
				return nil, err
			}
			yRangeList[index] = yRange
		}
	}
	return &basicRenderResult{
		xRange:     xRange,
		yRangeList: yRangeList,
		d:          d,
		titleBox:   titleBox,
	}, nil
}
