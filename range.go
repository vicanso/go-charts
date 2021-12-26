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
	"math"

	"github.com/wcharczuk/go-chart/v2"
)

type Range struct {
	TickPosition chart.TickPosition
	chart.ContinuousRange
}

func wrapRange(r chart.Range, tickPosition chart.TickPosition) chart.Range {
	xr, ok := r.(*chart.ContinuousRange)
	if !ok {
		return r
	}
	return &Range{
		TickPosition:    tickPosition,
		ContinuousRange: *xr,
	}
}

// Translate maps a given value into the ContinuousRange space.
func (r Range) Translate(value float64) int {
	v := r.ContinuousRange.Translate(value)
	if r.TickPosition == chart.TickPositionBetweenTicks {
		v -= int(float64(r.Domain) / (r.GetDelta() * 2))
	}
	return v
}

type HiddenRange struct {
	chart.ContinuousRange
}

func (r HiddenRange) GetDelta() float64 {
	return 0
}

// Y轴使用的continuous range
// min 与max只允许设置一次
// 如果是计算得出的max，增加20%的值并取整
type YContinuousRange struct {
	chart.ContinuousRange
}

func (m YContinuousRange) IsZero() bool {
	// 默认返回true，允许修改
	return true
}

func (m *YContinuousRange) SetMin(min float64) {
	// 如果已修改，则忽略
	if m.Min != -math.MaxFloat64 {
		return
	}
	m.Min = min
}

func (m *YContinuousRange) SetMax(max float64) {
	// 如果已修改，则忽略
	if m.Max != math.MaxFloat64 {
		return
	}
	// 此处为计算得来的最大值，放大20%
	v := int(max * 1.2)
	// TODO 是否要取整十整百
	m.Max = float64(v)
}
