package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vicanso/go-charts/v2"
)

func writeFile(buf []byte) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file := filepath.Join(tmpPath, "chinese-line-chart.png")
	err = os.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// 字体文件需要自行下载
	// https://github.com/googlefonts/noto-cjk
	buf, err := ioutil.ReadFile("./NotoSansSC.ttf")
	if err != nil {
		panic(err)
	}
	err = charts.InstallFont("noto", buf)
	if err != nil {
		panic(err)
	}
	font, _ := charts.GetFont("noto")
	charts.SetDefaultFont(font)

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
		{
			220,
			182,
			191,
			234,
			290,
			330,
			310,
		},
		{
			150,
			232,
			201,
			154,
			190,
			330,
			410,
		},
		{
			320,
			332,
			301,
			334,
			390,
			330,
			320,
		},
		{
			820,
			932,
			901,
			934,
			1290,
			1330,
			1320,
		},
	}
	p, err := charts.LineRender(
		values,
		charts.TitleTextOptionFunc("测试"),
		charts.XAxisDataOptionFunc([]string{
			"星期一",
			"星期二",
			"星期三",
			"星期四",
			"星期五",
			"星期六",
			"星期日",
		}),
		charts.LegendLabelsOptionFunc([]string{
			"邮件",
			"广告",
			"视频广告",
			"直接访问",
			"搜索引擎",
		}, charts.PositionCenter),
	)

	if err != nil {
		panic(err)
	}

	buf, err = p.Bytes()
	if err != nil {
		panic(err)
	}
	err = writeFile(buf)
	if err != nil {
		panic(err)
	}
}
