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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestChartSetDefaultWidthHeight(t *testing.T) {
	assert := assert.New(t)

	width := defaultChartWidth
	height := defaultChartHeight
	defer SetDefaultWidth(width)
	defer SetDefaultHeight(height)

	SetDefaultWidth(60)
	assert.Equal(60, defaultChartWidth)
	SetDefaultHeight(40)
	assert.Equal(40, defaultChartHeight)
}

func TestChartFillDefault(t *testing.T) {
	assert := assert.New(t)
	// default value
	opt := ChartOption{}
	opt.FillDefault("")
	// padding
	assert.Equal(chart.Box{
		Top:    10,
		Right:  10,
		Bottom: 10,
		Left:   10,
	}, opt.Padding)
	// background color
	assert.Equal(drawing.ColorWhite, opt.BackgroundColor)
	// title font color
	assert.Equal(drawing.Color{
		R: 70,
		G: 70,
		B: 70,
		A: 255,
	}, opt.Title.Style.FontColor)
	// title font size
	assert.Equal(float64(14), opt.Title.Style.FontSize)
	// sub title font color
	assert.Equal(drawing.Color{
		R: 70,
		G: 70,
		B: 70,
		A: 180,
	}, opt.Title.SubtextStyle.FontColor)
	// sub title font size
	assert.Equal(float64(10), opt.Title.SubtextStyle.FontSize)
	// legend font size
	assert.Equal(float64(10), opt.Legend.Style.FontSize)
	// legend position
	assert.Equal("center", opt.Legend.Left)
	assert.Equal(drawing.Color{
		R: 70,
		G: 70,
		B: 70,
		A: 255,
	}, opt.Legend.Style.FontColor)

	// y axis
	opt = ChartOption{
		SeriesList: SeriesList{
			{
				YAxisIndex: 1,
			},
		},
	}
	opt.FillDefault("")
	assert.Equal([]YAxisOption{
		{},
		{},
	}, opt.YAxisList)
	opt = ChartOption{}
	opt.FillDefault("")
	assert.Equal([]YAxisOption{
		{},
	}, opt.YAxisList)

	// legend get from series's name

	opt = ChartOption{
		SeriesList: SeriesList{
			{
				Name: "a",
			},
			{
				Name: "b",
			},
		},
	}
	opt.FillDefault("")
	assert.Equal([]string{
		"a",
		"b",
	}, opt.Legend.Data)
	// series name set by legend
	opt = ChartOption{
		Legend: LegendOption{
			Data: []string{
				"a",
				"b",
			},
		},
		SeriesList: SeriesList{
			{},
			{},
		},
	}
	opt.FillDefault("")
	assert.Equal("a", opt.SeriesList[0].Name)
	assert.Equal("b", opt.SeriesList[1].Name)
}

func TestChartGetWidthHeight(t *testing.T) {
	assert := assert.New(t)

	opt := ChartOption{
		Width: 10,
	}
	assert.Equal(10, opt.getWidth())
	opt.Width = 0
	assert.Equal(600, opt.getWidth())
	opt.Parent = &Draw{
		Box: chart.Box{
			Left:  10,
			Right: 50,
		},
	}
	assert.Equal(40, opt.getWidth())

	opt = ChartOption{
		Height: 20,
	}
	assert.Equal(20, opt.getHeight())
	opt.Height = 0
	assert.Equal(400, opt.getHeight())
	opt.Parent = &Draw{
		Box: chart.Box{
			Top:    20,
			Bottom: 80,
		},
	}
	assert.Equal(60, opt.getHeight())
}

