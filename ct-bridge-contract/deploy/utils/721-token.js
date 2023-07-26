const {ethers} = require('hardhat');

const approveMockERC721 = async (params) => {
  if (!params.tokenId) throw new Error("tokenId is not defined");
  if (!params.to) throw new Error("to is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MockERC721 = (await ethers.getContractFactory("MockERC721", params.signers[0]));
  const mockERC721 = await MockERC721.attach(params.tokenContractAddr);

  await mockERC721.approve(params.to, params.tokenId);
  console.log(`>> Approved MockERC721 to ${params.to}`);
};

const approveMirroredERC721 = async (params) => {
  if (!params.tokenId) throw new Error("tokenId is not defined");
  if (!params.to) throw new Error("to is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MirroredERC721 = (await ethers.getContractFactory("MirroredERC721", params.signers[0]));
  const mirroredERC721 = await MirroredERC721.attach(params.tokenContractAddr);

  await mirroredERC721.approve(params.to, params.tokenId);
  console.log(`>> Approved MirroredERC721 to ${params.to}`);
};

const mintMockERC721 = async (params) => {
  if (!params.tokenId) throw new Error("tokenId is not defined");
  if (!params.tokenUri) throw new Error("tokenUri is not defined");
  if (!params.to) throw new Error("to is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MockERC721 = (await ethers.getContractFactory("MockERC721", params.signers[0]));
  const mockERC721 = await MockERC721.attach(params.tokenContractAddr);

  // TODO: need to check for block confirmation
  await mockERC721.safeMint(params.to, params.tokenId);
  console.log(`>> Minted MockERC721 to ${params.to}`);

  await mockERC721.setTokenURI(params.tokenId, params.tokenUri);
  console.log(">> Set MockERC721 tokenURI");
}
module.exports = {
  approveMockERC721, approveMirroredERC721, mintMockERC721
}