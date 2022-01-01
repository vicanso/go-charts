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
	"errors"
	"io"
	"sync"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const (
	ThemeLight = "light"
	ThemeDark  = "dark"
)

const (
	DefaultChartWidth  = 800
	DefaultChartHeight = 400
)

type (
	Title struct {
		Text  string
		Style chart.Style
		Font  *truetype.Font
		Left  string
	}
	Legend struct {
		Data    []string
		Align   string
		Padding chart.Box
		Left    string
		Right   string
		Top     string
		Bottom  string
	}
	Options struct {
		Padding      chart.Box
		Width        int
		Height       int
		Theme        string
		XAxis        XAxis
		YAxisOptions []*YAxisOption
		Series       []Series
		Title        Title
		Legend       Legend
		TickPosition chart.TickPosition
		Log          chart.Logger
		Font         *truetype.Font
	}
)

var fonts = sync.Map{}
var ErrFontNotExists = errors.New("font is not exists")

// InstallFont installs the font for charts
func InstallFont(fontFamily string, data []byte) error {
	font, err := truetype.Parse(data)
	if err != nil {
		return err
	}
	fonts.Store(fontFamily, font)
	return nil
}

// GetFont returns the font of font family
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

type Graph interface {
	Render(rp chart.RendererProvider, w io.Writer) error
}

func (o *Options) validate() error {
	if len(o.Series) == 0 {
		return errors.New("series can not be empty")
	}
	xAxisCount := len(o.XAxis.Data)

	for _, item := range o.Series {
		if item.Type != SeriesPie && len(item.Data) != xAxisCount {
			return errors.New("series and xAxis is not matched")
		}
	}
	return nil
}

func (o *Options) getWidth() int {
	width := o.Width
	if width <= 0 {
		width = DefaultChartWidth
	}
	return width
}

func (o *Options) getHeight() int {
	height := o.Height
	if height <= 0 {
		height = DefaultChartHeight
	}
	return height
}

func (o *Options) getBackground() chart.Style {
	bg := chart.Style{
		Padding: o.Padding,
	}
	return bg
}

func render(g Graph, rp chart.RendererProvider) ([]byte, error) {
	buf := bytes.Buffer{}
	err := g.Render(rp, &buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ToPNG(g Graph) ([]byte, error) {
	return render(g, chart.PNG)
}

func ToSVG(g Graph) ([]byte, error) {
	return render(g, chart.SVG)
}

func newTitleRenderable(title Title, font *truetype.Font, textColor drawing.Color) chart.Renderable {
	if title.Text == "" || title.Style.Hidden {
		return nil
	}
	title.Font = font
	if title.Style.FontColor.IsZero() {
		title.Style.FontColor = textColor
	}
	return NewTitleCustomize(title)
}

func newPieChart(opt Options) *chart.PieChart {
	values := make(chart.Values, len(opt.Series))
	for index, item := range opt.Series {
		values[index] = chart.Value{
			Value: item.Data[0].Value,
			Label: item.Name,
		}
	}
	p := &chart.PieChart{
		Font:       opt.Font,
		Background: opt.getBackground(),
		Width:      opt.getWidth(),
		Height:     opt.getHeight(),
		Values:     values,
		ColorPalette: &PieThemeColorPalette{
			ThemeColorPalette: ThemeColorPalette{
				Theme: opt.Theme,
			},
		},
	}
	// pie 图表默认设置为居中
	if opt.Title.Left == "" {
		opt.Title.Left = "center"
	}
	themeColorPalette := &ThemeColorPalette{
		Theme: opt.Theme,
	}
	titleRender := newTitleRenderable(opt.Title, p.GetFont(), themeColorPalette.TextColor())
	if titleRender != nil {
		p.Elements = []chart.Renderable{
			titleRender,
		}
	}
	return p
}

func newChart(opt Options) *chart.Chart {
	tickPosition := opt.TickPosition

	xAxis, xValues := GetXAxisAndValues(opt.XAxis, tickPosition, opt.Theme)

	legendSize := len(opt.Legend.Data)
	for index, item := range opt.Series {
		if len(item.XValues) == 0 {
			opt.Series[index].XValues = xValues
		}
		if index < legendSize && opt.Series[index].Name == "" {
			opt.Series[index].Name = opt.Legend.Data[index]
		}
	}

	var secondaryYAxisOption *YAxisOption
	if len(opt.YAxisOptions) != 0 {
		secondaryYAxisOption = opt.YAxisOptions[0]
	}

	yAxisOption := &YAxisOption{
		Disabled: true,
	}
	if len(opt.YAxisOptions) > 1 {
		yAxisOption = opt.YAxisOptions[1]
	}

	c := &chart.Chart{
		Font:       opt.Font,
		Log:        opt.Log,
		Background: opt.getBackground(),
		ColorPalette: &ThemeColorPalette{
			Theme: opt.Theme,
		},
		Width:          opt.getWidth(),
		Height:         opt.getHeight(),
		XAxis:          xAxis,
		YAxis:          GetYAxis(opt.Theme, yAxisOption),
		YAxisSecondary: GetSecondaryYAxis(opt.Theme, secondaryYAxisOption),
		Series:         GetSeries(opt.Series, tickPosition, opt.Theme),
	}

	elements := make([]chart.Renderable, 0)

	if legendSize != 0 {
		elements = append(elements, NewLegendCustomize(c.Series, LegendOption{
			Theme:    opt.Theme,
			IconDraw: DefaultLegendIconDraw,
			Align:    opt.Legend.Align,
			Padding:  opt.Legend.Padding,
			Left:     opt.Legend.Left,
			Right:    opt.Legend.Right,
			Top:      opt.Legend.Top,
			Bottom:   opt.Legend.Bottom,
		}))
	}
	titleRender := newTitleRenderable(opt.Title, c.GetFont(), c.GetColorPalette().TextColor())
	if titleRender != nil {
		elements = append(elements, titleRender)
	}
	if len(elements) != 0 {
		c.Elements = elements
	}
	return c
}

func New(opt Options) (Graph, error) {
	err := opt.validate()
	if err != nil {
		return nil, err
	}
	if opt.Series[0].Type == SeriesPie {
		return newPieChart(opt), nil
	}

	return newChart(opt), nil
}
