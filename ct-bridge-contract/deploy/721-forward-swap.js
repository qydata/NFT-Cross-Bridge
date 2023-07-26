const {ethers} = require('hardhat');
const { swapERC721 } = require('./utils/721-agent');
const { get721Agent, get721MockToken } = require('./utils/721-cache');

const func = async function (hre ) {
  if (!process.env.DST_CHAIN_ID) {
    throw new Error("no DST_CHAIN_ID");
  }
  const dstChainId = parseInt(process.env.DST_CHAIN_ID, 10).toString();
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get721MockToken(chainId);
  const agent = get721Agent(chainId);

  await swapERC721({
    agentAddr: agent.address,
    dstChainId,
    tokenAddr: mockToken.address,
    recipient: signers[0].address,
    tokenId: mockToken.tokenId,
    signers,
  });
}

func.tags = ["ERC721ForwardSwap"];

module.exports = func;
