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
	"sort"

	"github.com/golang/freetype/truetype"
)

type ChartOption struct {
	theme ColorPalette
	font  *truetype.Font
	// The output type of chart, "svg" or "png", default value is "svg"
	Type string
	// The font family, which should be installed first
	FontFamily string
	// The theme of chart, "light" and "dark".
	// The default theme is "light"
	Theme string
	// The title option
	Title TitleOption
	// The legend option
	Legend LegendOption
	// The x axis option
	XAxis XAxisOption
	// The y axis option list
	YAxisOptions []YAxisOption
	// The width of chart, default width is 600
	Width int
	// The height of chart, default height is 400
	Height int
	Parent *Painter
	// The padding for chart, default padding is [20, 10, 10, 10]
	Padding Box
	// The canvas box for chart
	Box Box
	// The series list
	SeriesList SeriesList
	// The radar indicator list
	RadarIndicators []RadarIndicator
	// The background color of chart
	BackgroundColor Color
	// The flag for show symbol of line, set this to *false will hide symbol
	SymbolShow *bool
	// The stroke width of line chart
	LineStrokeWidth float64
	// The bar with of bar chart
	BarWidth int
	// The margin of each bar
	BarMargin int
	// The bar height of horizontal bar chart
	BarHeight int
	// Fill the area of line chart
	FillArea bool
	// background fill (alpha) opacity
	Opacity uint8
	// The child charts
	Children []ChartOption
	// The value formatter
	ValueFormatter ValueFormatter
}

// OptionFunc option function
type OptionFunc func(opt *ChartOption)

// SVGTypeOption set svg type of chart's output
func SVGTypeOption() OptionFunc {
	return TypeOptionFunc(ChartOutputSVG)
}

// PNGTypeOption set png type of chart's output
func PNGTypeOption() OptionFunc {
	return TypeOptionFunc(ChartOutputPNG)
}

// TypeOptionFunc set type of chart's output
func TypeOptionFunc(t string) OptionFunc {
	return func(opt *ChartOption) {
		opt.Type = t
	}
}

// FontFamilyOptionFunc set font family of chart
func FontFamilyOptionFunc(fontFamily string) OptionFunc {
	return func(opt *ChartOption) {
		opt.FontFamily = fontFamily
	}
}

// ThemeOptionFunc set them of chart
func ThemeOptionFunc(theme string) OptionFunc {
	return func(opt *ChartOption) {
		opt.Theme = theme
	}
}

// TitleOptionFunc set title of chart
func TitleOptionFunc(title TitleOption) OptionFunc {
	return func(opt *ChartOption) {
		opt.Title = title
	}
}

// TitleTextOptionFunc set title text of chart
func TitleTextOptionFunc(text string, subtext ...string) OptionFunc {
	return func(opt *ChartOption) {
		opt.Title.Text = text
		if len(subtext) != 0 {
			opt.Title.Subtext = subtext[0]
		}
	}
}

// LegendOptionFunc set legend of chart
func LegendOptionFunc(legend LegendOption) OptionFunc {
	return func(opt *ChartOption) {
		opt.Legend = legend
	}
}

// LegendLabelsOptionFunc set legend labels of chart
func LegendLabelsOptionFunc(labels []string, left ...string) OptionFunc {
	return func(opt *ChartOption) {
		opt.Legend = NewLegendOption(labels, left...)
	}
}

// XAxisOptionFunc set x axis of chart
func XAxisOptionFunc(xAxisOption XAxisOption) OptionFunc {
	return func(opt *ChartOption) {
		opt.XAxis = xAxisOption
	}
}

// XAxisDataOptionFunc set x axis data of chart
func XAxisDataOptionFunc(data []string, boundaryGap ...*bool) OptionFunc {
	return func(opt *ChartOption) {
		opt.XAxis = NewXAxisOption(data, boundaryGap...)
	}
}

// YAxisOptionFunc set y axis of chart, support two y axis
func YAxisOptionFunc(yAxisOption ...YAxisOption) OptionFunc {
	return func(opt *ChartOption) {
		opt.YAxisOptions = yAxisOption
	}
}

