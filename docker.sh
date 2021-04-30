#!/bin/bash

COMMAND=${1:-"usage"}
SERVER_NAME=${2:-"server1"}

usage() {
  cat <<EOF
------------------------------------------ docker.sh ------------------------------------------
Usage: ./docker.sh <command> [variables]
  build:                 build simple-server.          
  push:                  push to dockerhub.
  register:              build and push to dockerhub a image.

examples:
  ./docker.sh build server1
  ./docker.sh push server2
  ./docker.sh register server3
-----------------------------------------------------------------------------------------------
EOF
}

build() {
  docker build -t "morosawamikihito/simple-server:${SERVER_NAME}" . --build-arg servername=${SERVER_NAME}
}

push() {
  docker push "morosawamikihito/simple-server:${SERVER_NAME}"
}

register() {
  build
  push
}

case "${COMMAND}" in
  build) build ;;
  push) push ;;
  register) register ;;
  *) usage ;;
esac
