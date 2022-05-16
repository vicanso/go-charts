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
	"errors"
	"math"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
)

type Painter struct {
	render        Renderer
	box           Box
	font          *truetype.Font
	parent        *Painter
	style         Style
	previousStyle Style
	theme         *Theme
}

type PainterOptions struct {
	// Draw type, "svg" or "png", default type is "svg"
	Type string
	// The width of draw painter
	Width int
	// The height of draw painter
	Height int
	// The font for painter
	Font *truetype.Font
}

type PainterOption func(*Painter)

// PainterPaddingOption sets the padding of draw painter
func PainterPaddingOption(padding Box) PainterOption {
	return func(p *Painter) {
		p.box.Left += padding.Left
		p.box.Top += padding.Top
		p.box.Right -= padding.Right
		p.box.Bottom -= padding.Bottom
	}
}

// PainterBoxOption sets the box of draw painter
func PainterBoxOption(box Box) PainterOption {
	return func(p *Painter) {
		if box.IsZero() {
			return
		}
		p.box = box
	}
}

// PainterFontOption sets the font of draw painter
func PainterFontOption(font *truetype.Font) PainterOption {
	return func(p *Painter) {
		if font == nil {
			return
		}
		p.font = font
	}
}

// PainterStyleOption sets the style of draw painter
func PainterStyleOption(style Style) PainterOption {
	return func(p *Painter) {
		p.SetDrawingStyle(style)
	}
}

// PainterThemeOption sets the theme of draw painter
func PainterThemeOption(theme *Theme) PainterOption {
	return func(p *Painter) {
		if theme == nil {
			return
		}
		p.theme = theme
	}
}

func PainterWidthHeightOption(width, height int) PainterOption {
	return func(p *Painter) {
		if width > 0 {
			p.box.Right = p.box.Left + width
		}
		if height > 0 {
			p.box.Bottom = p.box.Top + height
		}
	}
}

// NewPainter creates a new painter
func NewPainter(opts PainterOptions, opt ...PainterOption) (*Painter, error) {
	if opts.Width <= 0 || opts.Height <= 0 {
		return nil, errors.New("width/height can not be nil")
	}
	font := opts.Font
	if font == nil {
		f, err := chart.GetDefaultFont()
		if err != nil {
			return nil, err
		}
		font = f
	}
	fn := chart.SVG
	if opts.Type == ChartOutputPNG {
		fn = chart.PNG
	}
	width := opts.Width
	height := opts.Height
	r, err := fn(width, height)
	if err != nil {
		return nil, err
	}
	r.SetFont(font)

	p := &Painter{
		render: r,
		box: Box{
			Right:  opts.Width,
			Bottom: opts.Height,
		},
		font: font,
	}
	p.setOptions(opt...)
	return p, nil
}
func (p *Painter) setOptions(opts ...PainterOption) {
	for _, fn := range opts {
		fn(p)
	}
}

func (p *Painter) Child(opt ...PainterOption) *Painter {
	child := &Painter{
		render:        p.render,
		box:           p.box.Clone(),
		font:          p.font,
		parent:        p,
		style:         p.style,
		previousStyle: p.previousStyle,
		theme:         p.theme,
	}
	child.setOptions(opt...)
	return child
}

func (p *Painter) SetStyle(style Style) {
	if style.Font == nil {
		style.Font = p.font
	}
	p.previousStyle = p.style
	p.style = style
	style.WriteToRenderer(p.render)
}

func (p *Painter) SetDrawingStyle(style Style) {
	p.previousStyle = p.style
	p.style = style
	style.WriteDrawingOptionsToRenderer(p.render)
}

func (p *Painter) SetTextStyle(style Style) {
	if style.Font == nil {
		style.Font = p.font
	}
	p.previousStyle = p.style
	p.style = style
	style.WriteTextOptionsToRenderer(p.render)
}

func (p *Painter) RestoreStyle() {
	p.style = p.previousStyle
}

// Bytes returns the data of draw canvas
func (p *Painter) Bytes() ([]byte, error) {
	buffer := bytes.Buffer{}
	err := p.render.Save(&buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), err
}

// MoveTo moves the cursor to a given point
func (p *Painter) MoveTo(x, y int) {
	p.render.MoveTo(x+p.box.Left, y+p.box.Top)
}

func (p *Painter) ArcTo(cx, cy int, rx, ry, startAngle, delta float64) {
	p.render.ArcTo(cx+p.box.Left, cy+p.box.Top, rx, ry, startAngle, delta)
}

func (p *Painter) LineTo(x, y int) {
	p.render.LineTo(x+p.box.Left, y+p.box.Top)
}

func (p *Painter) Pin(x, y, width int) {
	r := float64(width) / 2
	y -= width / 4
	angle := chart.DegreesToRadians(15)
	box := p.box

	startAngle := math.Pi/2 + angle
	delta := 2*math.Pi - 2*angle
	p.ArcTo(x, y, r, r, startAngle, delta)
	p.LineTo(x, y)
	p.Close()
	p.FillStroke()

	startX := x - int(r)
	startY := y
	endX := x + int(r)
	endY := y
	p.MoveTo(startX, startY)

	left := box.Left
	top := box.Top
	cx := x
	cy := y + int(r*2.5)
	p.render.QuadCurveTo(cx+left, cy+top, endX+left, endY+top)
	p.Close()
	p.Fill()
}

