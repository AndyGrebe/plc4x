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

// BACnetConstructedDataLastRestartReason is the data-structure of this message
type BACnetConstructedDataLastRestartReason struct {
	*BACnetConstructedData
	LastRestartReason *BACnetRestartReasonTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLastRestartReason is the corresponding interface of BACnetConstructedDataLastRestartReason
type IBACnetConstructedDataLastRestartReason interface {
	IBACnetConstructedData
	// GetLastRestartReason returns LastRestartReason (property field)
	GetLastRestartReason() *BACnetRestartReasonTagged
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

func (m *BACnetConstructedDataLastRestartReason) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLastRestartReason) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_RESTART_REASON
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLastRestartReason) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLastRestartReason) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLastRestartReason) GetLastRestartReason() *BACnetRestartReasonTagged {
	return m.LastRestartReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastRestartReason factory function for BACnetConstructedDataLastRestartReason
func NewBACnetConstructedDataLastRestartReason(lastRestartReason *BACnetRestartReasonTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLastRestartReason {
	_result := &BACnetConstructedDataLastRestartReason{
		LastRestartReason:     lastRestartReason,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLastRestartReason(structType interface{}) *BACnetConstructedDataLastRestartReason {
	if casted, ok := structType.(BACnetConstructedDataLastRestartReason); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastRestartReason); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLastRestartReason(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLastRestartReason(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLastRestartReason) GetTypeName() string {
	return "BACnetConstructedDataLastRestartReason"
}

func (m *BACnetConstructedDataLastRestartReason) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLastRestartReason) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lastRestartReason)
	lengthInBits += m.LastRestartReason.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLastRestartReason) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLastRestartReasonParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLastRestartReason, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastRestartReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastRestartReason")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastRestartReason)
	if pullErr := readBuffer.PullContext("lastRestartReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lastRestartReason")
	}
	_lastRestartReason, _lastRestartReasonErr := BACnetRestartReasonTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _lastRestartReasonErr != nil {
		return nil, errors.Wrap(_lastRestartReasonErr, "Error parsing 'lastRestartReason' field")
	}
	lastRestartReason := CastBACnetRestartReasonTagged(_lastRestartReason)
	if closeErr := readBuffer.CloseContext("lastRestartReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lastRestartReason")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastRestartReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastRestartReason")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLastRestartReason{
		LastRestartReason:     CastBACnetRestartReasonTagged(lastRestartReason),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLastRestartReason) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastRestartReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastRestartReason")
		}

		// Simple Field (lastRestartReason)
		if pushErr := writeBuffer.PushContext("lastRestartReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lastRestartReason")
		}
		_lastRestartReasonErr := writeBuffer.WriteSerializable(m.LastRestartReason)
		if popErr := writeBuffer.PopContext("lastRestartReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lastRestartReason")
		}
		if _lastRestartReasonErr != nil {
			return errors.Wrap(_lastRestartReasonErr, "Error serializing 'lastRestartReason' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastRestartReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastRestartReason")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLastRestartReason) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
