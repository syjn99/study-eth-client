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
[0xaa80c0a50eba709510f07bcbfc3be3632ece652133d539cd86b541276e0b7702](https://goerli.etherscan.io/tx/0xaa80c0a50eba709510f07bcbfc3be3632ece652133d539cd86b541276e0b7702)

- Emit 된 Event를 golang으로 로컬에서 받아와 Decoding하기

### Event Subscription

Event Subscription이 아래와 같이 가능합니다. 구독 중에 Tx를 보내고 receipt을 받는 순간 아래 로그가 나옵니다.

```dotnetcli
go run ./1-decode-event-log/decode-log/event-subscribe/main.go
```

Result:
```dotnetcli
{0xD40F0EeeAfb58C8d18E6b0469EFF02f0B737583d [0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118 0x0000000000000000000000001234567890123456789012345678901234567890 0x0000000000000000000000009234567890123456789012345678901234567890 0x00000000000000000000000000000000000000000000000000000000000001f4] [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 10 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 4 210] 8889365 0xaa80c0a50eba709510f07bcbfc3be3632ece652133d539cd86b541276e0b7702 73 0xa5359ee415d3645011a05f3e7fe8a0dcd0005c31e34eebd6380a524c550e4952 182 false}
```

### Read Event Logs

```dotnetcli
go run ./1-decode-event-log/decode-log/event-read/main.go
```

Result:
```dotnetcli
Block Number: 8889365
TxHash:  0xaa80c0a50eba709510f07bcbfc3be3632ece652133d539cd86b541276e0b7702
token0 address:  0x0000000000000000000000001234567890123456789012345678901234567890
token1 address:  0x0000000000000000000000009234567890123456789012345678901234567890
fee:  0x00000000000000000000000000000000000000000000000000000000000001f4
tickSpacing:  10
pool address:  0x00000000000000000000000000000000000004d2
First Topic and hash are equal
First Topic:  0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118
Hash:  0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118
```