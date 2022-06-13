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

// BACnetConstructedDataLogDeviceObjectProperty is the data-structure of this message
type BACnetConstructedDataLogDeviceObjectProperty struct {
	*BACnetConstructedData
	LogDeviceObjectProperty *BACnetDeviceObjectPropertyReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLogDeviceObjectProperty is the corresponding interface of BACnetConstructedDataLogDeviceObjectProperty
type IBACnetConstructedDataLogDeviceObjectProperty interface {
	IBACnetConstructedData
	// GetLogDeviceObjectProperty returns LogDeviceObjectProperty (property field)
	GetLogDeviceObjectProperty() *BACnetDeviceObjectPropertyReference
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

func (m *BACnetConstructedDataLogDeviceObjectProperty) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLogDeviceObjectProperty) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_DEVICE_OBJECT_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLogDeviceObjectProperty) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLogDeviceObjectProperty) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLogDeviceObjectProperty) GetLogDeviceObjectProperty() *BACnetDeviceObjectPropertyReference {
	return m.LogDeviceObjectProperty
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLogDeviceObjectProperty factory function for BACnetConstructedDataLogDeviceObjectProperty
func NewBACnetConstructedDataLogDeviceObjectProperty(logDeviceObjectProperty *BACnetDeviceObjectPropertyReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLogDeviceObjectProperty {
	_result := &BACnetConstructedDataLogDeviceObjectProperty{
		LogDeviceObjectProperty: logDeviceObjectProperty,
		BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLogDeviceObjectProperty(structType interface{}) *BACnetConstructedDataLogDeviceObjectProperty {
	if casted, ok := structType.(BACnetConstructedDataLogDeviceObjectProperty); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLogDeviceObjectProperty); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLogDeviceObjectProperty(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLogDeviceObjectProperty(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLogDeviceObjectProperty) GetTypeName() string {
	return "BACnetConstructedDataLogDeviceObjectProperty"
}

func (m *BACnetConstructedDataLogDeviceObjectProperty) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLogDeviceObjectProperty) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (logDeviceObjectProperty)
	lengthInBits += m.LogDeviceObjectProperty.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLogDeviceObjectProperty) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLogDeviceObjectPropertyParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLogDeviceObjectProperty, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLogDeviceObjectProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLogDeviceObjectProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (logDeviceObjectProperty)
	if pullErr := readBuffer.PullContext("logDeviceObjectProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logDeviceObjectProperty")
	}
	_logDeviceObjectProperty, _logDeviceObjectPropertyErr := BACnetDeviceObjectPropertyReferenceParse(readBuffer)
	if _logDeviceObjectPropertyErr != nil {
		return nil, errors.Wrap(_logDeviceObjectPropertyErr, "Error parsing 'logDeviceObjectProperty' field")
	}
	logDeviceObjectProperty := CastBACnetDeviceObjectPropertyReference(_logDeviceObjectProperty)
	if closeErr := readBuffer.CloseContext("logDeviceObjectProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logDeviceObjectProperty")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLogDeviceObjectProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLogDeviceObjectProperty")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLogDeviceObjectProperty{
		LogDeviceObjectProperty: CastBACnetDeviceObjectPropertyReference(logDeviceObjectProperty),
		BACnetConstructedData:   &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLogDeviceObjectProperty) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLogDeviceObjectProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLogDeviceObjectProperty")
		}

		// Simple Field (logDeviceObjectProperty)
		if pushErr := writeBuffer.PushContext("logDeviceObjectProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for logDeviceObjectProperty")
		}
		_logDeviceObjectPropertyErr := writeBuffer.WriteSerializable(m.LogDeviceObjectProperty)
		if popErr := writeBuffer.PopContext("logDeviceObjectProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for logDeviceObjectProperty")
		}
		if _logDeviceObjectPropertyErr != nil {
			return errors.Wrap(_logDeviceObjectPropertyErr, "Error serializing 'logDeviceObjectProperty' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLogDeviceObjectProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLogDeviceObjectProperty")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLogDeviceObjectProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
