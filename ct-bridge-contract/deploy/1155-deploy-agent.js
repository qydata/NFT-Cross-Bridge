const {ethers} = require('hardhat');
const { deployERC1155Agent } = require('./utils/1155-deploy');
const { set1155Agent } = require('./utils/1155-cache');

const func = async function (hre ) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const agent = await deployERC1155Agent({
    chainId,
    signers,
  });

  set1155Agent(chainId, {
    address: agent.address,
    registerTxHash: '',
    registerBlockNumber: '',
  });
}

func.tags = ["ERC1155SwapAgent"];

module.exports = func;
