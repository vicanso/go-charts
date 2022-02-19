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

func NewMarkLine(markLineTypes ...string) SeriesMarkLine {
	data := make([]SeriesMarkData, len(markLineTypes))
	for index, t := range markLineTypes {
		data[index] = SeriesMarkData{
			Type: t,
		}
	}
	return SeriesMarkLine{
		Data: data,
	}
}

type markLineRenderOption struct {
	Draw        *Draw
	FillColor   drawing.Color
	FontColor   drawing.Color
	StrokeColor drawing.Color
	Font        *truetype.Font
	Series      *Series
	Range       *Range
}

func markLineRender(opt markLineRenderOption) {
	d := opt.Draw
	s := opt.Series
	if len(s.MarkLine.Data) == 0 {
		return
	}
	r := d.Render
	summary := s.Summary()
	for _, markLine := range s.MarkLine.Data {
		// 由于mark line会修改style，因此每次重新设置
		chart.Style{
			FillColor:   opt.FillColor,
			FontColor:   opt.FontColor,
			FontSize:    labelFontSize,
			StrokeColor: opt.StrokeColor,
			StrokeWidth: 1,
			Font:        opt.Font,
			StrokeDashArray: []float64{
				4,
				2,
			},
		}.WriteToRenderer(r)
		value := float64(0)
		switch markLine.Type {
		case SeriesMarkDataTypeMax:
			value = summary.MaxValue
		case SeriesMarkDataTypeMin:
			value = summary.MinValue
		default:
			value = summary.AverageValue
		}
		y := opt.Range.getRestHeight(value)
		width := d.Box.Width()
		text := commafWithDigits(value)
		textBox := r.MeasureText(text)
		d.makeLine(0, y, width-2)
		d.text(text, width, y+textBox.Height()>>1-2)
	}

}
