#!/bin/bash

export CURDIR=$PWD
export GOPATH=$CURDIR/../../../../

cd $CURDIR/tests/c/bin
./ctest randpos
