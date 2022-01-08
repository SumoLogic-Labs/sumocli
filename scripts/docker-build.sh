#!/bin/sh

build="DEV"
version="DEV"

while getopts "bdv" opt; do
  case $opt in
    b)
      build=$OPTARG
      ;;
    d)
      echo "$build"
      echo "$version"
      echo "=> Compiling linux AMD64 binary for Docker image" >&2
      GOOS=linux GOARCH=amd64 go build -ldflags \
     "-X 'github.com/SumoLogic-Labs/sumocli/internal/build.Version=$version'
      -X 'github.com/SumoLogic-Labs/sumocli/internal/build.Build=$build'" \
      ./cmd/sumocli
      echo "=> Moving sumocli binary to $GOPATH/bin/"
      mv "$GOPATH/src/github.com/SumoLogic-Labs/sumocli/sumocli" "$GOPATH/bin/sumocli"
      ;;
    v)
      version=$OPTARG
      ;;
    \?)
      echo "Invalid option!" >&2
      ;;
  esac
done
