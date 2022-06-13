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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyWriteDefinition is the data-structure of this message
type BACnetPropertyWriteDefinition struct {
	PropertyIdentifier *BACnetPropertyIdentifierTagged
	ArrayIndex         *BACnetContextTagUnsignedInteger
	PropertyValue      *BACnetConstructedData
	Priority           *BACnetContextTagUnsignedInteger

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

// IBACnetPropertyWriteDefinition is the corresponding interface of BACnetPropertyWriteDefinition
type IBACnetPropertyWriteDefinition interface {
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() *BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() *BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() *BACnetConstructedData
	// GetPriority returns Priority (property field)
	GetPriority() *BACnetContextTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyWriteDefinition) GetPropertyIdentifier() *BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *BACnetPropertyWriteDefinition) GetArrayIndex() *BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *BACnetPropertyWriteDefinition) GetPropertyValue() *BACnetConstructedData {
	return m.PropertyValue
}

func (m *BACnetPropertyWriteDefinition) GetPriority() *BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyWriteDefinition factory function for BACnetPropertyWriteDefinition
func NewBACnetPropertyWriteDefinition(propertyIdentifier *BACnetPropertyIdentifierTagged, arrayIndex *BACnetContextTagUnsignedInteger, propertyValue *BACnetConstructedData, priority *BACnetContextTagUnsignedInteger, objectTypeArgument BACnetObjectType) *BACnetPropertyWriteDefinition {
	return &BACnetPropertyWriteDefinition{PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, PropertyValue: propertyValue, Priority: priority, ObjectTypeArgument: objectTypeArgument}
}

func CastBACnetPropertyWriteDefinition(structType interface{}) *BACnetPropertyWriteDefinition {
	if casted, ok := structType.(BACnetPropertyWriteDefinition); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyWriteDefinition); ok {
		return casted
	}
	return nil
}

func (m *BACnetPropertyWriteDefinition) GetTypeName() string {
	return "BACnetPropertyWriteDefinition"
}

func (m *BACnetPropertyWriteDefinition) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyWriteDefinition) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += (*m.ArrayIndex).GetLengthInBits()
	}

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += (*m.PropertyValue).GetLengthInBits()
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += (*m.Priority).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetPropertyWriteDefinition) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyWriteDefinitionParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (*BACnetPropertyWriteDefinition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyWriteDefinition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyWriteDefinition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := CastBACnetPropertyIdentifierTagged(_propertyIdentifier)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for arrayIndex")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field")
		default:
			arrayIndex = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if a given expression evaluates to false)
	var propertyValue *BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for propertyValue")
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(2), objectTypeArgument, propertyIdentifier.GetValue(), CastBACnetTagPayloadUnsignedInteger(CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() interface{} { return CastBACnetTagPayloadUnsignedInteger((*arrayIndex).GetPayload()) }, func() interface{} { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyValue' field")
		default:
			propertyValue = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for propertyValue")
			}
		}
	}

	// Optional Field (priority) (Can be skipped, if a given expression evaluates to false)
	var priority *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for priority")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priority' field")
		default:
			priority = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for priority")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyWriteDefinition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyWriteDefinition")
	}

	// Create the instance
	return NewBACnetPropertyWriteDefinition(propertyIdentifier, arrayIndex, propertyValue, priority, objectTypeArgument), nil
}

func (m *BACnetPropertyWriteDefinition) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPropertyWriteDefinition"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyWriteDefinition")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
	}
	_propertyIdentifierErr := writeBuffer.WriteSerializable(m.PropertyIdentifier)
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	if m.ArrayIndex != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for arrayIndex")
		}
		arrayIndex = m.ArrayIndex
		_arrayIndexErr := writeBuffer.WriteSerializable(arrayIndex)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for arrayIndex")
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if the value is null)
	var propertyValue *BACnetConstructedData = nil
	if m.PropertyValue != nil {
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyValue")
		}
		propertyValue = m.PropertyValue
		_propertyValueErr := writeBuffer.WriteSerializable(propertyValue)
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyValue")
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}
	}

	// Optional Field (priority) (Can be skipped, if the value is null)
	var priority *BACnetContextTagUnsignedInteger = nil
	if m.Priority != nil {
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priority")
		}
		priority = m.Priority
		_priorityErr := writeBuffer.WriteSerializable(priority)
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priority")
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyWriteDefinition"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyWriteDefinition")
	}
	return nil
}

func (m *BACnetPropertyWriteDefinition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
