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
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type lineChartOption struct {
	Theme      string
	SeriesList SeriesList
	Font       *truetype.Font
}

func lineChartRender(opt lineChartOption, result *basicRenderResult) ([]markPointRenderOption, error) {

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
	seriesNames := opt.SeriesList.Names()

	r := d.Render
	xRange := result.xRange
	markPointRenderOptions := make([]markPointRenderOption, 0)
	for i, s := range opt.SeriesList {
		// 由于series是for range，为同一个数据，因此需要clone
		// 后续需要使用，如mark point
		series := s
		index := series.index
		if index == 0 {
			index = i
		}
		seriesColor := theme.GetSeriesColor(index)

		yRange := result.getYRange(series.YAxisIndex)
		points := make([]Point, 0, len(series.Data))
		// mark line
		markLineRender(markLineRenderOption{
			Draw:        d,
			FillColor:   seriesColor,
			FontColor:   theme.GetTextColor(),
			StrokeColor: seriesColor,
			Font:        opt.Font,
			Series:      &series,
			Range:       yRange,
		})

		for j, item := range series.Data {
			if j >= xRange.divideCount {
				continue
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
			distance := series.Label.Distance
			if distance == 0 {
				distance = 5
			}
			text := NewValueLabelFormater(seriesNames, series.Label.Formatter)(i, item.Value, -1)
			labelStyle := chart.Style{
				FontColor: theme.GetTextColor(),
				FontSize:  labelFontSize,
				Font:      opt.Font,
			}
			if !series.Label.Color.IsZero() {
				labelStyle.FontColor = series.Label.Color
			}
			labelStyle.GetTextOptions().WriteToRenderer(r)
			textBox := r.MeasureText(text)
			d.text(text, x-textBox.Width()>>1, y-distance)
		}

		dotFillColor := drawing.ColorWhite
		if theme.IsDark() {
			dotFillColor = seriesColor
		}
		d.Line(points, LineStyle{
			StrokeColor:  seriesColor,
			StrokeWidth:  2,
			DotColor:     seriesColor,
			DotWidth:     defaultDotWidth,
			DotFillColor: dotFillColor,
		})
		// draw mark point
		markPointRenderOptions = append(markPointRenderOptions, markPointRenderOption{
			Draw:      d,
			FillColor: seriesColor,
			Font:      opt.Font,
			Points:    points,
			Series:    &series,
		})
	}

	return markPointRenderOptions, nil
}
