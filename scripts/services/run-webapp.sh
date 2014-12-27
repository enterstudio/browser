#!/bin/bash

# This script runs mounttables, proxy daemons, a http server and
# a few sample servers to run and demo the Veyron Browser.

source "${VANADIUM_ROOT}/release/projects/namespace_browser/scripts/services/common.sh"

main() {
  local -r MOUNTTABLE_PORT=5167
  local -r MOUNTTABLE_PORT_HOUSE=5168
  local -r MOUNTTABLE_PORT_COTTAGE=5169
  local -r WSPR_PORT=8124
  local -r PROXY_PORT=5164
  local -r VEYRON_IDENTITY_DIR="${TMPDIR}/app_credentials_dir"
  local -r HTTP_PORT=9000
  local -r SEEK_BLESSSING=true

  common::run "${MOUNTTABLE_PORT}" "${MOUNTTABLE_PORT_HOUSE}" "${MOUNTTABLE_PORT_COTTAGE}" "${WSPR_PORT}" "${PROXY_PORT}" "${VEYRON_IDENTITY_DIR}" "${SEEK_BLESSSING}"

  local -r SERVE="${VANADIUM_ROOT}/release/projects/namespace_browser/node_modules/.bin/serve"
  "${SERVE}" "${VANADIUM_ROOT}"/release/projects/namespace_browser/public/. --port "${HTTP_PORT}" --compress &
}

main "$@"
wait
