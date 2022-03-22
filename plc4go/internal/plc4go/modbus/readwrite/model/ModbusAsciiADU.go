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
type ModbusAsciiADU struct {
	*ModbusADU
	Address uint8
	Pdu     *ModbusPDU

	// Arguments.
	Response bool
}

// The corresponding interface
type IModbusAsciiADU interface {
	IModbusADU
	// GetAddress returns Address (property field)
	GetAddress() uint8
	// GetPdu returns Pdu (property field)
	GetPdu() *ModbusPDU
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
func (m *ModbusAsciiADU) GetDriverType() DriverType {
	return DriverType_MODBUS_ASCII
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusAsciiADU) InitializeParent(parent *ModbusADU) {}

func (m *ModbusAsciiADU) GetParent() *ModbusADU {
	return m.ModbusADU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *ModbusAsciiADU) GetAddress() uint8 {
	return m.Address
}

func (m *ModbusAsciiADU) GetPdu() *ModbusPDU {
	return m.Pdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusAsciiADU factory function for ModbusAsciiADU
func NewModbusAsciiADU(address uint8, pdu *ModbusPDU, response bool) *ModbusAsciiADU {
	_result := &ModbusAsciiADU{
		Address:   address,
		Pdu:       pdu,
		ModbusADU: NewModbusADU(response),
	}
	_result.Child = _result
	return _result
}

func CastModbusAsciiADU(structType interface{}) *ModbusAsciiADU {
	if casted, ok := structType.(ModbusAsciiADU); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusAsciiADU); ok {
		return casted
	}
	if casted, ok := structType.(ModbusADU); ok {
		return CastModbusAsciiADU(casted.Child)
	}
	if casted, ok := structType.(*ModbusADU); ok {
		return CastModbusAsciiADU(casted.Child)
	}
	return nil
}

func (m *ModbusAsciiADU) GetTypeName() string {
	return "ModbusAsciiADU"
}

func (m *ModbusAsciiADU) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusAsciiADU) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (address)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.GetLengthInBits()

	// Checksum Field (checksum)
	lengthInBits += 8

	return lengthInBits
}

func (m *ModbusAsciiADU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusAsciiADUParse(readBuffer utils.ReadBuffer, driverType DriverType, response bool) (*ModbusAsciiADU, error) {
	if pullErr := readBuffer.PullContext("ModbusAsciiADU"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint8("address", 8)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}
	address := _address

	// Simple Field (pdu)
	if pullErr := readBuffer.PullContext("pdu"); pullErr != nil {
		return nil, pullErr
	}
	_pdu, _pduErr := ModbusPDUParse(readBuffer, bool(response))
	if _pduErr != nil {
		return nil, errors.Wrap(_pduErr, "Error parsing 'pdu' field")
	}
	pdu := CastModbusPDU(_pdu)
	if closeErr := readBuffer.CloseContext("pdu"); closeErr != nil {
		return nil, closeErr
	}

	// Checksum Field (checksum)
	{
		checksumRef, _checksumRefErr := readBuffer.ReadUint8("checksum", 8)
		if _checksumRefErr != nil {
			return nil, errors.Wrap(_checksumRefErr, "Error parsing 'checksum' field")
		}
		checksum, _checksumErr := AsciiLrcCheck(address, pdu)
		if _checksumErr != nil {
			return nil, errors.Wrap(_checksumErr, "Checksum verification failed")
		}
		if checksum != checksumRef {
			return nil, errors.Errorf("Checksum verification failed. Expected %x but got %x", checksumRef, checksum)
		}
	}

	if closeErr := readBuffer.CloseContext("ModbusAsciiADU"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusAsciiADU{
		Address:   address,
		Pdu:       CastModbusPDU(pdu),
		ModbusADU: &ModbusADU{},
	}
	_child.ModbusADU.Child = _child
	return _child, nil
}

func (m *ModbusAsciiADU) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusAsciiADU"); pushErr != nil {
			return pushErr
		}

		// Simple Field (address)
		address := uint8(m.Address)
		_addressErr := writeBuffer.WriteUint8("address", 8, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (pdu)
		if pushErr := writeBuffer.PushContext("pdu"); pushErr != nil {
			return pushErr
		}
		_pduErr := m.Pdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("pdu"); popErr != nil {
			return popErr
		}
		if _pduErr != nil {
			return errors.Wrap(_pduErr, "Error serializing 'pdu' field")
		}

		// Checksum Field (checksum) (Calculated)
		{
			_checksum, _checksumErr := AsciiLrcCheck(m.GetAddress(), m.GetPdu())
			if _checksumErr != nil {
				return errors.Wrap(_checksumErr, "Checksum calculation failed")
			}
			_checksumWriteErr := writeBuffer.WriteUint8("checksum", 8, (_checksum))
			if _checksumWriteErr != nil {
				return errors.Wrap(_checksumWriteErr, "Error serializing 'checksum' field")
			}
		}

		if popErr := writeBuffer.PopContext("ModbusAsciiADU"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusAsciiADU) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}