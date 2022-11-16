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
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const ThemeDark = "dark"
const ThemeLight = "light"
const ThemeGrafana = "grafana"
const ThemeAnt = "ant"

type ColorPalette interface {
	IsDark() bool
	GetAxisStrokeColor() Color
	SetAxisStrokeColor(Color)
	GetAxisSplitLineColor() Color
	SetAxisSplitLineColor(Color)
	GetSeriesColor(int) Color
	SetSeriesColor([]Color)
	GetBackgroundColor() Color
	SetBackgroundColor(Color)
	GetTextColor() Color
	SetTextColor(Color)
	GetFontSize() float64
	SetFontSize(float64)
	GetFont() *truetype.Font
	SetFont(*truetype.Font)
}

type themeColorPalette struct {
	isDarkMode         bool
	axisStrokeColor    Color
	axisSplitLineColor Color
	backgroundColor    Color
	textColor          Color
	seriesColors       []Color
	fontSize           float64
	font               *truetype.Font
}

type ThemeOption struct {
	IsDarkMode         bool
	AxisStrokeColor    Color
	AxisSplitLineColor Color
	BackgroundColor    Color
	TextColor          Color
	SeriesColors       []Color
}

var palettes = map[string]*themeColorPalette{}

const defaultFontSize = 12.0

var defaultTheme ColorPalette

var defaultLightFontColor = drawing.Color{
	R: 70,
	G: 70,
	B: 70,
	A: 255,
}
var defaultDarkFontColor = drawing.Color{
	R: 238,
	G: 238,
	B: 238,
	A: 255,
}

func init() {
	echartSeriesColors := []Color{
		parseColor("#5470c6"),
		parseColor("#91cc75"),
		parseColor("#fac858"),
		parseColor("#ee6666"),
		parseColor("#73c0de"),
		parseColor("#3ba272"),
		parseColor("#fc8452"),
		parseColor("#9a60b4"),
		parseColor("#ea7ccc"),
	}
	grafanaSeriesColors := []Color{
		parseColor("#7EB26D"),
		parseColor("#EAB839"),
		parseColor("#6ED0E0"),
		parseColor("#EF843C"),
		parseColor("#E24D42"),
		parseColor("#1F78C1"),
		parseColor("#705DA0"),
		parseColor("#508642"),
	}
	antSeriesColors := []Color{
		parseColor("#5b8ff9"),
		parseColor("#5ad8a6"),
		parseColor("#5d7092"),
		parseColor("#f6bd16"),
		parseColor("#6f5ef9"),
		parseColor("#6dc8ec"),
		parseColor("#945fb9"),
		parseColor("#ff9845"),
	}
	AddTheme(
		ThemeDark,
		ThemeOption{
			IsDarkMode: true,
			AxisStrokeColor: Color{
				R: 185,
				G: 184,
				B: 206,
				A: 255,
			},
			AxisSplitLineColor: Color{
				R: 72,
				G: 71,
				B: 83,
				A: 255,
			},
			BackgroundColor: Color{
				R: 16,
				G: 12,
				B: 42,
				A: 255,
			},
			TextColor: Color{
				R: 238,
				G: 238,
				B: 238,
				A: 255,
			},
			SeriesColors: echartSeriesColors,
		},
	)

	AddTheme(
		ThemeLight,
		ThemeOption{
			IsDarkMode: false,
			AxisStrokeColor: Color{
				R: 110,
				G: 112,
				B: 121,
				A: 255,
			},
			AxisSplitLineColor: Color{
				R: 224,
				G: 230,
				B: 242,
				A: 255,
			},
			BackgroundColor: drawing.ColorWhite,
			TextColor: Color{
				R: 70,
				G: 70,
				B: 70,
				A: 255,
			},
			SeriesColors: echartSeriesColors,
		},
	)
	AddTheme(
		ThemeAnt,
		ThemeOption{
			IsDarkMode: false,
			AxisStrokeColor: Color{
				R: 110,
				G: 112,
				B: 121,
				A: 255,
			},
			AxisSplitLineColor: Color{
				R: 224,
				G: 230,
				B: 242,
				A: 255,
			},
			BackgroundColor: drawing.ColorWhite,
			TextColor: drawing.Color{
				R: 70,
				G: 70,
				B: 70,
				A: 255,
			},
			SeriesColors: antSeriesColors,
		},
	)
	AddTheme(
		ThemeGrafana,
		ThemeOption{
			IsDarkMode: true,
			AxisStrokeColor: Color{
				R: 185,
				G: 184,
				B: 206,
				A: 255,
			},
			AxisSplitLineColor: Color{
				R: 68,
				G: 67,
				B: 67,
				A: 255,
			},
			BackgroundColor: drawing.Color{
				R: 31,
				G: 29,
				B: 29,
				A: 255,
			},
			TextColor: Color{
				R: 216,
				G: 217,
				B: 218,
				A: 255,
			},
			SeriesColors: grafanaSeriesColors,
		},
	)
	SetDefaultTheme(ThemeLight)
}

// SetDefaultTheme sets default theme
func SetDefaultTheme(name string) {
	defaultTheme = NewTheme(name)
}

func AddTheme(name string, opt ThemeOption) {
	palettes[name] = &themeColorPalette{
		isDarkMode:         opt.IsDarkMode,
		axisStrokeColor:    opt.AxisStrokeColor,
		axisSplitLineColor: opt.AxisSplitLineColor,
		backgroundColor:    opt.BackgroundColor,
		textColor:          opt.TextColor,
		seriesColors:       opt.SeriesColors,
	}
}

func NewTheme(name string) ColorPalette {
	p, ok := palettes[name]
	if !ok {
		p = palettes[ThemeLight]
	}
	clone := *p
	return &clone
}

func (t *themeColorPalette) IsDark() bool {
	return t.isDarkMode
}

func (t *themeColorPalette) GetAxisStrokeColor() Color {
	return t.axisStrokeColor
}

func (t *themeColorPalette) SetAxisStrokeColor(c Color) {
	t.axisStrokeColor = c
}

func (t *themeColorPalette) GetAxisSplitLineColor() Color {
	return t.axisSplitLineColor
}

func (t *themeColorPalette) SetAxisSplitLineColor(c Color) {
	t.axisSplitLineColor = c
}

func (t *themeColorPalette) GetSeriesColor(index int) Color {
	colors := t.seriesColors
	return colors[index%len(colors)]
}
func (t *themeColorPalette) SetSeriesColor(colors []Color) {
	t.seriesColors = colors
}

func (t *themeColorPalette) GetBackgroundColor() Color {
	return t.backgroundColor
}

func (t *themeColorPalette) SetBackgroundColor(c Color) {
	t.backgroundColor = c
}

func (t *themeColorPalette) GetTextColor() Color {
	return t.textColor
}

func (t *themeColorPalette) SetTextColor(c Color) {
	t.textColor = c
}

func (t *themeColorPalette) GetFontSize() float64 {
	if t.fontSize != 0 {
		return t.fontSize
	}
	return defaultFontSize
}

func (t *themeColorPalette) SetFontSize(fontSize float64) {
	t.fontSize = fontSize
}

func (t *themeColorPalette) GetFont() *truetype.Font {
	if t.font != nil {
		return t.font
	}
	f, _ := GetDefaultFont()
	return f
}

func (t *themeColorPalette) SetFont(f *truetype.Font) {
	t.font = f
}
