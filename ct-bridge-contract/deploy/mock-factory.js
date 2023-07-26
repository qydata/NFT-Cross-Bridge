const {ethers} = require('hardhat');

const func = async function (hre) {
  const signers = await ethers.getSigners();

  const MockFactory = (await ethers.getContractFactory("MockFactory", signers[0]))
  const mockFactory = await MockFactory.deploy();

  console.log(`>> MockFactory Address: ${mockFactory.address}`);
}

func.tags = ["MockFactory"];

module.exports = func;
