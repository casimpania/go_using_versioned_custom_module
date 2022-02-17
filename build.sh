#!/bin/bash

## Define GOPATH env
export GOPATH=`pwd`

## Get the module(s)
cd src
go mod tidy

## Announce next step...
echo '---------------------------------------------'
echo 'Try and run the code (from within "src"):    '
echo '  go run gwashington.go'
echo '---------------------------------------------'
