const {ethers} = require('hardhat');
const {deployMockERC20} = require('./utils/20-deploy');
const {get20Agent, set20MockToken} = require('./utils/20-cache');

const TOKEN_NAME = process.env.ERC20_MOCK_TOKEN_NAME || '';
const TOKEN_SYMBOL = process.env.ERC20_MOCK_TOKEN_SYMBOL || '';

const func = async function (hre) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const cache = get20Agent(chainId);
  const mockToken = await deployMockERC20({
    name: TOKEN_NAME,
    symbol: TOKEN_SYMBOL,
    signers,
    agentAddr: cache.address,
  });

  set20MockToken(chainId, {
    address: mockToken.address,
    name: TOKEN_NAME,
    symbol: TOKEN_SYMBOL
  });
}

func.tags = ["ERC20MockToken"];

module.exports = func;
