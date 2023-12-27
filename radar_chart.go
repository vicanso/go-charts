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

	"github.com/dustin/go-humanize"
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type radarChart struct {
	p   *Painter
	opt *RadarChartOption
}

type RadarIndicator struct {
	// Indicator's name
	Name string
	// The maximum value of indicator
	Max float64
	// The minimum value of indicator
	Min float64
}

type RadarChartOption struct {
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
	// The radar indicator list
	RadarIndicators []RadarIndicator
	// background is filled
	backgroundIsFilled bool
}

// NewRadarIndicators returns a radar indicator list
func NewRadarIndicators(names []string, values []float64) []RadarIndicator {
	if len(names) != len(values) {
		return nil
	}
	indicators := make([]RadarIndicator, len(names))
	for index, name := range names {
		indicators[index] = RadarIndicator{
			Name: name,
			Max:  values[index],
		}
	}
	return indicators
}

// NewRadarChart returns a radar chart renderer
func NewRadarChart(p *Painter, opt RadarChartOption) *radarChart {
	if opt.Theme == nil {
		opt.Theme = defaultTheme
	}
	return &radarChart{
		p:   p,
		opt: &opt,
	}
}

func (r *radarChart) render(result *defaultRenderResult, seriesList SeriesList) (Box, error) {
	opt := r.opt
	indicators := opt.RadarIndicators
	sides := len(indicators)
	if sides < 3 {
		return BoxZero, errors.New("The count of indicator should be >= 3")
	}
	maxValues := make([]float64, len(indicators))
	for _, series := range seriesList {
		for index, item := range series.Data {
			if index < len(maxValues) && item.Value > maxValues[index] {
				maxValues[index] = item.Value
			}
		}
	}
	for index, indicator := range indicators {
		if indicator.Max <= 0 {
			indicators[index].Max = maxValues[index]
		}
	}

	radiusValue := ""
	for _, series := range seriesList {
		if len(series.Radius) != 0 {
			radiusValue = series.Radius
		}
	}

	seriesPainter := result.seriesPainter
	theme := opt.Theme

	cx := seriesPainter.Width() >> 1
	cy := seriesPainter.Height() >> 1
	diameter := chart.MinInt(seriesPainter.Width(), seriesPainter.Height())
	radius := getRadius(float64(diameter), radiusValue)

	divideCount := 5
	divideRadius := float64(int(radius / float64(divideCount)))
	radius = divideRadius * float64(divideCount)

	seriesPainter.OverrideDrawingStyle(Style{
		StrokeColor: theme.GetAxisSplitLineColor(),
		StrokeWidth: 1,
	})
	center := Point{
		X: cx,
		Y: cy,
	}
	for i := 0; i < divideCount; i++ {
		seriesPainter.Polygon(center, divideRadius*float64(i+1), sides)
	}
	points := getPolygonPoints(center, radius, sides)
	for _, p := range points {
		seriesPainter.MoveTo(center.X, center.Y)
		seriesPainter.LineTo(p.X, p.Y)
		seriesPainter.Stroke()
	}
	seriesPainter.OverrideTextStyle(Style{
		FontColor: theme.GetTextColor(),
		FontSize:  labelFontSize,
		Font:      opt.Font,
	})
	offset := 5
	// 文本生成
	for index, p := range points {
		name := indicators[index].Name
		b := seriesPainter.MeasureText(name)
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
		seriesPainter.Text(name, x, y)
	}

	// 雷达图
	angles := getPolygonPointAngles(sides)
	maxCount := len(indicators)
	for _, series := range seriesList {
		linePoints := make([]Point, 0, maxCount)
		for j, item := range series.Data {
			if j >= maxCount {
				continue
			}
			indicator := indicators[j]
			var percent float64
			offset := indicator.Max - indicator.Min
			if offset > 0 {
				percent = (item.Value - indicator.Min) / offset
			}
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
		seriesPainter.OverrideDrawingStyle(Style{
			StrokeColor: color,
			StrokeWidth: defaultStrokeWidth,
			DotWidth:    defaultDotWidth,
			DotColor:    color,
			FillColor:   color.WithAlpha(20),
		})
		seriesPainter.LineStroke(linePoints).
			FillArea(linePoints)
		dotWith := 2.0
		seriesPainter.OverrideDrawingStyle(Style{
			StrokeWidth: defaultStrokeWidth,
			StrokeColor: color,
			FillColor:   dotFillColor,
		})
		for index, point := range linePoints {
			seriesPainter.Circle(dotWith, point.X, point.Y)
			seriesPainter.FillStroke()
			if series.Label.Show && index < len(series.Data) {
				value := humanize.FtoaWithDigits(series.Data[index].Value, 2)
				b := seriesPainter.MeasureText(value)
				seriesPainter.Text(value, point.X-b.Width()/2, point.Y)
			}

		}
	}

	return r.p.box, nil
}

func (r *radarChart) Render() (Box, error) {
	p := r.p
	opt := r.opt
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
		TitleOption:        opt.Title,
		LegendOption:       opt.Legend,
		backgroundIsFilled: opt.backgroundIsFilled,
	})
	if err != nil {
		return BoxZero, err
	}
	seriesList := opt.SeriesList.Filter(ChartTypeRadar)
	return r.render(renderResult, seriesList)
}
