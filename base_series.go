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
	"fmt"

	"github.com/wcharczuk/go-chart/v2"
)

// Interface Assertions.
var (
	_ chart.Series              = (*BaseSeries)(nil)
	_ chart.FirstValuesProvider = (*BaseSeries)(nil)
	_ chart.LastValuesProvider  = (*BaseSeries)(nil)
)

// BaseSeries represents a line on a chart.
type BaseSeries struct {
	Name         string
	Style        chart.Style
	TickPosition chart.TickPosition

	YAxis chart.YAxisType

	XValueFormatter chart.ValueFormatter
	YValueFormatter chart.ValueFormatter

	XValues []float64
	YValues []float64
}

// GetName returns the name of the time series.
func (bs BaseSeries) GetName() string {
	return bs.Name
}

// GetStyle returns the line style.
func (bs BaseSeries) GetStyle() chart.Style {
	return bs.Style
}

// Len returns the number of elements in the series.
func (bs BaseSeries) Len() int {
	offset := 0
	if bs.TickPosition == chart.TickPositionBetweenTicks {
		offset = -1
	}
	return len(bs.XValues) + offset
}

// GetValues gets the x,y values at a given index.
func (bs BaseSeries) GetValues(index int) (float64, float64) {
	if bs.TickPosition == chart.TickPositionBetweenTicks {
		index++
	}
	return bs.XValues[index], bs.YValues[index]
}

// GetFirstValues gets the first x,y values.
func (bs BaseSeries) GetFirstValues() (float64, float64) {
	index := 0
	if bs.TickPosition == chart.TickPositionBetweenTicks {
		index++
	}
	return bs.XValues[index], bs.YValues[index]
}

// GetLastValues gets the last x,y values.
func (bs BaseSeries) GetLastValues() (float64, float64) {
	return bs.XValues[len(bs.XValues)-1], bs.YValues[len(bs.YValues)-1]
}

// GetValueFormatters returns value formatter defaults for the series.
func (bs BaseSeries) GetValueFormatters() (x, y chart.ValueFormatter) {
	if bs.XValueFormatter != nil {
		x = bs.XValueFormatter
	} else {
		x = chart.FloatValueFormatter
	}
	if bs.YValueFormatter != nil {
		y = bs.YValueFormatter
	} else {
		y = chart.FloatValueFormatter
	}
	return
}

// GetYAxis returns which YAxis the series draws on.
func (bs BaseSeries) GetYAxis() chart.YAxisType {
	return bs.YAxis
}

// Render renders the series.
func (bs BaseSeries) Render(r chart.Renderer, canvasBox chart.Box, xrange, yrange chart.Range, defaults chart.Style) {
	fmt.Println("should be override the function")
}

// Validate validates the series.
func (bs BaseSeries) Validate() error {
	if len(bs.XValues) == 0 {
		return fmt.Errorf("continuous series; must have xvalues set")
	}

	if len(bs.YValues) == 0 {
		return fmt.Errorf("continuous series; must have yvalues set")
	}

	if len(bs.XValues) != len(bs.YValues) {
		return fmt.Errorf("continuous series; must have same length xvalues as yvalues")
	}
	return nil
}
