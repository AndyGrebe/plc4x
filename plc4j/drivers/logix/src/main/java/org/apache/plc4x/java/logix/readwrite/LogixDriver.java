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
package org.apache.plc4x.java.logix.readwrite;

import io.netty.buffer.ByteBuf;
import org.apache.plc4x.java.logix.readwrite.configuration.LogixConfiguration;
import org.apache.plc4x.java.logix.readwrite.field.LogixField;
import org.apache.plc4x.java.logix.readwrite.field.LogixFieldHandler;
import org.apache.plc4x.java.logix.readwrite.protocol.LogixProtocolLogic;
import org.apache.plc4x.java.spi.values.IEC61131ValueHandler;
import org.apache.plc4x.java.api.value.PlcValueHandler;
import org.apache.plc4x.java.spi.configuration.Configuration;
import org.apache.plc4x.java.spi.connection.GeneratedDriverBase;
import org.apache.plc4x.java.spi.connection.PlcFieldHandler;
import org.apache.plc4x.java.spi.connection.ProtocolStackConfigurer;
import org.apache.plc4x.java.spi.connection.SingleProtocolStackConfigurer;

import java.util.function.Consumer;
import java.util.function.ToIntFunction;

public class LogixDriver extends GeneratedDriverBase<EipPacket> {
    public static final int PORT = 44818;
    @Override
    public String getProtocolCode() {
        return "logix";
    }

    @Override
    public String getProtocolName() {
        return "Allen Bradley Logix";
    }

    @Override
    protected Class<? extends Configuration> getConfigurationType() {
        return LogixConfiguration.class;
    }

    @Override
    protected PlcFieldHandler getFieldHandler() {
        return new LogixFieldHandler();
    }

    @Override
    protected PlcValueHandler getValueHandler() {
        return new IEC61131ValueHandler();
    }

    @Override
    protected boolean awaitDisconnectComplete() {
        return true;
    }

    @Override
    protected String getDefaultTransport() {
        return "tcp";
    }

    @Override
    protected boolean canRead() {
        return true;
    }

    @Override
    protected boolean canWrite() {
        return true;
    }

    @Override
    protected ProtocolStackConfigurer<EipPacket> getStackConfigurer() {
        return SingleProtocolStackConfigurer.builder(EipPacket.class, EipPacket::staticParse)
            .withProtocol(LogixProtocolLogic.class)
            .withPacketSizeEstimator(ByteLengthEstimator.class)
            .littleEndian()
            .build();
    }

    /** Estimate the Length of a Packet */
    public static class ByteLengthEstimator implements ToIntFunction<ByteBuf> {
        @Override
        public int applyAsInt(ByteBuf byteBuf) {
            if (byteBuf.readableBytes() >= 4) {
                //Second byte for the size and then add the header size 24
                int size = byteBuf.getUnsignedShort(byteBuf.readerIndex()+1)+24;
                return size;
            }
            return -1;
        }
    }

     /**Consumes all Bytes till another Magic Byte is found */
    public static class CorruptPackageCleaner implements Consumer<ByteBuf> {
        @Override
        public void accept(ByteBuf byteBuf) {
            while (byteBuf.getUnsignedByte(0) != 0x00) {
                // Just consume the bytes till the next possible start position.
                byteBuf.readByte();
            }
        }
    }

    @Override
    public LogixField prepareField(String query){
        return LogixField.of(query);
    }

}