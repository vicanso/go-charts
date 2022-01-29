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

type TitleOption struct {
	Text  string
	Style chart.Style
	Left  string
	Top   string
}
type titleMeasureOption struct {
	width  int
	height int
	text   string
}

func drawTitle(d *Draw, opt *TitleOption) (chart.Box, error) {
	if len(opt.Text) == 0 {
		return chart.BoxZero, nil
	}

	padding := opt.Style.Padding
	titleDraw, err := NewDraw(DrawOption{
		Parent: d,
	}, PaddingOption(padding))
	if err != nil {
		return chart.BoxZero, err
	}

	r := titleDraw.Render
	opt.Style.GetTextOptions().WriteToRenderer(r)
	arr := strings.Split(opt.Text, "\n")
	textMaxWidth := 0
	textMaxHeight := 0
	width := 0
	measureOptions := make([]titleMeasureOption, len(arr))
	for index, str := range arr {
		textBox := r.MeasureText(str)

		w := textBox.Width()
		h := textBox.Height()
		if w > textMaxWidth {
			textMaxWidth = w
		}
		if h > textMaxHeight {
			textMaxHeight = h
		}
		measureOptions[index] = titleMeasureOption{
			text:   str,
			width:  w,
			height: h,
		}
	}
	width = textMaxWidth
	titleX := 0
	b := titleDraw.Box
	switch opt.Left {
	case PositionRight:
		titleX = b.Width() - textMaxWidth
	case PositionCenter:
		titleX = b.Width()>>1 - (textMaxWidth >> 1)
	default:
		if strings.HasSuffix(opt.Left, "%") {
			value, _ := strconv.Atoi(strings.ReplaceAll(opt.Left, "%", ""))
			titleX = b.Width() * value / 100
		} else {
			value, _ := strconv.Atoi(opt.Left)
			titleX = value
		}
	}
	titleY := 0
	// TODO TOP 暂只支持数值
	if opt.Top != "" {
		value, _ := strconv.Atoi(opt.Top)
		titleY += value
	}
	for _, item := range measureOptions {
		x := titleX + (textMaxWidth-item.width)>>1
		titleDraw.text(item.text, x, titleY)
		titleY += textMaxHeight
	}
	height := titleY + padding.Top + padding.Bottom
	box := padding.Clone()
	box.Right = box.Left + width
	box.Bottom = box.Top + height

	return box, nil
}
