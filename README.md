# go-charts

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/vicanso/go-charts/blob/master/LICENSE)
[![Build Status](https://github.com/vicanso/go-charts/workflows/Test/badge.svg)](https://github.com/vicanso/go-charts/actions)

[中文](./README_zh.md)

`go-charts` base on [go-chart](https://github.com/wcharczuk/go-chart)，it is simpler way for generating charts, which supports `svg` and `png` format and themes: `light`, `dark`, `grafana` and `ant`.

`Apache ECharts` is popular among Front-end developers, so `go-charts` supports the option of `Apache ECharts`. Developers can generate charts almost the same as `Apache ECharts`. 

Screenshot of common charts, the left part is light theme, the right part is grafana theme.

<p align="center">
    <img src="./assets/go-charts.png" alt="go-charts">
</p>

## Chart Type

These chart types are supported: `line`, `bar`, `pie`, `radar` or `funnel`.

## Example

The example is for `golang option` and `echarts option`, more examples can be found in the `./examples/` directory.

```go
package main

import (
	"os"
	"path/filepath"

	charts "github.com/vicanso/go-charts"
)

func writeFile(file string, buf []byte) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file = filepath.Join(tmpPath, file)
	err = os.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func chartsRender() ([]byte, error) {
	values := [][]float64{
		{
			120,
			132,
			101,
			134,
			90,
			230,
			210,
		},
	}
	p, err := charts.LineRender(
		values,
		charts.TitleTextOptionFunc("Line"),
		charts.XAxisDataOptionFunc([]string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		}),
	)
	if err != nil {
		return nil, err
	}
	return p.Bytes()
}

func echartsRender() ([]byte, error) {
	return charts.RenderEChartsToPNG(`{
		"title": {
			"text": "Line"
		},
		"xAxis": {
			"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
		},
		"series": [
			{
				"data": [150, 230, 224, 218, 135, 147, 260]
			}
		]
	}`)
}

type Render func() ([]byte, error)

func main() {
	m := map[string]Render{
		"charts-line.png":  chartsRender,
		"echarts-line.png": echartsRender,
	}
	for name, fn := range m {
		buf, err := fn()
		if err != nil {
			panic(err)
		}
		err = writeFile(name, buf)
		if err != nil {
			panic(err)
		}
	}
}
```

## ECharts Option

The name with `[]` is new parameter, others are the same as `echarts`.

- `[type]` The canvas type, support `svg` and `png`, default is `svg`
- `[theme]` The theme, support `dark`, `light` and `grafana`, default is `light`
- `[fontFamily]` The font family for chart
- `[padding]` The padding of chart
- `[box]` The canvas box of chart
- `[width]` The width of chart
- `[height]` The height of chart
- `title` Title component, including main title and subtitle
  - `title.text` The main title text, supporting for \n for newlines
  - `title.subtext`Subtitle text, supporting for \n for newlines
  - `title.left` Distance between title component and the left side of the container. Left value can be instant pixel value like 20; it can also be a percentage value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
  - `title.top` Distance between title component and the top side of the container. Top value can be instant pixel value like 20
  - `title.textStyle.color` Text color for title
  - `title.textStyle.fontSize` Text font size for title 
  - `title.textStyle.fontFamily` Text font family for title, it will change the font family for chart
- `xAxis` The x axis in cartesian(rectangular) coordinate. `go-charts` only support one x axis.
  - `xAxis.boundaryGap` The boundary gap on both sides of a coordinate axis. The setting and behavior of category axes and non-category axes are different. If set `null` or `true`, the label appear in the center part of two axis ticks.
  - `xAxis.splitNumber` Number of segments that the axis is split into. Note that this number serves only as a recommendation, and the true segments may be adjusted based on readability
  - `xAxis.data` Category data, only support string array.
- `yAxis` The y axis in cartesian(rectangular) coordinate, it support 2 y axis
  - `yAxis.min` The minimum value of axis. It will be automatically computed to make sure axis tick is equally distributed when not set
  - `yAxis.max` The maximum value of axis. It will be automatically computed to make sure axis tick is equally distributed when not se.
  - `yAxis.axisLabel.formatter` Formatter of axis label, which supports string template: `"formatter": "{value} kg"`
  - `yAxis.axisLine.lineStyle.color` The color for line
- `legend` Legend component
  - `legend.show` Whether to show legend
  - `legend.data` Data array of legend, only support string array: ["Email", "Video Ads"]
  - `legend.align` Legend marker and text aligning. Support `left` and `right`, default is `left`
  - `legend.padding` legend space around content
  - `legend.left` Distance between legend component and the left side of the container. Left value can be instant pixel value like 20; it can also be a percentage value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
  - `legend.top` Distance between legend component and the top side of the container. Top value can be instant pixel value like 20
- `radar` Coordinate for radar charts
  - `radar.indicator` Indicator of radar chart, which is used to assign multiple variables(dimensions) in radar chart
    - `radar.indicator.name` Indicator's name
    - `radar.indicator.max` The maximum value of indicator
    - `radar.indicator.min` The minimum value of indicator, default value is 0.
- `series` The series for chart 
  - `series.name` Series name used for displaying in legend.
  - `series.type` Series type: `line`, `bar`, `pie`, `radar` or `funnel`
  - `series.radius` Radius of Pie chart:`50%`, default is `40%`
  - `series.yAxisIndex` Index of y axis to combine with, which is useful for multiple y axes in one chart
  - `series.label.show` Whether to show label
  - `series.label.distance` Distance to the host graphic element
  - `series.label.color` Label color
  - `series.itemStyle.color` Color for the series's item 
  - `series.markPoint` Mark point in a chart.
  - `series.markPoint.symbolSize` Symbol size, default is `30` 
  - `series.markPoint.data` Data array for mark points, each of which is an object and the type only support `max` and `min`: `[{"type": "max"}, {"type": "min"}]`
  - `series.markLine` Mark line in a chart
  - `series.markPoint.data`  Data array for mark points, each of which is an object and the type only support `max`, `min` and `average`: `[{"type": "max"}, {"type": "min"}, {"type": "average"}]``
  - `series.data` Data array of series, which can be in the following forms:
    - `value` It's a float array: [1.1, 2,3, 5.2]
    - `object` It's a object value array: [{"value": 1048, "name": "Search Engine"},{"value": 735,"name": "Direct"}]
- `[children]` The options of children chart


## Performance

Generate a png chart will be less than 20ms. It's better than using `chrome headless` with `echarts`.

```bash
BenchmarkMultiChartPNGRender-8                78          15216336 ns/op         2298308 B/op       1148 allocs/op
BenchmarkMultiChartSVGRender-8               367           3356325 ns/op        20597282 B/op       3088 allocs/op
```
