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

// BACnetPropertyStatesTimerState is the data-structure of this message
type BACnetPropertyStatesTimerState struct {
	*BACnetPropertyStates
	TimerState *BACnetTimerStateTagged
}

// IBACnetPropertyStatesTimerState is the corresponding interface of BACnetPropertyStatesTimerState
type IBACnetPropertyStatesTimerState interface {
	IBACnetPropertyStates
	// GetTimerState returns TimerState (property field)
	GetTimerState() *BACnetTimerStateTagged
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

func (m *BACnetPropertyStatesTimerState) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesTimerState) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesTimerState) GetTimerState() *BACnetTimerStateTagged {
	return m.TimerState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesTimerState factory function for BACnetPropertyStatesTimerState
func NewBACnetPropertyStatesTimerState(timerState *BACnetTimerStateTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesTimerState {
	_result := &BACnetPropertyStatesTimerState{
		TimerState:           timerState,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesTimerState(structType interface{}) *BACnetPropertyStatesTimerState {
	if casted, ok := structType.(BACnetPropertyStatesTimerState); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesTimerState); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesTimerState(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesTimerState(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesTimerState) GetTypeName() string {
	return "BACnetPropertyStatesTimerState"
}

func (m *BACnetPropertyStatesTimerState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesTimerState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timerState)
	lengthInBits += m.TimerState.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesTimerState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesTimerStateParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesTimerState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesTimerState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesTimerState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timerState)
	if pullErr := readBuffer.PullContext("timerState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timerState")
	}
	_timerState, _timerStateErr := BACnetTimerStateTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _timerStateErr != nil {
		return nil, errors.Wrap(_timerStateErr, "Error parsing 'timerState' field")
	}
	timerState := CastBACnetTimerStateTagged(_timerState)
	if closeErr := readBuffer.CloseContext("timerState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timerState")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesTimerState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesTimerState")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesTimerState{
		TimerState:           CastBACnetTimerStateTagged(timerState),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesTimerState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesTimerState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesTimerState")
		}

		// Simple Field (timerState)
		if pushErr := writeBuffer.PushContext("timerState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timerState")
		}
		_timerStateErr := writeBuffer.WriteSerializable(m.TimerState)
		if popErr := writeBuffer.PopContext("timerState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timerState")
		}
		if _timerStateErr != nil {
			return errors.Wrap(_timerStateErr, "Error serializing 'timerState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesTimerState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesTimerState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesTimerState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
