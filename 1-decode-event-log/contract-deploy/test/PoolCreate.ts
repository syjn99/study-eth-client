import { expect } from 'chai';
import { ethers } from 'hardhat';

describe("PoolCreate", function() {
  it("Should emit PoolCreate Event", async function() {
    const token1 = "0x1234567890123456789012345678901234567890";
    const token2 = "0x9234567890123456789012345678901234567890";
    const fee = 500;

    const PoolCreate = await ethers.getContractFactory("PoolCreate");
    const poolCreate = await PoolCreate.deploy();
    // Write token1 and token2 dummy address

    await poolCreate.deployed();

    const tx = await poolCreate.createPool(token1, token2, fee);
    const receipt = await tx.wait();
    const event = receipt.events && receipt.events[0];
    console.log(event);
    if (event) {
      expect(event.event).to.equal("PoolCreated");
    }
  }
  )
});