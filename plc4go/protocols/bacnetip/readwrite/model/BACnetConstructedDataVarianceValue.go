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

// BACnetConstructedDataVarianceValue is the corresponding interface of BACnetConstructedDataVarianceValue
type BACnetConstructedDataVarianceValue interface {
	BACnetConstructedData
	// GetVarianceValue returns VarianceValue (property field)
	GetVarianceValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConstructedDataVarianceValue is the data-structure of this message
type _BACnetConstructedDataVarianceValue struct {
	*_BACnetConstructedData
	VarianceValue BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVarianceValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVarianceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VARIANCE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVarianceValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataVarianceValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVarianceValue) GetVarianceValue() BACnetApplicationTagReal {
	return m.VarianceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataVarianceValue) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetVarianceValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVarianceValue factory function for _BACnetConstructedDataVarianceValue
func NewBACnetConstructedDataVarianceValue(varianceValue BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVarianceValue {
	_result := &_BACnetConstructedDataVarianceValue{
		VarianceValue:          varianceValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVarianceValue(structType interface{}) BACnetConstructedDataVarianceValue {
	if casted, ok := structType.(BACnetConstructedDataVarianceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVarianceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVarianceValue) GetTypeName() string {
	return "BACnetConstructedDataVarianceValue"
}

func (m *_BACnetConstructedDataVarianceValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataVarianceValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (varianceValue)
	lengthInBits += m.VarianceValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataVarianceValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataVarianceValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVarianceValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVarianceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVarianceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (varianceValue)
	if pullErr := readBuffer.PullContext("varianceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for varianceValue")
	}
	_varianceValue, _varianceValueErr := BACnetApplicationTagParse(readBuffer)
	if _varianceValueErr != nil {
		return nil, errors.Wrap(_varianceValueErr, "Error parsing 'varianceValue' field")
	}
	varianceValue := _varianceValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("varianceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for varianceValue")
	}

	// Virtual field
	_actualValue := varianceValue
	actualValue := _actualValue.(BACnetApplicationTagReal)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVarianceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVarianceValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataVarianceValue{
		VarianceValue:          varianceValue,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataVarianceValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVarianceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVarianceValue")
		}

		// Simple Field (varianceValue)
		if pushErr := writeBuffer.PushContext("varianceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for varianceValue")
		}
		_varianceValueErr := writeBuffer.WriteSerializable(m.GetVarianceValue())
		if popErr := writeBuffer.PopContext("varianceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for varianceValue")
		}
		if _varianceValueErr != nil {
			return errors.Wrap(_varianceValueErr, "Error serializing 'varianceValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVarianceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVarianceValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVarianceValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}