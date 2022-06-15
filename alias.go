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

type Box = chart.Box
type Style = chart.Style
type Color = drawing.Color

var BoxZero = chart.BoxZero

type Point struct {
	X int
	Y int
}

const (
	ChartTypeLine   = "line"
	ChartTypeBar    = "bar"
	ChartTypePie    = "pie"
	ChartTypeRadar  = "radar"
	ChartTypeFunnel = "funnel"
	// horizontal bar
	ChartTypeHorizontalBar = "horizontalBar"
)

const (
	ChartOutputSVG = "svg"
	ChartOutputPNG = "png"
)

const (
	PositionLeft   = "left"
	PositionRight  = "right"
	PositionCenter = "center"
	PositionTop    = "top"
	PositionBottom = "bottom"
)

const (
	AlignLeft   = "left"
	AlignRight  = "right"
	AlignCenter = "center"
)

const (
	OrientHorizontal = "horizontal"
	OrientVertical   = "vertical"
)
