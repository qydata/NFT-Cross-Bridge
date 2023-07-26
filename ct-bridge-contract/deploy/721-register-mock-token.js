const {ethers} = require('hardhat');
const { registerMockERC721 } = require('./utils/721-agent');
const { get721Agent, get721MockToken, set721Agent } = require('./utils/721-cache');

const func = async function (hre ) {
  if (!process.env.DST_CHAIN_ID) {
    throw new Error("no DST_CHAIN_ID");
  }
  const dstChainId = parseInt(process.env.DST_CHAIN_ID, 10).toString();
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get721MockToken(chainId);
  const agent = get721Agent(chainId);

  const receipt = await registerMockERC721({
    agentAddr: agent.address,
    dstChainId,
    tokenContractAddr: mockToken.address,
    signers,
  });

  agent.registerBlockNumber = receipt.blockNumber?.toString() || '';
  agent.registerTxHash = receipt.hash;

  set721Agent(chainId, agent)
}

func.tags = ["ERC721Register"];

module.exports = func;
