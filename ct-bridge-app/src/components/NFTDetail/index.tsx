import Title from 'antd/lib/typography/Title';
import Button from 'src/components/Button';
import Row from 'antd/lib/row';
import Col from 'antd/lib/col';
import { Image } from 'antd';
import CloseCircleOutlined from '@ant-design/icons/CloseCircleOutlined';
import ExclamationCircleOutlined from '@ant-design/icons/ExclamationCircleOutlined';
import { bridgeAddressState, nftState } from 'src/state/bridge';
import { useRecoilState, useRecoilValue } from 'recoil';
import { getChainDataByChainId, useChainList } from 'src/helpers/wallet';
import { getNFTStandard } from 'src/helpers/nft';
import { useEffect, useState } from 'react';
import {
  getIsNft1155Registered,
  getIsNft20Registered,
  getIsNft721Registered
} from 'src/apis/nft';
import contractErc20 from 'src/contract/erc20';
import contractErc721 from 'src/contract/erc721';
import contractErc1155 from 'src/contract/erc1155';
import Alert from 'antd/lib/alert';
import { message } from 'antd';
import { NFTStandard } from 'src/interfaces/nft';
import { EMPTY_NFT_DATA } from 'src/constants/nft';
import Input from 'antd/lib/input';
import NFTDetailStyle from './style';
import { ethers } from 'ethers';

import { useWeb3React } from '@web3-react/core';

enum NftStatus {
  Loading = 'loading',
  NotApprove = 'not_approve',
  NotRegister = 'not_register',
  Ready = 'Ready'
}

type NFTDetailPropType = {
  disabled: boolean;
  next: () => void;
};

