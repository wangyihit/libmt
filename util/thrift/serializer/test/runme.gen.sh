#!/bin/sh
dest=go
thrift --gen ${dest} ./test.thrift
