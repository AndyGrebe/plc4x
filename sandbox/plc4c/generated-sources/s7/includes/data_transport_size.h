/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#ifndef PLC4C_S7_READ_WRITE_DATA_TRANSPORT_SIZE_H_
#define PLC4C_S7_READ_WRITE_DATA_TRANSPORT_SIZE_H_

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

enum plc4c_s7_read_write_data_transport_size {
  plc4c_s7_read_write_data_transport_size_NULL = 0x00,
  plc4c_s7_read_write_data_transport_size_BIT = 0x03,
  plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD = 0x04,
  plc4c_s7_read_write_data_transport_size_INTEGER = 0x05,
  plc4c_s7_read_write_data_transport_size_DINTEGER = 0x06,
  plc4c_s7_read_write_data_transport_size_REAL = 0x07,
  plc4c_s7_read_write_data_transport_size_OCTET_STRING = 0x09
};
typedef enum plc4c_s7_read_write_data_transport_size plc4c_s7_read_write_data_transport_size;

// Get an empty NULL-struct
plc4c_s7_read_write_data_transport_size plc4c_s7_read_write_data_transport_size_null();


bool plc4c_s7_read_write_data_transport_size_get_size_in_bits(plc4c_s7_read_write_data_transport_size value);

#ifdef __cplusplus
}
#endif

#endif  // PLC4C_S7_READ_WRITE_DATA_TRANSPORT_SIZE_H_