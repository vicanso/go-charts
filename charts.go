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

	"github.com/wcharczuk/go-chart/v2"
)

type (
	Title struct {
		Text string
	}
	Legend struct {
		Data []string
	}
	Option struct {
		Width  int
		Height int
		Theme  string
		XAxis  XAxis
		// YAxis  Axis
		Series       []Series
		Title        Title
		Legend       Legend
		TickPosition chart.TickPosition
	}
	ECharOption struct {
		Title struct {
			Text      string
			TextStyle struct {
				Color      string
				FontFamily string
			}
		}
		XAxis struct {
			Type        string
			BoundaryGap bool
			Data        []string
		}
	}
)

const (
	ThemeLight = "light"
	ThemeDark  = "dark"
)

func render(c *chart.Chart, rp chart.RendererProvider) ([]byte, error) {
	buf := bytes.Buffer{}
	err := c.Render(rp, &buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ToPNG(c *chart.Chart) ([]byte, error) {
	return render(c, chart.PNG)
}

func ToSVG(c *chart.Chart) ([]byte, error) {
	return render(c, chart.SVG)
}
func New(opt Option) *chart.Chart {
	tickPosition := opt.TickPosition

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

	c := &chart.Chart{
		Title:  opt.Title.Text,
		Width:  opt.Width,
		Height: opt.Height,
		XAxis:  xAxis,
		YAxis:  GetYAxis(opt.Theme),
		Series: GetSeries(opt.Series, tickPosition, opt.Theme),
	}
	// TODO 校验xAxis与yAxis的数量是否一致
	// 设置secondary的样式
	c.YAxisSecondary.Style = c.YAxis.Style
	if legendSize != 0 {
		c.Elements = []chart.Renderable{
			DefaultLegend(c),
		}
	}
	return c
}
