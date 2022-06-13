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

// BACnetConstructedDataStateDescription is the data-structure of this message
type BACnetConstructedDataStateDescription struct {
	*BACnetConstructedData
	StateDescription *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataStateDescription is the corresponding interface of BACnetConstructedDataStateDescription
type IBACnetConstructedDataStateDescription interface {
	IBACnetConstructedData
	// GetStateDescription returns StateDescription (property field)
	GetStateDescription() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataStateDescription) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataStateDescription) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STATE_DESCRIPTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataStateDescription) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataStateDescription) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataStateDescription) GetStateDescription() *BACnetApplicationTagCharacterString {
	return m.StateDescription
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStateDescription factory function for BACnetConstructedDataStateDescription
func NewBACnetConstructedDataStateDescription(stateDescription *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataStateDescription {
	_result := &BACnetConstructedDataStateDescription{
		StateDescription:      stateDescription,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataStateDescription(structType interface{}) *BACnetConstructedDataStateDescription {
	if casted, ok := structType.(BACnetConstructedDataStateDescription); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStateDescription); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataStateDescription(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataStateDescription(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataStateDescription) GetTypeName() string {
	return "BACnetConstructedDataStateDescription"
}

func (m *BACnetConstructedDataStateDescription) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataStateDescription) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (stateDescription)
	lengthInBits += m.StateDescription.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataStateDescription) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataStateDescriptionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataStateDescription, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStateDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStateDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (stateDescription)
	if pullErr := readBuffer.PullContext("stateDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for stateDescription")
	}
	_stateDescription, _stateDescriptionErr := BACnetApplicationTagParse(readBuffer)
	if _stateDescriptionErr != nil {
		return nil, errors.Wrap(_stateDescriptionErr, "Error parsing 'stateDescription' field")
	}
	stateDescription := CastBACnetApplicationTagCharacterString(_stateDescription)
	if closeErr := readBuffer.CloseContext("stateDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for stateDescription")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStateDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStateDescription")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataStateDescription{
		StateDescription:      CastBACnetApplicationTagCharacterString(stateDescription),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataStateDescription) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStateDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStateDescription")
		}

		// Simple Field (stateDescription)
		if pushErr := writeBuffer.PushContext("stateDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for stateDescription")
		}
		_stateDescriptionErr := writeBuffer.WriteSerializable(m.StateDescription)
		if popErr := writeBuffer.PopContext("stateDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for stateDescription")
		}
		if _stateDescriptionErr != nil {
			return errors.Wrap(_stateDescriptionErr, "Error serializing 'stateDescription' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStateDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStateDescription")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataStateDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
