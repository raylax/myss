package main

import "github.com/raylax/myss/binlog"

func main() {
	reader, err := binlog.NewFileReader("_testdata/binlog.000001", true)
	if err != nil {
		panic(err)
	}
	for {
		e, err := reader.ReadEvent()
		println(e)
		if err != nil {
			panic(err)
		}
	}
}
