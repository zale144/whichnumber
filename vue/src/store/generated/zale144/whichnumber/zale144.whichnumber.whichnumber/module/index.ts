// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateParams } from "./types/whichnumber/tx";
import { MsgRevealNumber } from "./types/whichnumber/tx";
import { MsgNewGame } from "./types/whichnumber/tx";
import { MsgCommitNumber } from "./types/whichnumber/tx";


const types = [
  ["/zale144.whichnumber.whichnumber.MsgUpdateParams", MsgUpdateParams],
  ["/zale144.whichnumber.whichnumber.MsgRevealNumber", MsgRevealNumber],
  ["/zale144.whichnumber.whichnumber.MsgNewGame", MsgNewGame],
  ["/zale144.whichnumber.whichnumber.MsgCommitNumber", MsgCommitNumber],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgUpdateParams: (data: MsgUpdateParams): EncodeObject => ({ typeUrl: "/zale144.whichnumber.whichnumber.MsgUpdateParams", value: MsgUpdateParams.fromPartial( data ) }),
    msgRevealNumber: (data: MsgRevealNumber): EncodeObject => ({ typeUrl: "/zale144.whichnumber.whichnumber.MsgRevealNumber", value: MsgRevealNumber.fromPartial( data ) }),
    msgNewGame: (data: MsgNewGame): EncodeObject => ({ typeUrl: "/zale144.whichnumber.whichnumber.MsgNewGame", value: MsgNewGame.fromPartial( data ) }),
    msgCommitNumber: (data: MsgCommitNumber): EncodeObject => ({ typeUrl: "/zale144.whichnumber.whichnumber.MsgCommitNumber", value: MsgCommitNumber.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
