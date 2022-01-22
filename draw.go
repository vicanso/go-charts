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

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const (
	PositionLeft   = "left"
	PositionRight  = "right"
	PositionTop    = "top"
	PositionBottom = "bottom"
)

type draw struct {
	Render chart.Renderer
	Box    chart.Box
	parent *draw
}

type Point struct {
	X int
	Y int
}

type LineStyle struct {
	ClassName       string
	StrokeDashArray []float64
	StrokeColor     drawing.Color
	StrokeWidth     float64
	FillColor       drawing.Color
	DotWidth        float64
	DotColor        drawing.Color
	DotFillColor    drawing.Color
}

func (ls *LineStyle) Style() chart.Style {
	return chart.Style{
		ClassName:       ls.ClassName,
		StrokeDashArray: ls.StrokeDashArray,
		StrokeColor:     ls.StrokeColor,
		StrokeWidth:     ls.StrokeWidth,
		FillColor:       ls.FillColor,
		DotWidth:        ls.DotWidth,
		DotColor:        ls.DotColor,
	}
}

type BarStyle struct {
	ClassName       string
	StrokeDashArray []float64
	FillColor       drawing.Color
}

func (bs *BarStyle) Style() chart.Style {
	return chart.Style{
		ClassName:       bs.ClassName,
		StrokeDashArray: bs.StrokeDashArray,
		StrokeColor:     bs.FillColor,
		StrokeWidth:     1,
		FillColor:       bs.FillColor,
	}
}

type AxisStyle struct {
	BoundaryGap *bool
	Show        *bool
	Position    string
	Offset      int
	ClassName   string
	StrokeColor drawing.Color
	StrokeWidth float64
	TickLength  int
	TickShow    *bool
	LabelMargin int
	FontSize    float64
	Font        *truetype.Font
	FontColor   drawing.Color
}

func (as *AxisStyle) GetLabelMargin() int {
	return getDefaultInt(as.LabelMargin, 8)
}

func (as *AxisStyle) GetTickLength() int {
	return getDefaultInt(as.TickLength, 5)
}

