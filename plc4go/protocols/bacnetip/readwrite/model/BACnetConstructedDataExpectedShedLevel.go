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

// BACnetConstructedDataExpectedShedLevel is the data-structure of this message
type BACnetConstructedDataExpectedShedLevel struct {
	*BACnetConstructedData
	ExpectedShedLevel *BACnetShedLevel

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataExpectedShedLevel is the corresponding interface of BACnetConstructedDataExpectedShedLevel
type IBACnetConstructedDataExpectedShedLevel interface {
	IBACnetConstructedData
	// GetExpectedShedLevel returns ExpectedShedLevel (property field)
	GetExpectedShedLevel() *BACnetShedLevel
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

func (m *BACnetConstructedDataExpectedShedLevel) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataExpectedShedLevel) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXPECTED_SHED_LEVEL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataExpectedShedLevel) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataExpectedShedLevel) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataExpectedShedLevel) GetExpectedShedLevel() *BACnetShedLevel {
	return m.ExpectedShedLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataExpectedShedLevel factory function for BACnetConstructedDataExpectedShedLevel
func NewBACnetConstructedDataExpectedShedLevel(expectedShedLevel *BACnetShedLevel, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataExpectedShedLevel {
	_result := &BACnetConstructedDataExpectedShedLevel{
		ExpectedShedLevel:     expectedShedLevel,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataExpectedShedLevel(structType interface{}) *BACnetConstructedDataExpectedShedLevel {
	if casted, ok := structType.(BACnetConstructedDataExpectedShedLevel); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExpectedShedLevel); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataExpectedShedLevel(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataExpectedShedLevel(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataExpectedShedLevel) GetTypeName() string {
	return "BACnetConstructedDataExpectedShedLevel"
}

func (m *BACnetConstructedDataExpectedShedLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataExpectedShedLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (expectedShedLevel)
	lengthInBits += m.ExpectedShedLevel.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataExpectedShedLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataExpectedShedLevelParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataExpectedShedLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExpectedShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExpectedShedLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (expectedShedLevel)
	if pullErr := readBuffer.PullContext("expectedShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for expectedShedLevel")
	}
	_expectedShedLevel, _expectedShedLevelErr := BACnetShedLevelParse(readBuffer)
	if _expectedShedLevelErr != nil {
		return nil, errors.Wrap(_expectedShedLevelErr, "Error parsing 'expectedShedLevel' field")
	}
	expectedShedLevel := CastBACnetShedLevel(_expectedShedLevel)
	if closeErr := readBuffer.CloseContext("expectedShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for expectedShedLevel")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExpectedShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExpectedShedLevel")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataExpectedShedLevel{
		ExpectedShedLevel:     CastBACnetShedLevel(expectedShedLevel),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataExpectedShedLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExpectedShedLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExpectedShedLevel")
		}

		// Simple Field (expectedShedLevel)
		if pushErr := writeBuffer.PushContext("expectedShedLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for expectedShedLevel")
		}
		_expectedShedLevelErr := writeBuffer.WriteSerializable(m.ExpectedShedLevel)
		if popErr := writeBuffer.PopContext("expectedShedLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for expectedShedLevel")
		}
		if _expectedShedLevelErr != nil {
			return errors.Wrap(_expectedShedLevelErr, "Error serializing 'expectedShedLevel' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExpectedShedLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExpectedShedLevel")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataExpectedShedLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
