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

type LabelValue struct {
	Text  string
	Style Style
	X     int
	Y     int
}

type SeriesLabelPainter struct {
	p      *Painter
	values []LabelValue
}

func NewSeriesLabelPainter(p *Painter) *SeriesLabelPainter {
	return &SeriesLabelPainter{
		p:      p,
		values: make([]LabelValue, 0),
	}
}

func (o *SeriesLabelPainter) Add(value LabelValue) {
	o.values = append(o.values, value)
}

func (o *SeriesLabelPainter) Render() (Box, error) {
	for _, item := range o.values {
		o.p.OverrideTextStyle(item.Style)
		o.p.Text(item.Text, item.X, item.Y)
	}
	return chart.BoxZero, nil
}
