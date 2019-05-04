#!/usr/bin/env bash

set -e

if [ ! -f run.sh ]; then
    echo 'run.sh must be run within its container folder' 1>&2
    exit 1
fi

if [ ! -d log ]; then
	mkdir log
fi

if [ ! -d pid ]; then
	mkdir pid
fi

DIRPWD=`pwd`

export GOPATH=$DIRPWD

cd src/server

go run main.go  >> $DIRPWD/log/panic.log 2>&1 &

echo "finished"