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
)

// NewMarkPoint returns a series mark point
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

type markPointPainter struct {
	p       *Painter
	options []markPointRenderOption
}

func (m *markPointPainter) Add(opt markPointRenderOption) {
	m.options = append(m.options, opt)
}

type markPointRenderOption struct {
	FillColor Color
	Font      *truetype.Font
	Series    Series
	Points    []Point
}

// NewMarkPointPainter returns a mark point renderer
func NewMarkPointPainter(p *Painter) *markPointPainter {
	return &markPointPainter{
		p:       p,
		options: make([]markPointRenderOption, 0),
	}
}

func (m *markPointPainter) Render() (Box, error) {
	painter := m.p
	for _, opt := range m.options {
		s := opt.Series
		if len(s.MarkPoint.Data) == 0 {
			continue
		}
		points := opt.Points
		summary := s.Summary()
		symbolSize := s.MarkPoint.SymbolSize
		if symbolSize == 0 {
			symbolSize = 30
		}
		textStyle := Style{
			FontSize:    labelFontSize,
			StrokeWidth: 1,
			Font:        opt.Font,
		}
		if isLightColor(opt.FillColor) {
			textStyle.FontColor = defaultLightFontColor
		} else {
			textStyle.FontColor = defaultDarkFontColor
		}
		painter.OverrideDrawingStyle(Style{
			FillColor: opt.FillColor,
		}).OverrideTextStyle(textStyle)
		for _, markPointData := range s.MarkPoint.Data {
			textStyle.FontSize = labelFontSize
			painter.OverrideTextStyle(textStyle)
			p := points[summary.MinIndex]
			value := summary.MinValue
			switch markPointData.Type {
			case SeriesMarkDataTypeMax:
				p = points[summary.MaxIndex]
				value = summary.MaxValue
			}

			painter.Pin(p.X, p.Y-symbolSize>>1, symbolSize)
			text := commafWithDigits(value)
			textBox := painter.MeasureText(text)
			if textBox.Width() > symbolSize {
				textStyle.FontSize = smallLabelFontSize
				painter.OverrideTextStyle(textStyle)
				textBox = painter.MeasureText(text)
			}
			painter.Text(text, p.X-textBox.Width()>>1, p.Y-symbolSize>>1-2)
		}
	}
	return BoxZero, nil
}
