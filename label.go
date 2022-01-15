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
	"github.com/dustin/go-humanize"
	"github.com/wcharczuk/go-chart/v2"
)

type LabelRenderer struct {
	Options SeriesLabel
	Offset  chart.Box
}
type LabelValue struct {
	Left int
	Top  int
	Text string
}

func (l LabelRenderer) Render(r chart.Renderer, canvasBox chart.Box, xrange, yrange chart.Range, style chart.Style, vs chart.ValuesProvider) {
	if !l.Options.Show {
		return
	}
	r.SetFontColor(style.FontColor)
	r.SetFontSize(style.FontSize)
	r.SetFont(style.Font)
	offsetX := l.Options.Offset.Left + l.Offset.Left
	offsetY := l.Options.Offset.Top + l.Offset.Top
	for i := 0; i < vs.Len(); i++ {
		vx, vy := vs.GetValues(i)
		x := canvasBox.Left + xrange.Translate(vx) + offsetX
		y := canvasBox.Bottom - yrange.Translate(vy) + offsetY

		text := humanize.CommafWithDigits(vy, 2)
		// 往左移一半距离
		x -= r.MeasureText(text).Width() >> 1
		r.Text(text, x, y)
	}
}

func (l LabelRenderer) CustomizeRender(r chart.Renderer, style chart.Style, values []LabelValue) {
	if !l.Options.Show {
		return
	}
	r.SetFont(style.Font)
	r.SetFontColor(style.FontColor)
	r.SetFontSize(style.FontSize)
	offsetX := l.Options.Offset.Left + l.Offset.Left
	offsetY := l.Options.Offset.Top + l.Offset.Top
	for _, value := range values {
		x := value.Left + offsetX
		y := value.Top + offsetY
		text := value.Text
		x -= r.MeasureText(text).Width() >> 1
		r.Text(text, x, y)
	}
}
