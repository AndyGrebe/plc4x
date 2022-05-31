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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPriorityValueEnumerated is the data-structure of this message
type BACnetPriorityValueEnumerated struct {
	*BACnetPriorityValue
	EnumeratedValue *BACnetApplicationTagEnumerated

	// Arguments.
	ObjectType BACnetObjectType
}

// IBACnetPriorityValueEnumerated is the corresponding interface of BACnetPriorityValueEnumerated
type IBACnetPriorityValueEnumerated interface {
	IBACnetPriorityValue
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() *BACnetApplicationTagEnumerated
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

func (m *BACnetPriorityValueEnumerated) InitializeParent(parent *BACnetPriorityValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPriorityValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPriorityValueEnumerated) GetParent() *BACnetPriorityValue {
	return m.BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPriorityValueEnumerated) GetEnumeratedValue() *BACnetApplicationTagEnumerated {
	return m.EnumeratedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueEnumerated factory function for BACnetPriorityValueEnumerated
func NewBACnetPriorityValueEnumerated(enumeratedValue *BACnetApplicationTagEnumerated, peekedTagHeader *BACnetTagHeader, objectType BACnetObjectType) *BACnetPriorityValueEnumerated {
	_result := &BACnetPriorityValueEnumerated{
		EnumeratedValue:     enumeratedValue,
		BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectType),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPriorityValueEnumerated(structType interface{}) *BACnetPriorityValueEnumerated {
	if casted, ok := structType.(BACnetPriorityValueEnumerated); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPriorityValueEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPriorityValue); ok {
		return CastBACnetPriorityValueEnumerated(casted.Child)
	}
	if casted, ok := structType.(*BACnetPriorityValue); ok {
		return CastBACnetPriorityValueEnumerated(casted.Child)
	}
	return nil
}

func (m *BACnetPriorityValueEnumerated) GetTypeName() string {
	return "BACnetPriorityValueEnumerated"
}

func (m *BACnetPriorityValueEnumerated) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPriorityValueEnumerated) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (enumeratedValue)
	lengthInBits += m.EnumeratedValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPriorityValueEnumerated) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueEnumeratedParse(readBuffer utils.ReadBuffer, objectType BACnetObjectType) (*BACnetPriorityValueEnumerated, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueEnumerated"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (enumeratedValue)
	if pullErr := readBuffer.PullContext("enumeratedValue"); pullErr != nil {
		return nil, pullErr
	}
	_enumeratedValue, _enumeratedValueErr := BACnetApplicationTagParse(readBuffer)
	if _enumeratedValueErr != nil {
		return nil, errors.Wrap(_enumeratedValueErr, "Error parsing 'enumeratedValue' field")
	}
	enumeratedValue := CastBACnetApplicationTagEnumerated(_enumeratedValue)
	if closeErr := readBuffer.CloseContext("enumeratedValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueEnumerated"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPriorityValueEnumerated{
		EnumeratedValue:     CastBACnetApplicationTagEnumerated(enumeratedValue),
		BACnetPriorityValue: &BACnetPriorityValue{},
	}
	_child.BACnetPriorityValue.Child = _child
	return _child, nil
}

func (m *BACnetPriorityValueEnumerated) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueEnumerated"); pushErr != nil {
			return pushErr
		}

		// Simple Field (enumeratedValue)
		if pushErr := writeBuffer.PushContext("enumeratedValue"); pushErr != nil {
			return pushErr
		}
		_enumeratedValueErr := m.EnumeratedValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("enumeratedValue"); popErr != nil {
			return popErr
		}
		if _enumeratedValueErr != nil {
			return errors.Wrap(_enumeratedValueErr, "Error serializing 'enumeratedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueEnumerated"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPriorityValueEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}