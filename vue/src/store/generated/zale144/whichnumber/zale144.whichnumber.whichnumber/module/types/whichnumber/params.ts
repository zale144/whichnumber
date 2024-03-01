/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "zale144.whichnumber.whichnumber";

/** Params defines the parameters for the module. */
export interface Params {
  /** in seconds */
  commit_timeout: number;
  /** in seconds */
  reveal_timeout: number;
  max_players_per_game: number;
  min_distance_to_win: number;
  min_reward: Coin | undefined;
}

const baseParams: object = {
  commit_timeout: 0,
  reveal_timeout: 0,
  max_players_per_game: 0,
  min_distance_to_win: 0,
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.commit_timeout !== 0) {
      writer.uint32(8).uint64(message.commit_timeout);
    }
    if (message.reveal_timeout !== 0) {
      writer.uint32(16).uint64(message.reveal_timeout);
    }
    if (message.max_players_per_game !== 0) {
      writer.uint32(24).uint64(message.max_players_per_game);
    }
    if (message.min_distance_to_win !== 0) {
      writer.uint32(32).uint64(message.min_distance_to_win);
    }
    if (message.min_reward !== undefined) {
      Coin.encode(message.min_reward, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commit_timeout = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.reveal_timeout = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.max_players_per_game = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.min_distance_to_win = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.min_reward = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.commit_timeout !== undefined && object.commit_timeout !== null) {
      message.commit_timeout = Number(object.commit_timeout);
    } else {
      message.commit_timeout = 0;
    }
    if (object.reveal_timeout !== undefined && object.reveal_timeout !== null) {
      message.reveal_timeout = Number(object.reveal_timeout);
    } else {
      message.reveal_timeout = 0;
    }
    if (
      object.max_players_per_game !== undefined &&
      object.max_players_per_game !== null
    ) {
      message.max_players_per_game = Number(object.max_players_per_game);
    } else {
      message.max_players_per_game = 0;
    }
    if (
      object.min_distance_to_win !== undefined &&
      object.min_distance_to_win !== null
    ) {
      message.min_distance_to_win = Number(object.min_distance_to_win);
    } else {
      message.min_distance_to_win = 0;
    }
    if (object.min_reward !== undefined && object.min_reward !== null) {
      message.min_reward = Coin.fromJSON(object.min_reward);
    } else {
      message.min_reward = undefined;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.commit_timeout !== undefined &&
      (obj.commit_timeout = message.commit_timeout);
    message.reveal_timeout !== undefined &&
      (obj.reveal_timeout = message.reveal_timeout);
    message.max_players_per_game !== undefined &&
      (obj.max_players_per_game = message.max_players_per_game);
    message.min_distance_to_win !== undefined &&
      (obj.min_distance_to_win = message.min_distance_to_win);
    message.min_reward !== undefined &&
      (obj.min_reward = message.min_reward
        ? Coin.toJSON(message.min_reward)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.commit_timeout !== undefined && object.commit_timeout !== null) {
      message.commit_timeout = object.commit_timeout;
    } else {
      message.commit_timeout = 0;
    }
    if (object.reveal_timeout !== undefined && object.reveal_timeout !== null) {
      message.reveal_timeout = object.reveal_timeout;
    } else {
      message.reveal_timeout = 0;
    }
    if (
      object.max_players_per_game !== undefined &&
      object.max_players_per_game !== null
    ) {
      message.max_players_per_game = object.max_players_per_game;
    } else {
      message.max_players_per_game = 0;
    }
    if (
      object.min_distance_to_win !== undefined &&
      object.min_distance_to_win !== null
    ) {
      message.min_distance_to_win = object.min_distance_to_win;
    } else {
      message.min_distance_to_win = 0;
    }
    if (object.min_reward !== undefined && object.min_reward !== null) {
      message.min_reward = Coin.fromPartial(object.min_reward);
    } else {
      message.min_reward = undefined;
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
