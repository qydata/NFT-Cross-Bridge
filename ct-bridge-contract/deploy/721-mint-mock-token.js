const {ethers} = require('hardhat');
const { mintMockERC721 } = require('./utils/721-token');
const { get721MockToken, set721MockToken } = require('./utils/721-cache');

const func = async function (hre ) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get721MockToken(chainId);

  mockToken.tokenId = (mockToken.tokenId ? parseInt(mockToken.tokenId, 10) + 1 : 1).toString();

  let tokenUri = mockToken.baseUri + mockToken.tokenId;
  if (mockToken.baseUri) {
    tokenUri = mockToken.tokenId;
  }

  await mintMockERC721({
    tokenId: mockToken.tokenId,
    tokenUri,
    to: signers[0].address,
    tokenContractAddr: mockToken.address,
    signers,
  });

  set721MockToken(chainId, mockToken);
}

func.tags = ["ERC721MintMockToken"];

module.exports = func;
