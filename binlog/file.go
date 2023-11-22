package binlog

import (
	"bufio"
	"os"
)

func NewFileReader(path string) (Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return newReaderImpl(bufio.NewReader(file))
}
