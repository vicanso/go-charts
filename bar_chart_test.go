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

func TestBarChart(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				seriesList := NewSeriesListDataFromValues([][]float64{
					{
						2.0,
						4.9,
						7.0,
						23.2,
						25.6,
						76.7,
						135.6,
						162.2,
						32.6,
						20.0,
						6.4,
						3.3,
					},
					{
						2.6,
						5.9,
						9.0,
						26.4,
						28.7,
						70.7,
						175.6,
						182.2,
						48.7,
						18.8,
						6.0,
						2.3,
					},
				})
				for index := range seriesList {
					seriesList[index].Label.Show = true
				}
				_, err := NewBarChart(p, BarChartOption{
					Padding: Box{
						Left:   10,
						Top:    10,
						Right:  10,
						Bottom: 10,
					},
					SeriesList: seriesList,
					XAxis: NewXAxisOption([]string{
						"Jan",
						"Feb",
						"Mar",
						"Apr",
						"May",
						"Jun",
						"Jul",
						"Aug",
						"Sep",
						"Oct",
						"Nov",
						"Dec",
					}),
					YAxisOptions: NewYAxisOptions([]string{
						"Rainfall",
						"Evaporation",
					}),
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 400\nL 0 400\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><text x=\"10\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">240</text><text x=\"10\" y=\"75\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">200</text><text x=\"10\" y=\"133\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">160</text><text x=\"10\" y=\"192\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">120</text><text x=\"19\" y=\"250\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">80</text><text x=\"19\" y=\"308\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">40</text><text x=\"28\" y=\"367\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">0</text><path  d=\"M 47 10\nL 590 10\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 68\nL 590 68\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 126\nL 590 126\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 185\nL 590 185\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 243\nL 590 243\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 301\nL 590 301\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 365\nL 47 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 182 365\nL 182 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 318 365\nL 318 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 454 365\nL 454 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 590 365\nL 590 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 47 360\nL 590 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><text x=\"101\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Feb</text><text x=\"235\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">May</text><text x=\"372\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Aug</text><text x=\"507\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Nov</text><path  d=\"M 52 358\nL 68 358\nL 68 359\nL 52 359\nL 52 358\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 97 353\nL 113 353\nL 113 359\nL 97 359\nL 97 353\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 142 350\nL 158 350\nL 158 359\nL 142 359\nL 142 350\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 187 327\nL 203 327\nL 203 359\nL 187 359\nL 187 327\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 233 323\nL 249 323\nL 249 359\nL 233 359\nL 233 323\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 278 249\nL 294 249\nL 294 359\nL 278 359\nL 278 249\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 323 163\nL 339 163\nL 339 359\nL 323 359\nL 323 163\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 368 124\nL 384 124\nL 384 359\nL 368 359\nL 368 124\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 414 313\nL 430 313\nL 430 359\nL 414 359\nL 414 313\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 459 331\nL 475 331\nL 475 359\nL 459 359\nL 459 331\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 504 351\nL 520 351\nL 520 359\nL 504 359\nL 504 351\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 549 356\nL 565 356\nL 565 359\nL 549 359\nL 549 356\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 71 357\nL 87 357\nL 87 359\nL 71 359\nL 71 357\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 116 352\nL 132 352\nL 132 359\nL 116 359\nL 116 352\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 161 347\nL 177 347\nL 177 359\nL 161 359\nL 161 347\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 206 322\nL 222 322\nL 222 359\nL 206 359\nL 206 322\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 252 319\nL 268 319\nL 268 359\nL 252 359\nL 252 319\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 297 257\nL 313 257\nL 313 359\nL 297 359\nL 297 257\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 342 104\nL 358 104\nL 358 359\nL 342 359\nL 342 104\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 387 95\nL 403 95\nL 403 359\nL 387 359\nL 387 95\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 433 289\nL 449 289\nL 449 359\nL 433 359\nL 433 289\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 478 333\nL 494 333\nL 494 359\nL 478 359\nL 478 333\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 523 352\nL 539 352\nL 539 359\nL 523 359\nL 523 352\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 568 357\nL 584 357\nL 584 359\nL 568 359\nL 568 357\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><text x=\"57\" y=\"353\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2</text><text x=\"94\" y=\"348\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">4.9</text><text x=\"147\" y=\"345\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">7</text><text x=\"181\" y=\"322\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">23.2</text><text x=\"227\" y=\"318\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">25.6</text><text x=\"272\" y=\"244\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">76.7</text><text x=\"311\" y=\"158\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">135.6</text><text x=\"356\" y=\"119\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">162.2</text><text x=\"408\" y=\"308\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">32.6</text><text x=\"458\" y=\"326\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">20</text><text x=\"501\" y=\"346\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">6.4</text><text x=\"546\" y=\"351\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">3.3</text><text x=\"68\" y=\"352\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2.6</text><text x=\"113\" y=\"347\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">5.9</text><text x=\"166\" y=\"342\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">9</text><text x=\"200\" y=\"317\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">26.4</text><text x=\"246\" y=\"314\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">28.7</text><text x=\"291\" y=\"252\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">70.7</text><text x=\"330\" y=\"99\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">175.6</text><text x=\"375\" y=\"90\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">182.2</text><text x=\"427\" y=\"284\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">48.7</text><text x=\"472\" y=\"328\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">18.8</text><text x=\"528\" y=\"347\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">6</text><text x=\"565\" y=\"352\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2.3</text></svg>",
		},
		{
			render: func(p *Painter) ([]byte, error) {
				seriesList := NewSeriesListDataFromValues([][]float64{
					{
						2.0,
						4.9,
						7.0,
						23.2,
						25.6,
						76.7,
						135.6,
						162.2,
						32.6,
						20.0,
						6.4,
						3.3,
					},
					{
						2.6,
						5.9,
						9.0,
						26.4,
						28.7,
						70.7,
						175.6,
						182.2,
						48.7,
						18.8,
						6.0,
						2.3,
					},
				})
				for index := range seriesList {
					seriesList[index].Label.Show = true
					seriesList[index].RoundRadius = 5
				}
				_, err := NewBarChart(p, BarChartOption{
					Padding: Box{
						Left:   10,
						Top:    10,
						Right:  10,
						Bottom: 10,
					},
					SeriesList: seriesList,
					XAxis: NewXAxisOption([]string{
						"Jan",
						"Feb",
						"Mar",
						"Apr",
						"May",
						"Jun",
						"Jul",
						"Aug",
						"Sep",
						"Oct",
						"Nov",
						"Dec",
					}),
					YAxisOptions: NewYAxisOptions([]string{
						"Rainfall",
						"Evaporation",
					}),
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 400\nL 0 400\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><text x=\"10\" y=\"17\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">240</text><text x=\"10\" y=\"75\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">200</text><text x=\"10\" y=\"133\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">160</text><text x=\"10\" y=\"192\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">120</text><text x=\"19\" y=\"250\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">80</text><text x=\"19\" y=\"308\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">40</text><text x=\"28\" y=\"367\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">0</text><path  d=\"M 47 10\nL 590 10\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 68\nL 590 68\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 126\nL 590 126\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 185\nL 590 185\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 243\nL 590 243\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 301\nL 590 301\" style=\"stroke-width:1;stroke:rgba(224,230,242,1.0);fill:none\"/><path  d=\"M 47 365\nL 47 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 182 365\nL 182 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 318 365\nL 318 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 454 365\nL 454 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 590 365\nL 590 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><path  d=\"M 47 360\nL 590 360\" style=\"stroke-width:1;stroke:rgba(110,112,121,1.0);fill:none\"/><text x=\"101\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Feb</text><text x=\"235\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">May</text><text x=\"372\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Aug</text><text x=\"507\" y=\"385\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Nov</text><path  d=\"M 57 358\nL 63 358\nL 63 358\nA 5 5 90.00 0 1 68 363\nL 68 354\nL 68 354\nA 5 5 90.00 0 1 63 359\nL 57 359\nL 57 359\nA 5 5 90.00 0 1 52 354\nL 52 363\nL 52 363\nA 5 5 90.00 0 1 57 358\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 102 353\nL 108 353\nL 108 353\nA 5 5 90.00 0 1 113 358\nL 113 354\nL 113 354\nA 5 5 90.00 0 1 108 359\nL 102 359\nL 102 359\nA 5 5 90.00 0 1 97 354\nL 97 358\nL 97 358\nA 5 5 90.00 0 1 102 353\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 147 350\nL 153 350\nL 153 350\nA 5 5 90.00 0 1 158 355\nL 158 354\nL 158 354\nA 5 5 90.00 0 1 153 359\nL 147 359\nL 147 359\nA 5 5 90.00 0 1 142 354\nL 142 355\nL 142 355\nA 5 5 90.00 0 1 147 350\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 192 327\nL 198 327\nL 198 327\nA 5 5 90.00 0 1 203 332\nL 203 354\nL 203 354\nA 5 5 90.00 0 1 198 359\nL 192 359\nL 192 359\nA 5 5 90.00 0 1 187 354\nL 187 332\nL 187 332\nA 5 5 90.00 0 1 192 327\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 238 323\nL 244 323\nL 244 323\nA 5 5 90.00 0 1 249 328\nL 249 354\nL 249 354\nA 5 5 90.00 0 1 244 359\nL 238 359\nL 238 359\nA 5 5 90.00 0 1 233 354\nL 233 328\nL 233 328\nA 5 5 90.00 0 1 238 323\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 283 249\nL 289 249\nL 289 249\nA 5 5 90.00 0 1 294 254\nL 294 354\nL 294 354\nA 5 5 90.00 0 1 289 359\nL 283 359\nL 283 359\nA 5 5 90.00 0 1 278 354\nL 278 254\nL 278 254\nA 5 5 90.00 0 1 283 249\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 328 163\nL 334 163\nL 334 163\nA 5 5 90.00 0 1 339 168\nL 339 354\nL 339 354\nA 5 5 90.00 0 1 334 359\nL 328 359\nL 328 359\nA 5 5 90.00 0 1 323 354\nL 323 168\nL 323 168\nA 5 5 90.00 0 1 328 163\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 373 124\nL 379 124\nL 379 124\nA 5 5 90.00 0 1 384 129\nL 384 354\nL 384 354\nA 5 5 90.00 0 1 379 359\nL 373 359\nL 373 359\nA 5 5 90.00 0 1 368 354\nL 368 129\nL 368 129\nA 5 5 90.00 0 1 373 124\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 419 313\nL 425 313\nL 425 313\nA 5 5 90.00 0 1 430 318\nL 430 354\nL 430 354\nA 5 5 90.00 0 1 425 359\nL 419 359\nL 419 359\nA 5 5 90.00 0 1 414 354\nL 414 318\nL 414 318\nA 5 5 90.00 0 1 419 313\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 464 331\nL 470 331\nL 470 331\nA 5 5 90.00 0 1 475 336\nL 475 354\nL 475 354\nA 5 5 90.00 0 1 470 359\nL 464 359\nL 464 359\nA 5 5 90.00 0 1 459 354\nL 459 336\nL 459 336\nA 5 5 90.00 0 1 464 331\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 509 351\nL 515 351\nL 515 351\nA 5 5 90.00 0 1 520 356\nL 520 354\nL 520 354\nA 5 5 90.00 0 1 515 359\nL 509 359\nL 509 359\nA 5 5 90.00 0 1 504 354\nL 504 356\nL 504 356\nA 5 5 90.00 0 1 509 351\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 554 356\nL 560 356\nL 560 356\nA 5 5 90.00 0 1 565 361\nL 565 354\nL 565 354\nA 5 5 90.00 0 1 560 359\nL 554 359\nL 554 359\nA 5 5 90.00 0 1 549 354\nL 549 361\nL 549 361\nA 5 5 90.00 0 1 554 356\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(84,112,198,1.0)\"/><path  d=\"M 76 357\nL 82 357\nL 82 357\nA 5 5 90.00 0 1 87 362\nL 87 354\nL 87 354\nA 5 5 90.00 0 1 82 359\nL 76 359\nL 76 359\nA 5 5 90.00 0 1 71 354\nL 71 362\nL 71 362\nA 5 5 90.00 0 1 76 357\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 121 352\nL 127 352\nL 127 352\nA 5 5 90.00 0 1 132 357\nL 132 354\nL 132 354\nA 5 5 90.00 0 1 127 359\nL 121 359\nL 121 359\nA 5 5 90.00 0 1 116 354\nL 116 357\nL 116 357\nA 5 5 90.00 0 1 121 352\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 166 347\nL 172 347\nL 172 347\nA 5 5 90.00 0 1 177 352\nL 177 354\nL 177 354\nA 5 5 90.00 0 1 172 359\nL 166 359\nL 166 359\nA 5 5 90.00 0 1 161 354\nL 161 352\nL 161 352\nA 5 5 90.00 0 1 166 347\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 211 322\nL 217 322\nL 217 322\nA 5 5 90.00 0 1 222 327\nL 222 354\nL 222 354\nA 5 5 90.00 0 1 217 359\nL 211 359\nL 211 359\nA 5 5 90.00 0 1 206 354\nL 206 327\nL 206 327\nA 5 5 90.00 0 1 211 322\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 257 319\nL 263 319\nL 263 319\nA 5 5 90.00 0 1 268 324\nL 268 354\nL 268 354\nA 5 5 90.00 0 1 263 359\nL 257 359\nL 257 359\nA 5 5 90.00 0 1 252 354\nL 252 324\nL 252 324\nA 5 5 90.00 0 1 257 319\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 302 257\nL 308 257\nL 308 257\nA 5 5 90.00 0 1 313 262\nL 313 354\nL 313 354\nA 5 5 90.00 0 1 308 359\nL 302 359\nL 302 359\nA 5 5 90.00 0 1 297 354\nL 297 262\nL 297 262\nA 5 5 90.00 0 1 302 257\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 347 104\nL 353 104\nL 353 104\nA 5 5 90.00 0 1 358 109\nL 358 354\nL 358 354\nA 5 5 90.00 0 1 353 359\nL 347 359\nL 347 359\nA 5 5 90.00 0 1 342 354\nL 342 109\nL 342 109\nA 5 5 90.00 0 1 347 104\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 392 95\nL 398 95\nL 398 95\nA 5 5 90.00 0 1 403 100\nL 403 354\nL 403 354\nA 5 5 90.00 0 1 398 359\nL 392 359\nL 392 359\nA 5 5 90.00 0 1 387 354\nL 387 100\nL 387 100\nA 5 5 90.00 0 1 392 95\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 438 289\nL 444 289\nL 444 289\nA 5 5 90.00 0 1 449 294\nL 449 354\nL 449 354\nA 5 5 90.00 0 1 444 359\nL 438 359\nL 438 359\nA 5 5 90.00 0 1 433 354\nL 433 294\nL 433 294\nA 5 5 90.00 0 1 438 289\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 483 333\nL 489 333\nL 489 333\nA 5 5 90.00 0 1 494 338\nL 494 354\nL 494 354\nA 5 5 90.00 0 1 489 359\nL 483 359\nL 483 359\nA 5 5 90.00 0 1 478 354\nL 478 338\nL 478 338\nA 5 5 90.00 0 1 483 333\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 528 352\nL 534 352\nL 534 352\nA 5 5 90.00 0 1 539 357\nL 539 354\nL 539 354\nA 5 5 90.00 0 1 534 359\nL 528 359\nL 528 359\nA 5 5 90.00 0 1 523 354\nL 523 357\nL 523 357\nA 5 5 90.00 0 1 528 352\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"M 573 357\nL 579 357\nL 579 357\nA 5 5 90.00 0 1 584 362\nL 584 354\nL 584 354\nA 5 5 90.00 0 1 579 359\nL 573 359\nL 573 359\nA 5 5 90.00 0 1 568 354\nL 568 362\nL 568 362\nA 5 5 90.00 0 1 573 357\nZ\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:0;stroke:none;fill:rgba(145,204,117,1.0)\"/><text x=\"57\" y=\"353\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2</text><text x=\"94\" y=\"348\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">4.9</text><text x=\"147\" y=\"345\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">7</text><text x=\"181\" y=\"322\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">23.2</text><text x=\"227\" y=\"318\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">25.6</text><text x=\"272\" y=\"244\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">76.7</text><text x=\"311\" y=\"158\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">135.6</text><text x=\"356\" y=\"119\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">162.2</text><text x=\"408\" y=\"308\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">32.6</text><text x=\"458\" y=\"326\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">20</text><text x=\"501\" y=\"346\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">6.4</text><text x=\"546\" y=\"351\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">3.3</text><text x=\"68\" y=\"352\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2.6</text><text x=\"113\" y=\"347\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">5.9</text><text x=\"166\" y=\"342\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">9</text><text x=\"200\" y=\"317\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">26.4</text><text x=\"246\" y=\"314\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">28.7</text><text x=\"291\" y=\"252\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">70.7</text><text x=\"330\" y=\"99\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">175.6</text><text x=\"375\" y=\"90\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">182.2</text><text x=\"427\" y=\"284\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">48.7</text><text x=\"472\" y=\"328\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">18.8</text><text x=\"528\" y=\"347\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">6</text><text x=\"565\" y=\"352\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">2.3</text></svg>",
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
		assert.Equal(tt.result, string(data))
	}
}
