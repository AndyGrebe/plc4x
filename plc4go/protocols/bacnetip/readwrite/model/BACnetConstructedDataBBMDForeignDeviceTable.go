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

// BACnetConstructedDataBBMDForeignDeviceTable is the data-structure of this message
type BACnetConstructedDataBBMDForeignDeviceTable struct {
	*BACnetConstructedData
	BbmdForeignDeviceTable []*BACnetBDTEntry

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBBMDForeignDeviceTable is the corresponding interface of BACnetConstructedDataBBMDForeignDeviceTable
type IBACnetConstructedDataBBMDForeignDeviceTable interface {
	IBACnetConstructedData
	// GetBbmdForeignDeviceTable returns BbmdForeignDeviceTable (property field)
	GetBbmdForeignDeviceTable() []*BACnetBDTEntry
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

func (m *BACnetConstructedDataBBMDForeignDeviceTable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBBMDForeignDeviceTable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BBMD_FOREIGN_DEVICE_TABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBBMDForeignDeviceTable) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBBMDForeignDeviceTable) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBBMDForeignDeviceTable) GetBbmdForeignDeviceTable() []*BACnetBDTEntry {
	return m.BbmdForeignDeviceTable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBBMDForeignDeviceTable factory function for BACnetConstructedDataBBMDForeignDeviceTable
func NewBACnetConstructedDataBBMDForeignDeviceTable(bbmdForeignDeviceTable []*BACnetBDTEntry, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBBMDForeignDeviceTable {
	_result := &BACnetConstructedDataBBMDForeignDeviceTable{
		BbmdForeignDeviceTable: bbmdForeignDeviceTable,
		BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBBMDForeignDeviceTable(structType interface{}) *BACnetConstructedDataBBMDForeignDeviceTable {
	if casted, ok := structType.(BACnetConstructedDataBBMDForeignDeviceTable); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBBMDForeignDeviceTable); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBBMDForeignDeviceTable(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBBMDForeignDeviceTable(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBBMDForeignDeviceTable) GetTypeName() string {
	return "BACnetConstructedDataBBMDForeignDeviceTable"
}

func (m *BACnetConstructedDataBBMDForeignDeviceTable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBBMDForeignDeviceTable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.BbmdForeignDeviceTable) > 0 {
		for _, element := range m.BbmdForeignDeviceTable {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataBBMDForeignDeviceTable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBBMDForeignDeviceTableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBBMDForeignDeviceTable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBBMDForeignDeviceTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBBMDForeignDeviceTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (bbmdForeignDeviceTable)
	if pullErr := readBuffer.PullContext("bbmdForeignDeviceTable", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bbmdForeignDeviceTable")
	}
	// Terminated array
	bbmdForeignDeviceTable := make([]*BACnetBDTEntry, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetBDTEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'bbmdForeignDeviceTable' field")
			}
			bbmdForeignDeviceTable = append(bbmdForeignDeviceTable, CastBACnetBDTEntry(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("bbmdForeignDeviceTable", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bbmdForeignDeviceTable")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBBMDForeignDeviceTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBBMDForeignDeviceTable")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBBMDForeignDeviceTable{
		BbmdForeignDeviceTable: bbmdForeignDeviceTable,
		BACnetConstructedData:  &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBBMDForeignDeviceTable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBBMDForeignDeviceTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBBMDForeignDeviceTable")
		}

		// Array Field (bbmdForeignDeviceTable)
		if m.BbmdForeignDeviceTable != nil {
			if pushErr := writeBuffer.PushContext("bbmdForeignDeviceTable", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for bbmdForeignDeviceTable")
			}
			for _, _element := range m.BbmdForeignDeviceTable {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'bbmdForeignDeviceTable' field")
				}
			}
			if popErr := writeBuffer.PopContext("bbmdForeignDeviceTable", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for bbmdForeignDeviceTable")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBBMDForeignDeviceTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBBMDForeignDeviceTable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBBMDForeignDeviceTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
