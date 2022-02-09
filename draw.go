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
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const (
	PositionLeft   = "left"
	PositionRight  = "right"
	PositionCenter = "center"
	PositionTop    = "top"
	PositionBottom = "bottom"
)

const (
	OrientHorizontal = "horizontal"
	OrientVertical   = "vertical"
)

type Draw struct {
	Render chart.Renderer
	Box    chart.Box
	Font   *truetype.Font
	parent *Draw
}

type DrawOption struct {
	Type   string
	Parent *Draw
	Width  int
	Height int
}

type Option func(*Draw) error

func PaddingOption(padding chart.Box) Option {
	return func(d *Draw) error {
		d.Box.Left += padding.Left
		d.Box.Top += padding.Top
		d.Box.Right -= padding.Right
		d.Box.Bottom -= padding.Bottom
		return nil
	}
}

func NewDraw(opt DrawOption, opts ...Option) (*Draw, error) {
	if opt.Parent == nil && (opt.Width <= 0 || opt.Height <= 0) {
		return nil, errors.New("parent and width/height can not be nil")
	}
	font, _ := chart.GetDefaultFont()
	d := &Draw{
		Font: font,
	}
	width := opt.Width
	height := opt.Height
	if opt.Parent != nil {
		d.parent = opt.Parent
		d.Render = d.parent.Render
		d.Box = opt.Parent.Box.Clone()
	}
	if width != 0 && height != 0 {
		d.Box.Right = width + d.Box.Left
		d.Box.Bottom = height + d.Box.Top
	}
	// 创建render
	if d.parent == nil {
		fn := chart.SVG
		if opt.Type == "png" {
			fn = chart.PNG
		}
		r, err := fn(d.Box.Right, d.Box.Bottom)
		if err != nil {
			return nil, err
		}
		d.Render = r
	}

	for _, o := range opts {
		err := o(d)
		if err != nil {
			return nil, err
		}
	}
	return d, nil
}

func (d *Draw) Parent() *Draw {
	return d.parent
}

func (d *Draw) Top() *Draw {
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

func (d *Draw) Bytes() ([]byte, error) {
	buffer := bytes.Buffer{}
	err := d.Render.Save(&buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), err
}

func (d *Draw) moveTo(x, y int) {
	d.Render.MoveTo(x+d.Box.Left, y+d.Box.Top)
}

func (d *Draw) arcTo(cx, cy int, rx, ry, startAngle, delta float64) {
	d.Render.ArcTo(cx+d.Box.Left, cy+d.Box.Top, rx, ry, startAngle, delta)
}

func (d *Draw) lineTo(x, y int) {
	d.Render.LineTo(x+d.Box.Left, y+d.Box.Top)
}

func (d *Draw) pin(x, y, width int) {
	r := float64(width) / 2
	y -= width / 4
	angle := chart.DegreesToRadians(15)

	startAngle := math.Pi/2 + angle
	delta := 2*math.Pi - 2*angle
	d.arcTo(x, y, r, r, startAngle, delta)
	d.lineTo(x, y)
	d.Render.Fill()
	startX := x - int(r)
	startY := y
	endX := x + int(r)
	endY := y
	d.moveTo(startX, startY)

	left := d.Box.Left
	top := d.Box.Top
	cx := x
	cy := y + int(r*2.5)
	d.Render.QuadCurveTo(cx+left, cy+top, endX+left, endY+top)
	d.Render.Stroke()
}

func (d *Draw) arrowLeft(x, y, width, height int) {
	d.arrow(x, y, width, height, PositionLeft)
}

func (d *Draw) arrowRight(x, y, width, height int) {
	d.arrow(x, y, width, height, PositionRight)
}

func (d *Draw) arrowTop(x, y, width, height int) {
	d.arrow(x, y, width, height, PositionTop)
}
func (d *Draw) arrowBottom(x, y, width, height int) {
	d.arrow(x, y, width, height, PositionBottom)
}

func (d *Draw) arrow(x, y, width, height int, direction string) {
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
		d.moveTo(x0, y0)
		d.lineTo(x0+halfWidth, y1)
		d.lineTo(x1, y0)
		d.lineTo(x0+halfWidth, y+dy)
		d.lineTo(x0, y0)
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
		d.moveTo(x0, y0)
		d.lineTo(x1, y0+halfHeight)
		d.lineTo(x0, y0+height)
		d.lineTo(x0+dx, y0+halfHeight)
		d.lineTo(x0, y0)
	}
	d.Render.Stroke()
}

func (d *Draw) circle(radius float64, x, y int) {
	d.Render.Circle(radius, x+d.Box.Left, y+d.Box.Top)
}

func (d *Draw) text(body string, x, y int) {
	d.Render.Text(body, x+d.Box.Left, y+d.Box.Top)
}

func (d *Draw) lineStroke(points []Point, style LineStyle) {
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

func (d *Draw) setBackground(width, height int, color drawing.Color) {
	r := d.Render
	s := chart.Style{
		FillColor: color,
	}
	s.WriteToRenderer(r)
	r.MoveTo(0, 0)
	r.LineTo(width, 0)
	r.LineTo(width, height)
	r.LineTo(0, height)
	r.LineTo(0, 0)
	r.FillStroke()
}
