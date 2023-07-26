const {ethers} = require('hardhat');

const approveMockERC1155 = async (params) => {
  if (!params.to) throw new Error("to is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MockERC1155 = (await ethers.getContractFactory("MockERC1155", params.signers[0]));
  const mockERC1155 = await MockERC1155.attach(params.tokenContractAddr);

  await mockERC1155.setApprovalForAll(params.to, true);
  console.log(`>> Approved MockERC1155 to ${params.to}`);
};

const approveMirroredERC1155 = async (params) => {
  if (!params.to) throw new Error("to is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MirroredERC1155 = (await ethers.getContractFactory("MirroredERC1155", params.signers[0]));
  const mirroredERC1155 = await MirroredERC1155.attach(params.tokenContractAddr);

  await mirroredERC1155.setApprovalForAll(params.to, true);
  console.log(`>> Approved MirroredERC1155 to ${params.to}`);
};

const mintMockERC1155 = async (params) => {
  if (!params.ids.length) throw new Error("ids is not defined");
  if (!params.amounts.length) throw new Error("amounts is not defined");
  if (!params.to) throw new Error("to is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MockERC1155 = (await ethers.getContractFactory("MockERC1155", params.signers[0]));
  const mockERC1155 = await MockERC1155.attach(params.tokenContractAddr);

  await mockERC1155.mintBatch(params.to, params.ids, params.amounts, "0x00");
  console.log(`>> Minted MockERC1155 to ${params.to}`);
}
module.exports = {
  approveMockERC1155,
  approveMirroredERC1155,
  mintMockERC1155
}