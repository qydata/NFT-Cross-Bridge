import { Button, Drawer, Layout, PageHeader } from 'antd';
import PageLayoutStyle from './style';
import { useHistory } from 'react-router-dom';

import React, { useState } from 'react';
import ConnectWalletButton from '../ConnectWalletButton';
import { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons';

const { Content } = Layout;

const PageLayout: React.FC<{ showConnectWallet?: boolean }> = ({
  children,
  showConnectWallet = true
}) => {
  const history = useHistory();
  const [open, setOpen] = useState(false);
  const showDrawer = () => {
    setOpen(true);
  };

  const onClose = () => {
    setOpen(false);
  };
  const menuData = [
    { name: '首页', path: '/', key: '1' },
    { name: '桥', path: '/bridge', key: '2' },
    { name: '状态', path: '/status', key: '3' },
    { name: '我的NFT', path: '/my-nft', key: '4' }
  ];
  // @ts-ignore
  // @ts-ignore
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
            <div style={{}} className='mobileNav'>
              <Button type='primary' onClick={showDrawer}>
                {open ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
              </Button>
            </div>,
            <div className='pcNav'>
              {menuData.map((item, index) => (
                <Button
                  type='link'
                  key={index}
                  onClick={() => history.push(item.path)}
                >
                  {item.name}
                </Button>
              ))}
              <ConnectWalletButton
                style={{ visibility: showConnectWallet ? 'visible' : 'hidden' }}
              />
            </div>
          ]}
        />
        <Content>
          <div className='container'>{children}</div>
        </Content>
        <Drawer
          title='选择一个菜单'
          placement='right'
          width={'75%'}
          onClose={onClose}
          open={open}
        >
          {menuData.map((item, index) => (
            <div>
              <Button
                type='link'
                key={index}
                onClick={() => history.push(item.path)}
              >
                {item.name}
              </Button>
            </div>
          ))}
          <div>
            <ConnectWalletButton
              style={{ visibility: showConnectWallet ? 'visible' : 'hidden' }}
            />
          </div>
        </Drawer>
      </Layout>
    </PageLayoutStyle>
  );
};

export default PageLayout;
