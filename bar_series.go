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

import "github.com/wcharczuk/go-chart/v2"

type BarSeries struct {
	BaseSeries
}

func (bs BarSeries) Render(r chart.Renderer, canvasBox chart.Box, xrange, yrange chart.Range, defaults chart.Style) {
	if bs.Len() == 0 {
		return
	}
	style := bs.Style.InheritFrom(defaults)
	style.FillColor = style.StrokeColor
	if !style.ShouldDrawStroke() {
		return
	}

	cb := canvasBox.Bottom
	cl := canvasBox.Left

	for i := 0; i < bs.Len(); i++ {
		vx, vy := bs.GetValues(i)

		x := cl + xrange.Translate(vx)
		y := cb - yrange.Translate(vy)

		chart.Draw.Box(r, chart.Box{
			Left: x,
			Top:  y,
			// TODO 计算宽度
			Right:  x + 10,
			Bottom: canvasBox.Bottom - 1,
		}, style)
	}

}
