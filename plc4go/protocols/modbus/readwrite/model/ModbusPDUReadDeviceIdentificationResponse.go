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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const ModbusPDUReadDeviceIdentificationResponse_MEITYPE uint8 = 0x0E

// ModbusPDUReadDeviceIdentificationResponse is the data-structure of this message
type ModbusPDUReadDeviceIdentificationResponse struct {
	*ModbusPDU
	Level            ModbusDeviceInformationLevel
	IndividualAccess bool
	ConformityLevel  ModbusDeviceInformationConformityLevel
	MoreFollows      ModbusDeviceInformationMoreFollows
	NextObjectId     uint8
	Objects          []*ModbusDeviceInformationObject
}

// IModbusPDUReadDeviceIdentificationResponse is the corresponding interface of ModbusPDUReadDeviceIdentificationResponse
type IModbusPDUReadDeviceIdentificationResponse interface {
	IModbusPDU
	// GetLevel returns Level (property field)
	GetLevel() ModbusDeviceInformationLevel
	// GetIndividualAccess returns IndividualAccess (property field)
	GetIndividualAccess() bool
	// GetConformityLevel returns ConformityLevel (property field)
	GetConformityLevel() ModbusDeviceInformationConformityLevel
	// GetMoreFollows returns MoreFollows (property field)
	GetMoreFollows() ModbusDeviceInformationMoreFollows
	// GetNextObjectId returns NextObjectId (property field)
	GetNextObjectId() uint8
	// GetObjects returns Objects (property field)
	GetObjects() []*ModbusDeviceInformationObject
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

func (m *ModbusPDUReadDeviceIdentificationResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetFunctionFlag() uint8 {
	return 0x2B
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUReadDeviceIdentificationResponse) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ModbusPDUReadDeviceIdentificationResponse) GetLevel() ModbusDeviceInformationLevel {
	return m.Level
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetIndividualAccess() bool {
	return m.IndividualAccess
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetConformityLevel() ModbusDeviceInformationConformityLevel {
	return m.ConformityLevel
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetMoreFollows() ModbusDeviceInformationMoreFollows {
	return m.MoreFollows
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetNextObjectId() uint8 {
	return m.NextObjectId
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetObjects() []*ModbusDeviceInformationObject {
	return m.Objects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *ModbusPDUReadDeviceIdentificationResponse) GetMeiType() uint8 {
	return ModbusPDUReadDeviceIdentificationResponse_MEITYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadDeviceIdentificationResponse factory function for ModbusPDUReadDeviceIdentificationResponse
func NewModbusPDUReadDeviceIdentificationResponse(level ModbusDeviceInformationLevel, individualAccess bool, conformityLevel ModbusDeviceInformationConformityLevel, moreFollows ModbusDeviceInformationMoreFollows, nextObjectId uint8, objects []*ModbusDeviceInformationObject) *ModbusPDUReadDeviceIdentificationResponse {
	_result := &ModbusPDUReadDeviceIdentificationResponse{
		Level:            level,
		IndividualAccess: individualAccess,
		ConformityLevel:  conformityLevel,
		MoreFollows:      moreFollows,
		NextObjectId:     nextObjectId,
		Objects:          objects,
		ModbusPDU:        NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUReadDeviceIdentificationResponse(structType interface{}) *ModbusPDUReadDeviceIdentificationResponse {
	if casted, ok := structType.(ModbusPDUReadDeviceIdentificationResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUReadDeviceIdentificationResponse); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUReadDeviceIdentificationResponse(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUReadDeviceIdentificationResponse(casted.Child)
	}
	return nil
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetTypeName() string {
	return "ModbusPDUReadDeviceIdentificationResponse"
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (meiType)
	lengthInBits += 8

	// Simple field (level)
	lengthInBits += 8

	// Simple field (individualAccess)
	lengthInBits += 1

	// Simple field (conformityLevel)
	lengthInBits += 7

	// Simple field (moreFollows)
	lengthInBits += 8

	// Simple field (nextObjectId)
	lengthInBits += 8

	// Implicit Field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.Objects) > 0 {
		for i, element := range m.Objects {
			last := i == len(m.Objects)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadDeviceIdentificationResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUReadDeviceIdentificationResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadDeviceIdentificationResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (meiType)
	meiType, _meiTypeErr := readBuffer.ReadUint8("meiType", 8)
	if _meiTypeErr != nil {
		return nil, errors.Wrap(_meiTypeErr, "Error parsing 'meiType' field")
	}
	if meiType != ModbusPDUReadDeviceIdentificationResponse_MEITYPE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ModbusPDUReadDeviceIdentificationResponse_MEITYPE) + " but got " + fmt.Sprintf("%d", meiType))
	}

	// Simple Field (level)
	if pullErr := readBuffer.PullContext("level"); pullErr != nil {
		return nil, pullErr
	}
	_level, _levelErr := ModbusDeviceInformationLevelParse(readBuffer)
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field")
	}
	level := _level
	if closeErr := readBuffer.CloseContext("level"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (individualAccess)
	_individualAccess, _individualAccessErr := readBuffer.ReadBit("individualAccess")
	if _individualAccessErr != nil {
		return nil, errors.Wrap(_individualAccessErr, "Error parsing 'individualAccess' field")
	}
	individualAccess := _individualAccess

	// Simple Field (conformityLevel)
	if pullErr := readBuffer.PullContext("conformityLevel"); pullErr != nil {
		return nil, pullErr
	}
	_conformityLevel, _conformityLevelErr := ModbusDeviceInformationConformityLevelParse(readBuffer)
	if _conformityLevelErr != nil {
		return nil, errors.Wrap(_conformityLevelErr, "Error parsing 'conformityLevel' field")
	}
	conformityLevel := _conformityLevel
	if closeErr := readBuffer.CloseContext("conformityLevel"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (moreFollows)
	if pullErr := readBuffer.PullContext("moreFollows"); pullErr != nil {
		return nil, pullErr
	}
	_moreFollows, _moreFollowsErr := ModbusDeviceInformationMoreFollowsParse(readBuffer)
	if _moreFollowsErr != nil {
		return nil, errors.Wrap(_moreFollowsErr, "Error parsing 'moreFollows' field")
	}
	moreFollows := _moreFollows
	if closeErr := readBuffer.CloseContext("moreFollows"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (nextObjectId)
	_nextObjectId, _nextObjectIdErr := readBuffer.ReadUint8("nextObjectId", 8)
	if _nextObjectIdErr != nil {
		return nil, errors.Wrap(_nextObjectIdErr, "Error parsing 'nextObjectId' field")
	}
	nextObjectId := _nextObjectId

	// Implicit Field (numberOfObjects) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	_ = numberOfObjects
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field")
	}

	// Array field (objects)
	if pullErr := readBuffer.PullContext("objects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	objects := make([]*ModbusDeviceInformationObject, numberOfObjects)
	{
		for curItem := uint16(0); curItem < uint16(numberOfObjects); curItem++ {
			_item, _err := ModbusDeviceInformationObjectParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'objects' field")
			}
			objects[curItem] = CastModbusDeviceInformationObject(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("objects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadDeviceIdentificationResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadDeviceIdentificationResponse{
		Level:            level,
		IndividualAccess: individualAccess,
		ConformityLevel:  conformityLevel,
		MoreFollows:      moreFollows,
		NextObjectId:     nextObjectId,
		Objects:          objects,
		ModbusPDU:        &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUReadDeviceIdentificationResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadDeviceIdentificationResponse"); pushErr != nil {
			return pushErr
		}

		// Const Field (meiType)
		_meiTypeErr := writeBuffer.WriteUint8("meiType", 8, 0x0E)
		if _meiTypeErr != nil {
			return errors.Wrap(_meiTypeErr, "Error serializing 'meiType' field")
		}

		// Simple Field (level)
		if pushErr := writeBuffer.PushContext("level"); pushErr != nil {
			return pushErr
		}
		_levelErr := m.Level.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("level"); popErr != nil {
			return popErr
		}
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
		}

		// Simple Field (individualAccess)
		individualAccess := bool(m.IndividualAccess)
		_individualAccessErr := writeBuffer.WriteBit("individualAccess", (individualAccess))
		if _individualAccessErr != nil {
			return errors.Wrap(_individualAccessErr, "Error serializing 'individualAccess' field")
		}

		// Simple Field (conformityLevel)
		if pushErr := writeBuffer.PushContext("conformityLevel"); pushErr != nil {
			return pushErr
		}
		_conformityLevelErr := m.ConformityLevel.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("conformityLevel"); popErr != nil {
			return popErr
		}
		if _conformityLevelErr != nil {
			return errors.Wrap(_conformityLevelErr, "Error serializing 'conformityLevel' field")
		}

		// Simple Field (moreFollows)
		if pushErr := writeBuffer.PushContext("moreFollows"); pushErr != nil {
			return pushErr
		}
		_moreFollowsErr := m.MoreFollows.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("moreFollows"); popErr != nil {
			return popErr
		}
		if _moreFollowsErr != nil {
			return errors.Wrap(_moreFollowsErr, "Error serializing 'moreFollows' field")
		}

		// Simple Field (nextObjectId)
		nextObjectId := uint8(m.NextObjectId)
		_nextObjectIdErr := writeBuffer.WriteUint8("nextObjectId", 8, (nextObjectId))
		if _nextObjectIdErr != nil {
			return errors.Wrap(_nextObjectIdErr, "Error serializing 'nextObjectId' field")
		}

		// Implicit Field (numberOfObjects) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numberOfObjects := uint8(uint8(len(m.GetObjects())))
		_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, (numberOfObjects))
		if _numberOfObjectsErr != nil {
			return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
		}

		// Array Field (objects)
		if m.Objects != nil {
			if pushErr := writeBuffer.PushContext("objects", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Objects {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'objects' field")
				}
			}
			if popErr := writeBuffer.PopContext("objects", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadDeviceIdentificationResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadDeviceIdentificationResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}