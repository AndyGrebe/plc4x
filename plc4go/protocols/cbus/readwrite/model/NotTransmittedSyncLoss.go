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

// NotTransmittedSyncLoss is the data-structure of this message
type NotTransmittedSyncLoss struct {
	*Confirmation
}

// INotTransmittedSyncLoss is the corresponding interface of NotTransmittedSyncLoss
type INotTransmittedSyncLoss interface {
	IConfirmation
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

func (m *NotTransmittedSyncLoss) GetConfirmationType() byte {
	return 0x25
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *NotTransmittedSyncLoss) InitializeParent(parent *Confirmation, alpha *Alpha) {
	m.Confirmation.Alpha = alpha
}

func (m *NotTransmittedSyncLoss) GetParent() *Confirmation {
	return m.Confirmation
}

// NewNotTransmittedSyncLoss factory function for NotTransmittedSyncLoss
func NewNotTransmittedSyncLoss(alpha *Alpha) *NotTransmittedSyncLoss {
	_result := &NotTransmittedSyncLoss{
		Confirmation: NewConfirmation(alpha),
	}
	_result.Child = _result
	return _result
}

func CastNotTransmittedSyncLoss(structType interface{}) *NotTransmittedSyncLoss {
	if casted, ok := structType.(NotTransmittedSyncLoss); ok {
		return &casted
	}
	if casted, ok := structType.(*NotTransmittedSyncLoss); ok {
		return casted
	}
	if casted, ok := structType.(Confirmation); ok {
		return CastNotTransmittedSyncLoss(casted.Child)
	}
	if casted, ok := structType.(*Confirmation); ok {
		return CastNotTransmittedSyncLoss(casted.Child)
	}
	return nil
}

func (m *NotTransmittedSyncLoss) GetTypeName() string {
	return "NotTransmittedSyncLoss"
}

func (m *NotTransmittedSyncLoss) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NotTransmittedSyncLoss) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *NotTransmittedSyncLoss) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NotTransmittedSyncLossParse(readBuffer utils.ReadBuffer) (*NotTransmittedSyncLoss, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NotTransmittedSyncLoss"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NotTransmittedSyncLoss")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NotTransmittedSyncLoss"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NotTransmittedSyncLoss")
	}

	// Create a partially initialized instance
	_child := &NotTransmittedSyncLoss{
		Confirmation: &Confirmation{},
	}
	_child.Confirmation.Child = _child
	return _child, nil
}

func (m *NotTransmittedSyncLoss) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NotTransmittedSyncLoss"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NotTransmittedSyncLoss")
		}

		if popErr := writeBuffer.PopContext("NotTransmittedSyncLoss"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NotTransmittedSyncLoss")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NotTransmittedSyncLoss) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
