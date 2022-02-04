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
	"strconv"
	"strings"

	"github.com/wcharczuk/go-chart/v2"
)

func TrueFlag() *bool {
	t := true
	return &t
}

func FalseFlag() *bool {
	f := false
	return &f
}

func getDefaultInt(value, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}

func autoDivide(max, size int) []int {
	unit := max / size

	rest := max - unit*size
	values := make([]int, size+1)
	value := 0
	for i := 0; i < size; i++ {
		values[i] = value
		if i < rest {
			value++
		}
		value += unit
	}
	values[size] = max
	return values
}
func maxInt(values ...int) int {
	result := 0
	for _, v := range values {
		if v > result {
			result = v
		}
	}
	return result
}

// measureTextMaxWidthHeight returns maxWidth and maxHeight of text list
func measureTextMaxWidthHeight(textList []string, r chart.Renderer) (int, int) {
	maxWidth := 0
	maxHeight := 0
	for _, text := range textList {
		box := r.MeasureText(text)
		maxWidth = maxInt(maxWidth, box.Width())
		maxHeight = maxInt(maxHeight, box.Height())
	}
	return maxWidth, maxHeight
}

func reverseStringSlice(stringList []string) {
	for i, j := 0, len(stringList)-1; i < j; i, j = i+1, j-1 {
		stringList[i], stringList[j] = stringList[j], stringList[i]
	}
}

func reverseIntSlice(intList []int) {
	for i, j := 0, len(intList)-1; i < j; i, j = i+1, j-1 {
		intList[i], intList[j] = intList[j], intList[i]
	}
}

func convertPercent(value string) float64 {
	if !strings.HasSuffix(value, "%") {
		return -1
	}
	v, err := strconv.Atoi(strings.ReplaceAll(value, "%", ""))
	if err != nil {
		return -1
	}
	return float64(v) / 100
}

func isFalse(flag *bool) bool {
	if flag != nil && !*flag {
		return true
	}
	return false
}
