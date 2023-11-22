package binlog

type Event interface {
	noop()
}

type _event struct {
}

func (_ *_event) noop() {
	panic("do not call")
}

type StartEventV3 struct {
	_event
	BinlogVersion      uint16
	MysqlServerVersion string
	CreateTimestamp    uint32
}

type EventType int8

type EventHeader struct {
	Timestamp    uint32
	EventType    EventType
	ServerId     uint32
	EventSize    uint32
	NextPosition uint32
	Flags        uint16
}
