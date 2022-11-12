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
	"errors"
	"sync"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2/roboto"
)

var fonts = sync.Map{}
var ErrFontNotExists = errors.New("font is not exists")
var defaultFontFamily = "defaultFontFamily"

func init() {
	name := "roboto"
	_ = InstallFont(name, roboto.Roboto)
	font, _ := GetFont(name)
	SetDefaultFont(font)
}

// InstallFont installs the font for charts
func InstallFont(fontFamily string, data []byte) error {
	font, err := truetype.Parse(data)
	if err != nil {
		return err
	}
	fonts.Store(fontFamily, font)
	return nil
}

// GetDefaultFont get default font
func GetDefaultFont() (*truetype.Font, error) {
	return GetFont(defaultFontFamily)
}

// SetDefaultFont set default font
func SetDefaultFont(font *truetype.Font) {
	if font == nil {
		return
	}
	fonts.Store(defaultFontFamily, font)
}

// GetFont get the font by font family
func GetFont(fontFamily string) (*truetype.Font, error) {
	value, ok := fonts.Load(fontFamily)
	if !ok {
		return nil, ErrFontNotExists
	}
	f, ok := value.(*truetype.Font)
	if !ok {
		return nil, ErrFontNotExists
	}
	return f, nil
}
