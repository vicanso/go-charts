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
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type SeriesData struct {
	Value float64
	Style chart.Style
}

func NewSeriesFromValues(values []float64, chartType ...string) Series {
	s := Series{
		Data: NewSeriesDataFromValues(values),
	}
	if len(chartType) != 0 {
		s.Type = chartType[0]
	}
	return s
}

func NewSeriesDataFromValues(values []float64) []SeriesData {
	data := make([]SeriesData, len(values))
	for index, value := range values {
		data[index] = SeriesData{
			Value: value,
		}
	}
	return data
}

type SeriesLabel struct {
	Formatter string
	Color     drawing.Color
	Show      bool
}

const (
	SeriesMarkPointDataTypeMax = "max"
	SeriesMarkPointDataTypeMin = "min"
)

type SeriesMarkPointData struct {
	Type string
}
type SeriesMarkPoint struct {
	SymbolSize int
	Data       []SeriesMarkPointData
}
type Series struct {
	index      int
	Type       string
	Data       []SeriesData
	YAxisIndex int
	Style      chart.Style
	Label      SeriesLabel
	Name       string
	// Radius of Pie chart, e.g.: 40%
	Radius    string
	MarkPoint SeriesMarkPoint
}

type seriesSummary struct {
	MaxIndex     int
	MaxValue     float64
	MinIndex     int
	MinValue     float64
	AverageValue float64
}

func (s *Series) Summary() seriesSummary {
	minIndex := -1
	maxIndex := -1
	minValue := math.MaxFloat64
	maxValue := -math.MaxFloat64
	sum := float64(0)
	for j, item := range s.Data {
		if item.Value < minValue {
			minIndex = j
			minValue = item.Value
		}
		if item.Value > maxValue {
			maxIndex = j
			maxValue = item.Value
		}
		sum += item.Value
	}
	return seriesSummary{
		MaxIndex:     maxIndex,
		MaxValue:     maxValue,
		MinIndex:     minIndex,
		MinValue:     minValue,
		AverageValue: sum / float64(len(s.Data)),
	}
}

type LabelFormatter func(index int, value float64, percent float64) string

func NewPieLabelFormatter(seriesNames []string, layout string) LabelFormatter {
	if len(layout) == 0 {
		layout = "{b}: {d}"
	}
	return NewLabelFormatter(seriesNames, layout)
}

func NewValueLabelFormater(seriesNames []string, layout string) LabelFormatter {
	if len(layout) == 0 {
		layout = "{c}"
	}
	return NewLabelFormatter(seriesNames, layout)
}

func NewLabelFormatter(seriesNames []string, layout string) LabelFormatter {
	return func(index int, value, percent float64) string {
		// 如果无percent的则设置为<0
		percentText := ""
		if percent >= 0 {
			percentText = humanize.FtoaWithDigits(percent*100, 2) + "%"
		}
		valueText := humanize.FtoaWithDigits(value, 2)
		name := ""
		if len(seriesNames) > index {
			name = seriesNames[index]
		}
		text := strings.ReplaceAll(layout, "{c}", valueText)
		text = strings.ReplaceAll(text, "{d}", percentText)
		text = strings.ReplaceAll(text, "{b}", name)
		return text
	}
}
