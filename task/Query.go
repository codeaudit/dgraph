// automatically generated by the FlatBuffers compiler, do not modify

package task

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Query struct {
	_tab flatbuffers.Table
}

func (rcv *Query) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Query) Attr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Query) Uids(j int) uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint64(a + flatbuffers.UOffsetT(j * 8))
	}
	return 0
}

func (rcv *Query) UidsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Query) Count() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Query) MutateCount(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Query) Offset() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Query) MutateOffset(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Query) AfterUid() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Query) MutateAfterUid(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

func (rcv *Query) GetCount() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Query) MutateGetCount(n uint16) bool {
	return rcv._tab.MutateUint16Slot(14, n)
}

func QueryStart(builder *flatbuffers.Builder) { builder.StartObject(6) }
func QueryAddAttr(builder *flatbuffers.Builder, attr flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(attr), 0) }
func QueryAddUids(builder *flatbuffers.Builder, uids flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(uids), 0) }
func QueryStartUidsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(8, numElems, 8)
}
func QueryAddCount(builder *flatbuffers.Builder, count int32) { builder.PrependInt32Slot(2, count, 0) }
func QueryAddOffset(builder *flatbuffers.Builder, offset int32) { builder.PrependInt32Slot(3, offset, 0) }
func QueryAddAfterUid(builder *flatbuffers.Builder, afterUid uint64) { builder.PrependUint64Slot(4, afterUid, 0) }
func QueryAddGetCount(builder *flatbuffers.Builder, getCount uint16) { builder.PrependUint16Slot(5, getCount, 0) }
func QueryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
