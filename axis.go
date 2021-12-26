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
	"math"

	"github.com/dustin/go-humanize"
	"github.com/wcharczuk/go-chart/v2"
)

type (
	// AxisData string
	XAxis struct {
		// data value of axis
		Data []string
		// number of segments
		SplitNumber int
	}
)

type YAxisOption struct {
	// formater of axis
	Formater chart.ValueFormatter
	// disabled axis
	Disabled bool
	// min value of axis
	Min *float64
	// max value of axis
	Max *float64
}

const axisStrokeWidth = 1

func maxInt(values ...int) int {
	result := 0
	for _, v := range values {
		if v > result {
			result = v
		}
	}
	return result
}

// GetXAxisAndValues returns x axis by theme, and the values of axis.
func GetXAxisAndValues(xAxis XAxis, tickPosition chart.TickPosition, theme string) (chart.XAxis, []float64) {
	data := xAxis.Data
	originalSize := len(data)
	// 如果居中，则需要多添加一个值
	if tickPosition == chart.TickPositionBetweenTicks {
		data = append([]string{
			"",
		}, data...)
	}

	size := len(data)

	xValues := make([]float64, size)
	ticks := make([]chart.Tick, 0)

	// tick width
	maxTicks := maxInt(xAxis.SplitNumber, 10)

	// 计息最多每个unit至少放多个
	minUnitSize := originalSize / maxTicks
	if originalSize%maxTicks != 0 {
		minUnitSize++
	}
	unitSize := minUnitSize
	// 尽可能选择一格展示更多的块
	for i := minUnitSize; i < 2*minUnitSize; i++ {
		if originalSize%i == 0 {
			unitSize = i
		}
	}

	for index, key := range data {
		f := float64(index)
		xValues[index] = f
		if index%unitSize == 0 || index == size-1 {
			ticks = append(ticks, chart.Tick{
				Value: f,
				Label: key,
			})
		}
	}
	return chart.XAxis{
		Ticks:        ticks,
		TickPosition: tickPosition,
		Style: chart.Style{
			FontColor:   getAxisColor(theme),
			StrokeColor: getAxisColor(theme),
			StrokeWidth: axisStrokeWidth,
		},
	}, xValues
}

func defaultFloatFormater(v interface{}) string {
	value, ok := v.(float64)
	if !ok {
		return ""
	}
	// 大于10的则直接取整展示
	if value >= 10 {
		return humanize.CommafWithDigits(value, 0)
	}
	return humanize.CommafWithDigits(value, 2)
}

func newYContinuousRange(option *YAxisOption) *YContinuousRange {
	m := YContinuousRange{}
	m.Min = -math.MaxFloat64
	m.Max = math.MaxFloat64
	if option != nil {
		if option.Min != nil {
			m.Min = *option.Min
		}
		if option.Max != nil {
			m.Max = *option.Max
		}
	}
	return &m
}

// GetSecondaryYAxis returns the secondary y axis by theme
func GetSecondaryYAxis(theme string, option *YAxisOption) chart.YAxis {
	strokeColor := getGridColor(theme)
	yAxis := chart.YAxis{
		Range:          newYContinuousRange(option),
		ValueFormatter: defaultFloatFormater,
		AxisType:       chart.YAxisSecondary,
		GridMajorStyle: chart.Style{
			StrokeColor: strokeColor,
			StrokeWidth: axisStrokeWidth,
		},
		GridMinorStyle: chart.Style{
			StrokeColor: strokeColor,
			StrokeWidth: axisStrokeWidth,
		},
		Style: chart.Style{
			FontColor: getAxisColor(theme),
			// alpha 0，隐藏
			StrokeColor: hiddenColor,
			StrokeWidth: axisStrokeWidth,
		},
	}
	setYAxisOption(&yAxis, option)
	return yAxis
}

func setYAxisOption(yAxis *chart.YAxis, option *YAxisOption) {
	if option == nil {
		return
	}
	if option.Formater != nil {
		yAxis.ValueFormatter = option.Formater
	}
}

// GetYAxis returns the primary y axis by theme
func GetYAxis(theme string, option *YAxisOption) chart.YAxis {
	disabled := false
	if option != nil {
		disabled = option.Disabled
	}
	hidden := chart.Hidden()

	yAxis := chart.YAxis{
		Range:          newYContinuousRange(option),
		ValueFormatter: defaultFloatFormater,
		AxisType:       chart.YAxisPrimary,
		GridMajorStyle: hidden,
		GridMinorStyle: hidden,
		Style: chart.Style{
			FontColor: getAxisColor(theme),
			// alpha 0，隐藏
			StrokeColor: hiddenColor,
			StrokeWidth: axisStrokeWidth,
		},
	}
	// 如果禁用，则默认为隐藏，并设置range
	if disabled {
		yAxis.Range = &HiddenRange{}
		yAxis.Style.Hidden = true
	}
	setYAxisOption(&yAxis, option)
	return yAxis
}
