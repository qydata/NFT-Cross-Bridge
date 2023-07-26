/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../common";
import type {
  MirroredERC721,
  MirroredERC721Interface,
} from "../../../contracts/tokens/MirroredERC721";

const _abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "approved",
        type: "address",
      },
      {
        indexed: true,
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "Approval",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "operator",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bool",
        name: "approved",
        type: "bool",
      },
    ],
    name: "ApprovalForAll",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferred",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "account",
        type: "address",
      },
    ],
    name: "Paused",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        indexed: true,
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "Transfer",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "account",
        type: "address",
      },
    ],
    name: "Unpaused",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "approve",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "owner",
        type: "address",
      },
    ],
    name: "balanceOf",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "baseURI",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "burn",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "getApproved",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "name",
        type: "string",
      },
      {
        internalType: "string",
        name: "symbol",
        type: "string",
      },
    ],
    name: "initialize",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        internalType: "address",
        name: "operator",
        type: "address",
      },
    ],
    name: "isApprovedForAll",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "name",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "owner",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "ownerOf",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "pause",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "paused",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "renounceOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "safeMint",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "safeTransferFrom",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
      {
        internalType: "bytes",
        name: "_data",
        type: "bytes",
      },
    ],
    name: "safeTransferFrom",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "operator",
        type: "address",
      },
      {
        internalType: "bool",
        name: "approved",
        type: "bool",
      },
    ],
    name: "setApprovalForAll",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "baseURI_",
        type: "string",
      },
    ],
    name: "setBaseURI",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
      {
        internalType: "string",
        name: "_tokenURI",
        type: "string",
      },
    ],
    name: "setTokenURI",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes4",
        name: "interfaceId",
        type: "bytes4",
      },
    ],
    name: "supportsInterface",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "symbol",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "index",
        type: "uint256",
      },
    ],
    name: "tokenByIndex",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "index",
        type: "uint256",
      },
    ],
    name: "tokenOfOwnerByIndex",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "tokenURI",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "totalSupply",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "transferFrom",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "transferOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "unpause",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b50612e76806100206000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c80635c975abb116100f957806395d89b4111610097578063b88d4fde11610071578063b88d4fde14610381578063c87b56dd14610394578063e985e9c5146103a7578063f2fde38b146103e3576101c4565b806395d89b4114610353578063a14481941461035b578063a22cb4651461036e576101c4565b806370a08231116100d357806370a082311461031e578063715018a6146103315780638456cb59146103395780638da5cb5b14610341576101c4565b80635c975abb146102f85780636352211e146103035780636c0360eb14610316576101c4565b80632f745c591161016657806342966c681161014057806342966c68146102ac5780634cd88b76146102bf5780634f6ccce7146102d257806355f804b3146102e5576101c4565b80632f745c591461027e5780633f4ba83a1461029157806342842e0e14610299576101c4565b8063095ea7b3116101a2578063095ea7b314610231578063162094c41461024657806318160ddd1461025957806323b872dd1461026b576101c4565b806301ffc9a7146101c957806306fdde03146101f1578063081812fc14610206575b600080fd5b6101dc6101d7366004612b34565b6103f6565b60405190151581526020015b60405180910390f35b6101f9610409565b6040516101e89190612cfc565b610219610214366004612c08565b61049b565b6040516001600160a01b0390911681526020016101e8565b61024461023f366004612b0b565b610535565b005b610244610254366004612c20565b610667565b6099545b6040519081526020016101e8565b610244610279366004612a1d565b6106d0565b61025d61028c366004612b0b565b610758565b610244610800565b6102446102a7366004612a1d565b610865565b6102446102ba366004612c08565b610880565b6102446102cd366004612b6c565b610907565b61025d6102e0366004612c08565b610a62565b6102446102f3366004612bd5565b610b14565b60fb5460ff166101dc565b610219610311366004612c08565b610b28565b6101f9610bb3565b61025d61032c3660046129d1565b610c42565b610244610cdc565b610244610d41565b61012d546001600160a01b0316610219565b6101f9610da4565b610244610369366004612b0b565b610db3565b61024461037c366004612ad1565b610e18565b61024461038f366004612a58565b610eea565b6101f96103a2366004612c08565b610f78565b6101dc6103b53660046129eb565b6001600160a01b039182166000908152606a6020908152604080832093909416825291909152205460ff1690565b6102446103f13660046129d1565b610f83565b600061040182611063565b90505b919050565b60606065805461041890612d7e565b80601f016020809104026020016040519081016040528092919081815260200182805461044490612d7e565b80156104915780601f1061046657610100808354040283529160200191610491565b820191906000526020600020905b81548152906001019060200180831161047457829003601f168201915b5050505050905090565b6000818152606760205260408120546001600160a01b03166105195760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b60648201526084015b60405180910390fd5b506000908152606960205260409020546001600160a01b031690565b600061054082610b28565b9050806001600160a01b0316836001600160a01b031614156105ca5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560448201527f72000000000000000000000000000000000000000000000000000000000000006064820152608401610510565b336001600160a01b03821614806105e657506105e681336103b5565b6106585760405162461bcd60e51b815260206004820152603860248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760448201527f6e6572206e6f7220617070726f76656420666f7220616c6c00000000000000006064820152608401610510565b61066283836110a1565b505050565b61012d546001600160a01b031633146106c25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610510565b6106cc828261110f565b5050565b6106db335b826111b8565b61074d5760405162461bcd60e51b815260206004820152603160248201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f60448201527f776e6572206e6f7220617070726f7665640000000000000000000000000000006064820152608401610510565b6106628383836112af565b600061076383610c42565b82106107d75760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201527f74206f6620626f756e64730000000000000000000000000000000000000000006064820152608401610510565b506001600160a01b03919091166000908152609760209081526040808320938352929052205490565b61012d546001600160a01b0316331461085b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610510565b610863611487565b565b61066283838360405180602001604052806000815250610eea565b610889336106d5565b6108fb5760405162461bcd60e51b815260206004820152603060248201527f4552433732314275726e61626c653a2063616c6c6572206973206e6f74206f7760448201527f6e6572206e6f7220617070726f766564000000000000000000000000000000006064820152608401610510565b61090481611523565b50565b600054610100900460ff1680610920575060005460ff16155b6109835760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610510565b600054610100900460ff161580156109ae576000805460ff1961ff0019909116610100171660011790555b610a2185858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8901819004810282018101909252878152925087915086908190840183828082843760009201919091525061152c92505050565b610a29611603565b610a31611603565b610a396116d6565b610a4161178d565b610a49611603565b8015610a5b576000805461ff00191690555b5050505050565b6000610a6d60995490565b8210610ae15760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201527f7574206f6620626f756e647300000000000000000000000000000000000000006064820152608401610510565b60998281548110610b0257634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050919050565b80516106cc9061019190602084019061280d565b6000818152606760205260408120546001600160a01b0316806104015760405162461bcd60e51b815260206004820152602960248201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460448201527f656e7420746f6b656e00000000000000000000000000000000000000000000006064820152608401610510565b6101918054610bc190612d7e565b80601f0160208091040260200160405190810160405280929190818152602001828054610bed90612d7e565b8015610c3a5780601f10610c0f57610100808354040283529160200191610c3a565b820191906000526020600020905b815481529060010190602001808311610c1d57829003601f168201915b505050505081565b60006001600160a01b038216610cc05760405162461bcd60e51b815260206004820152602a60248201527f4552433732313a2062616c616e636520717565727920666f7220746865207a6560448201527f726f2061646472657373000000000000000000000000000000000000000000006064820152608401610510565b506001600160a01b031660009081526068602052604090205490565b61012d546001600160a01b03163314610d375760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610510565b6108636000611844565b61012d546001600160a01b03163314610d9c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610510565b610863611897565b60606066805461041890612d7e565b61012d546001600160a01b03163314610e0e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610510565b6106cc828261191f565b6001600160a01b038216331415610e715760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610510565b336000818152606a602090815260408083206001600160a01b0387168085529252909120805460ff1916841515179055906001600160a01b03167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610ede911515815260200190565b60405180910390a35050565b610ef433836111b8565b610f665760405162461bcd60e51b815260206004820152603160248201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f60448201527f776e6572206e6f7220617070726f7665640000000000000000000000000000006064820152608401610510565b610f7284848484611939565b50505050565b6060610401826119b7565b61012d546001600160a01b03163314610fde5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610510565b6001600160a01b03811661105a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610510565b61090481611844565b60006001600160e01b031982167f780e9d63000000000000000000000000000000000000000000000000000000001480610401575061040182611b36565b600081815260696020526040902080546001600160a01b0319166001600160a01b03841690811790915581906110d682610b28565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000828152606760205260409020546001600160a01b03166111995760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201527f6578697374656e7420746f6b656e0000000000000000000000000000000000006064820152608401610510565b600082815260c96020908152604090912082516106629284019061280d565b6000818152606760205260408120546001600160a01b03166112315760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b6064820152608401610510565b600061123c83610b28565b9050806001600160a01b0316846001600160a01b031614806112775750836001600160a01b031661126c8461049b565b6001600160a01b0316145b806112a757506001600160a01b038082166000908152606a602090815260408083209388168352929052205460ff165b949350505050565b826001600160a01b03166112c282610b28565b6001600160a01b03161461133e5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a207472616e73666572206f6620746f6b656e2074686174206960448201527f73206e6f74206f776e00000000000000000000000000000000000000000000006064820152608401610510565b6001600160a01b0382166113b95760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610510565b6113c4838383611bd1565b6113cf6000826110a1565b6001600160a01b03831660009081526068602052604081208054600192906113f8908490612d3b565b90915550506001600160a01b0382166000908152606860205260408120805460019290611426908490612d0f565b909155505060008181526067602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b60fb5460ff166114d95760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610510565b60fb805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b61090481611c2f565b600054610100900460ff1680611545575060005460ff16155b6115a85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610510565b600054610100900460ff161580156115d3576000805460ff1961ff0019909116610100171660011790555b6115db611c6f565b6115e3611c6f565b6115ed8383611d29565b8015610662576000805461ff0019169055505050565b600054610100900460ff168061161c575060005460ff16155b61167f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610510565b600054610100900460ff161580156116aa576000805460ff1961ff0019909116610100171660011790555b6116b2611c6f565b6116ba611c6f565b6116c2611c6f565b8015610904576000805461ff001916905550565b600054610100900460ff16806116ef575060005460ff16155b6117525760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610510565b600054610100900460ff1615801561177d576000805460ff1961ff0019909116610100171660011790555b611785611c6f565b6116c2611e0e565b600054610100900460ff16806117a6575060005460ff16155b6118095760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610510565b600054610100900460ff16158015611834576000805460ff1961ff0019909116610100171660011790555b61183c611c6f565b6116c2611ed3565b61012d80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60fb5460ff16156118ea5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610510565b60fb805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586115063390565b6106cc828260405180602001604052806000815250611f83565b6119448484846112af565b61195084848484612001565b610f725760405162461bcd60e51b815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b6064820152608401610510565b6000818152606760205260409020546060906001600160a01b0316611a445760405162461bcd60e51b815260206004820152603160248201527f45524337323155524953746f726167653a2055524920717565727920666f722060448201527f6e6f6e6578697374656e7420746f6b656e0000000000000000000000000000006064820152608401610510565b600082815260c9602052604081208054611a5d90612d7e565b80601f0160208091040260200160405190810160405280929190818152602001828054611a8990612d7e565b8015611ad65780601f10611aab57610100808354040283529160200191611ad6565b820191906000526020600020905b815481529060010190602001808311611ab957829003601f168201915b505050505090506000611ae7612159565b9050805160001415611afb57509050610404565b815115611b2d578082604051602001611b15929190612c91565b60405160208183030381529060405292505050610404565b6112a784612169565b60006001600160e01b031982167f80ac58cd000000000000000000000000000000000000000000000000000000001480611b9957506001600160e01b031982167f5b5e139f00000000000000000000000000000000000000000000000000000000145b8061040157507f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b0319831614610401565b60fb5460ff1615611c245760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610510565b610662838383612252565b611c388161230f565b600081815260c9602052604090208054611c5190612d7e565b15905061090457600081815260c96020526040812061090491612891565b600054610100900460ff1680611c88575060005460ff16155b611ceb5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610510565b600054610100900460ff161580156116c2576000805460ff1961ff0019909116610100171660011790558015610904576000805461ff001916905550565b600054610100900460ff1680611d42575060005460ff16155b611da55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610510565b600054610100900460ff16158015611dd0576000805460ff1961ff0019909116610100171660011790555b8251611de390606590602086019061280d565b508151611df790606690602085019061280d565b508015610662576000805461ff0019169055505050565b600054610100900460ff1680611e27575060005460ff16155b611e8a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610510565b600054610100900460ff16158015611eb5576000805460ff1961ff0019909116610100171660011790555b60fb805460ff191690558015610904576000805461ff001916905550565b600054610100900460ff1680611eec575060005460ff16155b611f4f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610510565b600054610100900460ff16158015611f7a576000805460ff1961ff0019909116610100171660011790555b6116c233611844565b611f8d83836123b6565b611f9a6000848484612001565b6106625760405162461bcd60e51b815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b6064820152608401610510565b60006001600160a01b0384163b1561214e57604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290612045903390899088908890600401612cc0565b602060405180830381600087803b15801561205f57600080fd5b505af192505050801561208f575060408051601f3d908101601f1916820190925261208c91810190612b50565b60015b612134573d8080156120bd576040519150601f19603f3d011682016040523d82523d6000602084013e6120c2565b606091505b50805161212c5760405162461bcd60e51b815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b6064820152608401610510565b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490506112a7565b506001949350505050565b6060610191805461041890612d7e565b6000818152606760205260409020546060906001600160a01b03166121f65760405162461bcd60e51b815260206004820152602f60248201527f4552433732314d657461646174613a2055524920717565727920666f72206e6f60448201527f6e6578697374656e7420746f6b656e00000000000000000000000000000000006064820152608401610510565b6000612200612159565b90506000815111612220576040518060200160405280600081525061224b565b8061222a84612504565b60405160200161223b929190612c91565b6040516020818303038152906040525b9392505050565b6001600160a01b0383166122ad576122a881609980546000838152609a60205260408120829055600182018355919091527f72a152ddfb8e864297c917af52ea6c1c68aead0fee1a62673fcc7e0c94979d000155565b6122d0565b816001600160a01b0316836001600160a01b0316146122d0576122d08382612653565b6001600160a01b0382166122ec576122e7816126f0565b610662565b826001600160a01b0316826001600160a01b0316146106625761066282826127c9565b600061231a82610b28565b905061232881600084611bd1565b6123336000836110a1565b6001600160a01b038116600090815260686020526040812080546001929061235c908490612d3b565b909155505060008281526067602052604080822080546001600160a01b0319169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6001600160a01b03821661240c5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610510565b6000818152606760205260409020546001600160a01b0316156124715760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610510565b61247d60008383611bd1565b6001600160a01b03821660009081526068602052604081208054600192906124a6908490612d0f565b909155505060008181526067602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b606081612545575060408051808201909152600181527f30000000000000000000000000000000000000000000000000000000000000006020820152610404565b8160005b811561256f578061255981612db9565b91506125689050600a83612d27565b9150612549565b60008167ffffffffffffffff81111561259857634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156125c2576020820181803683370190505b5090505b84156112a7576125d7600183612d3b565b91506125e4600a86612dd4565b6125ef906030612d0f565b60f81b81838151811061261257634e487b7160e01b600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061264c600a86612d27565b94506125c6565b6000600161266084610c42565b61266a9190612d3b565b6000838152609860205260409020549091508082146126bd576001600160a01b03841660009081526097602090815260408083208584528252808320548484528184208190558352609890915290208190555b5060009182526098602090815260408084208490556001600160a01b039094168352609781528383209183525290812055565b60995460009061270290600190612d3b565b6000838152609a60205260408120546099805493945090928490811061273857634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050806099838154811061276757634e487b7160e01b600052603260045260246000fd5b6000918252602080832090910192909255828152609a909152604080822084905585825281205560998054806127ad57634e487b7160e01b600052603160045260246000fd5b6001900381819060005260206000200160009055905550505050565b60006127d483610c42565b6001600160a01b039093166000908152609760209081526040808320868452825280832085905593825260989052919091209190915550565b82805461281990612d7e565b90600052602060002090601f01602090048101928261283b5760008555612881565b82601f1061285457805160ff1916838001178555612881565b82800160010185558215612881579182015b82811115612881578251825591602001919060010190612866565b5061288d9291506128c9565b5090565b50805461289d90612d7e565b6000825580601f106128af5750610904565b601f01602090049060005260206000209081019061090491905b5b8082111561288d57600081556001016128ca565b600067ffffffffffffffff808411156128f9576128f9612e14565b604051601f8501601f19908116603f0116810190828211818310171561292157612921612e14565b8160405280935085815286868601111561293a57600080fd5b858560208301376000602087830101525050509392505050565b80356001600160a01b038116811461040457600080fd5b60008083601f84011261297c578182fd5b50813567ffffffffffffffff811115612993578182fd5b6020830191508360208285010111156129ab57600080fd5b9250929050565b600082601f8301126129c2578081fd5b61224b838335602085016128de565b6000602082840312156129e2578081fd5b61224b82612954565b600080604083850312156129fd578081fd5b612a0683612954565b9150612a1460208401612954565b90509250929050565b600080600060608486031215612a31578081fd5b612a3a84612954565b9250612a4860208501612954565b9150604084013590509250925092565b60008060008060808587031215612a6d578081fd5b612a7685612954565b9350612a8460208601612954565b925060408501359150606085013567ffffffffffffffff811115612aa6578182fd5b8501601f81018713612ab6578182fd5b612ac5878235602084016128de565b91505092959194509250565b60008060408385031215612ae3578182fd5b612aec83612954565b915060208301358015158114612b00578182fd5b809150509250929050565b60008060408385031215612b1d578182fd5b612b2683612954565b946020939093013593505050565b600060208284031215612b45578081fd5b813561224b81612e2a565b600060208284031215612b61578081fd5b815161224b81612e2a565b60008060008060408587031215612b81578384fd5b843567ffffffffffffffff80821115612b98578586fd5b612ba48883890161296b565b90965094506020870135915080821115612bbc578384fd5b50612bc98782880161296b565b95989497509550505050565b600060208284031215612be6578081fd5b813567ffffffffffffffff811115612bfc578182fd5b6112a7848285016129b2565b600060208284031215612c19578081fd5b5035919050565b60008060408385031215612c32578182fd5b82359150602083013567ffffffffffffffff811115612c4f578182fd5b612c5b858286016129b2565b9150509250929050565b60008151808452612c7d816020860160208601612d52565b601f01601f19169290920160200192915050565b60008351612ca3818460208801612d52565b835190830190612cb7818360208801612d52565b01949350505050565b60006001600160a01b03808716835280861660208401525083604083015260806060830152612cf26080830184612c65565b9695505050505050565b60006020825261224b6020830184612c65565b60008219821115612d2257612d22612de8565b500190565b600082612d3657612d36612dfe565b500490565b600082821015612d4d57612d4d612de8565b500390565b60005b83811015612d6d578181015183820152602001612d55565b83811115610f725750506000910152565b600281046001821680612d9257607f821691505b60208210811415612db357634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415612dcd57612dcd612de8565b5060010190565b600082612de357612de3612dfe565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160e01b03198116811461090457600080fdfea2646970667358221220272a46e63d1c78d6a9f3b7304a5dffd42538caa5da8af399387d07539092391564736f6c63430008020033";

type MirroredERC721ConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: MirroredERC721ConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class MirroredERC721__factory extends ContractFactory {
  constructor(...args: MirroredERC721ConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<MirroredERC721> {
    return super.deploy(overrides || {}) as Promise<MirroredERC721>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): MirroredERC721 {
    return super.attach(address) as MirroredERC721;
  }
  override connect(signer: Signer): MirroredERC721__factory {
    return super.connect(signer) as MirroredERC721__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): MirroredERC721Interface {
    return new utils.Interface(_abi) as MirroredERC721Interface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): MirroredERC721 {
    return new Contract(address, _abi, signerOrProvider) as MirroredERC721;
  }
}
