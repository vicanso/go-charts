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

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type AxisOption struct {
	// The boundary gap on both sides of a coordinate axis.
	// Nil or *true means the center part of two axis ticks
	BoundaryGap *bool
	// The flag for show axis, set this to *false will hide axis
	Show *bool
	// The position of axis, it can be 'left', 'top', 'right' or 'bottom'
	Position string
	// Number of segments that the axis is split into. Note that this number serves only as a recommendation.
	SplitNumber int
	ClassName   string
	// The line color of axis
	StrokeColor drawing.Color
	// The line width
	StrokeWidth float64
	// The length of the axis tick
	TickLength int
	// The flag for show axis tick, set this to *false will hide axis tick
	TickShow *bool
	// The margin value of label
	LabelMargin int
	// The font size of label
	FontSize float64
	// The font of label
	Font *truetype.Font
	// The color of label
	FontColor drawing.Color
	// The flag for show axis split line, set this to true will show axis split line
	SplitLineShow bool
	// The color of split line
	SplitLineColor drawing.Color
}

type axis struct {
	d     *Draw
	data  *AxisDataList
	style *AxisOption
}

func NewAxis(d *Draw, data AxisDataList, style AxisOption) *axis {
	return &axis{
		d:     d,
		data:  &data,
		style: &style,
	}

}

func (as *AxisOption) GetLabelMargin() int {
	return getDefaultInt(as.LabelMargin, 8)
}

func (as *AxisOption) GetTickLength() int {
	return getDefaultInt(as.TickLength, 5)
}

func (as *AxisOption) Style(f *truetype.Font) chart.Style {
	s := chart.Style{
		ClassName:   as.ClassName,
		StrokeColor: as.StrokeColor,
		StrokeWidth: as.StrokeWidth,
		FontSize:    as.FontSize,
		FontColor:   as.FontColor,
		Font:        as.Font,
	}
	if s.FontSize == 0 {
		s.FontSize = chart.DefaultFontSize
	}
	if s.Font == nil {
		s.Font = f
	}
	return s
}

type AxisData struct {
	Text string
}
type AxisDataList []AxisData

func (l AxisDataList) TextList() []string {
	textList := make([]string, len(l))
	for index, item := range l {
		textList[index] = item.Text
	}
	return textList
}

type axisRenderOption struct {
	textMaxWith   int
	textMaxHeight int
	boundaryGap   bool
	unitCount     int
	modValue      int
}

func NewAxisDataListFromStringList(textList []string) AxisDataList {
	list := make(AxisDataList, len(textList))
	for index, text := range textList {
		list[index] = AxisData{
			Text: text,
		}
	}
	return list
}

func (a *axis) axisLabel(opt *axisRenderOption) {
	style := a.style
	data := *a.data
	d := a.d
	if style.FontColor.IsZero() || len(data) == 0 {
		return
	}
	r := d.Render

	s := style.Style(d.Font)
	s.GetTextOptions().WriteTextOptionsToRenderer(r)

	width := d.Box.Width()
	height := d.Box.Height()
	textList := data.TextList()
	count := len(textList)

	boundaryGap := opt.boundaryGap
	if !boundaryGap {
		count--
	}

	unitCount := opt.unitCount
	modValue := opt.modValue
	labelMargin := style.GetLabelMargin()

	// 轴线
	labelHeight := labelMargin + opt.textMaxHeight
	labelWidth := labelMargin + opt.textMaxWith

	// 坐标轴文本
	position := style.Position
	switch position {
	case PositionLeft:
		fallthrough
	case PositionRight:
		values := autoDivide(height, count)
		textList := data.TextList()
		// 由下往上
		reverseIntSlice(values)
		for index, text := range textList {
			y := values[index] - 2
			b := r.MeasureText(text)
			if boundaryGap {
				height := y - values[index+1]
				y -= (height - b.Height()) >> 1
			} else {
				y += b.Height() >> 1
			}
			// 左右位置的x不一样
			x := width - opt.textMaxWith
			if position == PositionLeft {
				x = labelWidth - b.Width() - 1
			}
			d.text(text, x, y)
		}
	default:
		// 定位bottom，重新计算y0的定位
		y0 := height - labelMargin
		if position == PositionTop {
			y0 = labelHeight - labelMargin
		}
		values := autoDivide(width, count)
		for index, text := range data.TextList() {
			if unitCount != 0 && index%unitCount != modValue {
				continue
			}
			x := values[index]
			leftOffset := 0
			b := r.MeasureText(text)
			if boundaryGap {
				width := values[index+1] - x
				leftOffset = (width - b.Width()) >> 1
			} else {
				// 左移文本长度
				leftOffset = -b.Width() >> 1
			}
			d.text(text, x+leftOffset, y0)
		}
	}
}

func (a *axis) axisLine(opt *axisRenderOption) {
	d := a.d
	r := d.Render
	style := a.style
	s := style.Style(d.Font)
	s.GetStrokeOptions().WriteDrawingOptionsToRenderer(r)

	x0 := 0
	y0 := 0
	x1 := 0
	y1 := 0
	width := d.Box.Width()
	height := d.Box.Height()
	labelMargin := style.GetLabelMargin()

	// 轴线
	labelHeight := labelMargin + opt.textMaxHeight
	labelWidth := labelMargin + opt.textMaxWith
	tickLength := style.GetTickLength()
	switch style.Position {
	case PositionLeft:
		x0 = tickLength + labelWidth
		x1 = x0
		y0 = 0
		y1 = height
	case PositionRight:
		x0 = width - labelWidth
		x1 = x0
		y0 = 0
		y1 = height
	case PositionTop:
		x0 = 0
		x1 = width
		y0 = labelHeight
		y1 = y0
	// bottom
	default:
		x0 = 0
		x1 = width
		y0 = height - tickLength - labelHeight
		y1 = y0
	}

	d.moveTo(x0, y0)
	d.lineTo(x1, y1)
	r.FillStroke()
}

