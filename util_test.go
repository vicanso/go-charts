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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
)

func TestGetDefaultInt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, getDefaultInt(0, 1))
	assert.Equal(10, getDefaultInt(10, 1))
}

func TestAutoDivide(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{
		0,
		86,
		172,
		258,
		344,
		430,
		515,
		600,
	}, autoDivide(600, 7))
}

func TestMaxInt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, maxInt(1, 3, 5, 2))
}

func TestMeasureTextMaxWidthHeight(t *testing.T) {
	assert := assert.New(t)
	r, err := chart.SVG(400, 300)
	assert.Nil(err)
	style := chart.Style{
		FontSize: 10,
	}
	style.Font, _ = chart.GetDefaultFont()
	style.WriteToRenderer(r)

	maxWidth, maxHeight := measureTextMaxWidthHeight([]string{
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
		"Sun",
	}, r)
	assert.Equal(26, maxWidth)
	assert.Equal(12, maxHeight)
}

func TestReverseSlice(t *testing.T) {
	assert := assert.New(t)

	arr := []string{
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
		"Sun",
	}
	reverseStringSlice(arr)
	assert.Equal([]string{
		"Sun",
		"Sat",
		"Fri",
		"Thu",
		"Wed",
		"Tue",
		"Mon",
	}, arr)

	numbers := []int{
		1,
		3,
		5,
		7,
		9,
	}
	reverseIntSlice(numbers)
	assert.Equal([]int{
		9,
		7,
		5,
		3,
		1,
	}, numbers)
}