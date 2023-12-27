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
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/wcharczuk/go-chart/v2"
)

func convertToArray(data []byte) []byte {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return nil
	}
	if data[0] != '[' {
		data = []byte("[" + string(data) + "]")
	}
	return data
}

type EChartsPosition string

func (p *EChartsPosition) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if regexp.MustCompile(`^\d+`).Match(data) {
		data = []byte(fmt.Sprintf(`"%s"`, string(data)))
	}
	s := (*string)(p)
	return json.Unmarshal(data, s)
}

type EChartStyle struct {
	Color string `json:"color"`
}

func (es *EChartStyle) ToStyle() Style {
	color := parseColor(es.Color)
	return Style{
		FillColor:   color,
		FontColor:   color,
		StrokeColor: color,
	}
}

type EChartsSeriesDataValue struct {
	values []float64
}

func (value *EChartsSeriesDataValue) UnmarshalJSON(data []byte) error {
	data = convertToArray(data)
	return json.Unmarshal(data, &value.values)
}
func (value *EChartsSeriesDataValue) First() float64 {
	if len(value.values) == 0 {
		return 0
	}
	return value.values[0]
}
func NewEChartsSeriesDataValue(values ...float64) EChartsSeriesDataValue {
	return EChartsSeriesDataValue{
		values: values,
	}
}

type EChartsSeriesData struct {
	Value     EChartsSeriesDataValue `json:"value"`
	Name      string                 `json:"name"`
	ItemStyle EChartStyle            `json:"itemStyle"`
}
type _EChartsSeriesData EChartsSeriesData

var numericRep = regexp.MustCompile(`^[-+]?[0-9]+(?:\.[0-9]+)?$`)

func (es *EChartsSeriesData) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return nil
	}
	if numericRep.Match(data) {
		v, err := strconv.ParseFloat(string(data), 64)
		if err != nil {
			return err
		}
		es.Value = EChartsSeriesDataValue{
			values: []float64{
				v,
			},
		}
		return nil
	}
	v := _EChartsSeriesData{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	es.Name = v.Name
	es.Value = v.Value
	es.ItemStyle = v.ItemStyle
	return nil
}

type EChartsXAxisData struct {
	BoundaryGap *bool    `json:"boundaryGap"`
	SplitNumber int      `json:"splitNumber"`
	Data        []string `json:"data"`
	Type        string   `json:"type"`
}
type EChartsXAxis struct {
	Data []EChartsXAxisData
}

func (ex *EChartsXAxis) UnmarshalJSON(data []byte) error {
	data = convertToArray(data)
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, &ex.Data)
}

type EChartsAxisLabel struct {
	Formatter string `json:"formatter"`
}
type EChartsYAxisData struct {
	Min       *float64         `json:"min"`
	Max       *float64         `json:"max"`
	AxisLabel EChartsAxisLabel `json:"axisLabel"`
	AxisLine  struct {
		LineStyle struct {
			Color string `json:"color"`
		} `json:"lineStyle"`
	} `json:"axisLine"`
	Data []string `json:"data"`
}
type EChartsYAxis struct {
	Data []EChartsYAxisData `json:"data"`
}

func (ey *EChartsYAxis) UnmarshalJSON(data []byte) error {
	data = convertToArray(data)
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, &ey.Data)
}

type EChartsPadding struct {
	Box chart.Box
}

func (eb *EChartsPadding) UnmarshalJSON(data []byte) error {
	data = convertToArray(data)
	if len(data) == 0 {
		return nil
	}
	arr := make([]int, 0)
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	if len(arr) == 0 {
		return nil
	}
	switch len(arr) {
	case 1:
		eb.Box = chart.Box{
			Left:   arr[0],
			Top:    arr[0],
			Bottom: arr[0],
			Right:  arr[0],
		}
	case 2:
		eb.Box = chart.Box{
			Top:    arr[0],
			Bottom: arr[0],
			Left:   arr[1],
			Right:  arr[1],
		}
	default:
		result := make([]int, 4)
		copy(result, arr)
		if len(arr) == 3 {
			result[3] = result[1]
		}
		// 上右下左
		eb.Box = chart.Box{
			Top:    result[0],
			Right:  result[1],
			Bottom: result[2],
			Left:   result[3],
		}
	}
	return nil
}

type EChartsLabelOption struct {
	Show     bool   `json:"show"`
	Distance int    `json:"distance"`
	Color    string `json:"color"`
}
type EChartsLegend struct {
	Show      *bool            `json:"show"`
	Data      []string         `json:"data"`
	Align     string           `json:"align"`
	Orient    string           `json:"orient"`
	Padding   EChartsPadding   `json:"padding"`
	Left      EChartsPosition  `json:"left"`
	Top       EChartsPosition  `json:"top"`
	TextStyle EChartsTextStyle `json:"textStyle"`
}

