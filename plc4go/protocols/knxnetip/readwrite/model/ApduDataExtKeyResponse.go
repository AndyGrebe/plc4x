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

// ApduDataExtKeyResponse is the data-structure of this message
type ApduDataExtKeyResponse struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtKeyResponse is the corresponding interface of ApduDataExtKeyResponse
type IApduDataExtKeyResponse interface {
	IApduDataExt
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

func (m *ApduDataExtKeyResponse) GetExtApciType() uint8 {
	return 0x14
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtKeyResponse) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtKeyResponse) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtKeyResponse factory function for ApduDataExtKeyResponse
func NewApduDataExtKeyResponse(length uint8) *ApduDataExtKeyResponse {
	_result := &ApduDataExtKeyResponse{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtKeyResponse(structType interface{}) *ApduDataExtKeyResponse {
	if casted, ok := structType.(ApduDataExtKeyResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtKeyResponse); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtKeyResponse(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtKeyResponse(casted.Child)
	}
	return nil
}

func (m *ApduDataExtKeyResponse) GetTypeName() string {
	return "ApduDataExtKeyResponse"
}

func (m *ApduDataExtKeyResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtKeyResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtKeyResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtKeyResponseParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtKeyResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtKeyResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtKeyResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtKeyResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtKeyResponse")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtKeyResponse{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtKeyResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtKeyResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtKeyResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtKeyResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtKeyResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtKeyResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