// YAxisDataOptionFunc set y axis data of chart
func YAxisDataOptionFunc(data []string) OptionFunc {
	return func(opt *ChartOption) {
		opt.YAxisOptions = NewYAxisOptions(data)
	}
}

// WidthOptionFunc set width of chart
func WidthOptionFunc(width int) OptionFunc {
	return func(opt *ChartOption) {
		opt.Width = width
	}
}

// HeightOptionFunc set height of chart
func HeightOptionFunc(height int) OptionFunc {
	return func(opt *ChartOption) {
		opt.Height = height
	}
}

// PaddingOptionFunc set padding of chart
func PaddingOptionFunc(padding Box) OptionFunc {
	return func(opt *ChartOption) {
		opt.Padding = padding
	}
}

// BoxOptionFunc set box of chart
func BoxOptionFunc(box Box) OptionFunc {
	return func(opt *ChartOption) {
		opt.Box = box
	}
}

// PieSeriesShowLabel set pie series show label
func PieSeriesShowLabel() OptionFunc {
	return func(opt *ChartOption) {
		for index := range opt.SeriesList {
			opt.SeriesList[index].Label.Show = true
		}
	}
}

// ChildOptionFunc add child chart
func ChildOptionFunc(child ...ChartOption) OptionFunc {
	return func(opt *ChartOption) {
		if opt.Children == nil {
			opt.Children = make([]ChartOption, 0)
		}
		opt.Children = append(opt.Children, child...)
	}
}

// RadarIndicatorOptionFunc set radar indicator of chart
func RadarIndicatorOptionFunc(names []string, values []float64) OptionFunc {
	return func(opt *ChartOption) {
		opt.RadarIndicators = NewRadarIndicators(names, values)
	}
}

// BackgroundColorOptionFunc set background color of chart
func BackgroundColorOptionFunc(color Color) OptionFunc {
	return func(opt *ChartOption) {
		opt.BackgroundColor = color
	}
}

// MarkLineOptionFunc set mark line for series of chart
func MarkLineOptionFunc(seriesIndex int, markLineTypes ...string) OptionFunc {
	return func(opt *ChartOption) {
		if len(opt.SeriesList) <= seriesIndex {
			return
		}
		opt.SeriesList[seriesIndex].MarkLine = NewMarkLine(markLineTypes...)
	}
}

// MarkPointOptionFunc set mark point for series of chart
func MarkPointOptionFunc(seriesIndex int, markPointTypes ...string) OptionFunc {
	return func(opt *ChartOption) {
		if len(opt.SeriesList) <= seriesIndex {
			return
		}
		opt.SeriesList[seriesIndex].MarkPoint = NewMarkPoint(markPointTypes...)
	}
}

func (o *ChartOption) fillDefault() {
	t := NewTheme(o.Theme)
	o.theme = t
	// 如果为空，初始化
	axisCount := 1
	for _, series := range o.SeriesList {
		if series.AxisIndex >= axisCount {
			axisCount++
		}
	}
	o.Width = getDefaultInt(o.Width, defaultChartWidth)
	o.Height = getDefaultInt(o.Height, defaultChartHeight)
	yAxisOptions := make([]YAxisOption, axisCount)
	copy(yAxisOptions, o.YAxisOptions)
	o.YAxisOptions = yAxisOptions
	o.font, _ = GetFont(o.FontFamily)

	if o.font == nil {
		o.font, _ = GetDefaultFont()
	} else {
		// 如果指定了字体，则设置主题的字体
		t.SetFont(o.font)
	}
	if o.BackgroundColor.IsZero() {
		o.BackgroundColor = t.GetBackgroundColor()
	}
	if o.Padding.IsZero() {
		o.Padding = Box{
			Top:    20,
			Right:  20,
			Bottom: 20,
			Left:   20,
		}
	}
	// legend与series name的关联
	if len(o.Legend.Data) == 0 {
		o.Legend.Data = o.SeriesList.Names()
	} else {
		seriesCount := len(o.SeriesList)
		for index, name := range o.Legend.Data {
			if index < seriesCount &&
				len(o.SeriesList[index].Name) == 0 {
				o.SeriesList[index].Name = name
			}
		}
		nameIndexDict := map[string]int{}
		for index, name := range o.Legend.Data {
			nameIndexDict[name] = index
		}
		// 保证series的顺序与legend一致
		sort.Slice(o.SeriesList, func(i, j int) bool {
			return nameIndexDict[o.SeriesList[i].Name] < nameIndexDict[o.SeriesList[j].Name]
		})
	}
}

