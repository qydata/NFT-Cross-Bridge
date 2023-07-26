const hre = require("hardhat");
async function migration() {
    let contractAddress = '0xAE2d7aF8c407503Ebf74f826D600CF29B7a1dF1a'
    await hre.run("verify:verify", {
        address: contractAddress,
        constructorArguments: [],
    });
};
migration();
