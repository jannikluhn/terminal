const { task } = require("hardhat/config");

require("@nomiclabs/hardhat-waffle");

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

task("deploySink", "Deploy and initialize a sink and friends")
  .setAction(async () => {
    const tokenFactory = await hre.ethers.getContractFactory("TestERC20");
    const token = await tokenFactory.deploy('TestToken', 'TTK');
    console.log("Deploying token...", token.deployTransaction.hash);
    await token.deployTransaction.wait();

    const oracle = (await hre.ethers.getSigners())[0].address; // TODO: deploy

    const sinkFactory = await hre.ethers.getContractFactory("ERC20Sink");
    const sink = await sinkFactory.deploy(sinkFactory.signer.address, oracle, token.address);
    console.log("Deploying sink...", sink.deployTransaction.hash);
    await sink.deployTransaction.wait();

    console.log("Deployment complete");
    console.log("Sink:", sink.address);
    console.log("Oracle:", oracle);
    console.log("Test token:", token.address);
  });

task("deploySource", "Deploy and initialize a source and friends")
  .setAction(async () => {
    const tokenFactory = await hre.ethers.getContractFactory("SourceToken");
    const token = await tokenFactory.deploy("SourceToken", "SST");
    console.log("Deploying token...", token.deployTransaction.hash);
    await token.deployTransaction.wait();

    const oracle = (await hre.ethers.getSigners())[0].address; // TODO: deploy

    const sourceFactory = await hre.ethers.getContractFactory("ERC20Source");
    const source = await sourceFactory.deploy(sourceFactory.signer.address, oracle, token.address);
    console.log("Deploying source...", source.deployTransaction.hash);
    await source.deployTransaction.wait();

    const setMinterTx = await token.setMinter(source.address);
    console.log("Setting source as minter...", setMinterTx.hash);
    await setMinterTx.wait();

    console.log("Deployment complete");
    console.log("Source:", source.address);
    console.log("Oracle:", oracle);
    console.log("Source token:", token.address);
  })

task("connectEndpoint", "Connect a sink or a source to another sink or source")
  .addParam("contract", "address of the contract to configure")
  .addParam("chainid", "chain id of the endpoint to connect to", undefined, types.int)
  .addParam("address", "contract address of the endpoint to connect to")
  .setAction(async (args) => {
    const factory = await hre.ethers.getContractFactory("HasConnectedEndpoints");
    const contract = await factory.attach(args.contract);
    const tx = await contract.connectToEndpoint([args.chainid, args.address]);
    console.log("waiting for tx...", tx.hash);
    await tx.wait();
    console.log("Done");
  })

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: {
    version: "0.8.6",
    optimizer: {
      enabled: true,
    },
  },

  networks: {
    localsink: {
      url: "http://localhost:8545",
      accounts: ["b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773"],
    },
    localsource: {
      url: "http://localhost:8555",
      accounts: ["b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3774"],
    },
  }
};
