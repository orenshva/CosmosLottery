/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "orenshva.cosmoslottery.cosmoslottery";

export interface BetChart {
  accountName: string;
  bet: number;
}

const baseBetChart: object = { accountName: "", bet: 0 };

export const BetChart = {
  encode(message: BetChart, writer: Writer = Writer.create()): Writer {
    if (message.accountName !== "") {
      writer.uint32(10).string(message.accountName);
    }
    if (message.bet !== 0) {
      writer.uint32(16).uint64(message.bet);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): BetChart {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBetChart } as BetChart;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountName = reader.string();
          break;
        case 2:
          message.bet = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BetChart {
    const message = { ...baseBetChart } as BetChart;
    if (object.accountName !== undefined && object.accountName !== null) {
      message.accountName = String(object.accountName);
    } else {
      message.accountName = "";
    }
    if (object.bet !== undefined && object.bet !== null) {
      message.bet = Number(object.bet);
    } else {
      message.bet = 0;
    }
    return message;
  },

  toJSON(message: BetChart): unknown {
    const obj: any = {};
    message.accountName !== undefined &&
      (obj.accountName = message.accountName);
    message.bet !== undefined && (obj.bet = message.bet);
    return obj;
  },

  fromPartial(object: DeepPartial<BetChart>): BetChart {
    const message = { ...baseBetChart } as BetChart;
    if (object.accountName !== undefined && object.accountName !== null) {
      message.accountName = object.accountName;
    } else {
      message.accountName = "";
    }
    if (object.bet !== undefined && object.bet !== null) {
      message.bet = object.bet;
    } else {
      message.bet = 0;
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
