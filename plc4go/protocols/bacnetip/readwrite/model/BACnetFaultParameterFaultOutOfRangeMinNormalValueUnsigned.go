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

// BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned is the data-structure of this message
type BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned struct {
	*BACnetFaultParameterFaultOutOfRangeMinNormalValue
	UnsignedValue *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
type IBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned interface {
	IBACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) InitializeParent(parent *BACnetFaultParameterFaultOutOfRangeMinNormalValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValue.OpeningTag = openingTag
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValue.PeekedTagHeader = peekedTagHeader
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValue.ClosingTag = closingTag
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetParent() *BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetUnsignedValue() *BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned factory function for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned(unsignedValue *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned {
	_result := &BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned{
		UnsignedValue: unsignedValue,
		BACnetFaultParameterFaultOutOfRangeMinNormalValue: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned(structType interface{}) *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return CastBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return CastBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unsignedValue")
	}
	_unsignedValue, _unsignedValueErr := BACnetApplicationTagParse(readBuffer)
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field")
	}
	unsignedValue := CastBACnetApplicationTagUnsignedInteger(_unsignedValue)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unsignedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned{
		UnsignedValue: CastBACnetApplicationTagUnsignedInteger(unsignedValue),
		BACnetFaultParameterFaultOutOfRangeMinNormalValue: &BACnetFaultParameterFaultOutOfRangeMinNormalValue{},
	}
	_child.BACnetFaultParameterFaultOutOfRangeMinNormalValue.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
		}

		// Simple Field (unsignedValue)
		if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unsignedValue")
		}
		_unsignedValueErr := writeBuffer.WriteSerializable(m.UnsignedValue)
		if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unsignedValue")
		}
		if _unsignedValueErr != nil {
			return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
