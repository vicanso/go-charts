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

// NewMarkLine returns a series mark line
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

type markLinePainter struct {
	p       *Painter
	options []markLineRenderOption
}

func (m *markLinePainter) Add(opt markLineRenderOption) {
	m.options = append(m.options, opt)
}

// NewMarkLinePainter returns a mark line renderer
func NewMarkLinePainter(p *Painter) *markLinePainter {
	return &markLinePainter{
		p:       p,
		options: make([]markLineRenderOption, 0),
	}
}

type markLineRenderOption struct {
	FillColor   Color
	FontColor   Color
	StrokeColor Color
	Font        *truetype.Font
	Series      Series
	Range       axisRange
}

func (m *markLinePainter) Render() (Box, error) {
	painter := m.p
	for _, opt := range m.options {
		s := opt.Series
		if len(s.MarkLine.Data) == 0 {
			continue
		}
		font := opt.Font
		if font == nil {
			font, _ = GetDefaultFont()
		}
		summary := s.Summary()
		for _, markLine := range s.MarkLine.Data {
			// 由于mark line会修改style，因此每次重新设置
			painter.OverrideDrawingStyle(Style{
				FillColor:   opt.FillColor,
				StrokeColor: opt.StrokeColor,
				StrokeWidth: 1,
				StrokeDashArray: []float64{
					4,
					2,
				},
			}).OverrideTextStyle(Style{
				Font:      font,
				FontColor: opt.FontColor,
				FontSize:  labelFontSize,
			})
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
			width := painter.Width()
			text := commafWithDigits(value)
			textBox := painter.MeasureText(text)
			painter.MarkLine(0, y, width-2)
			painter.Text(text, width, y+textBox.Height()>>1-2)
		}
	}
	return BoxZero, nil
}
