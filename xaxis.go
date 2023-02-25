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

type XAxisOption struct {
	// The font of x axis
	Font *truetype.Font
	// The boundary gap on both sides of a coordinate axis.
	// Nil or *true means the center part of two axis ticks
	BoundaryGap *bool
	// The data value of x axis
	Data []string
	// The theme of chart
	Theme ColorPalette
	// The font size of x axis label
	FontSize float64
	// The flag for show axis, set this to *false will hide axis
	Show *bool
	// Number of segments that the axis is split into. Note that this number serves only as a recommendation.
	SplitNumber int
	// The position of axis, it can be 'top' or 'bottom'
	Position string
	// The line color of axis
	StrokeColor Color
	// The color of label
	FontColor Color
	// The text rotation of label
	TextRotation float64
	// The first axis
	FirstAxis int
	// The offset of label
	LabelOffset Box
	isValueAxis bool
}

const defaultXAxisHeight = 30

// NewXAxisOption returns a x axis option
func NewXAxisOption(data []string, boundaryGap ...*bool) XAxisOption {
	opt := XAxisOption{
		Data: data,
	}
	if len(boundaryGap) != 0 {
		opt.BoundaryGap = boundaryGap[0]
	}
	return opt
}

func (opt *XAxisOption) ToAxisOption() AxisOption {
	position := PositionBottom
	if opt.Position == PositionTop {
		position = PositionTop
	}
	axisOpt := AxisOption{
		Theme:          opt.Theme,
		Data:           opt.Data,
		BoundaryGap:    opt.BoundaryGap,
		Position:       position,
		SplitNumber:    opt.SplitNumber,
		StrokeColor:    opt.StrokeColor,
		FontSize:       opt.FontSize,
		Font:           opt.Font,
		FontColor:      opt.FontColor,
		Show:           opt.Show,
		SplitLineColor: opt.Theme.GetAxisSplitLineColor(),
		TextRotation:   opt.TextRotation,
		LabelOffset:    opt.LabelOffset,
		FirstAxis:      opt.FirstAxis,
	}
	if opt.isValueAxis {
		axisOpt.SplitLineShow = true
		axisOpt.StrokeWidth = -1
		axisOpt.BoundaryGap = FalseFlag()
	}
	return axisOpt
}

// NewBottomXAxis returns a bottom x axis renderer
func NewBottomXAxis(p *Painter, opt XAxisOption) *axisPainter {
	return NewAxisPainter(p, opt.ToAxisOption())
}
