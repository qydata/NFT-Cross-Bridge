const {ethers} = require('hardhat');
const {deployERC721Agent} = require('./utils/721-deploy');
const {set721Agent} = require('./utils/721-cache');

const func = async function (hre) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const agent = await deployERC721Agent({
    chainId,
    signers,
  });

  set721Agent(chainId, {
    address: agent.address,
    registerTxHash: '',
    registerBlockNumber: '',
  });
}

func.tags = ["ERC721SwapAgent"];

module.exports = func;
