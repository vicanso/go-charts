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
	"fmt"
	"sort"

	"github.com/dustin/go-humanize"
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
)

type funnelChartOption struct {
	Theme      string
	Font       *truetype.Font
	SeriesList SeriesList
}

func funnelChartRender(opt funnelChartOption, result *basicRenderResult) error {
	d, err := NewDraw(DrawOption{
		Parent: result.d,
	}, PaddingOption(chart.Box{
		Top: result.titleBox.Height(),
	}))
	if err != nil {
		return err
	}
	seriesList := make([]Series, len(opt.SeriesList))
	copy(seriesList, opt.SeriesList)
	sort.Slice(seriesList, func(i, j int) bool {
		// 大的数据在前
		return seriesList[i].Data[0].Value > seriesList[j].Data[0].Value
	})
	max := float64(100)
	min := float64(0)
	for _, item := range seriesList {
		if item.Max != nil {
			max = *item.Max
		}
		if item.Min != nil {
			min = *item.Min
		}
	}

	theme := NewTheme(opt.Theme)
	gap := 2
	height := d.Box.Height()
	width := d.Box.Width()
	count := len(seriesList)

	h := (height - gap*(count-1)) / count

	y := 0
	widthList := make([]int, len(seriesList))
	textList := make([]string, len(seriesList))
	for index, item := range seriesList {
		value := item.Data[0].Value
		percent := (item.Data[0].Value - min) / (max - min)
		w := int(percent * float64(width))
		widthList[index] = w
		p := humanize.CommafWithDigits(value, 2) + "%"
		textList[index] = fmt.Sprintf("%s(%s)", item.Name, p)
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
		d.fill(points, chart.Style{
			FillColor: color,
		})

		// 文本
		text := textList[index]
		r := d.Render
		textStyle := chart.Style{
			FontColor: theme.GetTextColor(),
			FontSize:  labelFontSize,
			Font:      opt.Font,
		}
		textStyle.GetTextOptions().WriteToRenderer(r)
		textBox := r.MeasureText(text)
		textX := width>>1 - textBox.Width()>>1
		textY := y + h>>1
		d.text(text, textX, textY)

		y += (h + gap)
	}

	return nil
}
