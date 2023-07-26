const {expect} = require("chai");
const {ethers} = require('hardhat');

describe("SwapAgent", function () {
  this.timeout(50000);
  let swapAgent;
  let mockERC721;
  let deployerAddress;

  beforeEach(async () => {
    const [deployer] = await ethers.getSigners();
    deployerAddress = await deployer.getAddress();

    const swapAgentFactory = await ethers.getContractFactory("ERC721SwapAgent", deployer);
    const swapAgentContract = await swapAgentFactory.deploy();

    swapAgent = await swapAgentContract.deployed();

    const mockERC721Factory = await ethers.getContractFactory("MockERC721", deployer);
    const mockERC721Contract = await mockERC721Factory.deploy("MockERC721", "M721");

    mockERC721 = await mockERC721Contract.deployed();
  });

  describe('isChainSupport', () => {
    it('should support BSC', async () => {
      const res = await swapAgent.isChainSupport("BSC");

      expect(res).to.eq(true);
    });

    it('should support ETH', async () => {
      const res = await swapAgent.isChainSupport("ETH");

      expect(res).to.eq(true);
    });

    it('should not support other chains', async () => {
      const res = await swapAgent.isChainSupport("other");

      expect(res).to.eq(false);
    });
  });

  describe('registerSwapPair', () => {
    it('should not support other chains', async () => {
      await expect(
        swapAgent.registerSwapPair(mockERC721.address, "other")
      ).to.be.revertedWith("ERC721SwapAgent::registerSwapPair:: chain is not supported");
    });

    it('should register correctly', async () => {
      await swapAgent.registerSwapPair(mockERC721.address, "BSC");

      expect(
        await swapAgent.registeredERC721("BSC", mockERC721.address)
      ).to.eq(true);
    });

    it('should emit event', async () => {
      const tx = await swapAgent.registerSwapPair(mockERC721.address, "BSC");
      await expect(tx).to.emit(swapAgent, "SwapPairRegister").withArgs(
        deployerAddress,
        mockERC721.address,
        mockERC721.name(),
        mockERC721.symbol(),
      );
    });

    it('should not register if token is already registered', async () => {
      await swapAgent.registerSwapPair(mockERC721.address, "BSC");

      expect(
        swapAgent.registerSwapPair(mockERC721.address, "BSC")
      ).to.be.revertedWith("ERC721SwapAgent::registerSwapPair:: token is already registered");
    });

    it('should not register if token does not have name', async () => {
      const mockERC721Factory = await ethers.getContractFactory("MockERC721");
      const mockERC721Contract = await mockERC721Factory.deploy("", "M721");

      mockERC721 = await mockERC721Contract.deployed();

      expect(
        swapAgent.registerSwapPair(mockERC721.address, "BSC")
      ).to.be.revertedWith("ERC721SwapAgent::registerSwapPair:: empty token name");
    });

    it('should not register if token does not have symbol', async () => {
      const mockERC721Factory = await ethers.getContractFactory("MockERC721");
      const mockERC721Contract = await mockERC721Factory.deploy("MockERC721", "");

      mockERC721 = await mockERC721Contract.deployed();

      expect(
        swapAgent.registerSwapPair(mockERC721.address, "BSC")
      ).to.be.revertedWith("ERC721SwapAgent::registerSwapPair:: empty token symbol");
    });
  });
});
