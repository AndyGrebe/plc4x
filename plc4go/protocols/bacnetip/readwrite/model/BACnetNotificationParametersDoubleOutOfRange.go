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

// BACnetNotificationParametersDoubleOutOfRange is the data-structure of this message
type BACnetNotificationParametersDoubleOutOfRange struct {
	*BACnetNotificationParameters
	InnerOpeningTag *BACnetOpeningTag
	ExceedingValue  *BACnetContextTagDouble
	StatusFlags     *BACnetStatusFlagsTagged
	Deadband        *BACnetContextTagDouble
	ExceededLimit   *BACnetContextTagDouble
	InnerClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber          uint8
	ObjectTypeArgument BACnetObjectType
}

// IBACnetNotificationParametersDoubleOutOfRange is the corresponding interface of BACnetNotificationParametersDoubleOutOfRange
type IBACnetNotificationParametersDoubleOutOfRange interface {
	IBACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() *BACnetOpeningTag
	// GetExceedingValue returns ExceedingValue (property field)
	GetExceedingValue() *BACnetContextTagDouble
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() *BACnetStatusFlagsTagged
	// GetDeadband returns Deadband (property field)
	GetDeadband() *BACnetContextTagDouble
	// GetExceededLimit returns ExceededLimit (property field)
	GetExceededLimit() *BACnetContextTagDouble
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() *BACnetClosingTag
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetNotificationParametersDoubleOutOfRange) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParameters.OpeningTag = openingTag
	m.BACnetNotificationParameters.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParameters.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersDoubleOutOfRange) GetParent() *BACnetNotificationParameters {
	return m.BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersDoubleOutOfRange) GetInnerOpeningTag() *BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *BACnetNotificationParametersDoubleOutOfRange) GetExceedingValue() *BACnetContextTagDouble {
	return m.ExceedingValue
}

func (m *BACnetNotificationParametersDoubleOutOfRange) GetStatusFlags() *BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *BACnetNotificationParametersDoubleOutOfRange) GetDeadband() *BACnetContextTagDouble {
	return m.Deadband
}

func (m *BACnetNotificationParametersDoubleOutOfRange) GetExceededLimit() *BACnetContextTagDouble {
	return m.ExceededLimit
}

func (m *BACnetNotificationParametersDoubleOutOfRange) GetInnerClosingTag() *BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersDoubleOutOfRange factory function for BACnetNotificationParametersDoubleOutOfRange
func NewBACnetNotificationParametersDoubleOutOfRange(innerOpeningTag *BACnetOpeningTag, exceedingValue *BACnetContextTagDouble, statusFlags *BACnetStatusFlagsTagged, deadband *BACnetContextTagDouble, exceededLimit *BACnetContextTagDouble, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *BACnetNotificationParametersDoubleOutOfRange {
	_result := &BACnetNotificationParametersDoubleOutOfRange{
		InnerOpeningTag:              innerOpeningTag,
		ExceedingValue:               exceedingValue,
		StatusFlags:                  statusFlags,
		Deadband:                     deadband,
		ExceededLimit:                exceededLimit,
		InnerClosingTag:              innerClosingTag,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersDoubleOutOfRange(structType interface{}) *BACnetNotificationParametersDoubleOutOfRange {
	if casted, ok := structType.(BACnetNotificationParametersDoubleOutOfRange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersDoubleOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersDoubleOutOfRange(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersDoubleOutOfRange(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersDoubleOutOfRange) GetTypeName() string {
	return "BACnetNotificationParametersDoubleOutOfRange"
}

func (m *BACnetNotificationParametersDoubleOutOfRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersDoubleOutOfRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (exceedingValue)
	lengthInBits += m.ExceedingValue.GetLengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits()

	// Simple field (exceededLimit)
	lengthInBits += m.ExceededLimit.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersDoubleOutOfRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersDoubleOutOfRangeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParametersDoubleOutOfRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersDoubleOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersDoubleOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (exceedingValue)
	if pullErr := readBuffer.PullContext("exceedingValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for exceedingValue")
	}
	_exceedingValue, _exceedingValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_DOUBLE))
	if _exceedingValueErr != nil {
		return nil, errors.Wrap(_exceedingValueErr, "Error parsing 'exceedingValue' field")
	}
	exceedingValue := CastBACnetContextTagDouble(_exceedingValue)
	if closeErr := readBuffer.CloseContext("exceedingValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for exceedingValue")
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := CastBACnetStatusFlagsTagged(_statusFlags)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Simple Field (deadband)
	if pullErr := readBuffer.PullContext("deadband"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deadband")
	}
	_deadband, _deadbandErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_DOUBLE))
	if _deadbandErr != nil {
		return nil, errors.Wrap(_deadbandErr, "Error parsing 'deadband' field")
	}
	deadband := CastBACnetContextTagDouble(_deadband)
	if closeErr := readBuffer.CloseContext("deadband"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deadband")
	}

	// Simple Field (exceededLimit)
	if pullErr := readBuffer.PullContext("exceededLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for exceededLimit")
	}
	_exceededLimit, _exceededLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_DOUBLE))
	if _exceededLimitErr != nil {
		return nil, errors.Wrap(_exceededLimitErr, "Error parsing 'exceededLimit' field")
	}
	exceededLimit := CastBACnetContextTagDouble(_exceededLimit)
	if closeErr := readBuffer.CloseContext("exceededLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for exceededLimit")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersDoubleOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersDoubleOutOfRange")
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersDoubleOutOfRange{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		ExceedingValue:               CastBACnetContextTagDouble(exceedingValue),
		StatusFlags:                  CastBACnetStatusFlagsTagged(statusFlags),
		Deadband:                     CastBACnetContextTagDouble(deadband),
		ExceededLimit:                CastBACnetContextTagDouble(exceededLimit),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersDoubleOutOfRange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersDoubleOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersDoubleOutOfRange")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(m.InnerOpeningTag)
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (exceedingValue)
		if pushErr := writeBuffer.PushContext("exceedingValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for exceedingValue")
		}
		_exceedingValueErr := writeBuffer.WriteSerializable(m.ExceedingValue)
		if popErr := writeBuffer.PopContext("exceedingValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for exceedingValue")
		}
		if _exceedingValueErr != nil {
			return errors.Wrap(_exceedingValueErr, "Error serializing 'exceedingValue' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusFlags")
		}
		_statusFlagsErr := writeBuffer.WriteSerializable(m.StatusFlags)
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusFlags")
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (deadband)
		if pushErr := writeBuffer.PushContext("deadband"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deadband")
		}
		_deadbandErr := writeBuffer.WriteSerializable(m.Deadband)
		if popErr := writeBuffer.PopContext("deadband"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deadband")
		}
		if _deadbandErr != nil {
			return errors.Wrap(_deadbandErr, "Error serializing 'deadband' field")
		}

		// Simple Field (exceededLimit)
		if pushErr := writeBuffer.PushContext("exceededLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for exceededLimit")
		}
		_exceededLimitErr := writeBuffer.WriteSerializable(m.ExceededLimit)
		if popErr := writeBuffer.PopContext("exceededLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for exceededLimit")
		}
		if _exceededLimitErr != nil {
			return errors.Wrap(_exceededLimitErr, "Error serializing 'exceededLimit' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(m.InnerClosingTag)
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersDoubleOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersDoubleOutOfRange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersDoubleOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
