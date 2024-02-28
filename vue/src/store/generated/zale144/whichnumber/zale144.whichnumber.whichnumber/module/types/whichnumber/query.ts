/* eslint-disable */
import {
  GameStatus,
  NumberCommit,
  Winner,
  gameStatusFromJSON,
  gameStatusToJSON,
} from "../whichnumber/types";
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { SystemInfo } from "../whichnumber/system_info";
import { Params } from "../whichnumber/params";

export const protobufPackage = "zale144.whichnumber.whichnumber";

/** QueryGetGameRequest is the request type for the Query/GetGame RPC method. */
export interface QueryGetGameRequest {
  /** id defines the id of the game to query for. */
  id: number;
}

/** QueryGetGameResponse is the response type for the Query/GetGame RPC method. */
export interface QueryGetGameResponse {
  game: QueryGameResponse | undefined;
}

/** QueryGamesRequest is the request type for the Query/Games RPC method. */
export interface QueryGetGamesRequest {
  pagination: PageRequest | undefined;
}

/** QueryGamesResponse is the response type for the Query/Games RPC method. */
export interface QueryGetGamesResponse {
  games: QueryGameResponse[];
  pagination: PageResponse | undefined;
}

/** QueryGetGameResponse is the response type for the Query/GetGame RPC method. */
export interface QueryGameResponse {
  id: number;
  creator: string;
  player_commits: NumberCommit[];
  player_reveals: QueryPlayerReveal[];
  reward: string;
  entry_fee: string;
  commit_timeout: Date | undefined;
  reveal_timeout: Date | undefined;
  status: GameStatus;
  winners: Winner[];
}

export interface QueryPlayerReveal {
  player_address: string;
  proximity: number;
  created_at: Date | undefined;
}

/** QueryGetSystemRequest is the request type for the Query/GetSystem RPC method */
export interface QueryGetSystemInfoRequest {}

/** QueryGetSystemInfoResponse is the response type for the Query/GetSystemInfo RPC method */
export interface QueryGetSystemInfoResponse {
  SystemInfo: SystemInfo | undefined;
}

/** QueryGetParamsRequest is request type for the Query/Params RPC method. */
export interface QueryGetParamsRequest {}

/** QueryGetParamsResponse is response type for the Query/Params RPC method. */
export interface QueryGetParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

const baseQueryGetGameRequest: object = { id: 0 };

