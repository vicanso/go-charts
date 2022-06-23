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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableChart(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewTableChart(p, TableChartOption{
					Header: []string{
						"Name",
						"Age",
						"Address",
						"Tag",
						"Action",
					},
					Spans: []int{
						1,
						1,
						2,
						1,
						1,
					},
					Data: [][]string{
						{
							"John Brown",
							"32",
							"New York No. 1 Lake Park",
							"nice, developer",
							"Send Mail",
						},
						{
							"Jim Green	",
							"42",
							"London No. 1 Lake Park",
							"wow",
							"Send Mail",
						},
						{
							"Joe Black	",
							"32",
							"Sidney No. 1 Lake Park",
							"cool, teacher",
							"Send Mail",
						},
					},
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
		},
	}
	for _, tt := range tests {
		p, err := NewPainter(PainterOptions{
			Type:   ChartOutputSVG,
			Width:  600,
			Height: 400,
		}, PainterThemeOption(defaultTheme))
		assert.Nil(err)
		data, err := tt.render(p)
		assert.Nil(err)
		fmt.Println(string(data))
		assert.Equal(tt.result, string(data))
	}

}
