import { getDataFromTokenUri } from 'src/apis/nft';
import {
  CovalentData,
  ERROR_STATE,
  INFTParsedTokenAccount,
  NFT_STANDARD_OPTIONS,
  NFTStandard,
  NFTStandardData,
  SwapState,
  TransferStatus
} from 'src/interfaces/nft';

export const parseNFTData = async (
  walletAddress: string,
  tokenList: CovalentData[],
  standard: NFTStandard,
  chainId: number
): Promise<INFTParsedTokenAccount[]> => {
  const tokens = tokenList.reduce((arr, covalent) => {
    if (covalent) {
      arr.push({
        walletAddress,
        contractAddress: covalent.contract_address,
        amount: covalent.token_balance,
        decimals: covalent.contract_decimals,
        // uiAmount: Number(
        //   formatUnits(data.token_balance, covalent.contract_decimals)
        // ),
        uiAmount: Number(covalent.token_balance),
        // uiAmountString: formatUnits(
        //   data.token_balance,
        //   covalent.contract_decimals
        // ),
        uiAmountString: covalent.token_balance,
        symbol: covalent.contract_ticker_symbol,
        name: covalent.contract_name,
        logo: covalent.logo_url,
        tokenId: covalent.token_id,
        uri: covalent.token_url,
        animation_url: covalent.external_data?.animation_url,
        external_url: covalent.external_data?.external_url,
        image: covalent.external_data?.image,
        image_256: covalent.external_data?.image_256,
        nftName: covalent.external_data?.name,
        description: covalent.external_data?.description,
        standard,
        chainId
      });
    }
    return arr;
  }, [] as INFTParsedTokenAccount[]);
  return Promise.all(
    tokens.map(async (token) => {
      let image = token.image;
      if (!image) {
        const data = await getDataFromTokenUri(token.uri!);
        image = data.image;
      }
      return {
        ...token,
        image
      };
    })
  );
};

export const getNFTStandard = (standard: NFTStandard): NFTStandardData => {
  return NFT_STANDARD_OPTIONS.find((item) => item.value === standard)!;
};

export const getNFTStatusFromState = (state: SwapState) => {
  if (state === SwapState.FillTxConfirmed) {
    return TransferStatus.Done;
  } else if (ERROR_STATE.includes(state)) {
    return TransferStatus.Error;
  }
  return TransferStatus.InProgress;
};
