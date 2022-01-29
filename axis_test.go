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

	"github.com/golang/freetype/truetype"
	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestAxisStyle(t *testing.T) {
	assert := assert.New(t)

	as := AxisStyle{}

	assert.Equal(8, as.GetLabelMargin())
	as.LabelMargin = 10
	assert.Equal(10, as.GetLabelMargin())

	assert.Equal(5, as.GetTickLength())
	as.TickLength = 6
	assert.Equal(6, as.GetTickLength())

	f := &truetype.Font{}
	style := as.Style(f)
	assert.Equal(float64(10), style.FontSize)
	assert.Equal(f, style.Font)
}

func TestAxisDataList(t *testing.T) {
	assert := assert.New(t)

	textList := []string{
		"a",
		"b",
	}
	data := NewAxisDataListFromStringList(textList)
	assert.Equal(textList, data.TextList())
}

func TestAxis(t *testing.T) {
	assert := assert.New(t)

	data := NewAxisDataListFromStringList([]string{
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
		"Sun",
	})
	getDefaultOption := func() *axisOption {
		return &axisOption{
			data: &data,
			style: &AxisStyle{
				StrokeColor:    drawing.ColorBlack,
				StrokeWidth:    1,
				FontColor:      drawing.ColorBlack,
				Show:           TrueFlag(),
				TickShow:       TrueFlag(),
				SplitLineShow:  true,
				SplitLineColor: drawing.ColorBlack.WithAlpha(60),
			},
		}
	}
	tests := []struct {
		newOption func() *axisOption
		result    string
	}{
		// 文本按起始位置展示
		// axis位于bottom
		{
			newOption: func() *axisOption {
				opt := getDefaultOption()
				opt.style.BoundaryGap = FalseFlag()
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 5 270\nL 395 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 5 270\nL 5 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 70 270\nL 70 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 135 270\nL 135 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 200 270\nL 200 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 265 270\nL 265 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 330 270\nL 330 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 395 270\nL 395 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 5 5\nL 5 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 70 5\nL 70 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 135 5\nL 135 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 200 5\nL 200 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 265 5\nL 265 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 330 5\nL 330 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 395 5\nL 395 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><text x=\"-8\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><text x=\"59\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><text x=\"122\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text><text x=\"189\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Thu</text><text x=\"257\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Fri</text><text x=\"320\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sat</text><text x=\"384\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sun</text></svg>",
		},
		// 文本居中展示
		// axis位于bottom
		{
			newOption: func() *axisOption {
				opt := getDefaultOption()
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 5 270\nL 395 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 5 270\nL 5 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 61 270\nL 61 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 117 270\nL 117 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 173 270\nL 173 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 229 270\nL 229 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 285 270\nL 285 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 340 270\nL 340 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 395 270\nL 395 275\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 5 5\nL 5 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 61 5\nL 61 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 117 5\nL 117 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 173 5\nL 173 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 229 5\nL 229 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 285 5\nL 285 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 340 5\nL 340 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 395 5\nL 395 270\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><text x=\"20\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><text x=\"78\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><text x=\"132\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text><text x=\"190\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Thu</text><text x=\"249\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Fri</text><text x=\"303\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sat</text><text x=\"356\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sun</text></svg>",
		},
		// 文本按起始位置展示
		// axis位于top
		{
			newOption: func() *axisOption {
				opt := getDefaultOption()
				opt.style.Position = PositionTop
				opt.style.BoundaryGap = FalseFlag()
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 5 25\nL 395 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 5 20\nL 5 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 70 20\nL 70 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 135 20\nL 135 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 200 20\nL 200 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 265 20\nL 265 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 330 20\nL 330 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 395 20\nL 395 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 5 25\nL 5 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 70 25\nL 70 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 135 25\nL 135 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 200 25\nL 200 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 265 25\nL 265 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 330 25\nL 330 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 395 25\nL 395 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><text x=\"-8\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><text x=\"59\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><text x=\"122\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text><text x=\"189\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Thu</text><text x=\"257\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Fri</text><text x=\"320\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sat</text><text x=\"384\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sun</text></svg>",
		},
		// 文本居中展示
		// axis位于top
		{
			newOption: func() *axisOption {
				opt := getDefaultOption()
				opt.style.Position = PositionTop
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 5 25\nL 395 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 5 20\nL 5 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 61 20\nL 61 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 117 20\nL 117 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 173 20\nL 173 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 229 20\nL 229 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 285 20\nL 285 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 340 20\nL 340 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 395 20\nL 395 25\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 5 25\nL 5 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 61 25\nL 61 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 117 25\nL 117 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 173 25\nL 173 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 229 25\nL 229 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 285 25\nL 285 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 340 25\nL 340 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 395 25\nL 395 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><text x=\"20\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><text x=\"78\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><text x=\"132\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text><text x=\"190\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Thu</text><text x=\"249\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Fri</text><text x=\"303\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sat</text><text x=\"356\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sun</text></svg>",
		},
		// 文本按起始位置展示
		// axis位于left
		{
			newOption: func() *axisOption {
				opt := getDefaultOption()
				opt.style.Position = PositionLeft
				opt.style.BoundaryGap = FalseFlag()
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 44 5\nL 44 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 5\nL 44 5\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 54\nL 44 54\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 103\nL 44 103\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 151\nL 44 151\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 199\nL 44 199\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 247\nL 44 247\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 295\nL 44 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 5\nL 395 5\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 54\nL 395 54\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 103\nL 395 103\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 151\nL 395 151\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 199\nL 395 199\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 247\nL 395 247\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><text x=\"12\" y=\"299\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><text x=\"16\" y=\"251\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><text x=\"12\" y=\"203\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text><text x=\"16\" y=\"155\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Thu</text><text x=\"23\" y=\"107\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Fri</text><text x=\"19\" y=\"58\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sat</text><text x=\"16\" y=\"9\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sun</text></svg>",
		},
		// 文本居中展示
		// axis位于left
		{
			newOption: func() *axisOption {
				opt := getDefaultOption()
				opt.style.Position = PositionLeft
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 44 5\nL 44 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 5\nL 44 5\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 47\nL 44 47\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 89\nL 44 89\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 131\nL 44 131\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 172\nL 44 172\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 213\nL 44 213\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 254\nL 44 254\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 295\nL 44 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 39 5\nL 395 5\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 47\nL 395 47\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 89\nL 395 89\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 131\nL 395 131\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 172\nL 395 172\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 213\nL 395 213\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 39 254\nL 395 254\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><text x=\"12\" y=\"280\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><text x=\"16\" y=\"239\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><text x=\"12\" y=\"198\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text><text x=\"16\" y=\"157\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Thu</text><text x=\"23\" y=\"115\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Fri</text><text x=\"19\" y=\"73\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sat</text><text x=\"16\" y=\"31\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sun</text></svg>",
		},
		// 文本按起始位置展示
		// axis位于right
		{
			newOption: func() *axisOption {
				opt := getDefaultOption()
				opt.style.Position = PositionRight
				opt.style.BoundaryGap = FalseFlag()
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 361 5\nL 361 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 5\nL 366 5\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 54\nL 366 54\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 103\nL 366 103\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 151\nL 366 151\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 199\nL 366 199\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 247\nL 366 247\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 295\nL 366 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 5 5\nL 360 5\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 54\nL 360 54\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 103\nL 360 103\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 151\nL 360 151\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 199\nL 360 199\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 247\nL 360 247\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><text x=\"369\" y=\"299\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><text x=\"369\" y=\"251\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><text x=\"369\" y=\"203\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text><text x=\"369\" y=\"155\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Thu</text><text x=\"369\" y=\"107\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Fri</text><text x=\"369\" y=\"58\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sat</text><text x=\"369\" y=\"9\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sun</text></svg>",
		},
		// 文本居中展示
		// axis位于right
		{
			newOption: func() *axisOption {
				opt := getDefaultOption()
				opt.style.Position = PositionRight
				return opt
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 361 5\nL 361 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 5\nL 366 5\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 47\nL 366 47\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 89\nL 366 89\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 131\nL 366 131\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 172\nL 366 172\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 213\nL 366 213\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 254\nL 366 254\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 361 295\nL 366 295\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 5 5\nL 360 5\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 47\nL 360 47\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 89\nL 360 89\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 131\nL 360 131\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 172\nL 360 172\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 213\nL 360 213\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><path  d=\"M 5 254\nL 360 254\" style=\"stroke-width:1;stroke:rgba(0,0,0,0.2);fill:none\"/><text x=\"369\" y=\"280\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><text x=\"369\" y=\"239\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text><text x=\"369\" y=\"198\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Wed</text><text x=\"369\" y=\"157\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Thu</text><text x=\"369\" y=\"115\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Fri</text><text x=\"369\" y=\"73\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sat</text><text x=\"369\" y=\"31\" style=\"stroke-width:0;stroke:none;fill:rgba(0,0,0,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sun</text></svg>",
		},
	}
	for _, tt := range tests {
		d, err := NewDraw(DrawOption{
			Width:  400,
			Height: 300,
		}, PaddingOption(chart.Box{
			Left:   5,
			Top:    5,
			Right:  5,
			Bottom: 5,
		}))
		assert.Nil(err)
		opt := tt.newOption()

		d.Axis(*opt.data, *opt.style)

		result, err := d.Bytes()
		assert.Nil(err)
		assert.Equal(tt.result, string(result))
	}
}

func TestMeasureAxis(t *testing.T) {
	assert := assert.New(t)

	d, err := NewDraw(DrawOption{
		Width:  400,
		Height: 300,
	})
	assert.Nil(err)
	data := NewAxisDataListFromStringList([]string{
		"Mon",
		"Sun",
	})
	f, _ := chart.GetDefaultFont()
	width := d.measureAxis(data, AxisStyle{
		FontSize: 12,
		Font:     f,
		Position: PositionLeft,
	})
	assert.Equal(44, width)

	height := d.measureAxis(data, AxisStyle{
		FontSize: 12,
		Font:     f,
		Position: PositionTop,
	})
	assert.Equal(28, height)
}
