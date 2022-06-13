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

// FirmataMessageAnalogIO is the data-structure of this message
type FirmataMessageAnalogIO struct {
	*FirmataMessage
	Pin  uint8
	Data []int8

	// Arguments.
	Response bool
}

// IFirmataMessageAnalogIO is the corresponding interface of FirmataMessageAnalogIO
type IFirmataMessageAnalogIO interface {
	IFirmataMessage
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetData returns Data (property field)
	GetData() []int8
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

func (m *FirmataMessageAnalogIO) GetMessageType() uint8 {
	return 0xE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *FirmataMessageAnalogIO) InitializeParent(parent *FirmataMessage) {}

func (m *FirmataMessageAnalogIO) GetParent() *FirmataMessage {
	return m.FirmataMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *FirmataMessageAnalogIO) GetPin() uint8 {
	return m.Pin
}

func (m *FirmataMessageAnalogIO) GetData() []int8 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFirmataMessageAnalogIO factory function for FirmataMessageAnalogIO
func NewFirmataMessageAnalogIO(pin uint8, data []int8, response bool) *FirmataMessageAnalogIO {
	_result := &FirmataMessageAnalogIO{
		Pin:            pin,
		Data:           data,
		FirmataMessage: NewFirmataMessage(response),
	}
	_result.Child = _result
	return _result
}

func CastFirmataMessageAnalogIO(structType interface{}) *FirmataMessageAnalogIO {
	if casted, ok := structType.(FirmataMessageAnalogIO); ok {
		return &casted
	}
	if casted, ok := structType.(*FirmataMessageAnalogIO); ok {
		return casted
	}
	if casted, ok := structType.(FirmataMessage); ok {
		return CastFirmataMessageAnalogIO(casted.Child)
	}
	if casted, ok := structType.(*FirmataMessage); ok {
		return CastFirmataMessageAnalogIO(casted.Child)
	}
	return nil
}

func (m *FirmataMessageAnalogIO) GetTypeName() string {
	return "FirmataMessageAnalogIO"
}

func (m *FirmataMessageAnalogIO) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *FirmataMessageAnalogIO) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 4

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *FirmataMessageAnalogIO) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataMessageAnalogIOParse(readBuffer utils.ReadBuffer, response bool) (*FirmataMessageAnalogIO, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageAnalogIO"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageAnalogIO")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pin)
	_pin, _pinErr := readBuffer.ReadUint8("pin", 4)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field")
	}
	pin := _pin

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	// Count array
	data := make([]int8, uint16(2))
	{
		for curItem := uint16(0); curItem < uint16(uint16(2)); curItem++ {
			_item, _err := readBuffer.ReadInt8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	if closeErr := readBuffer.CloseContext("FirmataMessageAnalogIO"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageAnalogIO")
	}

	// Create a partially initialized instance
	_child := &FirmataMessageAnalogIO{
		Pin:            pin,
		Data:           data,
		FirmataMessage: &FirmataMessage{},
	}
	_child.FirmataMessage.Child = _child
	return _child, nil
}

func (m *FirmataMessageAnalogIO) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageAnalogIO"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageAnalogIO")
		}

		// Simple Field (pin)
		pin := uint8(m.Pin)
		_pinErr := writeBuffer.WriteUint8("pin", 4, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for data")
			}
			for _, _element := range m.Data {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for data")
			}
		}

		if popErr := writeBuffer.PopContext("FirmataMessageAnalogIO"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageAnalogIO")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *FirmataMessageAnalogIO) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
