#!/usr/bin/env bash

SCRIPT_PATH=$(dirname "$(readlink -f "${0}")")
cd "${SCRIPT_PATH}"

set -e

RELEASE_VOLUME="-v $(realpath "$(pwd)/../dist"):/test"

function clean() {
    docker stop "${CONTAINER_ID}"
    docker wait "${CONTAINER_ID}"
}

function start() {
  docker build . -t systemd-rest-e2e
  export CONTAINER_ID="$(docker run -d -p "${PORT}:${PORT}" --tmpfs /tmp --tmpfs /run --tmpfs /run/lock -v /sys/fs/cgroup:/sys/fs/cgroup:ro ${RELEASE_VOLUME} systemd-rest-e2e)"
  trap "clean" SIGTERM SIGINT
}

function exec_service(){
  docker exec -it "${CONTAINER_ID}" /bin/systemctl enable testing-unit.service
  docker exec -it "${CONTAINER_ID}" /bin/systemctl start testing-unit.service
  if [[ "${1}" == "follow" ]]; then
    docker exec -e MODE -e PORT -e ADDRESS=0.0.0.0 -it "${CONTAINER_ID}"  /test/systemd-rest
  else
    docker exec -e MODE -e PORT -e ADDRESS=0.0.0.0 -itd "${CONTAINER_ID}" /test/systemd-rest
  fi
}

start
set +e
exec_service "${1}"
pytest -sv
clean
