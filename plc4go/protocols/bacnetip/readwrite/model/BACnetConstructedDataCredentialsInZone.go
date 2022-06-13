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

// BACnetConstructedDataCredentialsInZone is the data-structure of this message
type BACnetConstructedDataCredentialsInZone struct {
	*BACnetConstructedData
	CredentialsInZone []*BACnetDeviceObjectReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataCredentialsInZone is the corresponding interface of BACnetConstructedDataCredentialsInZone
type IBACnetConstructedDataCredentialsInZone interface {
	IBACnetConstructedData
	// GetCredentialsInZone returns CredentialsInZone (property field)
	GetCredentialsInZone() []*BACnetDeviceObjectReference
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

func (m *BACnetConstructedDataCredentialsInZone) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataCredentialsInZone) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CREDENTIALS_IN_ZONE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataCredentialsInZone) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataCredentialsInZone) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataCredentialsInZone) GetCredentialsInZone() []*BACnetDeviceObjectReference {
	return m.CredentialsInZone
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCredentialsInZone factory function for BACnetConstructedDataCredentialsInZone
func NewBACnetConstructedDataCredentialsInZone(credentialsInZone []*BACnetDeviceObjectReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataCredentialsInZone {
	_result := &BACnetConstructedDataCredentialsInZone{
		CredentialsInZone:     credentialsInZone,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataCredentialsInZone(structType interface{}) *BACnetConstructedDataCredentialsInZone {
	if casted, ok := structType.(BACnetConstructedDataCredentialsInZone); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialsInZone); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataCredentialsInZone(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataCredentialsInZone(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataCredentialsInZone) GetTypeName() string {
	return "BACnetConstructedDataCredentialsInZone"
}

func (m *BACnetConstructedDataCredentialsInZone) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataCredentialsInZone) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.CredentialsInZone) > 0 {
		for _, element := range m.CredentialsInZone {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataCredentialsInZone) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCredentialsInZoneParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataCredentialsInZone, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialsInZone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialsInZone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (credentialsInZone)
	if pullErr := readBuffer.PullContext("credentialsInZone", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for credentialsInZone")
	}
	// Terminated array
	credentialsInZone := make([]*BACnetDeviceObjectReference, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'credentialsInZone' field")
			}
			credentialsInZone = append(credentialsInZone, CastBACnetDeviceObjectReference(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("credentialsInZone", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for credentialsInZone")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialsInZone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialsInZone")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataCredentialsInZone{
		CredentialsInZone:     credentialsInZone,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataCredentialsInZone) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialsInZone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialsInZone")
		}

		// Array Field (credentialsInZone)
		if m.CredentialsInZone != nil {
			if pushErr := writeBuffer.PushContext("credentialsInZone", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for credentialsInZone")
			}
			for _, _element := range m.CredentialsInZone {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'credentialsInZone' field")
				}
			}
			if popErr := writeBuffer.PopContext("credentialsInZone", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for credentialsInZone")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialsInZone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialsInZone")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataCredentialsInZone) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
