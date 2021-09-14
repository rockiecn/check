#!/bin/bash
solc  ~/work/check/cash/cash.sol  --abi  --overwrite  -o  ~/work/check/cash/

solc  ~/work/check/cash/cash.sol  --bin  --overwrite  -o  ~/work/check/cash/

abigen  --bin=/home/wy/work/check/cash/Cash.bin  --abi=/home/wy/work/check/cash/Cash.abi  --pkg=cash  --out=/home/wy/work/check/cash/cash.go
