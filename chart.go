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
	"errors"
	"math"
	"sort"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const (
	ChartTypeLine   = "line"
	ChartTypeBar    = "bar"
	ChartTypePie    = "pie"
	ChartTypeRadar  = "radar"
	ChartTypeFunnel = "funnel"
)

const (
	ChartOutputSVG = "svg"
	ChartOutputPNG = "png"
)

type Point struct {
	X int
	Y int
}

const labelFontSize = 10
const defaultDotWidth = 2.0
const defaultStrokeWidth = 2.0

var defaultChartWidth = 600
var defaultChartHeight = 400

type ChartOption struct {
	// The output type of chart, "svg" or "png", default value is "svg"
	Type string
	// The font family, which should be installed first
	FontFamily string
	// The font of chart, the default font is "roboto"
	Font *truetype.Font
	// The theme of chart, "light" and "dark".
	// The default theme is "light"
	Theme string
	// The title option
	Title TitleOption
	// The legend option
	Legend LegendOption
	// The x axis option
	XAxis XAxisOption
	// The y axis option list
	YAxisList []YAxisOption
	// The width of chart, default width is 600
	Width int
	// The height of chart, default height is 400
	Height int
	Parent *Draw
	// The padding for chart, default padding is [20, 10, 10, 10]
	Padding chart.Box
	// The canvas box for chart
	Box chart.Box
	// The series list
	SeriesList SeriesList
	// The radar indicator list
	RadarIndicators []RadarIndicator
	// The background color of chart
	BackgroundColor drawing.Color
	// The child charts
	Children []ChartOption
}

// FillDefault fills the default value for chart option
func (o *ChartOption) FillDefault(theme string) {
	t := NewTheme(theme)
	// 如果为空，初始化
	yAxisCount := 1
	for _, series := range o.SeriesList {
		if series.YAxisIndex >= yAxisCount {
			yAxisCount++
		}
	}
	yAxisList := make([]YAxisOption, yAxisCount)
	copy(yAxisList, o.YAxisList)
	o.YAxisList = yAxisList

	if o.Font == nil {
		o.Font, _ = chart.GetDefaultFont()
	}
	if o.BackgroundColor.IsZero() {
		o.BackgroundColor = t.GetBackgroundColor()
	}
	if o.Padding.IsZero() {
		o.Padding = chart.Box{
			Top:    10,
			Right:  10,
			Bottom: 10,
			Left:   10,
		}
	}

	// 标题的默认值
	if o.Title.Style.FontColor.IsZero() {
		o.Title.Style.FontColor = t.GetTextColor()
	}
	if o.Title.Style.FontSize == 0 {
		o.Title.Style.FontSize = 14
	}
	if o.Title.Style.Font == nil {
		o.Title.Style.Font = o.Font
	}
	if o.Title.Style.Padding.IsZero() {
		o.Title.Style.Padding = chart.Box{
			Bottom: 10,
		}
	}
	// 副标题
	if o.Title.SubtextStyle.FontColor.IsZero() {
		o.Title.SubtextStyle.FontColor = o.Title.Style.FontColor.WithAlpha(180)
	}
	if o.Title.SubtextStyle.FontSize == 0 {
		o.Title.SubtextStyle.FontSize = labelFontSize
	}
	if o.Title.SubtextStyle.Font == nil {
		o.Title.SubtextStyle.Font = o.Font
	}

	o.Legend.theme = theme
	if o.Legend.Style.FontSize == 0 {
		o.Legend.Style.FontSize = labelFontSize
	}
	if o.Legend.Left == "" {
		o.Legend.Left = PositionCenter
	}
	// legend与series name的关联
	if len(o.Legend.Data) == 0 {
		o.Legend.Data = o.SeriesList.Names()
	} else {
		seriesCount := len(o.SeriesList)
		for index, name := range o.Legend.Data {
			if index < seriesCount &&
				len(o.SeriesList[index].Name) == 0 {
				o.SeriesList[index].Name = name
			}
		}
		nameIndexDict := map[string]int{}
		for index, name := range o.Legend.Data {
			nameIndexDict[name] = index
		}
		// 保证series的顺序与legend一致
		sort.Slice(o.SeriesList, func(i, j int) bool {
			return nameIndexDict[o.SeriesList[i].Name] < nameIndexDict[o.SeriesList[j].Name]
		})
	}
	// 如果无legend数据，则隐藏
	if len(strings.Join(o.Legend.Data, "")) == 0 {
		o.Legend.Show = FalseFlag()
	}
	if o.Legend.Style.Font == nil {
		o.Legend.Style.Font = o.Font
	}
	if o.Legend.Style.FontColor.IsZero() {
		o.Legend.Style.FontColor = t.GetTextColor()
	}
	if o.XAxis.Theme == "" {
		o.XAxis.Theme = theme
	}
}

