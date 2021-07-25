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

const challengePeriod = 60
const startingStake = 1

task("deploySink", "Deploy and initialize a sink and friends")
  .setAction(async () => {
    const tokenFactory = await hre.ethers.getContractFactory("TestERC20");
    const token = await tokenFactory.deploy('TestToken', 'TTK');
    console.log("Deploying token...", token.deployTransaction.hash);
    await token.deployTransaction.wait();

    const oracleFactory = await hre.ethers.getContractFactory("Oracle");
    const oracle = await oracleFactory.deploy();
    console.log("Deploying oracle...", oracle.deployTransaction.hash);
    await oracle.deployTransaction.wait();

    const sinkFactory = await hre.ethers.getContractFactory("ERC20Sink");
    const sink = await sinkFactory.deploy(sinkFactory.signer.address, oracle.address, token.address);
    console.log("Deploying sink...", sink.deployTransaction.hash);
    await sink.deployTransaction.wait();

    const oracleInit = await oracle.init(challengePeriod, startingStake, token.address, token.address, sink.address);
    console.log("Initializing oracle...", oracleInit.hash);
    await oracleInit.wait()

    console.log("Deployment complete");
    console.log("Sink:", sink.address);
    console.log("Oracle:", oracle.address);
    console.log("Test token:", token.address);
  });

task("deploySource", "Deploy and initialize a source and friends")
  .setAction(async () => {
    const tokenFactory = await hre.ethers.getContractFactory("TestERC20");
    const token = await tokenFactory.deploy("SourceToken", "SST");
    console.log("Deploying token...", token.deployTransaction.hash);
    await token.deployTransaction.wait();

    const oracleFactory = await hre.ethers.getContractFactory("Oracle");
    const oracle = await oracleFactory.deploy();
    console.log("Deploying oracle...", oracle.deployTransaction.hash);
    await oracle.deployTransaction.wait();

    const sourceFactory = await hre.ethers.getContractFactory("ERC20Source");
    const source = await sourceFactory.deploy(sourceFactory.signer.address, oracle.address, token.address);
    console.log("Deploying source...", source.deployTransaction.hash);
    await source.deployTransaction.wait();

    const setMinterTx = await token.setMinter(source.address);
    console.log("Setting source as minter...", setMinterTx.hash);
    await setMinterTx.wait();

    const oracleInit = await oracle.init(challengePeriod, startingStake, token.address, token.address, source.address);
    console.log("Initializing oracle...", oracleInit.hash);
    await oracleInit.wait()

    console.log("Deployment complete");
    console.log("Source:", source.address);
    console.log("Oracle:", oracle.address);
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

task("initRelayer", "initialize a relayer for testing")
  .addParam("source", "source contract address")
  .addParam("oracle", "oracle contract address")
  .setAction(async (args) => {
    const source = (await hre.ethers.getContractFactory("ERC20Source")).attach(args.source)
    const token = (await hre.ethers.getContractFactory("SourceToken")).attach(await source.token())
    const address = (await hre.ethers.getSigners())[0].address
    const mintTx = await token.mint(address, "1000000000000000000000")
    await mintTx.wait()

    const oracle = (await hre.ethers.getContractFactory("Oracle")).attach(await args.oracle)
    const allowTx = await token.approve(args.oracle, "1000000000000000000000")
    await allowTx.wait()
  });

task("startTransfer", "start a transfer from sink to source")
  .addParam("sink", "address of the sink")
  .addParam("source", "address of the source")
  .addParam("chainid", "chain id of the source", undefined, types.int)
  .addParam("receiver", "receiver address")
  .setAction(async (args) => {
      const account = (await hre.ethers.getSigners())[0].address
      const transferValue = 123456

      const sink = await (await hre.ethers.getContractFactory("ERC20Sink")).attach(args.sink)
      const tokenAddress = await sink.token()
      const token = await (await hre.ethers.getContractFactory("TestERC20")).attach(tokenAddress)

      console.log("Minting tokens to user...", account)
      const mint = await token.mint(account, transferValue)
      console.log("Allowing transfer from sink...")
      const allow = await token.approve(args.sink, transferValue)
      await Promise.all([mint.wait(), allow.wait()])

      console.log("Requesting transfer on sink...")
      const requestTransfer = await sink.requestTransfer([args.chainid, args.source], args.receiver, transferValue)
      await requestTransfer.wait()
      console.log("done")
  })

task("queryBalance")
  .addParam("token")
  .addParam("address")
  .setAction(async (args) => {
    const token = await (await hre.ethers.getContractFactory("TestERC20")).attach(args.token)
    console.log("Current balance of", args.address, (await token.balanceOf(args.address)).toString())
  });

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
