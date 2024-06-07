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
	"sort"

	"github.com/wcharczuk/go-chart/v2"
)

const labelFontSize = 10
const smallLabelFontSize = 8
const defaultDotWidth = 2.0
const defaultStrokeWidth = 2.0

var defaultChartWidth = 600
var defaultChartHeight = 400

// SetDefaultWidth sets default width of chart
func SetDefaultWidth(width int) {
	if width > 0 {
		defaultChartWidth = width
	}
}

// SetDefaultHeight sets default height of chart
func SetDefaultHeight(height int) {
	if height > 0 {
		defaultChartHeight = height
	}
}

var nullValue = math.MaxFloat64

// SetNullValue sets the null value, default is MaxFloat64
func SetNullValue(v float64) {
	nullValue = v
}

// GetNullValue gets the null value
func GetNullValue() float64 {
	return nullValue
}

type Renderer interface {
	Render() (Box, error)
}

type renderHandler struct {
	list []func() error
}

func (rh *renderHandler) Add(fn func() error) {
	list := rh.list
	if len(list) == 0 {
		list = make([]func() error, 0)
	}
	rh.list = append(list, fn)
}

func (rh *renderHandler) Do() error {
	for _, fn := range rh.list {
		err := fn()
		if err != nil {
			return err
		}
	}
	return nil
}

type defaultRenderOption struct {
	Theme      ColorPalette
	Padding    Box
	SeriesList SeriesList
	// The y axis option
	YAxisOptions []YAxisOption
	// The x axis option
	XAxis XAxisOption
	// The title option
	TitleOption TitleOption
	// The legend option
	LegendOption LegendOption
	// background is filled
	backgroundIsFilled bool
	// x y axis is reversed
	axisReversed bool
}

type defaultRenderResult struct {
	axisRanges map[int]axisRange
	// 图例区域
	seriesPainter *Painter
}

