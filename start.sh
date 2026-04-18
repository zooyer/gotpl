#!/bin/sh
cd "$(dirname "$0")" || exit
./gotpl -nohup -config conf/test.yaml &