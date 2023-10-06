const {ethers, upgrades} = require("hardhat");

const deployMockERC20 = async (params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");
  if (!params.name) throw new Error("name is not defined");
  if (!params.symbol) throw new Error("symbol is not defined");
  if (!params.symbol) throw new Error("symbol is not defined");

  const MockERC20 = (await ethers.getContractFactory("MockERC20", params.signers[0]))

  console.log(`>> Deploying MockERC20`);

  const mockERC20 = await MockERC20.deploy(params.name, params.symbol);

  console.log(">> MockERC20 is deployed!");

  return mockERC20;
}

const deployERC20Agent = async (params) => {
  if (!params.chainId) throw new Error("chainId is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC20SwapAgent = (await ethers.getContractFactory("ERC20SwapAgent", params.signers[0]));

  console.log(`>> Deploying ERC20SwapAgent`);

  const agent = (await upgrades.deployProxy(ERC20SwapAgent, [], {initializer: 'initialize'}));

  await agent.deployed();

  console.log(">> ERC20SwapAgent is deployed!");

  console.log(agent.address);

  return agent;
}


module.exports = {
  deployMockERC20, deployERC20Agent
}