// LineRender line chart render
func LineRender(values [][]float64, opts ...OptionFunc) (*Painter, error) {
	seriesList := NewSeriesListDataFromValues(values, ChartTypeLine)
	return Render(ChartOption{
		SeriesList: seriesList,
	}, opts...)
}

// BarRender bar chart render
func BarRender(values [][]float64, opts ...OptionFunc) (*Painter, error) {
	seriesList := NewSeriesListDataFromValues(values, ChartTypeBar)
	return Render(ChartOption{
		SeriesList: seriesList,
	}, opts...)
}

// HorizontalBarRender horizontal bar chart render
func HorizontalBarRender(values [][]float64, opts ...OptionFunc) (*Painter, error) {
	seriesList := NewSeriesListDataFromValues(values, ChartTypeHorizontalBar)
	return Render(ChartOption{
		SeriesList: seriesList,
	}, opts...)
}

// PieRender pie chart render
func PieRender(values []float64, opts ...OptionFunc) (*Painter, error) {
	return Render(ChartOption{
		SeriesList: NewPieSeriesList(values),
	}, opts...)
}

// RadarRender radar chart render
func RadarRender(values [][]float64, opts ...OptionFunc) (*Painter, error) {
	seriesList := NewSeriesListDataFromValues(values, ChartTypeRadar)
	return Render(ChartOption{
		SeriesList: seriesList,
	}, opts...)
}

// FunnelRender funnel chart render
func FunnelRender(values []float64, opts ...OptionFunc) (*Painter, error) {
	seriesList := NewFunnelSeriesList(values)
	return Render(ChartOption{
		SeriesList: seriesList,
	}, opts...)
}

// TableRender table chart render
func TableRender(header []string, data [][]string, spanMaps ...map[int]int) (*Painter, error) {
	opt := TableChartOption{
		Header: header,
		Data:   data,
	}
	if len(spanMaps) != 0 {
		spanMap := spanMaps[0]
		spans := make([]int, len(opt.Header))
		for index := range spans {
			v, ok := spanMap[index]
			if !ok {
				v = 1
			}
			spans[index] = v
		}
		opt.Spans = spans
	}
	return TableOptionRender(opt)
}

// TableOptionRender table render with option
func TableOptionRender(opt TableChartOption) (*Painter, error) {
	if opt.Type == "" {
		opt.Type = ChartOutputPNG
	}
	if opt.Width <= 0 {
		opt.Width = defaultChartWidth
	}
	if opt.FontFamily != "" {
		opt.Font, _ = GetFont(opt.FontFamily)
	}
	if opt.Font == nil {
		opt.Font, _ = GetDefaultFont()
	}

	p, err := NewPainter(PainterOptions{
		Type:  opt.Type,
		Width: opt.Width,
		// 仅用于计算表格高度，因此随便设置即可
		Height: 100,
		Font:   opt.Font,
	})
	if err != nil {
		return nil, err
	}
	info, err := NewTableChart(p, opt).render()
	if err != nil {
		return nil, err
	}

	p, err = NewPainter(PainterOptions{
		Type:   opt.Type,
		Width:  info.Width,
		Height: info.Height,
		Font:   opt.Font,
	})
	if err != nil {
		return nil, err
	}
	_, err = NewTableChart(p, opt).renderWithInfo(info)
	if err != nil {
		return nil, err
	}
	return p, nil
}
