/* eslint-disable */
import {
  GameStatus,
  Winner,
  gameStatusFromJSON,
  gameStatusToJSON,
} from "../whichnumber/types";
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../whichnumber/params";

export const protobufPackage = "zale144.whichnumber.whichnumber";

export interface EventGameNew {
  creator: string;
  game_id: string;
  entry_fee: string;
  max_players: number;
  reward: string;
  commit_timeout: string;
  status: GameStatus;
}

export interface EventGameNewCommit {
  game_id: string;
  player: string;
  commit: string;
  number_of_commits: number;
  timestamp: string;
}

export interface EventGameNewReveal {
  game_id: string;
  player: string;
  reveal: string;
  number_of_reveals: number;
  number_of_commits: number;
  timestamp: string;
}

export interface EventGameCommitFinished {
  game_id: string;
  commit_timeout: string;
  number_of_commits: number;
}

export interface EventGameRevealFinished {
  game_id: string;
  reveal_timeout: string;
  number_of_reveals: number;
}

export interface EventGameEnd {
  game_id: string;
  winners: Winner[];
}

export interface EventGameCreatorDeposit {
  game_id: string;
  creator: string;
  amount: string;
}

export interface EventGamePlayerDeposit {
  game_id: string;
  player: string;
  amount: string;
}

export interface EventGameCreatorRefund {
  game_id: string;
  creator: string;
  amount: string;
}

export interface EventParamsUpdated {
  params: Params | undefined;
}

const baseEventGameNew: object = {
  creator: "",
  game_id: "",
  entry_fee: "",
  max_players: 0,
  reward: "",
  commit_timeout: "",
  status: 0,
};

