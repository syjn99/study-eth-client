import { ethers } from "hardhat";

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log(`Deploying contracts with the account: ${deployer.address}`);
  console.log(`Account balance: ${(await deployer.getBalance()).toString()}`);

  const PoolCreate = await ethers.getContractFactory("PoolCreate");
  const poolCreate = await PoolCreate.deploy();

  console.log(`PoolCreate deployed to ${poolCreate.address}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
