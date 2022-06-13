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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const AlarmMessageObjectQueryType_VARIABLESPEC uint8 = 0x12

// AlarmMessageObjectQueryType is the data-structure of this message
type AlarmMessageObjectQueryType struct {
	LengthDataset  uint8
	EventState     *State
	AckStateGoing  *State
	AckStateComing *State
	TimeComing     *DateAndTime
	ValueComing    *AssociatedValueType
	TimeGoing      *DateAndTime
	ValueGoing     *AssociatedValueType
}

// IAlarmMessageObjectQueryType is the corresponding interface of AlarmMessageObjectQueryType
type IAlarmMessageObjectQueryType interface {
	// GetLengthDataset returns LengthDataset (property field)
	GetLengthDataset() uint8
	// GetEventState returns EventState (property field)
	GetEventState() *State
	// GetAckStateGoing returns AckStateGoing (property field)
	GetAckStateGoing() *State
	// GetAckStateComing returns AckStateComing (property field)
	GetAckStateComing() *State
	// GetTimeComing returns TimeComing (property field)
	GetTimeComing() *DateAndTime
	// GetValueComing returns ValueComing (property field)
	GetValueComing() *AssociatedValueType
	// GetTimeGoing returns TimeGoing (property field)
	GetTimeGoing() *DateAndTime
	// GetValueGoing returns ValueGoing (property field)
	GetValueGoing() *AssociatedValueType
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *AlarmMessageObjectQueryType) GetLengthDataset() uint8 {
	return m.LengthDataset
}

func (m *AlarmMessageObjectQueryType) GetEventState() *State {
	return m.EventState
}

func (m *AlarmMessageObjectQueryType) GetAckStateGoing() *State {
	return m.AckStateGoing
}

func (m *AlarmMessageObjectQueryType) GetAckStateComing() *State {
	return m.AckStateComing
}

func (m *AlarmMessageObjectQueryType) GetTimeComing() *DateAndTime {
	return m.TimeComing
}

func (m *AlarmMessageObjectQueryType) GetValueComing() *AssociatedValueType {
	return m.ValueComing
}

func (m *AlarmMessageObjectQueryType) GetTimeGoing() *DateAndTime {
	return m.TimeGoing
}

