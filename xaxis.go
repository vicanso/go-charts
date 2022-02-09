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

type XAxisOption struct {
	// The boundary gap on both sides of a coordinate axis.
	// Nil or *true means the center part of two axis ticks
	BoundaryGap *bool
	// The data value of x axis
	Data []string
	// The theme of chart
	Theme string
	// Hidden x axis
	Hidden bool
	// Number of segments that the axis is split into. Note that this number serves only as a recommendation.
	SplitNumber int
}

// drawXAxis draws x axis, and returns the height, range of if.
func drawXAxis(p *Draw, opt *XAxisOption, yAxisCount int) (int, *Range, error) {
	if opt.Hidden {
		return 0, nil, nil
	}
	left := YAxisWidth
	right := (yAxisCount - 1) * YAxisWidth
	dXAxis, err := NewDraw(
		DrawOption{
			Parent: p,
		},
		PaddingOption(chart.Box{
			Left:  left,
			Right: right,
		}),
	)
	if err != nil {
		return 0, nil, err
	}
	theme := NewTheme(opt.Theme)
	data := NewAxisDataListFromStringList(opt.Data)
	style := AxisOption{
		BoundaryGap: opt.BoundaryGap,
		StrokeColor: theme.GetAxisStrokeColor(),
		FontColor:   theme.GetAxisStrokeColor(),
		StrokeWidth: 1,
		SplitNumber: opt.SplitNumber,
	}

	boundary := true
	max := float64(len(opt.Data))
	if isFalse(opt.BoundaryGap) {
		boundary = false
		max--
	}
	axis := NewAxis(dXAxis, data, style)
	axis.Render()

	return axis.measureAxis(), &Range{
		divideCount: len(opt.Data),
		Min:         0,
		Max:         max,
		Size:        dXAxis.Box.Width(),
		Boundary:    boundary,
	}, nil
}
