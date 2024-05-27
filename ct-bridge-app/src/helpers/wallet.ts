import { message } from 'antd';
import { useEffect } from 'react';
import { useRecoilState } from 'recoil';
import { getInfo } from 'src/apis/info';
import { ChainData } from 'src/interfaces/nft';
import { chainListState, infoState } from 'src/state/info';

export const useChainList = () => {
  const [info, setInfo] = useRecoilState(infoState);
  const [chainList, setChainList] = useRecoilState(chainListState);

  function capitalizeFirstLetter(str: string) {
    return str.charAt(0).toUpperCase() + str.slice(1);
  }

  useEffect(() => {
    if (!info) {
      getInfo().then((data) => {
        setInfo(data);
        const keys = Object.keys(data);
        const _chainKey = new Set<string>();
        const _chainList = [];
        for (const key of keys) {
          _chainKey.add(key.split('_')[0]);
        }

        // console.log(Array.from(_chainKey));
        const from: string[] = Array.from(_chainKey);
        for (const key1 of from) {
          _chainList.push({
            id: data[`${key1}_chain_id`],
            name: `${capitalizeFirstLetter(key1)} Chain`,
            value: key1,
            registerFee: 0.1,
            transferFee: 0.2,
            currency: key1.toUpperCase(),
            swapAgent20Address: data[`${key1}_erc_20_swap_agent`],
            swapAgent721Address: data[`${key1}_erc_721_swap_agent`],
            swapAgent1155Address: data[`${key1}_erc_1155_swap_agent`]
          });
        }

        setChainList(_chainList);
      });
    }
  }, [info]);

  return chainList;
};

export const getChainDataByChainId = (
  chainList: Array<ChainData>,
  chainId?: number
): ChainData => {
  return chainList.find((item) => item.id === Number(chainId))!;
};
export const AddEthereumChainParams: { [key: number]: any } = {
  27: {
    chainId: '0x1b',
    chainName: '草田链 Mainnet',
    nativeCurrency: {
      name: '草田链',
      symbol: 'CT',
      decimals: 18
    },
    rpcUrls: ['https://ctblock.cn/blockChain'],
    blockExplorerUrls: ['https://ctblock.cn/']
  },
  583: {
    chainId: '0x247',
    chainName: 'Coo Testnet',
    nativeCurrency: {
      name: 'Coo Testnet',
      symbol: 'Coo',
      decimals: 18
    },
    rpcUrls: ['http://182.43.26.165:7401'],
    blockExplorerUrls: ['http://182.43.26.165:4000/']
  }
};

export const requestChangeNetwork = async (
  chainId: number
): Promise<boolean> => {
  const provider = window.ethereum;
  if (provider && provider.request) {
    try {
      await provider.request({
        method: 'wallet_switchEthereumChain',
        params: [
          {
            chainId: `0x${chainId.toString(16)}`
          }
        ]
      });
      return true;
    } catch (error: any) {
      if (error.code === 4902) {
        // 这里添加网络
        message.error(`Please add chain id: ${chainId} chain to your network`);
        const params = AddEthereumChainParams[chainId];
        await provider.request({
          method: 'wallet_addEthereumChain',
          params: [params]
        });
      } else {
      }
      console.error('Failed to setup the network in Metamask:', error);
      return false;
    }
  } else {
    console.error(
      `Can't setup the BSC network on metamask because window.ethereum is undefined`
    );
    return false;
  }
};

export const formatAddress = (address: string, showLength: number): string => {
  const length = address.length;
  return `${address.substring(0, showLength)}...${address.substring(
    length - showLength,
    length
  )}`;
};
