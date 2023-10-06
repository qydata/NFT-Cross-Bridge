const path = require('path');
const fs = require('fs');

const erc20AgentFile = 'erc20-agent.json';
const erc20MockTokenFile = 'erc20-mock-token.json'
const erc20MirroredTokenFile = 'erc20-mirrored-token.json'

const readFileByChainId = (chainId, fileName) => {
  const filePath = path.join(__dirname, `../../chains/${chainId}/${fileName}`);
  return JSON.parse(fs.readFileSync(filePath).toString());
};

const writeFileByChainId = (chainId, fileName, data) => {
  const filePath = path.join(__dirname, `../../chains/${chainId}/${fileName}`);
  fs.writeFileSync(filePath, JSON.stringify(data));
};

const set20Agent = (chainId, data) => {
  writeFileByChainId(chainId, erc20AgentFile, data);
};

const get20Agent = (chainId) => {
  return readFileByChainId(chainId, erc20AgentFile);
};

const set20MockToken = (chainId, data) => {
  writeFileByChainId(chainId, erc20MockTokenFile, data);
};

const get20MockToken = (chainId) => {
  return readFileByChainId(chainId, erc20MockTokenFile);
};

const get20MirroredToken = (chainId) => {
  return readFileByChainId(chainId, erc20MirroredTokenFile);
};

module.exports = {
  set20Agent,
  get20Agent,
  set20MockToken,
  get20MockToken,
  get20MirroredToken
}