func defaultRender(p *Painter, opt defaultRenderOption) (*defaultRenderResult, error) {
	seriesList := opt.SeriesList
	seriesList.init()
	if !opt.backgroundIsFilled {
		p.SetBackground(p.Width(), p.Height(), opt.Theme.GetBackgroundColor())
	}

	if !opt.Padding.IsZero() {
		p = p.Child(PainterPaddingOption(opt.Padding))
	}

	legendHeight := 0
	if len(opt.LegendOption.Data) != 0 {
		if opt.LegendOption.Theme == nil {
			opt.LegendOption.Theme = opt.Theme
		}
		legendResult, err := NewLegendPainter(p, opt.LegendOption).Render()
		if err != nil {
			return nil, err
		}
		legendHeight = legendResult.Height()
	}

	// 如果有标题
	if opt.TitleOption.Text != "" {
		if opt.TitleOption.Theme == nil {
			opt.TitleOption.Theme = opt.Theme
		}
		titlePainter := NewTitlePainter(p, opt.TitleOption)

		titleBox, err := titlePainter.Render()
		if err != nil {
			return nil, err
		}

		top := chart.MaxInt(legendHeight, titleBox.Height())
		// 如果是垂直方式，则不计算legend高度
		if opt.LegendOption.Orient == OrientVertical {
			top = titleBox.Height()
		}
		p = p.Child(PainterPaddingOption(Box{
			// 标题下留白
			Top: top + 20,
		}))
	}

	result := defaultRenderResult{
		axisRanges: make(map[int]axisRange),
	}

	// 计算图表对应的轴有哪些
	axisIndexList := make([]int, 0)
	for _, series := range opt.SeriesList {
		if containsInt(axisIndexList, series.AxisIndex) {
			continue
		}
		axisIndexList = append(axisIndexList, series.AxisIndex)
	}
	// 高度需要减去x轴的高度
	rangeHeight := p.Height() - defaultXAxisHeight
	rangeWidthLeft := 0
	rangeWidthRight := 0

	// 倒序
	sort.Sort(sort.Reverse(sort.IntSlice(axisIndexList)))

	// 计算对应的axis range
	for _, index := range axisIndexList {
		yAxisOption := YAxisOption{}
		if len(opt.YAxisOptions) > index {
			yAxisOption = opt.YAxisOptions[index]
		}
		divideCount := yAxisOption.DivideCount
		if divideCount <= 0 {
			divideCount = defaultAxisDivideCount
		}
		max, min := opt.SeriesList.GetMaxMin(index)
		r := NewRange(AxisRangeOption{
			Painter: p,
			Min:     min,
			Max:     max,
			// 高度需要减去x轴的高度
			Size: rangeHeight,
			// 分隔数量
			DivideCount: divideCount,
		})
		if yAxisOption.Min != nil && *yAxisOption.Min <= min {
			r.min = *yAxisOption.Min
		}
		if yAxisOption.Max != nil && *yAxisOption.Max >= max {
			r.max = *yAxisOption.Max
		}
		result.axisRanges[index] = r

		if yAxisOption.Theme == nil {
			yAxisOption.Theme = opt.Theme
		}
		if !opt.axisReversed {
			yAxisOption.Data = r.Values()
		} else {
			yAxisOption.isCategoryAxis = true
			// 由于x轴为value部分，因此计算其label单独处理
			opt.XAxis.Data = NewRange(AxisRangeOption{
				Painter: p,
				Min:     min,
				Max:     max,
				// 高度需要减去x轴的高度
				Size: rangeHeight,
				// 分隔数量
				DivideCount: defaultAxisDivideCount,
			}).Values()
			opt.XAxis.isValueAxis = true
		}
		reverseStringSlice(yAxisOption.Data)
		// TODO生成其它位置既yAxis
		var yAxis *axisPainter
		child := p.Child(PainterPaddingOption(Box{
			Left:  rangeWidthLeft,
			Right: rangeWidthRight,
		}))
		if index == 0 {
			yAxis = NewLeftYAxis(child, yAxisOption)
		} else {
			yAxis = NewRightYAxis(child, yAxisOption)
		}
		yAxisBox, err := yAxis.Render()
		if err != nil {
			return nil, err
		}
		if index == 0 {
			rangeWidthLeft += yAxisBox.Width()
		} else {
			rangeWidthRight += yAxisBox.Width()
		}
	}

	if opt.XAxis.Theme == nil {
		opt.XAxis.Theme = opt.Theme
	}
	xAxis := NewBottomXAxis(p.Child(PainterPaddingOption(Box{
		Left:  rangeWidthLeft,
		Right: rangeWidthRight,
	})), opt.XAxis)
	_, err := xAxis.Render()
	if err != nil {
		return nil, err
	}

	result.seriesPainter = p.Child(PainterPaddingOption(Box{
		Bottom: defaultXAxisHeight,
		Left:   rangeWidthLeft,
		Right:  rangeWidthRight,
	}))
	return &result, nil
}

func doRender(renderers ...Renderer) error {
	for _, r := range renderers {
		_, err := r.Render()
		if err != nil {
			return err
		}
	}
	return nil
}