func (o *ChartOption) getWidth() int {
	if o.Width != 0 {
		return o.Width
	}
	if o.Parent != nil {
		return o.Parent.Box.Width()
	}
	return defaultChartWidth
}

func SetDefaultWidth(width int) {
	if width > 0 {
		defaultChartWidth = width
	}
}
func SetDefaultHeight(height int) {
	if height > 0 {
		defaultChartHeight = height
	}
}

func (o *ChartOption) getHeight() int {

	if o.Height != 0 {
		return o.Height
	}
	if o.Parent != nil {
		return o.Parent.Box.Height()
	}
	return defaultChartHeight
}

func (o *ChartOption) newYRange(axisIndex int) Range {
	min := math.MaxFloat64
	max := -math.MaxFloat64
	if axisIndex >= len(o.YAxisList) {
		axisIndex = 0
	}
	yAxis := o.YAxisList[axisIndex]

	for _, series := range o.SeriesList {
		if series.YAxisIndex != axisIndex {
			continue
		}
		for _, item := range series.Data {
			if item.Value > max {
				max = item.Value
			}
			if item.Value < min {
				min = item.Value
			}
		}
	}
	min = min * 0.9
	max = max * 1.1
	if yAxis.Min != nil {
		min = *yAxis.Min
	}
	if yAxis.Max != nil {
		max = *yAxis.Max
	}
	divideCount := 6
	// y轴分设置默认划分为6块
	r := NewRange(min, max, divideCount)

	// 由于NewRange会重新计算min max
	if yAxis.Min != nil {
		r.Min = min
	}
	if yAxis.Max != nil {
		r.Max = max
	}

	return r
}

type basicRenderResult struct {
	xRange     *Range
	yRangeList []*Range
	d          *Draw
	titleBox   chart.Box
}

func (r *basicRenderResult) getYRange(index int) *Range {
	if index >= len(r.yRangeList) {
		index = 0
	}
	return r.yRangeList[index]
}

