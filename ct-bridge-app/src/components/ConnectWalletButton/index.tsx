import { InjectedConnector } from '@web3-react/injected-connector';
import { useWeb3React } from '@web3-react/core';
import Button from 'src/components/Button';
import { formatAddress } from 'src/helpers/wallet';
import CloseCircleOutlined from '@ant-design/icons/CloseCircleOutlined';
import { SUPPORTED_CHAINS } from 'src/constants';
import Title from 'antd/lib/typography/Title';
import { useEffect, useState } from 'react';
import { setRedirectUrl } from 'src/helpers';
import { useRecoilState } from 'recoil';
import { DEFAULT_PROFILE, profileState } from 'src/state/profile';
import { ModalStyle } from './style';

export const injected = new InjectedConnector({
  supportedChainIds: SUPPORTED_CHAINS
});

const ProviderButton = ({ children, onClick, icon }: any) => (
  <Button shape='round' block onClick={onClick}>
    <img src={icon} alt={``} />
    {children}
  </Button>
);

const ConnectWalletButton = ({ block, style }: any) => {
  const { account, activate, deactivate } = useWeb3React();
  const [modalVisible, setModalVisible] = useState(false);
  const [profile, setProfile] = useRecoilState(profileState);

  useEffect(() => {
    if (account) {
      setProfile({
        ...DEFAULT_PROFILE,
        walletAddress: account
      });
    }
  }, [account]);

  const connect = async () => {
    try {
      await activate(injected);
    } catch (ex) {
      console.error(ex);
    }
    setModalVisible(false);
  };

  const disconnect = async () => {
    try {
      setRedirectUrl();
      if (account) {
        deactivate();
      } else {
        console.error('cannot logout');
      }
      setProfile(DEFAULT_PROFILE);
    } catch (ex) {
      console.error(ex);
    }
  };

  return (
    <>
      <ModalStyle
        visible={modalVisible}
        footer={null}
        closeIcon={<CloseCircleOutlined />}
        onCancel={() => setModalVisible(false)}
      >
        <Title level={4}>连接到钱包</Title>
        <ProviderButton icon={'/metamask.svg'} onClick={connect}>
          Metamask
        </ProviderButton>
        {/*<ProviderButton icon={'/ud.png'} onClick={loginWithUD}>*/}
        {/*  Login with Unstoppable*/}
        {/*</ProviderButton>*/}
      </ModalStyle>
      <Button
        type='primary'
        style={style}
        block={block}
        shape='round'
        onClick={
          profile.walletAddress ? disconnect : () => setModalVisible(true)
        }
      >
        {profile.walletAddress
          ? `断开连接 ${formatAddress(profile.walletAddress, 4)}`
          : '连接'}
      </Button>
    </>
  );
};

export default ConnectWalletButton;
