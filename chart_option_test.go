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

func TestOptionFunc(t *testing.T) {
	assert := assert.New(t)

	fns := []OptionFunc{
		TypeOptionFunc(ChartOutputPNG),
		FontFamilyOptionFunc("fontFamily"),
		ThemeOptionFunc("black"),
		TitleOptionFunc(TitleOption{
			Text: "title",
		}),
		LegendOptionFunc(LegendOption{
			Data: []string{
				"a",
				"b",
			},
		}),
		XAxisOptionFunc(NewXAxisOption([]string{
			"Mon",
			"Tue",
		})),
		YAxisOptionFunc(YAxisOption{
			Min: NewFloatPoint(0),
			Max: NewFloatPoint(100),
		}),
		WidthOptionFunc(400),
		HeightOptionFunc(300),
		PaddingOptionFunc(chart.Box{
			Top: 10,
		}),
		BoxOptionFunc(chart.Box{
			Left:  0,
			Right: 300,
		}),
		ChildOptionFunc(ChartOption{}),
		RadarIndicatorOptionFunc(RadarIndicator{
			Min: 0,
			Max: 10,
		}),
		BackgroundColorOptionFunc(drawing.ColorBlack),
	}

	opt := ChartOption{}
	for _, fn := range fns {
		fn(&opt)
	}

	assert.Equal("png", opt.Type)
	assert.Equal("fontFamily", opt.FontFamily)
	assert.Equal("black", opt.Theme)
	assert.Equal(TitleOption{
		Text: "title",
	}, opt.Title)
	assert.Equal(LegendOption{
		Data: []string{
			"a",
			"b",
		},
	}, opt.Legend)
	assert.Equal(NewXAxisOption([]string{
		"Mon",
		"Tue",
	}), opt.XAxis)
	assert.Equal([]YAxisOption{
		{
			Min: NewFloatPoint(0),
			Max: NewFloatPoint(100),
		},
	}, opt.YAxisList)
	assert.Equal(400, opt.Width)
	assert.Equal(300, opt.Height)
	assert.Equal(chart.Box{
		Top: 10,
	}, opt.Padding)
	assert.Equal(chart.Box{
		Left:  0,
		Right: 300,
	}, opt.Box)
	assert.Equal(1, len(opt.Children))
	assert.Equal([]RadarIndicator{
		{
			Min: 0,
			Max: 10,
		},
	}, opt.RadarIndicators)
	assert.Equal(drawing.ColorBlack, opt.BackgroundColor)
}

func TestLineRender(t *testing.T) {
	assert := assert.New(t)

	d, err := LineRender([][]float64{
		{
			1,
			2,
			3,
		},
		{
			1,
			5,
			2,
		},
	},
		XAxisOptionFunc(NewXAxisOption([]string{
			"01",
			"02",
			"03",
		})),
	)
	assert.Nil(err)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 400\nL 0 400\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 50 365\nL 590 365\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 50 365\nL 50 370\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 230 365\nL 230 370\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 410 365\nL 410 370\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 590 365\nL 590 370\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><text x=\"132\" y=\"382\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">01</text><text x=\"312\" y=\"382\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">02</text><text x=\"492\" y=\"382\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">03</text><path  d=\"M 50 10\nL 50 365\" style=\"stroke-width:1;stroke:none;fill:none\"/><path  d=\"M 50 10\nL 590 10\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 70\nL 590 70\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 129\nL 590 129\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 188\nL 590 188\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 247\nL 590 247\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 306\nL 590 306\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><text x=\"36\" y=\"369\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">0</text><text x=\"36\" y=\"310\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2</text><text x=\"36\" y=\"251\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">4</text><text x=\"36\" y=\"192\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">6</text><text x=\"36\" y=\"133\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">8</text><text x=\"29\" y=\"74\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">10</text><text x=\"29\" y=\"14\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">12</text><path  d=\"M 140 336\nL 320 306\nL 499 277\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:none\"/><circle cx=\"140\" cy=\"336\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"320\" cy=\"306\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"499\" cy=\"277\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"M 140 336\nL 320 218\nL 499 306\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:none\"/><circle cx=\"140\" cy=\"336\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"320\" cy=\"218\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"499\" cy=\"306\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/></svg>", string(data))
}

