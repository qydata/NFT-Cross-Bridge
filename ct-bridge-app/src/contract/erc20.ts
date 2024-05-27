import erc20AgentAbi from './abi/erc20Agent';
import erc20Agent from './abi/erc20Agent';
import erc20Abi from './abi/erc20';
import { getContract, getProvider } from '.';
import { message } from 'antd';
import { ethers } from 'ethers';
import { EVM } from 'evm';

class Contract20 {
  async getApprove(
    contractAddress: string,
    tokenAddress: string,
    userAddress: string | null | undefined
  ) {
    const contract = getContract(tokenAddress, erc20Abi);
    const approvedAmount = await contract.allowance(
      userAddress,
      contractAddress
    );
    return approvedAmount > 1;
  }

  async approve(contractAddress: string, tokenAddress: string) {
    return new Promise(async (reslove) => {
      const contract = getContract(tokenAddress, erc20Abi);
      try {
        const response = await contract.approve(
          contractAddress,
          '0xffffffffffffffffffffffffffffffffffffffff'
        );
        console.info(response.hash);
        contract.on('Approval', (sender, _spender, _value) => {
          const valueHex = ethers.BigNumber.from(
            _value.toString()
          ).toHexString();
          if (
            _spender.toLowerCase() === contractAddress.toLowerCase() &&
            '0xffffffffffffffffffffffffffffffffffffffff' === valueHex
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

  async name(tokenAddress: string) {
    const contract = getContract(tokenAddress, erc20Abi);
    try {
      const name = await contract.name();
      // console.log('name=>>>>', name);
      return name;
    } catch {
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
      const contract = getContract(agentAddress, erc20AgentAbi);

      async function decompile(contractAddr: string) {
        const code = await getProvider().getCode(contractAddr);
        const evm = new EVM(code);
        // console.log(evm.getFunctions()); /* Get functions */
        return evm.getFunctions();
      }

      try {
        const overrides = {
          value: ethers.utils.parseEther(fee.toString())
        };
        // console.log(tokenAddress, targetChainId);
        // console.log(agentAddress, overrides);

        // 这里因为合约中没有判断erc20, 所以这里进行判断
        const arrFun = await decompile(tokenAddress);

        const erc20Arr = [
          'name()',
          'symbol()',
          'totalSupply()',
          'balanceOf(address)',
          'transfer(address,uint256)',
          'transferFrom(address,address,uint256)',
          'approve(address,uint256)',
          'allowance(address,address)'
        ];
        if (
          arrFun.includes(erc20Arr[0]) === false ||
          arrFun.includes(erc20Arr[1]) === false ||
          arrFun.includes(erc20Arr[2]) === false ||
          arrFun.includes(erc20Arr[3]) === false ||
          arrFun.includes(erc20Arr[4]) === false ||
          arrFun.includes(erc20Arr[5]) === false ||
          arrFun.includes(erc20Arr[6]) === false ||
          arrFun.includes(erc20Arr[7]) === false
        ) {
          console.log(
            'ERC20SwapAgent::registerSwapPair:: token does not conform ERC20 standard '
          );
          reslove(false);
        } else {
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
        }
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
    amount: number,
    targetChainId: number,
    fee: number
  ): Promise<string> {
    const contract = getContract(agentAddress, erc20Agent);
    try {
      const overrides = {
        value: ethers.utils.parseEther(fee.toString())
      };
      const response = await contract.swap(
        tokenAddress,
        recipient,
        amount,
        targetChainId,
        overrides
      );
      return response.hash;
    } catch (error: any) {
      console.debug(error);
      message.error(error.data.message);
      return '';
    }
  }
}

export default new Contract20();
