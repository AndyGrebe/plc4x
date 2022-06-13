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

// BACnetConstructedDataTimeSynchronizationRecipients is the data-structure of this message
type BACnetConstructedDataTimeSynchronizationRecipients struct {
	*BACnetConstructedData
	TimeSynchronizationRecipients []*BACnetRecipient

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataTimeSynchronizationRecipients is the corresponding interface of BACnetConstructedDataTimeSynchronizationRecipients
type IBACnetConstructedDataTimeSynchronizationRecipients interface {
	IBACnetConstructedData
	// GetTimeSynchronizationRecipients returns TimeSynchronizationRecipients (property field)
	GetTimeSynchronizationRecipients() []*BACnetRecipient
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

func (m *BACnetConstructedDataTimeSynchronizationRecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataTimeSynchronizationRecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_SYNCHRONIZATION_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTimeSynchronizationRecipients) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTimeSynchronizationRecipients) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTimeSynchronizationRecipients) GetTimeSynchronizationRecipients() []*BACnetRecipient {
	return m.TimeSynchronizationRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeSynchronizationRecipients factory function for BACnetConstructedDataTimeSynchronizationRecipients
func NewBACnetConstructedDataTimeSynchronizationRecipients(timeSynchronizationRecipients []*BACnetRecipient, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataTimeSynchronizationRecipients {
	_result := &BACnetConstructedDataTimeSynchronizationRecipients{
		TimeSynchronizationRecipients: timeSynchronizationRecipients,
		BACnetConstructedData:         NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTimeSynchronizationRecipients(structType interface{}) *BACnetConstructedDataTimeSynchronizationRecipients {
	if casted, ok := structType.(BACnetConstructedDataTimeSynchronizationRecipients); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeSynchronizationRecipients); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimeSynchronizationRecipients(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimeSynchronizationRecipients(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTimeSynchronizationRecipients) GetTypeName() string {
	return "BACnetConstructedDataTimeSynchronizationRecipients"
}

func (m *BACnetConstructedDataTimeSynchronizationRecipients) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTimeSynchronizationRecipients) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.TimeSynchronizationRecipients) > 0 {
		for _, element := range m.TimeSynchronizationRecipients {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataTimeSynchronizationRecipients) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimeSynchronizationRecipientsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataTimeSynchronizationRecipients, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeSynchronizationRecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeSynchronizationRecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (timeSynchronizationRecipients)
	if pullErr := readBuffer.PullContext("timeSynchronizationRecipients", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeSynchronizationRecipients")
	}
	// Terminated array
	timeSynchronizationRecipients := make([]*BACnetRecipient, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetRecipientParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'timeSynchronizationRecipients' field")
			}
			timeSynchronizationRecipients = append(timeSynchronizationRecipients, CastBACnetRecipient(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("timeSynchronizationRecipients", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeSynchronizationRecipients")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeSynchronizationRecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeSynchronizationRecipients")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTimeSynchronizationRecipients{
		TimeSynchronizationRecipients: timeSynchronizationRecipients,
		BACnetConstructedData:         &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTimeSynchronizationRecipients) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeSynchronizationRecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeSynchronizationRecipients")
		}

		// Array Field (timeSynchronizationRecipients)
		if m.TimeSynchronizationRecipients != nil {
			if pushErr := writeBuffer.PushContext("timeSynchronizationRecipients", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for timeSynchronizationRecipients")
			}
			for _, _element := range m.TimeSynchronizationRecipients {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'timeSynchronizationRecipients' field")
				}
			}
			if popErr := writeBuffer.PopContext("timeSynchronizationRecipients", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for timeSynchronizationRecipients")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeSynchronizationRecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeSynchronizationRecipients")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTimeSynchronizationRecipients) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
