# Task 1

- Goerli Testnet에 간단한 Smart Contract 배포

Contract Address: [0xD40F0EeeAfb58C8d18E6b0469EFF02f0B737583d](https://goerli.etherscan.io/address/0xD40F0EeeAfb58C8d18E6b0469EFF02f0B737583d)

Contract Creation Tx Hash: [0xd845fe788da06e7df00eb0549084a6ba9971487489a8742263b8bba2e40ad675](https://goerli.etherscan.io/tx/0xd845fe788da06e7df00eb0549084a6ba9971487489a8742263b8bba2e40ad675)

- Contract의 Method 실행해서 Event Emit하기

```dotnetcli
event PoolCreated(
  address indexed token0,
  address indexed token1,
  uint24 indexed fee,
  int24 tickSpacing,
  address pool
);
```

Send Tx(`createPool()`) Hash: [0x802e69f2a119de3aa7e25f086934c4fbade63b14ab94ade98cd6bd914dc5ea34](https://goerli.etherscan.io/tx/0x802e69f2a119de3aa7e25f086934c4fbade63b14ab94ade98cd6bd914dc5ea34)

- Emit 된 Event를 golang으로 로컬에서 받아와 Decoding하기