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

// BACnetPropertyStatesSystemStatus is the data-structure of this message
type BACnetPropertyStatesSystemStatus struct {
	*BACnetPropertyStates
	SystemStatus *BACnetDeviceStatusTagged
}

// IBACnetPropertyStatesSystemStatus is the corresponding interface of BACnetPropertyStatesSystemStatus
type IBACnetPropertyStatesSystemStatus interface {
	IBACnetPropertyStates
	// GetSystemStatus returns SystemStatus (property field)
	GetSystemStatus() *BACnetDeviceStatusTagged
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

func (m *BACnetPropertyStatesSystemStatus) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesSystemStatus) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesSystemStatus) GetSystemStatus() *BACnetDeviceStatusTagged {
	return m.SystemStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesSystemStatus factory function for BACnetPropertyStatesSystemStatus
func NewBACnetPropertyStatesSystemStatus(systemStatus *BACnetDeviceStatusTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesSystemStatus {
	_result := &BACnetPropertyStatesSystemStatus{
		SystemStatus:         systemStatus,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesSystemStatus(structType interface{}) *BACnetPropertyStatesSystemStatus {
	if casted, ok := structType.(BACnetPropertyStatesSystemStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesSystemStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesSystemStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesSystemStatus(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesSystemStatus) GetTypeName() string {
	return "BACnetPropertyStatesSystemStatus"
}

func (m *BACnetPropertyStatesSystemStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesSystemStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (systemStatus)
	lengthInBits += m.SystemStatus.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesSystemStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesSystemStatusParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesSystemStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesSystemStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesSystemStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (systemStatus)
	if pullErr := readBuffer.PullContext("systemStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for systemStatus")
	}
	_systemStatus, _systemStatusErr := BACnetDeviceStatusTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _systemStatusErr != nil {
		return nil, errors.Wrap(_systemStatusErr, "Error parsing 'systemStatus' field")
	}
	systemStatus := CastBACnetDeviceStatusTagged(_systemStatus)
	if closeErr := readBuffer.CloseContext("systemStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for systemStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesSystemStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesSystemStatus")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesSystemStatus{
		SystemStatus:         CastBACnetDeviceStatusTagged(systemStatus),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesSystemStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesSystemStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesSystemStatus")
		}

		// Simple Field (systemStatus)
		if pushErr := writeBuffer.PushContext("systemStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for systemStatus")
		}
		_systemStatusErr := writeBuffer.WriteSerializable(m.SystemStatus)
		if popErr := writeBuffer.PopContext("systemStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for systemStatus")
		}
		if _systemStatusErr != nil {
			return errors.Wrap(_systemStatusErr, "Error serializing 'systemStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesSystemStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesSystemStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesSystemStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
