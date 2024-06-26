/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PayableOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../common";

export interface ERC721SwapAgentInterface extends utils.Interface {
  functions: {
    "createSwapPair(bytes32,address,uint256,string,string,string)": FunctionFragment;
    "fill(bytes32,address,address,uint256,uint256,string)": FunctionFragment;
    "filledSwap(bytes32)": FunctionFragment;
    "initialize()": FunctionFragment;
    "onERC721Received(address,address,uint256,bytes)": FunctionFragment;
    "owner()": FunctionFragment;
    "registerSwapPair(address,uint256)": FunctionFragment;
    "registeredToken(uint256,address)": FunctionFragment;
    "renounceOwnership()": FunctionFragment;
    "swap(address,address,uint256,uint256)": FunctionFragment;
    "swapMappingIncoming(uint256,address)": FunctionFragment;
    "swapMappingOutgoing(uint256,address)": FunctionFragment;
    "transferOwnership(address)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "createSwapPair"
      | "fill"
      | "filledSwap"
      | "initialize"
      | "onERC721Received"
      | "owner"
      | "registerSwapPair"
      | "registeredToken"
      | "renounceOwnership"
      | "swap"
      | "swapMappingIncoming"
      | "swapMappingOutgoing"
      | "transferOwnership"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "createSwapPair",
    values: [
      PromiseOrValue<BytesLike>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<string>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "fill",
    values: [
      PromiseOrValue<BytesLike>,
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "filledSwap",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "initialize",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "onERC721Received",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(functionFragment: "owner", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "registerSwapPair",
    values: [PromiseOrValue<string>, PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "registeredToken",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "renounceOwnership",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "swap",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "swapMappingIncoming",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "swapMappingOutgoing",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "transferOwnership",
    values: [PromiseOrValue<string>]
  ): string;

  decodeFunctionResult(
    functionFragment: "createSwapPair",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "fill", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "filledSwap", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "initialize", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "onERC721Received",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "owner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "registerSwapPair",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "registeredToken",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "renounceOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "swap", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "swapMappingIncoming",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "swapMappingOutgoing",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "transferOwnership",
    data: BytesLike
  ): Result;

  events: {
    "BackwardSwapFilled(bytes32,address,address,uint256,uint256)": EventFragment;
    "BackwardSwapStarted(address,address,address,uint256,uint256,uint256)": EventFragment;
    "OwnershipTransferred(address,address)": EventFragment;
    "SwapFilled(bytes32,address,address,address,uint256,uint256)": EventFragment;
    "SwapPairCreated(bytes32,address,address,uint256,string,string)": EventFragment;
    "SwapPairRegister(address,address,string,string,uint256,uint256)": EventFragment;
    "SwapStarted(address,address,address,uint256,uint256,uint256)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "BackwardSwapFilled"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "BackwardSwapStarted"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "OwnershipTransferred"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "SwapFilled"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "SwapPairCreated"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "SwapPairRegister"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "SwapStarted"): EventFragment;
}

export interface BackwardSwapFilledEventObject {
  swapTxHash: string;
  tokenAddr: string;
  recipient: string;
  fromChainId: BigNumber;
  tokenId: BigNumber;
}
export type BackwardSwapFilledEvent = TypedEvent<
  [string, string, string, BigNumber, BigNumber],
  BackwardSwapFilledEventObject
>;

export type BackwardSwapFilledEventFilter =
  TypedEventFilter<BackwardSwapFilledEvent>;

export interface BackwardSwapStartedEventObject {
  mirroredTokenAddr: string;
  sender: string;
  recipient: string;
  dstChainId: BigNumber;
  tokenId: BigNumber;
  feeAmount: BigNumber;
}
export type BackwardSwapStartedEvent = TypedEvent<
  [string, string, string, BigNumber, BigNumber, BigNumber],
  BackwardSwapStartedEventObject
>;

export type BackwardSwapStartedEventFilter =
  TypedEventFilter<BackwardSwapStartedEvent>;

export interface OwnershipTransferredEventObject {
  previousOwner: string;
  newOwner: string;
}
export type OwnershipTransferredEvent = TypedEvent<
  [string, string],
  OwnershipTransferredEventObject
>;

export type OwnershipTransferredEventFilter =
  TypedEventFilter<OwnershipTransferredEvent>;

export interface SwapFilledEventObject {
  swapTxHash: string;
  fromTokenAddr: string;
  recipient: string;
  mirroredTokenAddr: string;
  fromChainId: BigNumber;
  tokenId: BigNumber;
}
export type SwapFilledEvent = TypedEvent<
  [string, string, string, string, BigNumber, BigNumber],
  SwapFilledEventObject
>;

export type SwapFilledEventFilter = TypedEventFilter<SwapFilledEvent>;

export interface SwapPairCreatedEventObject {
  registerTxHash: string;
  fromTokenAddr: string;
  mirroredTokenAddr: string;
  fromChainId: BigNumber;
  tokenSymbol: string;
  tokenName: string;
}
export type SwapPairCreatedEvent = TypedEvent<
  [string, string, string, BigNumber, string, string],
  SwapPairCreatedEventObject
>;

export type SwapPairCreatedEventFilter = TypedEventFilter<SwapPairCreatedEvent>;

export interface SwapPairRegisterEventObject {
  sponsor: string;
  tokenAddress: string;
  tokenName: string;
  tokenSymbol: string;
  toChainId: BigNumber;
  feeAmount: BigNumber;
}
export type SwapPairRegisterEvent = TypedEvent<
  [string, string, string, string, BigNumber, BigNumber],
  SwapPairRegisterEventObject
>;

export type SwapPairRegisterEventFilter =
  TypedEventFilter<SwapPairRegisterEvent>;

export interface SwapStartedEventObject {
  tokenAddr: string;
  sender: string;
  recipient: string;
  dstChainId: BigNumber;
  tokenId: BigNumber;
  feeAmount: BigNumber;
}
export type SwapStartedEvent = TypedEvent<
  [string, string, string, BigNumber, BigNumber, BigNumber],
  SwapStartedEventObject
>;

export type SwapStartedEventFilter = TypedEventFilter<SwapStartedEvent>;

export interface ERC721SwapAgent extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: ERC721SwapAgentInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    createSwapPair(
      registerTxHash: PromiseOrValue<BytesLike>,
      fromTokenAddr: PromiseOrValue<string>,
      fromChainId: PromiseOrValue<BigNumberish>,
      baseURI_: PromiseOrValue<string>,
      tokenName: PromiseOrValue<string>,
      tokenSymbol: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    fill(
      swapTxHash: PromiseOrValue<BytesLike>,
      fromTokenAddr: PromiseOrValue<string>,
      recipient: PromiseOrValue<string>,
      fromChainId: PromiseOrValue<BigNumberish>,
      tokenId: PromiseOrValue<BigNumberish>,
      tokenURI: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    filledSwap(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    initialize(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    onERC721Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    owner(overrides?: CallOverrides): Promise<[string]>;

    registerSwapPair(
      tokenAddr: PromiseOrValue<string>,
      chainId: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    registeredToken(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    swap(
      tokenAddr: PromiseOrValue<string>,
      recipient: PromiseOrValue<string>,
      tokenId: PromiseOrValue<BigNumberish>,
      dstChainId: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    swapMappingIncoming(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    swapMappingOutgoing(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  createSwapPair(
    registerTxHash: PromiseOrValue<BytesLike>,
    fromTokenAddr: PromiseOrValue<string>,
    fromChainId: PromiseOrValue<BigNumberish>,
    baseURI_: PromiseOrValue<string>,
    tokenName: PromiseOrValue<string>,
    tokenSymbol: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  fill(
    swapTxHash: PromiseOrValue<BytesLike>,
    fromTokenAddr: PromiseOrValue<string>,
    recipient: PromiseOrValue<string>,
    fromChainId: PromiseOrValue<BigNumberish>,
    tokenId: PromiseOrValue<BigNumberish>,
    tokenURI: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  filledSwap(
    arg0: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  initialize(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  onERC721Received(
    arg0: PromiseOrValue<string>,
    arg1: PromiseOrValue<string>,
    arg2: PromiseOrValue<BigNumberish>,
    arg3: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  owner(overrides?: CallOverrides): Promise<string>;

  registerSwapPair(
    tokenAddr: PromiseOrValue<string>,
    chainId: PromiseOrValue<BigNumberish>,
    overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  registeredToken(
    arg0: PromiseOrValue<BigNumberish>,
    arg1: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  renounceOwnership(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  swap(
    tokenAddr: PromiseOrValue<string>,
    recipient: PromiseOrValue<string>,
    tokenId: PromiseOrValue<BigNumberish>,
    dstChainId: PromiseOrValue<BigNumberish>,
    overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  swapMappingIncoming(
    arg0: PromiseOrValue<BigNumberish>,
    arg1: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<string>;

  swapMappingOutgoing(
    arg0: PromiseOrValue<BigNumberish>,
    arg1: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<string>;

  transferOwnership(
    newOwner: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    createSwapPair(
      registerTxHash: PromiseOrValue<BytesLike>,
      fromTokenAddr: PromiseOrValue<string>,
      fromChainId: PromiseOrValue<BigNumberish>,
      baseURI_: PromiseOrValue<string>,
      tokenName: PromiseOrValue<string>,
      tokenSymbol: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    fill(
      swapTxHash: PromiseOrValue<BytesLike>,
      fromTokenAddr: PromiseOrValue<string>,
      recipient: PromiseOrValue<string>,
      fromChainId: PromiseOrValue<BigNumberish>,
      tokenId: PromiseOrValue<BigNumberish>,
      tokenURI: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    filledSwap(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    initialize(overrides?: CallOverrides): Promise<void>;

    onERC721Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    owner(overrides?: CallOverrides): Promise<string>;

    registerSwapPair(
      tokenAddr: PromiseOrValue<string>,
      chainId: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    registeredToken(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    renounceOwnership(overrides?: CallOverrides): Promise<void>;

    swap(
      tokenAddr: PromiseOrValue<string>,
      recipient: PromiseOrValue<string>,
      tokenId: PromiseOrValue<BigNumberish>,
      dstChainId: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    swapMappingIncoming(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<string>;

    swapMappingOutgoing(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<string>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "BackwardSwapFilled(bytes32,address,address,uint256,uint256)"(
      swapTxHash?: PromiseOrValue<BytesLike> | null,
      tokenAddr?: PromiseOrValue<string> | null,
      recipient?: PromiseOrValue<string> | null,
      fromChainId?: null,
      tokenId?: null
    ): BackwardSwapFilledEventFilter;
    BackwardSwapFilled(
      swapTxHash?: PromiseOrValue<BytesLike> | null,
      tokenAddr?: PromiseOrValue<string> | null,
      recipient?: PromiseOrValue<string> | null,
      fromChainId?: null,
      tokenId?: null
    ): BackwardSwapFilledEventFilter;

    "BackwardSwapStarted(address,address,address,uint256,uint256,uint256)"(
      mirroredTokenAddr?: PromiseOrValue<string> | null,
      sender?: PromiseOrValue<string> | null,
      recipient?: PromiseOrValue<string> | null,
      dstChainId?: null,
      tokenId?: null,
      feeAmount?: null
    ): BackwardSwapStartedEventFilter;
    BackwardSwapStarted(
      mirroredTokenAddr?: PromiseOrValue<string> | null,
      sender?: PromiseOrValue<string> | null,
      recipient?: PromiseOrValue<string> | null,
      dstChainId?: null,
      tokenId?: null,
      feeAmount?: null
    ): BackwardSwapStartedEventFilter;

    "OwnershipTransferred(address,address)"(
      previousOwner?: PromiseOrValue<string> | null,
      newOwner?: PromiseOrValue<string> | null
    ): OwnershipTransferredEventFilter;
    OwnershipTransferred(
      previousOwner?: PromiseOrValue<string> | null,
      newOwner?: PromiseOrValue<string> | null
    ): OwnershipTransferredEventFilter;

    "SwapFilled(bytes32,address,address,address,uint256,uint256)"(
      swapTxHash?: PromiseOrValue<BytesLike> | null,
      fromTokenAddr?: PromiseOrValue<string> | null,
      recipient?: PromiseOrValue<string> | null,
      mirroredTokenAddr?: null,
      fromChainId?: null,
      tokenId?: null
    ): SwapFilledEventFilter;
    SwapFilled(
      swapTxHash?: PromiseOrValue<BytesLike> | null,
      fromTokenAddr?: PromiseOrValue<string> | null,
      recipient?: PromiseOrValue<string> | null,
      mirroredTokenAddr?: null,
      fromChainId?: null,
      tokenId?: null
    ): SwapFilledEventFilter;

    "SwapPairCreated(bytes32,address,address,uint256,string,string)"(
      registerTxHash?: PromiseOrValue<BytesLike> | null,
      fromTokenAddr?: PromiseOrValue<string> | null,
      mirroredTokenAddr?: PromiseOrValue<string> | null,
      fromChainId?: null,
      tokenSymbol?: null,
      tokenName?: null
    ): SwapPairCreatedEventFilter;
    SwapPairCreated(
      registerTxHash?: PromiseOrValue<BytesLike> | null,
      fromTokenAddr?: PromiseOrValue<string> | null,
      mirroredTokenAddr?: PromiseOrValue<string> | null,
      fromChainId?: null,
      tokenSymbol?: null,
      tokenName?: null
    ): SwapPairCreatedEventFilter;

    "SwapPairRegister(address,address,string,string,uint256,uint256)"(
      sponsor?: PromiseOrValue<string> | null,
      tokenAddress?: PromiseOrValue<string> | null,
      tokenName?: null,
      tokenSymbol?: null,
      toChainId?: null,
      feeAmount?: null
    ): SwapPairRegisterEventFilter;
    SwapPairRegister(
      sponsor?: PromiseOrValue<string> | null,
      tokenAddress?: PromiseOrValue<string> | null,
      tokenName?: null,
      tokenSymbol?: null,
      toChainId?: null,
      feeAmount?: null
    ): SwapPairRegisterEventFilter;

    "SwapStarted(address,address,address,uint256,uint256,uint256)"(
      tokenAddr?: PromiseOrValue<string> | null,
      sender?: PromiseOrValue<string> | null,
      recipient?: PromiseOrValue<string> | null,
      dstChainId?: null,
      tokenId?: null,
      feeAmount?: null
    ): SwapStartedEventFilter;
    SwapStarted(
      tokenAddr?: PromiseOrValue<string> | null,
      sender?: PromiseOrValue<string> | null,
      recipient?: PromiseOrValue<string> | null,
      dstChainId?: null,
      tokenId?: null,
      feeAmount?: null
    ): SwapStartedEventFilter;
  };

  estimateGas: {
    createSwapPair(
      registerTxHash: PromiseOrValue<BytesLike>,
      fromTokenAddr: PromiseOrValue<string>,
      fromChainId: PromiseOrValue<BigNumberish>,
      baseURI_: PromiseOrValue<string>,
      tokenName: PromiseOrValue<string>,
      tokenSymbol: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    fill(
      swapTxHash: PromiseOrValue<BytesLike>,
      fromTokenAddr: PromiseOrValue<string>,
      recipient: PromiseOrValue<string>,
      fromChainId: PromiseOrValue<BigNumberish>,
      tokenId: PromiseOrValue<BigNumberish>,
      tokenURI: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    filledSwap(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    initialize(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    onERC721Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    owner(overrides?: CallOverrides): Promise<BigNumber>;

    registerSwapPair(
      tokenAddr: PromiseOrValue<string>,
      chainId: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    registeredToken(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    swap(
      tokenAddr: PromiseOrValue<string>,
      recipient: PromiseOrValue<string>,
      tokenId: PromiseOrValue<BigNumberish>,
      dstChainId: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    swapMappingIncoming(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    swapMappingOutgoing(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    createSwapPair(
      registerTxHash: PromiseOrValue<BytesLike>,
      fromTokenAddr: PromiseOrValue<string>,
      fromChainId: PromiseOrValue<BigNumberish>,
      baseURI_: PromiseOrValue<string>,
      tokenName: PromiseOrValue<string>,
      tokenSymbol: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    fill(
      swapTxHash: PromiseOrValue<BytesLike>,
      fromTokenAddr: PromiseOrValue<string>,
      recipient: PromiseOrValue<string>,
      fromChainId: PromiseOrValue<BigNumberish>,
      tokenId: PromiseOrValue<BigNumberish>,
      tokenURI: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    filledSwap(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    initialize(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    onERC721Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    owner(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    registerSwapPair(
      tokenAddr: PromiseOrValue<string>,
      chainId: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    registeredToken(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    swap(
      tokenAddr: PromiseOrValue<string>,
      recipient: PromiseOrValue<string>,
      tokenId: PromiseOrValue<BigNumberish>,
      dstChainId: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    swapMappingIncoming(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    swapMappingOutgoing(
      arg0: PromiseOrValue<BigNumberish>,
      arg1: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