func (m *AlarmMessageObjectQueryType) GetValueGoing() *AssociatedValueType {
	return m.ValueGoing
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *AlarmMessageObjectQueryType) GetVariableSpec() uint8 {
	return AlarmMessageObjectQueryType_VARIABLESPEC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessageObjectQueryType factory function for AlarmMessageObjectQueryType
func NewAlarmMessageObjectQueryType(lengthDataset uint8, eventState *State, ackStateGoing *State, ackStateComing *State, timeComing *DateAndTime, valueComing *AssociatedValueType, timeGoing *DateAndTime, valueGoing *AssociatedValueType) *AlarmMessageObjectQueryType {
	return &AlarmMessageObjectQueryType{LengthDataset: lengthDataset, EventState: eventState, AckStateGoing: ackStateGoing, AckStateComing: ackStateComing, TimeComing: timeComing, ValueComing: valueComing, TimeGoing: timeGoing, ValueGoing: valueGoing}
}

func CastAlarmMessageObjectQueryType(structType interface{}) *AlarmMessageObjectQueryType {
	if casted, ok := structType.(AlarmMessageObjectQueryType); ok {
		return &casted
	}
	if casted, ok := structType.(*AlarmMessageObjectQueryType); ok {
		return casted
	}
	return nil
}

func (m *AlarmMessageObjectQueryType) GetTypeName() string {
	return "AlarmMessageObjectQueryType"
}

func (m *AlarmMessageObjectQueryType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AlarmMessageObjectQueryType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (lengthDataset)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Const Field (variableSpec)
	lengthInBits += 8

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits()

	// Simple field (ackStateGoing)
	lengthInBits += m.AckStateGoing.GetLengthInBits()

	// Simple field (ackStateComing)
	lengthInBits += m.AckStateComing.GetLengthInBits()

	// Simple field (timeComing)
	lengthInBits += m.TimeComing.GetLengthInBits()

	// Simple field (valueComing)
	lengthInBits += m.ValueComing.GetLengthInBits()

	// Simple field (timeGoing)
	lengthInBits += m.TimeGoing.GetLengthInBits()

	// Simple field (valueGoing)
	lengthInBits += m.ValueGoing.GetLengthInBits()

	return lengthInBits
}

func (m *AlarmMessageObjectQueryType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AlarmMessageObjectQueryTypeParse(readBuffer utils.ReadBuffer) (*AlarmMessageObjectQueryType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageObjectQueryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageObjectQueryType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lengthDataset)
	_lengthDataset, _lengthDatasetErr := readBuffer.ReadUint8("lengthDataset", 8)
	if _lengthDatasetErr != nil {
		return nil, errors.Wrap(_lengthDatasetErr, "Error parsing 'lengthDataset' field")
	}
	lengthDataset := _lengthDataset

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Const Field (variableSpec)
	variableSpec, _variableSpecErr := readBuffer.ReadUint8("variableSpec", 8)
	if _variableSpecErr != nil {
		return nil, errors.Wrap(_variableSpecErr, "Error parsing 'variableSpec' field")
	}
	if variableSpec != AlarmMessageObjectQueryType_VARIABLESPEC {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AlarmMessageObjectQueryType_VARIABLESPEC) + " but got " + fmt.Sprintf("%d", variableSpec))
	}

	// Simple Field (eventState)
	if pullErr := readBuffer.PullContext("eventState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventState")
	}
	_eventState, _eventStateErr := StateParse(readBuffer)
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field")
	}
	eventState := CastState(_eventState)
	if closeErr := readBuffer.CloseContext("eventState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventState")
	}

	// Simple Field (ackStateGoing)
	if pullErr := readBuffer.PullContext("ackStateGoing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackStateGoing")
	}
	_ackStateGoing, _ackStateGoingErr := StateParse(readBuffer)
	if _ackStateGoingErr != nil {
		return nil, errors.Wrap(_ackStateGoingErr, "Error parsing 'ackStateGoing' field")
	}
	ackStateGoing := CastState(_ackStateGoing)
	if closeErr := readBuffer.CloseContext("ackStateGoing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackStateGoing")
	}

	// Simple Field (ackStateComing)
	if pullErr := readBuffer.PullContext("ackStateComing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackStateComing")
	}
	_ackStateComing, _ackStateComingErr := StateParse(readBuffer)
	if _ackStateComingErr != nil {
		return nil, errors.Wrap(_ackStateComingErr, "Error parsing 'ackStateComing' field")
	}
	ackStateComing := CastState(_ackStateComing)
	if closeErr := readBuffer.CloseContext("ackStateComing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackStateComing")
	}

	// Simple Field (timeComing)
	if pullErr := readBuffer.PullContext("timeComing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeComing")
	}
	_timeComing, _timeComingErr := DateAndTimeParse(readBuffer)
	if _timeComingErr != nil {
		return nil, errors.Wrap(_timeComingErr, "Error parsing 'timeComing' field")
	}
	timeComing := CastDateAndTime(_timeComing)
	if closeErr := readBuffer.CloseContext("timeComing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeComing")
	}

	// Simple Field (valueComing)
	if pullErr := readBuffer.PullContext("valueComing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for valueComing")
	}
	_valueComing, _valueComingErr := AssociatedValueTypeParse(readBuffer)
	if _valueComingErr != nil {
		return nil, errors.Wrap(_valueComingErr, "Error parsing 'valueComing' field")
	}
	valueComing := CastAssociatedValueType(_valueComing)
	if closeErr := readBuffer.CloseContext("valueComing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for valueComing")
	}

	// Simple Field (timeGoing)
	if pullErr := readBuffer.PullContext("timeGoing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeGoing")
	}
	_timeGoing, _timeGoingErr := DateAndTimeParse(readBuffer)
	if _timeGoingErr != nil {
		return nil, errors.Wrap(_timeGoingErr, "Error parsing 'timeGoing' field")
	}
	timeGoing := CastDateAndTime(_timeGoing)
	if closeErr := readBuffer.CloseContext("timeGoing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeGoing")
	}

	// Simple Field (valueGoing)
	if pullErr := readBuffer.PullContext("valueGoing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for valueGoing")
	}
	_valueGoing, _valueGoingErr := AssociatedValueTypeParse(readBuffer)
	if _valueGoingErr != nil {
		return nil, errors.Wrap(_valueGoingErr, "Error parsing 'valueGoing' field")
	}
	valueGoing := CastAssociatedValueType(_valueGoing)
	if closeErr := readBuffer.CloseContext("valueGoing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for valueGoing")
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageObjectQueryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageObjectQueryType")
	}

	// Create the instance
	return NewAlarmMessageObjectQueryType(lengthDataset, eventState, ackStateGoing, ackStateComing, timeComing, valueComing, timeGoing, valueGoing), nil
}

