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
	"github.com/wcharczuk/go-chart/v2"
)

type LegendOption struct {
	Style        chart.Style
	Align        string
	TextPosition string
	Theme        string
	IconDraw     LegendIconDraw
}

type LegendIconDrawOption struct {
	Box   chart.Box
	Style chart.Style
	Index int
	Theme string
}

const (
	LegendAlignLeft   = "left"
	LegendAlignCenter = "center"
	LegendAlignRight  = "right"
)

const (
	LegendTextPositionRight = "right"
)

type LegendIconDraw func(r chart.Renderer, opt LegendIconDrawOption)

func DefaultLegendIconDraw(r chart.Renderer, opt LegendIconDrawOption) {
	if opt.Box.IsZero() {
		return
	}
	r.SetStrokeColor(opt.Style.GetStrokeColor())
	stokeWidth := opt.Style.GetStrokeWidth()
	r.SetStrokeWidth(stokeWidth)
	height := opt.Box.Bottom - opt.Box.Top
	ly := (height / 2) + 1
	r.MoveTo(opt.Box.Left, ly)
	r.LineTo(opt.Box.Right, ly)
	r.Stroke()
	r.SetFillColor(getBackgroundColor(opt.Theme))
	r.Circle(5, (opt.Box.Left+opt.Box.Right)/2, ly)
	r.FillStroke()
}

func LegendCustomize(c *chart.Chart, opt LegendOption) chart.Renderable {
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
		for _, s := range c.Series {
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

		lineLengthMinimum := 25

		labelWidth += lineLengthMinimum * len(labels)

		left := 0
		switch opt.Align {
		case LegendAlignLeft:
			left = 0
		case LegendAlignRight:
			left = cb.Width() - labelWidth
		default:
			left = (cb.Width() - labelWidth) / 2
		}

		legendBox := chart.Box{
			Left:   left,
			Right:  left + labelWidth,
			Top:    0,
			Bottom: 0 + legendBoxHeight,
		}

		chart.Draw.Box(r, legendBox, legendDefaults)

		r.SetFont(legendStyle.GetFont())
		r.SetFontColor(legendStyle.GetFontColor())
		r.SetFontSize(legendStyle.GetFontSize())

		lineTextGap := 5

		startX := legendBox.Left + legendStyle.Padding.Left
		ty := legendYMargin + legendStyle.Padding.Top + textHeight
		var label string
		var x int
		iconDraw := opt.IconDraw
		if iconDraw == nil {
			iconDraw = DefaultLegendIconDraw
		}
		for index := range labels {
			label = labels[index]
			if len(label) > 0 {
				x = startX

				// 如果文本靠左显示
				if opt.TextPosition != LegendTextPositionRight {
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
						Right:  x + lineLengthMinimum,
						Bottom: ty + textHeight,
					},
				})
				x += (lineLengthMinimum + lineTextGap)

				// 如果文本靠右显示
				if opt.TextPosition == LegendTextPositionRight {
					textBox = r.MeasureText(label)
					r.Text(label, x, ty)
				}

				// 计算下一个legend的位置
				startX += textBox.Width() + chart.DefaultMinimumTickHorizontalSpacing + lineTextGap + lineLengthMinimum
			}
		}
	}
}
