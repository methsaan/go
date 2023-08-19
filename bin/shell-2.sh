#! /bin/bash

set -x

set -o pipefail
set -o nounset
set -o errexit

S=$(envdir $HOME/pkg/conf/env/test/test ./get_env_2 2>/dev/null||true)

[[ "${S}" ]]

echo made-it
