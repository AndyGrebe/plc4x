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

// BACnetConstructedDataBackupAndRestoreState is the data-structure of this message
type BACnetConstructedDataBackupAndRestoreState struct {
	*BACnetConstructedData
	BackupAndRestoreState *BACnetBackupStateTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBackupAndRestoreState is the corresponding interface of BACnetConstructedDataBackupAndRestoreState
type IBACnetConstructedDataBackupAndRestoreState interface {
	IBACnetConstructedData
	// GetBackupAndRestoreState returns BackupAndRestoreState (property field)
	GetBackupAndRestoreState() *BACnetBackupStateTagged
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

func (m *BACnetConstructedDataBackupAndRestoreState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBackupAndRestoreState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACKUP_AND_RESTORE_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBackupAndRestoreState) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBackupAndRestoreState) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBackupAndRestoreState) GetBackupAndRestoreState() *BACnetBackupStateTagged {
	return m.BackupAndRestoreState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBackupAndRestoreState factory function for BACnetConstructedDataBackupAndRestoreState
func NewBACnetConstructedDataBackupAndRestoreState(backupAndRestoreState *BACnetBackupStateTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBackupAndRestoreState {
	_result := &BACnetConstructedDataBackupAndRestoreState{
		BackupAndRestoreState: backupAndRestoreState,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBackupAndRestoreState(structType interface{}) *BACnetConstructedDataBackupAndRestoreState {
	if casted, ok := structType.(BACnetConstructedDataBackupAndRestoreState); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBackupAndRestoreState); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBackupAndRestoreState(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBackupAndRestoreState(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBackupAndRestoreState) GetTypeName() string {
	return "BACnetConstructedDataBackupAndRestoreState"
}

func (m *BACnetConstructedDataBackupAndRestoreState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBackupAndRestoreState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (backupAndRestoreState)
	lengthInBits += m.BackupAndRestoreState.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBackupAndRestoreState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBackupAndRestoreStateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBackupAndRestoreState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBackupAndRestoreState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBackupAndRestoreState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (backupAndRestoreState)
	if pullErr := readBuffer.PullContext("backupAndRestoreState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for backupAndRestoreState")
	}
	_backupAndRestoreState, _backupAndRestoreStateErr := BACnetBackupStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _backupAndRestoreStateErr != nil {
		return nil, errors.Wrap(_backupAndRestoreStateErr, "Error parsing 'backupAndRestoreState' field")
	}
	backupAndRestoreState := CastBACnetBackupStateTagged(_backupAndRestoreState)
	if closeErr := readBuffer.CloseContext("backupAndRestoreState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for backupAndRestoreState")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBackupAndRestoreState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBackupAndRestoreState")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBackupAndRestoreState{
		BackupAndRestoreState: CastBACnetBackupStateTagged(backupAndRestoreState),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBackupAndRestoreState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBackupAndRestoreState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBackupAndRestoreState")
		}

		// Simple Field (backupAndRestoreState)
		if pushErr := writeBuffer.PushContext("backupAndRestoreState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for backupAndRestoreState")
		}
		_backupAndRestoreStateErr := writeBuffer.WriteSerializable(m.BackupAndRestoreState)
		if popErr := writeBuffer.PopContext("backupAndRestoreState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for backupAndRestoreState")
		}
		if _backupAndRestoreStateErr != nil {
			return errors.Wrap(_backupAndRestoreStateErr, "Error serializing 'backupAndRestoreState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBackupAndRestoreState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBackupAndRestoreState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBackupAndRestoreState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
