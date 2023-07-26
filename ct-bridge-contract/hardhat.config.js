require("dotenv").config();

require("@nomiclabs/hardhat-etherscan");
require("@nomiclabs/hardhat-waffle");
require("@typechain/hardhat");
require("hardhat-gas-reporter");
require("solidity-coverage");
require('@openzeppelin/hardhat-upgrades');
require('hardhat-deploy');

module.exports = {
  solidity: {
    version: "0.8.2",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
      },
    },
  },
  defaultNetwork: "cooChain",
  // defaultNetwork: "ctChain",
  networks: {
    
    ctChain: {
      chainId: 27,
      url: process.env.CT_URL,
      accounts: [
        process.env.RELAY_CONTRACT_DEPLOY_ADDR,
        process.env.FIR_PRIVATE_KEY,
      ],
    },
    cooChain: {
      chainId: 583,
      url: process.env.COO_URL,
      accounts: [
        process.env.RELAY_CONTRACT_DEPLOY_ADDR,
        // process.env.FIR_PRIVATE_KEY,
      ],
    },
    hardhat: {
      chainId: 31337,
      gas: 12000000,
      blockGasLimit: 0x1fffffffffffff,
      allowUnlimitedContractSize: true,
      loggingEnabled: true,
      mining: {
        auto: true,
        interval: 10000
      },
      accounts: [{
        privateKey: process.env.LOCAL_PRIVATE_KEY_1 || '',
        balance: "10000000000000000000000",
      }, {
        privateKey: process.env.LOCAL_PRIVATE_KEY_2 || '',
        balance: "10000000000000000000000",
      }, {
        privateKey: process.env.LOCAL_PRIVATE_KEY_3 || '',
        balance: "10000000000000000000000",
      }],
    },
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: "USD",
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
    customChains: [
      {
        apiKey: process.env.ETHERSCAN_API_KEY,
        network: "ctChain",
        chainId: 27,
        urls: {
          apiURL: "https://ctblock.cn/api",
          browserURL: "https://ctblock.cn",
        },
      },
      {
        apiKey: process.env.ETHERSCAN_API_KEY,
        network: "cooChain",
        chainId: 583,
        urls: {
          apiURL: "http://182.43.26.165:4000/api",
          browserURL: "http://182.43.26.165:4000",
        },
      }]
  },
};
