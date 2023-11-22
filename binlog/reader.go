package binlog

import (
	"bufio"
	"encoding/binary"
	"io"
)

type rs int

const (
	rsStartEventV3 rs = iota
)

type Reader interface {
	ReadEvent() (Event, error)
}

func NewReader(r io.Reader) (Reader, error) {
	return newReaderImpl(bufio.NewReader(r))
}

func newReaderImpl(r io.Reader) (Reader, error) {
	reader := &readerImpl{r: r}
	if !reader.check() {
		return nil, ErrInvalidHeader
	}
	reader.s = rsStartEventV3
	return reader, nil
}

type readerImpl struct {
	r io.Reader
	s rs
}

func (r *readerImpl) ReadEvent() (Event, error) {
	switch r.s {
	case rsStartEventV3:

	}
	return nil, nil
}

func (r *readerImpl) readStartEventV3() (Event, error) {
	binlogVersion, err := r.readInt2()
	if err != nil {
		return nil, err
	}
	mysqlServerVersion, err := r.readFixedString(50)
	if err != nil {
		return nil, err
	}
	createTimestamp, err := r.readInt4()
	if err != nil {
		return nil, err
	}
	return &StartEventV3{
		BinlogVersion:      binlogVersion,
		MysqlServerVersion: mysqlServerVersion,
		CreateTimestamp:    createTimestamp,
	}, nil
}

func (r *readerImpl) readEventHeader() (*EventHeader, error) {
	timestamp, err := r.readInt4()
	if err != nil {
		return nil, err
	}
	eventType, err := r.readInt1()
	if err != nil {
		return nil, err
	}
	serverId, err := r.readInt4()
	if err != nil {
		return nil, err
	}
	eventSize, err := r.readInt4()
	if err != nil {
		return nil, err
	}
	nextPosition, err := r.readInt4()
	if err != nil {
		return nil, err
	}
	flags, err := r.readInt2()
	if err != nil {
		return nil, err
	}
	return &EventHeader{
		Timestamp:    timestamp,
		EventType:    EventType(eventType),
		ServerId:     serverId,
		EventSize:    eventSize,
		NextPosition: nextPosition,
		Flags:        flags,
	}, nil
}

func (r *readerImpl) check() bool {
	bytes := make([]byte, 4)
	_, err := r.r.Read(bytes)
	if err != nil {
		return false
	}
	return string(bytes) == "\xfebin"
}

// readInt1 reads 1 byte as uint8.
func (r *readerImpl) readInt1() (uint8, error) {
	bytes, err := r.readBytes(1)
	if err != nil {
		return 0, err
	}
	return bytes[0], nil
}

// readInt2 reads 2 bytes as uint16.
func (r *readerImpl) readInt2() (uint16, error) {
	bytes, err := r.readBytes(2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(bytes), nil
}

// readInt4 reads 4 bytes as uint32.
func (r *readerImpl) readInt4() (uint32, error) {
	bytes, err := r.readBytes(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(bytes), nil
}

// readInt8 reads 8 bytes as uint64.
func (r *readerImpl) readInt8() (uint64, error) {
	bytes, err := r.readBytes(8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(bytes), nil
}

func (r *readerImpl) readFixedString(n int) (string, error) {
	bytes, err := r.readBytes(n)
	if err != nil {
		return "", err
	}
	// remove null-terminated character
	for i, b := range bytes {
		if b == 0 {
			return string(bytes[:i]), nil
		}
	}
	return string(bytes), nil
}

func (r *readerImpl) readBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := r.r.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
