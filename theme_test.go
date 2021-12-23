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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestThemeColors(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AxisColorDark, getAxisColor(ThemeDark))
	assert.Equal(AxisColorLight, getAxisColor(""))

	assert.Equal(GridColorDark, getGridColor(ThemeDark))
	assert.Equal(GridColorLight, getGridColor(""))

	assert.Equal(BackgroundColorDark, getBackgroundColor(ThemeDark))
	assert.Equal(chart.DefaultBackgroundColor, getBackgroundColor(""))

	assert.Equal(TextColorDark, getTextColor(ThemeDark))
	assert.Equal(chart.DefaultTextColor, getTextColor(""))
}

func TestThemeColorPalette(t *testing.T) {
	assert := assert.New(t)

	dark := ThemeColorPalette{
		Theme: ThemeDark,
	}
	assert.Equal(BackgroundColorDark, dark.BackgroundColor())
	assert.Equal(chart.DefaultBackgroundStrokeColor, dark.BackgroundStrokeColor())
	assert.Equal(BackgroundColorDark, dark.CanvasColor())
	assert.Equal(chart.DefaultCanvasStrokeColor, dark.CanvasStrokeColor())
	assert.Equal(BackgroundColorDark, dark.AxisStrokeColor())
	assert.Equal(TextColorDark, dark.TextColor())
	// series 使用统一的color
	assert.Equal(SeriesColorsLight[0], dark.GetSeriesColor(0))

	light := ThemeColorPalette{}
	assert.Equal(chart.DefaultBackgroundColor, light.BackgroundColor())
	assert.Equal(chart.DefaultBackgroundStrokeColor, light.BackgroundStrokeColor())
	assert.Equal(chart.DefaultCanvasColor, light.CanvasColor())
	assert.Equal(chart.DefaultCanvasStrokeColor, light.CanvasStrokeColor())
	assert.Equal(chart.DefaultAxisColor, light.AxisStrokeColor())
	assert.Equal(chart.DefaultTextColor, light.TextColor())
	// series 使用统一的color
	assert.Equal(SeriesColorsLight[0], light.GetSeriesColor(0))
}

func TestPieThemeColorPalette(t *testing.T) {
	assert := assert.New(t)
	p := PieThemeColorPalette{}

	// pie无认哪种theme，文本的颜色都一样
	assert.Equal(chart.DefaultTextColor, p.TextColor())
	p.Theme = ThemeDark
	assert.Equal(chart.DefaultTextColor, p.TextColor())
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
