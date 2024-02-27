/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "zale144.whichnumber.whichnumber";

export interface EventNewGame {
  creator: string;
  game_id: string;
  entry_fee: string;
  max_players: number;
  reward: string;
  commit_timeout: string;
}

export interface EventRevealTimeout {
  game_id: string;
  reveal_timeout: string;
  number_of_players: number;
}

export interface EventGameEnd {
  game_id: string;
  winners: Winner[];
}

export interface Winner {
  player: string;
  proximity: number;
  reward: string;
}

const baseEventNewGame: object = {
  creator: "",
  game_id: "",
  entry_fee: "",
  max_players: 0,
  reward: "",
  commit_timeout: "",
};

export const EventNewGame = {
  encode(message: EventNewGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.game_id !== "") {
      writer.uint32(18).string(message.game_id);
    }
    if (message.entry_fee !== "") {
      writer.uint32(26).string(message.entry_fee);
    }
    if (message.max_players !== 0) {
      writer.uint32(32).uint64(message.max_players);
    }
    if (message.reward !== "") {
      writer.uint32(42).string(message.reward);
    }
    if (message.commit_timeout !== "") {
      writer.uint32(50).string(message.commit_timeout);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventNewGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventNewGame } as EventNewGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.game_id = reader.string();
          break;
        case 3:
          message.entry_fee = reader.string();
          break;
        case 4:
          message.max_players = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.reward = reader.string();
          break;
        case 6:
          message.commit_timeout = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventNewGame {
    const message = { ...baseEventNewGame } as EventNewGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = String(object.game_id);
    } else {
      message.game_id = "";
    }
    if (object.entry_fee !== undefined && object.entry_fee !== null) {
      message.entry_fee = String(object.entry_fee);
    } else {
      message.entry_fee = "";
    }
    if (object.max_players !== undefined && object.max_players !== null) {
      message.max_players = Number(object.max_players);
    } else {
      message.max_players = 0;
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = String(object.reward);
    } else {
      message.reward = "";
    }
    if (object.commit_timeout !== undefined && object.commit_timeout !== null) {
      message.commit_timeout = String(object.commit_timeout);
    } else {
      message.commit_timeout = "";
    }
    return message;
  },

  toJSON(message: EventNewGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.entry_fee !== undefined && (obj.entry_fee = message.entry_fee);
    message.max_players !== undefined &&
      (obj.max_players = message.max_players);
    message.reward !== undefined && (obj.reward = message.reward);
    message.commit_timeout !== undefined &&
      (obj.commit_timeout = message.commit_timeout);
    return obj;
  },

  fromPartial(object: DeepPartial<EventNewGame>): EventNewGame {
    const message = { ...baseEventNewGame } as EventNewGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = "";
    }
    if (object.entry_fee !== undefined && object.entry_fee !== null) {
      message.entry_fee = object.entry_fee;
    } else {
      message.entry_fee = "";
    }
    if (object.max_players !== undefined && object.max_players !== null) {
      message.max_players = object.max_players;
    } else {
      message.max_players = 0;
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = object.reward;
    } else {
      message.reward = "";
    }
    if (object.commit_timeout !== undefined && object.commit_timeout !== null) {
      message.commit_timeout = object.commit_timeout;
    } else {
      message.commit_timeout = "";
    }
    return message;
  },
};

const baseEventRevealTimeout: object = {
  game_id: "",
  reveal_timeout: "",
  number_of_players: 0,
};

