# 🧐 Running node guide(Testnet Sepolia)

---

## geth

1. Install geth. Consult [here](https://geth.ethereum.org/docs/getting-started/installing-geth)
2. Make your own geth data directory. In case of me, `./ethereum/execution/data`.
3. Create JWT token by prysm client, and place it into your own path. In case of me, `./ethereum/consensus/jwt.hex`. Consult [here](https://docs.prylabs.network/docs/install/install-with-script)
4. Config `./ethereum/execution/geth-config.toml`, and run geth. Just run:

```bash
geth --config ./ethereum/execution/geth-config.toml
```

5. Happy syncing!

---

## prysm

1. Install prysm. Consult [here](https://docs.prylabs.network/docs/install/install-with-script)
2. Make your own prysm data directory. In case of me, `./ethereum/consensus/data`.
3. Download `genesis.ssz` file from [here](https://github.com/eth-clients/merge-testnets/blob/main/sepolia/genesis.ssz).
4. Run prysm with your own config. We will leverage checkpoint sync for ease.

```bash
./prysm.sh beacon-chain
          --sepolia
          --execution-endpoint=http://localhost:8551 # geth default
          --jwt-secret=YOUR_OWN_JWT_SECRET_PATH
          --genesis-state=genesis.ssz
          --datadir=YOUR_OWN_DATA_DIRECTORY
          --checkpoint-sync-url=https://sepolia.beaconstate.info
          --genesis-beacon-api-url=https://sepolia.beaconstate.info
```

5. Happy syncing!

---

## erigon

[WIP]
