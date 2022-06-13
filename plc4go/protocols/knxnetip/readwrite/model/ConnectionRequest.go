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

// ConnectionRequest is the data-structure of this message
type ConnectionRequest struct {
	*KnxNetIpMessage
	HpaiDiscoveryEndpoint        *HPAIDiscoveryEndpoint
	HpaiDataEndpoint             *HPAIDataEndpoint
	ConnectionRequestInformation *ConnectionRequestInformation
}

// IConnectionRequest is the corresponding interface of ConnectionRequest
type IConnectionRequest interface {
	IKnxNetIpMessage
	// GetHpaiDiscoveryEndpoint returns HpaiDiscoveryEndpoint (property field)
	GetHpaiDiscoveryEndpoint() *HPAIDiscoveryEndpoint
	// GetHpaiDataEndpoint returns HpaiDataEndpoint (property field)
	GetHpaiDataEndpoint() *HPAIDataEndpoint
	// GetConnectionRequestInformation returns ConnectionRequestInformation (property field)
	GetConnectionRequestInformation() *ConnectionRequestInformation
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

func (m *ConnectionRequest) GetMsgType() uint16 {
	return 0x0205
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ConnectionRequest) InitializeParent(parent *KnxNetIpMessage) {}

func (m *ConnectionRequest) GetParent() *KnxNetIpMessage {
	return m.KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ConnectionRequest) GetHpaiDiscoveryEndpoint() *HPAIDiscoveryEndpoint {
	return m.HpaiDiscoveryEndpoint
}

func (m *ConnectionRequest) GetHpaiDataEndpoint() *HPAIDataEndpoint {
	return m.HpaiDataEndpoint
}

func (m *ConnectionRequest) GetConnectionRequestInformation() *ConnectionRequestInformation {
	return m.ConnectionRequestInformation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConnectionRequest factory function for ConnectionRequest
func NewConnectionRequest(hpaiDiscoveryEndpoint *HPAIDiscoveryEndpoint, hpaiDataEndpoint *HPAIDataEndpoint, connectionRequestInformation *ConnectionRequestInformation) *ConnectionRequest {
	_result := &ConnectionRequest{
		HpaiDiscoveryEndpoint:        hpaiDiscoveryEndpoint,
		HpaiDataEndpoint:             hpaiDataEndpoint,
		ConnectionRequestInformation: connectionRequestInformation,
		KnxNetIpMessage:              NewKnxNetIpMessage(),
	}
	_result.Child = _result
	return _result
}

func CastConnectionRequest(structType interface{}) *ConnectionRequest {
	if casted, ok := structType.(ConnectionRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ConnectionRequest); ok {
		return casted
	}
	if casted, ok := structType.(KnxNetIpMessage); ok {
		return CastConnectionRequest(casted.Child)
	}
	if casted, ok := structType.(*KnxNetIpMessage); ok {
		return CastConnectionRequest(casted.Child)
	}
	return nil
}

func (m *ConnectionRequest) GetTypeName() string {
	return "ConnectionRequest"
}

func (m *ConnectionRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ConnectionRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (hpaiDiscoveryEndpoint)
	lengthInBits += m.HpaiDiscoveryEndpoint.GetLengthInBits()

	// Simple field (hpaiDataEndpoint)
	lengthInBits += m.HpaiDataEndpoint.GetLengthInBits()

	// Simple field (connectionRequestInformation)
	lengthInBits += m.ConnectionRequestInformation.GetLengthInBits()

	return lengthInBits
}

func (m *ConnectionRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConnectionRequestParse(readBuffer utils.ReadBuffer) (*ConnectionRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (hpaiDiscoveryEndpoint)
	if pullErr := readBuffer.PullContext("hpaiDiscoveryEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hpaiDiscoveryEndpoint")
	}
	_hpaiDiscoveryEndpoint, _hpaiDiscoveryEndpointErr := HPAIDiscoveryEndpointParse(readBuffer)
	if _hpaiDiscoveryEndpointErr != nil {
		return nil, errors.Wrap(_hpaiDiscoveryEndpointErr, "Error parsing 'hpaiDiscoveryEndpoint' field")
	}
	hpaiDiscoveryEndpoint := CastHPAIDiscoveryEndpoint(_hpaiDiscoveryEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiDiscoveryEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hpaiDiscoveryEndpoint")
	}

	// Simple Field (hpaiDataEndpoint)
	if pullErr := readBuffer.PullContext("hpaiDataEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hpaiDataEndpoint")
	}
	_hpaiDataEndpoint, _hpaiDataEndpointErr := HPAIDataEndpointParse(readBuffer)
	if _hpaiDataEndpointErr != nil {
		return nil, errors.Wrap(_hpaiDataEndpointErr, "Error parsing 'hpaiDataEndpoint' field")
	}
	hpaiDataEndpoint := CastHPAIDataEndpoint(_hpaiDataEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiDataEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hpaiDataEndpoint")
	}

	// Simple Field (connectionRequestInformation)
	if pullErr := readBuffer.PullContext("connectionRequestInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for connectionRequestInformation")
	}
	_connectionRequestInformation, _connectionRequestInformationErr := ConnectionRequestInformationParse(readBuffer)
	if _connectionRequestInformationErr != nil {
		return nil, errors.Wrap(_connectionRequestInformationErr, "Error parsing 'connectionRequestInformation' field")
	}
	connectionRequestInformation := CastConnectionRequestInformation(_connectionRequestInformation)
	if closeErr := readBuffer.CloseContext("connectionRequestInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for connectionRequestInformation")
	}

	if closeErr := readBuffer.CloseContext("ConnectionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequest")
	}

	// Create a partially initialized instance
	_child := &ConnectionRequest{
		HpaiDiscoveryEndpoint:        CastHPAIDiscoveryEndpoint(hpaiDiscoveryEndpoint),
		HpaiDataEndpoint:             CastHPAIDataEndpoint(hpaiDataEndpoint),
		ConnectionRequestInformation: CastConnectionRequestInformation(connectionRequestInformation),
		KnxNetIpMessage:              &KnxNetIpMessage{},
	}
	_child.KnxNetIpMessage.Child = _child
	return _child, nil
}

