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
)

func TestNewXAxisOption(t *testing.T) {
	assert := assert.New(t)

	opt := NewXAxisOption([]string{
		"a",
		"b",
	}, FalseFlag())

	assert.Equal(XAxisOption{
		Data: []string{
			"a",
			"b",
		},
		BoundaryGap: FalseFlag(),
	}, opt)

}
func TestDrawXAxis(t *testing.T) {
	assert := assert.New(t)

	newDraw := func() *Draw {
		d, _ := NewDraw(DrawOption{
			Width:  400,
			Height: 300,
		})
		return d
	}

	tests := []struct {
		newDraw   func() *Draw
		newOption func() *XAxisOption
		result    string
	}{
		{
			newDraw: newDraw,
			newOption: func() *XAxisOption {
				return &XAxisOption{
					BoundaryGap: FalseFlag(),
					Data: []string{
						"Mon",
						"Tue",
					},
				}
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 40 275\nL 400 275\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 40 275\nL 40 280\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 400 275\nL 400 280\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><text x=\"27\" y=\"292\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Mon</text><text x=\"389\" y=\"292\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Tue</text></svg>",
		},
		{
			newDraw: newDraw,
			newOption: func() *XAxisOption {
				return &XAxisOption{
					Data: []string{
						"01-01",
						"01-02",
						"01-03",
						"01-04",
						"01-05",
						"01-06",
						"01-07",
						"01-08",
						"01-09",
					},
					SplitNumber: 3,
				}
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"400\" height=\"300\">\\n<path  d=\"M 40 275\nL 400 275\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 40 275\nL 40 280\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 160 275\nL 160 280\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 280 275\nL 280 280\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 400 275\nL 400 280\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><text x=\"83\" y=\"292\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">01-02</text><text x=\"203\" y=\"292\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">01-05</text><text x=\"323\" y=\"292\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">01-08</text></svg>",
		},
	}

	for _, tt := range tests {
		d := tt.newDraw()
		height, _, err := drawXAxis(d, tt.newOption(), 1)
		assert.Nil(err)
		assert.Equal(25, height)
		data, err := d.Bytes()
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}
