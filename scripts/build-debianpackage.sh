#!/bin/bash

architecture=""
version=""

while getopts ":a:v:" options; do
  case "${options}" in
    a)
      architecture=${OPTARG}
      ;;
    v)
      version=${OPTARG}
      ;;
    :)
      echo "Error: -${OPTARG} requires an argument"
      exit 1
      ;;
    *)
      exit 1
      ;;
  esac
done

echo "=> Creating debian package folder structure"
mkdir -p "./deb/sumocli_$version-1_$architecture/usr/bin"
echo "=> Copying sumocli binary"
chmod +x "./sumocli_linux_$architecture/sumocli/sumocli"
cp "./sumocli_linux_$architecture/sumocli/sumocli" "./deb/sumocli_$version-1_$architecture/usr/bin"
echo "=> Creating debian control file"
mkdir -p "./deb/sumocli_$version-1_$architecture/DEBIAN"
cat > "./deb/sumocli_$version-1_$architecture/DEBIAN/control" << EOF
Package: sumocli
Version: $version
Maintainer: kyle@thepublicclouds.com
Architecture: $architecture
Homepage: https://github.com/SumoLogic-Labs/sumocli
Description: Sumocli is a CLI application written in Go that allows you to manage your Sumo Logic tenancy from the command line.
EOF
echo "=> Building debian package"
dpkg --build "./deb/sumocli_$version-1_$architecture"
