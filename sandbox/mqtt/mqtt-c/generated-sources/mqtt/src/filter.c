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

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include "filter.h"

// Code generated by code-generation. DO NOT EDIT.


// Parse function.
plc4c_return_code plc4c_mqtt_read_write_filter_parse(plc4c_spi_read_buffer* readBuffer, plc4c_mqtt_read_write_filter** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_mqtt_read_write_filter));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Simple Field (filter)
  plc4c_mqtt_read_write_mqt_t__string* filter;
  _res = plc4c_mqtt_read_write_mqt_t__string_parse(readBuffer, (void*) &filter);
  if(_res != OK) {
    return _res;
  }
  (*_message)->filter = filter;

  // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
  {
    uint8_t _reserved = 0;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 2, (uint8_t*) &_reserved);
    if(_res != OK) {
      return _res;
    }
    if(_reserved != 0x0) {
      printf("Expected constant value '%d' but got '%d' for reserved field.", 0x0, _reserved);
    }
  }

  // Simple Field (retainHandling)
  plc4c_mqtt_read_write_mqt_t__retain_handling* retainHandling;
  _res = plc4c_mqtt_read_write_mqt_t__retain_handling_parse(readBuffer, (void*) &retainHandling);
  if(_res != OK) {
    return _res;
  }
  (*_message)->retain_handling = *retainHandling;

  // Simple Field (retain)
  bool retain = false;
  _res = plc4c_spi_read_bit(readBuffer, (bool*) &retain);
  if(_res != OK) {
    return _res;
  }
  (*_message)->retain = retain;

  // Simple Field (noLocal)
  bool noLocal = false;
  _res = plc4c_spi_read_bit(readBuffer, (bool*) &noLocal);
  if(_res != OK) {
    return _res;
  }
  (*_message)->no_local = noLocal;

  // Simple Field (maxQos)
  plc4c_mqtt_read_write_mqt_t__qos* maxQos;
  _res = plc4c_mqtt_read_write_mqt_t__qos_parse(readBuffer, (void*) &maxQos);
  if(_res != OK) {
    return _res;
  }
  (*_message)->max_qos = *maxQos;

  return OK;
}

plc4c_return_code plc4c_mqtt_read_write_filter_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_mqtt_read_write_filter* _message) {
  plc4c_return_code _res = OK;

  // Simple Field (filter)
  _res = plc4c_mqtt_read_write_mqt_t__string_serialize(writeBuffer, _message->filter);
  if(_res != OK) {
    return _res;
  }

  // Reserved Field
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 2, 0x0);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (retainHandling)
  _res = plc4c_mqtt_read_write_mqt_t__retain_handling_serialize(writeBuffer, &_message->retain_handling);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (retain)
  _res = plc4c_spi_write_bit(writeBuffer, _message->retain);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (noLocal)
  _res = plc4c_spi_write_bit(writeBuffer, _message->no_local);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (maxQos)
  _res = plc4c_mqtt_read_write_mqt_t__qos_serialize(writeBuffer, &_message->max_qos);
  if(_res != OK) {
    return _res;
  }

  return OK;
}

uint16_t plc4c_mqtt_read_write_filter_length_in_bytes(plc4c_mqtt_read_write_filter* _message) {
  return plc4c_mqtt_read_write_filter_length_in_bits(_message) / 8;
}

uint16_t plc4c_mqtt_read_write_filter_length_in_bits(plc4c_mqtt_read_write_filter* _message) {
  uint16_t lengthInBits = 0;

  // Simple field (filter)
  lengthInBits += plc4c_mqtt_read_write_mqt_t__string_length_in_bits(_message->filter);

  // Reserved Field (reserved)
  lengthInBits += 2;

  // Simple field (retainHandling)
  lengthInBits += plc4c_mqtt_read_write_mqt_t__retain_handling_length_in_bits(&_message->retain_handling);

  // Simple field (retain)
  lengthInBits += 1;

  // Simple field (noLocal)
  lengthInBits += 1;

  // Simple field (maxQos)
  lengthInBits += plc4c_mqtt_read_write_mqt_t__qos_length_in_bits(&_message->max_qos);

  return lengthInBits;
}

