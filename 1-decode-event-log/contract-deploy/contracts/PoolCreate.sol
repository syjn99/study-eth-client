// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

contract PoolCreate {
    mapping(uint24 => int24) public feeAmountTickSpacing;

    constructor() {
        feeAmountTickSpacing[500] = 10;
        feeAmountTickSpacing[3000] = 60;
        feeAmountTickSpacing[10000] = 200;
    }

    event PoolCreated(
        address indexed token0,
        address indexed token1,
        uint24 indexed fee,
        int24 tickSpacing,
        address pool
    );

    function createPool(
        address tokenA,
        address tokenB,
        uint24 fee
    ) external returns (address pool) {
        require(tokenA != tokenB);
        (address token0, address token1) = tokenA < tokenB
            ? (tokenA, tokenB)
            : (tokenB, tokenA);
        require(token0 != address(0));
        int24 tickSpacing = feeAmountTickSpacing[fee];
        require(tickSpacing != 0);
        pool = address(1234);
        emit PoolCreated(token0, token1, fee, tickSpacing, pool);
    }
}
