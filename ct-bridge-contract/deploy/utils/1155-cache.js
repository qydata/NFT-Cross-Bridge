const path = require('path');
const fs = require('fs');

const erc1155AgentFile = 'erc1155-agent.json';
const erc1155MockTokenFile = 'erc1155-mock-token.json'
const erc1155MirroredTokenFile = 'erc1155-mirrored-token.json'

const readFileByChainId = (chainId, fileName) => {
  const filePath = path.join(__dirname, `../../chains/${chainId}/${fileName}`);
  return JSON.parse(fs.readFileSync(filePath).toString());
};

const writeFileByChainId = (chainId, fileName, data) => {
  const filePath = path.join(__dirname, `../../chains/${chainId}/${fileName}`);
  fs.writeFileSync(filePath, JSON.stringify(data));
};

const set1155Agent = (chainId, data) => {
  writeFileByChainId(chainId, erc1155AgentFile, data);
};

const get1155Agent = (chainId) => {
  return readFileByChainId(chainId, erc1155AgentFile);
};

const set1155MockToken = (chainId, data) => {
  writeFileByChainId(chainId, erc1155MockTokenFile, data);
};

const get1155MockToken = (chainId) => {
  return readFileByChainId(chainId, erc1155MockTokenFile);
};

const get1155MirroredToken = (chainId) => {
  return readFileByChainId(chainId, erc1155MirroredTokenFile);
};

module.exports = {
  set1155Agent,
  get1155Agent,
  set1155MockToken,
  get1155MockToken,
  get1155MirroredToken
}