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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CALDataReplyStatusExtended is the data-structure of this message
type CALDataReplyStatusExtended struct {
	*CALData
	Encoding    uint8
	Application ApplicationIdContainer
	BlockStart  uint8
	Data        []byte
}

// ICALDataReplyStatusExtended is the corresponding interface of CALDataReplyStatusExtended
type ICALDataReplyStatusExtended interface {
	ICALData
	// GetEncoding returns Encoding (property field)
	GetEncoding() uint8
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetBlockStart returns BlockStart (property field)
	GetBlockStart() uint8
	// GetData returns Data (property field)
	GetData() []byte
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

func (m *CALDataReplyStatusExtended) InitializeParent(parent *CALData, commandTypeContainer CALCommandTypeContainer) {
	m.CALData.CommandTypeContainer = commandTypeContainer
}

func (m *CALDataReplyStatusExtended) GetParent() *CALData {
	return m.CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CALDataReplyStatusExtended) GetEncoding() uint8 {
	return m.Encoding
}

func (m *CALDataReplyStatusExtended) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *CALDataReplyStatusExtended) GetBlockStart() uint8 {
	return m.BlockStart
}

func (m *CALDataReplyStatusExtended) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataReplyStatusExtended factory function for CALDataReplyStatusExtended
func NewCALDataReplyStatusExtended(encoding uint8, application ApplicationIdContainer, blockStart uint8, data []byte, commandTypeContainer CALCommandTypeContainer) *CALDataReplyStatusExtended {
	_result := &CALDataReplyStatusExtended{
		Encoding:    encoding,
		Application: application,
		BlockStart:  blockStart,
		Data:        data,
		CALData:     NewCALData(commandTypeContainer),
	}
	_result.Child = _result
	return _result
}

func CastCALDataReplyStatusExtended(structType interface{}) *CALDataReplyStatusExtended {
	if casted, ok := structType.(CALDataReplyStatusExtended); ok {
		return &casted
	}
	if casted, ok := structType.(*CALDataReplyStatusExtended); ok {
		return casted
	}
	if casted, ok := structType.(CALData); ok {
		return CastCALDataReplyStatusExtended(casted.Child)
	}
	if casted, ok := structType.(*CALData); ok {
		return CastCALDataReplyStatusExtended(casted.Child)
	}
	return nil
}

func (m *CALDataReplyStatusExtended) GetTypeName() string {
	return "CALDataReplyStatusExtended"
}

func (m *CALDataReplyStatusExtended) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALDataReplyStatusExtended) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (encoding)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Simple field (blockStart)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *CALDataReplyStatusExtended) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataReplyStatusExtendedParse(readBuffer utils.ReadBuffer, commandTypeContainer CALCommandTypeContainer) (*CALDataReplyStatusExtended, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataReplyStatusExtended"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (encoding)
	_encoding, _encodingErr := readBuffer.ReadUint8("encoding", 8)
	if _encodingErr != nil {
		return nil, errors.Wrap(_encodingErr, "Error parsing 'encoding' field")
	}
	encoding := _encoding

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

	// Simple Field (blockStart)
	_blockStart, _blockStartErr := readBuffer.ReadUint8("blockStart", 8)
	if _blockStartErr != nil {
		return nil, errors.Wrap(_blockStartErr, "Error parsing 'blockStart' field")
	}
	blockStart := _blockStart
	// Byte Array field (data)
	numberOfBytesdata := int(commandTypeContainer.NumBytes())
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("CALDataReplyStatusExtended"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CALDataReplyStatusExtended{
		Encoding:    encoding,
		Application: application,
		BlockStart:  blockStart,
		Data:        data,
		CALData:     &CALData{},
	}
	_child.CALData.Child = _child
	return _child, nil
}

func (m *CALDataReplyStatusExtended) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataReplyStatusExtended"); pushErr != nil {
			return pushErr
		}

		// Simple Field (encoding)
		encoding := uint8(m.Encoding)
		_encodingErr := writeBuffer.WriteUint8("encoding", 8, (encoding))
		if _encodingErr != nil {
			return errors.Wrap(_encodingErr, "Error serializing 'encoding' field")
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

		// Simple Field (blockStart)
		blockStart := uint8(m.BlockStart)
		_blockStartErr := writeBuffer.WriteUint8("blockStart", 8, (blockStart))
		if _blockStartErr != nil {
			return errors.Wrap(_blockStartErr, "Error serializing 'blockStart' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("CALDataReplyStatusExtended"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CALDataReplyStatusExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}