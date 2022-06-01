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

func getPieStyle(theme *Theme, index int) chart.Style {
	seriesColor := theme.GetSeriesColor(index)
	return chart.Style{
		StrokeColor: seriesColor,
		StrokeWidth: 1,
		FillColor:   seriesColor,
	}
}

type pieChartOption struct {
	Theme      string
	Font       *truetype.Font
	SeriesList SeriesList
}

func pieChartRender(opt pieChartOption, result *basicRenderResult) error {
	d, err := NewDraw(DrawOption{
		Parent: result.d,
	}, PaddingOption(chart.Box{
		Top: result.titleBox.Height(),
	}))
	if err != nil {
		return err
	}

	values := make([]float64, len(opt.SeriesList))
	total := float64(0)
	radiusValue := ""
	for index, series := range opt.SeriesList {
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
		return errors.New("The sum value of pie chart should gt 0")
	}
	r := d.Render
	theme := NewTheme(opt.Theme)

	box := d.Box
	cx := box.Width() >> 1
	cy := box.Height() >> 1

	diameter := chart.MinInt(box.Width(), box.Height())
	radius := getRadius(float64(diameter), radiusValue)

	labelLineWidth := 15
	if radius < 50 {
		labelLineWidth = 10
	}
	labelRadius := radius + float64(labelLineWidth)

	seriesNames := opt.SeriesList.Names()

	if len(values) == 1 {
		getPieStyle(theme, 0).WriteToRenderer(r)
		d.moveTo(cx, cy)
		d.circle(radius, cx, cy)
	} else {
		currentValue := float64(0)
		prevEndX := 0
		prevEndY := 0
		for index, v := range values {

			pieStyle := getPieStyle(theme, index)
			pieStyle.WriteToRenderer(r)
			d.moveTo(cx, cy)
			start := chart.PercentToRadians(currentValue/total) - math.Pi/2
			currentValue += v
			percent := (v / total)
			delta := chart.PercentToRadians(percent)
			d.arcTo(cx, cy, radius, radius, start, delta)
			d.lineTo(cx, cy)
			r.Close()
			r.FillStroke()

			series := opt.SeriesList[index]
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
			if index != 0 &&
				math.Abs(float64(endx-prevEndX)) < labelFontSize &&
				math.Abs(float64(endy-prevEndY)) < labelFontSize {
				endy -= (labelFontSize << 1)
			}
			prevEndX = endx
			prevEndY = endy
			d.moveTo(startx, starty)
			d.lineTo(endx, endy)
			offset := labelLineWidth
			if endx < cx {
				offset *= -1
			}
			d.moveTo(endx, endy)
			endx += offset
			d.lineTo(endx, endy)
			r.Stroke()
			textStyle := chart.Style{
				FontColor: theme.GetTextColor(),
				FontSize:  labelFontSize,
				Font:      opt.Font,
			}
			if !series.Label.Color.IsZero() {
				textStyle.FontColor = series.Label.Color
			}
			textStyle.GetTextOptions().WriteToRenderer(r)
			text := NewPieLabelFormatter(seriesNames, series.Label.Formatter)(index, v, percent)
			textBox := r.MeasureText(text)
			textMargin := 3
			x := endx + textMargin
			y := endy + textBox.Height()>>1 - 1
			if offset < 0 {
				textWidth := textBox.Width()
				x = endx - textWidth - textMargin
			}
			d.text(text, x, y)
		}
	}
	return nil
}
