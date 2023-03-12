#!/bin/sh
solc --abi --bin contracts/HelloWorld.sol -o build
abigen --bin=build/Store.bin --abi=build/Store.abi --pkg=store --out=bindings/Store.go
go run cmd/web3/*.go