func Render(opt ChartOption, opts ...OptionFunc) (*Painter, error) {
	for _, fn := range opts {
		fn(&opt)
	}
	opt.fillDefault()

	isChild := true
	if opt.Parent == nil {
		isChild = false
		p, err := NewPainter(PainterOptions{
			Type:   opt.Type,
			Width:  opt.Width,
			Height: opt.Height,
			Font:   opt.font,
		})
		if err != nil {
			return nil, err
		}
		opt.Parent = p
	}
	p := opt.Parent
	if opt.ValueFormatter != nil {
		p.valueFormatter = opt.ValueFormatter
	}
	if !opt.Box.IsZero() {
		p = p.Child(PainterBoxOption(opt.Box))
	}
	if !isChild {
		p.SetBackground(p.Width(), p.Height(), opt.BackgroundColor)
	}
	seriesList := opt.SeriesList
	seriesList.init()

	seriesCount := len(seriesList)

	// line chart
	lineSeriesList := seriesList.Filter(ChartTypeLine)
	barSeriesList := seriesList.Filter(ChartTypeBar)
	horizontalBarSeriesList := seriesList.Filter(ChartTypeHorizontalBar)
	pieSeriesList := seriesList.Filter(ChartTypePie)
	radarSeriesList := seriesList.Filter(ChartTypeRadar)
	funnelSeriesList := seriesList.Filter(ChartTypeFunnel)

	if len(horizontalBarSeriesList) != 0 && len(horizontalBarSeriesList) != seriesCount {
		return nil, errors.New("Horizontal bar can not mix other charts")
	}
	if len(pieSeriesList) != 0 && len(pieSeriesList) != seriesCount {
		return nil, errors.New("Pie can not mix other charts")
	}
	if len(radarSeriesList) != 0 && len(radarSeriesList) != seriesCount {
		return nil, errors.New("Radar can not mix other charts")
	}
	if len(funnelSeriesList) != 0 && len(funnelSeriesList) != seriesCount {
		return nil, errors.New("Funnel can not mix other charts")
	}

	axisReversed := len(horizontalBarSeriesList) != 0
	renderOpt := defaultRenderOption{
		Theme:        opt.theme,
		Padding:      opt.Padding,
		SeriesList:   opt.SeriesList,
		XAxis:        opt.XAxis,
		YAxisOptions: opt.YAxisOptions,
		TitleOption:  opt.Title,
		LegendOption: opt.Legend,
		axisReversed: axisReversed,
		// 前置已设置背景色
		backgroundIsFilled: true,
	}
	if len(pieSeriesList) != 0 ||
		len(radarSeriesList) != 0 ||
		len(funnelSeriesList) != 0 {
		renderOpt.XAxis.Show = FalseFlag()
		renderOpt.YAxisOptions = []YAxisOption{
			{
				Show: FalseFlag(),
			},
		}
	}
	if len(horizontalBarSeriesList) != 0 {
		renderOpt.YAxisOptions[0].DivideCount = len(renderOpt.YAxisOptions[0].Data)
		renderOpt.YAxisOptions[0].Unit = 1
	}

	renderResult, err := defaultRender(p, renderOpt)
	if err != nil {
		return nil, err
	}

	handler := renderHandler{}

	// bar chart
	if len(barSeriesList) != 0 {
		handler.Add(func() error {
			_, err := NewBarChart(p, BarChartOption{
				Theme:     opt.theme,
				Font:      opt.font,
				XAxis:     opt.XAxis,
				BarWidth:  opt.BarWidth,
				BarMargin: opt.BarMargin,
			}).render(renderResult, barSeriesList)
			return err
		})
	}

	// horizontal bar chart
	if len(horizontalBarSeriesList) != 0 {
		handler.Add(func() error {
			_, err := NewHorizontalBarChart(p, HorizontalBarChartOption{
				Theme:        opt.theme,
				Font:         opt.font,
				BarHeight:    opt.BarHeight,
				BarMargin:    opt.BarMargin,
				YAxisOptions: opt.YAxisOptions,
			}).render(renderResult, horizontalBarSeriesList)
			return err
		})
	}

	// pie chart
	if len(pieSeriesList) != 0 {
		handler.Add(func() error {
			_, err := NewPieChart(p, PieChartOption{
				Theme: opt.theme,
				Font:  opt.font,
			}).render(renderResult, pieSeriesList)
			return err
		})
	}

	// line chart
	if len(lineSeriesList) != 0 {
		handler.Add(func() error {
			_, err := NewLineChart(p, LineChartOption{
				Theme:       opt.theme,
				Font:        opt.font,
				XAxis:       opt.XAxis,
				SymbolShow:  opt.SymbolShow,
				StrokeWidth: opt.LineStrokeWidth,
				FillArea:    opt.FillArea,
				Opacity:     opt.Opacity,
			}).render(renderResult, lineSeriesList)
			return err
		})
	}

	// radar chart
	if len(radarSeriesList) != 0 {
		handler.Add(func() error {
			_, err := NewRadarChart(p, RadarChartOption{
				Theme: opt.theme,
				Font:  opt.font,
				// 相应值
				RadarIndicators: opt.RadarIndicators,
			}).render(renderResult, radarSeriesList)
			return err
		})
	}

	// funnel chart
	if len(funnelSeriesList) != 0 {
		handler.Add(func() error {
			_, err := NewFunnelChart(p, FunnelChartOption{
				Theme: opt.theme,
				Font:  opt.font,
			}).render(renderResult, funnelSeriesList)
			return err
		})
	}

	err = handler.Do()

	if err != nil {
		return nil, err
	}
	for _, item := range opt.Children {
		item.Parent = p
		if item.Theme == "" {
			item.Theme = opt.Theme
		}
		if item.FontFamily == "" {
			item.FontFamily = opt.FontFamily
		}
		_, err = Render(item)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}
