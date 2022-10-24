/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cosmoslottery/params";
import { TxCounter } from "../cosmoslottery/tx_counter";
import { BetChart } from "../cosmoslottery/bet_chart";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { FeeCounter } from "../cosmoslottery/fee_counter";

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

export interface QueryGetBetChartRequest {
  accountName: string;
}

export interface QueryGetBetChartResponse {
  betChart: BetChart | undefined;
}

export interface QueryAllBetChartRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllBetChartResponse {
  betChart: BetChart[];
  pagination: PageResponse | undefined;
}

export interface QueryGetFeeCounterRequest {}

export interface QueryGetFeeCounterResponse {
  FeeCounter: FeeCounter | undefined;
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

const baseQueryGetBetChartRequest: object = { accountName: "" };

export const QueryGetBetChartRequest = {
  encode(
    message: QueryGetBetChartRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.accountName !== "") {
      writer.uint32(10).string(message.accountName);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBetChartRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBetChartRequest,
    } as QueryGetBetChartRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBetChartRequest {
    const message = {
      ...baseQueryGetBetChartRequest,
    } as QueryGetBetChartRequest;
    if (object.accountName !== undefined && object.accountName !== null) {
      message.accountName = String(object.accountName);
    } else {
      message.accountName = "";
    }
    return message;
  },

  toJSON(message: QueryGetBetChartRequest): unknown {
    const obj: any = {};
    message.accountName !== undefined &&
      (obj.accountName = message.accountName);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBetChartRequest>
  ): QueryGetBetChartRequest {
    const message = {
      ...baseQueryGetBetChartRequest,
    } as QueryGetBetChartRequest;
    if (object.accountName !== undefined && object.accountName !== null) {
      message.accountName = object.accountName;
    } else {
      message.accountName = "";
    }
    return message;
  },
};

const baseQueryGetBetChartResponse: object = {};

export const QueryGetBetChartResponse = {
  encode(
    message: QueryGetBetChartResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.betChart !== undefined) {
      BetChart.encode(message.betChart, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetBetChartResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBetChartResponse,
    } as QueryGetBetChartResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.betChart = BetChart.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBetChartResponse {
    const message = {
      ...baseQueryGetBetChartResponse,
    } as QueryGetBetChartResponse;
    if (object.betChart !== undefined && object.betChart !== null) {
      message.betChart = BetChart.fromJSON(object.betChart);
    } else {
      message.betChart = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetBetChartResponse): unknown {
    const obj: any = {};
    message.betChart !== undefined &&
      (obj.betChart = message.betChart
        ? BetChart.toJSON(message.betChart)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBetChartResponse>
  ): QueryGetBetChartResponse {
    const message = {
      ...baseQueryGetBetChartResponse,
    } as QueryGetBetChartResponse;
    if (object.betChart !== undefined && object.betChart !== null) {
      message.betChart = BetChart.fromPartial(object.betChart);
    } else {
      message.betChart = undefined;
    }
    return message;
  },
};

const baseQueryAllBetChartRequest: object = {};

export const QueryAllBetChartRequest = {
  encode(
    message: QueryAllBetChartRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBetChartRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllBetChartRequest,
    } as QueryAllBetChartRequest;
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

  fromJSON(object: any): QueryAllBetChartRequest {
    const message = {
      ...baseQueryAllBetChartRequest,
    } as QueryAllBetChartRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBetChartRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBetChartRequest>
  ): QueryAllBetChartRequest {
    const message = {
      ...baseQueryAllBetChartRequest,
    } as QueryAllBetChartRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllBetChartResponse: object = {};

export const QueryAllBetChartResponse = {
  encode(
    message: QueryAllBetChartResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.betChart) {
      BetChart.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllBetChartResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllBetChartResponse,
    } as QueryAllBetChartResponse;
    message.betChart = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.betChart.push(BetChart.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllBetChartResponse {
    const message = {
      ...baseQueryAllBetChartResponse,
    } as QueryAllBetChartResponse;
    message.betChart = [];
    if (object.betChart !== undefined && object.betChart !== null) {
      for (const e of object.betChart) {
        message.betChart.push(BetChart.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBetChartResponse): unknown {
    const obj: any = {};
    if (message.betChart) {
      obj.betChart = message.betChart.map((e) =>
        e ? BetChart.toJSON(e) : undefined
      );
    } else {
      obj.betChart = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBetChartResponse>
  ): QueryAllBetChartResponse {
    const message = {
      ...baseQueryAllBetChartResponse,
    } as QueryAllBetChartResponse;
    message.betChart = [];
    if (object.betChart !== undefined && object.betChart !== null) {
      for (const e of object.betChart) {
        message.betChart.push(BetChart.fromPartial(e));
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

const baseQueryGetFeeCounterRequest: object = {};

export const QueryGetFeeCounterRequest = {
  encode(
    _: QueryGetFeeCounterRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetFeeCounterRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetFeeCounterRequest,
    } as QueryGetFeeCounterRequest;
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

  fromJSON(_: any): QueryGetFeeCounterRequest {
    const message = {
      ...baseQueryGetFeeCounterRequest,
    } as QueryGetFeeCounterRequest;
    return message;
  },

  toJSON(_: QueryGetFeeCounterRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetFeeCounterRequest>
  ): QueryGetFeeCounterRequest {
    const message = {
      ...baseQueryGetFeeCounterRequest,
    } as QueryGetFeeCounterRequest;
    return message;
  },
};

const baseQueryGetFeeCounterResponse: object = {};

export const QueryGetFeeCounterResponse = {
  encode(
    message: QueryGetFeeCounterResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.FeeCounter !== undefined) {
      FeeCounter.encode(message.FeeCounter, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetFeeCounterResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetFeeCounterResponse,
    } as QueryGetFeeCounterResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.FeeCounter = FeeCounter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeCounterResponse {
    const message = {
      ...baseQueryGetFeeCounterResponse,
    } as QueryGetFeeCounterResponse;
    if (object.FeeCounter !== undefined && object.FeeCounter !== null) {
      message.FeeCounter = FeeCounter.fromJSON(object.FeeCounter);
    } else {
      message.FeeCounter = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetFeeCounterResponse): unknown {
    const obj: any = {};
    message.FeeCounter !== undefined &&
      (obj.FeeCounter = message.FeeCounter
        ? FeeCounter.toJSON(message.FeeCounter)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFeeCounterResponse>
  ): QueryGetFeeCounterResponse {
    const message = {
      ...baseQueryGetFeeCounterResponse,
    } as QueryGetFeeCounterResponse;
    if (object.FeeCounter !== undefined && object.FeeCounter !== null) {
      message.FeeCounter = FeeCounter.fromPartial(object.FeeCounter);
    } else {
      message.FeeCounter = undefined;
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
  /** Queries a BetChart by index. */
  BetChart(request: QueryGetBetChartRequest): Promise<QueryGetBetChartResponse>;
  /** Queries a list of BetChart items. */
  BetChartAll(
    request: QueryAllBetChartRequest
  ): Promise<QueryAllBetChartResponse>;
  /** Queries a FeeCounter by index. */
  FeeCounter(
    request: QueryGetFeeCounterRequest
  ): Promise<QueryGetFeeCounterResponse>;
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

  BetChart(
    request: QueryGetBetChartRequest
  ): Promise<QueryGetBetChartResponse> {
    const data = QueryGetBetChartRequest.encode(request).finish();
    const promise = this.rpc.request(
      "orenshva.cosmoslottery.cosmoslottery.Query",
      "BetChart",
      data
    );
    return promise.then((data) =>
      QueryGetBetChartResponse.decode(new Reader(data))
    );
  }

  BetChartAll(
    request: QueryAllBetChartRequest
  ): Promise<QueryAllBetChartResponse> {
    const data = QueryAllBetChartRequest.encode(request).finish();
    const promise = this.rpc.request(
      "orenshva.cosmoslottery.cosmoslottery.Query",
      "BetChartAll",
      data
    );
    return promise.then((data) =>
      QueryAllBetChartResponse.decode(new Reader(data))
    );
  }

  FeeCounter(
    request: QueryGetFeeCounterRequest
  ): Promise<QueryGetFeeCounterResponse> {
    const data = QueryGetFeeCounterRequest.encode(request).finish();
    const promise = this.rpc.request(
      "orenshva.cosmoslottery.cosmoslottery.Query",
      "FeeCounter",
      data
    );
    return promise.then((data) =>
      QueryGetFeeCounterResponse.decode(new Reader(data))
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
