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

import "github.com/golang/freetype/truetype"

type YAxisOption struct {
	// The minimun value of axis.
	Min *float64
	// The maximum value of axis.
	Max *float64
	// The font of y axis
	Font *truetype.Font
	// The data value of x axis
	Data []string
	// The theme of chart
	Theme ColorPalette
	// The font size of x axis label
	FontSize float64
	// The position of axis, it can be 'left' or 'right'
	Position string
	// The color of label
	FontColor Color
	// Formatter for y axis text value
	Formatter string
	// Color for y axis
	Color Color
	// The flag for show axis, set this to *false will hide axis
	Show           *bool
	DivideCount    int
	Unit           int
	isCategoryAxis bool
	// The flag for show axis split line, set this to true will show axis split line
	SplitLineShow *bool
}

// NewYAxisOptions returns a y axis option
func NewYAxisOptions(data []string, others ...[]string) []YAxisOption {
	arr := [][]string{
		data,
	}
	arr = append(arr, others...)
	opts := make([]YAxisOption, 0)
	for _, data := range arr {
		opts = append(opts, YAxisOption{
			Data: data,
		})
	}
	return opts
}

func (opt *YAxisOption) ToAxisOption(p *Painter) AxisOption {
	position := PositionLeft
	if opt.Position == PositionRight {
		position = PositionRight
	}
	theme := opt.Theme
	if theme == nil {
		theme = p.theme
	}
	axisOpt := AxisOption{
		Formatter:      opt.Formatter,
		Theme:          theme,
		Data:           opt.Data,
		Position:       position,
		FontSize:       opt.FontSize,
		StrokeWidth:    -1,
		Font:           opt.Font,
		FontColor:      opt.FontColor,
		BoundaryGap:    FalseFlag(),
		SplitLineShow:  true,
		SplitLineColor: theme.GetAxisSplitLineColor(),
		Show:           opt.Show,
		Unit:           opt.Unit,
	}
	if !opt.Color.IsZero() {
		axisOpt.FontColor = opt.Color
		axisOpt.StrokeColor = opt.Color
	}
	if opt.isCategoryAxis {
		axisOpt.BoundaryGap = TrueFlag()
		axisOpt.StrokeWidth = 1
		axisOpt.SplitLineShow = false
	}
	if opt.SplitLineShow != nil {
		axisOpt.SplitLineShow = *opt.SplitLineShow
	}
	return axisOpt
}

// NewLeftYAxis returns a left y axis renderer
func NewLeftYAxis(p *Painter, opt YAxisOption) *axisPainter {
	p = p.Child(PainterPaddingOption(Box{
		Bottom: defaultXAxisHeight,
	}))
	return NewAxisPainter(p, opt.ToAxisOption(p))
}

// NewRightYAxis returns a right y axis renderer
func NewRightYAxis(p *Painter, opt YAxisOption) *axisPainter {
	p = p.Child(PainterPaddingOption(Box{
		Bottom: defaultXAxisHeight,
	}))
	axisOpt := opt.ToAxisOption(p)
	axisOpt.Position = PositionRight
	axisOpt.SplitLineShow = false
	return NewAxisPainter(p, axisOpt)
}
