package binlog

//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType

type Event interface {
	noop()
}

type _event struct {
}

func (_ *_event) noop() {
	panic("do not call")
}

type ChecksumAlgo int8

const (
	ChecksumAlgoCrc32 ChecksumAlgo = 1
)

type EventType int8

const (
	EventTypeUnknown                 EventType = 0
	EventTypeStartEventV3            EventType = 1
	EventTypeQueryEvent              EventType = 2
	EventTypeStopEvent               EventType = 3
	EventTypeRotateEvent             EventType = 4
	EventTypeIntvarEvent             EventType = 5
	EventTypeSlaveEvent              EventType = 7
	EventTypeAppendBlockEvent        EventType = 9
	EventTypeDeleteFileEvent         EventType = 11
	EventTypeRandEvent               EventType = 13
	EventTypeUserVarEvent            EventType = 14
	EventTypeFormatDescriptionEvent  EventType = 15
	EventTypeXidEvent                EventType = 16
	EventTypeBeginLoadQueryEvent     EventType = 17
	EventTypeExecuteLoadQueryEvent   EventType = 18
	EventTypeTableMapEvent           EventType = 19
	EventTypeWriteRowsEventV1        EventType = 23
	EventTypeUpdateRowsEventV1       EventType = 24
	EventTypeDeleteRowsEventV1       EventType = 25
	EventTypeIncidentEvent           EventType = 26
	EventTypeHeartbeatLogEvent       EventType = 27
	EventTypeIgnorableLogEvent       EventType = 28
	EventTypeRowsQueryLogEvent       EventType = 29
	EventTypeWriteRowsEvent          EventType = 30
	EventTypeUpdateRowsEvent         EventType = 31
	EventTypeDeleteRowsEvent         EventType = 32
	EventTypeGtidLogEvent            EventType = 33
	EventTypeAnonymousGtidLogEvent   EventType = 34
	EventTypePreviousGtidsLogEvent   EventType = 35
	EventTypeTransactionContextEvent EventType = 36
	EventTypeViewChangeEvent         EventType = 37
	EventTypeXaPrepareLogEvent       EventType = 38
	EventTypePartialUpdateRowsEvent  EventType = 39
	EventTypeTransactionPayloadEvent EventType = 40
	EventTypeHeartbeatLogEventV2     EventType = 41
)

type EventHeader struct {
	Timestamp    uint32
	EventType    EventType
	ServerId     uint32
	EventSize    uint32
	NextPosition uint32
	Flags        uint16
}

type StartEventV3 struct {
	_event
	BinlogVersion      uint16
	MysqlServerVersion string
	CreateTimestamp    uint32
}

type FormatDescriptionEvent struct {
	_event
	BinlogVersion      uint16
	MysqlServerVersion string
	CreateTimestamp    uint32
	HeaderLength       uint8
	EventTypes         []byte
	ChecksumAlgo       ChecksumAlgo
	Crc32Bytes         []byte
}
