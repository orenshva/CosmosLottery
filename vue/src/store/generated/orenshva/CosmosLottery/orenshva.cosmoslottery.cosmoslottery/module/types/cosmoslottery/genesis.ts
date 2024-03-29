/* eslint-disable */
import { Params } from "../cosmoslottery/params";
import { TxCounter } from "../cosmoslottery/tx_counter";
import { BetChart } from "../cosmoslottery/bet_chart";
import { FeeCounter } from "../cosmoslottery/fee_counter";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "orenshva.cosmoslottery.cosmoslottery";

/** GenesisState defines the cosmoslottery module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  txCounter: TxCounter | undefined;
  betChartList: BetChart[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  feeCounter: FeeCounter | undefined;
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.txCounter !== undefined) {
      TxCounter.encode(message.txCounter, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.betChartList) {
      BetChart.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.feeCounter !== undefined) {
      FeeCounter.encode(message.feeCounter, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.betChartList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.txCounter = TxCounter.decode(reader, reader.uint32());
          break;
        case 3:
          message.betChartList.push(BetChart.decode(reader, reader.uint32()));
          break;
        case 4:
          message.feeCounter = FeeCounter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.betChartList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.txCounter !== undefined && object.txCounter !== null) {
      message.txCounter = TxCounter.fromJSON(object.txCounter);
    } else {
      message.txCounter = undefined;
    }
    if (object.betChartList !== undefined && object.betChartList !== null) {
      for (const e of object.betChartList) {
        message.betChartList.push(BetChart.fromJSON(e));
      }
    }
    if (object.feeCounter !== undefined && object.feeCounter !== null) {
      message.feeCounter = FeeCounter.fromJSON(object.feeCounter);
    } else {
      message.feeCounter = undefined;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.txCounter !== undefined &&
      (obj.txCounter = message.txCounter
        ? TxCounter.toJSON(message.txCounter)
        : undefined);
    if (message.betChartList) {
      obj.betChartList = message.betChartList.map((e) =>
        e ? BetChart.toJSON(e) : undefined
      );
    } else {
      obj.betChartList = [];
    }
    message.feeCounter !== undefined &&
      (obj.feeCounter = message.feeCounter
        ? FeeCounter.toJSON(message.feeCounter)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.betChartList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.txCounter !== undefined && object.txCounter !== null) {
      message.txCounter = TxCounter.fromPartial(object.txCounter);
    } else {
      message.txCounter = undefined;
    }
    if (object.betChartList !== undefined && object.betChartList !== null) {
      for (const e of object.betChartList) {
        message.betChartList.push(BetChart.fromPartial(e));
      }
    }
    if (object.feeCounter !== undefined && object.feeCounter !== null) {
      message.feeCounter = FeeCounter.fromPartial(object.feeCounter);
    } else {
      message.feeCounter = undefined;
    }
    return message;
  },
};

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
