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
	"math"

	"github.com/wcharczuk/go-chart/v2"
)

const defaultRadiusPercent = 0.4

func getPieStyle(theme *Theme, index int) chart.Style {
	seriesColor := theme.GetSeriesColor(index)
	return chart.Style{
		StrokeColor: seriesColor,
		StrokeWidth: 1,
		FillColor:   seriesColor,
	}

}

func pieChartRender(opt ChartOption, result *basicRenderResult) (*Draw, error) {
	d, err := NewDraw(DrawOption{
		Parent: result.d,
	}, PaddingOption(chart.Box{
		Top: result.titleBox.Height(),
	}))
	if err != nil {
		return nil, err
	}

	values := make([]float64, len(opt.SeriesList))
	total := float64(0)
	for index, series := range opt.SeriesList {
		value := float64(0)
		for _, item := range series.Data {
			value += item.Value
		}
		values[index] = value
		total += value
	}
	r := d.Render
	theme := NewTheme(opt.Theme)

	box := d.Box
	cx := box.Width() >> 1
	cy := box.Height() >> 1

	diameter := chart.MinInt(box.Width(), box.Height())
	radius := float64(diameter) * defaultRadiusPercent
	labelRadius := radius + 20

	if len(values) == 1 {
		getPieStyle(theme, 0).WriteToRenderer(r)
		d.moveTo(cx, cy)
		d.circle(radius, cx, cy)
	} else {
		currentValue := float64(0)
		for index, v := range values {
			getPieStyle(theme, index).WriteToRenderer(r)
			d.moveTo(cx, cy)
			start := chart.PercentToRadians(currentValue/total) - math.Pi/2
			currentValue += v
			delta := chart.PercentToRadians(v / total)
			d.arcTo(cx, cy, radius, radius, start, delta)
			d.lineTo(cx, cy)
			r.Close()
			r.FillStroke()

			// label的角度为饼块中间
			angle := start + delta/2
			startx := cx + int(radius*math.Cos(angle))
			starty := cy + int(radius*math.Sin(angle))

			endx := cx + int(labelRadius*math.Cos(angle))
			endy := cy + int(labelRadius*math.Sin(angle))
			d.moveTo(startx, starty)
			d.lineTo(endx, endy)
			offset := 30
			if endx < cx {
				offset *= -1
			}
			d.moveTo(endx, endy)
			d.lineTo(endx+offset, endy)
			r.Stroke()
			// TODO label show
		}
	}
	return result.d, nil
}
