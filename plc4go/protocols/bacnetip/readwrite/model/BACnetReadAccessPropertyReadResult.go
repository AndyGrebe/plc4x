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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetReadAccessPropertyReadResult is the data-structure of this message
type BACnetReadAccessPropertyReadResult struct {
	PeekedTagHeader     *BACnetTagHeader
	PropertyValue       *BACnetConstructedData
	PropertyAccessError *ErrorEnclosed

	// Arguments.
	ObjectType                 BACnetObjectType
	PropertyIdentifierArgument BACnetPropertyIdentifier
}

// IBACnetReadAccessPropertyReadResult is the corresponding interface of BACnetReadAccessPropertyReadResult
type IBACnetReadAccessPropertyReadResult interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() *BACnetTagHeader
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() *BACnetConstructedData
	// GetPropertyAccessError returns PropertyAccessError (property field)
	GetPropertyAccessError() *ErrorEnclosed
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
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

func (m *BACnetReadAccessPropertyReadResult) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *BACnetReadAccessPropertyReadResult) GetPropertyValue() *BACnetConstructedData {
	return m.PropertyValue
}

func (m *BACnetReadAccessPropertyReadResult) GetPropertyAccessError() *ErrorEnclosed {
	return m.PropertyAccessError
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetReadAccessPropertyReadResult) GetPeekedTagNumber() uint8 {
	propertyValue := m.PropertyValue
	_ = propertyValue
	propertyAccessError := m.PropertyAccessError
	_ = propertyAccessError
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetReadAccessPropertyReadResult factory function for BACnetReadAccessPropertyReadResult
func NewBACnetReadAccessPropertyReadResult(peekedTagHeader *BACnetTagHeader, propertyValue *BACnetConstructedData, propertyAccessError *ErrorEnclosed, objectType BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) *BACnetReadAccessPropertyReadResult {
	return &BACnetReadAccessPropertyReadResult{PeekedTagHeader: peekedTagHeader, PropertyValue: propertyValue, PropertyAccessError: propertyAccessError, ObjectType: objectType, PropertyIdentifierArgument: propertyIdentifierArgument}
}

func CastBACnetReadAccessPropertyReadResult(structType interface{}) *BACnetReadAccessPropertyReadResult {
	if casted, ok := structType.(BACnetReadAccessPropertyReadResult); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetReadAccessPropertyReadResult); ok {
		return casted
	}
	return nil
}

func (m *BACnetReadAccessPropertyReadResult) GetTypeName() string {
	return "BACnetReadAccessPropertyReadResult"
}

func (m *BACnetReadAccessPropertyReadResult) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetReadAccessPropertyReadResult) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += (*m.PropertyValue).GetLengthInBits()
	}

	// Optional Field (propertyAccessError)
	if m.PropertyAccessError != nil {
		lengthInBits += (*m.PropertyAccessError).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetReadAccessPropertyReadResult) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetReadAccessPropertyReadResultParse(readBuffer utils.ReadBuffer, objectType BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetReadAccessPropertyReadResult, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetReadAccessPropertyReadResult"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Optional Field (propertyValue) (Can be skipped, if a given expression evaluates to false)
	var propertyValue *BACnetConstructedData = nil
	if bool((peekedTagNumber) == (4)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(4), objectType, propertyIdentifierArgument)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyValue' field")
		default:
			propertyValue = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Validation
	if !(bool(bool(bool(bool((peekedTagNumber) == (4))) && bool(bool((propertyValue) != (nil))))) || bool(bool((peekedTagNumber) != (4)))) {
		return nil, utils.ParseValidationError{"failure parsing field 4"}
	}

	// Optional Field (propertyAccessError) (Can be skipped, if a given expression evaluates to false)
	var propertyAccessError *ErrorEnclosed = nil
	if bool((peekedTagNumber) == (5)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyAccessError"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := ErrorEnclosedParse(readBuffer, uint8(5))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyAccessError' field")
		default:
			propertyAccessError = CastErrorEnclosed(_val)
			if closeErr := readBuffer.CloseContext("propertyAccessError"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Validation
	if !(bool(bool(bool(bool((peekedTagNumber) == (5))) && bool(bool((propertyAccessError) != (nil))))) || bool(bool((peekedTagNumber) != (5)))) {
		return nil, utils.ParseValidationError{"failure parsing field 5"}
	}

	// Validation
	if !(bool(bool((peekedTagNumber) == (4))) || bool(bool((peekedTagNumber) == (5)))) {
		return nil, utils.ParseAssertError{"should be either 4 or 5"}
	}

	if closeErr := readBuffer.CloseContext("BACnetReadAccessPropertyReadResult"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetReadAccessPropertyReadResult(peekedTagHeader, propertyValue, propertyAccessError, objectType, propertyIdentifierArgument), nil
}

func (m *BACnetReadAccessPropertyReadResult) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetReadAccessPropertyReadResult"); pushErr != nil {
		return pushErr
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Optional Field (propertyValue) (Can be skipped, if the value is null)
	var propertyValue *BACnetConstructedData = nil
	if m.PropertyValue != nil {
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return pushErr
		}
		propertyValue = m.PropertyValue
		_propertyValueErr := propertyValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return popErr
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}
	}

	// Optional Field (propertyAccessError) (Can be skipped, if the value is null)
	var propertyAccessError *ErrorEnclosed = nil
	if m.PropertyAccessError != nil {
		if pushErr := writeBuffer.PushContext("propertyAccessError"); pushErr != nil {
			return pushErr
		}
		propertyAccessError = m.PropertyAccessError
		_propertyAccessErrorErr := propertyAccessError.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyAccessError"); popErr != nil {
			return popErr
		}
		if _propertyAccessErrorErr != nil {
			return errors.Wrap(_propertyAccessErrorErr, "Error serializing 'propertyAccessError' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetReadAccessPropertyReadResult"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetReadAccessPropertyReadResult) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}