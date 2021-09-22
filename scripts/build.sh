#!/bin/sh

build="DEV"
keyvaulturl=""
keyvaultclientid=""
keyvaultclientsecret=""
keyvaultcertificate=""
maintainer="kyle@thepublicclouds.com"
version="DEV"

args=(
  "version"
)
opts=$(getopt --long "$(printf "%s:," "${args[@]}")" -- "$@")
eval set -- "$opts"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --version)
      echo "we are setting the version $2"
      shift 2
      ;;
    *)
      echo "Its all broken"
      exit 1
      ;;
  esac
done


#while getopts "b:dlmv:w" opt; do
#  case $opt in
#    b)
#      build=$OPTARG
#      ;;
#    d)
#      echo "=> Compiling linux AMD64 binary for Docker image" >&2
#      GOOS=linux GOARCH=amd64 go build -ldflags \
#     "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
#      -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" \
#      ./cmd/sumocli
#      echo "=> Moving sumocli binary to $GOPATH/bin/"
#      mv "$GOPATH/src/github.com/SumoLogic-Incubator/sumocli/sumocli" "$GOPATH/bin/sumocli"
#      ;;
#    l)
#      echo "=> Compiling linux AMD64 binary" >&2
#      GOOS=linux GOARCH=amd64 go build -ldflags \
#      "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
#      -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" \
#      ./cmd/sumocli
#      if [ "$version" != "DEV" ] && [ "$build" != "DEV" ]; then
#        echo "=> Creating Deb package"
#          mkdir -p ~/deb/sumocli_$version-1_amd64/usr/bin
#          cp sumocli ~/deb/sumocli_$version-1_amd64/usr/bin
#          echo "=> Creating DEBIAN control file"
#          mkdir -p ~/deb/sumocli_$version-1_amd64/DEBIAN
#          cat > ~/deb/sumocli_$version-1_amd64/DEBIAN/control <<EOF
#Package: sumocli
#Version: $version
#Maintainer: $maintainer
#Architecture: amd64
#Homepage: https://github.com/SumoLogic-Incubator/sumocli
#Description: Sumocli is a CLI application written in Go that allows you to manage your Sumo Logic tenancy from the command line.
#EOF
#          echo "=> Building Deb package"
#          dpkg --build ~/deb/sumocli_$version-1_amd64
#      fi
#      ;;
#    m)
#      echo "=> Compiling macOS AMD64 binary" >&2
#      GOOS=darwin GOARCH=amd64 go build -ldflags \
#        "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
#        -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" \
#        ./cmd/sumocli
#
#      echo "=> Compiling macOS ARM64 binary" >&2
#      GOOS=darwin GOARCH=arm64 go build -ldflags \
#        "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
#        -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" \
#        ./cmd/sumocli
#      ;;
#    w)
#      echo "=> Compiling Windows AMD64 binary" >&2
#      GOOS=windows GOARCH=amd64 go build -ldflags \
#      "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
#      -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" \
#      ./cmd/sumocli
#      if [ "$version" != "DEV" ] && [ "$build" != "DEV" ]; then
#        echo "=> Installing azuresigntool"
#        dotnet tool install --global AzureSignTool --version 2.0.17
#        echo "=> Signing Windows binary with Azure Key Vault"
#        azuresigntool sign --description-url "https://github.com/SumoLogic-Incubator/sumocli" --file-digest sha256 \
#          --azure-key-vault-url "$keyvaulturl" \
#          --azure-key-vault-client-id "$keyvaultclientid" \
#          --azure-key-vault-client-secret "$keyvaultclientsecret" \
#          --azure-key-vault-certificate "$keyvaultcertificate" \
#          --timestamp-rfc3161 http://timestamp.sectigo.com \
#          --timestamp-digest sha256 \
#          sumocli.exe
#        echo "=> Creating zip archive"
#        zip -r sumocli-windows-amd64.zip sumocli.exe
#      fi
#      ;;
#    v)
#      version=$OPTARG
#      ;;
#    \?)
#      echo "Invalid option!" >&2
#      ;;
#  esac
#done
