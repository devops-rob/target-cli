#!/bin/bash -e

echo "################################"
echo "Installing Target CLI to /usr/local/bin/target"
echo ""
echo "Please note: You may be prompted for your password"

# Determine the OS and architecure
OS=$(uname)
ARCH=$(uname -m)
TARGET_OS="linux"
TARGET_ARCH="amd64"

if [ "${ARCH}" == "arm64" ]; then
  TARGET_ARCH="arm64"
fi

if [ "${OS}" == "Linux" ]; then
  TARGET_OS="linux"
  TARGET_EXT="tar.gz"
  TARGET_COMMAND="tar -xzf"
fi

if [ "${OS}" == "Darwin" ]; then
  TARGET_OS="darwin"
  TARGET_EXT="zip"
  TARGET_COMMAND="unzip"
fi

version=$(curl https://api.github.com/repos/devops-rob/target-cli/releases/latest | awk -F':|,' '/tag_name/{gsub(/[" ]/,"",$2); print $2}')
binary="target-cli_${version}_${TARGET_OS}_${TARGET_ARCH}.${TARGET_EXT}"



repo="https://github.com/devops-rob/target-cli/releases/download/${version}"

echo "Downloading $repo/$binary"
rm -f /tmp/target-cli.${TARGET_EXT}
rm -f /tmp/target-cli

curl -L -s -o /tmp/target-cli.${TARGET_EXT} "$repo/$binary"
cd /tmp

eval ${TARGET_COMMAND} target-cli.${TARGET_EXT}
rm target-cli.${TARGET_EXT}

if which sudo; then
  sudo mv target /usr/local/bin
  sudo chmod +x /usr/local/bin/target
else
  # try without sudo might be running in a container
  mv target /usr/local/bin
  chmod +x /usr/local/bin/target
fi

echo ""

target

target version