func (m *AlarmMessageObjectQueryType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AlarmMessageObjectQueryType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageObjectQueryType")
	}

	// Simple Field (lengthDataset)
	lengthDataset := uint8(m.LengthDataset)
	_lengthDatasetErr := writeBuffer.WriteUint8("lengthDataset", 8, (lengthDataset))
	if _lengthDatasetErr != nil {
		return errors.Wrap(_lengthDatasetErr, "Error serializing 'lengthDataset' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint16("reserved", 16, uint16(0x0000))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Const Field (variableSpec)
	_variableSpecErr := writeBuffer.WriteUint8("variableSpec", 8, 0x12)
	if _variableSpecErr != nil {
		return errors.Wrap(_variableSpecErr, "Error serializing 'variableSpec' field")
	}

	// Simple Field (eventState)
	if pushErr := writeBuffer.PushContext("eventState"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for eventState")
	}
	_eventStateErr := writeBuffer.WriteSerializable(m.EventState)
	if popErr := writeBuffer.PopContext("eventState"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for eventState")
	}
	if _eventStateErr != nil {
		return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
	}

	// Simple Field (ackStateGoing)
	if pushErr := writeBuffer.PushContext("ackStateGoing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ackStateGoing")
	}
	_ackStateGoingErr := writeBuffer.WriteSerializable(m.AckStateGoing)
	if popErr := writeBuffer.PopContext("ackStateGoing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ackStateGoing")
	}
	if _ackStateGoingErr != nil {
		return errors.Wrap(_ackStateGoingErr, "Error serializing 'ackStateGoing' field")
	}

	// Simple Field (ackStateComing)
	if pushErr := writeBuffer.PushContext("ackStateComing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ackStateComing")
	}
	_ackStateComingErr := writeBuffer.WriteSerializable(m.AckStateComing)
	if popErr := writeBuffer.PopContext("ackStateComing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ackStateComing")
	}
	if _ackStateComingErr != nil {
		return errors.Wrap(_ackStateComingErr, "Error serializing 'ackStateComing' field")
	}

	// Simple Field (timeComing)
	if pushErr := writeBuffer.PushContext("timeComing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeComing")
	}
	_timeComingErr := writeBuffer.WriteSerializable(m.TimeComing)
	if popErr := writeBuffer.PopContext("timeComing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeComing")
	}
	if _timeComingErr != nil {
		return errors.Wrap(_timeComingErr, "Error serializing 'timeComing' field")
	}

	// Simple Field (valueComing)
	if pushErr := writeBuffer.PushContext("valueComing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for valueComing")
	}
	_valueComingErr := writeBuffer.WriteSerializable(m.ValueComing)
	if popErr := writeBuffer.PopContext("valueComing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for valueComing")
	}
	if _valueComingErr != nil {
		return errors.Wrap(_valueComingErr, "Error serializing 'valueComing' field")
	}

	// Simple Field (timeGoing)
	if pushErr := writeBuffer.PushContext("timeGoing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeGoing")
	}
	_timeGoingErr := writeBuffer.WriteSerializable(m.TimeGoing)
	if popErr := writeBuffer.PopContext("timeGoing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeGoing")
	}
	if _timeGoingErr != nil {
		return errors.Wrap(_timeGoingErr, "Error serializing 'timeGoing' field")
	}

	// Simple Field (valueGoing)
	if pushErr := writeBuffer.PushContext("valueGoing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for valueGoing")
	}
	_valueGoingErr := writeBuffer.WriteSerializable(m.ValueGoing)
	if popErr := writeBuffer.PopContext("valueGoing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for valueGoing")
	}
	if _valueGoingErr != nil {
		return errors.Wrap(_valueGoingErr, "Error serializing 'valueGoing' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageObjectQueryType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageObjectQueryType")
	}
	return nil
}

func (m *AlarmMessageObjectQueryType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
