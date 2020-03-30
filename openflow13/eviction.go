package openflow13

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"

	"github.com/contiv/libOpenflow/common"
	"github.com/contiv/libOpenflow/util"
)

const (
	Type_EvictionSet     uint32 = 1925
	Type_EvictionRequest uint32 = 1926
	Type_EvictionReply   uint32 = 1927
)

var (
	_ util.Message = new(EvictionSet)
	_ util.Message = new(EvictionGet)
	_ util.Message = new(EvictionGetReply)
)

type EvictionSet struct {
	header           common.Header
	ExperimenterID   uint32
	ExperimenterType uint32
	TableID          uint8
	EvictionEnable   uint8
	pad2             [6]uint8
}

func (es *EvictionSet) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	es.header.Length = es.Len()
	headerBytes, err := es.header.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(headerBytes)
	if err := binary.Write(&buf, binary.BigEndian, es.ExperimenterID); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.BigEndian, es.ExperimenterType); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.BigEndian, es.TableID); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.BigEndian, es.EvictionEnable); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.BigEndian, es.pad2); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (es *EvictionSet) UnmarshalBinary(data []byte) error {
	if len(data) < int(es.Len()) {
		return fmt.Errorf("the []byte is too short to unmarshal a full Bundlegontrol message")
	}
	err := es.header.UnmarshalBinary(data)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data[es.header.Len():])
	if err = binary.Read(buf, binary.BigEndian, es.ExperimenterID); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, es.ExperimenterType); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, es.TableID); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, es.EvictionEnable); err != nil {
		return err
	}
	// Don't care padding data.
	return nil
}

func (es *EvictionSet) Len() uint16 {
	return es.header.Len() + uint16(unsafe.Sizeof(es.ExperimenterID)+unsafe.Sizeof(es.ExperimenterType)) + 8
}

func NewEvictionControl() *EvictionSet {
	b := new(EvictionSet)
	b.header = NewOfp13Header()
	b.header.Type = Type_Experimenter
	b.ExperimenterID = ONF_EXPERIMENTER_ID
	b.ExperimenterType = Type_EvictionSet
	return b
}

func (es *EvictionSet) SetTableID(id uint8) {
	es.TableID = id
}

type EvictionGet struct {
	header           common.Header
	ExperimenterID   uint32
	ExperimenterType uint32
	TableID          uint8
	pad              [7]uint8
}

func (eg *EvictionGet) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	eg.header.Length = eg.Len()
	headerBytes, err := eg.header.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(headerBytes)
	if err := binary.Write(&buf, binary.BigEndian, eg.ExperimenterID); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.BigEndian, eg.ExperimenterType); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.BigEndian, eg.TableID); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.BigEndian, eg.pad); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (eg *EvictionGet) UnmarshalBinary(data []byte) error {
	if len(data) < int(eg.Len()) {
		return fmt.Errorf("the []byte is too short to unmarshal a full Bundlegontrol message")
	}
	err := eg.header.UnmarshalBinary(data)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data[eg.header.Len():])
	if err = binary.Read(buf, binary.BigEndian, eg.ExperimenterID); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, eg.ExperimenterType); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, eg.TableID); err != nil {
		return err
	}
	// Don't care padding data.
	return nil
}

func (eg *EvictionGet) Len() uint16 {
	return eg.header.Len() + uint16(unsafe.Sizeof(eg.ExperimenterID)+unsafe.Sizeof(eg.ExperimenterType)) + 8
}

type EvictionGetReply struct {
	EvictionSet
}
