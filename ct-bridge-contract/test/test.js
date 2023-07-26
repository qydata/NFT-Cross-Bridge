const {expect} = require("chai");
const {ethers} = require('hardhat');
const hre = require('hardhat');

describe("SwapAgent", function () {
  this.timeout(50000)
  let swapAgent;
  let deployerAddress;

  beforeEach(async () => {

  });

  it('deploy', async () => {
    const [deployer] = await ethers.getSigners();
    deployerAddress = await deployer.getAddress();

    const swapAgentFactory = await ethers.getContractFactory("ERC721SwapAgent", deployer);
    const swapAgentContract = await swapAgentFactory.deploy();
    swapAgent = await swapAgentContract.deployed();
    console.log(swapAgent.address)
  });

  it('verify', async () => {
    let contractAddress = '0x79F23Cee5e2Ee911327157dF942D3cFe02dBF9D4'
    await hre.run("verify:verify", {
      address: contractAddress,
      constructorArguments: [],
    });
  });

  it('test', async () => {
    const [deployer] = await ethers.getSigners();
    deployerAddress = await deployer.getAddress();

    const swapAgentFactory = await ethers.getContractFactory("ERC721SwapAgent", deployer);
    const swapAgentContract = await swapAgentFactory.attach('0xEfB4bC4755A7569bC34aDEe01a4C2fe42b19554c');
    
    console.log(await swapAgentContract.owner())
  });

});