func TestChartRender(t *testing.T) {
	assert := assert.New(t)

	d, err := Render(ChartOption{
		Width:  800,
		Height: 600,
		Legend: LegendOption{
			Top: "-90",
			Data: []string{
				"Milk Tea",
				"Matcha Latte",
				"Cheese Cocoa",
				"Walnut Brownie",
			},
		},
		Padding: chart.Box{
			Top: 100,
		},
		XAxis: NewXAxisOption([]string{
			"2012",
			"2013",
			"2014",
			"2015",
			"2016",
			"2017",
		}),
		YAxisList: []YAxisOption{
			{

				Min: NewFloatPoint(0),
				Max: NewFloatPoint(90),
			},
		},
		SeriesList: []Series{
			NewSeriesFromValues([]float64{
				56.5,
				82.1,
				88.7,
				70.1,
				53.4,
				85.1,
			}),
			NewSeriesFromValues([]float64{
				51.1,
				51.4,
				55.1,
				53.3,
				73.8,
				68.7,
			}),
			NewSeriesFromValues([]float64{
				40.1,
				62.2,
				69.5,
				36.4,
				45.2,
				32.5,
			}, ChartTypeBar),
			NewSeriesFromValues([]float64{
				25.2,
				37.1,
				41.2,
				18,
				33.9,
				49.1,
			}, ChartTypeBar),
		},
		Children: []ChartOption{
			{
				Legend: LegendOption{
					Show: FalseFlag(),
					Data: []string{
						"Milk Tea",
						"Matcha Latte",
						"Cheese Cocoa",
						"Walnut Brownie",
					},
				},
				Box: chart.Box{
					Top:    20,
					Left:   400,
					Right:  500,
					Bottom: 120,
				},
				SeriesList: NewPieSeriesList([]float64{
					435.9,
					354.3,
					285.9,
					204.5,
				}, PieSeriesOption{
					Label: SeriesLabel{
						Show: true,
					},
					Radius: "35%",
				}),
			},
			{
				Legend: NewLegendOption([]string{
					"Allocated Budget",
					"Actual Spending",
				}),
				Box: chart.Box{
					Top:    20,
					Left:   0,
					Right:  200,
					Bottom: 120,
				},
				RadarIndicators: []RadarIndicator{
					{
						Name: "Sales",
						Max:  6500,
					},
					{
						Name: "Administration",
						Max:  16000,
					},
					{
						Name: "Information Technology",
						Max:  30000,
					},
					{
						Name: "Customer Support",
						Max:  38000,
					},
					{
						Name: "Development",
						Max:  52000,
					},
					{
						Name: "Marketing",
						Max:  25000,
					},
				},
				SeriesList: SeriesList{
					{
						Type: ChartTypeRadar,
						Data: NewSeriesDataFromValues([]float64{
							4200,
							3000,
							20000,
							35000,
							50000,
							18000,
						}),
					},
					{
						Type:  ChartTypeRadar,
						index: 1,
						Data: NewSeriesDataFromValues([]float64{
							5000,
							14000,
							28000,
							26000,
							42000,
							21000,
						}),
					},
				},
			},
		},
	})
	assert.Nil(err)
	data, err := d.Bytes()
	assert.Nil(err)
	assert.Equal("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"800\" height=\"600\">\\n<path  d=\"M 0 0\nL 800 0\nL 800 600\nL 0 600\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 40 575\nL 800 575\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 40 575\nL 40 580\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 167 575\nL 167 580\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 294 575\nL 294 580\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 421 575\nL 421 580\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 548 575\nL 548 580\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 674 575\nL 674 580\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 800 575\nL 800 580\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><text x=\"88\" y=\"592\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2012</text><text x=\"215\" y=\"592\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2013</text><text x=\"342\" y=\"592\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2014</text><text x=\"469\" y=\"592\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2015</text><text x=\"596\" y=\"592\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2016</text><text x=\"722\" y=\"592\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2017</text><path  d=\"M 40 100\nL 40 575\" style=\"stroke-width:1;stroke:none;fill:none\"/><path  d=\"M 40 100\nL 800 100\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 40 180\nL 800 180\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 40 259\nL 800 259\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 40 338\nL 800 338\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 40 417\nL 800 417\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 40 496\nL 800 496\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><text x=\"26\" y=\"579\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">0</text><text x=\"19\" y=\"500\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">15</text><text x=\"19\" y=\"421\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">30</text><text x=\"19\" y=\"342\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">45</text><text x=\"19\" y=\"263\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">60</text><text x=\"19\" y=\"184\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">75</text><text x=\"19\" y=\"104\" style=\"stroke-width:0;stroke:none;fill:rgba(110,112,121,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">90</text><path  d=\"M 50 364\nL 100 364\nL 100 574\nL 50 574\nL 50 364\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 177 247\nL 227 247\nL 227 574\nL 177 574\nL 177 247\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 304 209\nL 354 209\nL 354 574\nL 304 574\nL 304 209\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 431 383\nL 481 383\nL 481 574\nL 431 574\nL 431 383\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 558 337\nL 608 337\nL 608 574\nL 558 574\nL 558 337\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 684 404\nL 734 404\nL 734 574\nL 684 574\nL 684 404\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 105 442\nL 155 442\nL 155 574\nL 105 574\nL 105 442\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 232 380\nL 282 380\nL 282 574\nL 232 574\nL 232 380\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 359 358\nL 409 358\nL 409 574\nL 359 574\nL 359 358\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 486 480\nL 536 480\nL 536 574\nL 486 574\nL 486 480\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 613 397\nL 663 397\nL 663 574\nL 613 574\nL 613 397\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 739 316\nL 789 316\nL 789 574\nL 739 574\nL 739 316\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 103 277\nL 230 142\nL 356 107\nL 483 206\nL 610 294\nL 736 126\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:none\"/><circle cx=\"103\" cy=\"277\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"230\" cy=\"142\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"356\" cy=\"107\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"483\" cy=\"206\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"610\" cy=\"294\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"736\" cy=\"126\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"M 103 306\nL 230 304\nL 356 285\nL 483 294\nL 610 186\nL 736 213\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:none\"/><circle cx=\"103\" cy=\"306\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"230\" cy=\"304\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"356\" cy=\"285\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"483\" cy=\"294\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"610\" cy=\"186\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"736\" cy=\"213\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"M 155 20\nL 185 20\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"170\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"190\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Milk Tea</text><path  d=\"M 255 20\nL 285 20\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"270\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"290\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Matcha Latte</text><path  d=\"M 381 20\nL 411 20\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"396\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"416\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Cheese Cocoa</text><path  d=\"M 514 20\nL 544 20\" style=\"stroke-width:3;stroke:rgba(238,102,102,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"529\" cy=\"20\" r=\"5\" style=\"stroke-width:3;stroke:rgba(238,102,102,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(238,102,102,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"549\" y=\"25\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Walnut Brownie</text><path  d=\"M 450 70\nL 450 35\nA 35 35 122.54 0 1 479 88\nL 450 70\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 480 54\nL 489 49\nM 489 49\nL 499 49\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"502\" y=\"54\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Milk Tea: 34.03%</text><path  d=\"M 450 70\nL 479 88\nA 35 35 99.60 0 1 427 95\nL 450 70\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 454 104\nL 455 114\nM 455 114\nL 465 114\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"468\" y=\"119\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Matcha Latte: 27.66%</text><path  d=\"M 450 70\nL 427 95\nA 35 35 80.37 0 1 421 52\nL 450 70\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 416 74\nL 406 76\nM 406 76\nL 396 76\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"262\" y=\"81\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Cheese Cocoa: 22.32%</text><path  d=\"M 450 70\nL 421 52\nA 35 35 57.49 0 1 450 35\nL 450 70\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 434 40\nL 429 31\nM 429 31\nL 419 31\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"277\" y=\"36\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Walnut Brownie: 15.96%</text><path  d=\"M 100 62\nL 106 66\nL 106 73\nL 100 78\nL 94 74\nL 94 67\nL 100 62\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 100 54\nL 113 62\nL 113 77\nL 100 86\nL 87 78\nL 87 63\nL 100 54\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 100 46\nL 120 58\nL 120 81\nL 100 94\nL 80 82\nL 80 59\nL 100 46\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 100 38\nL 127 54\nL 127 85\nL 100 102\nL 73 86\nL 73 55\nL 100 38\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 100 30\nL 134 50\nL 134 89\nL 100 110\nL 66 90\nL 66 51\nL 100 30\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 100 70\nL 100 30\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 100 70\nL 134 50\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 100 70\nL 134 89\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 100 70\nL 100 110\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 100 70\nL 66 90\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 100 70\nL 66 51\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><text x=\"84\" y=\"23\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sales</text><text x=\"139\" y=\"55\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Administration</text><text x=\"139\" y=\"94\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Information Technology</text><text x=\"48\" y=\"127\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Customer Support</text><text x=\"-15\" y=\"95\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Development</text><text x=\"2\" y=\"56\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Marketing</text><path  d=\"M 100 45\nL 106 67\nL 123 83\nL 100 106\nL 67 89\nL 76 56\nL 100 45\" style=\"stroke-width:2;stroke:rgba(84,112,198,1.0);fill:none\"/><path  d=\"M 100 45\nL 106 67\nL 123 83\nL 100 106\nL 67 89\nL 76 56\nL 100 45\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,0.1)\"/><circle cx=\"100\" cy=\"45\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"106\" cy=\"67\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"123\" cy=\"83\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"100\" cy=\"106\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"67\" cy=\"89\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"76\" cy=\"56\" r=\"2\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"M 100 40\nL 130 53\nL 132 88\nL 100 97\nL 73 86\nL 71 54\nL 100 40\" style=\"stroke-width:2;stroke:rgba(145,204,117,1.0);fill:none\"/><path  d=\"M 100 40\nL 130 53\nL 132 88\nL 100 97\nL 73 86\nL 71 54\nL 100 40\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,0.1)\"/><circle cx=\"100\" cy=\"40\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"130\" cy=\"53\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"132\" cy=\"88\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"100\" cy=\"97\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"73\" cy=\"86\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"71\" cy=\"54\" r=\"2\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"M -42 30\nL -12 30\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"-27\" cy=\"30\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"-7\" y=\"35\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Allocated Budget</text><path  d=\"M 107 30\nL 137 30\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><circle cx=\"122\" cy=\"30\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(255,255,255,1.0)\"/><text x=\"142\" y=\"35\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Actual Spending</text></svg>", string(data))
}

func BenchmarkMultiChartPNGRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		opt := ChartOption{
			Type: ChartOutputPNG,
			Legend: LegendOption{
				Top: "-90",
				Data: []string{
					"Milk Tea",
					"Matcha Latte",
					"Cheese Cocoa",
					"Walnut Brownie",
				},
			},
			Padding: chart.Box{
				Top:    100,
				Right:  10,
				Bottom: 10,
				Left:   10,
			},
			XAxis: NewXAxisOption([]string{
				"2012",
				"2013",
				"2014",
				"2015",
				"2016",
				"2017",
			}),
			YAxisList: []YAxisOption{
				{

					Min: NewFloatPoint(0),
					Max: NewFloatPoint(90),
				},
			},
			SeriesList: []Series{
				NewSeriesFromValues([]float64{
					56.5,
					82.1,
					88.7,
					70.1,
					53.4,
					85.1,
				}),
				NewSeriesFromValues([]float64{
					51.1,
					51.4,
					55.1,
					53.3,
					73.8,
					68.7,
				}),
				NewSeriesFromValues([]float64{
					40.1,
					62.2,
					69.5,
					36.4,
					45.2,
					32.5,
				}, ChartTypeBar),
				NewSeriesFromValues([]float64{
					25.2,
					37.1,
					41.2,
					18,
					33.9,
					49.1,
				}, ChartTypeBar),
			},
			Children: []ChartOption{
				{
					Legend: LegendOption{
						Show: FalseFlag(),
						Data: []string{
							"Milk Tea",
							"Matcha Latte",
							"Cheese Cocoa",
							"Walnut Brownie",
						},
					},
					Box: chart.Box{
						Top:    20,
						Left:   400,
						Right:  500,
						Bottom: 120,
					},
					SeriesList: NewPieSeriesList([]float64{
						435.9,
						354.3,
						285.9,
						204.5,
					}, PieSeriesOption{
						Label: SeriesLabel{
							Show: true,
						},
						Radius: "35%",
					}),
				},
			},
		}
		d, err := Render(opt)
		if err != nil {
			panic(err)
		}
		buf, err := d.Bytes()
		if err != nil {
			panic(err)
		}
		if len(buf) == 0 {
			panic(errors.New("data is nil"))
		}
	}
}

func BenchmarkMultiChartSVGRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		opt := ChartOption{
			Legend: LegendOption{
				Top: "-90",
				Data: []string{
					"Milk Tea",
					"Matcha Latte",
					"Cheese Cocoa",
					"Walnut Brownie",
				},
			},
			Padding: chart.Box{
				Top:    100,
				Right:  10,
				Bottom: 10,
				Left:   10,
			},
			XAxis: NewXAxisOption([]string{
				"2012",
				"2013",
				"2014",
				"2015",
				"2016",
				"2017",
			}),
			YAxisList: []YAxisOption{
				{

					Min: NewFloatPoint(0),
					Max: NewFloatPoint(90),
				},
			},
			SeriesList: []Series{
				NewSeriesFromValues([]float64{
					56.5,
					82.1,
					88.7,
					70.1,
					53.4,
					85.1,
				}),
				NewSeriesFromValues([]float64{
					51.1,
					51.4,
					55.1,
					53.3,
					73.8,
					68.7,
				}),
				NewSeriesFromValues([]float64{
					40.1,
					62.2,
					69.5,
					36.4,
					45.2,
					32.5,
				}, ChartTypeBar),
				NewSeriesFromValues([]float64{
					25.2,
					37.1,
					41.2,
					18,
					33.9,
					49.1,
				}, ChartTypeBar),
			},
			Children: []ChartOption{
				{
					Legend: LegendOption{
						Show: FalseFlag(),
						Data: []string{
							"Milk Tea",
							"Matcha Latte",
							"Cheese Cocoa",
							"Walnut Brownie",
						},
					},
					Box: chart.Box{
						Top:    20,
						Left:   400,
						Right:  500,
						Bottom: 120,
					},
					SeriesList: NewPieSeriesList([]float64{
						435.9,
						354.3,
						285.9,
						204.5,
					}, PieSeriesOption{
						Label: SeriesLabel{
							Show: true,
						},
						Radius: "35%",
					}),
				},
			},
		}
		d, err := Render(opt)
		if err != nil {
			panic(err)
		}
		buf, err := d.Bytes()
		if err != nil {
			panic(err)
		}
		if len(buf) == 0 {
			panic(errors.New("data is nil"))
		}
	}
}
