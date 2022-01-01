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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
)

func TestRange(t *testing.T) {
	assert := assert.New(t)

	r := Range{
		ContinuousRange: chart.ContinuousRange{
			Min:    0,
			Max:    5,
			Domain: 500,
		},
	}

	assert.Equal(100, r.Translate(1))

	r.TickPosition = chart.TickPositionBetweenTicks
	assert.Equal(50, r.Translate(1))
}

func TestHiddenRange(t *testing.T) {
	assert := assert.New(t)
	r := HiddenRange{}

	assert.Equal(float64(0), r.GetDelta())
}

func TestYContinuousRange(t *testing.T) {
	assert := assert.New(t)
	r := YContinuousRange{}
	r.Min = -math.MaxFloat64
	r.Max = math.MaxFloat64

	assert.True(r.IsZero())

	r.SetMin(1.0)
	assert.Equal(1.0, r.GetMin())
	// 再次设置无效
	r.SetMin(2.0)
	assert.Equal(1.0, r.GetMin())

	r.SetMax(5.0)
	// *1.2
	assert.Equal(6.0, r.GetMax())
	// 再次设置无效
	r.SetMax(10.0)
	assert.Equal(6.0, r.GetMax())
}
