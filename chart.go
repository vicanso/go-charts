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
	"math"

	"github.com/dustin/go-humanize"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type SeriesData struct {
	Value float64
	Style chart.Style
}
type Point struct {
	X int
	Y int
}

type Series struct {
	Type       string
	Data       []SeriesData
	YAxisIndex int
	Style      chart.Style
}

type ChartOption struct {
	Theme           string
	Title           TitleOption
	Legend          LegendOption
	XAxis           XAxisOption
	Width           int
	Height          int
	Parent          *Draw
	Padding         chart.Box
	SeriesList      []Series
	BackgroundColor drawing.Color
}

func (o *ChartOption) FillDefault(t *Theme) {
	f, _ := chart.GetDefaultFont()
	if o.BackgroundColor.IsZero() {
		o.BackgroundColor = t.GetBackgroundColor()
	}
	if o.Title.Style.FontColor.IsZero() {
		o.Title.Style.FontColor = t.GetTextColor()
	}
	if o.Title.Style.FontSize == 0 {
		o.Title.Style.FontSize = 14
	}
	if o.Title.Style.Font == nil {
		o.Title.Style.Font = f
	}
	if o.Title.Style.Padding.IsZero() {
		o.Title.Style.Padding = chart.Box{
			Left:   5,
			Top:    5,
			Right:  5,
			Bottom: 5,
		}
	}
	if o.Legend.Style.FontSize == 0 {
		o.Legend.Style.FontSize = 8
	}
	if o.Legend.Style.Font == nil {
		o.Legend.Style.Font = f
	}
	if o.Legend.Style.FontColor.IsZero() {
		o.Legend.Style.FontColor = t.GetTextColor()
	}
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
	min := math.MaxFloat64
	max := -math.MaxFloat64

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
	// y轴分设置默认划分为6块
	r := NewRange(min*0.9, max*1.1, 6)
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
