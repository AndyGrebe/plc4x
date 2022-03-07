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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type MonitoredSALShortFormBasicMode struct {
	*MonitoredSAL
	Counts        byte
	BridgeCount   *BridgeCount
	NetworkNumber *NetworkNumber
	NoCounts      *byte
	Application   ApplicationIdContainer
}

// The corresponding interface
type IMonitoredSALShortFormBasicMode interface {
	// GetCounts returns Counts
	GetCounts() byte
	// GetBridgeCount returns BridgeCount
	GetBridgeCount() *BridgeCount
	// GetNetworkNumber returns NetworkNumber
	GetNetworkNumber() *NetworkNumber
	// GetNoCounts returns NoCounts
	GetNoCounts() *byte
	// GetApplication returns Application
	GetApplication() ApplicationIdContainer
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

func (m *MonitoredSALShortFormBasicMode) InitializeParent(parent *MonitoredSAL, salType byte, salData *SALData) {
	m.MonitoredSAL.SalType = salType
	m.MonitoredSAL.SalData = salData
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *MonitoredSALShortFormBasicMode) GetCounts() byte {
	return m.Counts
}

func (m *MonitoredSALShortFormBasicMode) GetBridgeCount() *BridgeCount {
	return m.BridgeCount
}

func (m *MonitoredSALShortFormBasicMode) GetNetworkNumber() *NetworkNumber {
	return m.NetworkNumber
}

func (m *MonitoredSALShortFormBasicMode) GetNoCounts() *byte {
	return m.NoCounts
}

func (m *MonitoredSALShortFormBasicMode) GetApplication() ApplicationIdContainer {
	return m.Application
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewMonitoredSALShortFormBasicMode factory function for MonitoredSALShortFormBasicMode
func NewMonitoredSALShortFormBasicMode(counts byte, bridgeCount *BridgeCount, networkNumber *NetworkNumber, noCounts *byte, application ApplicationIdContainer, salType byte, salData *SALData) *MonitoredSAL {
	child := &MonitoredSALShortFormBasicMode{
		Counts:        counts,
		BridgeCount:   bridgeCount,
		NetworkNumber: networkNumber,
		NoCounts:      noCounts,
		Application:   application,
		MonitoredSAL:  NewMonitoredSAL(salType, salData),
	}
	child.Child = child
	return child.MonitoredSAL
}

func CastMonitoredSALShortFormBasicMode(structType interface{}) *MonitoredSALShortFormBasicMode {
	castFunc := func(typ interface{}) *MonitoredSALShortFormBasicMode {
		if casted, ok := typ.(MonitoredSALShortFormBasicMode); ok {
			return &casted
		}
		if casted, ok := typ.(*MonitoredSALShortFormBasicMode); ok {
			return casted
		}
		if casted, ok := typ.(MonitoredSAL); ok {
			return CastMonitoredSALShortFormBasicMode(casted.Child)
		}
		if casted, ok := typ.(*MonitoredSAL); ok {
			return CastMonitoredSALShortFormBasicMode(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MonitoredSALShortFormBasicMode) GetTypeName() string {
	return "MonitoredSALShortFormBasicMode"
}

func (m *MonitoredSALShortFormBasicMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *MonitoredSALShortFormBasicMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Optional Field (bridgeCount)
	if m.BridgeCount != nil {
		lengthInBits += (*m.BridgeCount).GetLengthInBits()
	}

	// Optional Field (networkNumber)
	if m.NetworkNumber != nil {
		lengthInBits += (*m.NetworkNumber).GetLengthInBits()
	}

	// Optional Field (noCounts)
	if m.NoCounts != nil {
		lengthInBits += 8
	}

	// Simple field (application)
	lengthInBits += 8

	return lengthInBits
}

func (m *MonitoredSALShortFormBasicMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MonitoredSALShortFormBasicModeParse(readBuffer utils.ReadBuffer) (*MonitoredSAL, error) {
	if pullErr := readBuffer.PullContext("MonitoredSALShortFormBasicMode"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Peek Field (counts)
	currentPos = readBuffer.GetPos()
	counts, _err := readBuffer.ReadByte("counts")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'counts' field")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (bridgeCount) (Can be skipped, if a given expression evaluates to false)
	var bridgeCount *BridgeCount = nil
	if bool((counts) != (0x00)) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("bridgeCount"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BridgeCountParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'bridgeCount' field")
		default:
			bridgeCount = CastBridgeCount(_val)
			if closeErr := readBuffer.CloseContext("bridgeCount"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (networkNumber) (Can be skipped, if a given expression evaluates to false)
	var networkNumber *NetworkNumber = nil
	if bool((counts) != (0x00)) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("networkNumber"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := NetworkNumberParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'networkNumber' field")
		default:
			networkNumber = CastNetworkNumber(_val)
			if closeErr := readBuffer.CloseContext("networkNumber"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (noCounts) (Can be skipped, if a given expression evaluates to false)
	var noCounts *byte = nil
	if bool((counts) == (0x00)) {
		_val, _err := readBuffer.ReadByte("noCounts")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'noCounts' field")
		}
		noCounts = &_val
	}

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

	if closeErr := readBuffer.CloseContext("MonitoredSALShortFormBasicMode"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MonitoredSALShortFormBasicMode{
		Counts:        counts,
		BridgeCount:   CastBridgeCount(bridgeCount),
		NetworkNumber: CastNetworkNumber(networkNumber),
		NoCounts:      noCounts,
		Application:   application,
		MonitoredSAL:  &MonitoredSAL{},
	}
	_child.MonitoredSAL.Child = _child
	return _child.MonitoredSAL, nil
}

func (m *MonitoredSALShortFormBasicMode) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredSALShortFormBasicMode"); pushErr != nil {
			return pushErr
		}

		// Optional Field (bridgeCount) (Can be skipped, if the value is null)
		var bridgeCount *BridgeCount = nil
		if m.BridgeCount != nil {
			if pushErr := writeBuffer.PushContext("bridgeCount"); pushErr != nil {
				return pushErr
			}
			bridgeCount = m.BridgeCount
			_bridgeCountErr := bridgeCount.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("bridgeCount"); popErr != nil {
				return popErr
			}
			if _bridgeCountErr != nil {
				return errors.Wrap(_bridgeCountErr, "Error serializing 'bridgeCount' field")
			}
		}

		// Optional Field (networkNumber) (Can be skipped, if the value is null)
		var networkNumber *NetworkNumber = nil
		if m.NetworkNumber != nil {
			if pushErr := writeBuffer.PushContext("networkNumber"); pushErr != nil {
				return pushErr
			}
			networkNumber = m.NetworkNumber
			_networkNumberErr := networkNumber.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("networkNumber"); popErr != nil {
				return popErr
			}
			if _networkNumberErr != nil {
				return errors.Wrap(_networkNumberErr, "Error serializing 'networkNumber' field")
			}
		}

		// Optional Field (noCounts) (Can be skipped, if the value is null)
		var noCounts *byte = nil
		if m.NoCounts != nil {
			noCounts = m.NoCounts
			_noCountsErr := writeBuffer.WriteByte("noCounts", *(noCounts))
			if _noCountsErr != nil {
				return errors.Wrap(_noCountsErr, "Error serializing 'noCounts' field")
			}
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

		if popErr := writeBuffer.PopContext("MonitoredSALShortFormBasicMode"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MonitoredSALShortFormBasicMode) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