func (a *axis) axisTick(opt *axisRenderOption) {
	d := a.d
	r := d.Render

	style := a.style
	s := style.Style(d.Font)
	s.GetStrokeOptions().WriteDrawingOptionsToRenderer(r)

	width := d.Box.Width()
	height := d.Box.Height()
	data := *a.data
	tickCount := len(data)
	if !opt.boundaryGap {
		tickCount--
	}
	labelMargin := style.GetLabelMargin()
	tickShow := true
	if isFalse(style.TickShow) {
		tickShow = false
	}
	unitCount := opt.unitCount

	tickLengthValue := style.GetTickLength()
	labelHeight := labelMargin + opt.textMaxHeight
	labelWidth := labelMargin + opt.textMaxWith
	position := style.Position
	switch position {
	case PositionLeft:
		fallthrough
	case PositionRight:
		values := autoDivide(height, tickCount)
		// 左右仅是x0的位置不一样
		x0 := width - labelWidth
		if style.Position == PositionLeft {
			x0 = labelWidth
		}
		if tickShow {
			for _, v := range values {
				x := x0
				y := v
				d.moveTo(x, y)
				d.lineTo(x+tickLengthValue, y)
				r.Stroke()
			}
		}
		// 辅助线
		if style.SplitLineShow && !style.SplitLineColor.IsZero() {
			r.SetStrokeColor(style.SplitLineColor)
			splitLineWidth := width - labelWidth - tickLengthValue
			x0 = labelWidth + tickLengthValue
			if position == PositionRight {
				x0 = 0
				splitLineWidth = width - labelWidth - 1
			}
			for _, v := range values[0 : len(values)-1] {
				x := x0
				y := v
				d.moveTo(x, y)
				d.lineTo(x+splitLineWidth, y)
				r.Stroke()
			}
		}
	default:
		values := autoDivide(width, tickCount)
		// 上下仅是y0的位置不一样
		y0 := height - labelHeight
		if position == PositionTop {
			y0 = labelHeight
		}
		if tickShow {
			for index, v := range values {
				if index%unitCount != 0 {
					continue
				}
				x := v
				y := y0
				d.moveTo(x, y-tickLengthValue)
				d.lineTo(x, y)
				r.Stroke()
			}
		}
		// 辅助线
		if style.SplitLineShow && !style.SplitLineColor.IsZero() {
			r.SetStrokeColor(style.SplitLineColor)
			y0 = 0
			splitLineHeight := height - labelHeight - tickLengthValue
			if position == PositionTop {
				y0 = labelHeight
				splitLineHeight = height - labelHeight
			}

			for index, v := range values {
				if index%unitCount != 0 {
					continue
				}
				x := v
				y := y0

				d.moveTo(x, y)
				d.lineTo(x, y0+splitLineHeight)
				r.Stroke()
			}
		}
	}
}

func (a *axis) axisMeasureTextMaxWidthHeight() (int, int) {
	d := a.d
	r := d.Render
	s := a.style.Style(d.Font)
	data := a.data
	s.GetStrokeOptions().WriteDrawingOptionsToRenderer(r)
	s.GetTextOptions().WriteTextOptionsToRenderer(r)
	return measureTextMaxWidthHeight(data.TextList(), r)
}

// measureAxis returns the measurement of axis.
// If the position is left or right, it will be textMaxWidth + labelMargin + tickLength.
// If the position is top or bottom, it will be textMaxHeight + labelMargin + tickLength.
func (a *axis) measureAxis() int {
	style := a.style
	value := style.GetLabelMargin() + style.GetTickLength()
	textMaxWidth, textMaxHeight := a.axisMeasureTextMaxWidthHeight()
	if style.Position == PositionLeft ||
		style.Position == PositionRight {
		return textMaxWidth + value
	}
	return textMaxHeight + value
}

// Render renders the axis for chart
func (a *axis) Render() {
	style := a.style
	if isFalse(style.Show) {
		return
	}
	textMaxWidth, textMaxHeight := a.axisMeasureTextMaxWidthHeight()
	opt := &axisRenderOption{
		textMaxWith:   textMaxWidth,
		textMaxHeight: textMaxHeight,
		boundaryGap:   true,
	}
	if isFalse(style.BoundaryGap) {
		opt.boundaryGap = false
	}

	unitCount := chart.MaxInt(style.SplitNumber, 1)
	width := a.d.Box.Width()
	textList := a.data.TextList()
	count := len(textList)

	position := style.Position
	switch position {
	case PositionLeft:
		fallthrough
	case PositionRight:
	default:
		maxCount := width / (opt.textMaxWith + 10)
		// 可以显示所有
		if maxCount >= count {
			unitCount = 1
		} else if maxCount < count/unitCount {
			unitCount = int(math.Ceil(float64(count) / float64(maxCount)))
		}
	}

	boundaryGap := opt.boundaryGap
	modValue := 0
	if boundaryGap && unitCount > 1 {
		// 如果是居中，unit count需要设置为奇数
		if unitCount%2 == 0 {
			unitCount++
		}
		modValue = unitCount / 2
	}
	opt.modValue = modValue
	opt.unitCount = unitCount

	// 坐标轴线
	a.axisLine(opt)
	a.axisTick(opt)
	// 坐标文本
	a.axisLabel(opt)
}
