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
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
)

type horizontalBarChart struct {
	p   *Painter
	opt *HorizontalBarChartOption
}

type HorizontalBarChartOption struct {
	// The theme
	Theme ColorPalette
	// The font size
	Font *truetype.Font
	// The data series list
	SeriesList SeriesList
	// The x axis option
	XAxis XAxisOption
	// The padding of line chart
	Padding Box
	// The y axis option
	YAxisOptions []YAxisOption
	// The option of title
	Title TitleOption
	// The legend option
	Legend    LegendOption
	BarHeight int
}

// NewHorizontalBarChart returns a horizontal bar chart renderer
func NewHorizontalBarChart(p *Painter, opt HorizontalBarChartOption) *horizontalBarChart {
	if opt.Theme == nil {
		opt.Theme = defaultTheme
	}
	return &horizontalBarChart{
		p:   p,
		opt: &opt,
	}
}

func (h *horizontalBarChart) render(result *defaultRenderResult, seriesList SeriesList) (Box, error) {
	p := h.p
	opt := h.opt
	seriesPainter := result.seriesPainter
	yRange := result.axisRanges[0]
	y0, y1 := yRange.GetRange(0)
	height := int(y1 - y0)
	// 每一块之间的margin
	margin := 10
	// 每一个bar之间的margin
	barMargin := 5
	if height < 20 {
		margin = 2
		barMargin = 2
	} else if height < 50 {
		margin = 5
		barMargin = 3
	}
	seriesCount := len(seriesList)
	// 总的高度-两个margin-(总数-1)的barMargin
	barHeight := (height - 2*margin - barMargin*(seriesCount-1)) / seriesCount
	if opt.BarHeight > 0 && opt.BarHeight < barHeight {
		barHeight = opt.BarHeight
		margin = (height - seriesCount*barHeight - barMargin*(seriesCount-1)) / 2
	}

	theme := opt.Theme

	max, min := seriesList.GetMaxMin(0)
	xRange := NewRange(AxisRangeOption{
		Min:         min,
		Max:         max,
		DivideCount: defaultAxisDivideCount,
		Size:        seriesPainter.Width(),
	})

	for index := range seriesList {
		series := seriesList[index]
		seriesColor := theme.GetSeriesColor(series.index)
		divideValues := yRange.AutoDivide()
		for j, item := range series.Data {
			if j >= yRange.divideCount {
				continue
			}
			// 显示位置切换
			j = yRange.divideCount - j - 1
			y := divideValues[j]
			y += margin
			if index != 0 {
				y += index * (barHeight + barMargin)
			}

			w := int(xRange.getHeight(item.Value))
			fillColor := seriesColor
			if !item.Style.FillColor.IsZero() {
				fillColor = item.Style.FillColor
			}
			right := w
			seriesPainter.OverrideDrawingStyle(Style{
				FillColor: fillColor,
			}).Rect(chart.Box{
				Top:    y,
				Left:   0,
				Right:  right,
				Bottom: y + barHeight,
			})
		}
	}
	return p.box, nil
}

func (h *horizontalBarChart) Render() (Box, error) {
	p := h.p
	opt := h.opt
	renderResult, err := defaultRender(p, defaultRenderOption{
		Theme:        opt.Theme,
		Padding:      opt.Padding,
		SeriesList:   opt.SeriesList,
		XAxis:        opt.XAxis,
		YAxisOptions: opt.YAxisOptions,
		TitleOption:  opt.Title,
		LegendOption: opt.Legend,
		axisReversed: true,
	})
	if err != nil {
		return BoxZero, err
	}
	seriesList := opt.SeriesList.Filter(ChartTypeHorizontalBar)
	return h.render(renderResult, seriesList)
}
