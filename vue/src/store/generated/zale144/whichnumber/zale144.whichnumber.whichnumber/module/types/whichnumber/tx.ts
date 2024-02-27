/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Params } from "../whichnumber/params";

export const protobufPackage = "zale144.whichnumber.whichnumber";

export interface MsgNewGame {
  /** creator is the address of the player that created the game. */
  creator: string;
  /** secret_number is the number to guess. */
  secret_number: number;
  /** reward is the amount to be distributed to the winner(s). */
  reward: Coin | undefined;
  /** entry_fee is the amount to put into stake for the game. */
  entry_fee: Coin | undefined;
}

export interface MsgNewGameResponse {
  /** game_id is the ID of the created game. */
  game_id: number;
}

export interface MsgCommitNumber {
  player: string;
  /** game_id is the ID of the game to commit the number to. */
  game_id: number;
  /**
   * commit is the hex encoded commitment to the number.
   * SHA256("32byte-salt" + "number")
   */
  commit: string;
}

export interface MsgCommitNumberResponse {}

export interface MsgRevealNumber {
  player: string;
  /** game_id is the ID of the game to reveal the number for. */
  game_id: number;
  /** number is the number to reveal. */
  number: number;
  /** salt is the salt used to create the commitment. */
  salt: string;
}

export interface MsgRevealNumberResponse {}

/** MsgUpdateParams is the Msg/UpdateParams request type. */
export interface MsgUpdateParams {
  /**
   * authority is the address that controls the module
   * NOTE: Defaults to the governance module unless overwritten.
   */
  authority: string;
  /**
   * params defines the module parameters to update.
   * NOTE: All parameters must be supplied.
   */
  params: Params | undefined;
}

/**
 * MsgUpdateParamsResponse defines the response structure for executing a
 * MsgUpdateParams message.
 */
export interface MsgUpdateParamsResponse {}

const baseMsgNewGame: object = { creator: "", secret_number: 0 };

