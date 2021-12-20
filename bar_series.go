// MIT License

// Copyright (c) 2021 Tree Xie

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

const defaultBarMargin = 10

type BarSeriesCustomStyle struct {
	PointIndex int
	Index      int
	Style      chart.Style
}

type BarSeries struct {
	BaseSeries
	Count int
	Index int
	// 间隔
	Margin int
	// 偏移量
	Offset int
	// 宽度
	BarWidth     int
	CustomStyles []BarSeriesCustomStyle
}

type barSeriesWidthValues struct {
	columnWidth  int
	columnMargin int
	margin       int
	barWidth     int
}

func (bs BarSeries) GetBarStyle(index, pointIndex int) chart.Style {
	// 指定样式
	for _, item := range bs.CustomStyles {
		if item.Index == index && item.PointIndex == pointIndex {
			return item.Style
		}
	}
	// 其它非指定样式
	return chart.Style{}
}

func (bs BarSeries) getWidthValues(width int) barSeriesWidthValues {
	columnWidth := width / bs.Len()
	// 块间隔
	columnMargin := columnWidth / 10
	minColumnMargin := 2
	if columnMargin < minColumnMargin {
		columnMargin = minColumnMargin
	}
	margin := bs.Margin
	if margin <= 0 {
		margin = defaultBarMargin
	}
	// 如果margin大于column margin
	if margin > columnMargin {
		margin = columnMargin
	}

	allBarMarginWidth := (bs.Count - 1) * margin
	barWidth := ((columnWidth - 2*columnMargin) - allBarMarginWidth) / bs.Count
	if bs.BarWidth > 0 && bs.BarWidth < barWidth {
		barWidth = bs.BarWidth
		// 重新计息columnMargin
		columnMargin = (columnWidth - allBarMarginWidth - (bs.Count * barWidth)) / 2
	}
	return barSeriesWidthValues{
		columnWidth:  columnWidth,
		columnMargin: columnMargin,
		margin:       margin,
		barWidth:     barWidth,
	}
}

func (bs BarSeries) Render(r chart.Renderer, canvasBox chart.Box, xrange, yrange chart.Range, defaults chart.Style) {
	if bs.Len() == 0 || bs.Count <= 0 {
		return
	}
	style := bs.Style.InheritFrom(defaults)
	style.FillColor = style.StrokeColor
	if !style.ShouldDrawStroke() {
		return
	}

	cb := canvasBox.Bottom
	cl := canvasBox.Left
	widthValues := bs.getWidthValues(canvasBox.Width())

	for i := 0; i < bs.Len(); i++ {
		vx, vy := bs.GetValues(i)
		customStyle := bs.GetBarStyle(bs.Index, i)
		cloneStyle := style
		if !customStyle.IsZero() {
			cloneStyle.FillColor = customStyle.FillColor
			cloneStyle.StrokeColor = customStyle.StrokeColor
		}

		x := cl + xrange.Translate(vx)
		// 由于bar是居中展示，因此需要往前移一个显示块
		x += (-widthValues.columnWidth + widthValues.columnMargin)
		// 计算是第几个bar，位置右偏
		x += bs.Index * (widthValues.margin + widthValues.barWidth)
		y := cb - yrange.Translate(vy)

		chart.Draw.Box(r, chart.Box{
			Left:   x,
			Top:    y,
			Right:  x + widthValues.barWidth,
			Bottom: canvasBox.Bottom - 1,
		}, cloneStyle)
	}
}
