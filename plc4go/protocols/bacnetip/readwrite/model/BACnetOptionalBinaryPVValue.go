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

// BACnetOptionalBinaryPVValue is the data-structure of this message
type BACnetOptionalBinaryPVValue struct {
	*BACnetOptionalBinaryPV
	BinaryPv *BACnetBinaryPVTagged
}

// IBACnetOptionalBinaryPVValue is the corresponding interface of BACnetOptionalBinaryPVValue
type IBACnetOptionalBinaryPVValue interface {
	IBACnetOptionalBinaryPV
	// GetBinaryPv returns BinaryPv (property field)
	GetBinaryPv() *BACnetBinaryPVTagged
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

func (m *BACnetOptionalBinaryPVValue) InitializeParent(parent *BACnetOptionalBinaryPV, peekedTagHeader *BACnetTagHeader) {
	m.BACnetOptionalBinaryPV.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetOptionalBinaryPVValue) GetParent() *BACnetOptionalBinaryPV {
	return m.BACnetOptionalBinaryPV
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetOptionalBinaryPVValue) GetBinaryPv() *BACnetBinaryPVTagged {
	return m.BinaryPv
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalBinaryPVValue factory function for BACnetOptionalBinaryPVValue
func NewBACnetOptionalBinaryPVValue(binaryPv *BACnetBinaryPVTagged, peekedTagHeader *BACnetTagHeader) *BACnetOptionalBinaryPVValue {
	_result := &BACnetOptionalBinaryPVValue{
		BinaryPv:               binaryPv,
		BACnetOptionalBinaryPV: NewBACnetOptionalBinaryPV(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetOptionalBinaryPVValue(structType interface{}) *BACnetOptionalBinaryPVValue {
	if casted, ok := structType.(BACnetOptionalBinaryPVValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetOptionalBinaryPVValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetOptionalBinaryPV); ok {
		return CastBACnetOptionalBinaryPVValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetOptionalBinaryPV); ok {
		return CastBACnetOptionalBinaryPVValue(casted.Child)
	}
	return nil
}

func (m *BACnetOptionalBinaryPVValue) GetTypeName() string {
	return "BACnetOptionalBinaryPVValue"
}

func (m *BACnetOptionalBinaryPVValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetOptionalBinaryPVValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (binaryPv)
	lengthInBits += m.BinaryPv.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetOptionalBinaryPVValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetOptionalBinaryPVValueParse(readBuffer utils.ReadBuffer) (*BACnetOptionalBinaryPVValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalBinaryPVValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalBinaryPVValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (binaryPv)
	if pullErr := readBuffer.PullContext("binaryPv"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for binaryPv")
	}
	_binaryPv, _binaryPvErr := BACnetBinaryPVTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _binaryPvErr != nil {
		return nil, errors.Wrap(_binaryPvErr, "Error parsing 'binaryPv' field")
	}
	binaryPv := CastBACnetBinaryPVTagged(_binaryPv)
	if closeErr := readBuffer.CloseContext("binaryPv"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for binaryPv")
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalBinaryPVValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalBinaryPVValue")
	}

	// Create a partially initialized instance
	_child := &BACnetOptionalBinaryPVValue{
		BinaryPv:               CastBACnetBinaryPVTagged(binaryPv),
		BACnetOptionalBinaryPV: &BACnetOptionalBinaryPV{},
	}
	_child.BACnetOptionalBinaryPV.Child = _child
	return _child, nil
}

func (m *BACnetOptionalBinaryPVValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalBinaryPVValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalBinaryPVValue")
		}

		// Simple Field (binaryPv)
		if pushErr := writeBuffer.PushContext("binaryPv"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for binaryPv")
		}
		_binaryPvErr := writeBuffer.WriteSerializable(m.BinaryPv)
		if popErr := writeBuffer.PopContext("binaryPv"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for binaryPv")
		}
		if _binaryPvErr != nil {
			return errors.Wrap(_binaryPvErr, "Error serializing 'binaryPv' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalBinaryPVValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalBinaryPVValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetOptionalBinaryPVValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
