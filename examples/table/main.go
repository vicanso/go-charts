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

	file := filepath.Join(tmpPath, "table.png")
	err = ioutil.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	p, err := charts.TableRender(charts.TableChartOption{
		Header: []string{
			"Name",
			"Age",
			"Address",
			"Tag",
			"Action",
		},
		Spans: []int{
			1,
			1,
			2,
			1,
			1,
		},
		Data: [][]string{
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
		},
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
