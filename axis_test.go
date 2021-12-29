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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
)

func TestGetXAxisAndValues(t *testing.T) {
	assert := assert.New(t)

	genLabels := func(count int) []string {
		arr := make([]string, count)
		for i := 0; i < count; i++ {
			arr[i] = strconv.Itoa(i)
		}
		return arr
	}
	genValues := func(count int, betweenTicks bool) []float64 {
		if betweenTicks {
			count++
		}
		arr := make([]float64, count)
		for i := 0; i < count; i++ {
			arr[i] = float64(i)
		}
		return arr
	}
	genTicks := func(count int, betweenTicks bool) []chart.Tick {
		arr := make([]chart.Tick, 0)
		offset := 0
		if betweenTicks {
			offset = 1
			arr = append(arr, chart.Tick{})
		}
		for i := 0; i < count; i++ {
			arr = append(arr, chart.Tick{
				Value: float64(i + offset),
				Label: strconv.Itoa(i),
			})
		}
		return arr
	}

	tests := []struct {
		xAxis        XAxis
		tickPosition chart.TickPosition
		theme        string
		result       chart.XAxis
		values       []float64
	}{
		{
			xAxis: XAxis{
				Data: genLabels(5),
			},
			values: genValues(5, false),
			result: chart.XAxis{
				Ticks: genTicks(5, false),
			},
		},
		// 居中
		{
			xAxis: XAxis{
				Data: genLabels(5),
			},
			tickPosition: chart.TickPositionBetweenTicks,
			// 居中因此value多一个
			values: genValues(5, true),
			result: chart.XAxis{
				Ticks: genTicks(5, true),
			},
		},
		{
			xAxis: XAxis{
				Data: genLabels(20),
			},
			// 居中因此value多一个
			values: genValues(20, false),
			result: chart.XAxis{
				Ticks: []chart.Tick{
					{Value: 0, Label: "0"}, {Value: 2, Label: "2"}, {Value: 4, Label: "4"}, {Value: 6, Label: "6"}, {Value: 8, Label: "8"}, {Value: 10, Label: "10"}, {Value: 12, Label: "12"}, {Value: 14, Label: "14"}, {Value: 16, Label: "16"}, {Value: 18, Label: "18"}, {Value: 19, Label: "19"}},
			},
		},
	}

	for _, tt := range tests {
		xAxis, values := GetXAxisAndValues(tt.xAxis, tt.tickPosition, tt.theme)

		assert.Equal(tt.result.Ticks, xAxis.Ticks)
		assert.Equal(tt.tickPosition, xAxis.TickPosition)
		assert.Equal(tt.values, values)
	}
}

func TestDefaultFloatFormater(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("", defaultFloatFormater(1))

	assert.Equal("0.1", defaultFloatFormater(0.1))
	assert.Equal("0.12", defaultFloatFormater(0.123))
	assert.Equal("10", defaultFloatFormater(10.1))
}

func TestSetYAxisOption(t *testing.T) {
	assert := assert.New(t)
	min := 10.0
	max := 20.0
	opt := &YAxisOption{
		Formater: func(v interface{}) string {
			return ""
		},
		Min: &min,
		Max: &max,
	}
	yAxis := &chart.YAxis{
		Range: newYContinuousRange(opt),
	}
	setYAxisOption(yAxis, opt)

	assert.NotEmpty(yAxis.ValueFormatter)
	assert.Equal(max, yAxis.Range.GetMax())
	assert.Equal(min, yAxis.Range.GetMin())
}

func TestGetYAxis(t *testing.T) {
	assert := assert.New(t)

	yAxis := GetYAxis(ThemeDark, nil)

	assert.True(yAxis.GridMajorStyle.Hidden)
	assert.True(yAxis.GridMajorStyle.Hidden)
	assert.False(yAxis.Style.Hidden)

	yAxis = GetYAxis(ThemeDark, &YAxisOption{
		Disabled: true,
	})

	assert.True(yAxis.GridMajorStyle.Hidden)
	assert.True(yAxis.GridMajorStyle.Hidden)
	assert.True(yAxis.Style.Hidden)

	// secondary yAxis
	yAxis = GetSecondaryYAxis(ThemeDark, nil)
	assert.False(yAxis.GridMajorStyle.Hidden)
	assert.False(yAxis.GridMajorStyle.Hidden)
	assert.True(yAxis.Style.StrokeColor.IsTransparent())
}
