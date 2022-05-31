package main

import (
	"io/ioutil"
	"log"

	charts "github.com/vicanso/go-charts"
)

func echartsRender() ([]byte, error) {
	return charts.RenderEChartsToPNG(`{
		"title": {
			"text": "用户访问次数",
			"textStyle": {
				"fontFamily": "chinese"
			}
		},
		"xAxis": {
			"data": ["周一", "周二", "周三", "周四", "周五", "周六", "周日"]
		},
		"series": [
			{
				"data": [150, 230, 224, 218, 135, 147, 260],
				"label": {
					"show": true
				}
			}
		]
	}`)
}

func main() {
	fontData, err := ioutil.ReadFile("/Users/darcy/Downloads/NotoSansCJKsc-VF.ttf")
	if err != nil {
		log.Fatalln("Error when reading font file:", err)
	}

	if err := charts.InstallFont("chinese", fontData); err != nil {
		log.Fatalln("Error when instaling font:", err)
	}

	fileData, err := echartsRender()

	if err != nil {
		log.Fatalln("Error when rendering image:", err)
	}
	if err := ioutil.WriteFile("chinese.png", fileData, 0644); err != nil {
		log.Fatalln("Error when save image to chinese.png:", err)
	}
}
