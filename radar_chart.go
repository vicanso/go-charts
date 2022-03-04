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

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type RadarIndicator struct {
	// Indicator's name
	Name string
	// The maximum value of indicator
	Max float64
	// The minimum value of indicator
	Min float64
}

type radarChartOption struct {
	Theme      string
	Font       *truetype.Font
	SeriesList SeriesList
	Indicators []RadarIndicator
}

func radarChartRender(opt radarChartOption, result *basicRenderResult) error {
	sides := len(opt.Indicators)
	if sides < 3 {
		return errors.New("The count of indicator should be >= 3")
	}
	maxValues := make([]float64, len(opt.Indicators))
	for _, series := range opt.SeriesList {
		for index, item := range series.Data {
			if index < len(maxValues) && item.Value > maxValues[index] {
				maxValues[index] = item.Value
			}
		}
	}
	for index, indicator := range opt.Indicators {
		if indicator.Max <= 0 {
			opt.Indicators[index].Max = maxValues[index]
		}
	}
	d, err := NewDraw(DrawOption{
		Parent: result.d,
	}, PaddingOption(chart.Box{
		Top: result.titleBox.Height(),
	}))
	if err != nil {
		return err
	}
	radiusValue := ""
	for _, series := range opt.SeriesList {
		if len(series.Radius) != 0 {
			radiusValue = series.Radius
		}
	}

	box := d.Box
	cx := box.Width() >> 1
	cy := box.Height() >> 1
	diameter := chart.MinInt(box.Width(), box.Height())
	radius := getRadius(float64(diameter), radiusValue)

	theme := NewTheme(opt.Theme)

	divideCount := 5
	divideRadius := float64(int(radius / float64(divideCount)))
	radius = divideRadius * float64(divideCount)

	style := chart.Style{
		StrokeColor: theme.GetAxisSplitLineColor(),
		StrokeWidth: 1,
	}
	r := d.Render
	style.WriteToRenderer(r)
	center := Point{
		X: cx,
		Y: cy,
	}
	for i := 0; i < divideCount; i++ {
		d.polygon(center, divideRadius*float64(i+1), sides)
	}
	points := getPolygonPoints(center, radius, sides)
	for _, p := range points {
		d.moveTo(center.X, center.Y)
		d.lineTo(p.X, p.Y)
		d.Render.Stroke()
	}
	// 文本
	textStyle := chart.Style{
		FontColor: theme.GetTextColor(),
		FontSize:  labelFontSize,
		Font:      opt.Font,
	}
	textStyle.GetTextOptions().WriteToRenderer(r)
	offset := 5
	// 文本生成
	for index, p := range points {
		name := opt.Indicators[index].Name
		b := r.MeasureText(name)
		isXCenter := p.X == center.X
		isYCenter := p.Y == center.Y
		isRight := p.X > center.X
		isLeft := p.X < center.X
		isTop := p.Y < center.Y
		isBottom := p.Y > center.Y
		x := p.X
		y := p.Y
		if isXCenter {
			x -= b.Width() >> 1
			if isTop {
				y -= b.Height()
			} else {
				y += b.Height()
			}
		}
		if isYCenter {
			y += b.Height() >> 1
		}
		if isTop {
			y += offset
		}
		if isBottom {
			y += offset
		}
		if isRight {
			x += offset
		}
		if isLeft {
			x -= (b.Width() + offset)
		}
		d.text(name, x, y)
	}

	// 雷达图
	angles := getPolygonPointAngles(sides)
	maxCount := len(opt.Indicators)
	for _, series := range opt.SeriesList {
		linePoints := make([]Point, 0, maxCount)
		for j, item := range series.Data {
			if j >= maxCount {
				continue
			}
			indicator := opt.Indicators[j]
			percent := (item.Value - indicator.Min) / (indicator.Max - indicator.Min)
			r := percent * radius
			p := getPolygonPoint(center, r, angles[j])
			linePoints = append(linePoints, p)
		}
		color := theme.GetSeriesColor(series.index)
		dotFillColor := drawing.ColorWhite
		if theme.IsDark() {
			dotFillColor = color
		}
		linePoints = append(linePoints, linePoints[0])
		s := LineStyle{
			StrokeColor:  color,
			StrokeWidth:  defaultStrokeWidth,
			DotWidth:     defaultDotWidth,
			DotColor:     color,
			DotFillColor: dotFillColor,
			FillColor:    color.WithAlpha(20),
		}
		d.lineStroke(linePoints, s)
		d.fill(linePoints, s.Style())
		d.lineDot(linePoints[0:len(linePoints)-1], s)
	}
	return nil
}