type EChartsMarkData struct {
	Type string `json:"type"`
}
type _EChartsMarkData EChartsMarkData

func (emd *EChartsMarkData) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return nil
	}
	data = convertToArray(data)
	ds := make([]*_EChartsMarkData, 0)
	err := json.Unmarshal(data, &ds)
	if err != nil {
		return err
	}
	for _, d := range ds {
		if d.Type != "" {
			emd.Type = d.Type
		}
	}
	return nil
}

type EChartsMarkPoint struct {
	SymbolSize int               `json:"symbolSize"`
	Data       []EChartsMarkData `json:"data"`
}

func (emp *EChartsMarkPoint) ToSeriesMarkPoint() SeriesMarkPoint {
	sp := SeriesMarkPoint{
		SymbolSize: emp.SymbolSize,
	}
	if len(emp.Data) == 0 {
		return sp
	}
	data := make([]SeriesMarkData, len(emp.Data))
	for index, item := range emp.Data {
		data[index].Type = item.Type
	}
	sp.Data = data
	return sp
}

type EChartsMarkLine struct {
	Data []EChartsMarkData `json:"data"`
}

func (eml *EChartsMarkLine) ToSeriesMarkLine() SeriesMarkLine {
	sl := SeriesMarkLine{}
	if len(eml.Data) == 0 {
		return sl
	}
	data := make([]SeriesMarkData, len(eml.Data))
	for index, item := range eml.Data {
		data[index].Type = item.Type
	}
	sl.Data = data
	return sl
}

type EChartsSeries struct {
	Data       []EChartsSeriesData `json:"data"`
	Name       string              `json:"name"`
	Type       string              `json:"type"`
	Radius     string              `json:"radius"`
	YAxisIndex int                 `json:"yAxisIndex"`
	ItemStyle  EChartStyle         `json:"itemStyle"`
	// label的配置
	Label     EChartsLabelOption `json:"label"`
	MarkPoint EChartsMarkPoint   `json:"markPoint"`
	MarkLine  EChartsMarkLine    `json:"markLine"`
	Max       *float64           `json:"max"`
	Min       *float64           `json:"min"`
}
type EChartsSeriesList []EChartsSeries

func (esList EChartsSeriesList) ToSeriesList() SeriesList {
	seriesList := make(SeriesList, 0, len(esList))
	for _, item := range esList {
		// 如果是pie，则每个子荐生成一个series
		if item.Type == ChartTypePie {
			for _, dataItem := range item.Data {
				seriesList = append(seriesList, Series{
					Type: item.Type,
					Name: dataItem.Name,
					Label: SeriesLabel{
						Show: true,
					},
					Radius: item.Radius,
					Data: []SeriesData{
						{
							Value: dataItem.Value.First(),
						},
					},
				})
			}
			continue
		}
		// 如果是radar或funnel
		if item.Type == ChartTypeRadar ||
			item.Type == ChartTypeFunnel {
			for _, dataItem := range item.Data {
				seriesList = append(seriesList, Series{
					Name: dataItem.Name,
					Type: item.Type,
					Data: NewSeriesDataFromValues(dataItem.Value.values),
					Max:  item.Max,
					Min:  item.Min,
					Label: SeriesLabel{
						Color:    parseColor(item.Label.Color),
						Show:     item.Label.Show,
						Distance: item.Label.Distance,
					},
				})
			}
			continue
		}
		data := make([]SeriesData, len(item.Data))
		for j, dataItem := range item.Data {
			data[j] = SeriesData{
				Value: dataItem.Value.First(),
				Style: dataItem.ItemStyle.ToStyle(),
			}
		}
		seriesList = append(seriesList, Series{
			Type:      item.Type,
			Data:      data,
			AxisIndex: item.YAxisIndex,
			Style:     item.ItemStyle.ToStyle(),
			Label: SeriesLabel{
				Color:    parseColor(item.Label.Color),
				Show:     item.Label.Show,
				Distance: item.Label.Distance,
			},
			Name:      item.Name,
			MarkPoint: item.MarkPoint.ToSeriesMarkPoint(),
			MarkLine:  item.MarkLine.ToSeriesMarkLine(),
		})
	}
	return seriesList
}

type EChartsTextStyle struct {
	Color      string  `json:"color"`
	FontFamily string  `json:"fontFamily"`
	FontSize   float64 `json:"fontSize"`
}

func (et *EChartsTextStyle) ToStyle() chart.Style {
	s := chart.Style{
		FontSize:  et.FontSize,
		FontColor: parseColor(et.Color),
	}
	if et.FontFamily != "" {
		s.Font, _ = GetFont(et.FontFamily)
	}
	return s
}

