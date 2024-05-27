import { useEffect, useState } from 'react';
import Spin from 'antd/lib/spin';
import PageLayout from 'src/components/PageLayout';
import Row from 'antd/lib/row';
import Col from 'antd/lib/col';
import Card from 'antd/lib/card';
import Title from 'antd/lib/typography/Title';
import { getNFTList } from 'src/apis/nft';
import { Image } from 'antd';
import { formatAddress } from 'src/helpers/wallet';
import { useWeb3React } from '@web3-react/core';
import {
  INFTParsedTokenAccount,
  NFT_STANDARD_OPTIONS,
  NFTStandard
} from 'src/interfaces/nft';
import Button from 'src/components/Button';
import ReloadOutlined from '@ant-design/icons/ReloadOutlined';
import MyNFTPageStyle from './style';
import Radio from 'antd/lib/radio';
const MyNFTPage: React.FC = () => {
  const [loading, setLoading] = useState(false);
  const { account, chainId } = useWeb3React();
  const [nftStandard, setNftStandard] = useState(NFTStandard.ERC_721);
  const [items, setItems] = useState<INFTParsedTokenAccount[]>([]);

  // 暂时只显示ERC721
  const fetchNFT = async () => {
    if (account && chainId) {
      setLoading(true);
      getNFTList(chainId, account, nftStandard)
        .then((data) => {
          setItems(data);
          setLoading(false);
        })
        .catch((error) => {
          setLoading(false);
          console.error(error.response);
        });
    }
  };

  useEffect(() => {
    fetchNFT();
  }, [nftStandard, account, chainId]);

  return (
    <MyNFTPageStyle>
      <PageLayout>
        <div className='header'>
          <Title level={2}>我的 NFT</Title>
        </div>
        <div className='search-container'>
          <Radio.Group
            options={NFT_STANDARD_OPTIONS}
            onChange={(e) => setNftStandard(e.target.value)}
            value={nftStandard}
            optionType='button'
            buttonStyle='solid'
          />
          <Button className='reload-button' shape='round' onClick={fetchNFT}>
            <ReloadOutlined /> 重新加载
          </Button>
        </div>
        {loading ? (
          <div className='loading-container'>
            <Spin size='large' />
            <Title level={5}>加载可用NFT</Title>
          </div>
        ) : items.length > 0 ? (
          <Row gutter={16} className='image-container'>
            {items.map((item, index) => (
              <Col lg={6} md={12} sm={12} xs={12} key={index}>
                <Card
                  hoverable
                  cover={
                    <>
                      {nftStandard == NFTStandard.ERC_20 ? (
                        <></>
                      ) : (
                        <Image
                          width='100%'
                          height={150}
                          alt={item.name!}
                          src={item.image!}
                          fallback={`/加载失败.svg`}
                        />
                      )}
                    </>
                  }
                >
                  <Title level={5} ellipsis>
                    {item.name}
                  </Title>
                  {nftStandard != NFTStandard.ERC_20 && (
                    <>
                      <Title level={5}>#{item.tokenId}</Title>
                    </>
                  )}
                  <p>{formatAddress(item.contractAddress, 6)}</p>
                </Card>
              </Col>
            ))}
          </Row>
        ) : (
          <div className='no-data-container'>
            <Title level={2} style={{ color: '#a3a3a3' }}>
              暂无数据
            </Title>
          </div>
        )}
      </PageLayout>
    </MyNFTPageStyle>
  );
};

export default MyNFTPage;
