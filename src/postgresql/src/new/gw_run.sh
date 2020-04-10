#! /bin/bash

set -o pipefail
set -o nounset
set -o errexit

export PATH=$PATH:.

BASEDIR="$HOME/basedir"

JDBC_HOST=127.0.0.1
JDBC_PORT=5432
JDBC_USERNAME=_owner
JDBC_PASSWD=welcome
JDBC_NAME=_database

envdir $BASEDIR/conf/env/test/test \
    pg_select_now \
        "${JDBC_HOST}" \
        "${JDBC_PORT}" \
        "${JDBC_USERNAME}" \
        "${JDBC_PASSWD}" \
        "${JDBC_NAME}" \
        | egrep --line-buffered --max-count=1 "40"

echo SUCCESS
