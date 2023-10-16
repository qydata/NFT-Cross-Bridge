import erc721AgentAbi from './abi/erc721Agent';
import erc721Abi from './abi/erc721';
import { getContract } from '.';
import { message } from 'antd';
import { ethers } from 'ethers';

class Contract721 {
  async getApprove(
    contractAddress: string,
    tokenAddress: string,
    tokenId: string
  ) {
    const contract = getContract(tokenAddress, erc721Abi);
    const approvedContract = await contract.getApproved(tokenId);
    return approvedContract.toLowerCase() === contractAddress.toLowerCase();
  }

  async approve(
    contractAddress: string,
    tokenAddress: string,
    tokenId: string
  ) {
    return new Promise(async (reslove) => {
      const contract = getContract(tokenAddress, erc721Abi);
      try {
        const response = await contract.approve(contractAddress, tokenId);
        console.info(response.hash);
        contract.on('Approval', (owner, approved, _tokenId) => {
          const tokenIdHex = ethers.BigNumber.from(
            _tokenId.toString()
          ).toHexString();
          if (
            approved.toLowerCase() === contractAddress.toLowerCase() &&
            tokenId === tokenIdHex
          ) {
            contract.removeAllListeners('Approval');
            reslove(true);
          }
        });
      } catch {
        contract.removeAllListeners('Approval');
        reslove(false);
      }
    });
  }

  async getTokenUri(tokenAddress: string, tokenId: string) {
    const contract = getContract(tokenAddress, erc721Abi);
    try {
      // console.log(tokenId);
      // console.log(typeof tokenId);
      const tokenIdHex = ethers.BigNumber.from(
        tokenId.toString()
      ).toHexString();
      // console.log(tokenIdHex);
      const tokenUri = await contract.tokenURI(tokenIdHex);
      return tokenUri;
    } catch (e) {
      console.log(e);
      message.error('Cannot find token!');
      return null;
    }
  }

  async registerToken(
    agentAddress: string,
    tokenAddress: string,
    targetChainId: number,
    fee: number
  ): Promise<boolean> {
    return new Promise(async (reslove) => {
      const contract = getContract(agentAddress, erc721AgentAbi);
      try {
        const overrides = {
          value: ethers.utils.parseEther(fee.toString())
        };
        // console.log(tokenAddress, targetChainId);
        // console.log(agentAddress, overrides);
        const response = await contract.registerSwapPair(
          tokenAddress,
          targetChainId,
          overrides
        );
        contract.on(
          'SwapPairRegister',
          (
            sponsor,
            _tokenAddress,
            tokenName,
            tokenSymbol,
            toChainId,
            feeAmount
          ) => {
            console.info(
              sponsor,
              _tokenAddress,
              tokenName,
              tokenSymbol,
              toChainId,
              feeAmount
            );
            reslove(true);
            contract.removeAllListeners('SwapPairRegister');
          }
        );
        console.info('tx: ', response.hash);
      } catch {
        contract.removeAllListeners('SwapPairRegister');
        reslove(false);
      }
    });
  }

  async transferToken(
    agentAddress: string,
    tokenAddress: string,
    recipient: string,
    tokenId: string,
    targetChainId: number,
    fee: number
  ): Promise<string> {
    const contract = getContract(agentAddress, erc721AgentAbi);
    try {
      const overrides = {
        value: ethers.utils.parseEther(fee.toString())
      };
      const response = await contract.swap(
        tokenAddress,
        recipient,
        tokenId,
        targetChainId,
        overrides
      );
      console.debug('tx: ', response.hash);
      return response.hash;
    } catch (error) {
      console.error(error);
      return '';
    }
  }
}

export default new Contract721();
