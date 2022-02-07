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

func TestSplitTitleText(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{
		"a",
		"b",
	}, splitTitleText("a\nb"))
	assert.Equal([]string{
		"a",
	}, splitTitleText("a\n "))
}

func TestDrawTitle(t *testing.T) {
	assert := assert.New(t)

	newOption := func() *TitleOption {
		f, _ := chart.GetDefaultFont()
		return &TitleOption{
			Text:    "title\nHello",
			Subtext: "subtitle\nWorld!",
			Style: chart.Style{
				FontSize:  14,
				Font:      f,
				FontColor: drawing.ColorBlack,
			},
			SubtextStyle: chart.Style{
				FontSize:  10,
				Font:      f,
				FontColor: drawing.ColorBlue,
			},
		}
	}
	newDraw := func() *Draw {
		d, _ := NewDraw(DrawOption{
			Width:  400,
			Height: 300,
		})
		return d
	}

	tests := []struct {
		newDraw   func() *Draw
		newOption func() *TitleOption
		result    string
		box       chart.Box
	}{
		{
			newDraw:   newDraw,
			newOption: newOption,
			result:    "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<text x=\"6\" y=\"0\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:17.9px;font-family:'Roboto Medium',sans-serif\">title</text><text x=\"0\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:17.9px;font-family:'Roboto Medium',sans-serif\">Hello</text><text x=\"0\" y=\"34\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">subtitle</text><text x=\"3\" y=\"46\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">World!</text></svg>",
			box: chart.Box{
				Right:  43,
				Bottom: 58,
			},
		},
		{
			newDraw: newDraw,
			newOption: func() *TitleOption {
				opt := newOption()
				opt.Left = PositionRight
				opt.Top = "50"
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<text x=\"363\" y=\"50\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:17.9px;font-family:'Roboto Medium',sans-serif\">title</text><text x=\"357\" y=\"67\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:17.9px;font-family:'Roboto Medium',sans-serif\">Hello</text><text x=\"357\" y=\"84\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">subtitle</text><text x=\"360\" y=\"96\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">World!</text></svg>",
			box: chart.Box{
				Right:  400,
				Bottom: 108,
			},
		},
		{
			newDraw: newDraw,
			newOption: func() *TitleOption {
				opt := newOption()
				opt.Left = PositionCenter
				opt.Top = "10"
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<text x=\"185\" y=\"10\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:17.9px;font-family:'Roboto Medium',sans-serif\">title</text><text x=\"179\" y=\"27\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:17.9px;font-family:'Roboto Medium',sans-serif\">Hello</text><text x=\"179\" y=\"44\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">subtitle</text><text x=\"182\" y=\"56\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">World!</text></svg>",
			box: chart.Box{
				Right:  222,
				Bottom: 68,
			},
		},
		{
			newDraw: newDraw,
			newOption: func() *TitleOption {
				opt := newOption()
				opt.Left = "10%"
				opt.Top = "10"
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<text x=\"46\" y=\"10\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:17.9px;font-family:'Roboto Medium',sans-serif\">title</text><text x=\"40\" y=\"27\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:17.9px;font-family:'Roboto Medium',sans-serif\">Hello</text><text x=\"40\" y=\"44\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">subtitle</text><text x=\"43\" y=\"56\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,255,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">World!</text></svg>",
			box: chart.Box{
				Right:  83,
				Bottom: 68,
			},
		},
	}
	for _, tt := range tests {
		d := tt.newDraw()
		o := tt.newOption()
		b, err := drawTitle(d, o)
		assert.Nil(err)
		assert.Equal(tt.box, b)
		data, err := d.Bytes()
		assert.Nil(err)
		assert.NotEmpty(data)
		assert.Equal(tt.result, string(data))
	}
}
