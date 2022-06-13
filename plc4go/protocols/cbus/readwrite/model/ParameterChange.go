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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const ParameterChange_SPECIALCHAR1 byte = 0x3D
const ParameterChange_SPECIALCHAR2 byte = 0x3D
const ParameterChange_CR byte = 0x0D
const ParameterChange_LF byte = 0x0A

// ParameterChange is the data-structure of this message
type ParameterChange struct {
}

// IParameterChange is the corresponding interface of ParameterChange
type IParameterChange interface {
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *ParameterChange) GetSpecialChar1() byte {
	return ParameterChange_SPECIALCHAR1
}

func (m *ParameterChange) GetSpecialChar2() byte {
	return ParameterChange_SPECIALCHAR2
}

func (m *ParameterChange) GetCr() byte {
	return ParameterChange_CR
}

func (m *ParameterChange) GetLf() byte {
	return ParameterChange_LF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterChange factory function for ParameterChange
func NewParameterChange() *ParameterChange {
	return &ParameterChange{}
}

func CastParameterChange(structType interface{}) *ParameterChange {
	if casted, ok := structType.(ParameterChange); ok {
		return &casted
	}
	if casted, ok := structType.(*ParameterChange); ok {
		return casted
	}
	return nil
}

func (m *ParameterChange) GetTypeName() string {
	return "ParameterChange"
}

func (m *ParameterChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ParameterChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (specialChar1)
	lengthInBits += 8

	// Const Field (specialChar2)
	lengthInBits += 8

	// Const Field (cr)
	lengthInBits += 8

	// Const Field (lf)
	lengthInBits += 8

	return lengthInBits
}

func (m *ParameterChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ParameterChangeParse(readBuffer utils.ReadBuffer) (*ParameterChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (specialChar1)
	specialChar1, _specialChar1Err := readBuffer.ReadByte("specialChar1")
	if _specialChar1Err != nil {
		return nil, errors.Wrap(_specialChar1Err, "Error parsing 'specialChar1' field")
	}
	if specialChar1 != ParameterChange_SPECIALCHAR1 {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ParameterChange_SPECIALCHAR1) + " but got " + fmt.Sprintf("%d", specialChar1))
	}

	// Const Field (specialChar2)
	specialChar2, _specialChar2Err := readBuffer.ReadByte("specialChar2")
	if _specialChar2Err != nil {
		return nil, errors.Wrap(_specialChar2Err, "Error parsing 'specialChar2' field")
	}
	if specialChar2 != ParameterChange_SPECIALCHAR2 {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ParameterChange_SPECIALCHAR2) + " but got " + fmt.Sprintf("%d", specialChar2))
	}

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field")
	}
	if cr != ParameterChange_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ParameterChange_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	// Const Field (lf)
	lf, _lfErr := readBuffer.ReadByte("lf")
	if _lfErr != nil {
		return nil, errors.Wrap(_lfErr, "Error parsing 'lf' field")
	}
	if lf != ParameterChange_LF {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ParameterChange_LF) + " but got " + fmt.Sprintf("%d", lf))
	}

	if closeErr := readBuffer.CloseContext("ParameterChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterChange")
	}

	// Create the instance
	return NewParameterChange(), nil
}

func (m *ParameterChange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ParameterChange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ParameterChange")
	}

	// Const Field (specialChar1)
	_specialChar1Err := writeBuffer.WriteByte("specialChar1", 0x3D)
	if _specialChar1Err != nil {
		return errors.Wrap(_specialChar1Err, "Error serializing 'specialChar1' field")
	}

	// Const Field (specialChar2)
	_specialChar2Err := writeBuffer.WriteByte("specialChar2", 0x3D)
	if _specialChar2Err != nil {
		return errors.Wrap(_specialChar2Err, "Error serializing 'specialChar2' field")
	}

	// Const Field (cr)
	_crErr := writeBuffer.WriteByte("cr", 0x0D)
	if _crErr != nil {
		return errors.Wrap(_crErr, "Error serializing 'cr' field")
	}

	// Const Field (lf)
	_lfErr := writeBuffer.WriteByte("lf", 0x0A)
	if _lfErr != nil {
		return errors.Wrap(_lfErr, "Error serializing 'lf' field")
	}

	if popErr := writeBuffer.PopContext("ParameterChange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ParameterChange")
	}
	return nil
}

func (m *ParameterChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
