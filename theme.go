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

import "github.com/wcharczuk/go-chart/v2/drawing"

const ThemeDark = "dark"
const ThemeLight = "light"

type Theme struct {
	mode string
}

func (t *Theme) IsDark() bool {
	return t.mode == ThemeDark
}

func (t *Theme) GetAxisStrokeColor() drawing.Color {
	if t.IsDark() {
		return drawing.Color{
			R: 185,
			G: 184,
			B: 206,
			A: 255,
		}
	}
	return drawing.Color{
		R: 110,
		G: 112,
		B: 121,
		A: 255,
	}
}

func (t *Theme) GetAxisSplitLineColor() drawing.Color {
	if t.IsDark() {
		return drawing.Color{
			R: 72,
			G: 71,
			B: 83,
			A: 255,
		}
	}
	return drawing.Color{
		R: 224,
		G: 230,
		B: 242,
		A: 255,
	}
}

func (t *Theme) GetSeriesColor(index int) drawing.Color {
	colors := t.GetSeriesColors()
	return colors[index%len(colors)]
}

func (t *Theme) GetSeriesColors() []drawing.Color {
	return []drawing.Color{
		{
			R: 84,
			G: 112,
			B: 198,
			A: 255,
		},
		{
			R: 145,
			G: 204,
			B: 117,
			A: 255,
		},
		{
			R: 250,
			G: 200,
			B: 88,
			A: 255,
		},
		{
			R: 238,
			G: 102,
			B: 102,
			A: 255,
		},
		{
			R: 115,
			G: 192,
			B: 222,
			A: 255,
		},
	}
}

func (t *Theme) GetBackgroundColor() drawing.Color {
	if t.IsDark() {
		return drawing.Color{
			R: 16,
			G: 12,
			B: 42,
			A: 255,
		}
	}
	return drawing.ColorWhite
}
