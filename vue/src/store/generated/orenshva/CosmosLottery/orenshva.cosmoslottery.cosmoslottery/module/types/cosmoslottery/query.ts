/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cosmoslottery/params";
import { TxCounter } from "../cosmoslottery/tx_counter";

export const protobufPackage = "orenshva.cosmoslottery.cosmoslottery";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetTxCounterRequest {}

export interface QueryGetTxCounterResponse {
  TxCounter: TxCounter | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetTxCounterRequest: object = {};

export const QueryGetTxCounterRequest = {
  encode(
    _: QueryGetTxCounterRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTxCounterRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTxCounterRequest,
    } as QueryGetTxCounterRequest;
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

  fromJSON(_: any): QueryGetTxCounterRequest {
    const message = {
      ...baseQueryGetTxCounterRequest,
    } as QueryGetTxCounterRequest;
    return message;
  },

  toJSON(_: QueryGetTxCounterRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetTxCounterRequest>
  ): QueryGetTxCounterRequest {
    const message = {
      ...baseQueryGetTxCounterRequest,
    } as QueryGetTxCounterRequest;
    return message;
  },
};

const baseQueryGetTxCounterResponse: object = {};

export const QueryGetTxCounterResponse = {
  encode(
    message: QueryGetTxCounterResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.TxCounter !== undefined) {
      TxCounter.encode(message.TxCounter, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTxCounterResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTxCounterResponse,
    } as QueryGetTxCounterResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TxCounter = TxCounter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTxCounterResponse {
    const message = {
      ...baseQueryGetTxCounterResponse,
    } as QueryGetTxCounterResponse;
    if (object.TxCounter !== undefined && object.TxCounter !== null) {
      message.TxCounter = TxCounter.fromJSON(object.TxCounter);
    } else {
      message.TxCounter = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTxCounterResponse): unknown {
    const obj: any = {};
    message.TxCounter !== undefined &&
      (obj.TxCounter = message.TxCounter
        ? TxCounter.toJSON(message.TxCounter)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTxCounterResponse>
  ): QueryGetTxCounterResponse {
    const message = {
      ...baseQueryGetTxCounterResponse,
    } as QueryGetTxCounterResponse;
    if (object.TxCounter !== undefined && object.TxCounter !== null) {
      message.TxCounter = TxCounter.fromPartial(object.TxCounter);
    } else {
      message.TxCounter = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a TxCounter by index. */
  TxCounter(
    request: QueryGetTxCounterRequest
  ): Promise<QueryGetTxCounterResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "orenshva.cosmoslottery.cosmoslottery.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  TxCounter(
    request: QueryGetTxCounterRequest
  ): Promise<QueryGetTxCounterResponse> {
    const data = QueryGetTxCounterRequest.encode(request).finish();
    const promise = this.rpc.request(
      "orenshva.cosmoslottery.cosmoslottery.Query",
      "TxCounter",
      data
    );
    return promise.then((data) =>
      QueryGetTxCounterResponse.decode(new Reader(data))
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
