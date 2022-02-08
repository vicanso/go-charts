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
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func lineChartRender(opt ChartOption, result *basicRenderResult) (*Draw, error) {

	theme := NewTheme(opt.Theme)

	d, err := NewDraw(DrawOption{
		Parent: result.d,
	}, PaddingOption(chart.Box{
		Top:  result.titleBox.Height(),
		Left: YAxisWidth,
	}))
	if err != nil {
		return nil, err
	}
	seriesNames := opt.Legend.Data

	r := d.Render
	yRange := result.yRange
	xRange := result.xRange
	for i, series := range opt.SeriesList {
		points := make([]Point, 0)
		minIndex := -1
		maxIndex := -1
		minValue := math.MaxFloat64
		maxValue := -math.MaxFloat64
		for j, item := range series.Data {
			if item.Value < minValue {
				minIndex = j
				minValue = item.Value
			}
			if item.Value > maxValue {
				maxIndex = j
				maxValue = item.Value
			}
			y := yRange.getRestHeight(item.Value)
			x := xRange.getWidth(float64(j))
			points = append(points, Point{
				Y: y,
				X: x,
			})
			if !series.Label.Show {
				continue
			}
			text := NewValueLabelFormater(seriesNames, series.Label.Formatter)(i, item.Value, -1)
			labelStyle := chart.Style{
				FontColor: theme.GetTextColor(),
				FontSize:  10,
				Font:      opt.Font,
			}
			if !series.Label.Color.IsZero() {
				labelStyle.FontColor = series.Label.Color
			}
			labelStyle.GetTextOptions().WriteToRenderer(r)
			textBox := r.MeasureText(text)
			d.text(text, x-textBox.Width()>>1, y-5)
		}
		index := series.index
		if index == 0 {
			index = i
		}
		seriesColor := theme.GetSeriesColor(index)
		dotFillColor := drawing.ColorWhite
		if theme.IsDark() {
			dotFillColor = seriesColor
		}
		d.Line(points, LineStyle{
			StrokeColor:  seriesColor,
			StrokeWidth:  2,
			DotColor:     seriesColor,
			DotWidth:     2,
			DotFillColor: dotFillColor,
		})
		// draw mark point
		symbolSize := 30
		if series.MarkPoint.SymbolSize > 0 {
			symbolSize = series.MarkPoint.SymbolSize
		}
		for _, markPointData := range series.MarkPoint.Data {
			p := points[minIndex]
			value := minValue
			switch markPointData.Type {
			case SeriesMarkPointDataTypeMax:
				p = points[maxIndex]
				value = maxValue
			}
			chart.Style{
				FillColor: seriesColor,
			}.WriteToRenderer(r)
			d.pin(p.X, p.Y-symbolSize>>1, symbolSize)

			chart.Style{
				FontColor:   NewTheme(ThemeDark).GetTextColor(),
				FontSize:    10,
				StrokeWidth: 1,
				Font:        opt.Font,
			}.WriteTextOptionsToRenderer(d.Render)
			text := commafWithDigits(value)
			textBox := r.MeasureText(text)
			d.text(text, p.X-textBox.Width()>>1, p.Y-symbolSize>>1-2)
		}
	}

	return result.d, nil
}
