#!/bin/sh
dest=go
thrift --gen ${dest} ./video_info.thrift
