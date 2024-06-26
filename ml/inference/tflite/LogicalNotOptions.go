// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LogicalNotOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsLogicalNotOptions(buf []byte, offset flatbuffers.UOffsetT) *LogicalNotOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LogicalNotOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLogicalNotOptions(buf []byte, offset flatbuffers.UOffsetT) *LogicalNotOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LogicalNotOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LogicalNotOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LogicalNotOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func LogicalNotOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func LogicalNotOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
