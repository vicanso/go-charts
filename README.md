# go-charts

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/vicanso/go-charts/blob/master/LICENSE)
[![Build Status](https://github.com/vicanso/go-charts/workflows/Test/badge.svg)](https://github.com/vicanso/go-charts/actions)

`go-charts`基于[go-chart](https://github.com/wcharczuk/go-chart)生成数据图表，支持`svg`与`png`两种方式的输出，支持三种主题`light`, `dark`以及`grafana`。

`Apache ECharts`在前端开发中得到众多开发者的认可，因此`go-charts`提供了兼容`Apache ECharts`的配置参数，简单快捷的生成相似的图表(`svg`或`png`)，方便插入至Email或分享使用。下面为常用的图表截图(主题为light与grafana)：


![go-charts](./assets/go-charts.png)

## 支持图表类型

暂仅支持三种的图表类型：`line`, `bar` 以及 `pie`


## 示例


下面的示例为`go-charts`两种方式的参数配置：golang的参数配置、echarts的JSON配置，输出相同的折线图。
更多的示例参考：`./examples/`目录

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
	d, err := charts.Render(charts.ChartOption{
		Type: charts.ChartOutputPNG,
		Title: charts.TitleOption{
			Text: "Line",
		},
		XAxis: charts.NewXAxisOption([]string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		}),
		SeriesList: charts.SeriesList{
			charts.NewSeriesFromValues([]float64{
				150,
				230,
				224,
				218,
				135,
				147,
				260,
			}),
		},
	})
	if err != nil {
		return nil, err
	}
	return d.Bytes()
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

## ECharts参数说明

标记为[非echarts配置]的参数为新增参数，可根据实际使用场景添加。

- `type` 画布类型，支持`svg`与`png`，默认为`svg`
- `theme` 颜色主题，支持`dark`、`light`以及`grafana`模式，默认为`light`
- `fontFamily` 字体，全局的字体设置[非echarts配置]
- `padding` 图表的内边距，单位px。支持以下几种模式的设置[非echarts配置]
  - `padding: 5` 设置内边距为5
  - `padding: [5, 10]` 设置上下的内边距为 5，左右的内边距为 10
  - `padding:[5, 10, 5, 10]` 分别设置`上右下左`边距
- `box` 图表的区域，以{"left": Int, "right": Int, "top": Int, "bottom": Int}的形式配置，[非echarts配置]
- `width` 画布宽度，默认为600[非echarts配置]
- `height` 画布高度，默认为400[非echarts配置]
- `title` 图表标题，包括标题内容、高度、颜色等
  - `title.text` 标题文本，支持以`\n`的形式换行
  - `title.subtext` 副标题文本，支持以`\n`的形式换行
  - `title.left` 标题与容器左侧的距离，可设置为`left`, `right`, `center`, `20%` 以及 `20` 这样的具体数值
  - `title.top` 标题与容器顶部的距离，暂仅支持具体数值，如`20`
  - `title.textStyle.color` 标题文字颜色
  - `title.textStyle.fontSize` 标题文字字体大小
  - `title.textStyle.fontFamily` 标题文字的字体系列，需要注意此配置是会影响整个图表的字体
- `xAxis` 直角坐标系grid中的x轴，由于go-charts仅支持单一个x轴，因此若参数为数组多个x轴，只使用第一个配置
  - `xAxis.boundaryGap` 坐标轴两边留白策略，仅支持三种设置方式`null`, `true`或者`false`。`null`或`true`时则数据点展示在两个刻度中间
  - `xAxis.splitNumber` 坐标轴的分割段数，需要注意的是这个分割段数只是个预估值，最后实际显示的段数会在这个基础上根据分割后坐标轴刻度显示的易读程度作调整
  - `xAxis.data` x轴的展示文案，暂只支持字符串数组，如["Mon", "Tue"]，其数量需要与展示点一致
- `yAxis` 直角坐标系grid中的y轴，最多支持两个y轴
  - `yAxis.min` 坐标轴刻度最小值，若不设置则自动计算
  - `yAxis.max` 坐标轴刻度最大值，若不设置则自动计算
  - `yAxis.axisLabel.formatter` 刻度标签的内容格式器，如`"formatter": "{value} kg"`
  - `yAxis.axisLine.lineStyle.color` 坐标轴颜色
- `legend` 图表中不同系列的标记
  - `legend.show` 图例是否显示，如果不需要展示需要设置为`false`
  - `legend.data` 图例的数据数组，为字符串数组，如["Email", "Video Ads"]
  - `legend.align` 图例标记和文本的对齐，可设置为`left`或者`right`，默认为标记靠左`left`
  - `legend.padding` legend的padding，配置方式与图表的`padding`一致
  - `legend.left` legend离容器左侧的距离，其值可以为具体的像素值(20)或百分比(20%)、`left`或者`right`
  - `legend.top` legend离容器顶部的距离，暂仅支持数值形式
- `series` 图表的数据项列表
  - `series.name` 图表的名称，与`legend.data`对应，两者只只设置其一
  - `series.type` 图表的展示类型，暂支持`line`, `bar`以及`pie`，需要注意`pie`只能单独使用
  - `series.radius` 饼图的半径值，如`50%`，默认为`40%`
  - `series.yAxisIndex` 该数据项使用的y轴，默认为0，对yAxis的配置对应
  - `series.label.show` 是否显示文本标签(默认为对应的值)
  - `series.label.distance` 距离图形元素的距离
  - `series.label.color` 文本标签的颜色
  - `series.itemStyle.color` 该数据项展示时使用的颜色
  - `series.markPoint` 图表的标注配置
  - `series.markPoint.symbolSize` 标注的大小，默认为30
  - `series.markPoint.data` 标注类型，仅支持数组形式，其类型只支持`max`与`min`，如：`[{"type": "max"}, {"type": "min"}]
  - `series.markLine` 图表的标线配置 
  - `series.markPoint.data` 标线类型，仅支持数组形式，其类型只支持`max`、`min`以及`average`，如：`[{"type": "max"}, {"type": "min"}, {"type": "average"}]
  - `series.data` 数据项对应的数据数组，支持以下形式的数据：
    - `数值` 常用形式，数组数据为浮点数组，如[1.1, 2,3, 5.2]
    - `结构体` pie图表或bar图表中指定样式使用，如[{"value": 1048, "name": "Search Engine"},{"value": 735,"name": "Direct"}]
- `children` 嵌套的子图表参数列表，图表支持嵌套的形式非echarts配置]

## 性能


简单的图表生成PNG在20ms左右，而SVG的性能则更快，性能上比起使用`chrome headless`加载`echarts`图表展示页面再截图生成的方式大幅度提升，满足简单的图表生成需求。

```bash
BenchmarkMultiChartPNGRender-8                78          15216336 ns/op         2298308 B/op       1148 allocs/op
BenchmarkMultiChartSVGRender-8               367           3356325 ns/op        20597282 B/op       3088 allocs/op
```

## 中文字符

默认使用的字符为`roboto`为英文字体库，因此如果需要显示中文字符需要增加中文字体库，`InstallFont`函数可添加对应的字体库，成功添加之后则指定`title.textStyle.fontFamily`即可。
在浏览器中使用`svg`时，如果指定的`fontFamily`不支持中文字符，展示的中文并不会乱码，但是会导致在计算字符宽度等错误。
