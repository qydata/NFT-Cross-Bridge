const {ethers} = require('hardhat');
const { swapERC1155 } = require('./utils/1155-agent');
const { get1155Agent, get1155MockToken } = require('./utils/1155-cache');

const func = async function (hre ) {
  if (!process.env.DST_CHAIN_ID) {
    throw new Error("no DST_CHAIN_ID");
  }
  const dstChainId = parseInt(process.env.DST_CHAIN_ID, 10).toString();
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get1155MockToken(chainId);
  const agent = get1155Agent(chainId);

  await swapERC1155({
    agentAddr: agent.address,
    dstChainId,
    tokenAddr: mockToken.address,
    recipient: signers[0].address,
    ids: mockToken.ids,
    amounts: mockToken.amounts,
    signers,
  });
}

func.tags = ["ERC1155ForwardSwap"];

module.exports = func;