const NFTDetail: React.FC<NFTDetailPropType> = ({
  disabled,
  next
}: {
  disabled: boolean;
  next: () => void;
}) => {
  const [nft, setNft] = useRecoilState(nftState);
  const bridgeAddress = useRecoilValue(bridgeAddressState);
  const [nftStatus, setNftStatus] = useState<NftStatus>(NftStatus.Loading);
  const chainList = useChainList();
  const { account } = useWeb3React();

  const {
    name,
    image,
    contractAddress: tokenAddress,
    tokenId,
    chainId,
    standard,
    walletAddress,
    uiAmount
  } = nft;
  const chainData = getChainDataByChainId(chainList, chainId);
  const validate = (): boolean => {
    if (standard === NFTStandard.ERC_1155 && !uiAmount) {
      message.error('Please fill token amount!');
      return false;
    }
    return true;
  };

  const validateAndNext = () => {
    if (validate()) {
      next();
    }
  };

  const checkIsApproved = async () => {
    if (standard === NFTStandard.ERC_20) {
      return contractErc20.getApprove(
        chainData.swapAgent20Address,
        tokenAddress,
        account
      );
    } else if (standard === NFTStandard.ERC_721) {
      const tokenIdHex = ethers.BigNumber.from(
        tokenId!.toString()
      ).toHexString();
      return contractErc721.getApprove(
        chainData.swapAgent721Address,
        tokenAddress,
        tokenIdHex
      );
    } else if (standard === NFTStandard.ERC_1155) {
      return contractErc1155.getApprove(
        tokenAddress,
        walletAddress,
        chainData.swapAgent1155Address
      );
    }
    return false;
  };

  async function checkNftStatus() {
    const isApproved = await checkIsApproved();
    if (isApproved) {
      let isRegister;
      if (standard === NFTStandard.ERC_20) {
        isRegister = await getIsNft20Registered(
          bridgeAddress.sourceChain!,
          bridgeAddress.targetChain!,
          tokenAddress
        );
      } else if (standard === NFTStandard.ERC_721) {
        isRegister = await getIsNft721Registered(
          bridgeAddress.sourceChain!,
          bridgeAddress.targetChain!,
          tokenAddress
        );
      } else if (standard === NFTStandard.ERC_1155) {
        isRegister = await getIsNft1155Registered(
          bridgeAddress.sourceChain!,
          bridgeAddress.targetChain!,
          tokenAddress
        );
      }

      if (isRegister) {
        setNftStatus(NftStatus.Ready);
      } else {
        setNftStatus(NftStatus.NotRegister);
      }
    } else {
      setNftStatus(NftStatus.NotApprove);
    }
  }

  const approve = async () => {
    setNftStatus(NftStatus.Loading);
    let isApproved;

    if (standard === NFTStandard.ERC_20) {
      isApproved = await contractErc20.approve(
        chainData.swapAgent20Address,
        tokenAddress
      );
    } else if (standard === NFTStandard.ERC_721) {
      const tokenIdHex = ethers.BigNumber.from(
        tokenId!.toString()
      ).toHexString();
      isApproved = await contractErc721.approve(
        chainData.swapAgent721Address,
        tokenAddress,
        tokenIdHex
      );
    } else if (standard === NFTStandard.ERC_1155) {
      isApproved = await contractErc1155.approve(
        chainData.swapAgent1155Address,
        tokenAddress
      );
    }
    if (isApproved) {
      message.success('成功批准 NFT');
      checkNftStatus();
    } else {
      setNftStatus(NftStatus.NotApprove);
    }
  };

  const register = async () => {
    setNftStatus(NftStatus.Loading);
    let isRegistered;
    if (standard === NFTStandard.ERC_20) {
      isRegistered = await contractErc20.registerToken(
        chainData.swapAgent20Address,
        tokenAddress,
        bridgeAddress.targetChain!,
        chainData.registerFee
      );
    } else if (standard === NFTStandard.ERC_721) {
      isRegistered = await contractErc721.registerToken(
        chainData.swapAgent721Address,
        tokenAddress,
        bridgeAddress.targetChain!,
        chainData.registerFee
      );
    } else if (standard === NFTStandard.ERC_1155) {
      isRegistered = await contractErc1155.registerToken(
        chainData.swapAgent1155Address,
        tokenAddress,
        bridgeAddress.targetChain!,
        chainData.registerFee
      );
    }
    if (isRegistered) {
      setNftStatus(NftStatus.Ready);
      message.success('成功注册 NFT');
    } else {
      // TODO 这里可能是已经注册失败得情况
      setNftStatus(NftStatus.NotRegister);
      checkNftStatus();
    }
  };

  const setAmount = (value: number) => {
    setNft({
      ...nft,
      uiAmount: value
    });
  };

  useEffect(() => {
    if (nft) {
      setNftStatus(NftStatus.Loading);
      checkNftStatus();
    }
  }, [nft]);

  if (!nft) return null;

  return (
    // @ts-ignore
    <NFTDetailStyle disabled={disabled}>
      <div className='nft-detail-headar'>
        <Title level={4}>NFT 详情</Title>
        <Button
          type='text'
          shape='circle'
          disabled={disabled}
          onClick={() => setNft(EMPTY_NFT_DATA)}
        >
          <CloseCircleOutlined />
        </Button>
      </div>
      <Row gutter={24}>
        <Col lg={12} md={24} sm={24} xs={24}>
          <Image
            src={image!}
            alt={name!}
            width='100%'
            fallback={`/加载失败.svg`}
          />
        </Col>
        <Col lg={12} md={24} sm={24} xs={24} className='detail-container'>
          <Title level={5}>名称</Title>
          <p className='detail'>{name}</p>
          <Title level={5}>NFT地址</Title>
          <p className='detail'>{tokenAddress}</p>
          {standard != NFTStandard.ERC_20 && (
            <>
              <Title level={5}>NFT ID</Title>
              <p className='detail'>{tokenId?.toString()}</p>
            </>
          )}
          {standard != NFTStandard.ERC_721 && (
            <>
              <Title level={5}>NFT数量</Title>
              <Input
                type='text'
                disabled={disabled}
                className='token-amount-input'
                value={uiAmount}
                onChange={(e) => setAmount(Number(e.target.value))}
                suffix={standard != NFTStandard.ERC_20 ? '' : 'WEI'}
              />
            </>
          )}
          <Title level={5}>NFT 标准</Title>
          <p className='detail'>{getNFTStandard(standard).label}</p>
          <Title level={5}>链</Title>
          <p className='detail'>{chainData.name}</p>
        </Col>
      </Row>
      <div className='action-container'>
        {nftStatus === NftStatus.Loading && (
          <Button
            type='primary'
            shape='round'
            loading
            hidden={disabled}
          ></Button>
        )}
        {nftStatus === NftStatus.NotApprove && (
          <Button
            type='primary'
            shape='round'
            onClick={approve}
            hidden={disabled}
          >
            授权 NFT
          </Button>
        )}
        {nftStatus === NftStatus.NotRegister && (
          <>
            <Button
              type='primary'
              shape='round'
              onClick={register}
              hidden={disabled}
            >
              注册NFT
            </Button>
            <p className='register-info'>
              <Alert
                type='warning'
                message={
                  <>
                    <ExclamationCircleOutlined /> 您必须支付注册费{' '}
                    {chainData.registerFee} {chainData.currency}
                  </>
                }
              />
            </p>
          </>
        )}

        {nftStatus === NftStatus.Ready && (
          <>
            <Button
              type='primary'
              shape='round'
              onClick={validateAndNext}
              hidden={disabled}
            >
              下一步
            </Button>
          </>
        )}
      </div>
    </NFTDetailStyle>
  );
};

export default NFTDetail;