export const EventGameNew = {
  encode(message: EventGameNew, writer: Writer = Writer.create()): Writer {
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
    if (message.status !== 0) {
      writer.uint32(56).int32(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventGameNew {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventGameNew } as EventGameNew;
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
        case 7:
          message.status = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventGameNew {
    const message = { ...baseEventGameNew } as EventGameNew;
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
    if (object.status !== undefined && object.status !== null) {
      message.status = gameStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: EventGameNew): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.entry_fee !== undefined && (obj.entry_fee = message.entry_fee);
    message.max_players !== undefined &&
      (obj.max_players = message.max_players);
    message.reward !== undefined && (obj.reward = message.reward);
    message.commit_timeout !== undefined &&
      (obj.commit_timeout = message.commit_timeout);
    message.status !== undefined &&
      (obj.status = gameStatusToJSON(message.status));
    return obj;
  },

  fromPartial(object: DeepPartial<EventGameNew>): EventGameNew {
    const message = { ...baseEventGameNew } as EventGameNew;
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
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseEventGameNewCommit: object = {
  game_id: "",
  player: "",
  commit: "",
  number_of_commits: 0,
  timestamp: "",
};

export const EventGameNewCommit = {
  encode(
    message: EventGameNewCommit,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.game_id !== "") {
      writer.uint32(10).string(message.game_id);
    }
    if (message.player !== "") {
      writer.uint32(18).string(message.player);
    }
    if (message.commit !== "") {
      writer.uint32(26).string(message.commit);
    }
    if (message.number_of_commits !== 0) {
      writer.uint32(32).uint64(message.number_of_commits);
    }
    if (message.timestamp !== "") {
      writer.uint32(42).string(message.timestamp);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventGameNewCommit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventGameNewCommit } as EventGameNewCommit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.game_id = reader.string();
          break;
        case 2:
          message.player = reader.string();
          break;
        case 3:
          message.commit = reader.string();
          break;
        case 4:
          message.number_of_commits = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.timestamp = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventGameNewCommit {
    const message = { ...baseEventGameNewCommit } as EventGameNewCommit;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = String(object.game_id);
    } else {
      message.game_id = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.commit !== undefined && object.commit !== null) {
      message.commit = String(object.commit);
    } else {
      message.commit = "";
    }
    if (
      object.number_of_commits !== undefined &&
      object.number_of_commits !== null
    ) {
      message.number_of_commits = Number(object.number_of_commits);
    } else {
      message.number_of_commits = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = String(object.timestamp);
    } else {
      message.timestamp = "";
    }
    return message;
  },

  toJSON(message: EventGameNewCommit): unknown {
    const obj: any = {};
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.player !== undefined && (obj.player = message.player);
    message.commit !== undefined && (obj.commit = message.commit);
    message.number_of_commits !== undefined &&
      (obj.number_of_commits = message.number_of_commits);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<EventGameNewCommit>): EventGameNewCommit {
    const message = { ...baseEventGameNewCommit } as EventGameNewCommit;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.commit !== undefined && object.commit !== null) {
      message.commit = object.commit;
    } else {
      message.commit = "";
    }
    if (
      object.number_of_commits !== undefined &&
      object.number_of_commits !== null
    ) {
      message.number_of_commits = object.number_of_commits;
    } else {
      message.number_of_commits = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = "";
    }
    return message;
  },
};

const baseEventGameNewReveal: object = {
  game_id: "",
  player: "",
  reveal: "",
  number_of_reveals: 0,
  number_of_commits: 0,
  timestamp: "",
};

export const EventGameNewReveal = {
  encode(
    message: EventGameNewReveal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.game_id !== "") {
      writer.uint32(10).string(message.game_id);
    }
    if (message.player !== "") {
      writer.uint32(18).string(message.player);
    }
    if (message.reveal !== "") {
      writer.uint32(26).string(message.reveal);
    }
    if (message.number_of_reveals !== 0) {
      writer.uint32(32).uint64(message.number_of_reveals);
    }
    if (message.number_of_commits !== 0) {
      writer.uint32(40).uint64(message.number_of_commits);
    }
    if (message.timestamp !== "") {
      writer.uint32(50).string(message.timestamp);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventGameNewReveal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventGameNewReveal } as EventGameNewReveal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.game_id = reader.string();
          break;
        case 2:
          message.player = reader.string();
          break;
        case 3:
          message.reveal = reader.string();
          break;
        case 4:
          message.number_of_reveals = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.number_of_commits = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.timestamp = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventGameNewReveal {
    const message = { ...baseEventGameNewReveal } as EventGameNewReveal;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = String(object.game_id);
    } else {
      message.game_id = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.reveal !== undefined && object.reveal !== null) {
      message.reveal = String(object.reveal);
    } else {
      message.reveal = "";
    }
    if (
      object.number_of_reveals !== undefined &&
      object.number_of_reveals !== null
    ) {
      message.number_of_reveals = Number(object.number_of_reveals);
    } else {
      message.number_of_reveals = 0;
    }
    if (
      object.number_of_commits !== undefined &&
      object.number_of_commits !== null
    ) {
      message.number_of_commits = Number(object.number_of_commits);
    } else {
      message.number_of_commits = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = String(object.timestamp);
    } else {
      message.timestamp = "";
    }
    return message;
  },

  toJSON(message: EventGameNewReveal): unknown {
    const obj: any = {};
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.player !== undefined && (obj.player = message.player);
    message.reveal !== undefined && (obj.reveal = message.reveal);
    message.number_of_reveals !== undefined &&
      (obj.number_of_reveals = message.number_of_reveals);
    message.number_of_commits !== undefined &&
      (obj.number_of_commits = message.number_of_commits);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<EventGameNewReveal>): EventGameNewReveal {
    const message = { ...baseEventGameNewReveal } as EventGameNewReveal;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.reveal !== undefined && object.reveal !== null) {
      message.reveal = object.reveal;
    } else {
      message.reveal = "";
    }
    if (
      object.number_of_reveals !== undefined &&
      object.number_of_reveals !== null
    ) {
      message.number_of_reveals = object.number_of_reveals;
    } else {
      message.number_of_reveals = 0;
    }
    if (
      object.number_of_commits !== undefined &&
      object.number_of_commits !== null
    ) {
      message.number_of_commits = object.number_of_commits;
    } else {
      message.number_of_commits = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = "";
    }
    return message;
  },
};

const baseEventGameCommitFinished: object = {
  game_id: "",
  commit_timeout: "",
  number_of_commits: 0,
};

export const EventGameCommitFinished = {
  encode(
    message: EventGameCommitFinished,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.game_id !== "") {
      writer.uint32(10).string(message.game_id);
    }
    if (message.commit_timeout !== "") {
      writer.uint32(18).string(message.commit_timeout);
    }
    if (message.number_of_commits !== 0) {
      writer.uint32(24).uint64(message.number_of_commits);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventGameCommitFinished {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventGameCommitFinished,
    } as EventGameCommitFinished;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.game_id = reader.string();
          break;
        case 2:
          message.commit_timeout = reader.string();
          break;
        case 3:
          message.number_of_commits = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventGameCommitFinished {
    const message = {
      ...baseEventGameCommitFinished,
    } as EventGameCommitFinished;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = String(object.game_id);
    } else {
      message.game_id = "";
    }
    if (object.commit_timeout !== undefined && object.commit_timeout !== null) {
      message.commit_timeout = String(object.commit_timeout);
    } else {
      message.commit_timeout = "";
    }
    if (
      object.number_of_commits !== undefined &&
      object.number_of_commits !== null
    ) {
      message.number_of_commits = Number(object.number_of_commits);
    } else {
      message.number_of_commits = 0;
    }
    return message;
  },

  toJSON(message: EventGameCommitFinished): unknown {
    const obj: any = {};
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.commit_timeout !== undefined &&
      (obj.commit_timeout = message.commit_timeout);
    message.number_of_commits !== undefined &&
      (obj.number_of_commits = message.number_of_commits);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventGameCommitFinished>
  ): EventGameCommitFinished {
    const message = {
      ...baseEventGameCommitFinished,
    } as EventGameCommitFinished;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = "";
    }
    if (object.commit_timeout !== undefined && object.commit_timeout !== null) {
      message.commit_timeout = object.commit_timeout;
    } else {
      message.commit_timeout = "";
    }
    if (
      object.number_of_commits !== undefined &&
      object.number_of_commits !== null
    ) {
      message.number_of_commits = object.number_of_commits;
    } else {
      message.number_of_commits = 0;
    }
    return message;
  },
};

