# go-charts

`go-charts`主要分为了下几个模块：

- `标题`：图表的标题，包括主副标题，位置为图表的顶部
- `图例`：图表的图例列表，用于标识每个图例对应的颜色与名称信息，默认为图表的顶部，可自定义位置
- `X轴`：图表的x轴，用于折线图、柱状图中，表示每个点对应的时间，位置图表的底部
- `Y轴`：图表的y轴，用于折线图、柱状图中，最多可使用两组y轴（一左一右），默认位置图表的左侧
- `内容`: 图表的内容，折线图、柱状图、饼图等，在图表的中间区域

## 标题

### 常用设置

标题一般仅需要设置主副标题即可，其它的属性均会设置默认值，常用的方式是使用`TitleTextOptionFunc`设置，其中副标题为可选值，方式如下：

```go
    charts.TitleTextOptionFunc("Text", "Subtext"),
```

### 个性化设置

```go
func(opt *charts.ChartOption) {
    opt.Title = charts.TitleOption{
        // 主标题
        Text: "Text",
        // 副标题
        Subtext: "Subtext",
        // 标题左侧位置，可设置为"center"，"right"，数值("20")或百份比("20%")
        Left: charts.PositionRight,
        // 标题顶部位置，只可调为数值
        Top: "20",
        // 主标题文字大小
        FontSize: 14,
        // 副标题文字大小
        SubtextFontSize: 12,
        // 主标题字体颜色
        FontColor: charts.Color{
            R: 100,
            G: 100,
            B: 100,
            A: 255,
        },
        // 副标题字体影响
        SubtextFontColor: charts.Color{
            R: 200,
            G: 200,
            B: 200,
            A: 255,
        },
    }
},
```

### 部分属性个性化设置

```go
charts.TitleTextOptionFunc("Text", "Subtext"),
func(opt *charts.ChartOption) {
    // 修改top的值
    opt.Title.Top = "20"
},
```

## 图例

### 常用设置

图例组件与图表中的数据一一对应，常用仅设置其名称及左侧的值即可(可选)，方式如下：


```go
charts.LegendLabelsOptionFunc([]string{
    "Email",
    "Union Ads",
    "Video Ads",
    "Direct",
    "Search Engine",
}, "50"),
```

### 个性化设置

```go
func(opt *charts.ChartOption) {
    opt.Legend = charts.LegendOption{
        // 图例名称
        Data: []string{
            "Email",
            "Union Ads",
            "Video Ads",
            "Direct",
            "Search Engine",
        },
        // 图例左侧位置，可设置为"center"，"right"，数值("20")或百份比("20%")
        // 如果示例有多行，只影响第一行，而且对于多行的示例，设置"center", "right"无效
        Left: "50",
        // 图例顶部位置，只可调为数值
        Top: "10",
        // 图例图标的位置，默认为左侧，只允许左或右
        Align: charts.AlignRight,
        // 图例排列方式，默认为水平，只允许水平或垂直
        Orient: charts.OrientVertical,
        // 图标类型，提供"rect"与"lineDot"两种类型
        Icon: charts.IconRect,
        // 字体大小
        FontSize: 14,
        // 字体颜色
        FontColor: charts.Color{
            R: 150,
            G: 150,
            B: 150,
            A: 255,
        },
        // 是否展示，如果不需要展示则设置
        // Show: charts.FalseFlag(),
        // 图例区域的padding值
        Padding: charts.Box{
            Top:  10,
            Left: 10,
        },
    }
},
```

### 部分属性个性化设置

```go
charts.LegendLabelsOptionFunc([]string{
    "Email",
    "Union Ads",
    "Video Ads",
    "Direct",
    "Search Engine",
}, "50"),
func(opt *charts.ChartOption) {
    opt.Legend.Top = "10"
},
```

## X轴

### 常用设置

图表中X轴的展示，常用的设置方式是指定数组即可：

```go
charts.XAxisDataOptionFunc([]string{
    "Mon",
    "Tue",
    "Wed",
    "Thu",
    "Fri",
    "Sat",
    "Sun",
}),
```

### 个性化设置

```go
func(opt *charts.ChartOption) {
    opt.XAxis = charts.XAxisOption{
        // X轴内容
        Data: []string{
            "01",
            "02",
            "03",
            "04",
            "05",
            "06",
            "07",
            "08",
            "09",
        },
        // 如果数据点不居中，则设置为false
        BoundaryGap: charts.FalseFlag(),
        // 字体大小
        FontSize: 14,
        // 是否展示，如果不需要展示则设置
        // Show: charts.FalseFlag(),
        // 会根据文本内容以及此值选择适合的分块大小，一般不需要设置
        // SplitNumber: 3,
        // 线条颜色
        StrokeColor: charts.Color{
            R: 200,
            G: 200,
            B: 200,
            A: 255,
        },
        // 文字颜色
        FontColor: charts.Color{
            R: 100,
            G: 100,
            B: 100,
            A: 255,
        },
    }
},
```

### 部分属性个性化设置

```go
charts.XAxisDataOptionFunc([]string{
    "Mon",
    "Tue",
    "Wed",
    "Thu",
    "Fri",
    "Sat",
    "Sun",
}),
func(opt *charts.ChartOption) {
    opt.XAxis.FontColor = charts.Color{
        R: 100,
        G: 100,
        B: 100,
        A: 255,
    },
},
```

## Y轴

图表中的y轴展示的相关数据会根据图表中的数据自动生成适合的值，如果需要自定义，则可自定义以下部分数据：

```go
func(opt *charts.ChartOption) {
    opt.YAxisOptions = []charts.YAxisOption{
        {
            // 字体大小
            FontSize: 16,
            // 字体颜色
            FontColor: charts.Color{
                R: 100,
                G: 100,
                B: 100,
                A: 255,
            },
            // 内容，{value}会替换为对应的值
            Formatter: "{value} ml",
            // Y轴颜色，如果设置此值，会覆盖font color
            Color: charts.Color{
                R: 255,
                G: 0,
                B: 0,
                A: 255,
            },
        },
    }
},
```
