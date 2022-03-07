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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CBusPointToMultiPointCommandNormal_CR byte = 0xD

// The data-structure of this message
type CBusPointToMultiPointCommandNormal struct {
	*CBusPointToMultiPointCommand
	Application ApplicationIdContainer
	SalData     *SALData
	Crc         *Checksum
	PeekAlpha   byte
	Alpha       *Alpha

	// Arguments.
	Srchk bool
}

// The corresponding interface
type ICBusPointToMultiPointCommandNormal interface {
	// GetApplication returns Application
	GetApplication() ApplicationIdContainer
	// GetSalData returns SalData
	GetSalData() *SALData
	// GetCrc returns Crc
	GetCrc() *Checksum
	// GetPeekAlpha returns PeekAlpha
	GetPeekAlpha() byte
	// GetAlpha returns Alpha
	GetAlpha() *Alpha
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////

func (m *CBusPointToMultiPointCommandNormal) InitializeParent(parent *CBusPointToMultiPointCommand, peekedApplication byte) {
	m.CBusPointToMultiPointCommand.PeekedApplication = peekedApplication
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CBusPointToMultiPointCommandNormal) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *CBusPointToMultiPointCommandNormal) GetSalData() *SALData {
	return m.SalData
}

func (m *CBusPointToMultiPointCommandNormal) GetCrc() *Checksum {
	return m.Crc
}

func (m *CBusPointToMultiPointCommandNormal) GetPeekAlpha() byte {
	return m.PeekAlpha
}

func (m *CBusPointToMultiPointCommandNormal) GetAlpha() *Alpha {
	return m.Alpha
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCBusPointToMultiPointCommandNormal factory function for CBusPointToMultiPointCommandNormal
func NewCBusPointToMultiPointCommandNormal(application ApplicationIdContainer, salData *SALData, crc *Checksum, peekAlpha byte, alpha *Alpha, peekedApplication byte, srchk bool) *CBusPointToMultiPointCommand {
	child := &CBusPointToMultiPointCommandNormal{
		Application:                  application,
		SalData:                      salData,
		Crc:                          crc,
		PeekAlpha:                    peekAlpha,
		Alpha:                        alpha,
		CBusPointToMultiPointCommand: NewCBusPointToMultiPointCommand(peekedApplication, srchk),
	}
	child.Child = child
	return child.CBusPointToMultiPointCommand
}

func CastCBusPointToMultiPointCommandNormal(structType interface{}) *CBusPointToMultiPointCommandNormal {
	castFunc := func(typ interface{}) *CBusPointToMultiPointCommandNormal {
		if casted, ok := typ.(CBusPointToMultiPointCommandNormal); ok {
			return &casted
		}
		if casted, ok := typ.(*CBusPointToMultiPointCommandNormal); ok {
			return casted
		}
		if casted, ok := typ.(CBusPointToMultiPointCommand); ok {
			return CastCBusPointToMultiPointCommandNormal(casted.Child)
		}
		if casted, ok := typ.(*CBusPointToMultiPointCommand); ok {
			return CastCBusPointToMultiPointCommandNormal(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CBusPointToMultiPointCommandNormal) GetTypeName() string {
	return "CBusPointToMultiPointCommandNormal"
}

func (m *CBusPointToMultiPointCommandNormal) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CBusPointToMultiPointCommandNormal) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (application)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (salData)
	lengthInBits += m.SalData.GetLengthInBits()

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += (*m.Crc).GetLengthInBits()
	}

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += (*m.Alpha).GetLengthInBits()
	}

	// Const Field (cr)
	lengthInBits += 8

	return lengthInBits
}

func (m *CBusPointToMultiPointCommandNormal) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusPointToMultiPointCommandNormalParse(readBuffer utils.ReadBuffer, srchk bool) (*CBusPointToMultiPointCommand, error) {
	if pullErr := readBuffer.PullContext("CBusPointToMultiPointCommandNormal"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, pullErr
	}
	_application, _applicationErr := ApplicationIdContainerParse(readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, closeErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != byte(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (salData)
	if pullErr := readBuffer.PullContext("salData"); pullErr != nil {
		return nil, pullErr
	}
	_salData, _salDataErr := SALDataParse(readBuffer)
	if _salDataErr != nil {
		return nil, errors.Wrap(_salDataErr, "Error parsing 'salData' field")
	}
	salData := CastSALData(_salData)
	if closeErr := readBuffer.CloseContext("salData"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (crc) (Can be skipped, if a given expression evaluates to false)
	var crc *Checksum = nil
	if srchk {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("crc"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := ChecksumParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'crc' field")
		default:
			crc = CastChecksum(_val)
			if closeErr := readBuffer.CloseContext("crc"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Peek Field (peekAlpha)
	currentPos = readBuffer.GetPos()
	peekAlpha, _err := readBuffer.ReadByte("peekAlpha")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'peekAlpha' field")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (alpha) (Can be skipped, if a given expression evaluates to false)
	var alpha *Alpha = nil
	if bool(bool(bool((peekAlpha) >= (0x67)))) && bool(bool(bool((peekAlpha) <= (0x7A)))) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := AlphaParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'alpha' field")
		default:
			alpha = CastAlpha(_val)
			if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field")
	}
	if cr != CBusPointToMultiPointCommandNormal_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CBusPointToMultiPointCommandNormal_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	if closeErr := readBuffer.CloseContext("CBusPointToMultiPointCommandNormal"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CBusPointToMultiPointCommandNormal{
		Application:                  application,
		SalData:                      CastSALData(salData),
		Crc:                          CastChecksum(crc),
		PeekAlpha:                    peekAlpha,
		Alpha:                        CastAlpha(alpha),
		CBusPointToMultiPointCommand: &CBusPointToMultiPointCommand{},
	}
	_child.CBusPointToMultiPointCommand.Child = _child
	return _child.CBusPointToMultiPointCommand, nil
}

func (m *CBusPointToMultiPointCommandNormal) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToMultiPointCommandNormal"); pushErr != nil {
			return pushErr
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return pushErr
		}
		_applicationErr := m.Application.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return popErr
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (salData)
		if pushErr := writeBuffer.PushContext("salData"); pushErr != nil {
			return pushErr
		}
		_salDataErr := m.SalData.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("salData"); popErr != nil {
			return popErr
		}
		if _salDataErr != nil {
			return errors.Wrap(_salDataErr, "Error serializing 'salData' field")
		}

		// Optional Field (crc) (Can be skipped, if the value is null)
		var crc *Checksum = nil
		if m.Crc != nil {
			if pushErr := writeBuffer.PushContext("crc"); pushErr != nil {
				return pushErr
			}
			crc = m.Crc
			_crcErr := crc.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("crc"); popErr != nil {
				return popErr
			}
			if _crcErr != nil {
				return errors.Wrap(_crcErr, "Error serializing 'crc' field")
			}
		}

		// Optional Field (alpha) (Can be skipped, if the value is null)
		var alpha *Alpha = nil
		if m.Alpha != nil {
			if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
				return pushErr
			}
			alpha = m.Alpha
			_alphaErr := alpha.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
				return popErr
			}
			if _alphaErr != nil {
				return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
			}
		}

		// Const Field (cr)
		_crErr := writeBuffer.WriteByte("cr", 0xD)
		if _crErr != nil {
			return errors.Wrap(_crErr, "Error serializing 'cr' field")
		}

		if popErr := writeBuffer.PopContext("CBusPointToMultiPointCommandNormal"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CBusPointToMultiPointCommandNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
