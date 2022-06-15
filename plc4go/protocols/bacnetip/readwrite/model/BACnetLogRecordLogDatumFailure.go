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

// BACnetLogRecordLogDatumFailure is the corresponding interface of BACnetLogRecordLogDatumFailure
type BACnetLogRecordLogDatumFailure interface {
	BACnetLogRecordLogDatum
	// GetFailure returns Failure (property field)
	GetFailure() ErrorEnclosed
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetLogRecordLogDatumFailure is the data-structure of this message
type _BACnetLogRecordLogDatumFailure struct {
	*_BACnetLogRecordLogDatum
	Failure ErrorEnclosed

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumFailure) InitializeParent(parent BACnetLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogRecordLogDatumFailure) GetParent() BACnetLogRecordLogDatum {
	return m._BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumFailure) GetFailure() ErrorEnclosed {
	return m.Failure
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumFailure factory function for _BACnetLogRecordLogDatumFailure
func NewBACnetLogRecordLogDatumFailure(failure ErrorEnclosed, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatumFailure {
	_result := &_BACnetLogRecordLogDatumFailure{
		Failure:                  failure,
		_BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumFailure(structType interface{}) BACnetLogRecordLogDatumFailure {
	if casted, ok := structType.(BACnetLogRecordLogDatumFailure); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumFailure); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumFailure) GetTypeName() string {
	return "BACnetLogRecordLogDatumFailure"
}

func (m *_BACnetLogRecordLogDatumFailure) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogRecordLogDatumFailure) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (failure)
	lengthInBits += m.Failure.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumFailure) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogRecordLogDatumFailureParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogRecordLogDatumFailure, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (failure)
	if pullErr := readBuffer.PullContext("failure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for failure")
	}
	_failure, _failureErr := ErrorEnclosedParse(readBuffer, uint8(uint8(8)))
	if _failureErr != nil {
		return nil, errors.Wrap(_failureErr, "Error parsing 'failure' field")
	}
	failure := _failure.(ErrorEnclosed)
	if closeErr := readBuffer.CloseContext("failure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for failure")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumFailure")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogRecordLogDatumFailure{
		Failure:                  failure,
		_BACnetLogRecordLogDatum: &_BACnetLogRecordLogDatum{},
	}
	_child._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogRecordLogDatumFailure) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumFailure")
		}

		// Simple Field (failure)
		if pushErr := writeBuffer.PushContext("failure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for failure")
		}
		_failureErr := writeBuffer.WriteSerializable(m.GetFailure())
		if popErr := writeBuffer.PopContext("failure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for failure")
		}
		if _failureErr != nil {
			return errors.Wrap(_failureErr, "Error serializing 'failure' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumFailure")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}