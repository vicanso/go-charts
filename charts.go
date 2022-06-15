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

import "errors"

const labelFontSize = 10
const defaultDotWidth = 2.0
const defaultStrokeWidth = 2.0

var defaultChartWidth = 600
var defaultChartHeight = 400

func SetDefaultWidth(width int) {
	if width > 0 {
		defaultChartWidth = width
	}
}
func SetDefaultHeight(height int) {
	if height > 0 {
		defaultChartHeight = height
	}
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

	if len(opt.LegendOption.Data) != 0 {
		if opt.LegendOption.Theme == nil {
			opt.LegendOption.Theme = opt.Theme
		}
		_, err := NewLegendPainter(p, opt.LegendOption).Render()
		if err != nil {
			return nil, err
		}
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
		p = p.Child(PainterPaddingOption(Box{
			// 标题下留白
			Top: titleBox.Height() + 20,
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
		axisIndexList = append(axisIndexList, series.index)
	}
	// 高度需要减去x轴的高度
	rangeHeight := p.Height() - defaultXAxisHeight
	rangeWidthLeft := 0
	rangeWidthRight := 0

	// 计算对应的axis range
	for _, index := range axisIndexList {
		max, min := opt.SeriesList.GetMaxMin(index)
		r := NewRange(AxisRangeOption{
			Min: min,
			Max: max,
			// 高度需要减去x轴的高度
			Size: rangeHeight,
			// 分隔数量
			DivideCount: defaultAxisDivideCount,
		})
		result.axisRanges[index] = r
		yAxisOption := YAxisOption{}
		if len(opt.YAxisOptions) > index {
			yAxisOption = opt.YAxisOptions[index]
		}
		if yAxisOption.Theme == nil {
			yAxisOption.Theme = opt.Theme
		}
		if !opt.axisReversed {
			yAxisOption.Data = r.Values()
		} else {
			yAxisOption.isCategoryAxis = true
			opt.XAxis.Data = r.Values()
			opt.XAxis.isValueAxis = true
		}
		reverseStringSlice(yAxisOption.Data)
		// TODO生成其它位置既yAxis
		yAxis := NewLeftYAxis(p, yAxisOption)
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
		Left: rangeWidthLeft,
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

func Render(opt ChartOption) (*Painter, error) {
	opt.fillDefault()

	if opt.Parent == nil {
		p, err := NewPainter(PainterOptions{
			Type:   opt.Type,
			Width:  opt.Width,
			Height: opt.Height,
		})
		if err != nil {
			return nil, err
		}
		opt.Parent = p
	}
	p := opt.Parent
	p.SetBackground(p.Width(), p.Height(), opt.BackgroundColor)
	seriesList := opt.SeriesList
	seriesList.init()

	// line chart
	lineSeriesList := seriesList.Filter(ChartTypeLine)
	barSeriesList := seriesList.Filter(ChartTypeBar)
	horizontalBarSeriesList := seriesList.Filter(ChartTypeHorizontalBar)
	if len(horizontalBarSeriesList) != 0 && len(horizontalBarSeriesList) != len(seriesList) {
		return nil, errors.New("Horizontal bar can not mix other charts")
	}

	axisReversed := len(horizontalBarSeriesList) != 0

	renderResult, err := defaultRender(p, defaultRenderOption{
		Theme:        opt.theme,
		Padding:      opt.Padding,
		SeriesList:   opt.SeriesList,
		XAxis:        opt.XAxis,
		YAxisOptions: opt.YAxisOptions,
		TitleOption:  opt.Title,
		LegendOption: opt.Legend,
		axisReversed: axisReversed,
	})
	if err != nil {
		return nil, err
	}

	handler := renderHandler{}

	if len(lineSeriesList) != 0 {
		handler.Add(func() error {
			_, err := NewLineChart(p, LineChartOption{
				Theme: opt.theme,
				Font:  opt.font,
				XAxis: opt.XAxis,
			}).render(renderResult, lineSeriesList)
			return err
		})
	}

	// bar chart
	if len(barSeriesList) != 0 {
		handler.Add(func() error {
			_, err := NewBarChart(p, BarChartOption{
				Theme: opt.theme,
				Font:  opt.font,
				XAxis: opt.XAxis,
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
				YAxisOptions: opt.YAxisOptions,
			}).render(renderResult, horizontalBarSeriesList)
			return err
		})
	}

	err = handler.Do()

	if err != nil {
		return nil, err
	}

	return p, nil
}
