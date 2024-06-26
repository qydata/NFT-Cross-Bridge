const {ethers} = require('hardhat');
const { mintMockERC1155 } = require('./utils/1155-token');
const { get1155MockToken, set1155MockToken } = require('./utils/1155-cache');

const func = async function (hre ) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get1155MockToken(chainId);

  let nextId = 1;
  if (mockToken.ids.length) {
    const lastIdx = mockToken.ids.length - 1;
    nextId = mockToken.ids[lastIdx] + 1;
  }

  mockToken.ids = [...mockToken.ids, nextId];
  mockToken.amounts = [...mockToken.amounts, 1];
  await mintMockERC1155({
    ids: mockToken.ids,
    amounts: mockToken.amounts,
    to: signers[0].address,
    tokenContractAddr: mockToken.address,
    signers,
  });

  set1155MockToken(chainId, mockToken);
}

func.tags = ["ERC1155MintMockToken"];

module.exports = func;
