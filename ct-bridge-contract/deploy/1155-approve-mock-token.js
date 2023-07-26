const {ethers} = require('hardhat');
const { approveMockERC1155 } = require('./utils/1155-token');
const { get1155MockToken, get1155Agent } = require('./utils/1155-cache');

const func = async function (hre ) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get1155MockToken(chainId);
  const agent = get1155Agent(chainId);

  await approveMockERC1155({
    to: agent.address,
    tokenContractAddr: mockToken.address,
    signers,
  });
}

func.tags = ["ERC1155ApproveMockToken"];

module.exports = func;
