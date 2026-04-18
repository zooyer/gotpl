#!/bin/bash

set -e

NAME="gotpl"
UPX_ENABLE=1

function setenv() {
    local goos=$1
    local goarch=$2
    local goenv=$3

    export GOOS="$goos"
    export GOARCH="$goarch"

    unset GOARM
    unset GOMIPS
    unset GOMIPS64
    unset GOAMD64

    if [ -z "$goenv" ]; then
        return
    fi

    case "$goarch" in
        arm)
            # GOARM=5/6/7
            export GOARM="$goenv"
            ;;

        mips|mipsle)
            # GOMIPS=softfloat/hardfloat
            export GOMIPS="$goenv"
            ;;

        mips64|mips64le)
            # GOMIPS64=softfloat/hardfloat
            export GOMIPS64="$goenv"
            ;;

        amd64)
            # GOAMD64=v1/v2/v3/v4
            export GOAMD64="$goenv"
            ;;

        arm64)
            # arm64 没有任何细分变量
            # 第三个参数会被忽略，这是刻意行为
            ;;
        *)
            echo "unsupported GOARCH: $goarch" >&2
            exit 1
            ;;
    esac
}

function build() {
    setenv "$@"

    local goos=$1
    local goarch=$2
    local goenv=$3

    local suffix=""
    if [ -n "$goenv" ]; then
        suffix="-$goenv"
    fi

    local exe="$NAME-$goos-$goarch$suffix"

    echo "==> build $exe"

    go build -ldflags "-s -w" -o "$exe"

    if [ "$UPX_ENABLE" = "1" ]; then
        upx -qq --brute "$exe" || true # UPX 对部分架构可能失败，这里不让它中断整个脚本
    fi
}

echo "Build info:"
echo "  NAME=$NAME"
echo "  UPX_ENABLE=$UPX_ENABLE"

rm -rf $NAME
rm -rf deploy

build linux arm 5
build linux arm 6
build linux arm 7
build linux arm64

mkdir "$NAME"
mv "$NAME-"* ./$NAME

tar -zcf ./$NAME.tgz ./$NAME/*

mv $NAME deploy
mv $NAME.tgz deploy