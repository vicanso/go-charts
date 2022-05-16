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
	"regexp"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TrueFlag() *bool {
	t := true
	return &t
}

func FalseFlag() *bool {
	f := false
	return &f
}

func ceilFloatToInt(value float64) int {
	i := int(value)
	if value == float64(i) {
		return i
	}
	return i + 1
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

// measureTextMaxWidthHeight returns maxWidth and maxHeight of text list
func measureTextMaxWidthHeight(textList []string, p *Painter) (int, int) {
	maxWidth := 0
	maxHeight := 0
	for _, text := range textList {
		box := p.MeasureText(text)
		maxWidth = chart.MaxInt(maxWidth, box.Width())
		maxHeight = chart.MaxInt(maxHeight, box.Height())
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

func NewFloatPoint(f float64) *float64 {
	v := f
	return &v
}
func commafWithDigits(value float64) string {
	decimals := 2
	m := float64(1000 * 1000)
	if value >= m {
		return humanize.CommafWithDigits(value/m, decimals) + "M"
	}
	k := float64(1000)
	if value >= k {
		return humanize.CommafWithDigits(value/k, decimals) + "k"
	}
	return humanize.CommafWithDigits(value, decimals)
}

func parseColor(color string) drawing.Color {
	c := drawing.Color{}
	if color == "" {
		return c
	}
	if strings.HasPrefix(color, "#") {
		return drawing.ColorFromHex(color[1:])
	}
	reg := regexp.MustCompile(`\((\S+)\)`)
	result := reg.FindAllStringSubmatch(color, 1)
	if len(result) == 0 || len(result[0]) != 2 {
		return c
	}
	arr := strings.Split(result[0][1], ",")
	if len(arr) < 3 {
		return c
	}
	// 设置默认为255
	c.A = 255
	for index, v := range arr {
		value, _ := strconv.Atoi(strings.TrimSpace(v))
		ui8 := uint8(value)
		switch index {
		case 0:
			c.R = ui8
		case 1:
			c.G = ui8
		case 2:
			c.B = ui8
		default:
			c.A = ui8
		}
	}
	return c
}

const defaultRadiusPercent = 0.4

func getRadius(diameter float64, radiusValue string) float64 {
	var radius float64
	if len(radiusValue) != 0 {
		v := convertPercent(radiusValue)
		if v != -1 {
			radius = float64(diameter) * v
		} else {
			radius, _ = strconv.ParseFloat(radiusValue, 64)
		}
	}
	if radius <= 0 {
		radius = float64(diameter) * defaultRadiusPercent
	}
	return radius
}

func getPolygonPointAngles(sides int) []float64 {
	angles := make([]float64, sides)
	for i := 0; i < sides; i++ {
		angle := 2*math.Pi/float64(sides)*float64(i) - (math.Pi / 2)
		angles[i] = angle
	}
	return angles
}

func getPolygonPoint(center Point, radius, angle float64) Point {
	x := center.X + int(radius*math.Cos(angle))
	y := center.Y + int(radius*math.Sin(angle))
	return Point{
		X: x,
		Y: y,
	}
}

func getPolygonPoints(center Point, radius float64, sides int) []Point {
	points := make([]Point, sides)
	for i, angle := range getPolygonPointAngles(sides) {
		points[i] = getPolygonPoint(center, radius, angle)
	}
	return points
}
