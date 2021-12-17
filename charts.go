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
	"errors"
	"io"

	"github.com/wcharczuk/go-chart/v2"
)

const (
	ThemeLight = "light"
	ThemeDark  = "dark"
)

const (
	DefaultChartWidth  = 800
	DefaultChartHeight = 400
)

type (
	Title struct {
		Text  string
		Style chart.Style
	}
	Legend struct {
		Data []string
	}
	Options struct {
		Width        int
		Height       int
		Theme        string
		XAxis        XAxis
		YAxisOptions []*YAxisOption
		Series       []Series
		Title        Title
		Legend       Legend
		TickPosition chart.TickPosition
	}
)

type Graph interface {
	Render(rp chart.RendererProvider, w io.Writer) error
}

func (o *Options) validate() error {
	xAxisCount := len(o.XAxis.Data)
	if len(o.Series) == 0 {
		return errors.New("series can not be empty")
	}

	for _, item := range o.Series {
		if item.Type != SeriesPie && len(item.Data) != xAxisCount {
			return errors.New("series and xAxis is not matched")
		}
	}
	return nil
}

func render(g Graph, rp chart.RendererProvider) ([]byte, error) {
	buf := bytes.Buffer{}
	err := g.Render(rp, &buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ToPNG(g Graph) ([]byte, error) {
	return render(g, chart.PNG)
}

func ToSVG(g Graph) ([]byte, error) {
	return render(g, chart.SVG)
}
func New(opt Options) (Graph, error) {
	err := opt.validate()
	if err != nil {
		return nil, err
	}
	tickPosition := opt.TickPosition
	width := opt.Width
	if width <= 0 {
		width = DefaultChartWidth
	}
	height := opt.Height
	if height <= 0 {
		height = DefaultChartHeight
	}
	if opt.Series[0].Type == SeriesPie {
		values := make(chart.Values, len(opt.Series))
		for index, item := range opt.Series {
			values[index] = chart.Value{
				Value: item.Data[0].Value,
				Label: item.Name,
			}
		}
		g := &chart.PieChart{
			Title:      opt.Title.Text,
			TitleStyle: opt.Title.Style,
			Width:      width,
			Height:     height,
			Values:     values,
			ColorPalette: &ThemeColorPalette{
				Theme: opt.Theme,
			},
		}
		return g, nil
	}

	xAxis, xValues := GetXAxisAndValues(opt.XAxis, tickPosition, opt.Theme)

	legendSize := len(opt.Legend.Data)
	for index, item := range opt.Series {
		if len(item.XValues) == 0 {
			opt.Series[index].XValues = xValues
		}
		if index < legendSize && opt.Series[index].Name == "" {
			opt.Series[index].Name = opt.Legend.Data[index]
		}
	}

	var yAxisOption *YAxisOption
	if len(opt.YAxisOptions) != 0 {
		yAxisOption = opt.YAxisOptions[0]
	}
	var secondaryYAxisOption *YAxisOption
	if len(opt.YAxisOptions) > 1 {
		secondaryYAxisOption = opt.YAxisOptions[1]
	}

	g := &chart.Chart{
		ColorPalette: &ThemeColorPalette{
			Theme: opt.Theme,
		},
		Title:          opt.Title.Text,
		TitleStyle:     opt.Title.Style,
		Width:          width,
		Height:         height,
		XAxis:          xAxis,
		YAxis:          GetYAxis(opt.Theme, yAxisOption),
		YAxisSecondary: GetSecondaryYAxis(opt.Theme, secondaryYAxisOption),
		Series:         GetSeries(opt.Series, tickPosition, opt.Theme),
	}

	// 设置secondary的样式
	if legendSize != 0 {
		g.Elements = []chart.Renderable{
			LegendCustomize(g, LegendOption{
				Theme:        opt.Theme,
				TextPosition: LegendTextPositionRight,
				IconDraw:     DefaultLegendIconDraw,
			}),
		}
	}
	return g, nil
}
