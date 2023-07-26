const {randomBytes} = require('crypto');
const {ethers} = require('hardhat');

const registerMockERC1155 = async (params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.dstChainId) throw new Error("dstChainId is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MockERC1155 = (await ethers.getContractFactory("MockERC1155", params.signers[0]));
  const mockERC1155 = await MockERC1155.attach(params.tokenContractAddr);

  const ERC1155SwapAgent = (await ethers.getContractFactory("ERC1155SwapAgent", params.signers[0]));
  const erc1155SwapAgent = await ERC1155SwapAgent.attach(params.agentAddr);

  const receipt = await erc1155SwapAgent.registerSwapPair(mockERC1155.address, params.dstChainId);
  console.log(">> Registered!!");

  return receipt;
}

const createSwapPairERC1155 = async (params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.fromChainId) throw new Error("fromChainId is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC1155SwapAgent = (await ethers.getContractFactory("ERC1155SwapAgent", params.signers[0]));
  const erc1155SwapAgent = await ERC1155SwapAgent.attach(params.agentAddr);

  await erc1155SwapAgent.createSwapPair(
    randomBytes(32),
    params.tokenContractAddr,
    params.fromChainId,
    params.tokenBaseUri,
  );

  console.log(">> SwapPair created !!");
};

const swapERC1155 = async (params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.tokenAddr) throw new Error("tokenAddr is not defined");
  if (!params.recipient) throw new Error("recipient is not defined");
  if (!params.ids.length) throw new Error("ids is not defined");
  if (!params.amounts.length) throw new Error("amounts is not defined");
  if (!params.dstChainId) throw new Error("dstChainId is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC1155SwapAgent = (await ethers.getContractFactory("ERC1155SwapAgent", params.signers[0]));
  const erc1155SwapAgent = await ERC1155SwapAgent.attach(params.agentAddr);

  await erc1155SwapAgent.swap(
    params.tokenAddr,
    params.recipient,
    params.ids,
    params.amounts,
    params.dstChainId,
  );

  console.log(">> ERC1155 swaped !!");
};

const fillERC1155 = async (params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.fromTokenAddr) throw new Error("fromTokenAddr is not defined");
  if (!params.recipient) throw new Error("recipient is not defined");
  if (!params.ids.length) throw new Error("ids is not defined");
  if (!params.amounts.length) throw new Error("amounts is not defined");
  if (!params.fromChainId) throw new Error("fromChainId is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC1155SwapAgent = (await ethers.getContractFactory("ERC1155SwapAgent", params.signers[0]));
  const erc1155SwapAgent = await ERC1155SwapAgent.attach(params.agentAddr);

  await erc1155SwapAgent.fill(
    randomBytes(32),
    params.fromTokenAddr,
    params.recipient,
    params.fromChainId,
    params.ids,
    params.amounts,
  );

  console.log(">> ERC1155 swaped !!");
};

module.exports = {
  registerMockERC1155, createSwapPairERC1155, swapERC1155, fillERC1155
}