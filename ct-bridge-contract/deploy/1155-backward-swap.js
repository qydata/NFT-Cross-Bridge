const {ethers} = require('hardhat');
const { swapERC1155 } = require('./utils/1155-agent');
const { get1155Agent, get1155MirroredToken } = require('./utils/1155-cache');

const func = async function (hre ) {
  if (!process.env.DST_CHAIN_ID) {
    throw new Error("no DST_CHAIN_ID");
  }

  const chainId = hre.network.config.chainId?.toString() || '';
  const dstChainId = parseInt(process.env.DST_CHAIN_ID, 10).toString();

  console.warn(`[${func.tags[0]}]: chains/${chainId}/erc1155-mirrored-token.json must be prior created with completed details and the token should be prepared before swapping`);

  const signers = await ethers.getSigners();
  const mirrored = get1155MirroredToken(chainId);
  const agent = get1155Agent(chainId);

  if (!mirrored.address) {
    throw new Error("no mirrored token address");
  }
  if (!mirrored.ids) {
    throw new Error("no mirrored tokenId");
  }
  if (!mirrored.amounts) {
    throw new Error("no mirrored tokenId");
  }

  await swapERC1155({
    agentAddr: agent.address,
    dstChainId,
    tokenAddr: mirrored.address,
    recipient: signers[0].address,
    ids: mirrored.ids,
    amounts: mirrored.amounts,
    signers,
  });
}

func.tags = ["ERC1155BackwardSwap"];

module.exports = func;
