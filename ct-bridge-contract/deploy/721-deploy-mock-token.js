
const { deployMockERC721 } = require('./utils/721-deploy');
const { set721MockToken } = require('./utils/721-cache');const {ethers} = require('hardhat');

const BASE_URI = process.env.ERC721_MOCK_TOKEN_BASE_URL || '';

const func = async function (hre ) {
  const baseUri = process.env.BASE_URI === 'true' ? BASE_URI : '';
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = await deployMockERC721({
    baseUri,
    signers,
  });

  set721MockToken(chainId, {
    address: mockToken.address,
    baseUri,
    symbol: await mockToken.symbol(),
    name: await mockToken.name(),
    tokenId: '',
  });
}

func.tags = ["ERC721MockToken"];

module.exports = func;
