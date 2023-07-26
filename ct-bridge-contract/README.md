# avax-evm-compatible-bridge-contr

### Prepare configuration

Please refer to example env file

```shell
$ cat .env.example > .env
```

### Install Dependencies

```shell
$ nvm use
$ npm ci
```

### Deploy local chains to test the contracts locally

Need to have 3 terminals to run the scripts simultaneously

```shell
$ npm run chain1
```

```shell
$ npm run chain2
```

After 2 chains are running, deploy ERC721SwapAgent and ERC1155SwapAgent contracts

```shell
$ npm run chain1:erc721-deploy-agent
$ npm run chain1:erc1155-deploy-agent
$ npm run chain2:erc721-deploy-agent
$ npm run chain2:erc1155-deploy-agent
```

There are other useful commands to deploy contracts, see them in the package.json

### Deploy to CT and COO Test

```shell
$ npm run cootest:erc721-deploy-agent
$ npm run cootest:erc1155-deploy-agent
$ npm run ct:erc721-deploy-agent
$ npm run ct:erc1155-deploy-agent
```
 
The addresses of the deployed contracts will be in the chains folder which has chain ids as subfolders.
