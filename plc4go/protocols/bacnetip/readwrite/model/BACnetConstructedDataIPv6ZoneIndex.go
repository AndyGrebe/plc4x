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

// BACnetConstructedDataIPv6ZoneIndex is the data-structure of this message
type BACnetConstructedDataIPv6ZoneIndex struct {
	*BACnetConstructedData
	Ipv6ZoneIndex *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataIPv6ZoneIndex is the corresponding interface of BACnetConstructedDataIPv6ZoneIndex
type IBACnetConstructedDataIPv6ZoneIndex interface {
	IBACnetConstructedData
	// GetIpv6ZoneIndex returns Ipv6ZoneIndex (property field)
	GetIpv6ZoneIndex() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataIPv6ZoneIndex) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataIPv6ZoneIndex) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_ZONE_INDEX
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataIPv6ZoneIndex) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataIPv6ZoneIndex) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataIPv6ZoneIndex) GetIpv6ZoneIndex() *BACnetApplicationTagCharacterString {
	return m.Ipv6ZoneIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6ZoneIndex factory function for BACnetConstructedDataIPv6ZoneIndex
func NewBACnetConstructedDataIPv6ZoneIndex(ipv6ZoneIndex *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataIPv6ZoneIndex {
	_result := &BACnetConstructedDataIPv6ZoneIndex{
		Ipv6ZoneIndex:         ipv6ZoneIndex,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataIPv6ZoneIndex(structType interface{}) *BACnetConstructedDataIPv6ZoneIndex {
	if casted, ok := structType.(BACnetConstructedDataIPv6ZoneIndex); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6ZoneIndex); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPv6ZoneIndex(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPv6ZoneIndex(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataIPv6ZoneIndex) GetTypeName() string {
	return "BACnetConstructedDataIPv6ZoneIndex"
}

func (m *BACnetConstructedDataIPv6ZoneIndex) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataIPv6ZoneIndex) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipv6ZoneIndex)
	lengthInBits += m.Ipv6ZoneIndex.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataIPv6ZoneIndex) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPv6ZoneIndexParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataIPv6ZoneIndex, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6ZoneIndex"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6ZoneIndex)
	if pullErr := readBuffer.PullContext("ipv6ZoneIndex"); pullErr != nil {
		return nil, pullErr
	}
	_ipv6ZoneIndex, _ipv6ZoneIndexErr := BACnetApplicationTagParse(readBuffer)
	if _ipv6ZoneIndexErr != nil {
		return nil, errors.Wrap(_ipv6ZoneIndexErr, "Error parsing 'ipv6ZoneIndex' field")
	}
	ipv6ZoneIndex := CastBACnetApplicationTagCharacterString(_ipv6ZoneIndex)
	if closeErr := readBuffer.CloseContext("ipv6ZoneIndex"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6ZoneIndex"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataIPv6ZoneIndex{
		Ipv6ZoneIndex:         CastBACnetApplicationTagCharacterString(ipv6ZoneIndex),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataIPv6ZoneIndex) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6ZoneIndex"); pushErr != nil {
			return pushErr
		}

		// Simple Field (ipv6ZoneIndex)
		if pushErr := writeBuffer.PushContext("ipv6ZoneIndex"); pushErr != nil {
			return pushErr
		}
		_ipv6ZoneIndexErr := m.Ipv6ZoneIndex.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("ipv6ZoneIndex"); popErr != nil {
			return popErr
		}
		if _ipv6ZoneIndexErr != nil {
			return errors.Wrap(_ipv6ZoneIndexErr, "Error serializing 'ipv6ZoneIndex' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6ZoneIndex"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataIPv6ZoneIndex) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}