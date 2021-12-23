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
	"github.com/wcharczuk/go-chart/v2"
)

type SeriesData struct {
	Value float64
	Style chart.Style
}

type Series struct {
	Type       string
	Name       string
	Data       []SeriesData
	XValues    []float64
	YAxisIndex int
}

const lineStrokeWidth = 2
const dotWith = 2

const (
	SeriesBar  = "bar"
	SeriesLine = "line"
	SeriesPie  = "pie"
)

func NewSeriesDataListFromFloat(values []float64) []SeriesData {
	dataList := make([]SeriesData, len(values))
	for index, value := range values {
		dataList[index] = SeriesData{
			Value: value,
		}
	}
	return dataList
}
func GetSeries(series []Series, tickPosition chart.TickPosition, theme string) []chart.Series {
	arr := make([]chart.Series, len(series))
	barCount := 0
	barIndex := 0
	for _, item := range series {
		if item.Type == SeriesBar {
			barCount++
		}
	}

	for index, item := range series {
		style := chart.Style{
			StrokeWidth: lineStrokeWidth,
			StrokeColor: getSeriesColor(theme, index),
			// TODO 调整为通过dot with color 生成
			DotColor: getSeriesColor(theme, index),
			DotWidth: dotWith,
		}
		pointIndexOffset := 0
		// 如果居中，需要多增加一个点
		if tickPosition == chart.TickPositionBetweenTicks {
			item.Data = append([]SeriesData{
				{
					Value: 0.0,
				},
			}, item.Data...)
			pointIndexOffset = -1
		}
		yValues := make([]float64, len(item.Data))
		barCustomStyles := make([]BarSeriesCustomStyle, 0)
		for i, item := range item.Data {
			yValues[i] = item.Value
			if !item.Style.IsZero() {
				barCustomStyles = append(barCustomStyles, BarSeriesCustomStyle{
					PointIndex: i + pointIndexOffset,
					Index:      barIndex,
					Style:      item.Style,
				})
			}
		}
		baseSeries := BaseSeries{
			Name:         item.Name,
			XValues:      item.XValues,
			Style:        style,
			YValues:      yValues,
			TickPosition: tickPosition,
			YAxis:        chart.YAxisSecondary,
		}
		if item.YAxisIndex != 0 {
			baseSeries.YAxis = chart.YAxisPrimary
		}
		switch item.Type {
		case SeriesBar:
			arr[index] = BarSeries{
				Count:        barCount,
				Index:        barIndex,
				BaseSeries:   baseSeries,
				CustomStyles: barCustomStyles,
			}
			barIndex++
		default:
			arr[index] = LineSeries{
				BaseSeries: baseSeries,
			}
		}
	}
	return arr
}