type EChartsOption struct {
	Type       string         `json:"type"`
	Theme      string         `json:"theme"`
	FontFamily string         `json:"fontFamily"`
	Padding    EChartsPadding `json:"padding"`
	Box        chart.Box      `json:"box"`
	Width      int            `json:"width"`
	Height     int            `json:"height"`
	Title      struct {
		Text         string           `json:"text"`
		Subtext      string           `json:"subtext"`
		Left         EChartsPosition  `json:"left"`
		Top          EChartsPosition  `json:"top"`
		TextStyle    EChartsTextStyle `json:"textStyle"`
		SubtextStyle EChartsTextStyle `json:"subtextStyle"`
	} `json:"title"`
	XAxis  EChartsXAxis  `json:"xAxis"`
	YAxis  EChartsYAxis  `json:"yAxis"`
	Legend EChartsLegend `json:"legend"`
	Radar  struct {
		Indicator []RadarIndicator `json:"indicator"`
	} `json:"radar"`
	Series   EChartsSeriesList `json:"series"`
	Children []EChartsOption   `json:"children"`
}

func (eo *EChartsOption) ToOption() ChartOption {
	fontFamily := eo.FontFamily
	if len(fontFamily) == 0 {
		fontFamily = eo.Title.TextStyle.FontFamily
	}
	titleTextStyle := eo.Title.TextStyle.ToStyle()
	titleSubtextStyle := eo.Title.SubtextStyle.ToStyle()
	legendTextStyle := eo.Legend.TextStyle.ToStyle()
	o := ChartOption{
		Type:       eo.Type,
		FontFamily: fontFamily,
		Theme:      eo.Theme,
		Title: TitleOption{
			Text:             eo.Title.Text,
			Subtext:          eo.Title.Subtext,
			FontColor:        titleTextStyle.FontColor,
			FontSize:         titleTextStyle.FontSize,
			SubtextFontSize:  titleSubtextStyle.FontSize,
			SubtextFontColor: titleSubtextStyle.FontColor,
			Left:             string(eo.Title.Left),
			Top:              string(eo.Title.Top),
		},
		Legend: LegendOption{
			Show:      eo.Legend.Show,
			FontSize:  legendTextStyle.FontSize,
			FontColor: legendTextStyle.FontColor,
			Data:      eo.Legend.Data,
			Left:      string(eo.Legend.Left),
			Top:       string(eo.Legend.Top),
			Align:     eo.Legend.Align,
			Orient:    eo.Legend.Orient,
		},
		RadarIndicators: eo.Radar.Indicator,
		Width:           eo.Width,
		Height:          eo.Height,
		Padding:         eo.Padding.Box,
		Box:             eo.Box,
		SeriesList:      eo.Series.ToSeriesList(),
	}
	isHorizontalChart := false
	for _, item := range eo.XAxis.Data {
		if item.Type == "value" {
			isHorizontalChart = true
		}
	}
	if isHorizontalChart {
		for index := range o.SeriesList {
			series := o.SeriesList[index]
			if series.Type == ChartTypeBar {
				o.SeriesList[index].Type = ChartTypeHorizontalBar
			}
		}
	}

	if len(eo.XAxis.Data) != 0 {
		xAxisData := eo.XAxis.Data[0]
		o.XAxis = XAxisOption{
			BoundaryGap: xAxisData.BoundaryGap,
			Data:        xAxisData.Data,
			SplitNumber: xAxisData.SplitNumber,
		}
	}
	yAxisOptions := make([]YAxisOption, len(eo.YAxis.Data))
	for index, item := range eo.YAxis.Data {
		yAxisOptions[index] = YAxisOption{
			Min:       item.Min,
			Max:       item.Max,
			Formatter: item.AxisLabel.Formatter,
			Color:     parseColor(item.AxisLine.LineStyle.Color),
			Data:      item.Data,
		}
	}
	o.YAxisOptions = yAxisOptions

	if len(eo.Children) != 0 {
		o.Children = make([]ChartOption, len(eo.Children))
		for index, item := range eo.Children {
			o.Children[index] = item.ToOption()
		}
	}
	return o
}

func renderEcharts(options, outputType string) ([]byte, error) {
	o := EChartsOption{}
	err := json.Unmarshal([]byte(options), &o)
	if err != nil {
		return nil, err
	}
	opt := o.ToOption()
	opt.Type = outputType
	d, err := Render(opt)
	if err != nil {
		return nil, err
	}
	return d.Bytes()
}

func RenderEChartsToPNG(options string) ([]byte, error) {
	return renderEcharts(options, "png")
}

func RenderEChartsToSVG(options string) ([]byte, error) {
	return renderEcharts(options, "svg")
}
