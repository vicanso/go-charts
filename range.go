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
	"math"
)

const defaultAxisDivideCount = 6

type axisRange struct {
	p           *Painter
	divideCount int
	min         float64
	max         float64
	size        int
	boundary    bool
}

type AxisRangeOption struct {
	Painter *Painter
	// The min value of axis
	Min float64
	// The max value of axis
	Max float64
	// The size of axis
	Size int
	// Boundary gap
	Boundary bool
	// The count of divide
	DivideCount int
}

// NewRange returns a axis range
func NewRange(opt AxisRangeOption) axisRange {
	max := opt.Max
	min := opt.Min

	max += math.Abs(max * 0.1)
	min -= math.Abs(min * 0.1)
	divideCount := opt.DivideCount
	r := math.Abs(max - min)

	// 最小单位计算
	unit := 1
	if r > 5 {
		unit = 2
	}
	if r > 10 {
		unit = 4
	}
	if r > 30 {
		unit = 5
	}
	if r > 100 {
		unit = 10
	}
	if r > 200 {
		unit = 20
	}
	unit = int((r/float64(divideCount))/float64(unit))*unit + unit

	if min != 0 {
		isLessThanZero := min < 0
		min = float64(int(min/float64(unit)) * unit)
		// 如果是小于0，int的时候向上取整了，因此调整
		if min < 0 ||
			(isLessThanZero && min == 0) {
			min -= float64(unit)
		}
	}
	max = min + float64(unit*divideCount)
	expectMax := opt.Max * 2
	if max > expectMax {
		max = float64(ceilFloatToInt(expectMax))
	}
	return axisRange{
		p:           opt.Painter,
		divideCount: divideCount,
		min:         min,
		max:         max,
		size:        opt.Size,
		boundary:    opt.Boundary,
	}
}

// Values returns values of range
func (r axisRange) Values() []string {
	offset := (r.max - r.min) / float64(r.divideCount)
	values := make([]string, 0)
	formatter := commafWithDigits
	if r.p != nil && r.p.valueFormatter != nil {
		formatter = r.p.valueFormatter
	}
	for i := 0; i <= r.divideCount; i++ {
		v := r.min + float64(i)*offset
		value := formatter(v)
		values = append(values, value)
	}
	return values
}

func (r *axisRange) getHeight(value float64) int {
	if r.max <= r.min {
		return 0
	}
	v := (value - r.min) / (r.max - r.min)
	return int(v * float64(r.size))
}

func (r *axisRange) getRestHeight(value float64) int {
	return r.size - r.getHeight(value)
}

// GetRange returns a range of index
func (r *axisRange) GetRange(index int) (float64, float64) {
	unit := float64(r.size) / float64(r.divideCount)
	return unit * float64(index), unit * float64(index+1)
}

// AutoDivide divides the axis
func (r *axisRange) AutoDivide() []int {
	return autoDivide(r.size, r.divideCount)
}
