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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
)

func TestBaseSeries(t *testing.T) {
	assert := assert.New(t)

	bs := BaseSeries{
		XValues: []float64{
			1,
			2,
			3,
		},
		YValues: []float64{
			10,
			20,
			30,
		},
	}
	assert.Equal(3, bs.Len())
	bs.TickPosition = chart.TickPositionBetweenTicks
	assert.Equal(2, bs.Len())

	bs.TickPosition = chart.TickPositionUnset
	x, y := bs.GetValues(1)
	assert.Equal(float64(2), x)
	assert.Equal(float64(20), y)
	bs.TickPosition = chart.TickPositionBetweenTicks
	x, y = bs.GetValues(1)
	assert.Equal(float64(3), x)
	assert.Equal(float64(30), y)

	bs.TickPosition = chart.TickPositionUnset
	x, y = bs.GetFirstValues()
	assert.Equal(float64(1), x)
	assert.Equal(float64(10), y)
	bs.TickPosition = chart.TickPositionBetweenTicks
	x, y = bs.GetFirstValues()
	assert.Equal(float64(2), x)
	assert.Equal(float64(20), y)

	bs.TickPosition = chart.TickPositionUnset
	x, y = bs.GetLastValues()
	assert.Equal(float64(3), x)
	assert.Equal(float64(30), y)
	bs.TickPosition = chart.TickPositionBetweenTicks
	x, y = bs.GetLastValues()
	assert.Equal(float64(3), x)
	assert.Equal(float64(30), y)

	xFormater, yFormater := bs.GetValueFormatters()
	assert.Equal(reflect.ValueOf(chart.FloatValueFormatter).Pointer(), reflect.ValueOf(xFormater).Pointer())
	assert.Equal(reflect.ValueOf(chart.FloatValueFormatter).Pointer(), reflect.ValueOf(yFormater).Pointer())
	formater := func(v interface{}) string {
		return ""
	}
	bs.XValueFormatter = formater
	bs.YValueFormatter = formater
	xFormater, yFormater = bs.GetValueFormatters()
	assert.Equal(reflect.ValueOf(formater).Pointer(), reflect.ValueOf(xFormater).Pointer())
	assert.Equal(reflect.ValueOf(formater).Pointer(), reflect.ValueOf(yFormater).Pointer())

	assert.Equal(chart.YAxisPrimary, bs.GetYAxis())

	assert.Nil(bs.Validate())
}
