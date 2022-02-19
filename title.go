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
	// Title text, support \n for new line
	Text string
	// Subtitle text, support \n for new line
	Subtext string
	// Title style
	Style chart.Style
	// Subtitle style
	SubtextStyle chart.Style
	// Distance between title component and the left side of the container.
	// It can be pixel value: 20, percentage value: 20%,
	// or position value: right, center.
	Left string
	// Distance between title component and the top side of the container.
	// It can be pixel value: 20.
	Top string
}
type titleMeasureOption struct {
	width  int
	height int
	text   string
	style  chart.Style
}

func splitTitleText(text string) []string {
	arr := strings.Split(text, "\n")
	result := make([]string, 0)
	for _, v := range arr {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		result = append(result, v)
	}
	return result
}

func drawTitle(p *Draw, opt *TitleOption) (chart.Box, error) {
	if len(opt.Text) == 0 {
		return chart.BoxZero, nil
	}

	padding := opt.Style.Padding
	d, err := NewDraw(DrawOption{
		Parent: p,
	}, PaddingOption(padding))
	if err != nil {
		return chart.BoxZero, err
	}

	r := d.Render

	measureOptions := make([]titleMeasureOption, 0)

	// 主标题
	for _, v := range splitTitleText(opt.Text) {
		measureOptions = append(measureOptions, titleMeasureOption{
			text:  v,
			style: opt.Style.GetTextOptions(),
		})
	}
	// 副标题
	for _, v := range splitTitleText(opt.Subtext) {
		measureOptions = append(measureOptions, titleMeasureOption{
			text:  v,
			style: opt.SubtextStyle.GetTextOptions(),
		})
	}

	textMaxWidth := 0
	textMaxHeight := 0
	width := 0
	for index, item := range measureOptions {
		item.style.WriteTextOptionsToRenderer(r)
		textBox := r.MeasureText(item.text)

		w := textBox.Width()
		h := textBox.Height()
		if w > textMaxWidth {
			textMaxWidth = w
		}
		if h > textMaxHeight {
			textMaxHeight = h
		}
		measureOptions[index].height = h
		measureOptions[index].width = w
	}
	width = textMaxWidth
	titleX := 0
	b := d.Box
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
		item.style.WriteTextOptionsToRenderer(r)
		x := titleX + (textMaxWidth-item.width)>>1
		y := titleY + item.height
		d.text(item.text, x, y)
		titleY += item.height
	}
	height := titleY + padding.Top + padding.Bottom
	box := padding.Clone()
	box.Right = box.Left + titleX + width
	box.Bottom = box.Top + height

	return box, nil
}
