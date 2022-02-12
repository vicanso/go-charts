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
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func NewMarkPoint(markPointTypes ...string) SeriesMarkPoint {
	data := make([]SeriesMarkData, len(markPointTypes))
	for index, t := range markPointTypes {
		data[index] = SeriesMarkData{
			Type: t,
		}
	}
	return SeriesMarkPoint{
		Data: data,
	}
}

type markPointRenderOption struct {
	Draw      *Draw
	FillColor drawing.Color
	Font      *truetype.Font
	Series    *Series
	Points    []Point
}

func markPointRender(opt *markPointRenderOption) {
	d := opt.Draw
	s := opt.Series
	if len(s.MarkPoint.Data) == 0 {
		return
	}
	points := opt.Points
	summary := s.Summary()
	symbolSize := s.MarkPoint.SymbolSize
	if symbolSize == 0 {
		symbolSize = 30
	}
	r := d.Render
	// 设置填充样式
	chart.Style{
		FillColor: opt.FillColor,
	}.WriteToRenderer(r)
	// 设置文本样式
	chart.Style{
		FontColor:   NewTheme(ThemeDark).GetTextColor(),
		FontSize:    labelFontSize,
		StrokeWidth: 1,
		Font:        opt.Font,
	}.WriteTextOptionsToRenderer(r)
	for _, markPointData := range s.MarkPoint.Data {
		p := points[summary.MinIndex]
		value := summary.MinValue
		switch markPointData.Type {
		case SeriesMarkDataTypeMax:
			p = points[summary.MaxIndex]
			value = summary.MaxValue
		}

		d.pin(p.X, p.Y-symbolSize>>1, symbolSize)
		text := commafWithDigits(value)
		textBox := r.MeasureText(text)
		d.text(text, p.X-textBox.Width()>>1, p.Y-symbolSize>>1-2)
	}
}
