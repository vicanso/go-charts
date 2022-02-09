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
	"github.com/wcharczuk/go-chart/v2"
)

func barChartRender(opt ChartOption, result *basicRenderResult) (*Draw, error) {

	d, err := NewDraw(DrawOption{
		Parent: result.d,
	}, PaddingOption(chart.Box{
		Top:  result.titleBox.Height(),
		Left: YAxisWidth,
	}))
	if err != nil {
		return nil, err
	}
	// yRange := result.yRange
	xRange := result.xRange
	x0, x1 := xRange.GetRange(0)
	width := int(x1 - x0)
	// 每一块之间的margin
	margin := 10
	// 每一个bar之间的margin
	barMargin := 5

	seriesCount := len(opt.SeriesList)
	// 总的宽度-两个margin-(总数-1)的barMargin
	barWidth := (width - 2*margin - barMargin*(seriesCount-1)) / len(opt.SeriesList)

	barMaxHeight := result.getYRange(0).Size
	theme := NewTheme(opt.Theme)

	seriesNames := opt.Legend.Data

	r := d.Render

	for i, series := range opt.SeriesList {
		yRange := result.getYRange(series.YAxisIndex)
		points := make([]Point, len(series.Data))
		seriesColor := theme.GetSeriesColor(i)
		for j, item := range series.Data {
			x0, _ := xRange.GetRange(j)
			x := int(x0)
			x += margin
			if i != 0 {
				x += i * (barWidth + barMargin)
			}

			h := int(yRange.getHeight(item.Value))
			fillColor := seriesColor
			if !item.Style.FillColor.IsZero() {
				fillColor = item.Style.FillColor
			}
			top := barMaxHeight - h
			d.Bar(chart.Box{
				Top:    top,
				Left:   x,
				Right:  x + barWidth,
				Bottom: barMaxHeight - 1,
			}, BarStyle{
				FillColor: fillColor,
			})
			// 用于生成marker point
			points[j] = Point{
				// 居中的位置
				X: x + barWidth>>1,
				Y: top,
			}
			if !series.Label.Show {
				continue
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
			d.text(text, x+(barWidth-textBox.Width())>>1, barMaxHeight-h-5)
		}
		markPointRender(d, markPointRenderOption{
			FillColor: seriesColor,
			Font:      opt.Font,
			Points:    points,
			Series:    &series,
		})
	}

	return result.d, nil
}
