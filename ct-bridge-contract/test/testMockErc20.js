const {expect} = require("chai");
const {ethers} = require('hardhat');
const hre = require('hardhat');

describe("SwapAgent", function () {
  this.timeout(50000)
  let swapAgent;
  let deployerAddress;

  beforeEach(async () => {

  });
 
  it('deploy MockERC20', async () => {
    const [deployer] = await ethers.getSigners();
    deployerAddress = await deployer.getAddress();

    const swapAgentFactory = await ethers.getContractFactory("MockERC20", deployer);
    const swapAgentContract = await swapAgentFactory.deploy('test1', "test2", 4);
    swapAgent = await swapAgentContract.deployed();
    console.log(swapAgent.address)
  });

  it('test mint', async () => {
    const [deployer] = await ethers.getSigners();
    deployerAddress = await deployer.getAddress();

    const swapAgentFactory = await ethers.getContractFactory("MockERC20");
    const swapAgentContract = await swapAgentFactory.attach('0x53ADC783b127452da27861F49c5998931C2529d4');
    console.log(await swapAgentContract.mint(deployer.address,  4))
  });

  it('test burn', async () => {
    const [deployer] = await ethers.getSigners();
    deployerAddress = await deployer.getAddress();

    const swapAgentFactory = await ethers.getContractFactory("MockERC20");
    const swapAgentContract = await swapAgentFactory.attach('0x53ADC783b127452da27861F49c5998931C2529d4');
    console.log(deployer.address)
    console.log(swapAgentContract)
    console.log(await swapAgentContract.burn(deployer.address, 1))
    // console.log(await swapAgentContract.burn(1))
  });

  it('test decmial', async () => {
    const [deployer] = await ethers.getSigners();
    deployerAddress = await deployer.getAddress();

    const swapAgentFactory = await ethers.getContractFactory("MockERC20");
    const swapAgentContract = await swapAgentFactory.attach('0x53ADC783b127452da27861F49c5998931C2529d4');

    console.log(await swapAgentContract.decimals())
  });

  it('verify', async () => {
    let contractAddress = '0x53ADC783b127452da27861F49c5998931C2529d4'
    await hre.run("verify:verify", {
      address: contractAddress,
      // constructorArguments: ['test_name', 'test_symbol'],
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
