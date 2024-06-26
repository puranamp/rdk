// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SoftmaxOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsSoftmaxOptions(buf []byte, offset flatbuffers.UOffsetT) *SoftmaxOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SoftmaxOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSoftmaxOptions(buf []byte, offset flatbuffers.UOffsetT) *SoftmaxOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SoftmaxOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SoftmaxOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SoftmaxOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SoftmaxOptions) Beta() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *SoftmaxOptions) MutateBeta(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func SoftmaxOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SoftmaxOptionsAddBeta(builder *flatbuffers.Builder, beta float32) {
	builder.PrependFloat32Slot(0, beta, 0.0)
}
func SoftmaxOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
