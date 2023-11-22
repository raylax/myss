package main

import (
	"errors"
	"github.com/raylax/myss/binlog"
	"io"
)

func main() {
	reader, err := binlog.NewFileReader("_testdata/binlog.000001")
	if err != nil {
		panic(err)
	}
	for {
		_, err := reader.ReadEvent()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
	}
}
