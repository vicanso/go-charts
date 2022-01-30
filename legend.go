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

type LegendOption struct {
	Style chart.Style
	Data  []string
	Left  string
	Right string
	Align string
}

func drawLegend(p *Draw, opt *LegendOption, theme *Theme) (chart.Box, error) {
	if len(opt.Data) == 0 {
		return chart.BoxZero, nil
	}
	padding := opt.Style.Padding
	legendDraw, err := NewDraw(DrawOption{
		Parent: p,
	}, PaddingOption(padding))
	if err != nil {
		return chart.BoxZero, err
	}
	r := legendDraw.Render
	opt.Style.GetTextOptions().WriteToRenderer(r)

	x := 0
	y := 0
	legendWidth := 30
	legendDotHeight := 5
	textPadding := 5
	legendMargin := 10

	widthCount := 0
	// 文本宽度
	for _, text := range opt.Data {
		b := r.MeasureText(text)
		widthCount += b.Width()
	}
	// 加上标记
	widthCount += legendWidth * len(opt.Data)
	// 文本的padding
	widthCount += 2 * textPadding * len(opt.Data)
	// margin的宽度
	widthCount += legendMargin * (len(opt.Data) - 1)

	// TODO 支持更多的定位方式
	// 居中
	x = (legendDraw.Box.Width() - widthCount) >> 1
	for index, text := range opt.Data {
		if index != 0 {
			x += legendMargin
		}
		style := chart.Style{
			StrokeColor: theme.GetSeriesColor(index),
			FillColor:   theme.GetSeriesColor(index),
			StrokeWidth: 3,
		}
		textBox := r.MeasureText(text)
		renderText := func() {
			x += textPadding
			legendDraw.text(text, x, y+legendDotHeight-2)
			x += textBox.Width()
			x += textPadding
		}

		if opt.Align == PositionRight {
			renderText()
		}

		style.GetFillAndStrokeOptions().WriteDrawingOptionsToRenderer(r)
		legendDraw.moveTo(x, y)
		legendDraw.lineTo(x+legendWidth, y)
		r.Stroke()
		legendDraw.circle(float64(legendDotHeight), x+legendWidth>>1, y)
		r.FillStroke()
		x += legendWidth

		if opt.Align != PositionRight {
			renderText()
		}
	}
	legendBox := padding.Clone()
	legendBox.Right = legendBox.Left + x
	legendBox.Bottom = legendBox.Top + 2*legendDotHeight

	return legendBox, nil
}
