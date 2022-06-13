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

// BACnetConstructedDataCarPosition is the data-structure of this message
type BACnetConstructedDataCarPosition struct {
	*BACnetConstructedData
	CarPosition *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataCarPosition is the corresponding interface of BACnetConstructedDataCarPosition
type IBACnetConstructedDataCarPosition interface {
	IBACnetConstructedData
	// GetCarPosition returns CarPosition (property field)
	GetCarPosition() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataCarPosition) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataCarPosition) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_POSITION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataCarPosition) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataCarPosition) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataCarPosition) GetCarPosition() *BACnetApplicationTagUnsignedInteger {
	return m.CarPosition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCarPosition factory function for BACnetConstructedDataCarPosition
func NewBACnetConstructedDataCarPosition(carPosition *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataCarPosition {
	_result := &BACnetConstructedDataCarPosition{
		CarPosition:           carPosition,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataCarPosition(structType interface{}) *BACnetConstructedDataCarPosition {
	if casted, ok := structType.(BACnetConstructedDataCarPosition); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarPosition); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataCarPosition(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataCarPosition(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataCarPosition) GetTypeName() string {
	return "BACnetConstructedDataCarPosition"
}

func (m *BACnetConstructedDataCarPosition) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataCarPosition) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (carPosition)
	lengthInBits += m.CarPosition.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataCarPosition) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCarPositionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataCarPosition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarPosition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarPosition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (carPosition)
	if pullErr := readBuffer.PullContext("carPosition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for carPosition")
	}
	_carPosition, _carPositionErr := BACnetApplicationTagParse(readBuffer)
	if _carPositionErr != nil {
		return nil, errors.Wrap(_carPositionErr, "Error parsing 'carPosition' field")
	}
	carPosition := CastBACnetApplicationTagUnsignedInteger(_carPosition)
	if closeErr := readBuffer.CloseContext("carPosition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for carPosition")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarPosition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarPosition")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataCarPosition{
		CarPosition:           CastBACnetApplicationTagUnsignedInteger(carPosition),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataCarPosition) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarPosition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarPosition")
		}

		// Simple Field (carPosition)
		if pushErr := writeBuffer.PushContext("carPosition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for carPosition")
		}
		_carPositionErr := writeBuffer.WriteSerializable(m.CarPosition)
		if popErr := writeBuffer.PopContext("carPosition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for carPosition")
		}
		if _carPositionErr != nil {
			return errors.Wrap(_carPositionErr, "Error serializing 'carPosition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarPosition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarPosition")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataCarPosition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
