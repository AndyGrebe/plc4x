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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetApplicationTagBoolean struct {
	*BACnetApplicationTag
	Value   bool
	IsTrue  bool
	IsFalse bool
}

// The corresponding interface
type IBACnetApplicationTagBoolean interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetApplicationTagBoolean) ActualTagNumber() uint8 {
	return 0x1
}

func (m *BACnetApplicationTagBoolean) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader, actualTagNumber uint8, actualLength uint32) {
	m.BACnetApplicationTag.Header = header
	m.BACnetApplicationTag.ActualTagNumber = actualTagNumber
	m.BACnetApplicationTag.ActualLength = actualLength
}

func NewBACnetApplicationTagBoolean(value bool, isTrue bool, isFalse bool, header *BACnetTagHeader, actualTagNumber uint8, actualLength uint32) *BACnetApplicationTag {
	child := &BACnetApplicationTagBoolean{
		Value:                value,
		IsTrue:               isTrue,
		IsFalse:              isFalse,
		BACnetApplicationTag: NewBACnetApplicationTag(header, actualTagNumber, actualLength),
	}
	child.Child = child
	return child.BACnetApplicationTag
}

func CastBACnetApplicationTagBoolean(structType interface{}) *BACnetApplicationTagBoolean {
	castFunc := func(typ interface{}) *BACnetApplicationTagBoolean {
		if casted, ok := typ.(BACnetApplicationTagBoolean); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetApplicationTagBoolean); ok {
			return casted
		}
		if casted, ok := typ.(BACnetApplicationTag); ok {
			return CastBACnetApplicationTagBoolean(casted.Child)
		}
		if casted, ok := typ.(*BACnetApplicationTag); ok {
			return CastBACnetApplicationTagBoolean(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetApplicationTagBoolean) GetTypeName() string {
	return "BACnetApplicationTagBoolean"
}

func (m *BACnetApplicationTagBoolean) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetApplicationTagBoolean) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetApplicationTagBoolean) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetApplicationTagBooleanParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetApplicationTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTagBoolean"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_value := bool((actualLength) == (1))
	value := bool(_value)

	// Virtual field
	_isTrue := value
	isTrue := bool(_isTrue)

	// Virtual field
	_isFalse := !(value)
	isFalse := bool(_isFalse)

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagBoolean"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagBoolean{
		Value:                value,
		IsTrue:               isTrue,
		IsFalse:              isFalse,
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child.BACnetApplicationTag, nil
}

func (m *BACnetApplicationTagBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagBoolean"); pushErr != nil {
			return pushErr
		}
		// Virtual field
		if _valueErr := writeBuffer.WriteVirtual("value", m.Value); _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
		// Virtual field
		if _isTrueErr := writeBuffer.WriteVirtual("isTrue", m.IsTrue); _isTrueErr != nil {
			return errors.Wrap(_isTrueErr, "Error serializing 'isTrue' field")
		}
		// Virtual field
		if _isFalseErr := writeBuffer.WriteVirtual("isFalse", m.IsFalse); _isFalseErr != nil {
			return errors.Wrap(_isFalseErr, "Error serializing 'isFalse' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagBoolean"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}