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
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestTitleCustomize(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		title Title
		svg   string
	}{
		// 单行标题
		{
			title: Title{
				Text: "Hello World!",
				Style: chart.Style{
					FontColor: drawing.ColorBlack,
				},
			},
			svg: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"800\" height=\"600\">\\n<text x=\"50\" y=\"71\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:23.0px;font-family:'Roboto Medium',sans-serif\">Hello World!</text></svg>",
		},
		// 多行标题，靠右
		{
			title: Title{
				Text: "Hello World!\nHello World",
				Style: chart.Style{
					FontColor: drawing.ColorBlack,
				},
				Left: "right",
			},
			svg: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"800\" height=\"600\">\\n<text x=\"474\" y=\"71\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:23.0px;font-family:'Roboto Medium',sans-serif\">Hello World!</text><text x=\"477\" y=\"94\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:23.0px;font-family:'Roboto Medium',sans-serif\">Hello World</text></svg>",
		},
		// 标题居中
		{
			title: Title{
				Text: "Hello World!",
				Style: chart.Style{
					FontColor: drawing.ColorBlack,
				},
				Left: "center",
			},
			svg: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"800\" height=\"600\">\\n<text x=\"262\" y=\"71\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:23.0px;font-family:'Roboto Medium',sans-serif\">Hello World!</text></svg>",
		},
	}
	for _, tt := range tests {
		r, err := chart.SVG(800, 600)
		assert.Nil(err)
		fn := NewTitleCustomize(tt.title)
		fn(r, chart.NewBox(50, 50, 600, 400), chart.Style{
			Font: chart.StyleTextDefaults().Font,
		})
		buf := bytes.Buffer{}
		err = r.Save(&buf)
		assert.Nil(err)
		assert.Equal(tt.svg, buf.String())
	}
}