export const QueryGetGameRequest = {
  encode(
    message: QueryGetGameRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetGameRequest } as QueryGetGameRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGameRequest {
    const message = { ...baseQueryGetGameRequest } as QueryGetGameRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetGameRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetGameRequest>): QueryGetGameRequest {
    const message = { ...baseQueryGetGameRequest } as QueryGetGameRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetGameResponse: object = {};

export const QueryGetGameResponse = {
  encode(
    message: QueryGetGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.game !== undefined) {
      QueryGameResponse.encode(message.game, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetGameResponse } as QueryGetGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.game = QueryGameResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGameResponse {
    const message = { ...baseQueryGetGameResponse } as QueryGetGameResponse;
    if (object.game !== undefined && object.game !== null) {
      message.game = QueryGameResponse.fromJSON(object.game);
    } else {
      message.game = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetGameResponse): unknown {
    const obj: any = {};
    message.game !== undefined &&
      (obj.game = message.game
        ? QueryGameResponse.toJSON(message.game)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetGameResponse>): QueryGetGameResponse {
    const message = { ...baseQueryGetGameResponse } as QueryGetGameResponse;
    if (object.game !== undefined && object.game !== null) {
      message.game = QueryGameResponse.fromPartial(object.game);
    } else {
      message.game = undefined;
    }
    return message;
  },
};

const baseQueryGetGamesRequest: object = {};

export const QueryGetGamesRequest = {
  encode(
    message: QueryGetGamesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGamesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetGamesRequest } as QueryGetGamesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGamesRequest {
    const message = { ...baseQueryGetGamesRequest } as QueryGetGamesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetGamesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetGamesRequest>): QueryGetGamesRequest {
    const message = { ...baseQueryGetGamesRequest } as QueryGetGamesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetGamesResponse: object = {};

export const QueryGetGamesResponse = {
  encode(
    message: QueryGetGamesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.games) {
      QueryGameResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGamesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetGamesResponse } as QueryGetGamesResponse;
    message.games = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.games.push(QueryGameResponse.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGamesResponse {
    const message = { ...baseQueryGetGamesResponse } as QueryGetGamesResponse;
    message.games = [];
    if (object.games !== undefined && object.games !== null) {
      for (const e of object.games) {
        message.games.push(QueryGameResponse.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetGamesResponse): unknown {
    const obj: any = {};
    if (message.games) {
      obj.games = message.games.map((e) =>
        e ? QueryGameResponse.toJSON(e) : undefined
      );
    } else {
      obj.games = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetGamesResponse>
  ): QueryGetGamesResponse {
    const message = { ...baseQueryGetGamesResponse } as QueryGetGamesResponse;
    message.games = [];
    if (object.games !== undefined && object.games !== null) {
      for (const e of object.games) {
        message.games.push(QueryGameResponse.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGameResponse: object = {
  id: 0,
  creator: "",
  reward: "",
  entry_fee: "",
  status: 0,
};

export const QueryGameResponse = {
  encode(message: QueryGameResponse, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    for (const v of message.player_commits) {
      NumberCommit.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.player_reveals) {
      QueryPlayerReveal.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.reward !== "") {
      writer.uint32(50).string(message.reward);
    }
    if (message.entry_fee !== "") {
      writer.uint32(58).string(message.entry_fee);
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
    if (message.status !== 0) {
      writer.uint32(80).int32(message.status);
    }
    for (const v of message.winners) {
      Winner.encode(v!, writer.uint32(90).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGameResponse } as QueryGameResponse;
    message.player_commits = [];
    message.player_reveals = [];
    message.winners = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 4:
          message.player_commits.push(
            NumberCommit.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.player_reveals.push(
            QueryPlayerReveal.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.reward = reader.string();
          break;
        case 7:
          message.entry_fee = reader.string();
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
          message.status = reader.int32() as any;
          break;
        case 11:
          message.winners.push(Winner.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGameResponse {
    const message = { ...baseQueryGameResponse } as QueryGameResponse;
    message.player_commits = [];
    message.player_reveals = [];
    message.winners = [];
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
    if (object.player_commits !== undefined && object.player_commits !== null) {
      for (const e of object.player_commits) {
        message.player_commits.push(NumberCommit.fromJSON(e));
      }
    }
    if (object.player_reveals !== undefined && object.player_reveals !== null) {
      for (const e of object.player_reveals) {
        message.player_reveals.push(QueryPlayerReveal.fromJSON(e));
      }
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = String(object.reward);
    } else {
      message.reward = "";
    }
    if (object.entry_fee !== undefined && object.entry_fee !== null) {
      message.entry_fee = String(object.entry_fee);
    } else {
      message.entry_fee = "";
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
    if (object.status !== undefined && object.status !== null) {
      message.status = gameStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (object.winners !== undefined && object.winners !== null) {
      for (const e of object.winners) {
        message.winners.push(Winner.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryGameResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.player_commits) {
      obj.player_commits = message.player_commits.map((e) =>
        e ? NumberCommit.toJSON(e) : undefined
      );
    } else {
      obj.player_commits = [];
    }
    if (message.player_reveals) {
      obj.player_reveals = message.player_reveals.map((e) =>
        e ? QueryPlayerReveal.toJSON(e) : undefined
      );
    } else {
      obj.player_reveals = [];
    }
    message.reward !== undefined && (obj.reward = message.reward);
    message.entry_fee !== undefined && (obj.entry_fee = message.entry_fee);
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
    message.status !== undefined &&
      (obj.status = gameStatusToJSON(message.status));
    if (message.winners) {
      obj.winners = message.winners.map((e) =>
        e ? Winner.toJSON(e) : undefined
      );
    } else {
      obj.winners = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGameResponse>): QueryGameResponse {
    const message = { ...baseQueryGameResponse } as QueryGameResponse;
    message.player_commits = [];
    message.player_reveals = [];
    message.winners = [];
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
    if (object.player_commits !== undefined && object.player_commits !== null) {
      for (const e of object.player_commits) {
        message.player_commits.push(NumberCommit.fromPartial(e));
      }
    }
    if (object.player_reveals !== undefined && object.player_reveals !== null) {
      for (const e of object.player_reveals) {
        message.player_reveals.push(QueryPlayerReveal.fromPartial(e));
      }
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = object.reward;
    } else {
      message.reward = "";
    }
    if (object.entry_fee !== undefined && object.entry_fee !== null) {
      message.entry_fee = object.entry_fee;
    } else {
      message.entry_fee = "";
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
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (object.winners !== undefined && object.winners !== null) {
      for (const e of object.winners) {
        message.winners.push(Winner.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryPlayerReveal: object = { player_address: "", proximity: 0 };

export const QueryPlayerReveal = {
  encode(message: QueryPlayerReveal, writer: Writer = Writer.create()): Writer {
    if (message.player_address !== "") {
      writer.uint32(10).string(message.player_address);
    }
    if (message.proximity !== 0) {
      writer.uint32(16).uint64(message.proximity);
    }
    if (message.created_at !== undefined) {
      Timestamp.encode(
        toTimestamp(message.created_at),
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryPlayerReveal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryPlayerReveal } as QueryPlayerReveal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.player_address = reader.string();
          break;
        case 2:
          message.proximity = longToNumber(reader.uint64() as Long);
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

  fromJSON(object: any): QueryPlayerReveal {
    const message = { ...baseQueryPlayerReveal } as QueryPlayerReveal;
    if (object.player_address !== undefined && object.player_address !== null) {
      message.player_address = String(object.player_address);
    } else {
      message.player_address = "";
    }
    if (object.proximity !== undefined && object.proximity !== null) {
      message.proximity = Number(object.proximity);
    } else {
      message.proximity = 0;
    }
    if (object.created_at !== undefined && object.created_at !== null) {
      message.created_at = fromJsonTimestamp(object.created_at);
    } else {
      message.created_at = undefined;
    }
    return message;
  },

  toJSON(message: QueryPlayerReveal): unknown {
    const obj: any = {};
    message.player_address !== undefined &&
      (obj.player_address = message.player_address);
    message.proximity !== undefined && (obj.proximity = message.proximity);
    message.created_at !== undefined &&
      (obj.created_at =
        message.created_at !== undefined
          ? message.created_at.toISOString()
          : null);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryPlayerReveal>): QueryPlayerReveal {
    const message = { ...baseQueryPlayerReveal } as QueryPlayerReveal;
    if (object.player_address !== undefined && object.player_address !== null) {
      message.player_address = object.player_address;
    } else {
      message.player_address = "";
    }
    if (object.proximity !== undefined && object.proximity !== null) {
      message.proximity = object.proximity;
    } else {
      message.proximity = 0;
    }
    if (object.created_at !== undefined && object.created_at !== null) {
      message.created_at = object.created_at;
    } else {
      message.created_at = undefined;
    }
    return message;
  },
};

const baseQueryGetSystemInfoRequest: object = {};

export const QueryGetSystemInfoRequest = {
  encode(
    _: QueryGetSystemInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    return message;
  },

  toJSON(_: QueryGetSystemInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetSystemInfoRequest>
  ): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    return message;
  },
};

const baseQueryGetSystemInfoResponse: object = {};

export const QueryGetSystemInfoResponse = {
  encode(
    message: QueryGetSystemInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.SystemInfo !== undefined) {
      SystemInfo.encode(message.SystemInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SystemInfo = SystemInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.SystemInfo !== undefined && object.SystemInfo !== null) {
      message.SystemInfo = SystemInfo.fromJSON(object.SystemInfo);
    } else {
      message.SystemInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSystemInfoResponse): unknown {
    const obj: any = {};
    message.SystemInfo !== undefined &&
      (obj.SystemInfo = message.SystemInfo
        ? SystemInfo.toJSON(message.SystemInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSystemInfoResponse>
  ): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.SystemInfo !== undefined && object.SystemInfo !== null) {
      message.SystemInfo = SystemInfo.fromPartial(object.SystemInfo);
    } else {
      message.SystemInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetParamsRequest: object = {};

export const QueryGetParamsRequest = {
  encode(_: QueryGetParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetParamsRequest } as QueryGetParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetParamsRequest {
    const message = { ...baseQueryGetParamsRequest } as QueryGetParamsRequest;
    return message;
  },

  toJSON(_: QueryGetParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryGetParamsRequest>): QueryGetParamsRequest {
    const message = { ...baseQueryGetParamsRequest } as QueryGetParamsRequest;
    return message;
  },
};

const baseQueryGetParamsResponse: object = {};

export const QueryGetParamsResponse = {
  encode(
    message: QueryGetParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetParamsResponse } as QueryGetParamsResponse;
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

  fromJSON(object: any): QueryGetParamsResponse {
    const message = { ...baseQueryGetParamsResponse } as QueryGetParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetParamsResponse>
  ): QueryGetParamsResponse {
    const message = { ...baseQueryGetParamsResponse } as QueryGetParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  GetGame(request: QueryGetGameRequest): Promise<QueryGetGameResponse>;
  GetGames(request: QueryGetGamesRequest): Promise<QueryGetGamesResponse>;
  GetSystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse>;
  /** Parameters queries the parameters of the module. */
  GetParams(request: QueryGetParamsRequest): Promise<QueryGetParamsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  GetGame(request: QueryGetGameRequest): Promise<QueryGetGameResponse> {
    const data = QueryGetGameRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zale144.whichnumber.whichnumber.Query",
      "GetGame",
      data
    );
    return promise.then((data) =>
      QueryGetGameResponse.decode(new Reader(data))
    );
  }

  GetGames(request: QueryGetGamesRequest): Promise<QueryGetGamesResponse> {
    const data = QueryGetGamesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zale144.whichnumber.whichnumber.Query",
      "GetGames",
      data
    );
    return promise.then((data) =>
      QueryGetGamesResponse.decode(new Reader(data))
    );
  }

  GetSystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse> {
    const data = QueryGetSystemInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zale144.whichnumber.whichnumber.Query",
      "GetSystemInfo",
      data
    );
    return promise.then((data) =>
      QueryGetSystemInfoResponse.decode(new Reader(data))
    );
  }

  GetParams(request: QueryGetParamsRequest): Promise<QueryGetParamsResponse> {
    const data = QueryGetParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zale144.whichnumber.whichnumber.Query",
      "GetParams",
      data
    );
    return promise.then((data) =>
      QueryGetParamsResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
