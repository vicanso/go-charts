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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieChart(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				values := []float64{
					1048,
					735,
					580,
					484,
					300,
				}
				_, err := NewPieChart(p, PieChartOption{
					SeriesList: NewPieSeriesList(values, PieSeriesOption{
						Label: SeriesLabel{
							Show: true,
						},
					}),
					Title: TitleOption{
						Text:    "Rainfall vs Evaporation",
						Subtext: "Fake Data",
						Left:    PositionCenter,
					},
					Padding: Box{
						Top:    20,
						Right:  20,
						Bottom: 20,
						Left:   20,
					},
					Legend: LegendOption{
						Orient: OrientVertical,
						Data: []string{
							"Search Engine",
							"Direct",
							"Email",
							"Union Ads",
							"Video Ads",
						},
						Left: PositionLeft,
					},
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 560 0\nL 560 360\nL 0 360\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 40 49\nL 70 49\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><circle cx=\"55\" cy=\"49\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"72\" y=\"55\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Search Engine</text><path  d=\"M 40 69\nL 70 69\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><circle cx=\"55\" cy=\"69\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"72\" y=\"75\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Direct</text><path  d=\"M 40 89\nL 70 89\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><circle cx=\"55\" cy=\"89\" r=\"5\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"72\" y=\"95\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Email</text><path  d=\"M 40 109\nL 70 109\" style=\"stroke-width:3;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><circle cx=\"55\" cy=\"109\" r=\"5\" style=\"stroke-width:3;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"72\" y=\"115\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Union Ads</text><path  d=\"M 40 129\nL 70 129\" style=\"stroke-width:3;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><circle cx=\"55\" cy=\"129\" r=\"5\" style=\"stroke-width:3;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"72\" y=\"135\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Video Ads</text><text x=\"222\" y=\"55\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Rainfall vs Evaporation</text><text x=\"266\" y=\"70\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Fake Data</text><path  d=\"M 300 210\nL 300 114\nA 96 96 119.89 0 1 383 257\nL 300 210\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 383 162\nL 396 155\nM 396 155\nL 411 155\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"414\" y=\"160\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Search Engine: 33.3%</text><path  d=\"M 300 210\nL 383 257\nA 96 96 84.08 0 1 262 297\nL 300 210\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 329 301\nL 334 315\nM 334 315\nL 349 315\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"352\" y=\"320\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Direct: 23.35%</text><path  d=\"M 300 210\nL 262 297\nA 96 96 66.35 0 1 205 210\nL 300 210\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 220 262\nL 207 270\nM 207 270\nL 192 270\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"108\" y=\"275\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Email: 18.43%</text><path  d=\"M 300 210\nL 205 210\nA 96 96 55.37 0 1 246 131\nL 300 210\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 216 165\nL 202 158\nM 202 158\nL 187 158\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"76\" y=\"163\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Union Ads: 15.37%</text><path  d=\"M 300 210\nL 246 131\nA 96 96 34.32 0 1 300 114\nL 300 210\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 272 119\nL 268 104\nM 268 104\nL 253 104\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"150\" y=\"109\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Video Ads: 9.53%</text></svg>",
		},
	}
	for _, tt := range tests {
		p, err := NewPainter(PainterOptions{
			Type:   ChartOutputSVG,
			Width:  600,
			Height: 400,
		}, PainterThemeOption(defaultTheme))
		assert.Nil(err)
		data, err := tt.render(p.Child(PainterPaddingOption(Box{
			Left:   20,
			Top:    20,
			Right:  20,
			Bottom: 20,
		})))
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}

func TestPieChartWithLabelsValuesSortedDescending(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				values := []float64{
					84358845,
					68070697,
					58850717,
					48059777,
					36753736,
					19051562,
					17947406,
					11754004,
					10827529,
					10521556,
					10467366,
					10394055,
					9597085,
					9104772,
					6447710,
					5932654,
					5563970,
					5428792,
					5194336,
					3850894,
					2857279,
					2116792,
					1883008,
					1373101,
					920701,
					660809,
					542051,
				}
				_, err := NewPieChart(p, PieChartOption{
					SeriesList: NewPieSeriesList(values, PieSeriesOption{
						Label: SeriesLabel{
							Show:      true,
							Formatter: "{b} ({c} ≅ {d})",
						},
						Radius: "200",
					}),
					Title: TitleOption{
						Text: "European Union member states by population",
						Left: PositionRight,
					},
					Padding: Box{
						Top:    20,
						Right:  20,
						Bottom: 20,
						Left:   20,
					},
					Legend: LegendOption{
						Data: []string{
							"Germany",
							"France",
							"Italy",
							"Spain",
							"Poland",
							"Romania",
							"Netherlands",
							"Belgium",
							"Czech Republic",
							"Sweden",
							"Portugal",
							"Greece",
							"Hungary",
							"Austria",
							"Bulgaria",
							"Denmark",
							"Finland",
							"Slovakia",
							"Ireland",
							"Croatia",
							"Lithuania",
							"Slovenia",
							"Latvia",
							"Estonia",
							"Cyprus",
							"Luxembourg",
							"Malta",
						},
						Show: FalseFlag(),
					},
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"1000\" height=\"800\">\\n<path  d=\"M 0 0\nL 960 0\nL 960 760\nL 0 760\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><text x=\"647\" y=\"55\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">European Union member states by population</text><path  d=\"M 500 402\nL 500 202\nA 200 200 67.71 0 1 685 327\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 611 236\nL 619 224\nM 619 224\nL 634 224\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"637\" y=\"229\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Germany (84358845 ≅ 18.8%)</text><path  d=\"M 500 402\nL 685 327\nA 200 200 54.63 0 1 668 508\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 699 419\nL 714 420\nM 714 420\nL 729 420\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"732\" y=\"425\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">France (68070697 ≅ 15.17%)</text><path  d=\"M 500 402\nL 668 508\nA 200 200 47.23 0 1 536 598\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 611 567\nL 620 580\nM 620 580\nL 635 580\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"638\" y=\"585\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Italy (58850717 ≅ 13.12%)</text><path  d=\"M 500 402\nL 309 460\nA 200 200 14.40 0 1 301 411\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 303 436\nL 289 438\nM 289 438\nL 274 438\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"102\" y=\"443\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Netherlands (17947406 ≅ 4%)</text><path  d=\"M 500 402\nL 332 509\nA 200 200 15.29 0 1 309 460\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 319 485\nL 305 491\nM 305 491\nL 290 491\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"119\" y=\"496\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Romania (19051562 ≅ 4.24%)</text><path  d=\"M 500 402\nL 406 578\nA 200 200 29.50 0 1 332 509\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 364 548\nL 354 559\nM 354 559\nL 339 559\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"178\" y=\"564\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Poland (36753736 ≅ 8.19%)</text><path  d=\"M 500 402\nL 536 598\nA 200 200 38.57 0 1 406 578\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 470 599\nL 467 614\nM 467 614\nL 452 614\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"292\" y=\"619\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Spain (48059777 ≅ 10.71%)</text><path  d=\"M 500 402\nL 301 411\nA 200 200 9.43 0 1 302 379\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 301 395\nL 286 395\nM 286 395\nL 271 395\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"103\" y=\"400\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Belgium (11754004 ≅ 2.62%)</text><path  d=\"M 500 402\nL 302 379\nA 200 200 8.69 0 1 308 349\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 304 364\nL 290 361\nM 290 361\nL 275 361\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"66\" y=\"366\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Czech Republic (10827529 ≅ 2.41%)</text><path  d=\"M 500 402\nL 308 349\nA 200 200 8.44 0 1 318 321\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 312 335\nL 298 330\nM 298 330\nL 283 330\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"117\" y=\"335\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sweden (10521556 ≅ 2.34%)</text><path  d=\"M 500 402\nL 318 321\nA 200 200 8.40 0 1 331 296\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 324 308\nL 311 301\nM 311 301\nL 296 301\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"127\" y=\"306\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Portugal (10467366 ≅ 2.33%)</text><path  d=\"M 500 402\nL 331 296\nA 200 200 8.34 0 1 349 272\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 340 284\nL 328 275\nM 328 275\nL 313 275\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"152\" y=\"280\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Greece (10394055 ≅ 2.31%)</text><path  d=\"M 500 402\nL 349 272\nA 200 200 7.70 0 1 368 253\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 358 262\nL 347 252\nM 347 252\nL 332 252\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"170\" y=\"257\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Hungary (9597085 ≅ 2.13%)</text><path  d=\"M 500 402\nL 368 253\nA 200 200 7.31 0 1 388 237\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 377 245\nL 368 233\nM 368 233\nL 353 233\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"199\" y=\"238\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Austria (9104772 ≅ 2.02%)</text><path  d=\"M 500 402\nL 388 237\nA 200 200 5.18 0 1 403 228\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 395 232\nL 387 217\nM 387 217\nL 372 217\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"211\" y=\"222\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Bulgaria (6447710 ≅ 1.43%)</text><path  d=\"M 500 402\nL 403 228\nA 200 200 4.76 0 1 418 220\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 410 224\nL 404 201\nM 404 201\nL 389 201\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"224\" y=\"206\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Denmark (5932654 ≅ 1.32%)</text><path  d=\"M 500 402\nL 418 220\nA 200 200 4.47 0 1 432 214\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 425 217\nL 419 185\nM 419 185\nL 404 185\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"249\" y=\"190\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Finland (5563970 ≅ 1.24%)</text><path  d=\"M 500 402\nL 432 214\nA 200 200 4.36 0 1 447 210\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 440 212\nL 435 169\nM 435 169\nL 420 169\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"258\" y=\"174\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Slovakia (5428792 ≅ 1.21%)</text><path  d=\"M 500 402\nL 447 210\nA 200 200 4.17 0 1 461 206\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 454 208\nL 450 153\nM 450 153\nL 435 153\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"282\" y=\"158\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Ireland (5194336 ≅ 1.15%)</text><path  d=\"M 500 402\nL 461 206\nA 200 200 3.09 0 1 472 205\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 466 205\nL 464 137\nM 464 137\nL 449 137\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"294\" y=\"142\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Croatia (3850894 ≅ 0.85%)</text><path  d=\"M 500 402\nL 472 205\nA 200 200 2.29 0 1 480 204\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 476 204\nL 474 121\nM 474 121\nL 459 121\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"293\" y=\"126\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Lithuania (2857279 ≅ 0.63%)</text><path  d=\"M 500 402\nL 480 204\nA 200 200 1.70 0 1 485 203\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 482 203\nL 481 105\nM 481 105\nL 466 105\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"304\" y=\"110\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Slovenia (2116792 ≅ 0.47%)</text><path  d=\"M 500 402\nL 485 203\nA 200 200 1.51 0 1 491 203\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 488 203\nL 487 89\nM 487 89\nL 472 89\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"324\" y=\"94\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Latvia (1883008 ≅ 0.41%)</text><path  d=\"M 500 402\nL 491 203\nA 200 200 1.10 0 1 495 203\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 493 203\nL 492 73\nM 492 73\nL 477 73\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"328\" y=\"78\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Estonia (1373101 ≅ 0.3%)</text><path  d=\"M 500 402\nL 495 203\nA 200 200 0.74 0 1 497 203\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 496 203\nL 495 57\nM 495 57\nL 480 57\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"341\" y=\"62\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Cyprus (920701 ≅ 0.2%)</text><path  d=\"M 500 402\nL 497 203\nA 200 200 0.53 0 1 499 203\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 498 203\nL 498 41\nM 498 41\nL 483 41\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"305\" y=\"46\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Luxembourg (660809 ≅ 0.14%)</text><path  d=\"M 500 402\nL 499 203\nA 200 200 0.44 0 1 500 202\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 500 203\nL 500 25\nM 500 25\nL 485 25\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"346\" y=\"30\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Malta (542051 ≅ 0.12%)</text></svg>",
		},
	}
	for _, tt := range tests {
		p, err := NewPainter(PainterOptions{
			Type:   ChartOutputSVG,
			Width:  1000,
			Height: 800,
		}, PainterThemeOption(defaultTheme))
		assert.Nil(err)
		data, err := tt.render(p.Child(PainterPaddingOption(Box{
			Left:   20,
			Top:    20,
			Right:  20,
			Bottom: 20,
		})))
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}

func TestPieChartWithLabelsValuesUnsorted(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				values := []float64{
					9104772,
					11754004,
					6447710,
					3850894,
					920701,
					10827529,
					5932654,
					1373101,
					5563970,
					68070697,
					84358845,
					10394055,
					9597085,
					5194336,
					58850717,
					1883008,
					2857279,
					660809,
					542051,
					17947406,
					36753736,
					10467366,
					19051562,
					5428792,
					2116792,
					48059777,
					10521556,
				}
				_, err := NewPieChart(p, PieChartOption{
					SeriesList: NewPieSeriesList(values, PieSeriesOption{
						Label: SeriesLabel{
							Show:      true,
							Formatter: "{b} ({c} ≅ {d})",
						},
						Radius: "200",
					}),
					Title: TitleOption{
						Text: "European Union member states by population",
						Left: PositionRight,
					},
					Padding: Box{
						Top:    20,
						Right:  20,
						Bottom: 20,
						Left:   20,
					},
					Legend: LegendOption{
						Data: []string{
							"Austria",
							"Belgium",
							"Bulgaria",
							"Croatia",
							"Cyprus",
							"Czech Republic",
							"Denmark",
							"Estonia",
							"Finland",
							"France",
							"Germany",
							"Greece",
							"Hungary",
							"Ireland",
							"Italy",
							"Latvia",
							"Lithuania",
							"Luxembourg",
							"Malta",
							"Netherlands",
							"Poland",
							"Portugal",
							"Romania",
							"Slovakia",
							"Slovenia",
							"Spain",
							"Sweden",
						},
						Show: FalseFlag(),
					},
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"1000\" height=\"800\">\\n<path  d=\"M 0 0\nL 960 0\nL 960 760\nL 0 760\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><text x=\"647\" y=\"55\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">European Union member states by population</text><path  d=\"M 500 402\nL 640 261\nA 200 200 54.63 0 1 697 434\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 690 341\nL 704 336\nM 704 336\nL 719 336\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"722\" y=\"341\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">France (68070697 ≅ 15.17%)</text><path  d=\"M 500 402\nL 629 250\nA 200 200 4.47 0 1 640 261\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 635 255\nL 645 244\nM 645 244\nL 660 244\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"663\" y=\"249\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Finland (5563970 ≅ 1.24%)</text><path  d=\"M 500 402\nL 626 248\nA 200 200 1.10 0 1 629 250\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 627 249\nL 637 228\nM 637 228\nL 652 228\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"655\" y=\"233\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Estonia (1373101 ≅ 0.3%)</text><path  d=\"M 500 402\nL 613 238\nA 200 200 4.76 0 1 626 248\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 619 242\nL 628 212\nM 628 212\nL 643 212\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"646\" y=\"217\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Denmark (5932654 ≅ 1.32%)</text><path  d=\"M 500 402\nL 586 222\nA 200 200 8.69 0 1 613 238\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 600 229\nL 607 196\nM 607 196\nL 622 196\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"625\" y=\"201\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Czech Republic (10827529 ≅ 2.41%)</text><path  d=\"M 500 402\nL 584 221\nA 200 200 0.74 0 1 586 222\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 585 222\nL 592 180\nM 592 180\nL 607 180\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"610\" y=\"185\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Cyprus (920701 ≅ 0.2%)</text><path  d=\"M 500 402\nL 574 217\nA 200 200 3.09 0 1 584 221\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 579 219\nL 585 164\nM 585 164\nL 600 164\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"603\" y=\"169\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Croatia (3850894 ≅ 0.85%)</text><path  d=\"M 500 402\nL 557 211\nA 200 200 5.18 0 1 574 217\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 566 214\nL 571 148\nM 571 148\nL 586 148\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"589\" y=\"153\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Bulgaria (6447710 ≅ 1.43%)</text><path  d=\"M 500 402\nL 525 204\nA 200 200 9.43 0 1 557 211\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 541 207\nL 544 132\nM 544 132\nL 559 132\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"562\" y=\"137\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Belgium (11754004 ≅ 2.62%)</text><path  d=\"M 500 402\nL 500 202\nA 200 200 7.31 0 1 525 204\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 512 203\nL 513 116\nM 513 116\nL 528 116\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"531\" y=\"121\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Austria (9104772 ≅ 2.02%)</text><path  d=\"M 500 402\nL 697 434\nA 200 200 67.71 0 1 544 596\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 645 539\nL 656 549\nM 656 549\nL 671 549\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"674\" y=\"554\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Germany (84358845 ≅ 18.8%)</text><path  d=\"M 500 402\nL 544 596\nA 200 200 8.34 0 1 515 601\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 530 599\nL 532 614\nM 532 614\nL 547 614\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"550\" y=\"619\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Greece (10394055 ≅ 2.31%)</text><path  d=\"M 500 402\nL 515 601\nA 200 200 7.70 0 1 489 601\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 502 601\nL 502 630\nM 502 630\nL 517 630\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"520\" y=\"635\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Hungary (9597085 ≅ 2.13%)</text><path  d=\"M 500 402\nL 309 458\nA 200 200 29.50 0 1 306 357\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 301 407\nL 286 407\nM 286 407\nL 271 407\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"110\" y=\"412\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Poland (36753736 ≅ 8.19%)</text><path  d=\"M 500 402\nL 328 504\nA 200 200 14.40 0 1 309 458\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 317 481\nL 303 487\nM 303 487\nL 288 487\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"116\" y=\"492\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Netherlands (17947406 ≅ 4%)</text><path  d=\"M 500 402\nL 329 505\nA 200 200 0.44 0 1 328 504\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 329 504\nL 316 512\nM 316 512\nL 301 512\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"162\" y=\"517\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Malta (542051 ≅ 0.12%)</text><path  d=\"M 500 402\nL 330 506\nA 200 200 0.53 0 1 329 505\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 330 506\nL 317 528\nM 317 528\nL 302 528\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"124\" y=\"533\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Luxembourg (660809 ≅ 0.14%)</text><path  d=\"M 500 402\nL 335 513\nA 200 200 2.29 0 1 330 506\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 332 510\nL 320 544\nM 320 544\nL 305 544\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"139\" y=\"549\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Lithuania (2857279 ≅ 0.63%)</text><path  d=\"M 500 402\nL 338 517\nA 200 200 1.51 0 1 335 513\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 336 515\nL 324 560\nM 324 560\nL 309 560\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"161\" y=\"565\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Latvia (1883008 ≅ 0.41%)</text><path  d=\"M 500 402\nL 475 600\nA 200 200 47.23 0 1 338 517\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 398 573\nL 390 586\nM 390 586\nL 375 586\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"223\" y=\"591\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Italy (58850717 ≅ 13.12%)</text><path  d=\"M 500 402\nL 489 601\nA 200 200 4.17 0 1 475 600\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 482 601\nL 481 616\nM 481 616\nL 466 616\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"313\" y=\"621\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Ireland (5194336 ≅ 1.15%)</text><path  d=\"M 500 402\nL 306 357\nA 200 200 8.40 0 1 315 329\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 310 343\nL 295 338\nM 295 338\nL 280 338\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"111\" y=\"343\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Portugal (10467366 ≅ 2.33%)</text><path  d=\"M 500 402\nL 315 329\nA 200 200 15.29 0 1 341 282\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 326 305\nL 313 297\nM 313 297\nL 298 297\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"127\" y=\"302\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Romania (19051562 ≅ 4.24%)</text><path  d=\"M 500 402\nL 341 282\nA 200 200 4.36 0 1 350 271\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 345 276\nL 334 267\nM 334 267\nL 319 267\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"157\" y=\"272\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Slovakia (5428792 ≅ 1.21%)</text><path  d=\"M 500 402\nL 350 271\nA 200 200 1.70 0 1 354 266\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 352 268\nL 341 251\nM 341 251\nL 326 251\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"164\" y=\"256\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Slovenia (2116792 ≅ 0.47%)</text><path  d=\"M 500 402\nL 354 266\nA 200 200 38.57 0 1 471 205\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 407 225\nL 400 212\nM 400 212\nL 385 212\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"225\" y=\"217\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Spain (48059777 ≅ 10.71%)</text><path  d=\"M 500 402\nL 471 205\nA 200 200 8.44 0 1 500 202\nL 500 402\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 486 203\nL 485 188\nM 485 188\nL 470 188\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"304\" y=\"193\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Sweden (10521556 ≅ 2.34%)</text></svg>",
		},
	}
	for _, tt := range tests {
		p, err := NewPainter(PainterOptions{
			Type:   ChartOutputSVG,
			Width:  1000,
			Height: 800,
		}, PainterThemeOption(defaultTheme))
		assert.Nil(err)
		data, err := tt.render(p.Child(PainterPaddingOption(Box{
			Left:   20,
			Top:    20,
			Right:  20,
			Bottom: 20,
		})))
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}

func TestPieChartWith100Labels(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				var values []float64
				var labels []string
				for i := 1; i <= 100; i++ {
					values = append(values, float64(1))
					labels = append(labels, "Label "+strconv.Itoa(i))
				}
				_, err := NewPieChart(p, PieChartOption{
					SeriesList: NewPieSeriesList(values, PieSeriesOption{
						Label: SeriesLabel{
							Show: true,
						},
						Radius: "200",
					}),
					Title: TitleOption{
						Text: "Test with 100 labels",
						Left: PositionRight,
					},
					Padding: Box{
						Top:    20,
						Right:  20,
						Bottom: 20,
						Left:   20,
					},
					Legend: LegendOption{
						Data: labels,
						Show: FalseFlag(),
					},
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"1000\" height=\"900\">\\n<path  d=\"M 0 0\nL 960 0\nL 960 860\nL 0 860\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><text x=\"822\" y=\"55\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Test with 100 labels</text><path  d=\"M 500 452\nL 699 440\nA 200 200 3.60 0 1 700 452\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 699 446\nL 714 446\nM 714 446\nL 729 446\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"732\" y=\"451\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 25: 1%</text><path  d=\"M 500 452\nL 698 427\nA 200 200 3.60 0 1 699 440\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 699 434\nL 714 430\nM 714 430\nL 729 430\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"732\" y=\"435\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 24: 1%</text><path  d=\"M 500 452\nL 696 415\nA 200 200 3.60 0 1 698 427\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 697 421\nL 712 414\nM 712 414\nL 727 414\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"730\" y=\"419\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 23: 1%</text><path  d=\"M 500 452\nL 693 403\nA 200 200 3.60 0 1 696 415\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 695 409\nL 709 398\nM 709 398\nL 724 398\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"727\" y=\"403\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 22: 1%</text><path  d=\"M 500 452\nL 690 391\nA 200 200 3.60 0 1 693 403\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 692 397\nL 706 382\nM 706 382\nL 721 382\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"724\" y=\"387\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 21: 1%</text><path  d=\"M 500 452\nL 685 379\nA 200 200 3.60 0 1 690 391\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 688 385\nL 702 366\nM 702 366\nL 717 366\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"720\" y=\"371\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 20: 1%</text><path  d=\"M 500 452\nL 680 367\nA 200 200 3.60 0 1 685 379\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 683 373\nL 697 350\nM 697 350\nL 712 350\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"715\" y=\"355\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 19: 1%</text><path  d=\"M 500 452\nL 675 356\nA 200 200 3.60 0 1 680 367\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 678 362\nL 691 334\nM 691 334\nL 706 334\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"709\" y=\"339\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 18: 1%</text><path  d=\"M 500 452\nL 668 345\nA 200 200 3.60 0 1 675 356\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 672 351\nL 685 318\nM 685 318\nL 700 318\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"703\" y=\"323\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 17: 1%</text><path  d=\"M 500 452\nL 661 335\nA 200 200 3.60 0 1 668 345\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 665 340\nL 677 302\nM 677 302\nL 692 302\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"695\" y=\"307\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 16: 1%</text><path  d=\"M 500 452\nL 654 325\nA 200 200 3.60 0 1 661 335\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 658 330\nL 669 286\nM 669 286\nL 684 286\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"687\" y=\"291\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 15: 1%</text><path  d=\"M 500 452\nL 645 316\nA 200 200 3.60 0 1 654 325\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 650 320\nL 661 270\nM 661 270\nL 676 270\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"679\" y=\"275\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 14: 1%</text><path  d=\"M 500 452\nL 636 307\nA 200 200 3.60 0 1 645 316\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 641 311\nL 652 254\nM 652 254\nL 667 254\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"670\" y=\"259\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 13: 1%</text><path  d=\"M 500 452\nL 627 298\nA 200 200 3.60 0 1 636 307\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 632 302\nL 642 238\nM 642 238\nL 657 238\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"660\" y=\"243\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 12: 1%</text><path  d=\"M 500 452\nL 617 291\nA 200 200 3.60 0 1 627 298\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 622 294\nL 631 222\nM 631 222\nL 646 222\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"649\" y=\"227\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 11: 1%</text><path  d=\"M 500 452\nL 607 284\nA 200 200 3.60 0 1 617 291\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 612 287\nL 620 206\nM 620 206\nL 635 206\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"638\" y=\"211\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 10: 1%</text><path  d=\"M 500 452\nL 596 277\nA 200 200 3.60 0 1 607 284\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 601 280\nL 609 190\nM 609 190\nL 624 190\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"627\" y=\"195\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 9: 1%</text><path  d=\"M 500 452\nL 585 272\nA 200 200 3.60 0 1 596 277\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 590 274\nL 597 174\nM 597 174\nL 612 174\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"615\" y=\"179\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 8: 1%</text><path  d=\"M 500 452\nL 573 267\nA 200 200 3.60 0 1 585 272\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 579 269\nL 585 158\nM 585 158\nL 600 158\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"603\" y=\"163\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 7: 1%</text><path  d=\"M 500 452\nL 561 262\nA 200 200 3.60 0 1 573 267\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 567 264\nL 572 142\nM 572 142\nL 587 142\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"590\" y=\"147\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 6: 1%</text><path  d=\"M 500 452\nL 549 259\nA 200 200 3.60 0 1 561 262\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 555 260\nL 559 126\nM 559 126\nL 574 126\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"577\" y=\"131\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 5: 1%</text><path  d=\"M 500 452\nL 537 256\nA 200 200 3.60 0 1 549 259\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 543 257\nL 546 110\nM 546 110\nL 561 110\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"564\" y=\"115\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 4: 1%</text><path  d=\"M 500 452\nL 525 254\nA 200 200 3.60 0 1 537 256\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 531 255\nL 533 94\nM 533 94\nL 548 94\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"551\" y=\"99\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 3: 1%</text><path  d=\"M 500 452\nL 512 253\nA 200 200 3.60 0 1 525 254\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 518 253\nL 520 78\nM 520 78\nL 535 78\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"538\" y=\"83\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 2: 1%</text><path  d=\"M 500 452\nL 500 252\nA 200 200 3.60 0 1 512 253\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 506 253\nL 506 62\nM 506 62\nL 521 62\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"524\" y=\"67\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 1: 1%</text><path  d=\"M 500 452\nL 700 452\nA 200 200 3.60 0 1 699 464\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 699 458\nL 714 462\nM 714 462\nL 729 462\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"732\" y=\"467\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 26: 1%</text><path  d=\"M 500 452\nL 699 464\nA 200 200 3.60 0 1 698 477\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 699 470\nL 714 478\nM 714 478\nL 729 478\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"732\" y=\"483\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 27: 1%</text><path  d=\"M 500 452\nL 698 477\nA 200 200 3.60 0 1 696 489\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 697 483\nL 712 494\nM 712 494\nL 727 494\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"730\" y=\"499\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 28: 1%</text><path  d=\"M 500 452\nL 696 489\nA 200 200 3.60 0 1 693 501\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 695 495\nL 709 510\nM 709 510\nL 724 510\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"727\" y=\"515\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 29: 1%</text><path  d=\"M 500 452\nL 693 501\nA 200 200 3.60 0 1 690 513\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 692 507\nL 706 526\nM 706 526\nL 721 526\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"724\" y=\"531\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 30: 1%</text><path  d=\"M 500 452\nL 690 513\nA 200 200 3.60 0 1 685 525\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 688 519\nL 702 542\nM 702 542\nL 717 542\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"720\" y=\"547\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 31: 1%</text><path  d=\"M 500 452\nL 685 525\nA 200 200 3.60 0 1 680 537\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 683 531\nL 697 558\nM 697 558\nL 712 558\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"715\" y=\"563\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 32: 1%</text><path  d=\"M 500 452\nL 680 537\nA 200 200 3.60 0 1 675 548\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 678 542\nL 691 574\nM 691 574\nL 706 574\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"709\" y=\"579\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 33: 1%</text><path  d=\"M 500 452\nL 675 548\nA 200 200 3.60 0 1 668 559\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 672 553\nL 685 590\nM 685 590\nL 700 590\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"703\" y=\"595\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 34: 1%</text><path  d=\"M 500 452\nL 668 559\nA 200 200 3.60 0 1 661 569\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 665 564\nL 677 606\nM 677 606\nL 692 606\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"695\" y=\"611\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 35: 1%</text><path  d=\"M 500 452\nL 661 569\nA 200 200 3.60 0 1 654 579\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 658 574\nL 669 622\nM 669 622\nL 684 622\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"687\" y=\"627\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 36: 1%</text><path  d=\"M 500 452\nL 654 579\nA 200 200 3.60 0 1 645 588\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 650 584\nL 661 638\nM 661 638\nL 676 638\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"679\" y=\"643\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 37: 1%</text><path  d=\"M 500 452\nL 645 588\nA 200 200 3.60 0 1 636 597\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 641 593\nL 652 654\nM 652 654\nL 667 654\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"670\" y=\"659\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 38: 1%</text><path  d=\"M 500 452\nL 636 597\nA 200 200 3.60 0 1 627 606\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 632 602\nL 642 670\nM 642 670\nL 657 670\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"660\" y=\"675\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 39: 1%</text><path  d=\"M 500 452\nL 627 606\nA 200 200 3.60 0 1 617 613\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 622 610\nL 631 686\nM 631 686\nL 646 686\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"649\" y=\"691\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 40: 1%</text><path  d=\"M 500 452\nL 617 613\nA 200 200 3.60 0 1 607 620\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 612 617\nL 620 702\nM 620 702\nL 635 702\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"638\" y=\"707\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 41: 1%</text><path  d=\"M 500 452\nL 607 620\nA 200 200 3.60 0 1 596 627\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 601 624\nL 609 718\nM 609 718\nL 624 718\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"627\" y=\"723\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 42: 1%</text><path  d=\"M 500 452\nL 596 627\nA 200 200 3.60 0 1 585 632\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 590 630\nL 597 734\nM 597 734\nL 612 734\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"615\" y=\"739\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 43: 1%</text><path  d=\"M 500 452\nL 585 632\nA 200 200 3.60 0 1 573 637\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 579 635\nL 585 750\nM 585 750\nL 600 750\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"603\" y=\"755\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 44: 1%</text><path  d=\"M 500 452\nL 573 637\nA 200 200 3.60 0 1 561 642\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 567 640\nL 572 766\nM 572 766\nL 587 766\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"590\" y=\"771\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 45: 1%</text><path  d=\"M 500 452\nL 561 642\nA 200 200 3.60 0 1 549 645\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 555 644\nL 559 782\nM 559 782\nL 574 782\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"577\" y=\"787\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 46: 1%</text><path  d=\"M 500 452\nL 549 645\nA 200 200 3.60 0 1 537 648\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 543 647\nL 546 798\nM 546 798\nL 561 798\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"564\" y=\"803\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 47: 1%</text><path  d=\"M 500 452\nL 537 648\nA 200 200 3.60 0 1 525 650\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 531 649\nL 533 814\nM 533 814\nL 548 814\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"551\" y=\"819\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 48: 1%</text><path  d=\"M 500 452\nL 525 650\nA 200 200 3.60 0 1 512 651\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 518 651\nL 520 830\nM 520 830\nL 535 830\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"538\" y=\"835\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 49: 1%</text><path  d=\"M 500 452\nL 512 651\nA 200 200 3.60 0 1 500 652\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 506 651\nL 506 846\nM 506 846\nL 521 846\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"524\" y=\"851\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 50: 1%</text><path  d=\"M 500 452\nL 301 464\nA 200 200 3.60 0 1 300 452\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 301 458\nL 286 458\nM 286 458\nL 271 458\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"195\" y=\"463\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 75: 1%</text><path  d=\"M 500 452\nL 302 477\nA 200 200 3.60 0 1 301 464\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 301 470\nL 286 474\nM 286 474\nL 271 474\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"195\" y=\"479\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 74: 1%</text><path  d=\"M 500 452\nL 304 489\nA 200 200 3.60 0 1 302 477\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 303 483\nL 288 490\nM 288 490\nL 273 490\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"197\" y=\"495\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 73: 1%</text><path  d=\"M 500 452\nL 307 501\nA 200 200 3.60 0 1 304 489\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 305 495\nL 291 506\nM 291 506\nL 276 506\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"200\" y=\"511\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 72: 1%</text><path  d=\"M 500 452\nL 310 513\nA 200 200 3.60 0 1 307 501\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 308 507\nL 294 522\nM 294 522\nL 279 522\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"203\" y=\"527\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 71: 1%</text><path  d=\"M 500 452\nL 315 525\nA 200 200 3.60 0 1 310 513\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 312 519\nL 298 538\nM 298 538\nL 283 538\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"207\" y=\"543\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 70: 1%</text><path  d=\"M 500 452\nL 320 537\nA 200 200 3.60 0 1 315 525\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 317 531\nL 303 554\nM 303 554\nL 288 554\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"212\" y=\"559\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 69: 1%</text><path  d=\"M 500 452\nL 325 548\nA 200 200 3.60 0 1 320 537\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 322 542\nL 309 570\nM 309 570\nL 294 570\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"218\" y=\"575\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 68: 1%</text><path  d=\"M 500 452\nL 332 559\nA 200 200 3.60 0 1 325 548\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 328 553\nL 315 586\nM 315 586\nL 300 586\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"224\" y=\"591\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 67: 1%</text><path  d=\"M 500 452\nL 339 569\nA 200 200 3.60 0 1 332 559\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 335 564\nL 323 602\nM 323 602\nL 308 602\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"232\" y=\"607\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 66: 1%</text><path  d=\"M 500 452\nL 346 579\nA 200 200 3.60 0 1 339 569\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 342 574\nL 331 618\nM 331 618\nL 316 618\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"240\" y=\"623\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 65: 1%</text><path  d=\"M 500 452\nL 355 588\nA 200 200 3.60 0 1 346 579\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 350 584\nL 339 634\nM 339 634\nL 324 634\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"248\" y=\"639\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 64: 1%</text><path  d=\"M 500 452\nL 364 597\nA 200 200 3.60 0 1 355 588\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 359 593\nL 348 650\nM 348 650\nL 333 650\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"257\" y=\"655\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 63: 1%</text><path  d=\"M 500 452\nL 373 606\nA 200 200 3.60 0 1 364 597\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 368 602\nL 358 666\nM 358 666\nL 343 666\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"267\" y=\"671\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 62: 1%</text><path  d=\"M 500 452\nL 383 613\nA 200 200 3.60 0 1 373 606\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 378 610\nL 369 682\nM 369 682\nL 354 682\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"278\" y=\"687\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 61: 1%</text><path  d=\"M 500 452\nL 393 620\nA 200 200 3.60 0 1 383 613\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 388 617\nL 380 698\nM 380 698\nL 365 698\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"289\" y=\"703\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 60: 1%</text><path  d=\"M 500 452\nL 404 627\nA 200 200 3.60 0 1 393 620\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 399 624\nL 391 714\nM 391 714\nL 376 714\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"300\" y=\"719\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 59: 1%</text><path  d=\"M 500 452\nL 415 632\nA 200 200 3.60 0 1 404 627\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 410 630\nL 403 730\nM 403 730\nL 388 730\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"312\" y=\"735\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 58: 1%</text><path  d=\"M 500 452\nL 427 637\nA 200 200 3.60 0 1 415 632\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 421 635\nL 415 746\nM 415 746\nL 400 746\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"324\" y=\"751\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 57: 1%</text><path  d=\"M 500 452\nL 439 642\nA 200 200 3.60 0 1 427 637\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 433 640\nL 428 762\nM 428 762\nL 413 762\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"337\" y=\"767\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 56: 1%</text><path  d=\"M 500 452\nL 451 645\nA 200 200 3.60 0 1 439 642\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 445 644\nL 441 778\nM 441 778\nL 426 778\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"350\" y=\"783\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 55: 1%</text><path  d=\"M 500 452\nL 463 648\nA 200 200 3.60 0 1 451 645\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 457 647\nL 454 794\nM 454 794\nL 439 794\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"363\" y=\"799\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 54: 1%</text><path  d=\"M 500 452\nL 475 650\nA 200 200 3.60 0 1 463 648\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 469 649\nL 467 810\nM 467 810\nL 452 810\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"376\" y=\"815\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 53: 1%</text><path  d=\"M 500 452\nL 488 651\nA 200 200 3.60 0 1 475 650\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 482 651\nL 480 826\nM 480 826\nL 465 826\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"389\" y=\"831\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 52: 1%</text><path  d=\"M 500 452\nL 500 652\nA 200 200 3.60 0 1 488 651\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 494 651\nL 494 842\nM 494 842\nL 479 842\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"403\" y=\"847\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 51: 1%</text><path  d=\"M 500 452\nL 300 452\nA 200 200 3.60 0 1 301 440\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 301 446\nL 286 442\nM 286 442\nL 271 442\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"195\" y=\"447\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 76: 1%</text><path  d=\"M 500 452\nL 301 440\nA 200 200 3.60 0 1 302 427\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 301 434\nL 286 426\nM 286 426\nL 271 426\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"195\" y=\"431\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 77: 1%</text><path  d=\"M 500 452\nL 302 427\nA 200 200 3.60 0 1 304 415\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 303 421\nL 288 410\nM 288 410\nL 273 410\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"197\" y=\"415\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 78: 1%</text><path  d=\"M 500 452\nL 304 415\nA 200 200 3.60 0 1 307 403\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 305 409\nL 291 394\nM 291 394\nL 276 394\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"200\" y=\"399\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 79: 1%</text><path  d=\"M 500 452\nL 307 403\nA 200 200 3.60 0 1 310 391\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 308 397\nL 294 378\nM 294 378\nL 279 378\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"203\" y=\"383\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 80: 1%</text><path  d=\"M 500 452\nL 310 391\nA 200 200 3.60 0 1 315 379\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 312 385\nL 298 362\nM 298 362\nL 283 362\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"207\" y=\"367\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 81: 1%</text><path  d=\"M 500 452\nL 315 379\nA 200 200 3.60 0 1 320 367\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 317 373\nL 303 346\nM 303 346\nL 288 346\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"212\" y=\"351\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 82: 1%</text><path  d=\"M 500 452\nL 320 367\nA 200 200 3.60 0 1 325 356\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 322 362\nL 309 330\nM 309 330\nL 294 330\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"218\" y=\"335\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 83: 1%</text><path  d=\"M 500 452\nL 325 356\nA 200 200 3.60 0 1 332 345\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 328 351\nL 315 314\nM 315 314\nL 300 314\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"224\" y=\"319\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 84: 1%</text><path  d=\"M 500 452\nL 332 345\nA 200 200 3.60 0 1 339 335\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 335 340\nL 323 298\nM 323 298\nL 308 298\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"232\" y=\"303\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 85: 1%</text><path  d=\"M 500 452\nL 339 335\nA 200 200 3.60 0 1 346 325\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 342 330\nL 331 282\nM 331 282\nL 316 282\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"240\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 86: 1%</text><path  d=\"M 500 452\nL 346 325\nA 200 200 3.60 0 1 355 316\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 350 320\nL 339 266\nM 339 266\nL 324 266\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"248\" y=\"271\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 87: 1%</text><path  d=\"M 500 452\nL 355 316\nA 200 200 3.60 0 1 364 307\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 359 311\nL 348 250\nM 348 250\nL 333 250\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"257\" y=\"255\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 88: 1%</text><path  d=\"M 500 452\nL 364 307\nA 200 200 3.60 0 1 373 298\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 368 302\nL 358 234\nM 358 234\nL 343 234\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"267\" y=\"239\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 89: 1%</text><path  d=\"M 500 452\nL 373 298\nA 200 200 3.60 0 1 383 291\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 378 294\nL 369 218\nM 369 218\nL 354 218\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"278\" y=\"223\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 90: 1%</text><path  d=\"M 500 452\nL 383 291\nA 200 200 3.60 0 1 393 284\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 388 287\nL 380 202\nM 380 202\nL 365 202\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"289\" y=\"207\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 91: 1%</text><path  d=\"M 500 452\nL 393 284\nA 200 200 3.60 0 1 404 277\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 399 280\nL 391 186\nM 391 186\nL 376 186\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"300\" y=\"191\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 92: 1%</text><path  d=\"M 500 452\nL 404 277\nA 200 200 3.60 0 1 415 272\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 410 274\nL 403 170\nM 403 170\nL 388 170\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"312\" y=\"175\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 93: 1%</text><path  d=\"M 500 452\nL 415 272\nA 200 200 3.60 0 1 427 267\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 421 269\nL 415 154\nM 415 154\nL 400 154\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"324\" y=\"159\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 94: 1%</text><path  d=\"M 500 452\nL 427 267\nA 200 200 3.60 0 1 439 262\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 433 264\nL 428 138\nM 428 138\nL 413 138\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"337\" y=\"143\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 95: 1%</text><path  d=\"M 500 452\nL 439 262\nA 200 200 3.60 0 1 451 259\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 445 260\nL 441 122\nM 441 122\nL 426 122\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"350\" y=\"127\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 96: 1%</text><path  d=\"M 500 452\nL 451 259\nA 200 200 3.60 0 1 463 256\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 457 257\nL 454 106\nM 454 106\nL 439 106\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"363\" y=\"111\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 97: 1%</text><path  d=\"M 500 452\nL 463 256\nA 200 200 3.60 0 1 475 254\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 469 255\nL 467 90\nM 467 90\nL 452 90\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"376\" y=\"95\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 98: 1%</text><path  d=\"M 500 452\nL 475 254\nA 200 200 3.60 0 1 488 253\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 482 253\nL 480 74\nM 480 74\nL 465 74\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"389\" y=\"79\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 99: 1%</text><path  d=\"M 500 452\nL 488 253\nA 200 200 3.60 0 1 500 252\nL 500 452\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 494 253\nL 494 58\nM 494 58\nL 479 58\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"396\" y=\"63\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Label 100: 1%</text></svg>",
		},
	}
	for _, tt := range tests {
		p, err := NewPainter(PainterOptions{
			Type:   ChartOutputSVG,
			Width:  1000,
			Height: 900,
		}, PainterThemeOption(defaultTheme))
		assert.Nil(err)
		data, err := tt.render(p.Child(PainterPaddingOption(Box{
			Left:   20,
			Top:    20,
			Right:  20,
			Bottom: 20,
		})))
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}

func TestPieChartFixLabelPos72586(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				values := []float64{
					397594,
					185596,
					149086,
					144258,
					120194,
					117514,
					99412,
					91135,
					87282,
					76790,
					72586,
					58818,
					58270,
					56306,
					55486,
					54792,
					53746,
					51460,
					41242,
					39476,
					37414,
					36644,
					33784,
					32788,
					32566,
					29608,
					29558,
					29384,
					28166,
					26998,
					26948,
					26054,
					25804,
					25730,
					24438,
					23782,
					22896,
					21404,
					428978,
				}
				_, err := NewPieChart(p, PieChartOption{
					SeriesList: NewPieSeriesList(values, PieSeriesOption{
						Label: SeriesLabel{
							Show:      true,
							Formatter: "{b} ({c} ≅ {d})",
						},
						Radius: "150",
					}),
					Title: TitleOption{
						Text: "Fix label K (72586)",
						Left: PositionRight,
					},
					Padding: Box{
						Top:    20,
						Right:  20,
						Bottom: 20,
						Left:   20,
					},
					Legend: LegendOption{
						Data: []string{
							"A",
							"B",
							"C",
							"D",
							"E",
							"F",
							"G",
							"H",
							"I",
							"J",
							"K",
							"L",
							"M",
							"N",
							"O",
							"P",
							"Q",
							"R",
							"S",
							"T",
							"U",
							"V",
							"W",
							"X",
							"Y",
							"Z",
							"AA",
							"AB",
							"AC",
							"AD",
							"AE",
							"AF",
							"AG",
							"AH",
							"AI",
							"AJ",
							"AK",
							"AL",
							"AM",
						},
						Show: FalseFlag(),
					},
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"1150\" height=\"550\">\\n<path  d=\"M 0 0\nL 1110 0\nL 1110 510\nL 0 510\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><text x=\"981\" y=\"55\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Fix label K (72586)</text><path  d=\"M 575 277\nL 716 229\nA 150 150 18.17 0 1 724 276\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 722 252\nL 737 249\nM 737 249\nL 752 249\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"755\" y=\"254\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">C (149086 ≅ 5.04%)</text><path  d=\"M 575 277\nL 687 178\nA 150 150 22.62 0 1 716 229\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 704 202\nL 717 194\nM 717 194\nL 732 194\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"735\" y=\"199\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">B (185596 ≅ 6.28%)</text><path  d=\"M 575 277\nL 575 127\nA 150 150 48.45 0 1 687 178\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 636 141\nL 642 127\nM 642 127\nL 657 127\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"660\" y=\"132\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">A (397594 ≅ 13.45%)</text><path  d=\"M 575 277\nL 724 276\nA 150 150 17.58 0 1 718 320\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 723 297\nL 738 300\nM 738 300\nL 753 300\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"756\" y=\"305\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">D (144258 ≅ 4.88%)</text><path  d=\"M 575 277\nL 718 320\nA 150 150 14.65 0 1 702 355\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 711 338\nL 725 344\nM 725 344\nL 740 344\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"743\" y=\"349\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">E (120194 ≅ 4.06%)</text><path  d=\"M 575 277\nL 702 355\nA 150 150 14.32 0 1 679 384\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 692 370\nL 703 380\nM 703 380\nL 718 380\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"721\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">F (117514 ≅ 3.97%)</text><path  d=\"M 575 277\nL 679 384\nA 150 150 12.12 0 1 654 404\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 667 394\nL 676 406\nM 676 406\nL 691 406\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"694\" y=\"411\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">G (99412 ≅ 3.36%)</text><path  d=\"M 575 277\nL 654 404\nA 150 150 11.11 0 1 628 417\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 642 411\nL 648 424\nM 648 424\nL 663 424\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"666\" y=\"429\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">H (91135 ≅ 3.08%)</text><path  d=\"M 575 277\nL 628 417\nA 150 150 10.64 0 1 601 424\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 615 421\nL 619 440\nM 619 440\nL 634 440\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"637\" y=\"445\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">I (87282 ≅ 2.95%)</text><path  d=\"M 575 277\nL 601 424\nA 150 150 9.36 0 1 577 426\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 589 426\nL 591 456\nM 591 456\nL 606 456\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"609\" y=\"461\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">J (76790 ≅ 2.59%)</text><path  d=\"M 575 277\nL 426 286\nA 150 150 3.61 0 1 426 277\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 426 282\nL 411 282\nM 411 282\nL 396 282\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"308\" y=\"287\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Z (29608 ≅ 1%)</text><path  d=\"M 575 277\nL 427 297\nA 150 150 3.97 0 1 426 286\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 426 292\nL 411 298\nM 411 298\nL 396 298\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"297\" y=\"303\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Y (32566 ≅ 1.1%)</text><path  d=\"M 575 277\nL 429 307\nA 150 150 4.00 0 1 427 297\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 428 302\nL 413 314\nM 413 314\nL 398 314\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"298\" y=\"319\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">X (32788 ≅ 1.1%)</text><path  d=\"M 575 277\nL 431 318\nA 150 150 4.12 0 1 429 307\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 430 312\nL 415 330\nM 415 330\nL 400 330\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"290\" y=\"335\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">W (33784 ≅ 1.14%)</text><path  d=\"M 575 277\nL 435 329\nA 150 150 4.47 0 1 431 318\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 433 323\nL 419 346\nM 419 346\nL 404 346\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"297\" y=\"351\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">V (36644 ≅ 1.24%)</text><path  d=\"M 575 277\nL 439 340\nA 150 150 4.56 0 1 435 329\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 437 334\nL 423 362\nM 423 362\nL 408 362\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"301\" y=\"367\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">U (37414 ≅ 1.26%)</text><path  d=\"M 575 277\nL 445 351\nA 150 150 4.81 0 1 439 340\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 442 345\nL 429 378\nM 429 378\nL 414 378\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"307\" y=\"383\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">T (39476 ≅ 1.33%)</text><path  d=\"M 575 277\nL 452 362\nA 150 150 5.03 0 1 445 351\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 449 357\nL 436 394\nM 436 394\nL 421 394\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"314\" y=\"399\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">S (41242 ≅ 1.39%)</text><path  d=\"M 575 277\nL 462 375\nA 150 150 6.27 0 1 452 362\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 457 369\nL 445 410\nM 445 410\nL 430 410\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"323\" y=\"415\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">R (51460 ≅ 1.74%)</text><path  d=\"M 575 277\nL 474 387\nA 150 150 6.55 0 1 462 375\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 468 381\nL 457 426\nM 457 426\nL 442 426\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"334\" y=\"431\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Q (53746 ≅ 1.81%)</text><path  d=\"M 575 277\nL 488 398\nA 150 150 6.68 0 1 474 387\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 481 393\nL 471 442\nM 471 442\nL 456 442\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"349\" y=\"447\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">P (54792 ≅ 1.85%)</text><path  d=\"M 575 277\nL 503 408\nA 150 150 6.76 0 1 488 398\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 495 403\nL 487 458\nM 487 458\nL 472 458\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"364\" y=\"463\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">O (55486 ≅ 1.87%)</text><path  d=\"M 575 277\nL 519 415\nA 150 150 6.86 0 1 503 408\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 511 412\nL 504 474\nM 504 474\nL 489 474\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"388\" y=\"479\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">N (56306 ≅ 1.9%)</text><path  d=\"M 575 277\nL 537 421\nA 150 150 7.10 0 1 519 415\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 528 419\nL 523 490\nM 523 490\nL 508 490\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"398\" y=\"495\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">M (58270 ≅ 1.97%)</text><path  d=\"M 575 277\nL 555 425\nA 150 150 7.17 0 1 537 421\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 546 424\nL 543 506\nM 543 506\nL 528 506\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"422\" y=\"511\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">L (58818 ≅ 1.99%)</text><path  d=\"M 575 277\nL 577 426\nA 150 150 8.85 0 1 555 425\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 567 426\nL 566 522\nM 566 522\nL 551 522\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"444\" y=\"527\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">K (72586 ≅ 2.45%)</text><path  d=\"M 575 277\nL 426 277\nA 150 150 3.60 0 1 426 269\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 426 273\nL 411 266\nM 411 266\nL 396 266\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"298\" y=\"271\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AA (29558 ≅ 1%)</text><path  d=\"M 575 277\nL 426 269\nA 150 150 3.58 0 1 427 259\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 426 264\nL 411 250\nM 411 250\nL 396 250\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"281\" y=\"255\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AB (29384 ≅ 0.99%)</text><path  d=\"M 575 277\nL 427 259\nA 150 150 3.43 0 1 428 250\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 427 255\nL 412 234\nM 412 234\nL 397 234\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"281\" y=\"239\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AC (28166 ≅ 0.95%)</text><path  d=\"M 575 277\nL 428 250\nA 150 150 3.29 0 1 430 242\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 429 246\nL 414 218\nM 414 218\nL 399 218\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"283\" y=\"223\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AD (26998 ≅ 0.91%)</text><path  d=\"M 575 277\nL 430 242\nA 150 150 3.28 0 1 432 234\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 431 238\nL 416 202\nM 416 202\nL 401 202\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"286\" y=\"207\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AE (26948 ≅ 0.91%)</text><path  d=\"M 575 277\nL 432 234\nA 150 150 3.18 0 1 435 226\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 433 230\nL 419 186\nM 419 186\nL 404 186\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"290\" y=\"191\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AF (26054 ≅ 0.88%)</text><path  d=\"M 575 277\nL 435 226\nA 150 150 3.14 0 1 438 218\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><path  d=\"M 436 222\nL 422 170\nM 422 170\nL 407 170\" style=\"stroke-width:1;stroke:rgba(59,162,114,1.0);fill:rgba(59,162,114,1.0)\"/><text x=\"291\" y=\"175\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AG (25804 ≅ 0.87%)</text><path  d=\"M 575 277\nL 438 218\nA 150 150 3.14 0 1 441 211\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><path  d=\"M 439 214\nL 426 154\nM 426 154\nL 411 154\" style=\"stroke-width:1;stroke:rgba(252,132,82,1.0);fill:rgba(252,132,82,1.0)\"/><text x=\"295\" y=\"159\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AH (25730 ≅ 0.87%)</text><path  d=\"M 575 277\nL 441 211\nA 150 150 2.98 0 1 445 204\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><path  d=\"M 443 207\nL 430 138\nM 430 138\nL 415 138\" style=\"stroke-width:1;stroke:rgba(154,96,180,1.0);fill:rgba(154,96,180,1.0)\"/><text x=\"304\" y=\"143\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AI (24438 ≅ 0.82%)</text><path  d=\"M 575 277\nL 445 204\nA 150 150 2.90 0 1 449 197\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><path  d=\"M 447 201\nL 434 122\nM 434 122\nL 419 122\" style=\"stroke-width:1;stroke:rgba(234,124,204,1.0);fill:rgba(234,124,204,1.0)\"/><text x=\"312\" y=\"127\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AJ (23782 ≅ 0.8%)</text><path  d=\"M 575 277\nL 449 197\nA 150 150 2.79 0 1 453 191\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 451 194\nL 438 106\nM 438 106\nL 423 106\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"308\" y=\"111\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AK (22896 ≅ 0.77%)</text><path  d=\"M 575 277\nL 453 191\nA 150 150 2.61 0 1 457 186\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 455 188\nL 443 90\nM 443 90\nL 428 90\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"314\" y=\"95\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AL (21404 ≅ 0.72%)</text><path  d=\"M 575 277\nL 457 186\nA 150 150 52.28 0 1 575 127\nL 575 277\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 509 143\nL 503 74\nM 503 74\nL 488 74\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"355\" y=\"79\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">AM (428978 ≅ 14.52%)</text></svg>",
		},
	}
	for _, tt := range tests {
		p, err := NewPainter(PainterOptions{
			Type:   ChartOutputSVG,
			Width:  1150,
			Height: 550,
		}, PainterThemeOption(defaultTheme))
		assert.Nil(err)
		data, err := tt.render(p.Child(PainterPaddingOption(Box{
			Left:   20,
			Top:    20,
			Right:  20,
			Bottom: 20,
		})))
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}
