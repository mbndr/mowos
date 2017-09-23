#!/bin/bash
# compiles mowos-agent and mowos-monitor for all supported platforms
# currently only platforms I need

version="0.0.0-unknown"
rootPath="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/.."
binPath="$rootPath/bin"

# for now use git hash as version (dev)
version="$(git rev-parse --short HEAD)"

# compile agent and monitor
# $1 GOOS
# $2 GOARCH
# $3 output dir
compile() {
    mkdir -p $3
    GOOS=$1 GOARCH=$2 go build -ldflags "-X github.com/mbndr/mowos/agent.Version=$version" -o "$3/mowos-agent" "$rootPath/cmd/agent.go"
    GOOS=$1 GOARCH=$2 go build -ldflags "-X github.com/mbndr/mowos/monitor.Version=$version" -o "$3/mowos-monitor" "$rootPath/cmd/monitor.go"
}

# clean
rm -rf $binPath

echo "cross compiling mowos"

# Linux
echo "linux 64"
compile "linux" "amd64" "$binPath/mowos-linux64"

echo "linux 32"
compile "linux" "386" "$binPath/mowos-linux32"

echo "linux arm"
compile "linux" "arm" "$binPath/mowos-linux-amd"
