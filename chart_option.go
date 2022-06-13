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
	"github.com/wcharczuk/go-chart/v2"
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
	// RadarIndicators []RadarIndicator
	// The background color of chart
	BackgroundColor Color
	// The child charts
	Children []ChartOption
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
		o.font, _ = chart.GetDefaultFont()
	}
	if o.BackgroundColor.IsZero() {
		o.BackgroundColor = t.GetBackgroundColor()
	}
	if o.Padding.IsZero() {
		o.Padding = chart.Box{
			Top:    10,
			Right:  10,
			Bottom: 10,
			Left:   10,
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
