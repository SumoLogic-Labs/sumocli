#!/bin/bash

releaseArchitectures=""
version=""

generate_hash() {
  hashname=$1
  hashcmd=$2
  echo "$hashname:"
  for file in $(find -type f); do
    file=$(echo "$file" | cut -c3-)
    if [ "$file" = "Release" ]; then
      continue
    fi
    echo " $(${hashcmd} "$file" | cut -d" " -f1) $(wc -c "$file")"
  done
}

while getopts ":a:v:" options; do
  case "${options}" in
    a)
      architectures+=("$OPTARG")
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

echo "=> Creating apt repo folder"
mkdir ./aptsumocli
echo "=> Syncing S3 bucket locally"
aws s3 sync s3://aptsumocli ./aptsumocli
echo "=> Checking pools directory"
if [ -d "./aptsumocli/pool/main" ]; then
  echo "=> ./aptsumocli/pool/main already exists"
else
  echo "=> Creating pools directory"
  mkdir -p ./aptsumocli/pool/main
fi
for architecture in "${architectures[@]}"; do
  releaseArchitectures+=$architecture
  releaseArchitectures+=" "
  echo "=> Moving $architecture debian package to local apt repo"
  mv "./sumocli_$version-1_$architecture.deb" "./aptsumocli/pool/main/sumocli_$version-1_$architecture.deb"
  echo "=> Checking for $architecture packages directory"
  if [ -d "./aptsumocli/dists/stable/main/binary-$architecture" ]; then
    echo "=> ./aptsumocli/dists/stable/main/binary-$architecture already exists"
  else
    mkdir -p "./aptsumocli/dists/stable/main/binary-$architecture"
  fi
  echo "=> Checking for old $architecture package files"
  if [ -f "./aptsumocli/dists/stable/main/binary-$architecture/Packages" ]; then
    echo "=> Removing ./aptsumocli/dists/stable/main/binary-$architecture/Packages"
    rm "./aptsumocli/dists/stable/main/binary-$architecture/Packages"
  else
    echo "=> ./aptsumocli/dists/stable/main/binary-$architecture/Packages does not exist"
  fi
  if [ -f "./aptsumocli/dists/stable/main/binary-$architecture/Packages.gz" ]; then
    echo "=> Removing ./aptsumocli/dists/stable/main/binary-$architecture/Packages.gz"
    rm "./aptsumocli/dists/stable/main/binary-$architecture/Packages.gz"
  else
    echo "=> ./aptsumocli/dists/stable/main/binary-$architecture/Packages.gz does not exist"
  fi
  echo "=> Generate new $architecture package file"
  cd ./aptsumocli || exit
  dpkg-scanpackages --arch "$architecture" pool/ > "Packages"
  cd ..
  mv -f "./aptsumocli/Packages" "./aptsumocli/dists/stable/main/binary-$architecture/Packages"
  echo "=> Compressing $architecture package file"
  gzip -k "./aptsumocli/dists/stable/main/binary-$architecture/Packages"
done
echo "=> Checking for old release files"
if [ -f "./aptsumocli/dists/stable/Release" ]; then
  echo "=> Removing ./aptsumocli/dists/stable/Release"
  rm ./aptsumocli/dists/stable/Release
else
  echo "=> ./aptsumocli/dists/stable/Release does not exist"
fi
if [ -f "./aptsumocli/dists/stable/Release.gpg" ]; then
  echo "=> Removing ./aptsumocli/dists/stable/Release.gpg"
  rm ./aptsumocli/dists/stable/Release.gpg
else
  echo "=> ./aptsumocli/dists/stable/Release.gpg does not exist"
fi
if [ -f "./aptsumocli/dists/stable/InRelease" ]; then
  echo "=> Removing ./aptsumocli/dists/stable/InRelease"
  rm ./aptsumocli/dists/stable/InRelease
else
  echo "=> ./aptsumocli/dists/stable/InRelease does not exist"
fi
echo "=> Moving directories to ./aptsumocli/dists/stable"
cd ./aptsumocli/dists/stable || exit
cat << EOF > Release
Origin: apt.sumocli.app
Suite: stable
Codename: stable
Version: $version
Architectures: $releaseArchitectures
Components: main
Description: Sumocli is a CLI application written in Go that allows you to manage your Sumo Logic tenancy from the command line.
Date: $(date -Ru)
$(generate_hash "MD5Sum" "md5sum")
$(generate_hash "SHA1" "sha1sum")
$(generate_hash "SHA256" "sha256sum")
EOF
echo "=> Signing release file"
cat ./Release | gpg --default-key "Kyle Jackson" -abs > Release.gpg
echo "=> Creating InRelease file"
cat ./Release | gpg --default-key "Kyle Jackson" -abs --clearsign > InRelease
echo "=> Moving back to root directory"
cd -
echo "=> Syncing local apt repo to S3"
aws s3 sync ./aptsumocli/ s3://aptsumocli
