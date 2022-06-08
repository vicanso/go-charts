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

type axisPainter struct {
	p   *Painter
	opt *AxisPainterOption
}

func NewAxisPainter(p *Painter, opt AxisPainterOption) *axisPainter {
	return &axisPainter{
		p:   p,
		opt: &opt,
	}
}

type AxisPainterOption struct {
	// The label of axis
	Data []string
	// The boundary gap on both sides of a coordinate axis.
	// Nil or *true means the center part of two axis ticks
	BoundaryGap *bool
	// The position of axis, it can be 'left', 'top', 'right' or 'bottom'
	Position string
	// Number of segments that the axis is split into. Note that this number serves only as a recommendation.
	SplitNumber int
	// The line color of axis
	StrokeColor Color
	// The line width
	StrokeWidth float64
	// The length of the axis tick
	TickLength int
	// The margin value of label
	LabelMargin int
	// The font size of label
	FontSize float64
	// The font of label
	Font *truetype.Font
	// The color of label
	FontColor Color
	// The flag for show axis split line, set this to true will show axis split line
	SplitLineShow bool
	// The color of split line
	SplitLineColor Color
}

func (a *axisPainter) Render() (Box, error) {
	opt := a.opt
	p := a.p

	strokeWidth := opt.StrokeWidth
	if strokeWidth == 0 {
		strokeWidth = 1
	}

	tickCount := opt.SplitNumber
	if tickCount == 0 {
		tickCount = len(opt.Data)
	}

	boundaryGap := true
	if opt.BoundaryGap != nil && !*opt.BoundaryGap {
		boundaryGap = false
	}

	labelPosition := ""
	if !boundaryGap {
		tickCount--
		labelPosition = PositionLeft
	}

	// TODO 计算unit
	unit := 1
	// 如果小于0，则表示不处理
	tickLength := getDefaultInt(opt.TickLength, 5)
	labelMargin := getDefaultInt(opt.LabelMargin, 5)

	textMaxWidth, textMaxHeight := p.MeasureTextMaxWidthHeight(opt.Data)

	width := 0
	height := 0
	// 垂直
	if opt.Position == PositionLeft ||
		opt.Position == PositionRight {
		width = textMaxWidth + tickLength<<1
		height = p.Height()
	} else {
		width = p.Width()
		height = tickLength<<1 + textMaxHeight
	}
	padding := Box{}
	switch opt.Position {
	case PositionTop:
		padding.Top = p.Height() - height
	case PositionLeft:
		padding.Right = p.Width() - width
	}
	p = p.Child(PainterPaddingOption(padding))
	p.SetDrawingStyle(Style{
		StrokeColor: opt.StrokeColor,
		StrokeWidth: strokeWidth,
	}).OverrideTextStyle(Style{
		Font:      opt.Font,
		FontColor: opt.FontColor,
		FontSize:  opt.FontSize,
	})

	x0 := 0
	y0 := 0
	x1 := 0
	y1 := 0
	ticksPadding := 0
	labelPadding := 0
	orient := ""
	textAlign := ""

	switch opt.Position {
	case PositionTop:
		labelPadding = labelMargin
		x1 = p.Width()
		y0 = labelMargin + int(opt.FontSize)
		ticksPadding = int(opt.FontSize)
		y1 = y0
		orient = OrientHorizontal
	case PositionLeft:
		orient = OrientVertical
		textAlign = AlignRight
	default:
		labelPadding = height
		x1 = p.Width()
		orient = OrientHorizontal
	}

	p.Child(PainterPaddingOption(Box{
		Top: ticksPadding,
	})).Ticks(TicksOption{
		Count:  tickCount,
		Length: tickLength,
		Unit:   unit,
		Orient: orient,
	})

	p.LineStroke([]Point{
		{
			X: x0,
			Y: y0,
		},
		{
			X: x1,
			Y: y1,
		},
	})

	p.Child(PainterPaddingOption(Box{
		Top: labelPadding,
	})).MultiText(MultiTextOption{
		Align:    textAlign,
		TextList: opt.Data,
		Orient:   orient,
		Unit:     unit,
		Position: labelPosition,
	})

	return Box{
		Bottom: height,
		Right:  width,
	}, nil
}