func (as *AxisStyle) Style() chart.Style {
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
		s.Font, _ = chart.GetDefaultFont()
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

type axisOption struct {
	data          *AxisDataList
	style         *AxisStyle
	textMaxWith   int
	textMaxHeight int
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

type Option func(*draw)

func ParentOption(p *draw) Option {
	return func(d *draw) {
		d.parent = p
	}
}

func NewDraw(r chart.Renderer, box chart.Box, opts ...Option) *draw {
	d := &draw{
		Render: r,
		Box:    box,
	}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

func (d *draw) Parent() *draw {
	return d.parent
}

func (d *draw) Top() *draw {
	if d.parent == nil {
		return nil
	}
	t := d.parent
	// 限制最多查询次数，避免嵌套引用
	for i := 50; i > 0; i-- {
		if t.parent == nil {
			break
		}
		t = t.parent
	}
	return t
}

func (d *draw) Bytes() ([]byte, error) {
	buffer := bytes.Buffer{}
	err := d.Render.Save(&buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), err
}

func (d *draw) moveTo(x, y int) {
	d.Render.MoveTo(x+d.Box.Left, y+d.Box.Top)
}

func (d *draw) lineTo(x, y int) {
	d.Render.LineTo(x+d.Box.Left, y+d.Box.Top)
}

func (d *draw) circle(radius float64, x, y int) {
	d.Render.Circle(radius, x+d.Box.Left, y+d.Box.Top)
}

func (d *draw) text(body string, x, y int) {
	d.Render.Text(body, x+d.Box.Left, y+d.Box.Top)
}

func (d *draw) lineStroke(points []Point, style LineStyle) {
	s := style.Style()
	if !s.ShouldDrawStroke() {
		return
	}
	r := d.Render
	s.GetStrokeOptions().WriteDrawingOptionsToRenderer(r)
	for index, point := range points {
		x := point.X
		y := point.Y
		if index == 0 {
			d.moveTo(x, y)
		} else {
			d.lineTo(x, y)
		}
	}
	r.Stroke()
}

func (d *draw) lineFill(points []Point, style LineStyle) {
	s := style.Style()
	if !(s.ShouldDrawStroke() && s.ShouldDrawFill()) {
		return
	}
	r := d.Render
	var x, y int
	s.GetFillOptions().WriteDrawingOptionsToRenderer(r)
	for index, point := range points {
		x = point.X
		y = point.Y
		if index == 0 {
			d.moveTo(x, y)
		} else {
			d.lineTo(x, y)
		}
	}
	height := d.Box.Height()
	d.lineTo(x, height)
	x0 := points[0].X
	y0 := points[0].Y
	d.lineTo(x0, height)
	d.lineTo(x0, y0)
	r.Fill()
}

func (d *draw) lineDot(points []Point, style LineStyle) {
	s := style.Style()
	if !s.ShouldDrawDot() {
		return
	}
	r := d.Render
	dotWith := s.GetDotWidth()

	s.GetDotOptions().WriteDrawingOptionsToRenderer(r)
	for _, point := range points {
		if !style.DotFillColor.IsZero() {
			r.SetFillColor(style.DotFillColor)
		}
		r.SetStrokeColor(s.DotColor)
		d.circle(dotWith, point.X, point.Y)
		r.FillStroke()
	}
}

func (d *draw) Line(points []Point, style LineStyle) {
	if len(points) == 0 {
		return
	}
	d.lineFill(points, style)
	d.lineStroke(points, style)
	d.lineDot(points, style)
}

func (d *draw) Bar(b chart.Box, style BarStyle) {
	s := style.Style()

	r := d.Render
	s.GetFillAndStrokeOptions().WriteToRenderer(r)

	d.moveTo(b.Left, b.Top)
	d.lineTo(b.Right, b.Top)
	d.lineTo(b.Right, b.Bottom)
	d.lineTo(b.Left, b.Bottom)
	d.lineTo(b.Left, b.Top)
	r.FillStroke()
}

func (d *draw) axisLabel(opt *axisOption) {
	style := opt.style
	data := *opt.data
	if style.FontColor.IsZero() || len(data) == 0 {
		return
	}
	r := d.Render

	s := style.Style()
	s.GetTextOptions().WriteTextOptionsToRenderer(r)

	width := d.Box.Width()
	height := d.Box.Height()
	textList := data.TextList()
	count := len(textList)
	x0 := 0
	y0 := 0
	tickLength := style.GetTickLength()

	// 坐标轴文本
	switch style.Position {
	case PositionLeft:
		values := autoDivide(height, count)
		textList := data.TextList()
		// 由下往上
		reverseIntSlice(values)
		for index, text := range textList {
			y := values[index]
			height := y - values[index+1]
			b := r.MeasureText(text)
			y -= (height - b.Height()) >> 1
			x := x0 + opt.textMaxWith - (b.Width())
			d.text(text, x, y)
		}
	case PositionRight:
		values := autoDivide(height, count)
		textList := data.TextList()
		// 由下往上
		reverseIntSlice(values)
		for index, text := range textList {
			y := values[index]
			height := y - values[index+1]
			b := r.MeasureText(text)
			y -= (height - b.Height()) >> 1
			x := width - opt.textMaxWith
			d.text(text, x, y)
		}
	case PositionTop:
		y0 = tickLength + style.Offset
		values := autoDivide(width, count)
		maxIndex := len(values) - 2
		for index, text := range data.TextList() {
			if index > maxIndex {
				break
			}
			x := values[index]
			width := values[index+1] - x
			b := r.MeasureText(text)
			leftOffset := (width - b.Width()) >> 1
			d.text(text, x+leftOffset, y0)
		}
	default:
		// 定位bottom，重新计算y0的定位
		y0 = height - tickLength + style.Offset
		values := autoDivide(width, count)
		maxIndex := len(values) - 2
		for index, text := range data.TextList() {
			if index > maxIndex {
				break
			}
			x := values[index]
			width := values[index+1] - x
			b := r.MeasureText(text)
			leftOffset := (width - b.Width()) >> 1
			d.text(text, x+leftOffset, y0)
		}
	}
}

func (d *draw) axisLine(opt *axisOption) {

	r := d.Render
	style := opt.style
	s := style.Style()
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
		x0 = tickLength + style.Offset + labelWidth
		x1 = x0
		y0 = 0
		y1 = height
	case PositionRight:
		x0 = width + style.Offset - labelWidth
		x1 = x0
		y0 = 0
		y1 = height
	case PositionTop:
		x0 = 0
		x1 = width
		y0 = style.Offset + labelHeight
		y1 = y0
	// bottom
	default:
		x0 = 0
		x1 = width
		y0 = height - tickLength + style.Offset - labelHeight
		y1 = y0
	}

	d.moveTo(x0, y0)
	d.lineTo(x1, y1)
	r.FillStroke()
}

func (d *draw) axisTick(opt *axisOption) {
	r := d.Render

	style := opt.style
	s := style.Style()
	s.GetStrokeOptions().WriteDrawingOptionsToRenderer(r)

	x0 := 0
	y0 := 0
	width := d.Box.Width()
	height := d.Box.Height()
	data := *opt.data
	tickCount := len(data)
	labelMargin := style.GetLabelMargin()

	tickLengthValue := style.GetTickLength()
	labelHeight := labelMargin + opt.textMaxHeight
	labelWidth := labelMargin + opt.textMaxWith
	switch style.Position {
	case PositionLeft:
		x0 += labelWidth
		values := autoDivide(height, tickCount)
		for _, v := range values {
			x := x0
			y := v
			d.moveTo(x, y)
			d.lineTo(x+tickLengthValue, y)
			r.Stroke()
		}
	case PositionRight:
		values := autoDivide(height, tickCount)
		x0 = width - labelWidth
		for _, v := range values {
			x := x0
			y := v
			d.moveTo(x, y)
			d.lineTo(x+tickLengthValue, y)
			r.Stroke()
		}
	case PositionTop:
		values := autoDivide(width, tickCount)
		y0 = style.Offset + labelHeight
		for _, v := range values {
			x := v
			y := y0
			d.moveTo(x, y-tickLengthValue)
			d.lineTo(x, y)
			r.Stroke()
		}

	default:
		values := autoDivide(width, tickCount)
		y0 = height + style.Offset - labelHeight
		for _, v := range values {
			x := v
			y := y0
			d.moveTo(x, y-tickLengthValue)
			d.lineTo(x, y)
			r.Stroke()
		}
	}
}

func (d *draw) axisMeasureTextMaxWidthHeight(data AxisDataList, style AxisStyle) (int, int) {
	r := d.Render
	s := style.Style()
	s.GetStrokeOptions().WriteDrawingOptionsToRenderer(r)
	s.GetTextOptions().WriteTextOptionsToRenderer(r)
	return measureTextMaxWidthHeight(data.TextList(), r)
}

func (d *draw) Axis(data AxisDataList, style AxisStyle) {
	if style.Show != nil && !*style.Show {
		return
	}
	r := d.Render
	s := style.Style()
	s.GetTextOptions().WriteTextOptionsToRenderer(r)
	textMaxWidth, textMaxHeight := d.axisMeasureTextMaxWidthHeight(data, style)
	opt := &axisOption{
		data:          &data,
		style:         &style,
		textMaxWith:   textMaxWidth,
		textMaxHeight: textMaxHeight,
	}

	// 坐标轴线
	d.axisLine(opt)
	d.axisTick(opt)
	// 坐标文本
	d.axisLabel(opt)
}
