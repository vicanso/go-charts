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

func TestRange(t *testing.T) {
	assert := assert.New(t)

	r := NewRange(0, 8, 6)
	assert.Equal(0.0, r.Min)
	assert.Equal(12.0, r.Max)

	r = NewRange(0, 12, 6)
	assert.Equal(0.0, r.Min)
	assert.Equal(30.0, r.Max)

	r = NewRange(-13, 18, 6)
	assert.Equal(-20.0, r.Min)
	assert.Equal(40.0, r.Max)

	r = NewRange(0, 400, 6)
	assert.Equal(0.0, r.Min)
	assert.Equal(480.0, r.Max)
}

func TestRangeHeightWidth(t *testing.T) {
	assert := assert.New(t)
	r := NewRange(0, 8, 6)
	r.Size = 100

	assert.Equal(33, r.getHeight(4))
	assert.Equal(67, r.getRestHeight(4))

	assert.Equal(33, r.getWidth(4))
	r.Boundary = true
	assert.Equal(41, r.getWidth(4))
}

func TestRangeGetRange(t *testing.T) {
	assert := assert.New(t)
	r := NewRange(0, 8, 6)
	r.Size = 120

	f1, f2 := r.GetRange(0)
	assert.Equal(0.0, f1)
	assert.Equal(20.0, f2)

	f1, f2 = r.GetRange(2)
	assert.Equal(40.0, f1)
	assert.Equal(60.0, f2)
}
