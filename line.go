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
)

type LineStyle struct {
	ClassName       string
	StrokeDashArray []float64
	StrokeColor     Color
	StrokeWidth     float64
	FillColor       Color
	DotWidth        float64
	DotColor        Color
	DotFillColor    Color
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

	newPoints := make([]Point, len(points))
	copy(newPoints, points)
	x0 := points[0].X
	y0 := points[0].Y
	height := d.Box.Height()
	newPoints = append(newPoints, Point{
		X: points[len(points)-1].X,
		Y: height,
	}, Point{
		X: x0,
		Y: height,
	}, Point{
		X: x0,
		Y: y0,
	})
	d.fill(newPoints, style.Style())
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
