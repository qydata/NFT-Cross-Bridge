import { Button, Layout, PageHeader } from 'antd';
import PageLayoutStyle from './style';
import { useHistory } from 'react-router-dom';
import ConnectWalletButton from '../ConnectWalletButton';

const { Content } = Layout;

const PageLayout: React.FC<{ showConnectWallet?: boolean }> = ({
  children,
  showConnectWallet = true
}) => {
  const history = useHistory();
  return (
    <PageLayoutStyle>
      <Layout>
        <PageHeader
          title={
            <span className='app-name' onClick={() => history.push('/')}>
              草田链
            </span>
          }
          subTitle='NFT桥'
          avatar={{
            shape: 'square',
            src: '/ct.svg'
          }}
          extra={[
            <Button type='link' key='1' onClick={() => history.push('/')}>
              首页
            </Button>,
            <Button type='link' key='2' onClick={() => history.push('/bridge')}>
              桥
            </Button>,
            <Button type='link' key='3' onClick={() => history.push('/status')}>
              状态
            </Button>,
            <Button type='link' key='3' onClick={() => history.push('/my-nft')}>
              我的NFT
            </Button>,
            <ConnectWalletButton
              key='4'
              style={{ visibility: showConnectWallet ? 'visible' : 'hidden' }}
            />
          ]}
        ></PageHeader>
        <Content>
          <div className='container'>{children}</div>
        </Content>
      </Layout>
    </PageLayoutStyle>
  );
};

export default PageLayout;
