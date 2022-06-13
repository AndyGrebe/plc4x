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

// BACnetConstructedDataOperationDirection is the data-structure of this message
type BACnetConstructedDataOperationDirection struct {
	*BACnetConstructedData
	OperationDirection *BACnetEscalatorOperationDirectionTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataOperationDirection is the corresponding interface of BACnetConstructedDataOperationDirection
type IBACnetConstructedDataOperationDirection interface {
	IBACnetConstructedData
	// GetOperationDirection returns OperationDirection (property field)
	GetOperationDirection() *BACnetEscalatorOperationDirectionTagged
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

func (m *BACnetConstructedDataOperationDirection) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataOperationDirection) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OPERATION_DIRECTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataOperationDirection) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataOperationDirection) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataOperationDirection) GetOperationDirection() *BACnetEscalatorOperationDirectionTagged {
	return m.OperationDirection
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOperationDirection factory function for BACnetConstructedDataOperationDirection
func NewBACnetConstructedDataOperationDirection(operationDirection *BACnetEscalatorOperationDirectionTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataOperationDirection {
	_result := &BACnetConstructedDataOperationDirection{
		OperationDirection:    operationDirection,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataOperationDirection(structType interface{}) *BACnetConstructedDataOperationDirection {
	if casted, ok := structType.(BACnetConstructedDataOperationDirection); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOperationDirection); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataOperationDirection(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataOperationDirection(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataOperationDirection) GetTypeName() string {
	return "BACnetConstructedDataOperationDirection"
}

func (m *BACnetConstructedDataOperationDirection) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataOperationDirection) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (operationDirection)
	lengthInBits += m.OperationDirection.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataOperationDirection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOperationDirectionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataOperationDirection, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOperationDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOperationDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (operationDirection)
	if pullErr := readBuffer.PullContext("operationDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for operationDirection")
	}
	_operationDirection, _operationDirectionErr := BACnetEscalatorOperationDirectionTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _operationDirectionErr != nil {
		return nil, errors.Wrap(_operationDirectionErr, "Error parsing 'operationDirection' field")
	}
	operationDirection := CastBACnetEscalatorOperationDirectionTagged(_operationDirection)
	if closeErr := readBuffer.CloseContext("operationDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for operationDirection")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOperationDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOperationDirection")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataOperationDirection{
		OperationDirection:    CastBACnetEscalatorOperationDirectionTagged(operationDirection),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataOperationDirection) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOperationDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOperationDirection")
		}

		// Simple Field (operationDirection)
		if pushErr := writeBuffer.PushContext("operationDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for operationDirection")
		}
		_operationDirectionErr := writeBuffer.WriteSerializable(m.OperationDirection)
		if popErr := writeBuffer.PopContext("operationDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for operationDirection")
		}
		if _operationDirectionErr != nil {
			return errors.Wrap(_operationDirectionErr, "Error serializing 'operationDirection' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOperationDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOperationDirection")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataOperationDirection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
