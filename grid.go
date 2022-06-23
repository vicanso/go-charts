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

type gridPainter struct {
	p   *Painter
	opt *GridPainterOption
}

type GridPainterOption struct {
	// The stroke width
	StrokeWidth float64
	// The stroke color
	StrokeColor Color
	// The spans of column
	ColumnSpans []int
	// The column of grid
	Column int
	// The row of grid
	Row int
	// Ignore first row
	IgnoreFirstRow bool
	// Ignore last row
	IgnoreLastRow bool
	// Ignore first column
	IgnoreFirstColumn bool
	// Ignore last column
	IgnoreLastColumn bool
}

// NewGridPainter returns new a grid renderer
func NewGridPainter(p *Painter, opt GridPainterOption) *gridPainter {
	return &gridPainter{
		p:   p,
		opt: &opt,
	}
}

func (g *gridPainter) Render() (Box, error) {
	opt := g.opt
	ignoreColumnLines := make([]int, 0)
	if opt.IgnoreFirstColumn {
		ignoreColumnLines = append(ignoreColumnLines, 0)
	}
	if opt.IgnoreLastColumn {
		ignoreColumnLines = append(ignoreColumnLines, opt.Column)
	}
	ignoreRowLines := make([]int, 0)
	if opt.IgnoreFirstRow {
		ignoreRowLines = append(ignoreRowLines, 0)
	}
	if opt.IgnoreLastRow {
		ignoreRowLines = append(ignoreRowLines, opt.Row)
	}
	strokeWidth := opt.StrokeWidth
	if strokeWidth <= 0 {
		strokeWidth = 1
	}

	g.p.SetDrawingStyle(Style{
		StrokeWidth: strokeWidth,
		StrokeColor: opt.StrokeColor,
	})
	g.p.Grid(GridOption{
		Column:            opt.Column,
		ColumnSpans:       opt.ColumnSpans,
		Row:               opt.Row,
		IgnoreColumnLines: ignoreColumnLines,
		IgnoreRowLines:    ignoreRowLines,
	})
	return g.p.box, nil
}
