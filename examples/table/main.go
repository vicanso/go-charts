package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/vicanso/go-charts/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func writeFile(buf []byte, filename string) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file := filepath.Join(tmpPath, filename)
	err = ioutil.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// charts.SetDefaultTableSetting(charts.TableDarkThemeSetting)
	charts.SetDefaultWidth(810)
	header := []string{
		"Name",
		"Age",
		"Address",
		"Tag",
		"Action",
	}
	data := [][]string{
		{
			"John Brown",
			"32",
			"New York No. 1 Lake Park",
			"nice, developer",
			"Send Mail",
		},
		{
			"Jim Green	",
			"42",
			"London No. 1 Lake Park",
			"wow",
			"Send Mail",
		},
		{
			"Joe Black	",
			"32",
			"Sidney No. 1 Lake Park",
			"cool, teacher",
			"Send Mail",
		},
	}
	spans := map[int]int{
		0: 2,
		1: 1,
		// 设置第三列的span
		2: 3,
		3: 2,
		4: 2,
	}
	p, err := charts.TableRender(
		header,
		data,
		spans,
	)
	if err != nil {
		panic(err)
	}

	buf, err := p.Bytes()
	if err != nil {
		panic(err)
	}
	err = writeFile(buf, "table.png")
	if err != nil {
		panic(err)
	}

	p, err = charts.TableOptionRender(charts.TableChartOption{
		Header: header,
		Data:   data,
		CellTextStyle: func(tc charts.TableCell) *charts.Style {
			row := tc.Row
			column := tc.Column
			style := tc.Style
			if column == 1 && row != 0 {
				age, _ := strconv.Atoi(tc.Text)
				if age < 40 {
					style.FontColor = drawing.ColorGreen
				} else {
					style.FontColor = drawing.ColorRed
				}
				return &style
			}
			return nil
		},
		CellStyle: func(tc charts.TableCell) *charts.Style {
			row := tc.Row
			column := tc.Column
			if row == 2 && column == 1 {
				return &charts.Style{
					FillColor: drawing.ColorBlue,
				}
			}
			if row == 3 && column == 4 {
				return &charts.Style{
					FillColor: drawing.ColorRed.WithAlpha(100),
				}
			}
			return nil
		},
	})
	if err != nil {
		panic(err)
	}

	buf, err = p.Bytes()
	if err != nil {
		panic(err)
	}
	err = writeFile(buf, "table-color.png")
	if err != nil {
		panic(err)
	}
}
