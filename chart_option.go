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

// OptionFunc option function
type OptionFunc func(opt *ChartOption)

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

// LegendOptionFunc set legend of chart
func LegendOptionFunc(legend LegendOption) OptionFunc {
	return func(opt *ChartOption) {
		opt.Legend = legend
	}
}

// XAxisOptionFunc set x axis of chart
func XAxisOptionFunc(xAxisOption XAxisOption) OptionFunc {
	return func(opt *ChartOption) {
		opt.XAxis = xAxisOption
	}
}

// YAxisOptionFunc set y axis of chart, support two y axis
func YAxisOptionFunc(yAxisOption ...YAxisOption) OptionFunc {
	return func(opt *ChartOption) {
		opt.YAxisList = yAxisOption
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
func PaddingOptionFunc(padding chart.Box) OptionFunc {
	return func(opt *ChartOption) {
		opt.Padding = padding
	}
}

// BoxOptionFunc set box of chart
func BoxOptionFunc(box chart.Box) OptionFunc {
	return func(opt *ChartOption) {
		opt.Box = box
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
func RadarIndicatorOptionFunc(radarIndicator ...RadarIndicator) OptionFunc {
	return func(opt *ChartOption) {
		opt.RadarIndicators = radarIndicator
	}
}

// BackgroundColorOptionFunc set background color of chart
func BackgroundColorOptionFunc(color drawing.Color) OptionFunc {
	return func(opt *ChartOption) {
		opt.BackgroundColor = color
	}
}

// LineRender line chart render
func LineRender(values [][]float64, opts ...OptionFunc) (*Draw, error) {
	seriesList := make(SeriesList, len(values))
	for index, value := range values {
		seriesList[index] = NewSeriesFromValues(value, ChartTypeLine)
	}
	return Render(ChartOption{
		SeriesList: seriesList,
	}, opts...)
}

// BarRender bar chart render
func BarRender(values [][]float64, opts ...OptionFunc) (*Draw, error) {
	seriesList := make(SeriesList, len(values))
	for index, value := range values {
		seriesList[index] = NewSeriesFromValues(value, ChartTypeBar)
	}
	return Render(ChartOption{
		SeriesList: seriesList,
	}, opts...)
}

// PieRender pie chart render
func PieRender(values []float64, opts ...OptionFunc) (*Draw, error) {
	return Render(ChartOption{
		SeriesList: NewPieSeriesList(values),
	}, opts...)
}

// RadarRender radar chart render
func RadarRender(values [][]float64, opts ...OptionFunc) (*Draw, error) {
	seriesList := make(SeriesList, len(values))
	for index, value := range values {
		seriesList[index] = NewSeriesFromValues(value, ChartTypeRadar)
	}
	return Render(ChartOption{
		SeriesList: seriesList,
	}, opts...)
}

// FunnelRender funnel chart render
func FunnelRender(values []float64, opts ...OptionFunc) (*Draw, error) {
	seriesList := make(SeriesList, len(values))
	for index, value := range values {
		seriesList[index] = NewSeriesFromValues([]float64{
			value,
		}, ChartTypeFunnel)
	}
	return Render(ChartOption{
		SeriesList: seriesList,
	}, opts...)
}
