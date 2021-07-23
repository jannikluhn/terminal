const { expect } = require("chai");

describe("Oracle", function () {

  let oracle
  let token
  const challengePeriod = 60*60
  const startingStake = 100
  const transferIdentifier = "0x57db4c0a87b1c702638c18e4134aa0c52062a7bf7430a30a3202fef8c34db53c"
  const transferValue = 123

  let requestInitiator
  let requestDenier

  before(async () => {
    const tokenFactory = await ethers.getContractFactory("TestERC20")
    token = await tokenFactory.deploy('TestToken', 'TTK');
    await token.deployed();

    const oracleFactory = await ethers.getContractFactory("Oracle");
    oracle = await oracleFactory.deploy(challengePeriod, startingStake, token.address, token.address);
    await oracle.deployed();

    [requestInitiator, requestDenier] = await ethers.getSigners();

    // arbitrary token amount for testing
    tokenAmount = startingStake * 100
    await token.mint(requestInitiator.address, tokenAmount)
    await token.mint(requestDenier.address, tokenAmount)
    await token.connect(requestInitiator).approve(oracle.address, tokenAmount)
    await token.connect(requestDenier).approve(oracle.address, tokenAmount)
  })

  it("Should initiate a request", async function () {

    await oracle.connect(requestInitiator).startRequest(transferIdentifier, transferValue, requestInitiator.address)

    const [
      legit,
      value,
      lastStake,
      lastChallengeTime,
      receiver,
      lastModifier
    ] = await oracle.getRequest(transferIdentifier)

    expect(legit).to.equal(true)
    expect(value).to.equal(transferValue)
    expect(lastStake).to.equal(startingStake)
    expect(lastChallengeTime).to.equal((await ethers.provider.getBlock()).timestamp)
    expect(receiver).to.equal(requestInitiator.address)
    expect(lastModifier).to.equal(requestInitiator.address)
  });
});
