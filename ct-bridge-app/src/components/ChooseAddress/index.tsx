import Title from 'antd/lib/typography/Title';
import Button from 'src/components/Button';
import { useEffect } from 'react';
import { useWeb3React } from '@web3-react/core';
import { Col, Image, Input, message, Row, Select, Space, Tooltip } from 'antd';
import {
  getChainDataByChainId,
  requestChangeNetwork,
  useChainList
} from 'src/helpers/wallet';
import ArrowRightOutlined from '@ant-design/icons/ArrowRightOutlined';
import { bridgeAddressState } from 'src/state/bridge';
import { useRecoilState } from 'recoil';
import ChooseAddressStyle from './style';
import Web3ConnectButton from '../ConnectWalletButton/web3';
import chainData from './chainData.json';

const { Option } = Select;

type ChooseAccountPropType = {
  active: boolean;
  next: () => void;
};

const ChooseAccount: React.FC<ChooseAccountPropType> = ({ active, next }) => {
  const chainList = useChainList();
  const [bridgeAddress, setBridgeAddress] = useRecoilState(bridgeAddressState);
  const { account, chainId } = useWeb3React();

  const validate = (): boolean => {
    if (
      !bridgeAddress.targetChain ||
      !bridgeAddress.targetAddress ||
      !bridgeAddress.sourceChain ||
      !bridgeAddress.sourceAddress
    ) {
      message.error('请选择地址和链！');
      return false;
    }

    if (chainId !== bridgeAddress.sourceChain) {
      const chain = getChainDataByChainId(chainList, bridgeAddress.sourceChain);
      message.error(`请将链切换到 ${chain.name}`);
      return false;
    }
    return true;
  };

  const validateAndNext = () => {
    if (validate()) {
      next();
    }
  };

  function getLogo(val: string) {
    for (const chainDateElement of chainData) {
      if (chainDateElement['简称'] == val) {
        return chainDateElement['官网地址'] + chainDateElement['图标'];
      }
    }
  }

  useEffect(() => {
    if (bridgeAddress.sourceChain) {
      requestChangeNetwork(bridgeAddress.sourceChain);
    }
  }, [chainId, bridgeAddress.sourceChain]);

  useEffect(() => {
    if (account) {
      setBridgeAddress({
        ...bridgeAddress,
        sourceAddress: account || '',
        targetAddress: account || ''
      });
    }
  }, [account, chainId]);

  return (
    <ChooseAddressStyle>
      <div className='box'>
        {active ? (
          <>
            <Row gutter={{ sm: 20, md: 20, xs: 20, lg: 48 }}>
              <Col lg={11} md={24} xs={24} sm={24}>
                <Space direction='vertical'>
                  <Title level={5}>从</Title>
                  <Select
                    placeholder='选择源链'
                    value={bridgeAddress.sourceChain}
                    onChange={(value) => {
                      let targetChain = bridgeAddress.targetChain;
                      if (bridgeAddress.targetChain === value) {
                        targetChain = undefined;
                      }
                      setBridgeAddress({
                        ...bridgeAddress,
                        sourceChain: value,
                        targetChain
                      });
                    }}
                  >
                    {chainList.map((chainItem, index) => (
                      <Option value={chainItem.id} key={index}>
                        <Image
                          width={35}
                          src={getLogo(chainItem.value)}
                          alt={chainItem.value}
                          fallback={`/加载失败.svg`}
                        />{' '}
                        {chainItem.name}
                      </Option>
                    ))}
                  </Select>
                  <div />
                  <Title level={5}>源地址</Title>
                  <Web3ConnectButton block={true} />
                </Space>
              </Col>
              <Col lg={2} md={24} sm={24} xs={24} className='arrow-container'>
                <ArrowRightOutlined />
              </Col>
              <Col lg={11} md={24} sm={24} xs={24}>
                <Space direction='vertical'>
                  <Title level={5}>到</Title>
                  <Select
                    placeholder='选择目标链'
                    value={bridgeAddress.targetChain}
                    onChange={(value) => {
                      let sourceChain = bridgeAddress.sourceChain;
                      // console.log(bridgeAddress.sourceChain);
                      if (bridgeAddress.sourceChain === value) {
                        sourceChain = undefined;
                      }
                      setBridgeAddress({
                        ...bridgeAddress,
                        targetChain: value,
                        sourceChain
                      });
                    }}
                  >
                    {chainList.map((chainItem, index) => (
                      <Option value={chainItem.id} key={index}>
                        <Image
                          width={35}
                          src={getLogo(chainItem.value)}
                          alt={chainItem.value}
                          fallback={`/加载失败.svg`}
                        />{' '}
                        {chainItem.name}
                      </Option>
                    ))}
                  </Select>
                  <div />
                  <Title level={5}>目标地址</Title>
                  <Tooltip
                    placement='bottomLeft'
                    trigger={['hover']}
                    title={bridgeAddress.targetAddress}
                    overlayInnerStyle={{
                      width: 370,
                      borderRadius: 8
                    }}
                  >
                    <Input
                      placeholder='填写接收地址'
                      value={bridgeAddress.targetAddress}
                      onChange={(e) =>
                        setBridgeAddress({
                          ...bridgeAddress,
                          targetAddress: e.target.value
                        })
                      }
                    />
                  </Tooltip>
                </Space>
              </Col>
            </Row>
            <div className='next-container'>
              <Button type='primary' shape='round' onClick={validateAndNext}>
                下一步
              </Button>
            </div>
          </>
        ) : (
          <div>
            <p className='address-detail'>
              从: <span className='address'>{bridgeAddress.sourceAddress}</span>{' '}
              (
              {
                getChainDataByChainId(chainList, bridgeAddress.sourceChain!)
                  ?.name
              }
              )
            </p>
            <p className='address-detail'>
              到: <span className='address'>{bridgeAddress.targetAddress}</span>{' '}
              (
              {
                getChainDataByChainId(chainList, bridgeAddress.targetChain!)
                  ?.name
              }
              )
            </p>
          </div>
        )}
      </div>
    </ChooseAddressStyle>
  );
};

export default ChooseAccount;
