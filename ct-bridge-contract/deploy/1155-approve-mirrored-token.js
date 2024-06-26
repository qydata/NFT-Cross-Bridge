const {ethers} = require('hardhat');
const {approveMirroredERC1155} = require('./utils/1155-token');
const {get1155Agent, get1155MirroredToken} = require('./utils/1155-cache');

const func = async function (hre) {
  const chainId = hre.network.config.chainId?.toString() || '';

  console.warn(`[${func.tags[0]}]: chains/${chainId}/erc1155-mirrored-token.json must be prior created with completed details and the token should be prepared before swapping`);

  const signers = await ethers.getSigners();
  const mirrroredToken = get1155MirroredToken(chainId);
  const agent = get1155Agent(chainId);
  console.log(approveMirroredERC1155)
  await approveMirroredERC1155({
    to: agent.address,
    tokenContractAddr: mirrroredToken.address,
    signers,
  });
}

func.tags = ["ERC1155ApproveMirroredToken"];

module.exports = func;
