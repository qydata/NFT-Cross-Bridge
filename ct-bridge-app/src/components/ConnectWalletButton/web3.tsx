import { InjectedConnector } from '@web3-react/injected-connector';
import { useWeb3React } from '@web3-react/core';
import Button from 'src/components/Button';
import { formatAddress } from 'src/helpers/wallet';
import { SUPPORTED_CHAINS } from 'src/constants';

export const injected = new InjectedConnector({
  supportedChainIds: SUPPORTED_CHAINS
});

const Web3ConnectButton = ({ block }: any) => {
  const { account, activate, deactivate } = useWeb3React();

  const connect = async () => {
    try {
      await activate(injected);
    } catch (ex) {
      console.error(ex);
    }
  };

  const disconnect = async () => {
    try {
      deactivate();
    } catch (ex) {
      console.error(ex);
    }
  };
  return (
    <Button
      type='primary'
      block={block}
      shape='round'
      onClick={account ? disconnect : connect}
    >
      {account ? `断开连接 ${formatAddress(account, 5)}` : '连接钱包'}
    </Button>
  );
};

export default Web3ConnectButton;
