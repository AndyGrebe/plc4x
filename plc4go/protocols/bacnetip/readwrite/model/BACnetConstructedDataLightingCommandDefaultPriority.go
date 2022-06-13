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

// BACnetConstructedDataLightingCommandDefaultPriority is the data-structure of this message
type BACnetConstructedDataLightingCommandDefaultPriority struct {
	*BACnetConstructedData
	LightingCommandDefaultPriority *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLightingCommandDefaultPriority is the corresponding interface of BACnetConstructedDataLightingCommandDefaultPriority
type IBACnetConstructedDataLightingCommandDefaultPriority interface {
	IBACnetConstructedData
	// GetLightingCommandDefaultPriority returns LightingCommandDefaultPriority (property field)
	GetLightingCommandDefaultPriority() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataLightingCommandDefaultPriority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLightingCommandDefaultPriority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIGHTING_COMMAND_DEFAULT_PRIORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLightingCommandDefaultPriority) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLightingCommandDefaultPriority) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLightingCommandDefaultPriority) GetLightingCommandDefaultPriority() *BACnetApplicationTagUnsignedInteger {
	return m.LightingCommandDefaultPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLightingCommandDefaultPriority factory function for BACnetConstructedDataLightingCommandDefaultPriority
func NewBACnetConstructedDataLightingCommandDefaultPriority(lightingCommandDefaultPriority *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLightingCommandDefaultPriority {
	_result := &BACnetConstructedDataLightingCommandDefaultPriority{
		LightingCommandDefaultPriority: lightingCommandDefaultPriority,
		BACnetConstructedData:          NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLightingCommandDefaultPriority(structType interface{}) *BACnetConstructedDataLightingCommandDefaultPriority {
	if casted, ok := structType.(BACnetConstructedDataLightingCommandDefaultPriority); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingCommandDefaultPriority); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLightingCommandDefaultPriority(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLightingCommandDefaultPriority(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLightingCommandDefaultPriority) GetTypeName() string {
	return "BACnetConstructedDataLightingCommandDefaultPriority"
}

func (m *BACnetConstructedDataLightingCommandDefaultPriority) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLightingCommandDefaultPriority) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lightingCommandDefaultPriority)
	lengthInBits += m.LightingCommandDefaultPriority.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLightingCommandDefaultPriority) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLightingCommandDefaultPriorityParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLightingCommandDefaultPriority, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingCommandDefaultPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingCommandDefaultPriority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lightingCommandDefaultPriority)
	if pullErr := readBuffer.PullContext("lightingCommandDefaultPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lightingCommandDefaultPriority")
	}
	_lightingCommandDefaultPriority, _lightingCommandDefaultPriorityErr := BACnetApplicationTagParse(readBuffer)
	if _lightingCommandDefaultPriorityErr != nil {
		return nil, errors.Wrap(_lightingCommandDefaultPriorityErr, "Error parsing 'lightingCommandDefaultPriority' field")
	}
	lightingCommandDefaultPriority := CastBACnetApplicationTagUnsignedInteger(_lightingCommandDefaultPriority)
	if closeErr := readBuffer.CloseContext("lightingCommandDefaultPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lightingCommandDefaultPriority")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingCommandDefaultPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingCommandDefaultPriority")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLightingCommandDefaultPriority{
		LightingCommandDefaultPriority: CastBACnetApplicationTagUnsignedInteger(lightingCommandDefaultPriority),
		BACnetConstructedData:          &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLightingCommandDefaultPriority) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingCommandDefaultPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingCommandDefaultPriority")
		}

		// Simple Field (lightingCommandDefaultPriority)
		if pushErr := writeBuffer.PushContext("lightingCommandDefaultPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lightingCommandDefaultPriority")
		}
		_lightingCommandDefaultPriorityErr := writeBuffer.WriteSerializable(m.LightingCommandDefaultPriority)
		if popErr := writeBuffer.PopContext("lightingCommandDefaultPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lightingCommandDefaultPriority")
		}
		if _lightingCommandDefaultPriorityErr != nil {
			return errors.Wrap(_lightingCommandDefaultPriorityErr, "Error serializing 'lightingCommandDefaultPriority' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingCommandDefaultPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingCommandDefaultPriority")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLightingCommandDefaultPriority) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