export const MsgNewGame = {
  encode(message: MsgNewGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.secret_number !== 0) {
      writer.uint32(16).int64(message.secret_number);
    }
    if (message.reward !== undefined) {
      Coin.encode(message.reward, writer.uint32(26).fork()).ldelim();
    }
    if (message.entry_fee !== undefined) {
      Coin.encode(message.entry_fee, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgNewGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgNewGame } as MsgNewGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.secret_number = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.reward = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.entry_fee = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgNewGame {
    const message = { ...baseMsgNewGame } as MsgNewGame;
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
    return message;
  },

  toJSON(message: MsgNewGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.secret_number !== undefined &&
      (obj.secret_number = message.secret_number);
    message.reward !== undefined &&
      (obj.reward = message.reward ? Coin.toJSON(message.reward) : undefined);
    message.entry_fee !== undefined &&
      (obj.entry_fee = message.entry_fee
        ? Coin.toJSON(message.entry_fee)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgNewGame>): MsgNewGame {
    const message = { ...baseMsgNewGame } as MsgNewGame;
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
    return message;
  },
};

const baseMsgNewGameResponse: object = { game_id: 0 };

export const MsgNewGameResponse = {
  encode(
    message: MsgNewGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.game_id !== 0) {
      writer.uint32(8).int64(message.game_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgNewGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgNewGameResponse } as MsgNewGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.game_id = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgNewGameResponse {
    const message = { ...baseMsgNewGameResponse } as MsgNewGameResponse;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = Number(object.game_id);
    } else {
      message.game_id = 0;
    }
    return message;
  },

  toJSON(message: MsgNewGameResponse): unknown {
    const obj: any = {};
    message.game_id !== undefined && (obj.game_id = message.game_id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgNewGameResponse>): MsgNewGameResponse {
    const message = { ...baseMsgNewGameResponse } as MsgNewGameResponse;
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = 0;
    }
    return message;
  },
};

const baseMsgCommitNumber: object = { player: "", game_id: 0, commit: "" };

export const MsgCommitNumber = {
  encode(message: MsgCommitNumber, writer: Writer = Writer.create()): Writer {
    if (message.player !== "") {
      writer.uint32(10).string(message.player);
    }
    if (message.game_id !== 0) {
      writer.uint32(16).int64(message.game_id);
    }
    if (message.commit !== "") {
      writer.uint32(26).string(message.commit);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCommitNumber {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCommitNumber } as MsgCommitNumber;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.player = reader.string();
          break;
        case 2:
          message.game_id = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.commit = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCommitNumber {
    const message = { ...baseMsgCommitNumber } as MsgCommitNumber;
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = Number(object.game_id);
    } else {
      message.game_id = 0;
    }
    if (object.commit !== undefined && object.commit !== null) {
      message.commit = String(object.commit);
    } else {
      message.commit = "";
    }
    return message;
  },

  toJSON(message: MsgCommitNumber): unknown {
    const obj: any = {};
    message.player !== undefined && (obj.player = message.player);
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.commit !== undefined && (obj.commit = message.commit);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCommitNumber>): MsgCommitNumber {
    const message = { ...baseMsgCommitNumber } as MsgCommitNumber;
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = 0;
    }
    if (object.commit !== undefined && object.commit !== null) {
      message.commit = object.commit;
    } else {
      message.commit = "";
    }
    return message;
  },
};

const baseMsgCommitNumberResponse: object = {};

export const MsgCommitNumberResponse = {
  encode(_: MsgCommitNumberResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCommitNumberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCommitNumberResponse,
    } as MsgCommitNumberResponse;
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

  fromJSON(_: any): MsgCommitNumberResponse {
    const message = {
      ...baseMsgCommitNumberResponse,
    } as MsgCommitNumberResponse;
    return message;
  },

  toJSON(_: MsgCommitNumberResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCommitNumberResponse>
  ): MsgCommitNumberResponse {
    const message = {
      ...baseMsgCommitNumberResponse,
    } as MsgCommitNumberResponse;
    return message;
  },
};

const baseMsgRevealNumber: object = {
  player: "",
  game_id: 0,
  number: 0,
  salt: "",
};

export const MsgRevealNumber = {
  encode(message: MsgRevealNumber, writer: Writer = Writer.create()): Writer {
    if (message.player !== "") {
      writer.uint32(10).string(message.player);
    }
    if (message.game_id !== 0) {
      writer.uint32(16).int64(message.game_id);
    }
    if (message.number !== 0) {
      writer.uint32(24).int64(message.number);
    }
    if (message.salt !== "") {
      writer.uint32(34).string(message.salt);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRevealNumber {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRevealNumber } as MsgRevealNumber;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.player = reader.string();
          break;
        case 2:
          message.game_id = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.number = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.salt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRevealNumber {
    const message = { ...baseMsgRevealNumber } as MsgRevealNumber;
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = Number(object.game_id);
    } else {
      message.game_id = 0;
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
    return message;
  },

  toJSON(message: MsgRevealNumber): unknown {
    const obj: any = {};
    message.player !== undefined && (obj.player = message.player);
    message.game_id !== undefined && (obj.game_id = message.game_id);
    message.number !== undefined && (obj.number = message.number);
    message.salt !== undefined && (obj.salt = message.salt);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRevealNumber>): MsgRevealNumber {
    const message = { ...baseMsgRevealNumber } as MsgRevealNumber;
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.game_id !== undefined && object.game_id !== null) {
      message.game_id = object.game_id;
    } else {
      message.game_id = 0;
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
    return message;
  },
};

const baseMsgRevealNumberResponse: object = {};

export const MsgRevealNumberResponse = {
  encode(_: MsgRevealNumberResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRevealNumberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRevealNumberResponse,
    } as MsgRevealNumberResponse;
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

  fromJSON(_: any): MsgRevealNumberResponse {
    const message = {
      ...baseMsgRevealNumberResponse,
    } as MsgRevealNumberResponse;
    return message;
  },

  toJSON(_: MsgRevealNumberResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRevealNumberResponse>
  ): MsgRevealNumberResponse {
    const message = {
      ...baseMsgRevealNumberResponse,
    } as MsgRevealNumberResponse;
    return message;
  },
};

const baseMsgUpdateParams: object = { authority: "" };

export const MsgUpdateParams = {
  encode(message: MsgUpdateParams, writer: Writer = Writer.create()): Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateParams {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateParams } as MsgUpdateParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.authority = reader.string();
          break;
        case 2:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateParams {
    const message = { ...baseMsgUpdateParams } as MsgUpdateParams;
    if (object.authority !== undefined && object.authority !== null) {
      message.authority = String(object.authority);
    } else {
      message.authority = "";
    }
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: MsgUpdateParams): unknown {
    const obj: any = {};
    message.authority !== undefined && (obj.authority = message.authority);
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateParams>): MsgUpdateParams {
    const message = { ...baseMsgUpdateParams } as MsgUpdateParams;
    if (object.authority !== undefined && object.authority !== null) {
      message.authority = object.authority;
    } else {
      message.authority = "";
    }
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseMsgUpdateParamsResponse: object = {};

export const MsgUpdateParamsResponse = {
  encode(_: MsgUpdateParamsResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateParamsResponse,
    } as MsgUpdateParamsResponse;
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

  fromJSON(_: any): MsgUpdateParamsResponse {
    const message = {
      ...baseMsgUpdateParamsResponse,
    } as MsgUpdateParamsResponse;
    return message;
  },

  toJSON(_: MsgUpdateParamsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateParamsResponse>
  ): MsgUpdateParamsResponse {
    const message = {
      ...baseMsgUpdateParamsResponse,
    } as MsgUpdateParamsResponse;
    return message;
  },
};

/** Msg defines the module Msg service. */
export interface Msg {
  NewGame(request: MsgNewGame): Promise<MsgNewGameResponse>;
  CommitNumber(request: MsgCommitNumber): Promise<MsgCommitNumberResponse>;
  RevealNumber(request: MsgRevealNumber): Promise<MsgRevealNumberResponse>;
  /** UpdateParams updates the module parameters. */
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  NewGame(request: MsgNewGame): Promise<MsgNewGameResponse> {
    const data = MsgNewGame.encode(request).finish();
    const promise = this.rpc.request(
      "zale144.whichnumber.whichnumber.Msg",
      "NewGame",
      data
    );
    return promise.then((data) => MsgNewGameResponse.decode(new Reader(data)));
  }

  CommitNumber(request: MsgCommitNumber): Promise<MsgCommitNumberResponse> {
    const data = MsgCommitNumber.encode(request).finish();
    const promise = this.rpc.request(
      "zale144.whichnumber.whichnumber.Msg",
      "CommitNumber",
      data
    );
    return promise.then((data) =>
      MsgCommitNumberResponse.decode(new Reader(data))
    );
  }

  RevealNumber(request: MsgRevealNumber): Promise<MsgRevealNumberResponse> {
    const data = MsgRevealNumber.encode(request).finish();
    const promise = this.rpc.request(
      "zale144.whichnumber.whichnumber.Msg",
      "RevealNumber",
      data
    );
    return promise.then((data) =>
      MsgRevealNumberResponse.decode(new Reader(data))
    );
  }

  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse> {
    const data = MsgUpdateParams.encode(request).finish();
    const promise = this.rpc.request(
      "zale144.whichnumber.whichnumber.Msg",
      "UpdateParams",
      data
    );
    return promise.then((data) =>
      MsgUpdateParamsResponse.decode(new Reader(data))
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
