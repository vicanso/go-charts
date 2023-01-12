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
)

type funnelChart struct {
	p   *Painter
	opt *FunnelChartOption
}

// NewFunnelSeriesList returns a series list for funnel
func NewFunnelSeriesList(values []float64) SeriesList {
	seriesList := make(SeriesList, len(values))
	for index, value := range values {
		seriesList[index] = NewSeriesFromValues([]float64{
			value,
		}, ChartTypeFunnel)
	}
	return seriesList
}

// NewFunnelChart returns a funnel chart renderer
func NewFunnelChart(p *Painter, opt FunnelChartOption) *funnelChart {
	if opt.Theme == nil {
		opt.Theme = defaultTheme
	}
	return &funnelChart{
		p:   p,
		opt: &opt,
	}
}

type FunnelChartOption struct {
	// The theme
	Theme ColorPalette
	// The font size
	Font *truetype.Font
	// The data series list
	SeriesList SeriesList
	// The padding of line chart
	Padding Box
	// The option of title
	Title TitleOption
	// The legend option
	Legend LegendOption
}

func (f *funnelChart) render(result *defaultRenderResult, seriesList SeriesList) (Box, error) {
	opt := f.opt
	seriesPainter := result.seriesPainter
	max := seriesList[0].Data[0].Value
	min := float64(0)
	for _, item := range seriesList {
		if item.Max != nil {
			max = *item.Max
		}
		if item.Min != nil {
			min = *item.Min
		}
	}
	theme := opt.Theme
	gap := 2
	height := seriesPainter.Height()
	width := seriesPainter.Width()
	count := len(seriesList)

	h := (height - gap*(count-1)) / count

	y := 0
	widthList := make([]int, len(seriesList))
	textList := make([]string, len(seriesList))
	seriesNames := seriesList.Names()
	offset := max - min
	for index, item := range seriesList {
		value := item.Data[0].Value
		// 最大最小值一致则为100%
		widthPercent := 100.0
		if offset != 0 {
			widthPercent = (value - min) / offset
		}
		w := int(widthPercent * float64(width))
		widthList[index] = w
		// 如果最大值为0，则占比100%
		percent := 1.0
		if max != 0 {
			percent = value / max
		}
		textList[index] = NewFunnelLabelFormatter(seriesNames, item.Label.Formatter)(index, value, percent)
	}

	for index, w := range widthList {
		series := seriesList[index]
		nextWidth := 0
		if index+1 < len(widthList) {
			nextWidth = widthList[index+1]
		}
		topStartX := (width - w) >> 1
		topEndX := topStartX + w
		bottomStartX := (width - nextWidth) >> 1
		bottomEndX := bottomStartX + nextWidth
		points := []Point{
			{
				X: topStartX,
				Y: y,
			},
			{
				X: topEndX,
				Y: y,
			},
			{
				X: bottomEndX,
				Y: y + h,
			},
			{
				X: bottomStartX,
				Y: y + h,
			},
			{
				X: topStartX,
				Y: y,
			},
		}
		color := theme.GetSeriesColor(series.index)

		seriesPainter.OverrideDrawingStyle(Style{
			FillColor: color,
		}).FillArea(points)

		// 文本
		text := textList[index]
		seriesPainter.OverrideTextStyle(Style{
			FontColor: theme.GetTextColor(),
			FontSize:  labelFontSize,
			Font:      opt.Font,
		})
		textBox := seriesPainter.MeasureText(text)
		textX := width>>1 - textBox.Width()>>1
		textY := y + h>>1
		seriesPainter.Text(text, textX, textY)
		y += (h + gap)
	}

	return f.p.box, nil
}

func (f *funnelChart) Render() (Box, error) {
	p := f.p
	opt := f.opt
	renderResult, err := defaultRender(p, defaultRenderOption{
		Theme:      opt.Theme,
		Padding:    opt.Padding,
		SeriesList: opt.SeriesList,
		XAxis: XAxisOption{
			Show: FalseFlag(),
		},
		YAxisOptions: []YAxisOption{
			{
				Show: FalseFlag(),
			},
		},
		TitleOption:  opt.Title,
		LegendOption: opt.Legend,
	})
	if err != nil {
		return BoxZero, err
	}
	seriesList := opt.SeriesList.Filter(ChartTypeFunnel)
	return f.render(renderResult, seriesList)
}
