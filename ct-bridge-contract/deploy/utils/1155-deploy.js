const {ethers, upgrades} = require("hardhat");

const deployMockERC1155 = async (params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");
  if (!params.uri) throw new Error("uri is not defined");

  const MockERC1155 = (await ethers.getContractFactory("MockERC1155", params.signers[0]))

  console.log(`>> Deploying MockERC1155`);

  const mockERC1155 = await MockERC1155.deploy(params.uri);

  console.log(">> MockERC1155 is deployed!");

  return mockERC1155;
}

const deployERC1155Agent = async (params) => {
  if (!params.chainId) throw new Error("chainId is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC1155SwapAgent = (await ethers.getContractFactory("ERC1155SwapAgent", params.signers[0]));

  console.log(`>> Deploying ERC1155SwapAgent`);

  const agent = (await upgrades.deployProxy(ERC1155SwapAgent, [], {initializer: 'initialize'}));

  await agent.deployed();

  console.log(">> ERC1155SwapAgent is deployed!");

  console.log(agent.address);

  return agent;
}


module.exports = {
  deployMockERC1155, deployERC1155Agent
}