const path = require('path');
const fs = require('fs');

const erc721AgentFile = 'erc721-agent.json';
const erc721MockTokenFile = 'erc721-mock-token.json'
const erc721MirroredTokenFile = 'erc721-mirrored-token.json'

const readFileByChainId = (chainId, fileName) => {
  const filePath = path.join(__dirname, `../../chains/${chainId}/${fileName}`);
  return JSON.parse(fs.readFileSync(filePath).toString());
};

const writeFileByChainId = (chainId, fileName, data) => {
  const filePath = path.join(__dirname, `../../chains/${chainId}/${fileName}`);
  fs.writeFileSync(filePath, JSON.stringify(data));
};

const set721Agent = (chainId, data) => {
  writeFileByChainId(chainId, erc721AgentFile, data);
};

const get721Agent = (chainId) => {
  return readFileByChainId(chainId, erc721AgentFile);
};

const set721MockToken = (chainId, data) => {
  writeFileByChainId(chainId, erc721MockTokenFile, data);
};

const get721MockToken = (chainId) => {
  return readFileByChainId(chainId, erc721MockTokenFile);
};

const get721MirroredToken = (chainId) => {
  return readFileByChainId(chainId, erc721MirroredTokenFile);
};

module.exports = {
  set721Agent,
  get721Agent, set721MockToken, get721MockToken,
  get721MirroredToken
}