func TestBarRender(t *testing.T) {
	assert := assert.New(t)

	d, err := BarRender([][]float64{
		{
			1,
			2,
			3,
		},
		{
			1,
			5,
			2,
		},
	},
		XAxisOptionFunc(NewXAxisOption([]string{
			"01",
			"02",
			"03",
		})),
	)
	assert.Nil(err)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 400\nL 0 400\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 50 365\nL 590 365\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 50 365\nL 50 370\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 230 365\nL 230 370\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 410 365\nL 410 370\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 590 365\nL 590 370\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><text x=\"132\" y=\"382\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">01</text><text x=\"312\" y=\"382\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">02</text><text x=\"492\" y=\"382\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">03</text><path  d=\"M 50 10\nL 50 365\" style=\"stroke-width:1;stroke:none;fill:none\"/><path  d=\"M 50 10\nL 590 10\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 70\nL 590 70\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 129\nL 590 129\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 188\nL 590 188\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 247\nL 590 247\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 50 306\nL 590 306\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><text x=\"36\" y=\"369\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">0</text><text x=\"36\" y=\"310\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2</text><text x=\"36\" y=\"251\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">4</text><text x=\"36\" y=\"192\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">6</text><text x=\"36\" y=\"133\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">8</text><text x=\"29\" y=\"74\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">10</text><text x=\"29\" y=\"14\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">12</text><path  d=\"M 60 336\nL 137 336\nL 137 364\nL 60 364\nL 60 336\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 240 306\nL 317 306\nL 317 364\nL 240 364\nL 240 306\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 420 277\nL 497 277\nL 497 364\nL 420 364\nL 420 277\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 142 336\nL 219 336\nL 219 364\nL 142 364\nL 142 336\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 322 218\nL 399 218\nL 399 364\nL 322 364\nL 322 218\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 502 306\nL 579 306\nL 579 364\nL 502 364\nL 502 306\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/></svg>", string(data))
}

func TestPieRender(t *testing.T) {
	assert := assert.New(t)

	d, err := PieRender([]float64{
		1,
		3,
		5,
	})
	assert.Nil(err)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 400\nL 0 400\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 300 200\nL 300 48\nA 152 152 40.00 0 1 397 84\nL 300 200\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 300 200\nL 397 84\nA 152 152 120.00 0 1 351 342\nL 300 200\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 300 200\nL 351 342\nA 152 152 200.00 1 1 300 48\nL 300 200\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/></svg>", string(data))
}

func TestRadarRender(t *testing.T) {
	assert := assert.New(t)
	d, err := RadarRender([][]float64{
		{
			1,
			2,
			3,
		},
		{
			1,
			5,
			2,
		},
	},
		RadarIndicatorOptionFunc([]RadarIndicator{
			{
				Name: "A",
				Min:  0,
				Max:  10,
			},
			{
				Name: "B",
				Min:  0,
				Max:  10,
			},
			{
				Name: "C",
				Min:  0,
				Max:  10,
			},
		}...),
	)
	assert.Nil(err)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 400\nL 0 400\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 300 170\nL 325 214\nL 275 215\nL 300 170\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 140\nL 351 229\nL 249 230\nL 300 140\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 110\nL 377 244\nL 223 245\nL 300 110\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 80\nL 403 259\nL 197 260\nL 300 80\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 50\nL 429 274\nL 171 275\nL 300 50\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 200\nL 300 50\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 200\nL 429 274\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 300 200\nL 171 275\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><text x=\"296\" y=\"43\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">A</text><text x=\"434\" y=\"279\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">B</text><text x=\"157\" y=\"280\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">C</text><path  d=\"M 300 185\nL 325 214\nL 262 222\nL 300 185\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:none\"/><path  d=\"M 300 185\nL 325 214\nL 262 222\nL 300 185\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,0.1)\"/><circle cx=\"300\" cy=\"185\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"325\" cy=\"214\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"262\" cy=\"222\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"M 300 185\nL 364 237\nL 275 215\nL 300 185\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:none\"/><path  d=\"M 300 185\nL 364 237\nL 275 215\nL 300 185\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,0.1)\"/><circle cx=\"300\" cy=\"185\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"364\" cy=\"237\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"275\" cy=\"215\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/></svg>", string(data))
}

func TestFunnelRender(t *testing.T) {
	assert := assert.New(t)

	d, err := FunnelRender([]float64{
		5,
		3,
		1,
	})
	assert.Nil(err)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 400\nL 0 400\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 10 10\nL 590 10\nL 474 135\nL 126 135\nL 10 10\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><text x=\"280\" y=\"72\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">(100%)</text><path  d=\"M 126 137\nL 474 137\nL 358 262\nL 242 262\nL 126 137\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><text x=\"284\" y=\"199\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">(60%)</text><path  d=\"M 242 264\nL 358 264\nL 300 389\nL 300 389\nL 242 264\" style=\"stroke-width:0;stroke:none;fill:rgba(250,200,88,1.0)\"/><text x=\"284\" y=\"326\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">(20%)</text></svg>", string(data))
}