export const EventRevealTimeout = {
  encode(
    message: EventRevealTimeout,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.game_id !== "") {
      writer.uint32(10).string(message.game_id);
    }
    if (message.reveal_timeout !== "") {
      writer.uint32(18).string(message.reveal_timeout);
    }
    if (message.number_of_players !== 0) {
      writer.uint32(24).uint64(message.number_of_players);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventRevealTimeout {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventRevealTimeout } as EventRevealTimeout;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.game_id = reader.string();
          break;
        case 2:
          message.reveal_timeout = reader.string();
          break;
        case 3:
          message.number_of_players = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventRevealTimeout {
    const message = { ...baseEventRevealTimeout } as EventRevealTimeout;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = String(object.game_id);
    } else {
      message.game_id = "";
    }
    if (object.reveal_timeout !== undefined && object.reveal_timeout !== null) {
      message.reveal_timeout = String(object.reveal_timeout);
    } else {
      message.reveal_timeout = "";
    }
    if (
      object.number_of_players !== undefined &&
      object.number_of_players !== null
    ) {
      message.number_of_players = Number(object.number_of_players);
    } else {
      message.number_of_players = 0;
    }
    return message;
  },

  toJSON(message: EventRevealTimeout): unknown {
    const obj: any = {};
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.reveal_timeout !== undefined &&
      (obj.reveal_timeout = message.reveal_timeout);
    message.number_of_players !== undefined &&
      (obj.number_of_players = message.number_of_players);
    return obj;
  },

  fromPartial(object: DeepPartial<EventRevealTimeout>): EventRevealTimeout {
    const message = { ...baseEventRevealTimeout } as EventRevealTimeout;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = "";
    }
    if (object.reveal_timeout !== undefined && object.reveal_timeout !== null) {
      message.reveal_timeout = object.reveal_timeout;
    } else {
      message.reveal_timeout = "";
    }
    if (
      object.number_of_players !== undefined &&
      object.number_of_players !== null
    ) {
      message.number_of_players = object.number_of_players;
    } else {
      message.number_of_players = 0;
    }
    return message;
  },
};

const baseEventGameEnd: object = { game_id: "" };

export const EventGameEnd = {
  encode(message: EventGameEnd, writer: Writer = Writer.create()): Writer {
    if (message.game_id !== "") {
      writer.uint32(10).string(message.game_id);
    }
    for (const v of message.winners) {
      Winner.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventGameEnd {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventGameEnd } as EventGameEnd;
    message.winners = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.game_id = reader.string();
          break;
        case 2:
          message.winners.push(Winner.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventGameEnd {
    const message = { ...baseEventGameEnd } as EventGameEnd;
    message.winners = [];
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = String(object.game_id);
    } else {
      message.game_id = "";
    }
    if (object.winners !== undefined && object.winners !== null) {
      for (const e of object.winners) {
        message.winners.push(Winner.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: EventGameEnd): unknown {
    const obj: any = {};
    message.game_id !== undefined && (obj.game_id = message.game_id);
    if (message.winners) {
      obj.winners = message.winners.map((e) =>
        e ? Winner.toJSON(e) : undefined
      );
    } else {
      obj.winners = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<EventGameEnd>): EventGameEnd {
    const message = { ...baseEventGameEnd } as EventGameEnd;
    message.winners = [];
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = "";
    }
    if (object.winners !== undefined && object.winners !== null) {
      for (const e of object.winners) {
        message.winners.push(Winner.fromPartial(e));
      }
    }
    return message;
  },
};

const baseWinner: object = { player: "", proximity: 0, reward: "" };

export const Winner = {
  encode(message: Winner, writer: Writer = Writer.create()): Writer {
    if (message.player !== "") {
      writer.uint32(10).string(message.player);
    }
    if (message.proximity !== 0) {
      writer.uint32(16).uint64(message.proximity);
    }
    if (message.reward !== "") {
      writer.uint32(26).string(message.reward);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Winner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWinner } as Winner;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.player = reader.string();
          break;
        case 2:
          message.proximity = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.reward = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Winner {
    const message = { ...baseWinner } as Winner;
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.proximity !== undefined && object.proximity !== null) {
      message.proximity = Number(object.proximity);
    } else {
      message.proximity = 0;
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = String(object.reward);
    } else {
      message.reward = "";
    }
    return message;
  },

  toJSON(message: Winner): unknown {
    const obj: any = {};
    message.player !== undefined && (obj.player = message.player);
    message.proximity !== undefined && (obj.proximity = message.proximity);
    message.reward !== undefined && (obj.reward = message.reward);
    return obj;
  },

  fromPartial(object: DeepPartial<Winner>): Winner {
    const message = { ...baseWinner } as Winner;
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.proximity !== undefined && object.proximity !== null) {
      message.proximity = object.proximity;
    } else {
      message.proximity = 0;
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = object.reward;
    } else {
      message.reward = "";
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
