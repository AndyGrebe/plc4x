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

// IdentifyReplyCommandGAVValuesStored is the corresponding interface of IdentifyReplyCommandGAVValuesStored
type IdentifyReplyCommandGAVValuesStored interface {
	IdentifyReplyCommand
	// GetValues returns Values (property field)
	GetValues() []byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _IdentifyReplyCommandGAVValuesStored is the data-structure of this message
type _IdentifyReplyCommandGAVValuesStored struct {
	*_IdentifyReplyCommand
	Values []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandGAVValuesStored) GetAttribute() Attribute {
	return Attribute_GAVValuesStored
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandGAVValuesStored) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandGAVValuesStored) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandGAVValuesStored) GetValues() []byte {
	return m.Values
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandGAVValuesStored factory function for _IdentifyReplyCommandGAVValuesStored
func NewIdentifyReplyCommandGAVValuesStored(values []byte) *_IdentifyReplyCommandGAVValuesStored {
	_result := &_IdentifyReplyCommandGAVValuesStored{
		Values:                values,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandGAVValuesStored(structType interface{}) IdentifyReplyCommandGAVValuesStored {
	if casted, ok := structType.(IdentifyReplyCommandGAVValuesStored); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandGAVValuesStored); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandGAVValuesStored) GetTypeName() string {
	return "IdentifyReplyCommandGAVValuesStored"
}

func (m *_IdentifyReplyCommandGAVValuesStored) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandGAVValuesStored) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Values) > 0 {
		lengthInBits += 8 * uint16(len(m.Values))
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandGAVValuesStored) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandGAVValuesStoredParse(readBuffer utils.ReadBuffer, attribute Attribute) (IdentifyReplyCommandGAVValuesStored, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandGAVValuesStored"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandGAVValuesStored")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (values)
	numberOfBytesvalues := int(uint16(16))
	values, _readArrayErr := readBuffer.ReadByteArray("values", numberOfBytesvalues)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'values' field")
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandGAVValuesStored"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandGAVValuesStored")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandGAVValuesStored{
		Values:                values,
		_IdentifyReplyCommand: &_IdentifyReplyCommand{},
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandGAVValuesStored) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandGAVValuesStored"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandGAVValuesStored")
		}

		// Array Field (values)
		if m.GetValues() != nil {
			// Byte Array field (values)
			_writeArrayErr := writeBuffer.WriteByteArray("values", m.GetValues())
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'values' field")
			}
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandGAVValuesStored"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandGAVValuesStored")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandGAVValuesStored) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
