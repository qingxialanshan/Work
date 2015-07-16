#!/bin/bash

# generate the binary for PageAction on 4 platforms

GOPATH=`pwd`
if [ ! -d $GOPATH/bin/windows ]
then
    mkdir -p $GOPATH/bin/windows
fi

if [ ! -d $GOPATH/bin/linux-x86 ]
then
    mkdir -p $GOPATH/bin/linux-x86
fi

if [ ! -d $GOPATH/bin/linux-x64 ]
then
    mkdir -p $GOPATH/bin/linux-x64
fi

if [ ! -d $GOPATH/bin/osx ]
then
    mkdir -p $GOPATH/bin/osx
fi
export GOROOT=/home/amyl/software/go

export GOPATH=$GOPATH

CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build PageAction.go pageMap.go
mv PageAction.exe $GOPATH/bin/windows/

CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build PageAction.go pageMap.go
mv PageAction $GOPATH/bin/linux-x64/

CGO_ENABLE=0 GOOS=linux GOARCH=386 go build PageAction.go pageMap.go
mv PageAction $GOPATH/bin/linux-x86/

CGO_ENABLE=0 GOOS=darwin GOARCH=amd64 go build PageAction.go pageMap.go
mv PageAction $GOPATH/bin/osx/
