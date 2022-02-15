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
	"strconv"
	"strings"

	"github.com/wcharczuk/go-chart/v2"
)

type LegendOption struct {
	theme string
	// Legend show flag, if nil or true, the legend will be shown
	Show *bool
	// Legend text style
	Style chart.Style
	// Text array of legend
	Data []string
	// Distance between legend component and the left side of the container.
	// It can be pixel value: 20, percentage value: 20%,
	// or position value: right, center.
	Left string
	// Distance between legend component and the top side of the container.
	// It can be pixel value: 20.
	Top string
	// Legend marker and text aligning, it can be left or right, default is left.
	Align string
	// The layout orientation of legend, it can be horizontal or vertical, default is horizontal.
	Orient string
}

// NewLegendOption creates a new legend option by legend text list
func NewLegendOption(data []string, position ...string) LegendOption {
	opt := LegendOption{
		Data: data,
	}
	if len(position) != 0 {
		opt.Left = position[0]
	}
	return opt
}

type legend struct {
	d   *Draw
	opt *LegendOption
}

func NewLegend(d *Draw, opt LegendOption) *legend {
	return &legend{
		d:   d,
		opt: &opt,
	}
}

func (l *legend) Render() (chart.Box, error) {
	d := l.d
	opt := l.opt
	if len(opt.Data) == 0 || isFalse(opt.Show) {
		return chart.BoxZero, nil
	}
	theme := NewTheme(opt.theme)
	padding := opt.Style.Padding
	legendDraw, err := NewDraw(DrawOption{
		Parent: d,
	}, PaddingOption(padding))
	if err != nil {
		return chart.BoxZero, err
	}
	r := legendDraw.Render
	opt.Style.GetTextOptions().WriteToRenderer(r)

	x := 0
	y := 0
	top := 0
	// TODO TOP 暂只支持数值
	if opt.Top != "" {
		top, _ = strconv.Atoi(opt.Top)
		y += top
	}
	legendWidth := 30
	legendDotHeight := 5
	textPadding := 5
	legendMargin := 10
	// 往下移2倍dot的高度
	y += 2 * legendDotHeight

	widthCount := 0
	maxTextWidth := 0
	// 文本宽度
	for _, text := range opt.Data {
		b := r.MeasureText(text)
		if b.Width() > maxTextWidth {
			maxTextWidth = b.Width()
		}
		widthCount += b.Width()
	}
	if opt.Orient == OrientVertical {
		widthCount = maxTextWidth + legendWidth + textPadding
	} else {
		// 加上标记
		widthCount += legendWidth * len(opt.Data)
		// 文本的padding
		widthCount += 2 * textPadding * len(opt.Data)
		// margin的宽度
		widthCount += legendMargin * (len(opt.Data) - 1)
	}

	left := 0
	switch opt.Left {
	case PositionRight:
		left = legendDraw.Box.Width() - widthCount
	case PositionCenter:
		left = (legendDraw.Box.Width() - widthCount) >> 1
	default:
		if strings.HasSuffix(opt.Left, "%") {
			value, _ := strconv.Atoi(strings.ReplaceAll(opt.Left, "%", ""))
			left = legendDraw.Box.Width() * value / 100
		} else {
			value, _ := strconv.Atoi(opt.Left)
			left = value
		}
	}
	x = left
	for index, text := range opt.Data {
		seriesColor := theme.GetSeriesColor(index)
		fillColor := seriesColor
		if !theme.IsDark() {
			fillColor = theme.GetBackgroundColor()
		}
		style := chart.Style{
			StrokeColor: seriesColor,
			FillColor:   fillColor,
			StrokeWidth: 3,
		}
		style.GetFillAndStrokeOptions().WriteDrawingOptionsToRenderer(r)

		textBox := r.MeasureText(text)
		var renderText func()
		if opt.Orient == OrientVertical {
			// 垂直
			// 重置x的位置
			x = left
			renderText = func() {
				x += textPadding
				legendDraw.text(text, x, y+legendDotHeight)
				x += textBox.Width()
				y += (2*legendDotHeight + legendMargin)
			}

		} else {
			// 水平
			if index != 0 {
				x += legendMargin
			}
			renderText = func() {
				x += textPadding
				legendDraw.text(text, x, y+legendDotHeight)
				x += textBox.Width()
				x += textPadding
			}
		}
		if opt.Align == PositionRight {
			renderText()
		}

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
	// 计算展示区域
	if opt.Orient == OrientVertical {
		legendBox.Right = legendBox.Left + left + maxTextWidth + legendWidth + textPadding
		legendBox.Bottom = legendBox.Top + y
	} else {
		legendBox.Right = legendBox.Left + x
		legendBox.Bottom = legendBox.Top + 2*legendDotHeight + top + textPadding
	}
	return legendBox, nil
}
