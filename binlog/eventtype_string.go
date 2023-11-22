// Code generated by "stringer -type=EventType"; DO NOT EDIT.

package binlog

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EventTypeUnknown-0]
	_ = x[EventTypeStartEventV3-1]
	_ = x[EventTypeQueryEvent-2]
	_ = x[EventTypeStopEvent-3]
	_ = x[EventTypeRotateEvent-4]
	_ = x[EventTypeIntvarEvent-5]
	_ = x[EventTypeSlaveEvent-7]
	_ = x[EventTypeAppendBlockEvent-9]
	_ = x[EventTypeDeleteFileEvent-11]
	_ = x[EventTypeRandEvent-13]
	_ = x[EventTypeUserVarEvent-14]
	_ = x[EventTypeFormatDescriptionEvent-15]
	_ = x[EventTypeXidEvent-16]
	_ = x[EventTypeBeginLoadQueryEvent-17]
	_ = x[EventTypeExecuteLoadQueryEvent-18]
	_ = x[EventTypeTableMapEvent-19]
	_ = x[EventTypeWriteRowsEventV1-23]
	_ = x[EventTypeUpdateRowsEventV1-24]
	_ = x[EventTypeDeleteRowsEventV1-25]
	_ = x[EventTypeIncidentEvent-26]
	_ = x[EventTypeHeartbeatLogEvent-27]
	_ = x[EventTypeIgnorableLogEvent-28]
	_ = x[EventTypeRowsQueryLogEvent-29]
	_ = x[EventTypeWriteRowsEvent-30]
	_ = x[EventTypeUpdateRowsEvent-31]
	_ = x[EventTypeDeleteRowsEvent-32]
	_ = x[EventTypeGtidLogEvent-33]
	_ = x[EventTypeAnonymousGtidLogEvent-34]
	_ = x[EventTypePreviousGtidsLogEvent-35]
	_ = x[EventTypeTransactionContextEvent-36]
	_ = x[EventTypeViewChangeEvent-37]
	_ = x[EventTypeXaPrepareLogEvent-38]
	_ = x[EventTypePartialUpdateRowsEvent-39]
	_ = x[EventTypeTransactionPayloadEvent-40]
	_ = x[EventTypeHeartbeatLogEventV2-41]
}

const (
	_EventType_name_0 = "EventTypeUnknownEventTypeStartEventV3EventTypeQueryEventEventTypeStopEventEventTypeRotateEventEventTypeIntvarEvent"
	_EventType_name_1 = "EventTypeSlaveEvent"
	_EventType_name_2 = "EventTypeAppendBlockEvent"
	_EventType_name_3 = "EventTypeDeleteFileEvent"
	_EventType_name_4 = "EventTypeRandEventEventTypeUserVarEventEventTypeFormatDescriptionEventEventTypeXidEventEventTypeBeginLoadQueryEventEventTypeExecuteLoadQueryEventEventTypeTableMapEvent"
	_EventType_name_5 = "EventTypeWriteRowsEventV1EventTypeUpdateRowsEventV1EventTypeDeleteRowsEventV1EventTypeIncidentEventEventTypeHeartbeatLogEventEventTypeIgnorableLogEventEventTypeRowsQueryLogEventEventTypeWriteRowsEventEventTypeUpdateRowsEventEventTypeDeleteRowsEventEventTypeGtidLogEventEventTypeAnonymousGtidLogEventEventTypePreviousGtidsLogEventEventTypeTransactionContextEventEventTypeViewChangeEventEventTypeXaPrepareLogEventEventTypePartialUpdateRowsEventEventTypeTransactionPayloadEventEventTypeHeartbeatLogEventV2"
)

var (
	_EventType_index_0 = [...]uint8{0, 16, 37, 56, 74, 94, 114}
	_EventType_index_4 = [...]uint8{0, 18, 39, 70, 87, 115, 145, 167}
	_EventType_index_5 = [...]uint16{0, 25, 51, 77, 99, 125, 151, 177, 200, 224, 248, 269, 299, 329, 361, 385, 411, 442, 474, 502}
)

func (i EventType) String() string {
	switch {
	case 0 <= i && i <= 5:
		return _EventType_name_0[_EventType_index_0[i]:_EventType_index_0[i+1]]
	case i == 7:
		return _EventType_name_1
	case i == 9:
		return _EventType_name_2
	case i == 11:
		return _EventType_name_3
	case 13 <= i && i <= 19:
		i -= 13
		return _EventType_name_4[_EventType_index_4[i]:_EventType_index_4[i+1]]
	case 23 <= i && i <= 41:
		i -= 23
		return _EventType_name_5[_EventType_index_5[i]:_EventType_index_5[i+1]]
	default:
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
