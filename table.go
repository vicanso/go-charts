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
	"errors"

	"github.com/wcharczuk/go-chart/v2"
)

type TableOption struct {
	// draw
	Draw *Draw
	// The width of table
	Width int
	// The header of table
	Header []string
	// The style of table
	Style        chart.Style
	ColumnWidths []float64
	// 是否仅测量高度
	measurement bool
}

var ErrTableColumnWidthInvalid = errors.New("table column width is invalid")

func tableDivideWidth(width, size int, columnWidths []float64) ([]int, error) {
	widths := make([]int, size)

	autoFillCount := size
	restWidth := width
	if len(columnWidths) != 0 {
		for index, v := range columnWidths {
			if v <= 0 {
				continue
			}
			autoFillCount--
			// 小于1的表示占比
			if v < 1 {
				widths[index] = int(v * float64(width))
			} else {
				widths[index] = int(v)
			}
			restWidth -= widths[index]
		}
	}
	if restWidth < 0 {
		return nil, ErrTableColumnWidthInvalid
	}
	// 填充其它未指定的宽度
	if autoFillCount > 0 {
		autoWidth := restWidth / autoFillCount
		for index, v := range widths {
			if v == 0 {
				widths[index] = autoWidth
			}
		}
	}
	widthSum := 0
	for _, v := range widths {
		widthSum += v
	}
	if widthSum > width {
		return nil, ErrTableColumnWidthInvalid
	}
	return widths, nil
}

func TableMeasure(opt TableOption) (chart.Box, error) {
	d, err := NewDraw(DrawOption{
		Width:  opt.Width,
		Height: 600,
	})
	if err != nil {
		return chart.BoxZero, err
	}
	opt.Draw = d
	opt.measurement = true
	return tableRender(opt)
}

func tableRender(opt TableOption) (chart.Box, error) {
	if opt.Draw == nil {
		return chart.BoxZero, errors.New("draw can not be nil")
	}
	if len(opt.Header) == 0 {
		return chart.BoxZero, errors.New("header can not be nil")
	}
	width := opt.Width
	if width == 0 {
		width = opt.Draw.Box.Width()
	}

	columnWidths, err := tableDivideWidth(width, len(opt.Header), opt.ColumnWidths)
	if err != nil {
		return chart.BoxZero, err
	}

	d := opt.Draw
	style := opt.Style
	y := 0
	x := 0

	headerMaxHeight := 0
	for index, text := range opt.Header {
		var box chart.Box
		w := columnWidths[index]
		y0 := y + int(opt.Style.FontSize)
		if opt.measurement {
			box = d.measureTextFit(text, x, y0, w, style)
		} else {
			box = d.textFit(text, x, y0, w, style)
		}
		if box.Height() > headerMaxHeight {
			headerMaxHeight = box.Height()
		}
		x += w
	}
	y += headerMaxHeight

	return chart.Box{
		Right:  width,
		Bottom: y,
	}, nil
}