func (p *Painter) arrow(x, y, width, height int, direction string) {
	halfWidth := width >> 1
	halfHeight := height >> 1
	if direction == PositionTop || direction == PositionBottom {
		x0 := x - halfWidth
		x1 := x0 + width
		dy := -height / 3
		y0 := y
		y1 := y0 - height
		if direction == PositionBottom {
			y0 = y - height
			y1 = y
			dy = 2 * dy
		}
		p.MoveTo(x0, y0)
		p.LineTo(x0+halfWidth, y1)
		p.LineTo(x1, y0)
		p.LineTo(x0+halfWidth, y+dy)
		p.LineTo(x0, y0)
	} else {
		x0 := x + width
		x1 := x0 - width
		y0 := y - halfHeight
		dx := -width / 3
		if direction == PositionRight {
			x0 = x - width
			dx = -dx
			x1 = x0 + width
		}
		p.MoveTo(x0, y0)
		p.LineTo(x1, y0+halfHeight)
		p.LineTo(x0, y0+height)
		p.LineTo(x0+dx, y0+halfHeight)
		p.LineTo(x0, y0)
	}
	p.FillStroke()
}

func (p *Painter) ArrowLeft(x, y, width, height int) {
	p.arrow(x, y, width, height, PositionLeft)
}

func (p *Painter) ArrowRight(x, y, width, height int) {
	p.arrow(x, y, width, height, PositionRight)
}

func (p *Painter) ArrowTop(x, y, width, height int) {
	p.arrow(x, y, width, height, PositionTop)
}
func (p *Painter) ArrowBottom(x, y, width, height int) {
	p.arrow(x, y, width, height, PositionBottom)
}

func (p *Painter) Circle(radius float64, x, y int) {
	p.render.Circle(radius, x+p.box.Left, y+p.box.Top)
}

func (p *Painter) Stroke() {
	p.render.Stroke()
}

func (p *Painter) Close() {
	p.render.Close()
}

func (p *Painter) FillStroke() {
	p.render.FillStroke()
}

func (p *Painter) Fill() {
	p.render.Fill()
}

func (p *Painter) Width() int {
	return p.box.Width()
}

func (p *Painter) Height() int {
	return p.box.Height()
}

func (p *Painter) MeasureText(text string) Box {
	return p.render.MeasureText(text)
}

func (p *Painter) SetStrokeColor(color Color) {
	p.render.SetStrokeColor(color)
}

func (p *Painter) LineStroke(points []Point, style LineStyle) {
	s := style.Style()
	if !s.ShouldDrawStroke() {
		return
	}
	defer p.RestoreStyle()
	p.SetDrawingStyle(s.GetStrokeOptions())
	for index, point := range points {
		x := point.X
		y := point.Y
		if index == 0 {
			p.MoveTo(x, y)
		} else {
			p.LineTo(x, y)
		}
	}
	p.Stroke()
}

func (p *Painter) SetBackground(width, height int, color Color) {
	r := p.render
	s := chart.Style{
		FillColor: color,
	}
	defer p.RestoreStyle()
	p.SetStyle(s)
	// 设置背景色不使用box，因此不直接使用Painter
	r.MoveTo(0, 0)
	r.LineTo(width, 0)
	r.LineTo(width, height)
	r.LineTo(0, height)
	r.LineTo(0, 0)
	p.FillStroke()
}
func (p *Painter) MarkLine(x, y, width int) {
	arrowWidth := 16
	arrowHeight := 10
	endX := x + width
	p.Circle(3, x, y)
	p.render.Fill()
	p.MoveTo(x+5, y)
	p.LineTo(endX-arrowWidth, y)
	p.Stroke()
	p.render.SetStrokeDashArray([]float64{})
	p.ArrowRight(endX, y, arrowWidth, arrowHeight)
}

func (p *Painter) Polygon(center Point, radius float64, sides int) {
	points := getPolygonPoints(center, radius, sides)
	for i, item := range points {
		if i == 0 {
			p.MoveTo(item.X, item.Y)
		} else {
			p.LineTo(item.X, item.Y)
		}
	}
	p.LineTo(points[0].X, points[0].Y)
	p.Stroke()
}

func (p *Painter) FillArea(points []Point, s Style) {
	if !s.ShouldDrawFill() {
		return
	}
	defer p.RestoreStyle()
	var x, y int
	p.SetDrawingStyle(s.GetFillOptions())
	for index, point := range points {
		x = point.X
		y = point.Y
		if index == 0 {
			p.MoveTo(x, y)
		} else {
			p.LineTo(x, y)
		}
	}
	p.Fill()
}

func (p *Painter) Text(body string, x, y int) {
	p.render.Text(body, x+p.box.Left, y+p.box.Top)
}

func (p *Painter) TextFit(body string, x, y, width int) chart.Box {
	style := p.style
	textWarp := style.TextWrap
	style.TextWrap = chart.TextWrapWord
	r := p.render
	lines := chart.Text.WrapFit(r, body, width, style)
	p.SetTextStyle(style)
	var output chart.Box

	for index, line := range lines {
		x0 := x
		y0 := y + output.Height()
		p.Text(line, x0, y0)
		lineBox := r.MeasureText(line)
		output.Right = chart.MaxInt(lineBox.Right, output.Right)
		output.Bottom += lineBox.Height()
		if index < len(lines)-1 {
			output.Bottom += +style.GetTextLineSpacing()
		}
	}
	p.style.TextWrap = textWarp
	return output
}
