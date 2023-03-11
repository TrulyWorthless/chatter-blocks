#!/bin/sh
solc --abi --bin solidity/helloworld.sol -o build
go run cmd/web3/*.go