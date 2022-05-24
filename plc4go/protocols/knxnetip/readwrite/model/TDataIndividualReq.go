/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// TDataIndividualReq is the data-structure of this message
type TDataIndividualReq struct {
	*CEMI

	// Arguments.
	Size uint16
}

// ITDataIndividualReq is the corresponding interface of TDataIndividualReq
type ITDataIndividualReq interface {
	ICEMI
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *TDataIndividualReq) GetMessageCode() uint8 {
	return 0x4A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *TDataIndividualReq) InitializeParent(parent *CEMI) {}

func (m *TDataIndividualReq) GetParent() *CEMI {
	return m.CEMI
}

// NewTDataIndividualReq factory function for TDataIndividualReq
func NewTDataIndividualReq(size uint16) *TDataIndividualReq {
	_result := &TDataIndividualReq{
		CEMI: NewCEMI(size),
	}
	_result.Child = _result
	return _result
}

func CastTDataIndividualReq(structType interface{}) *TDataIndividualReq {
	if casted, ok := structType.(TDataIndividualReq); ok {
		return &casted
	}
	if casted, ok := structType.(*TDataIndividualReq); ok {
		return casted
	}
	if casted, ok := structType.(CEMI); ok {
		return CastTDataIndividualReq(casted.Child)
	}
	if casted, ok := structType.(*CEMI); ok {
		return CastTDataIndividualReq(casted.Child)
	}
	return nil
}

func (m *TDataIndividualReq) GetTypeName() string {
	return "TDataIndividualReq"
}

func (m *TDataIndividualReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *TDataIndividualReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *TDataIndividualReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TDataIndividualReqParse(readBuffer utils.ReadBuffer, size uint16) (*TDataIndividualReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TDataIndividualReq"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TDataIndividualReq"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &TDataIndividualReq{
		CEMI: &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child, nil
}

func (m *TDataIndividualReq) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TDataIndividualReq"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("TDataIndividualReq"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *TDataIndividualReq) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}