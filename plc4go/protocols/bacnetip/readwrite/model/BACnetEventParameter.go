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

// BACnetEventParameter is the data-structure of this message
type BACnetEventParameter struct {
	PeekedTagHeader *BACnetTagHeader
	Child           IBACnetEventParameterChild
}

// IBACnetEventParameter is the corresponding interface of BACnetEventParameter
type IBACnetEventParameter interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() *BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetEventParameterParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetEventParameter, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetEventParameterChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetEventParameter, peekedTagHeader *BACnetTagHeader)
	GetParent() *BACnetEventParameter

	GetTypeName() string
	IBACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventParameter) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetEventParameter) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameter factory function for BACnetEventParameter
func NewBACnetEventParameter(peekedTagHeader *BACnetTagHeader) *BACnetEventParameter {
	return &BACnetEventParameter{PeekedTagHeader: peekedTagHeader}
}

func CastBACnetEventParameter(structType interface{}) *BACnetEventParameter {
	if casted, ok := structType.(BACnetEventParameter); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventParameter); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetEventParameterChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetEventParameter) GetTypeName() string {
	return "BACnetEventParameter"
}

func (m *BACnetEventParameter) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventParameter) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetEventParameter) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetEventParameter) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterParse(readBuffer utils.ReadBuffer) (*BACnetEventParameter, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetEventParameterChild interface {
		InitializeParent(*BACnetEventParameter, *BACnetTagHeader)
		GetParent() *BACnetEventParameter
	}
	var _child BACnetEventParameterChild
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetEventParameterChangeOfBitstring
		_child, typeSwitchError = BACnetEventParameterChangeOfBitstringParse(readBuffer)
	case peekedTagNumber == uint8(1): // BACnetEventParameterChangeOfState
		_child, typeSwitchError = BACnetEventParameterChangeOfStateParse(readBuffer)
	case peekedTagNumber == uint8(2): // BACnetEventParameterChangeOfValue
		_child, typeSwitchError = BACnetEventParameterChangeOfValueParse(readBuffer)
	case peekedTagNumber == uint8(3): // BACnetEventParameterCommandFailure
		_child, typeSwitchError = BACnetEventParameterCommandFailureParse(readBuffer)
	case peekedTagNumber == uint8(4): // BACnetEventParameterFloatingLimit
		_child, typeSwitchError = BACnetEventParameterFloatingLimitParse(readBuffer)
	case peekedTagNumber == uint8(5): // BACnetEventParameterOutOfRange
		_child, typeSwitchError = BACnetEventParameterOutOfRangeParse(readBuffer)
	case peekedTagNumber == uint8(8): // BACnetEventParameterChangeOfLifeSavety
		_child, typeSwitchError = BACnetEventParameterChangeOfLifeSavetyParse(readBuffer)
	case peekedTagNumber == uint8(9): // BACnetEventParameterExtended
		_child, typeSwitchError = BACnetEventParameterExtendedParse(readBuffer)
	case peekedTagNumber == uint8(10): // BACnetEventParameterBufferReady
		_child, typeSwitchError = BACnetEventParameterBufferReadyParse(readBuffer)
	case peekedTagNumber == uint8(11): // BACnetEventParameterUnsignedRange
		_child, typeSwitchError = BACnetEventParameterUnsignedRangeParse(readBuffer)
	case peekedTagNumber == uint8(13): // BACnetEventParameterAccessEvent
		_child, typeSwitchError = BACnetEventParameterAccessEventParse(readBuffer)
	case peekedTagNumber == uint8(14): // BACnetEventParameterDoubleOutOfRange
		_child, typeSwitchError = BACnetEventParameterDoubleOutOfRangeParse(readBuffer)
	case peekedTagNumber == uint8(15): // BACnetEventParameterSignedOutOfRange
		_child, typeSwitchError = BACnetEventParameterSignedOutOfRangeParse(readBuffer)
	case peekedTagNumber == uint8(16): // BACnetEventParameterUnsignedOutOfRange
		_child, typeSwitchError = BACnetEventParameterUnsignedOutOfRangeParse(readBuffer)
	case peekedTagNumber == uint8(17): // BACnetEventParameterChangeOfCharacterString
		_child, typeSwitchError = BACnetEventParameterChangeOfCharacterStringParse(readBuffer)
	case peekedTagNumber == uint8(18): // BACnetEventParameterChangeOfStatusFlags
		_child, typeSwitchError = BACnetEventParameterChangeOfStatusFlagsParse(readBuffer)
	case peekedTagNumber == uint8(20): // BACnetEventParameterNone
		_child, typeSwitchError = BACnetEventParameterNoneParse(readBuffer)
	case peekedTagNumber == uint8(21): // BACnetEventParameterChangeOfDiscreteValue
		_child, typeSwitchError = BACnetEventParameterChangeOfDiscreteValueParse(readBuffer)
	case peekedTagNumber == uint8(22): // BACnetEventParameterChangeOfTimer
		_child, typeSwitchError = BACnetEventParameterChangeOfTimerParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameter")
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), peekedTagHeader)
	return _child.GetParent(), nil
}

func (m *BACnetEventParameter) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetEventParameter) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetEventParameter, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventParameter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameter")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameter")
	}
	return nil
}

func (m *BACnetEventParameter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
