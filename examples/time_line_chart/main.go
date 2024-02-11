package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/vicanso/go-charts/v2"
)

func writeFile(buf []byte) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file := filepath.Join(tmpPath, "time-line-chart.png")
	err = os.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	xAxisValue := []string{}
	values := []float64{}
	now := time.Now()
	firstAxis := 0
	for i := 0; i < 300; i++ {
		// 设置首个axis为xx:00的时间点
		if firstAxis == 0 && now.Minute() == 0 {
			firstAxis = i
		}
		xAxisValue = append(xAxisValue, now.Format("15:04"))
		now = now.Add(time.Minute)
		value, _ := rand.Int(rand.Reader, big.NewInt(100))
		values = append(values, float64(value.Int64()))
	}
	p, err := charts.LineRender(
		[][]float64{
			values,
		},
		charts.TitleTextOptionFunc("Line"),
		charts.XAxisDataOptionFunc(xAxisValue, charts.FalseFlag()),
		charts.LegendLabelsOptionFunc([]string{
			"Demo",
		}, "50"),
		func(opt *charts.ChartOption) {
			opt.XAxis.FirstAxis = firstAxis
			// 必须要比计算得来的最小值更大(每60分钟)
			opt.XAxis.SplitNumber = 60
			opt.Legend.Padding = charts.Box{
				Top:    5,
				Bottom: 10,
			}
			opt.SymbolShow = charts.FalseFlag()
			opt.LineStrokeWidth = 1
			opt.ValueFormatter = func(f float64) string {
				return fmt.Sprintf("%.0f", f)
			}
		},
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
