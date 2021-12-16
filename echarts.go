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
	"bytes"
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/wcharczuk/go-chart/v2"
)

type EChartStyle struct {
	Color string `json:"color"`
}
type ECharsSeriesData struct {
	Value     float64     `json:"value"`
	ItemStyle EChartStyle `json:"itemStyle"`
}
type _ECharsSeriesData ECharsSeriesData

func (es *ECharsSeriesData) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return nil
	}
	if regexp.MustCompile(`^\d+`).Match(data) {
		v, err := strconv.ParseFloat(string(data), 64)
		if err != nil {
			return err
		}
		es.Value = v
		return nil
	}
	v := _ECharsSeriesData{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	es.Value = v.Value
	es.ItemStyle = v.ItemStyle
	return nil
}

type EChartsYAxis struct {
	Data []struct {
		Min       int `json:"min"`
		Max       int `json:"max"`
		Interval  int `json:"interval"`
		AxisLabel struct {
			Formatter string `json:"formatter"`
		} `json:"axisLabel"`
	} `json:"data"`
}

func (ey *EChartsYAxis) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return nil
	}
	if data[0] != '[' {
		data = []byte("[" + string(data) + "]")
	}
	return json.Unmarshal(data, &ey.Data)
}

type ECharsOptions struct {
	Title struct {
		Text      string `json:"text"`
		TextStyle struct {
			Color      string `json:"color"`
			FontFamily string `json:"fontFamily"`
		} `json:"textStyle"`
	} `json:"title"`
	XAxis struct {
		Type        string   `json:"type"`
		BoundaryGap *bool    `json:"boundaryGap"`
		SplitNumber int      `json:"splitNumber"`
		Data        []string `json:"data"`
	} `json:"xAxis"`
	YAxis  EChartsYAxis `json:"yAxis"`
	Legend struct {
		Data []string `json:"data"`
	} `json:"legend"`
	Series []struct {
		Data []ECharsSeriesData `json:"data"`
		Type string             `json:"type"`
	} `json:"series"`
}

func (e *ECharsOptions) ToOptions() Options {
	o := Options{}
	o.Title = Title{
		Text: e.Title.Text,
	}

	o.XAxis = XAxis{
		Type:        e.XAxis.Type,
		Data:        e.XAxis.Data,
		SplitNumber: e.XAxis.SplitNumber,
	}

	o.Legend = Legend{
		Data: e.Legend.Data,
	}

	// TODO 生成yAxis

	tickPosition := chart.TickPositionUnset
	series := make([]Series, len(e.Series))
	for index, item := range e.Series {
		// bar默认tick居中
		if item.Type == SeriesBar {
			tickPosition = chart.TickPositionBetweenTicks
		}
		data := make([]SeriesData, len(item.Data))
		for j, itemData := range item.Data {
			sd := SeriesData{
				Value: itemData.Value,
			}
			if itemData.ItemStyle.Color != "" {
				c := parseColor(itemData.ItemStyle.Color)
				sd.Style.FillColor = c
				sd.Style.StrokeColor = c
			}
			data[j] = sd
		}
		series[index] = Series{
			Data: data,
			Type: item.Type,
		}
	}
	o.Series = series

	if e.XAxis.BoundaryGap == nil || *e.XAxis.BoundaryGap {
		tickPosition = chart.TickPositionBetweenTicks
	}
	o.TickPosition = tickPosition
	return o
}

func ParseECharsOptions(options string) (Options, error) {
	e := ECharsOptions{}
	err := json.Unmarshal([]byte(options), &e)
	if err != nil {
		return Options{}, err
	}

	return e.ToOptions(), nil
}
