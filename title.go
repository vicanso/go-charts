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

type titleMeasureOption struct {
	width  int
	height int
	text   string
}

func NewTitleCustomize(title Title) chart.Renderable {
	return func(r chart.Renderer, cb chart.Box, chartDefaults chart.Style) {
		if len(title.Text) == 0 || title.Style.Hidden {
			return
		}
		font := title.Font
		if font == nil {
			font, _ = chart.GetDefaultFont()
		}
		r.SetFont(font)
		r.SetFontColor(title.Style.FontColor)
		titleFontSize := title.Style.GetFontSize(chart.DefaultTitleFontSize)
		r.SetFontSize(titleFontSize)

		arr := strings.Split(title.Text, "\n")
		textWidth := 0
		textHeight := 0
		measureOptions := make([]titleMeasureOption, len(arr))
		for index, str := range arr {
			textBox := r.MeasureText(str)

			w := textBox.Width()
			h := textBox.Height()
			if w > textWidth {
				textWidth = w
			}
			if h > textHeight {
				textHeight = h
			}
			measureOptions[index] = titleMeasureOption{
				text:   str,
				width:  w,
				height: h,
			}
		}

		titleX := 0
		switch title.Left {
		case "right":
			titleX = cb.Left + cb.Width() - textWidth
		case "center":
			titleX = cb.Left + cb.Width()>>1 - (textWidth >> 1)
		default:
			if strings.HasSuffix(title.Left, "%") {
				value, _ := strconv.Atoi(strings.ReplaceAll(title.Left, "%", ""))
				titleX = cb.Left + cb.Width()*value/100
			} else {
				value, _ := strconv.Atoi(title.Left)
				titleX = cb.Left + value
			}
		}

		titleY := cb.Top + title.Style.Padding.GetTop(chart.DefaultTitleTop) + (textHeight >> 1)

		for _, item := range measureOptions {
			x := titleX + (textWidth-item.width)>>1
			r.Text(item.text, x, titleY)
			titleY += textHeight
		}
	}
}
