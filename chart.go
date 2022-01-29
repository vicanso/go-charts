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
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type XAxisOption struct {
	BoundaryGap *bool
	Data        []string
	// TODO split number
}

type SeriesData struct {
	Value float64
	Style chart.Style
}
type Point struct {
	X int
	Y int
}

type Range struct {
	originalMin float64
	originalMax float64
	divideCount int
	Min         float64
	Max         float64
	Size        int
	Boundary    bool
}

func (r *Range) getHeight(value float64) int {
	v := 1 - value/(r.Max-r.Min)
	return int(v * float64(r.Size))
}

func (r *Range) getWidth(value float64) int {
	v := value / (r.Max - r.Min)
	// 移至居中
	if r.Boundary &&
		r.divideCount != 0 {
		v += 1 / float64(r.divideCount*2)
	}
	return int(v * float64(r.Size))
}

type Series struct {
	Type       string
	Name       string
	Data       []SeriesData
	YAxisIndex int
	Style      chart.Style
}

type ChartOption struct {
	Theme           string
	XAxis           XAxisOption
	Width           int
	Height          int
	Parent          *Draw
	Padding         chart.Box
	SeriesList      []Series
	BackgroundColor drawing.Color
}

func (o *ChartOption) getWidth() int {
	if o.Width == 0 {
		return 600
	}
	return o.Width
}

func (o *ChartOption) getHeight() int {
	if o.Height == 0 {
		return 400
	}
	return o.Height
}

func (o *ChartOption) getYRange(axisIndex int) Range {
	min := float64(0)
	max := float64(0)

	for _, series := range o.SeriesList {
		if series.YAxisIndex != axisIndex {
			continue
		}
		for _, item := range series.Data {
			if item.Value > max {
				max = item.Value
			}
			if item.Value < min {
				min = item.Value
			}
		}
	}
	// TODO 对于小数的处理

	divideCount := 6
	r := Range{
		originalMin: min,
		originalMax: max,
		Min:         float64(int(min * 0.8)),
		Max:         max * 1.2,
		divideCount: divideCount,
	}
	value := int((r.Max - r.Min) / float64(divideCount))
	r.Max = float64(int(float64(value*divideCount) + r.Min))
	return r
}

func (r Range) Values() []string {
	offset := (r.Max - r.Min) / float64(r.divideCount)
	values := make([]string, 0)
	for i := 0; i <= r.divideCount; i++ {
		v := r.Min + float64(i)*offset
		value := humanize.CommafWithDigits(v, 2)
		values = append(values, value)
	}
	return values
}
