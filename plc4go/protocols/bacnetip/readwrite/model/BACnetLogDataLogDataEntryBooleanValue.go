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

// BACnetLogDataLogDataEntryBooleanValue is the data-structure of this message
type BACnetLogDataLogDataEntryBooleanValue struct {
	*BACnetLogDataLogDataEntry
	BooleanValue *BACnetContextTagBoolean
}

// IBACnetLogDataLogDataEntryBooleanValue is the corresponding interface of BACnetLogDataLogDataEntryBooleanValue
type IBACnetLogDataLogDataEntryBooleanValue interface {
	IBACnetLogDataLogDataEntry
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() *BACnetContextTagBoolean
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetLogDataLogDataEntryBooleanValue) InitializeParent(parent *BACnetLogDataLogDataEntry, peekedTagHeader *BACnetTagHeader) {
	m.BACnetLogDataLogDataEntry.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetLogDataLogDataEntryBooleanValue) GetParent() *BACnetLogDataLogDataEntry {
	return m.BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetLogDataLogDataEntryBooleanValue) GetBooleanValue() *BACnetContextTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryBooleanValue factory function for BACnetLogDataLogDataEntryBooleanValue
func NewBACnetLogDataLogDataEntryBooleanValue(booleanValue *BACnetContextTagBoolean, peekedTagHeader *BACnetTagHeader) *BACnetLogDataLogDataEntryBooleanValue {
	_result := &BACnetLogDataLogDataEntryBooleanValue{
		BooleanValue:              booleanValue,
		BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetLogDataLogDataEntryBooleanValue(structType interface{}) *BACnetLogDataLogDataEntryBooleanValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryBooleanValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryBooleanValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetLogDataLogDataEntry); ok {
		return CastBACnetLogDataLogDataEntryBooleanValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntry); ok {
		return CastBACnetLogDataLogDataEntryBooleanValue(casted.Child)
	}
	return nil
}

func (m *BACnetLogDataLogDataEntryBooleanValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryBooleanValue"
}

func (m *BACnetLogDataLogDataEntryBooleanValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLogDataLogDataEntryBooleanValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLogDataLogDataEntryBooleanValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogDataLogDataEntryBooleanValueParse(readBuffer utils.ReadBuffer) (*BACnetLogDataLogDataEntryBooleanValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryBooleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryBooleanValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for booleanValue")
	}
	_booleanValue, _booleanValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field")
	}
	booleanValue := CastBACnetContextTagBoolean(_booleanValue)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for booleanValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryBooleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryBooleanValue")
	}

	// Create a partially initialized instance
	_child := &BACnetLogDataLogDataEntryBooleanValue{
		BooleanValue:              CastBACnetContextTagBoolean(booleanValue),
		BACnetLogDataLogDataEntry: &BACnetLogDataLogDataEntry{},
	}
	_child.BACnetLogDataLogDataEntry.Child = _child
	return _child, nil
}

func (m *BACnetLogDataLogDataEntryBooleanValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryBooleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryBooleanValue")
		}

		// Simple Field (booleanValue)
		if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for booleanValue")
		}
		_booleanValueErr := writeBuffer.WriteSerializable(m.BooleanValue)
		if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for booleanValue")
		}
		if _booleanValueErr != nil {
			return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryBooleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryBooleanValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetLogDataLogDataEntryBooleanValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
