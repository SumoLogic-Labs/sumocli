#!/bin/sh

echo "=> Downloading go dependencies"
go mod download

echo "=> Installing gox"
go get github.com/mitchellh/gox

if [ "$DEV" = "true" ]; then
  echo "=> Building sumocli"
  gox -osarch="linux/amd64" ./cmd/sumocli
  ls -lah
  mv "$GOPATH/src/github.com/wizedkyle/sumocli/sumocli_linux_amd64" "$GOPATH/bin/sumocli"
fi
