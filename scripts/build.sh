#!/bin/sh

build="DEV"
maintainer="kyle@thepublicclouds.com"
version="DEV"

while getopts "b:dlmv:w" opt; do
  case $opt in
    b)
      build=$OPTARG
      ;;
    d)
      echo "=> Compiling linux AMD64 binary for Docker image" >&2
      GOOS=linux GOARCH=amd64 go build -ldflags \
      "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
      -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" \
      ./cmd/sumocli
      echo "=> Moving sumocli binary to $GOPATH/bin/"
      mv "$GOPATH/src/github.com/SumoLogic-Incubator/sumocli/sumocli" "$GOPATH/bin/sumocli"
      ;;
    l)
      echo "=> Compiling linux AMD64 binary" >&2
      GOOS=linux GOARCH=amd64 go build -ldflags \
      "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
      -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" \
      ./cmd/sumocli
      if [ "$version" != "DEV" ] && [ "$build" != "DEV" ]; then
        echo "=> Creating Deb package"
          mkdir -p ~/deb/sumocli_$version-1_amd64/usr/bin
          cp sumocli ~/deb/sumocli_$version-1_amd64/usr/bin
          echo "=> Creating DEBIAN control file"
          mkdir -p ~/deb/sumocli_$version-1_amd64/DEBIAN
          cat > ~/deb/sumocli_$version-1_amd64/DEBIAN/control <<EOL
          Package: sumocli
          Version: $version
          Maintainer: $maintainer
          Architecture: amd64
          Homepage: https://github.com/SumoLogic-Incubator/sumocli
          Description: Sumocli is a CLI application written in Go that allows you to manage your Sumo Logic tenancy from the command line.
          EOL
          echo "=> Building Deb package"
          dpkg --build ~/deb/sumocli_$version-1_amd64
      fi
      ;;
    m)
      echo "=> Compiling macOS AMD64 binary" >&2
      GOOS=darwin GOARCH=amd64 go build -ldflags \
        "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
        -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" \
        ./cmd/sumocli

      echo "=> Compiling macOS ARM64 binary" >&2
      GOOS=darwin GOARCH=arm64 go build -ldflags \
        "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
        -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" \
        ./cmd/sumocli
      ;;
    w)
      echo "=> Compiling Windows AMD64 binary" >&2
      GOOS=windows GOARCH=amd64 go build -ldflags \
      "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
      -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" \
      ./cmd/sumocli
      ;;
    v)
      version=$OPTARG
      ;;
    \?)
      echo "Invalid option!" >&2
      ;;
  esac
done
