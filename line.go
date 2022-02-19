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
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

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

func (d *Draw) lineFill(points []Point, style LineStyle) {
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

func (d *Draw) lineDot(points []Point, style LineStyle) {
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

func (d *Draw) Line(points []Point, style LineStyle) {
	if len(points) == 0 {
		return
	}
	d.lineFill(points, style)
	d.lineStroke(points, style)
	d.lineDot(points, style)
}
