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

// BACnetConfirmedServiceRequestConfirmedUnknown is the data-structure of this message
type BACnetConfirmedServiceRequestConfirmedUnknown struct {
	*BACnetConfirmedServiceRequest
	UnknownBytes []byte

	// Arguments.
	ServiceRequestLength        uint16
	ServiceRequestPayloadLength uint16
}

// IBACnetConfirmedServiceRequestConfirmedUnknown is the corresponding interface of BACnetConfirmedServiceRequestConfirmedUnknown
type IBACnetConfirmedServiceRequestConfirmedUnknown interface {
	IBACnetConfirmedServiceRequest
	// GetUnknownBytes returns UnknownBytes (property field)
	GetUnknownBytes() []byte
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

func (m *BACnetConfirmedServiceRequestConfirmedUnknown) GetServiceChoice() BACnetConfirmedServiceChoice {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestConfirmedUnknown) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestConfirmedUnknown) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestConfirmedUnknown) GetUnknownBytes() []byte {
	return m.UnknownBytes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedUnknown factory function for BACnetConfirmedServiceRequestConfirmedUnknown
func NewBACnetConfirmedServiceRequestConfirmedUnknown(unknownBytes []byte, serviceRequestLength uint16, serviceRequestPayloadLength uint16) *BACnetConfirmedServiceRequestConfirmedUnknown {
	_result := &BACnetConfirmedServiceRequestConfirmedUnknown{
		UnknownBytes:                  unknownBytes,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestConfirmedUnknown(structType interface{}) *BACnetConfirmedServiceRequestConfirmedUnknown {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedUnknown); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedUnknown); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestConfirmedUnknown(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestConfirmedUnknown(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestConfirmedUnknown) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedUnknown"
}

func (m *BACnetConfirmedServiceRequestConfirmedUnknown) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestConfirmedUnknown) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.UnknownBytes) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownBytes))
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestConfirmedUnknown) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedUnknownParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16, serviceRequestPayloadLength uint16) (*BACnetConfirmedServiceRequestConfirmedUnknown, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedUnknown"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedUnknown")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (unknownBytes)
	numberOfBytesunknownBytes := int(serviceRequestPayloadLength)
	unknownBytes, _readArrayErr := readBuffer.ReadByteArray("unknownBytes", numberOfBytesunknownBytes)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownBytes' field")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedUnknown"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedUnknown")
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestConfirmedUnknown{
		UnknownBytes:                  unknownBytes,
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestConfirmedUnknown) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedUnknown"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedUnknown")
		}

		// Array Field (unknownBytes)
		if m.UnknownBytes != nil {
			// Byte Array field (unknownBytes)
			_writeArrayErr := writeBuffer.WriteByteArray("unknownBytes", m.UnknownBytes)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'unknownBytes' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedUnknown"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedUnknown")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestConfirmedUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
