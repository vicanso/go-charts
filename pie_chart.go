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
)

type pieChart struct {
	p   *Painter
	opt *PieChartOption
}

type PieChartOption struct {
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
	// background is filled
	backgroundIsFilled bool
}

// NewPieChart returns a pie chart renderer
func NewPieChart(p *Painter, opt PieChartOption) *pieChart {
	if opt.Theme == nil {
		opt.Theme = defaultTheme
	}
	return &pieChart{
		p:   p,
		opt: &opt,
	}
}

func (p *pieChart) render(result *defaultRenderResult, seriesList SeriesList) (Box, error) {
	opt := p.opt
	values := make([]float64, len(seriesList))
	total := float64(0)
	radiusValue := ""
	for index, series := range seriesList {
		if len(series.Radius) != 0 {
			radiusValue = series.Radius
		}
		value := float64(0)
		for _, item := range series.Data {
			value += item.Value
		}
		values[index] = value
		total += value
	}
	if total <= 0 {
		return BoxZero, errors.New("The sum value of pie chart should gt 0")
	}
	seriesPainter := result.seriesPainter
	cx := seriesPainter.Width() >> 1
	cy := seriesPainter.Height() >> 1

	diameter := chart.MinInt(seriesPainter.Width(), seriesPainter.Height())
	radius := getRadius(float64(diameter), radiusValue)

	labelLineWidth := 15
	if radius < 50 {
		labelLineWidth = 10
	}
	labelRadius := radius + float64(labelLineWidth)
	seriesNames := opt.Legend.Data
	if len(seriesNames) == 0 {
		seriesNames = seriesList.Names()
	}
	theme := opt.Theme

	currentValue := float64(0)
	prevPoints := make([]Point, 0)

	isOverride := func(x, y int) bool {
		for _, p := range prevPoints {
			if math.Abs(float64(p.Y-y)) > labelFontSize {
				continue
			}
			// label可能较多内容，不好计算横向占用空间
			// 因此x的位置需要中间位置两侧，否则认为override
			if (p.X <= cx && x <= cx) ||
				(p.X > cx && x > cx) {
				return true
			}
		}
		return false
	}

	for index, v := range values {
		seriesPainter.OverrideDrawingStyle(Style{
			StrokeWidth: 1,
			StrokeColor: theme.GetSeriesColor(index),
			FillColor:   theme.GetSeriesColor(index),
		})
		seriesPainter.MoveTo(cx, cy)
		start := chart.PercentToRadians(currentValue/total) - math.Pi/2
		currentValue += v
		percent := (v / total)
		delta := chart.PercentToRadians(percent)
		seriesPainter.ArcTo(cx, cy, radius, radius, start, delta).
			LineTo(cx, cy).
			Close().
			FillStroke()

		series := seriesList[index]
		// 是否显示label
		showLabel := series.Label.Show
		if !showLabel {
			continue
		}

		// label的角度为饼块中间
		angle := start + delta/2
		startx := cx + int(radius*math.Cos(angle))
		starty := cy + int(radius*math.Sin(angle))

		endx := cx + int(labelRadius*math.Cos(angle))
		endy := cy + int(labelRadius*math.Sin(angle))
		// 计算是否有重叠，如果有则调整y坐标位置
		// 最多只尝试5次
		for i := 0; i < 5; i++ {
			if !isOverride(endx, endy) {
				break
			}
			endy -= (labelFontSize << 1)
		}
		prevPoints = append(prevPoints, Point{
			X: endx,
			Y: endy,
		})

		seriesPainter.MoveTo(startx, starty)
		seriesPainter.LineTo(endx, endy)
		offset := labelLineWidth
		if endx < cx {
			offset *= -1
		}
		seriesPainter.MoveTo(endx, endy)
		endx += offset
		seriesPainter.LineTo(endx, endy)
		seriesPainter.Stroke()

		textStyle := Style{
			FontColor: theme.GetTextColor(),
			FontSize:  labelFontSize,
			Font:      opt.Font,
		}
		if !series.Label.Color.IsZero() {
			textStyle.FontColor = series.Label.Color
		}
		seriesPainter.OverrideTextStyle(textStyle)
		text := NewPieLabelFormatter(seriesNames, series.Label.Formatter)(index, v, percent)
		textBox := seriesPainter.MeasureText(text)
		textMargin := 3
		x := endx + textMargin
		y := endy + textBox.Height()>>1 - 1
		if offset < 0 {
			textWidth := textBox.Width()
			x = endx - textWidth - textMargin
		}
		seriesPainter.Text(text, x, y)
	}

	return p.p.box, nil
}

func (p *pieChart) Render() (Box, error) {
	opt := p.opt

	renderResult, err := defaultRender(p.p, defaultRenderOption{
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
		TitleOption:        opt.Title,
		LegendOption:       opt.Legend,
		backgroundIsFilled: opt.backgroundIsFilled,
	})
	if err != nil {
		return BoxZero, err
	}
	seriesList := opt.SeriesList.Filter(ChartTypePie)
	return p.render(renderResult, seriesList)
}
