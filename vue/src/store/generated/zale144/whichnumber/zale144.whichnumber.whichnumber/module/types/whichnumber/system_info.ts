/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "zale144.whichnumber.whichnumber";

export interface SystemInfo {
  nextId: number;
  fifoHeadId: number;
  fifoTailId: number;
}

const baseSystemInfo: object = { nextId: 0, fifoHeadId: 0, fifoTailId: 0 };

export const SystemInfo = {
  encode(message: SystemInfo, writer: Writer = Writer.create()): Writer {
    if (message.nextId !== 0) {
      writer.uint32(8).int64(message.nextId);
    }
    if (message.fifoHeadId !== 0) {
      writer.uint32(16).int64(message.fifoHeadId);
    }
    if (message.fifoTailId !== 0) {
      writer.uint32(24).int64(message.fifoTailId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SystemInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSystemInfo } as SystemInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nextId = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.fifoHeadId = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.fifoTailId = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SystemInfo {
    const message = { ...baseSystemInfo } as SystemInfo;
    if (object.nextId !== undefined && object.nextId !== null) {
      message.nextId = Number(object.nextId);
    } else {
      message.nextId = 0;
    }
    if (object.fifoHeadId !== undefined && object.fifoHeadId !== null) {
      message.fifoHeadId = Number(object.fifoHeadId);
    } else {
      message.fifoHeadId = 0;
    }
    if (object.fifoTailId !== undefined && object.fifoTailId !== null) {
      message.fifoTailId = Number(object.fifoTailId);
    } else {
      message.fifoTailId = 0;
    }
    return message;
  },

  toJSON(message: SystemInfo): unknown {
    const obj: any = {};
    message.nextId !== undefined && (obj.nextId = message.nextId);
    message.fifoHeadId !== undefined && (obj.fifoHeadId = message.fifoHeadId);
    message.fifoTailId !== undefined && (obj.fifoTailId = message.fifoTailId);
    return obj;
  },

  fromPartial(object: DeepPartial<SystemInfo>): SystemInfo {
    const message = { ...baseSystemInfo } as SystemInfo;
    if (object.nextId !== undefined && object.nextId !== null) {
      message.nextId = object.nextId;
    } else {
      message.nextId = 0;
    }
    if (object.fifoHeadId !== undefined && object.fifoHeadId !== null) {
      message.fifoHeadId = object.fifoHeadId;
    } else {
      message.fifoHeadId = 0;
    }
    if (object.fifoTailId !== undefined && object.fifoTailId !== null) {
      message.fifoTailId = object.fifoTailId;
    } else {
      message.fifoTailId = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
