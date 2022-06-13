/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SysexCommandPinStateQuery is the data-structure of this message
type SysexCommandPinStateQuery struct {
	*SysexCommand
	Pin uint8
}

// ISysexCommandPinStateQuery is the corresponding interface of SysexCommandPinStateQuery
type ISysexCommandPinStateQuery interface {
	ISysexCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
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

func (m *SysexCommandPinStateQuery) GetCommandType() uint8 {
	return 0x6D
}

func (m *SysexCommandPinStateQuery) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *SysexCommandPinStateQuery) InitializeParent(parent *SysexCommand) {}

func (m *SysexCommandPinStateQuery) GetParent() *SysexCommand {
	return m.SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *SysexCommandPinStateQuery) GetPin() uint8 {
	return m.Pin
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSysexCommandPinStateQuery factory function for SysexCommandPinStateQuery
func NewSysexCommandPinStateQuery(pin uint8) *SysexCommandPinStateQuery {
	_result := &SysexCommandPinStateQuery{
		Pin:          pin,
		SysexCommand: NewSysexCommand(),
	}
	_result.Child = _result
	return _result
}

func CastSysexCommandPinStateQuery(structType interface{}) *SysexCommandPinStateQuery {
	if casted, ok := structType.(SysexCommandPinStateQuery); ok {
		return &casted
	}
	if casted, ok := structType.(*SysexCommandPinStateQuery); ok {
		return casted
	}
	if casted, ok := structType.(SysexCommand); ok {
		return CastSysexCommandPinStateQuery(casted.Child)
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return CastSysexCommandPinStateQuery(casted.Child)
	}
	return nil
}

func (m *SysexCommandPinStateQuery) GetTypeName() string {
	return "SysexCommandPinStateQuery"
}

func (m *SysexCommandPinStateQuery) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SysexCommandPinStateQuery) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	return lengthInBits
}

func (m *SysexCommandPinStateQuery) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandPinStateQueryParse(readBuffer utils.ReadBuffer, response bool) (*SysexCommandPinStateQuery, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandPinStateQuery"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandPinStateQuery")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pin)
	_pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field")
	}
	pin := _pin

	if closeErr := readBuffer.CloseContext("SysexCommandPinStateQuery"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandPinStateQuery")
	}

	// Create a partially initialized instance
	_child := &SysexCommandPinStateQuery{
		Pin:          pin,
		SysexCommand: &SysexCommand{},
	}
	_child.SysexCommand.Child = _child
	return _child, nil
}

func (m *SysexCommandPinStateQuery) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandPinStateQuery"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandPinStateQuery")
		}

		// Simple Field (pin)
		pin := uint8(m.Pin)
		_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandPinStateQuery"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandPinStateQuery")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandPinStateQuery) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
