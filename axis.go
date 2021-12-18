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
	"github.com/dustin/go-humanize"
	"github.com/wcharczuk/go-chart/v2"
)

type (
	// AxisData string
	XAxis struct {
		// 'value' 数值轴，适用于连续数据。
		// 'category' 类目轴，适用于离散的类目数据。为该类型时类目数据可自动从 series.data 或 dataset.source 中取，或者可通过 xAxis.data 设置类目数据。
		// 'time' 时间轴，适用于连续的时序数据，与数值轴相比时间轴带有时间的格式化，在刻度计算上也有所不同，例如会根据跨度的范围来决定使用月，星期，日还是小时范围的刻度。
		// 'log' 对数轴。适用于对数数据。
		Type        string
		Data        []string
		SplitNumber int
	}
)

type YAxisOption struct {
	Formater chart.ValueFormatter
	Disabled bool
	Min      *float64
	Max      *float64
}

const axisStrokeWidth = 1

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

	maxTicks := xAxis.SplitNumber
	if maxTicks == 0 {
		maxTicks = 10
	}

	// 计息最多每个unit至少放多个
	minUnitSize := originalSize / maxTicks
	if originalSize%maxTicks != 0 {
		minUnitSize++
	}
	unitSize := minUnitSize
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
	if value >= 10 {
		return humanize.CommafWithDigits(value, 0)
	}
	return humanize.CommafWithDigits(value, 2)
}

func GetSecondaryYAxis(theme string, option *YAxisOption) chart.YAxis {
	strokeColor := getGridColor(theme)
	yAxis := chart.YAxis{
		Range:          &chart.ContinuousRange{},
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
	setYAxis(&yAxis, option)
	return yAxis
}

func setYAxis(yAxis *chart.YAxis, option *YAxisOption) {
	if option == nil {
		return
	}
	if option.Formater != nil {
		yAxis.ValueFormatter = option.Formater
	}
	if option.Max != nil {
		yAxis.Range.SetMax(*option.Max)
	}
	if option.Min != nil {
		yAxis.Range.SetMin(*option.Min)
	}
}

func GetYAxis(theme string, option *YAxisOption) chart.YAxis {
	disabled := false
	if option != nil {
		disabled = option.Disabled
	}
	hidden := chart.Hidden()

	yAxis := chart.YAxis{
		Range:          &chart.ContinuousRange{},
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
	if disabled {
		yAxis.Range = &HiddenRange{}
		yAxis.Style.Hidden = true
	}
	setYAxis(&yAxis, option)
	return yAxis
}
