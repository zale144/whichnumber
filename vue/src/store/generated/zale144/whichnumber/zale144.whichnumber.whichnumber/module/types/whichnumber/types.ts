/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "zale144.whichnumber.whichnumber";

export interface Game {
  id: number;
  /** Address of the player who created the game */
  creator: string;
  /** The secret number that the players are guessing */
  secret_number: number;
  /** The guesses submitted by the players */
  player_commits: NumberCommit[];
  /** The reveals submitted by the players */
  player_reveals: NumberReveal[];
  /** The reward for the winner */
  reward: Coin | undefined;
  entry_fee: Coin | undefined;
  commit_timeout: Date | undefined;
  reveal_timeout: Date | undefined;
  beforeId: number;
  afterId: number;
}

export interface NumberCommit {
  /** hex encoded sha256 of "salt:number" */
  commit: string;
  /** Address of the player who submitted the guess */
  player_address: string;
  created_at: Date | undefined;
}

export interface NumberReveal {
  player_address: string;
  number: number;
  /** hex encoded 32 bytes salt */
  salt: string;
  created_at: Date | undefined;
}

const baseGame: object = {
  id: 0,
  creator: "",
  secret_number: 0,
  beforeId: 0,
  afterId: 0,
};

export const Game = {
  encode(message: Game, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.secret_number !== 0) {
      writer.uint32(24).int64(message.secret_number);
    }
    for (const v of message.player_commits) {
      NumberCommit.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.player_reveals) {
      NumberReveal.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.reward !== undefined) {
      Coin.encode(message.reward, writer.uint32(50).fork()).ldelim();
    }
    if (message.entry_fee !== undefined) {
      Coin.encode(message.entry_fee, writer.uint32(58).fork()).ldelim();
    }
    if (message.commit_timeout !== undefined) {
      Timestamp.encode(
        toTimestamp(message.commit_timeout),
        writer.uint32(66).fork()
      ).ldelim();
    }
    if (message.reveal_timeout !== undefined) {
      Timestamp.encode(
        toTimestamp(message.reveal_timeout),
        writer.uint32(74).fork()
      ).ldelim();
    }
    if (message.beforeId !== 0) {
      writer.uint32(80).int64(message.beforeId);
    }
    if (message.afterId !== 0) {
      writer.uint32(88).int64(message.afterId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Game {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGame } as Game;
    message.player_commits = [];
    message.player_reveals = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.secret_number = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.player_commits.push(
            NumberCommit.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.player_reveals.push(
            NumberReveal.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.reward = Coin.decode(reader, reader.uint32());
          break;
        case 7:
          message.entry_fee = Coin.decode(reader, reader.uint32());
          break;
        case 8:
          message.commit_timeout = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 9:
          message.reveal_timeout = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.beforeId = longToNumber(reader.int64() as Long);
          break;
        case 11:
          message.afterId = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Game {
    const message = { ...baseGame } as Game;
    message.player_commits = [];
    message.player_reveals = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.secret_number !== undefined && object.secret_number !== null) {
      message.secret_number = Number(object.secret_number);
    } else {
      message.secret_number = 0;
    }
    if (object.player_commits !== undefined && object.player_commits !== null) {
      for (const e of object.player_commits) {
        message.player_commits.push(NumberCommit.fromJSON(e));
      }
    }
    if (object.player_reveals !== undefined && object.player_reveals !== null) {
      for (const e of object.player_reveals) {
        message.player_reveals.push(NumberReveal.fromJSON(e));
      }
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = Coin.fromJSON(object.reward);
    } else {
      message.reward = undefined;
    }
    if (object.entry_fee !== undefined && object.entry_fee !== null) {
      message.entry_fee = Coin.fromJSON(object.entry_fee);
    } else {
      message.entry_fee = undefined;
    }
    if (object.commit_timeout !== undefined && object.commit_timeout !== null) {
      message.commit_timeout = fromJsonTimestamp(object.commit_timeout);
    } else {
      message.commit_timeout = undefined;
    }
    if (object.reveal_timeout !== undefined && object.reveal_timeout !== null) {
      message.reveal_timeout = fromJsonTimestamp(object.reveal_timeout);
    } else {
      message.reveal_timeout = undefined;
    }
    if (object.beforeId !== undefined && object.beforeId !== null) {
      message.beforeId = Number(object.beforeId);
    } else {
      message.beforeId = 0;
    }
    if (object.afterId !== undefined && object.afterId !== null) {
      message.afterId = Number(object.afterId);
    } else {
      message.afterId = 0;
    }
    return message;
  },

  toJSON(message: Game): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creator !== undefined && (obj.creator = message.creator);
    message.secret_number !== undefined &&
      (obj.secret_number = message.secret_number);
    if (message.player_commits) {
      obj.player_commits = message.player_commits.map((e) =>
        e ? NumberCommit.toJSON(e) : undefined
      );
    } else {
      obj.player_commits = [];
    }
    if (message.player_reveals) {
      obj.player_reveals = message.player_reveals.map((e) =>
        e ? NumberReveal.toJSON(e) : undefined
      );
    } else {
      obj.player_reveals = [];
    }
    message.reward !== undefined &&
      (obj.reward = message.reward ? Coin.toJSON(message.reward) : undefined);
    message.entry_fee !== undefined &&
      (obj.entry_fee = message.entry_fee
        ? Coin.toJSON(message.entry_fee)
        : undefined);
    message.commit_timeout !== undefined &&
      (obj.commit_timeout =
        message.commit_timeout !== undefined
          ? message.commit_timeout.toISOString()
          : null);
    message.reveal_timeout !== undefined &&
      (obj.reveal_timeout =
        message.reveal_timeout !== undefined
          ? message.reveal_timeout.toISOString()
          : null);
    message.beforeId !== undefined && (obj.beforeId = message.beforeId);
    message.afterId !== undefined && (obj.afterId = message.afterId);
    return obj;
  },

  fromPartial(object: DeepPartial<Game>): Game {
    const message = { ...baseGame } as Game;
    message.player_commits = [];
    message.player_reveals = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.secret_number !== undefined && object.secret_number !== null) {
      message.secret_number = object.secret_number;
    } else {
      message.secret_number = 0;
    }
    if (object.player_commits !== undefined && object.player_commits !== null) {
      for (const e of object.player_commits) {
        message.player_commits.push(NumberCommit.fromPartial(e));
      }
    }
    if (object.player_reveals !== undefined && object.player_reveals !== null) {
      for (const e of object.player_reveals) {
        message.player_reveals.push(NumberReveal.fromPartial(e));
      }
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = Coin.fromPartial(object.reward);
    } else {
      message.reward = undefined;
    }
    if (object.entry_fee !== undefined && object.entry_fee !== null) {
      message.entry_fee = Coin.fromPartial(object.entry_fee);
    } else {
      message.entry_fee = undefined;
    }
    if (object.commit_timeout !== undefined && object.commit_timeout !== null) {
      message.commit_timeout = object.commit_timeout;
    } else {
      message.commit_timeout = undefined;
    }
    if (object.reveal_timeout !== undefined && object.reveal_timeout !== null) {
      message.reveal_timeout = object.reveal_timeout;
    } else {
      message.reveal_timeout = undefined;
    }
    if (object.beforeId !== undefined && object.beforeId !== null) {
      message.beforeId = object.beforeId;
    } else {
      message.beforeId = 0;
    }
    if (object.afterId !== undefined && object.afterId !== null) {
      message.afterId = object.afterId;
    } else {
      message.afterId = 0;
    }
    return message;
  },
};

const baseNumberCommit: object = { commit: "", player_address: "" };

export const NumberCommit = {
  encode(message: NumberCommit, writer: Writer = Writer.create()): Writer {
    if (message.commit !== "") {
      writer.uint32(10).string(message.commit);
    }
    if (message.player_address !== "") {
      writer.uint32(18).string(message.player_address);
    }
    if (message.created_at !== undefined) {
      Timestamp.encode(
        toTimestamp(message.created_at),
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NumberCommit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNumberCommit } as NumberCommit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commit = reader.string();
          break;
        case 2:
          message.player_address = reader.string();
          break;
        case 3:
          message.created_at = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NumberCommit {
    const message = { ...baseNumberCommit } as NumberCommit;
    if (object.commit !== undefined && object.commit !== null) {
      message.commit = String(object.commit);
    } else {
      message.commit = "";
    }
    if (object.player_address !== undefined && object.player_address !== null) {
      message.player_address = String(object.player_address);
    } else {
      message.player_address = "";
    }
    if (object.created_at !== undefined && object.created_at !== null) {
      message.created_at = fromJsonTimestamp(object.created_at);
    } else {
      message.created_at = undefined;
    }
    return message;
  },

  toJSON(message: NumberCommit): unknown {
    const obj: any = {};
    message.commit !== undefined && (obj.commit = message.commit);
    message.player_address !== undefined &&
      (obj.player_address = message.player_address);
    message.created_at !== undefined &&
      (obj.created_at =
        message.created_at !== undefined
          ? message.created_at.toISOString()
          : null);
    return obj;
  },

  fromPartial(object: DeepPartial<NumberCommit>): NumberCommit {
    const message = { ...baseNumberCommit } as NumberCommit;
    if (object.commit !== undefined && object.commit !== null) {
      message.commit = object.commit;
    } else {
      message.commit = "";
    }
    if (object.player_address !== undefined && object.player_address !== null) {
      message.player_address = object.player_address;
    } else {
      message.player_address = "";
    }
    if (object.created_at !== undefined && object.created_at !== null) {
      message.created_at = object.created_at;
    } else {
      message.created_at = undefined;
    }
    return message;
  },
};

const baseNumberReveal: object = { player_address: "", number: 0, salt: "" };

export const NumberReveal = {
  encode(message: NumberReveal, writer: Writer = Writer.create()): Writer {
    if (message.player_address !== "") {
      writer.uint32(10).string(message.player_address);
    }
    if (message.number !== 0) {
      writer.uint32(16).int64(message.number);
    }
    if (message.salt !== "") {
      writer.uint32(26).string(message.salt);
    }
    if (message.created_at !== undefined) {
      Timestamp.encode(
        toTimestamp(message.created_at),
        writer.uint32(34).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NumberReveal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNumberReveal } as NumberReveal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.player_address = reader.string();
          break;
        case 2:
          message.number = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.salt = reader.string();
          break;
        case 4:
          message.created_at = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NumberReveal {
    const message = { ...baseNumberReveal } as NumberReveal;
    if (object.player_address !== undefined && object.player_address !== null) {
      message.player_address = String(object.player_address);
    } else {
      message.player_address = "";
    }
    if (object.number !== undefined && object.number !== null) {
      message.number = Number(object.number);
    } else {
      message.number = 0;
    }
    if (object.salt !== undefined && object.salt !== null) {
      message.salt = String(object.salt);
    } else {
      message.salt = "";
    }
    if (object.created_at !== undefined && object.created_at !== null) {
      message.created_at = fromJsonTimestamp(object.created_at);
    } else {
      message.created_at = undefined;
    }
    return message;
  },

  toJSON(message: NumberReveal): unknown {
    const obj: any = {};
    message.player_address !== undefined &&
      (obj.player_address = message.player_address);
    message.number !== undefined && (obj.number = message.number);
    message.salt !== undefined && (obj.salt = message.salt);
    message.created_at !== undefined &&
      (obj.created_at =
        message.created_at !== undefined
          ? message.created_at.toISOString()
          : null);
    return obj;
  },

  fromPartial(object: DeepPartial<NumberReveal>): NumberReveal {
    const message = { ...baseNumberReveal } as NumberReveal;
    if (object.player_address !== undefined && object.player_address !== null) {
      message.player_address = object.player_address;
    } else {
      message.player_address = "";
    }
    if (object.number !== undefined && object.number !== null) {
      message.number = object.number;
    } else {
      message.number = 0;
    }
    if (object.salt !== undefined && object.salt !== null) {
      message.salt = object.salt;
    } else {
      message.salt = "";
    }
    if (object.created_at !== undefined && object.created_at !== null) {
      message.created_at = object.created_at;
    } else {
      message.created_at = undefined;
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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
