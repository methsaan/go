#! /bin/bash

set -o pipefail
set -o nounset
set -o errexit

OUT=$(envdir $HOME/pkg/conf/env/test/test \
 ./pg_find_primary_no_exit_error 5432 postgres welcome postgres 2>/dev/null || true)

[[ -z "$OUT" ]] && exit 1

echo "ALL is Good, running java application"
echo "$OUT"
