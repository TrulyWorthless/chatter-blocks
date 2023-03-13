#!/bin/sh
solc --abi --bin contracts/SimpleChannel.sol -o build
abigen --bin=build/SimpleChannel.bin --abi=build/SimpleChannel.abi --pkg=store --out=bindings/SimpleChannel.go