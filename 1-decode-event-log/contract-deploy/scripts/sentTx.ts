import { ethers } from "hardhat";

async function main() {
  const token1 = "0x1234567890123456789012345678901234567890";
  const token2 = "0x9234567890123456789012345678901234567890";
  const fee = 500;

  const contractAddress = "0xD40F0EeeAfb58C8d18E6b0469EFF02f0B737583d";
  const poolCreate = await ethers.getContractAt("PoolCreate", contractAddress);

  const createPool = await poolCreate.createPool(token1, token2, fee);
  console.log(createPool.hash);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
