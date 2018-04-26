#!/usr/bin/env bash
set -e

SRC_ROOT=$(git rev-parse --show-toplevel)
VERSION=$(git describe --tags --dirty)
COMMIT_HASH=$(git rev-parse --short HEAD 2>/dev/null)
DATE=$(date "+%Y-%m-%d")
IMPORT_DURING_SOLVE=${IMPORT_DURING_SOLVE:-false}

if [[ -z "${BASENAME}" ]]; then
    BASENAME="out"
fi

if [[ -z "${RELEASE}" ]]; then
    RELEASE="release"
fi

if [[ "$(pwd)" != "${SRC_ROOT}" ]]; then
  echo "you are not in the root of the repo" 1>&2
  echo "please cd to ${SRC_ROOT} before running this script" 1>&2
  exit 1
fi

if [[ -z "${GO_BUILD_FLAGS}" ]]; then
    GO_BUILD_FLAGS="-a -installsuffix cgo"
fi

GO_BUILD_CMD="go build ${GO_BUILD_FLAGS} "
GO_BUILD_LDFLAGS="-s -w -X main.commitHash=${COMMIT_HASH} -X main.buildDate=${DATE} -X main.version=${VERSION} -X main.flagImportDuringSolve=${IMPORT_DURING_SOLVE}"

if [[ -z "${SRC_BUILD_PLATFORMS}" ]]; then
    SRC_BUILD_PLATFORMS="linux windows darwin freebsd"
fi

if [[ -z "${SRC_BUILD_ARCHS}" ]]; then
    SRC_BUILD_ARCHS="amd64 386"
fi

mkdir -p "${SRC_ROOT}/${RELEASE}"

for OS in ${SRC_BUILD_PLATFORMS[@]}; do
  for ARCH in ${SRC_BUILD_ARCHS[@]}; do
    NAME="${BASENAME}_${OS}_${ARCH}"
    if [[ "${OS}" == "windows" ]]; then
      NAME="${NAME}.exe"
    fi
    NAME="${SRC_ROOT}/${RELEASE}/${NAME}"
    echo "Building to ${NAME} for ${OS}/${ARCH}"
    GOARCH=${ARCH} GOOS=${OS} CGO_ENABLED=0 ${GO_BUILD_CMD} -ldflags "${GO_BUILD_LDFLAGS}" -o "${NAME}" ./cmd/${BASENAME}
    shasum -a 256 "${NAME}" > "${NAME}".sha256
  done
done
