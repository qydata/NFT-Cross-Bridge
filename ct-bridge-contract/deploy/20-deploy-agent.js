const {ethers} = require('hardhat');
const { deployERC20Agent } = require('./utils/20-deploy');
const { set20Agent } = require('./utils/20-cache');

const func = async function (hre ) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const agent = await deployERC20Agent({
    chainId,
    signers,
  });

  set20Agent(chainId, {
    address: agent.address,
    registerTxHash: '',
    registerBlockNumber: '',
  });
}

func.tags = ["ERC20SwapAgent"];

module.exports = func;
