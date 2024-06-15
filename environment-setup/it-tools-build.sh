#!/usr/bin/env bash

TMP_DIR=$(mktemp -d -t "${1}.XXXXXXX")
LATEST_TAG=$(curl -sS  https://api.github.com/repos/CorentinTh/it-tools/releases/latest | jq -r '.tag_name')

pushd ${TMP_DIR} 1>/dev/null

echo "[*] Download ${TMP_DIR}/${LATEST_TAG}.tar.gz"
curl -sSOL "https://github.com/CorentinTh/it-tools/archive/refs/tags/${LATEST_TAG}.tar.gz"
DIR_NAME=$(tar -tf ${LATEST_TAG}.tar.gz | head -1)

echo "[*] Extract ${LATEST_TAG}.tar.gz"
tar -xf ${LATEST_TAG}.tar.gz

pushd $DIR_NAME 1>/dev/null

echo "[*] Update BASE_URL in Dockerfile"
# Update your BASE_URL below!
sed -i '/RUN pnpm build/i ENV BASE_URL="/it-tools/"' Dockerfile

echo "[*] Run docker build"
docker build -t it-tools .

echo "[*] Remove ${TMP_DIR}"
rm -rf ${TMP_DIR}

