// MIT License

// Copyright (c) 2021 Tree Xie

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
	Style    chart.Style
	Padding  chart.Box
	Left     string
	Right    string
	Top      string
	Bottom   string
	Align    string
	Theme    string
	IconDraw LegendIconDraw
}

type LegendIconDrawOption struct {
	Box   chart.Box
	Style chart.Style
	Index int
	Theme string
}

const (
	LegendAlignLeft  = "left"
	LegendAlignRight = "right"
)

type LegendIconDraw func(r chart.Renderer, opt LegendIconDrawOption)

func DefaultLegendIconDraw(r chart.Renderer, opt LegendIconDrawOption) {
	if opt.Box.IsZero() {
		return
	}
	r.SetStrokeColor(opt.Style.GetStrokeColor())
	strokeWidth := opt.Style.GetStrokeWidth()
	r.SetStrokeWidth(strokeWidth)
	height := opt.Box.Bottom - opt.Box.Top
	ly := opt.Box.Top - (height / 2) + 2
	r.MoveTo(opt.Box.Left, ly)
	r.LineTo(opt.Box.Right, ly)
	r.Stroke()
	r.SetFillColor(getBackgroundColor(opt.Theme))
	r.Circle(5, (opt.Box.Left+opt.Box.Right)/2, ly)
	r.FillStroke()
}

func covertPercent(value string) float64 {
	if !strings.HasSuffix(value, "%") {
		return -1
	}
	v, err := strconv.Atoi(strings.ReplaceAll(value, "%", ""))
	if err != nil {
		return -1
	}
	return float64(v) / 100
}

func getLegendLeft(width, legendBoxWidth int, opt LegendOption) int {
	left := (width - legendBoxWidth) / 2
	leftValue := opt.Left
	if leftValue == "auto" || leftValue == "center" {
		leftValue = ""
	}
	if leftValue == "left" {
		leftValue = "0"
	}

	rightValue := opt.Right
	if rightValue == "auto" || leftValue == "center" {
		rightValue = ""
	}
	if rightValue == "right" {
		rightValue = "0"
	}
	if leftValue == "" && rightValue == "" {
		return left
	}
	if leftValue != "" {
		percent := covertPercent(leftValue)
		if percent >= 0 {
			return int(float64(width) * percent)
		}
		v, _ := strconv.Atoi(leftValue)
		return v
	}
	if rightValue != "" {
		percent := covertPercent(rightValue)
		if percent >= 0 {
			return width - legendBoxWidth - int(float64(width)*percent)
		}
		v, _ := strconv.Atoi(rightValue)
		return width - legendBoxWidth - v
	}
	return left
}

func getLegendTop(height, legendBoxHeight int, opt LegendOption) int {
	// TODO 支持top的处理
	return 0
}

func LegendCustomize(series []chart.Series, opt LegendOption) chart.Renderable {
	return func(r chart.Renderer, cb chart.Box, chartDefaults chart.Style) {
		legendDefaults := chart.Style{
			FontColor:   getTextColor(opt.Theme),
			FontSize:    8.0,
			StrokeColor: chart.DefaultAxisColor,
		}

		legendStyle := opt.Style.InheritFrom(chartDefaults.InheritFrom(legendDefaults))

		r.SetFont(legendStyle.GetFont())
		r.SetFontColor(legendStyle.GetFontColor())
		r.SetFontSize(legendStyle.GetFontSize())

		var labels []string
		var lines []chart.Style
		// 计算label和lines
		for _, s := range series {
			if !s.GetStyle().Hidden {
				if _, isAnnotationSeries := s.(chart.AnnotationSeries); !isAnnotationSeries {
					labels = append(labels, s.GetName())
					lines = append(lines, s.GetStyle())
				}
			}
		}

		var textHeight int
		var textWidth int
		var textBox chart.Box
		labelWidth := 0
		// 计算文本宽度与高度（取最大值）
		for x := 0; x < len(labels); x++ {
			if len(labels[x]) > 0 {
				textBox = r.MeasureText(labels[x])
				labelWidth += textBox.Width()
				textHeight = chart.MaxInt(textBox.Height(), textHeight)
				textWidth = chart.MaxInt(textBox.Width(), textWidth)
			}
		}

		legendBoxHeight := textHeight + legendStyle.Padding.Top + legendStyle.Padding.Bottom
		chartPadding := cb.Top
		legendYMargin := (chartPadding - legendBoxHeight) >> 1

		iconWidth := 25
		lineTextGap := 5

		iconAllWidth := iconWidth * len(labels)
		spaceAllWidth := chart.DefaultMinimumTickHorizontalSpacing * (len(labels) - 1)

		legendBoxWidth := labelWidth + iconAllWidth + spaceAllWidth

		left := getLegendLeft(cb.Width(), legendBoxWidth, opt)
		top := getLegendTop(cb.Height(), legendBoxHeight, opt)

		left += opt.Padding.Left
		top += opt.Padding.Top

		legendBox := chart.Box{
			Left:   left,
			Right:  left + legendBoxWidth,
			Top:    top,
			Bottom: top + legendBoxHeight,
		}

		chart.Draw.Box(r, legendBox, legendDefaults)

		r.SetFont(legendStyle.GetFont())
		r.SetFontColor(legendStyle.GetFontColor())
		r.SetFontSize(legendStyle.GetFontSize())

		startX := legendBox.Left + legendStyle.Padding.Left
		ty := top + legendYMargin + legendStyle.Padding.Top + textHeight
		var label string
		var x int
		iconDraw := opt.IconDraw
		if iconDraw == nil {
			iconDraw = DefaultLegendIconDraw
		}
		align := opt.Align
		if align == "" {
			align = LegendAlignLeft
		}
		for index := range labels {
			label = labels[index]
			if len(label) > 0 {
				x = startX

				// 如果图例标记靠右展示
				if align == LegendAlignRight {
					textBox = r.MeasureText(label)
					r.Text(label, x, ty)
					x = startX + textBox.Width() + lineTextGap
				}

				// 图标
				iconDraw(r, LegendIconDrawOption{
					Theme: opt.Theme,
					Index: index,
					Style: lines[index],
					Box: chart.Box{
						Left:   x,
						Top:    ty,
						Right:  x + iconWidth,
						Bottom: ty + textHeight,
					},
				})
				x += (iconWidth + lineTextGap)

				// 如果图例标记靠左展示
				if align == LegendAlignLeft {
					textBox = r.MeasureText(label)
					r.Text(label, x, ty)
					x += textBox.Width()
				}

				// 计算下一个legend的位置
				startX = x + chart.DefaultMinimumTickHorizontalSpacing
			}
		}
	}
}
