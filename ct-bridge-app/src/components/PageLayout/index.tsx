import { Button, Layout, Menu, PageHeader, Drawer, Radio, Space } from 'antd';
import PageLayoutStyle from './style';
import { useHistory } from 'react-router-dom';

import React, { useState } from 'react';
import ConnectWalletButton from '../ConnectWalletButton';
import {
  AppstoreOutlined,
  ContainerOutlined,
  DesktopOutlined,
  MailOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  PieChartOutlined
} from '@ant-design/icons';

const { Content } = Layout;

// type MenuItem = Required<MenuProps>['items'][number];

function getItem(
  label: React.ReactNode,
  key: React.Key,
  icon?: React.ReactNode,
  children?: any[],
  type?: 'group'
) {
  return {
    key,
    icon,
    children,
    label,
    type
  };
}

const items = [
  getItem('Option 1', '1', <PieChartOutlined />),
  getItem('Option 2', '2', <DesktopOutlined />),
  getItem('Option 3', '3', <ContainerOutlined />),

  getItem('Navigation One', 'sub1', <MailOutlined />, [
    getItem('Option 5', '5'),
    getItem('Option 6', '6'),
    getItem('Option 7', '7'),
    getItem('Option 8', '8')
  ]),

  getItem('Navigation Two', 'sub2', <AppstoreOutlined />, [
    getItem('Option 9', '9'),
    getItem('Option 10', '10'),

    getItem('Submenu', 'sub3', null, [
      getItem('Option 11', '11'),
      getItem('Option 12', '12')
    ])
  ])
];

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
              {menuData.map((item) => (
                <Button
                  type='link'
                  key={item.key}
                  onClick={() => history.push(item.path)}
                >
                  {item.name}
                </Button>
              ))}
              <ConnectWalletButton
                key='5'
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
          {menuData.map((item) => (
            <div>
              <Button
                type='link'
                key={item.key}
                onClick={() => history.push(item.path)}
              >
                {item.name}
              </Button>
            </div>
          ))}
          <div>
            <ConnectWalletButton
              key='5'
              style={{ visibility: showConnectWallet ? 'visible' : 'hidden' }}
            />
          </div>
        </Drawer>
      </Layout>
    </PageLayoutStyle>
  );
};

export default PageLayout;