// Render renders the chart by option
func Render(opt ChartOption) (*Draw, error) {
	if len(opt.SeriesList) == 0 {
		return nil, errors.New("series can not be nil")
	}
	if len(opt.FontFamily) != 0 {
		f, err := GetFont(opt.FontFamily)
		if err != nil {
			return nil, err
		}
		opt.Font = f
	}
	opt.FillDefault(opt.Theme)

	lineSeries := make([]Series, 0)
	barSeries := make([]Series, 0)
	isPieChart := false
	isRadarChart := false
	isFunnelChart := false
	for index := range opt.SeriesList {
		opt.SeriesList[index].index = index
		item := opt.SeriesList[index]
		switch item.Type {
		case ChartTypePie:
			isPieChart = true
		case ChartTypeRadar:
			isRadarChart = true
		case ChartTypeFunnel:
			isFunnelChart = true
		case ChartTypeBar:
			barSeries = append(barSeries, item)
		default:
			lineSeries = append(lineSeries, item)
		}
	}
	// 如果指定了pie，则以pie的形式处理，pie不支持多类型图表
	// pie不需要axis
	// radar 同样处理
	if isPieChart ||
		isRadarChart ||
		isFunnelChart {
		opt.XAxis.Hidden = true
		for index := range opt.YAxisList {
			opt.YAxisList[index].Hidden = true
		}
	}
	result, err := chartBasicRender(&opt)
	if err != nil {
		return nil, err
	}
	markPointRenderOptions := make([]markPointRenderOption, 0)
	fns := []func() error{
		// pie render
		func() error {
			if !isPieChart {
				return nil
			}
			return pieChartRender(pieChartOption{
				SeriesList: opt.SeriesList,
				Theme:      opt.Theme,
				Font:       opt.Font,
			}, result)
		},
		// radar render
		func() error {
			if !isRadarChart {
				return nil
			}
			return radarChartRender(radarChartOption{
				SeriesList: opt.SeriesList,
				Theme:      opt.Theme,
				Font:       opt.Font,
				Indicators: opt.RadarIndicators,
			}, result)
		},
		// funnel render
		func() error {
			if !isFunnelChart {
				return nil
			}
			return funnelChartRender(funnelChartOption{
				SeriesList: opt.SeriesList,
				Theme:      opt.Theme,
				Font:       opt.Font,
			}, result)
		},
		// bar render
		func() error {
			// 如果无bar类型的series
			if len(barSeries) == 0 {
				return nil
			}
			options, err := barChartRender(barChartOption{
				SeriesList: barSeries,
				Theme:      opt.Theme,
				Font:       opt.Font,
			}, result)
			if err != nil {
				return err
			}
			markPointRenderOptions = append(markPointRenderOptions, options...)
			return nil
		},
		// line render
		func() error {
			// 如果无line类型的series
			if len(lineSeries) == 0 {
				return nil
			}
			options, err := lineChartRender(lineChartOption{
				Theme:      opt.Theme,
				SeriesList: lineSeries,
				Font:       opt.Font,
			}, result)
			if err != nil {
				return err
			}
			markPointRenderOptions = append(markPointRenderOptions, options...)
			return nil
		},
		// legend需要在顶层，因此此处render
		func() error {
			_, err := NewLegend(result.d, opt.Legend).Render()
			return err
		},
		// mark point最后render
		func() error {
			// mark point render不会出错
			for _, opt := range markPointRenderOptions {
				markPointRender(opt)
			}
			return nil
		},
	}

	for _, fn := range fns {
		err = fn()
		if err != nil {
			return nil, err
		}
	}
	for _, child := range opt.Children {
		child.Parent = result.d
		if len(child.Theme) == 0 {
			child.Theme = opt.Theme
		}
		_, err = Render(child)
		if err != nil {
			return nil, err
		}
	}
	return result.d, nil
}

func chartBasicRender(opt *ChartOption) (*basicRenderResult, error) {
	d, err := NewDraw(
		DrawOption{
			Type:   opt.Type,
			Parent: opt.Parent,
			Width:  opt.getWidth(),
			Height: opt.getHeight(),
		},
		BoxOption(opt.Box),
		PaddingOption(opt.Padding),
	)
	if err != nil {
		return nil, err
	}

	if len(opt.YAxisList) > 2 {
		return nil, errors.New("y axis should not be gt 2")
	}
	if opt.Parent == nil {
		d.setBackground(opt.getWidth(), opt.getHeight(), opt.BackgroundColor)
	}

	// 标题
	titleBox, err := drawTitle(d, &opt.Title)
	if err != nil {
		return nil, err
	}

	xAxisHeight := 0
	var xRange *Range

	if !opt.XAxis.Hidden {
		// xAxis
		xAxisHeight, xRange, err = drawXAxis(d, &opt.XAxis, len(opt.YAxisList))
		if err != nil {
			return nil, err
		}
	}

	yRangeList := make([]*Range, len(opt.YAxisList))

	for index, yAxis := range opt.YAxisList {
		var yRange *Range
		if !yAxis.Hidden {
			yRange, err = drawYAxis(d, opt, index, xAxisHeight, chart.Box{
				Top: titleBox.Height(),
			})
			if err != nil {
				return nil, err
			}
			yRangeList[index] = yRange
		}
	}
	return &basicRenderResult{
		xRange:     xRange,
		yRangeList: yRangeList,
		d:          d,
		titleBox:   titleBox,
	}, nil
}
