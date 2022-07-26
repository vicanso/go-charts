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
)

type legendPainter struct {
	p   *Painter
	opt *LegendOption
}

const IconRect = "rect"
const IconLineDot = "lineDot"

type LegendOption struct {
	// The theme
	Theme ColorPalette
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
	// Icon of the legend.
	Icon string
	// Font size of legend text
	FontSize float64
	// FontColor color of legend text
	FontColor Color
	// The flag for show legend, set this to *false will hide legend
	Show *bool
	// The padding of legend
	Padding Box
}

// NewLegendOption returns a legend option
func NewLegendOption(labels []string, left ...string) LegendOption {
	opt := LegendOption{
		Data: labels,
	}
	if len(left) != 0 {
		opt.Left = left[0]
	}
	return opt
}

// IsEmpty checks legend is empty
func (opt *LegendOption) IsEmpty() bool {
	isEmpty := true
	for _, v := range opt.Data {
		if v != "" {
			isEmpty = false
			break
		}
	}
	return isEmpty
}

// NewLegendPainter returns a legend renderer
func NewLegendPainter(p *Painter, opt LegendOption) *legendPainter {
	return &legendPainter{
		p:   p,
		opt: &opt,
	}
}

func (l *legendPainter) Render() (Box, error) {
	opt := l.opt
	theme := opt.Theme
	if opt.IsEmpty() ||
		isFalse(opt.Show) {
		return BoxZero, nil
	}
	if theme == nil {
		theme = l.p.theme
	}
	if opt.FontSize == 0 {
		opt.FontSize = theme.GetFontSize()
	}
	if opt.FontColor.IsZero() {
		opt.FontColor = theme.GetTextColor()
	}
	if opt.Left == "" {
		opt.Left = PositionCenter
	}
	padding := opt.Padding
	if padding.IsZero() {
		padding.Top = 5
	}
	p := l.p.Child(PainterPaddingOption(padding))
	p.SetTextStyle(Style{
		FontSize:  opt.FontSize,
		FontColor: opt.FontColor,
	})
	measureList := make([]Box, len(opt.Data))
	maxTextWidth := 0
	for index, text := range opt.Data {
		b := p.MeasureText(text)
		if b.Width() > maxTextWidth {
			maxTextWidth = b.Width()
		}
		measureList[index] = b
	}

	// 计算展示的宽高
	width := 0
	height := 0
	offset := 20
	textOffset := 2
	legendWidth := 30
	legendHeight := 20
	itemMaxHeight := 0
	for _, item := range measureList {
		if item.Height() > itemMaxHeight {
			itemMaxHeight = item.Height()
		}
		if opt.Orient == OrientVertical {
			height += item.Height()
		} else {
			width += item.Width()
		}
	}
	// 增加padding
	itemMaxHeight += 10
	if opt.Orient == OrientVertical {
		width = maxTextWidth + textOffset + legendWidth
		height = offset * len(opt.Data)
	} else {
		height = legendHeight
		offsetValue := (len(opt.Data) - 1) * (offset + textOffset)
		allLegendWidth := len(opt.Data) * legendWidth
		width += (offsetValue + allLegendWidth)
	}

	// 计算开始的位置
	left := 0
	switch opt.Left {
	case PositionRight:
		left = p.Width() - width
	case PositionCenter:
		left = (p.Width() - width) >> 1
	default:
		if strings.HasSuffix(opt.Left, "%") {
			value, _ := strconv.Atoi(strings.ReplaceAll(opt.Left, "%", ""))
			left = p.Width() * value / 100
		} else {
			value, _ := strconv.Atoi(opt.Left)
			left = value
		}
	}
	top, _ := strconv.Atoi(opt.Top)

	if left < 0 {
		left = 0
	}

	x := int(left)
	y := int(top) + 10
	startY := y
	x0 := x
	y0 := y

	drawIcon := func(top, left int) int {
		if opt.Icon == IconRect {
			p.Rect(Box{
				Top:    top - legendHeight + 8,
				Left:   left,
				Right:  left + legendWidth,
				Bottom: top + 1,
			})
		} else {
			p.LegendLineDot(Box{
				Top:    top + 1,
				Left:   left,
				Right:  left + legendWidth,
				Bottom: top + legendHeight + 1,
			})
		}
		return left + legendWidth
	}
	lastIndex := len(opt.Data) - 1
	for index, text := range opt.Data {
		color := theme.GetSeriesColor(index)
		p.SetDrawingStyle(Style{
			FillColor:   color,
			StrokeColor: color,
		})
		itemWidth := x0 + measureList[index].Width() + textOffset + offset + legendWidth
		if lastIndex == index {
			itemWidth = x0 + measureList[index].Width() + legendWidth
		}
		if itemWidth > p.Width() {
			x0 = 0
			y += itemMaxHeight
			y0 = y
		}
		if opt.Align != AlignRight {
			x0 = drawIcon(y0, x0)
			x0 += textOffset
		}
		p.Text(text, x0, y0)
		x0 += measureList[index].Width()
		if opt.Align == AlignRight {
			x0 += textOffset
			x0 = drawIcon(y0, x0)
		}
		if opt.Orient == OrientVertical {
			y0 += offset
			x0 = x
		} else {
			x0 += offset
			y0 = y
		}
		height = y0 - startY + 10
	}

	return Box{
		Right:  width,
		Bottom: height + padding.Bottom + padding.Top,
	}, nil
}
