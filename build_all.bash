#!/usr/bin/env bash
set -e

BASENAME=pic2ascii
SRC_ROOT=$(git rev-parse --show-toplevel)
VERSION=$(git describe --tags --dirty)
COMMIT_HASH=$(git rev-parse --short HEAD 2>/dev/null)
DATE=$(date "+%Y-%m-%d")
IMPORT_DURING_SOLVE=${IMPORT_DURING_SOLVE:-false}

if [[ "$(pwd)" != "${SRC_ROOT}" ]]; then
  echo "you are not in the root of the repo" 1>&2
  echo "please cd to ${SRC_ROOT} before running this script" 1>&2
  exit 1
fi

GO_BUILD_CMD="go build -a -installsuffix cgo"
GO_BUILD_LDFLAGS="-s -w -X main.commitHash=${COMMIT_HASH} -X main.buildDate=${DATE} -X main.version=${VERSION} -X main.flagImportDuringSolve=${IMPORT_DURING_SOLVE}"

if [[ -z "${SRC_BUILD_PLATFORMS}" ]]; then
    SRC_BUILD_PLATFORMS="linux windows darwin freebsd"
fi

if [[ -z "${SRC_BUILD_ARCHS}" ]]; then
    SRC_BUILD_ARCHS="amd64 386"
fi

mkdir -p "${SRC_ROOT}/release"

for OS in ${SRC_BUILD_PLATFORMS[@]}; do
  for ARCH in ${SRC_BUILD_ARCHS[@]}; do
    NAME="${BASENAME}_${OS}_${ARCH}"
    if [[ "${OS}" == "windows" ]]; then
      NAME="${NAME}.exe"
    fi
    echo "Building for ${OS}/${ARCH}"
    GOARCH=${ARCH} GOOS=${OS} CGO_ENABLED=0 ${GO_BUILD_CMD} -ldflags "${GO_BUILD_LDFLAGS}"\
     -o "${SRC_ROOT}/release/${NAME}" ./cmd/${BASENAME}
    shasum -a 256 "${SRC_ROOT}/release/${NAME}" > "${SRC_ROOT}/release/${NAME}".sha256
  done
done