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

	file := filepath.Join(tmpPath, "horizontal-bar-chart.png")
	err = ioutil.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	values := [][]float64{
		{
			18203,
			23489,
			29034,
			104970,
			131744,
			630230,
		},
		{
			19325,
			23438,
			31000,
			121594,
			134141,
			681807,
		},
	}
	p, err := charts.HorizontalBarRender(
		values,
		charts.TitleTextOptionFunc("World Population"),
		charts.PaddingOptionFunc(charts.Box{
			Top:    20,
			Right:  40,
			Bottom: 20,
			Left:   20,
		}),
		charts.LegendLabelsOptionFunc([]string{
			"2011",
			"2012",
		}),
		charts.YAxisDataOptionFunc([]string{
			"Brazil",
			"Indonesia",
			"USA",
			"India",
			"China",
			"World",
		}),
	)
	if err != nil {
		panic(err)
	}

	buf, err := p.Bytes()
	if err != nil {
		panic(err)
	}
	err = writeFile(buf)
	if err != nil {
		panic(err)
	}
}
