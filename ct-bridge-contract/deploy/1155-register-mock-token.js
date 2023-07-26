const {ethers} = require('hardhat');
const { registerMockERC1155 } = require('./utils/1155-agent');
const { get1155Agent, get1155MockToken, set1155Agent } = require('./utils/1155-cache');

const func = async function (hre ) {
  if (!process.env.DST_CHAIN_ID) {
    throw new Error("no DST_CHAIN_ID");
  }
  const dstChainId = parseInt(process.env.DST_CHAIN_ID, 10).toString();
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get1155MockToken(chainId);
  const agent = get1155Agent(chainId);

  const receipt = await registerMockERC1155({
    agentAddr: agent.address,
    dstChainId,
    tokenContractAddr: mockToken.address,
    signers,
  });

  agent.registerBlockNumber = receipt.blockNumber?.toString() || '';
  agent.registerTxHash = receipt.hash;

  set1155Agent(chainId, agent)
}

func.tags = ["ERC1155Register"];

module.exports = func;
