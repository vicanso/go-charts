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

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type tableChart struct {
	p   *Painter
	opt *TableChartOption
}

func NewTableChart(p *Painter, opt TableChartOption) *tableChart {
	if opt.Theme == nil {
		opt.Theme = defaultTheme
	}
	return &tableChart{
		p:   p,
		opt: &opt,
	}
}

type TableChartOption struct {
	// The output type
	Type string
	// The width of table
	Width int
	// The height of table
	Height int
	// The theme
	Theme ColorPalette
	// The padding of table cell
	Padding Box
	// The header data of table
	Header []string
	// The data of table
	Data [][]string
	// The span of table column
	Spans []int
	// The font size of table
	FontSize float64
	Font     *truetype.Font
	// The font color of table
	FontColor Color
	// The background color of header
	HeaderBackgroundColor Color
	// The header font color
	HeaderFontColor Color
	// The background color of row
	RowBackgroundColors []Color
	// The background color
	BackgroundColor Color
}

var defaultTableHeaderColor = Color{
	R: 34,
	G: 34,
	B: 34,
	A: 255,
}
var defaultTableRowColors = []Color{
	drawing.ColorWhite,
	{
		R: 242,
		G: 242,
		B: 242,
		A: 255,
	},
}
var defaultTablePadding = Box{
	Left:   10,
	Top:    10,
	Right:  10,
	Bottom: 10,
}

type renderInfo struct {
	Width        int
	Height       int
	HeaderHeight int
	RowHeights   []int
}

func (c *tableChart) render() (*renderInfo, error) {
	info := renderInfo{
		RowHeights: make([]int, 0),
	}
	p := c.p
	opt := c.opt
	if len(opt.Header) == 0 {
		return nil, errors.New("header can not be nil")
	}
	theme := opt.Theme
	if theme == nil {
		theme = p.theme
	}
	fontSize := opt.FontSize
	if fontSize == 0 {
		fontSize = 12
	}
	fontColor := opt.FontColor
	if fontColor.IsZero() {
		fontColor = theme.GetTextColor()
	}
	font := opt.Font
	if font == nil {
		font = theme.GetFont()
	}
	headerFontColor := opt.HeaderFontColor
	if opt.HeaderFontColor.IsZero() {
		headerFontColor = drawing.ColorWhite
	}

	spans := opt.Spans
	if len(spans) != 0 && len(spans) != len(opt.Header) {
		newSpans := make([]int, len(opt.Header))
		for index := range opt.Header {
			if len(spans) < index {
				newSpans[index] = 1
			} else {
				newSpans[index] = spans[index]
			}
		}
		spans = newSpans
	}

	values := autoDivideSpans(p.Width(), len(opt.Header)+1, spans)
	height := 0
	textStyle := Style{
		FontSize:  fontSize,
		FontColor: headerFontColor,
		FillColor: headerFontColor,
		Font:      font,
	}
	p.SetStyle(textStyle)

	headerHeight := 0
	padding := opt.Padding
	if padding.IsZero() {
		padding = defaultTablePadding
	}

	renderTableCells := func(textList []string, currentHeight int, cellPadding Box) int {
		cellMaxHeight := 0
		paddingHeight := cellPadding.Top + cellPadding.Bottom
		paddingWidth := cellPadding.Left + cellPadding.Right
		for index, text := range textList {
			x := values[index]
			y := currentHeight + cellPadding.Top
			width := values[index+1] - x
			x += cellPadding.Left
			width -= paddingWidth
			box := p.TextFit(text, x, y+int(fontSize), width)
			if box.Height()+paddingHeight > cellMaxHeight {
				cellMaxHeight = box.Height() + paddingHeight
			}
		}
		return cellMaxHeight
	}

	headerHeight = renderTableCells(opt.Header, height, padding)
	height += headerHeight
	info.HeaderHeight = headerHeight

	textStyle.FontColor = fontColor
	textStyle.FillColor = fontColor
	p.SetStyle(textStyle)
	for _, textList := range opt.Data {
		cellHeight := renderTableCells(textList, height, padding)
		info.RowHeights = append(info.RowHeights, cellHeight)
		height += cellHeight
	}

	info.Width = p.Width()
	info.Height = height
	return &info, nil
}

func (c *tableChart) Render() (Box, error) {
	p := c.p
	opt := c.opt
	if !opt.BackgroundColor.IsZero() {
		p.SetBackground(p.Width(), p.Height(), opt.BackgroundColor)
	}
	headerBGColor := opt.HeaderBackgroundColor
	if headerBGColor.IsZero() {
		headerBGColor = defaultTableHeaderColor
	}

	r := p.render
	newRender, err := chart.SVG(p.Width(), 100)
	if err != nil {
		return BoxZero, err
	}
	p.render = newRender
	info, err := c.render()
	if err != nil {
		return BoxZero, err
	}
	p.render = r
	// 如果设置表头背景色
	p.SetBackground(info.Width, info.HeaderHeight, headerBGColor, true)
	currentHeight := info.HeaderHeight
	rowColors := opt.RowBackgroundColors
	if len(rowColors) == 0 {
		rowColors = defaultTableRowColors
	}
	for index, h := range info.RowHeights {
		color := rowColors[index%len(rowColors)]
		child := p.Child(PainterPaddingOption(Box{
			Top: currentHeight,
		}))
		child.SetBackground(p.Width(), h, color, true)
		currentHeight += h
	}
	_, err = c.render()
	if err != nil {
		return BoxZero, err
	}

	return Box{
		Right:  info.Width,
		Bottom: info.Height,
	}, nil
}
