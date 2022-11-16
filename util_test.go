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
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestGetDefaultInt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, getDefaultInt(0, 1))
	assert.Equal(10, getDefaultInt(10, 1))
}

func TestCeilFloatToInt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, ceilFloatToInt(0.8))
	assert.Equal(1, ceilFloatToInt(1.0))
	assert.Equal(2, ceilFloatToInt(1.2))
}

func TestCommafWithDigits(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("1.2", commafWithDigits(1.2))
	assert.Equal("1.21", commafWithDigits(1.21231))

	assert.Equal("1.20k", commafWithDigits(1200.121))
	assert.Equal("1.20M", commafWithDigits(1200000.121))
}

func TestAutoDivide(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{
		0,
		85,
		171,
		257,
		342,
		428,
		514,
		600,
	}, autoDivide(600, 7))
}

func TestGetRadius(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(50.0, getRadius(100, "50%"))
	assert.Equal(30.0, getRadius(100, "30"))
	assert.Equal(40.0, getRadius(100, ""))
}

func TestMeasureTextMaxWidthHeight(t *testing.T) {
	assert := assert.New(t)
	p, err := NewPainter(PainterOptions{
		Width:  400,
		Height: 300,
	})
	assert.Nil(err)
	style := chart.Style{
		FontSize: 10,
	}
	p.SetStyle(style)

	maxWidth, maxHeight := measureTextMaxWidthHeight([]string{
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
		"Sun",
	}, p)
	assert.Equal(31, maxWidth)
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

func TestConvertPercent(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(-1.0, convertPercent("1"))
	assert.Equal(-1.0, convertPercent("a%"))
	assert.Equal(0.1, convertPercent("10%"))
}

func TestParseColor(t *testing.T) {
	assert := assert.New(t)

	c := parseColor("")
	assert.True(c.IsZero())

	c = parseColor("#333")
	assert.Equal(drawing.Color{
		R: 51,
		G: 51,
		B: 51,
		A: 255,
	}, c)

	c = parseColor("#313233")
	assert.Equal(drawing.Color{
		R: 49,
		G: 50,
		B: 51,
		A: 255,
	}, c)

	c = parseColor("rgb(31,32,33)")
	assert.Equal(drawing.Color{
		R: 31,
		G: 32,
		B: 33,
		A: 255,
	}, c)

	c = parseColor("rgba(50,51,52,250)")
	assert.Equal(drawing.Color{
		R: 50,
		G: 51,
		B: 52,
		A: 250,
	}, c)
}

func TestIsLightColor(t *testing.T) {
	assert := assert.New(t)

	assert.True(isLightColor(drawing.Color{
		R: 255,
		G: 255,
		B: 255,
	}))
	assert.True(isLightColor(drawing.Color{
		R: 145,
		G: 204,
		B: 117,
	}))

	assert.False(isLightColor(drawing.Color{
		R: 88,
		G: 112,
		B: 198,
	}))

	assert.False(isLightColor(drawing.Color{
		R: 0,
		G: 0,
		B: 0,
	}))
	assert.False(isLightColor(drawing.Color{
		R: 16,
		G: 12,
		B: 42,
	}))
}
