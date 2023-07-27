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

    const swapAgentFactory = await ethers.getContractFactory("MirroredERC1155");
    const swapAgentContract = await swapAgentFactory.attach('0x275ffC2A8459F0145145c08027f28B5f910f9db1');
    
    console.log(await swapAgentContract.owner())
    console.log(await swapAgentContract.uri('26067885190938743157339440157251197165479395358752782221930628278050914522887'))
    // console.log(await swapAgentContract.name())
    console.log(await swapAgentContract.symbol())
  });

});
