const {ethers} = require('hardhat');
const {deployMockERC1155} = require('./utils/1155-deploy');
const {get1155Agent, set1155MockToken} = require('./utils/1155-cache');

const BASE_URI = process.env.ERC1151_MOCK_TOKEN_BASE_URL || '';

const func = async function (hre) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const cache = get1155Agent(chainId);
  const mockToken = await deployMockERC1155({
    uri: BASE_URI,
    signers,
    agentAddr: cache.address,
  });

  set1155MockToken(chainId, {
    address: mockToken.address,
    uri: BASE_URI,
    ids: [],
    amounts: [],
  });
}

func.tags = ["ERC1155MockToken"];

module.exports = func;
