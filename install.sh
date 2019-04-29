#!/usr/bin/env bash

if [ ! -f install.sh ]; then
    echo 'install must be run within its container folder' 1>&2
    exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

if [ ! -d log ]; then
	mkdir log
fi

gofmt -w -s src

go install goblog

export GOPATH="$OLDGOPATH"

echo 'finished'