// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PredictScoreResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsPredictScoreResponse(buf []byte, offset flatbuffers.UOffsetT) *PredictScoreResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PredictScoreResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPredictScoreResponse(buf []byte, offset flatbuffers.UOffsetT) *PredictScoreResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PredictScoreResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PredictScoreResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PredictScoreResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PredictScoreResponse) Afc() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PredictScoreResponse) MutateAfc(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *PredictScoreResponse) Nfc() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PredictScoreResponse) MutateNfc(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *PredictScoreResponse) Error() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PredictScoreResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func PredictScoreResponseAddAfc(builder *flatbuffers.Builder, afc int32) {
	builder.PrependInt32Slot(0, afc, 0)
}
func PredictScoreResponseAddNfc(builder *flatbuffers.Builder, nfc int32) {
	builder.PrependInt32Slot(1, nfc, 0)
}
func PredictScoreResponseAddError(builder *flatbuffers.Builder, error flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(error), 0)
}
func PredictScoreResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}