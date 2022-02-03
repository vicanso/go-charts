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

type YAxisOption struct {
	Min    *float64
	Max    *float64
	Hidden bool
}

const YAxisWidth = 40

func drawYAxis(p *Draw, opt *ChartOption, xAxisHeight int, padding chart.Box) (*Range, error) {
	theme := NewTheme(opt.Theme)
	yRange := opt.getYRange(0)
	data := NewAxisDataListFromStringList(yRange.Values())
	style := AxisStyle{
		Position:       PositionLeft,
		BoundaryGap:    FalseFlag(),
		FontColor:      theme.GetAxisStrokeColor(),
		TickShow:       FalseFlag(),
		StrokeWidth:    1,
		SplitLineColor: theme.GetAxisSplitLineColor(),
		SplitLineShow:  true,
	}
	width := NewAxis(p, data, style).measureAxis()

	padding.Left += (YAxisWidth - width)

	dYAxis, err := NewDraw(
		DrawOption{
			Parent: p,
			Width:  p.Box.Width(),
			// 减去x轴的高
			Height: p.Box.Height() - xAxisHeight,
		},
		PaddingOption(padding),
	)
	if err != nil {
		return nil, err
	}
	NewAxis(dYAxis, data, style).Render()
	yRange.Size = dYAxis.Box.Height()
	return &yRange, nil
}
