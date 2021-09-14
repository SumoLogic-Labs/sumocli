#!/bin/sh

architecture="arm64"
build_number="DEV"
version_number="DEV"

echo "=> Installing gox"

while getopts "ab:lmwv:" opt; do
  case $opt in
    a)
      architecture=$OPTARG
      ;;
    b)
      build_number=$OPTARG
      ;;
    l)
      echo "$build_number"
      echo "$version_number"
      echo "linux" >&2
      ;;
    m)
      echo "macos" >&2
      ;;
    w)
      echo "windows" >&2
      pwd
      GOOS=windows go build --ldflags "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version_number' -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build_number'" ./cmd/sumocli
      ;;
    v)
      version_number=$OPTARG
      ;;
    \?)
      echo "Invalid option!" >&2
      ;;
  esac
done

#echo "=> Downloading go dependencies"
#go mod download

#if [ "$DEV" = "true" ]; then
  #echo "=> Building sumocli"
  #gox -osarch="linux/amd64" ./cmd/sumocli
  #ls -lah
  #mv "$GOPATH/src/github.com/SumoLogic-Incubator/sumocli/sumocli_linux_amd64" "$GOPATH/bin/sumocli"
#fi