const baseEventGameRevealFinished: object = {
  game_id: "",
  reveal_timeout: "",
  number_of_reveals: 0,
};

export const EventGameRevealFinished = {
  encode(
    message: EventGameRevealFinished,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.game_id !== "") {
      writer.uint32(10).string(message.game_id);
    }
    if (message.reveal_timeout !== "") {
      writer.uint32(18).string(message.reveal_timeout);
    }
    if (message.number_of_reveals !== 0) {
      writer.uint32(24).uint64(message.number_of_reveals);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventGameRevealFinished {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventGameRevealFinished,
    } as EventGameRevealFinished;
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
          message.number_of_reveals = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventGameRevealFinished {
    const message = {
      ...baseEventGameRevealFinished,
    } as EventGameRevealFinished;
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
      object.number_of_reveals !== undefined &&
      object.number_of_reveals !== null
    ) {
      message.number_of_reveals = Number(object.number_of_reveals);
    } else {
      message.number_of_reveals = 0;
    }
    return message;
  },

  toJSON(message: EventGameRevealFinished): unknown {
    const obj: any = {};
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.reveal_timeout !== undefined &&
      (obj.reveal_timeout = message.reveal_timeout);
    message.number_of_reveals !== undefined &&
      (obj.number_of_reveals = message.number_of_reveals);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventGameRevealFinished>
  ): EventGameRevealFinished {
    const message = {
      ...baseEventGameRevealFinished,
    } as EventGameRevealFinished;
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
      object.number_of_reveals !== undefined &&
      object.number_of_reveals !== null
    ) {
      message.number_of_reveals = object.number_of_reveals;
    } else {
      message.number_of_reveals = 0;
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

const baseEventGameCreatorDeposit: object = {
  game_id: "",
  creator: "",
  amount: "",
};

export const EventGameCreatorDeposit = {
  encode(
    message: EventGameCreatorDeposit,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.game_id !== "") {
      writer.uint32(10).string(message.game_id);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventGameCreatorDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventGameCreatorDeposit,
    } as EventGameCreatorDeposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.game_id = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventGameCreatorDeposit {
    const message = {
      ...baseEventGameCreatorDeposit,
    } as EventGameCreatorDeposit;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = String(object.game_id);
    } else {
      message.game_id = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: EventGameCreatorDeposit): unknown {
    const obj: any = {};
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventGameCreatorDeposit>
  ): EventGameCreatorDeposit {
    const message = {
      ...baseEventGameCreatorDeposit,
    } as EventGameCreatorDeposit;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseEventGamePlayerDeposit: object = {
  game_id: "",
  player: "",
  amount: "",
};

export const EventGamePlayerDeposit = {
  encode(
    message: EventGamePlayerDeposit,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.game_id !== "") {
      writer.uint32(10).string(message.game_id);
    }
    if (message.player !== "") {
      writer.uint32(18).string(message.player);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventGamePlayerDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventGamePlayerDeposit } as EventGamePlayerDeposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.game_id = reader.string();
          break;
        case 2:
          message.player = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventGamePlayerDeposit {
    const message = { ...baseEventGamePlayerDeposit } as EventGamePlayerDeposit;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = String(object.game_id);
    } else {
      message.game_id = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: EventGamePlayerDeposit): unknown {
    const obj: any = {};
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.player !== undefined && (obj.player = message.player);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventGamePlayerDeposit>
  ): EventGamePlayerDeposit {
    const message = { ...baseEventGamePlayerDeposit } as EventGamePlayerDeposit;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseEventGameCreatorRefund: object = {
  game_id: "",
  creator: "",
  amount: "",
};

export const EventGameCreatorRefund = {
  encode(
    message: EventGameCreatorRefund,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.game_id !== "") {
      writer.uint32(10).string(message.game_id);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventGameCreatorRefund {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventGameCreatorRefund } as EventGameCreatorRefund;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.game_id = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventGameCreatorRefund {
    const message = { ...baseEventGameCreatorRefund } as EventGameCreatorRefund;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = String(object.game_id);
    } else {
      message.game_id = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: EventGameCreatorRefund): unknown {
    const obj: any = {};
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventGameCreatorRefund>
  ): EventGameCreatorRefund {
    const message = { ...baseEventGameCreatorRefund } as EventGameCreatorRefund;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseEventParamsUpdated: object = {};

export const EventParamsUpdated = {
  encode(
    message: EventParamsUpdated,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventParamsUpdated {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventParamsUpdated } as EventParamsUpdated;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventParamsUpdated {
    const message = { ...baseEventParamsUpdated } as EventParamsUpdated;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: EventParamsUpdated): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<EventParamsUpdated>): EventParamsUpdated {
    const message = { ...baseEventParamsUpdated } as EventParamsUpdated;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
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
