#!/bin/bash

#auto gen .go from .sol

solc  ~/work/check/internal/cash/cash.sol  --abi  --overwrite  -o  ~/work/check/internal/cash/
solc  ~/work/check/internal/cash/cash.sol  --bin  --overwrite  -o  ~/work/check/internal/cash/
abigen  --bin=/home/wy/work/check/internal/cash/Cash.bin  --abi=/home/wy/work/check/internal/cash/Cash.abi  --pkg=cash  --out=/home/wy/work/check/internal/cash/cash.go

rm /home/wy/work/check/internal/cash/*.bin
rm /home/wy/work/check/internal/cash/*.abi
