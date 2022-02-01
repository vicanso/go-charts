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

import "github.com/wcharczuk/go-chart/v2"

type XAxisOption struct {
	BoundaryGap *bool
	Data        []string
	// TODO split number
}

// drawXAxis draws x axis, and returns the height, range of if.
func drawXAxis(p *Draw, opt *XAxisOption, theme *Theme) (int, *Range, error) {
	dXAxis, err := NewDraw(
		DrawOption{
			Parent: p,
		},
		PaddingOption(chart.Box{
			Left: YAxisWidth,
		}),
	)
	if err != nil {
		return 0, nil, err
	}
	data := NewAxisDataListFromStringList(opt.Data)
	style := AxisStyle{
		BoundaryGap: opt.BoundaryGap,
		StrokeColor: theme.GetAxisStrokeColor(),
		FontColor:   theme.GetAxisStrokeColor(),
		StrokeWidth: 1,
	}

	boundary := true
	max := float64(len(opt.Data))
	if opt.BoundaryGap != nil && !*opt.BoundaryGap {
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
