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

type Renderer interface {
	Render() (Box, error)
}

type defaultRenderOption struct {
	Theme      ColorPalette
	Padding    Box
	SeriesList SeriesList
	// The y axis option
	YAxisOptions []YAxisOption
	// The x axis option
	XAxis XAxisOption
}

type defaultRenderResult struct {
	axisRanges map[int]axisRange
	p          *Painter
}

func defaultRender(p *Painter, opt defaultRenderOption) (*defaultRenderResult, error) {
	p.SetBackground(p.Width(), p.Height(), opt.Theme.GetBackgroundColor())
	if !opt.Padding.IsZero() {
		p = p.Child(PainterPaddingOption(opt.Padding))
	}
	result := defaultRenderResult{
		axisRanges: make(map[int]axisRange),
	}

	// 计算图表对应的轴有哪些
	axisIndexList := make([]int, 0)
	for _, series := range opt.SeriesList {
		if containsInt(axisIndexList, series.AxisIndex) {
			continue
		}
		axisIndexList = append(axisIndexList, series.index)
	}
	// 高度需要减去x轴的高度
	rangeHeight := p.Height() - defaultXAxisHeight
	rangeWidth := 0

	// 计算对应的axis range
	for _, index := range axisIndexList {
		max, min := opt.SeriesList.GetMaxMin(index)
		r := NewRange(AxisRangeOption{
			Min: min,
			Max: max,
			// 高度需要减去x轴的高度
			Size: rangeHeight,
			// 分隔数量
			DivideCount: defaultAxisDivideCount,
		})
		result.axisRanges[index] = r
		yAxisOption := YAxisOption{}
		if len(opt.YAxisOptions) > index {
			yAxisOption = opt.YAxisOptions[index]
		}
		if yAxisOption.Theme == nil {
			yAxisOption.Theme = opt.Theme
		}
		yAxisOption.Data = r.Values()
		reverseStringSlice(yAxisOption.Data)
		// TODO生成其它位置既yAxis
		yAxis := NewLeftYAxis(p, yAxisOption)
		yAxisBox, err := yAxis.Render()
		if err != nil {
			return nil, err
		}
		rangeWidth += yAxisBox.Width()
	}

	if opt.XAxis.Theme == nil {
		opt.XAxis.Theme = opt.Theme
	}
	xAxis := NewBottomXAxis(p.Child(PainterPaddingOption(Box{
		Left: rangeWidth,
	})), opt.XAxis)
	_, err := xAxis.Render()
	if err != nil {
		return nil, err
	}

	// // 生成Y轴
	// for _, yAxisOption := range opt.YAxisOptions {

	// }

	result.p = p.Child(PainterPaddingOption(Box{
		Bottom: rangeHeight,
		Left:   rangeWidth,
	}))
	return &result, nil
}

func doRender(renderers ...Renderer) error {
	for _, r := range renderers {
		_, err := r.Render()
		if err != nil {
			return err
		}
	}
	return nil
}
