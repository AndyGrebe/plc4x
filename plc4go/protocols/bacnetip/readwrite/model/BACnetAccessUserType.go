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

// BACnetAccessUserType is an enum
type BACnetAccessUserType uint16

type IBACnetAccessUserType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetAccessUserType_ASSET                    BACnetAccessUserType = 0
	BACnetAccessUserType_GROUP                    BACnetAccessUserType = 1
	BACnetAccessUserType_PERSON                   BACnetAccessUserType = 2
	BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE BACnetAccessUserType = 0xFFFF
)

var BACnetAccessUserTypeValues []BACnetAccessUserType

func init() {
	_ = errors.New
	BACnetAccessUserTypeValues = []BACnetAccessUserType{
		BACnetAccessUserType_ASSET,
		BACnetAccessUserType_GROUP,
		BACnetAccessUserType_PERSON,
		BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAccessUserTypeByValue(value uint16) BACnetAccessUserType {
	switch value {
	case 0:
		return BACnetAccessUserType_ASSET
	case 0xFFFF:
		return BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetAccessUserType_GROUP
	case 2:
		return BACnetAccessUserType_PERSON
	}
	return 0
}

func BACnetAccessUserTypeByName(value string) BACnetAccessUserType {
	switch value {
	case "ASSET":
		return BACnetAccessUserType_ASSET
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE
	case "GROUP":
		return BACnetAccessUserType_GROUP
	case "PERSON":
		return BACnetAccessUserType_PERSON
	}
	return 0
}

func BACnetAccessUserTypeKnows(value uint16) bool {
	for _, typeValue := range BACnetAccessUserTypeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessUserType(structType interface{}) BACnetAccessUserType {
	castFunc := func(typ interface{}) BACnetAccessUserType {
		if sBACnetAccessUserType, ok := typ.(BACnetAccessUserType); ok {
			return sBACnetAccessUserType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessUserType) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetAccessUserType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAccessUserTypeParse(readBuffer utils.ReadBuffer) (BACnetAccessUserType, error) {
	val, err := readBuffer.ReadUint16("BACnetAccessUserType", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetAccessUserTypeByValue(val), nil
}

func (e BACnetAccessUserType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetAccessUserType", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetAccessUserType) name() string {
	switch e {
	case BACnetAccessUserType_ASSET:
		return "ASSET"
	case BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAccessUserType_GROUP:
		return "GROUP"
	case BACnetAccessUserType_PERSON:
		return "PERSON"
	}
	return ""
}

func (e BACnetAccessUserType) String() string {
	return e.name()
}
