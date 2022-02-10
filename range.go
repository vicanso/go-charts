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

type Range struct {
	divideCount int
	Min         float64
	Max         float64
	Size        int
	Boundary    bool
}

func NewRange(min, max float64, divideCount int) Range {
	r := math.Abs(max - min)

	// 最小单位计算
	unit := 2
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
	return Range{
		Min:         min,
		Max:         max,
		divideCount: divideCount,
	}
}

func (r Range) Values() []string {
	offset := (r.Max - r.Min) / float64(r.divideCount)
	values := make([]string, 0)
	for i := 0; i <= r.divideCount; i++ {
		v := r.Min + float64(i)*offset
		value := commafWithDigits(v)
		values = append(values, value)
	}
	return values
}

func (r *Range) getHeight(value float64) int {
	v := (value - r.Min) / (r.Max - r.Min)
	return int(v * float64(r.Size))
}

func (r *Range) getRestHeight(value float64) int {
	return r.Size - r.getHeight(value)
}

func (r *Range) GetRange(index int) (float64, float64) {
	unit := float64(r.Size) / float64(r.divideCount)
	return unit * float64(index), unit * float64(index+1)
}

func (r *Range) getWidth(value float64) int {
	v := value / (r.Max - r.Min)
	// 移至居中
	if r.Boundary &&
		r.divideCount != 0 {
		v += 1 / float64(r.divideCount*2)
	}
	return int(v * float64(r.Size))
}
