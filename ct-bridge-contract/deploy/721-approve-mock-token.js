const {ethers} = require('hardhat');
const {approveMockERC721} = require('./utils/721-token');
const {get721MockToken, get721Agent} = require('./utils/721-cache');

const func = async function (hre) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get721MockToken(chainId);
  const agent = get721Agent(chainId);

  await approveMockERC721({
    tokenId: mockToken.tokenId,
    to: agent.address,
    tokenContractAddr: mockToken.address,
    signers,
  });
}

func.tags = ["ERC721ApproveMockToken"];
module.exports = func;
