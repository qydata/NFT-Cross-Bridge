const {randomBytes} = require('crypto');
const {ethers} = require('hardhat');

const registerMockERC721 = async (params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.dstChainId) throw new Error("dstChainId is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MockERC721 = (await ethers.getContractFactory("MockERC721", params.signers[0]));
  const mockERC721 = await MockERC721.attach(params.tokenContractAddr);

  const ERC721SwapAgent = (await ethers.getContractFactory("ERC721SwapAgent", params.signers[0]));
  const erc721SwapAgent = await ERC721SwapAgent.attach(params.agentAddr);

  const receipt = await erc721SwapAgent.registerSwapPair(mockERC721.address, params.dstChainId);
  console.log(">> Registered!!")

  return receipt;
}

const createSwapPairERC721 = async (params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.fromChainId) throw new Error("fromChainId is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.tokenName) throw new Error("tokenName is not defined");
  if (!params.tokenSymbol) throw new Error("tokenSymbol is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC721SwapAgent = (await ethers.getContractFactory("ERC721SwapAgent", params.signers[0]));
  const erc721SwapAgent = await ERC721SwapAgent.attach(params.agentAddr);

  await erc721SwapAgent.createSwapPair(
    randomBytes(32),
    params.tokenContractAddr,
    params.fromChainId,
    params.tokenBaseUri,
    params.tokenName,
    params.tokenSymbol,
  );

  console.log(">> SwapPair created !!");
};

const swapERC721 = async (params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.tokenAddr) throw new Error("tokenAddr is not defined");
  if (!params.recipient) throw new Error("recipient is not defined");
  if (!params.tokenId) throw new Error("tokenId is not defined");
  if (!params.dstChainId) throw new Error("dstChainId is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC721SwapAgent = (await ethers.getContractFactory("ERC721SwapAgent", params.signers[0]));
  const erc721SwapAgent = await ERC721SwapAgent.attach(params.agentAddr);

  await erc721SwapAgent.swap(
    params.tokenAddr,
    params.recipient,
    params.tokenId,
    params.dstChainId,
  );

  console.log(">> ERC721 swaped !!");
};

const fillERC721 = async (params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.fromTokenAddr) throw new Error("fromTokenAddr is not defined");
  if (!params.recipient) throw new Error("recipient is not defined");
  if (!params.tokenId) throw new Error("tokenId is not defined");
  if (!params.fromChainId) throw new Error("fromChainId is not defined");
  if (!params.tokenURI) throw new Error("tokenURI is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC721SwapAgent = (await ethers.getContractFactory("ERC721SwapAgent", params.signers[0]));
  const erc721SwapAgent = await ERC721SwapAgent.attach(params.agentAddr);

  await erc721SwapAgent.fill(
    randomBytes(32),
    params.fromTokenAddr,
    params.recipient,
    params.fromChainId,
    params.tokenId,
    params.tokenURI,
  );

  console.log(">> ERC721 swaped !!");
};

module.exports = {registerMockERC721, createSwapPairERC721, swapERC721, fillERC721}