func (m *ConnectionRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionRequest")
		}

		// Simple Field (hpaiDiscoveryEndpoint)
		if pushErr := writeBuffer.PushContext("hpaiDiscoveryEndpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hpaiDiscoveryEndpoint")
		}
		_hpaiDiscoveryEndpointErr := writeBuffer.WriteSerializable(m.HpaiDiscoveryEndpoint)
		if popErr := writeBuffer.PopContext("hpaiDiscoveryEndpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hpaiDiscoveryEndpoint")
		}
		if _hpaiDiscoveryEndpointErr != nil {
			return errors.Wrap(_hpaiDiscoveryEndpointErr, "Error serializing 'hpaiDiscoveryEndpoint' field")
		}

		// Simple Field (hpaiDataEndpoint)
		if pushErr := writeBuffer.PushContext("hpaiDataEndpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hpaiDataEndpoint")
		}
		_hpaiDataEndpointErr := writeBuffer.WriteSerializable(m.HpaiDataEndpoint)
		if popErr := writeBuffer.PopContext("hpaiDataEndpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hpaiDataEndpoint")
		}
		if _hpaiDataEndpointErr != nil {
			return errors.Wrap(_hpaiDataEndpointErr, "Error serializing 'hpaiDataEndpoint' field")
		}

		// Simple Field (connectionRequestInformation)
		if pushErr := writeBuffer.PushContext("connectionRequestInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for connectionRequestInformation")
		}
		_connectionRequestInformationErr := writeBuffer.WriteSerializable(m.ConnectionRequestInformation)
		if popErr := writeBuffer.PopContext("connectionRequestInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for connectionRequestInformation")
		}
		if _connectionRequestInformationErr != nil {
			return errors.Wrap(_connectionRequestInformationErr, "Error serializing 'connectionRequestInformation' field")
		}

		if popErr := writeBuffer.PopContext("ConnectionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ConnectionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
