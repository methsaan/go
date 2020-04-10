#! /bin/bash

set -o pipefail
set -o errexit
set -o nounset

export PATH=$PATH:.

echo ATTEMPTING ALL OK
timeout --kill-after=90 --signal=9 60 pg_pass_primary 127.0.0.1 5432 databaseX_user welcome databaseX
echo

### All errors
#echo BAD IP
#timeout --kill-after=90 --signal=9 60 pg_pass_primary 127.0.0.18  5432 databaseX_user welcome databaseX
#echo
#
#echo BAD USER
#timeout --kill-after=90 --signal=9 60 pg_pass_primary 127.0.0.1 5432 atabaseX_user  welcome databaseX
#echo
#
#echo BAD PASS
#timeout --kill-after=90 --signal=9 60 pg_pass_primary 127.0.0.1 5432 databaseX_user elcome  databaseX
#echo
#
#echo BAD DBNAME
#timeout --kill-after=90 --signal=9 60 pg_pass_primary 127.0.0.1 5432 databaseX_user welcome atabaseX
#echo
#
#echo BAD PORT
#timeout --kill-after=90 --signal=9 60 pg_pass_primary 127.0.0.1 432 databaseX_user welcome atabaseX
#echo
#
#echo BAD PORT NAME
#timeout --kill-after=90 --signal=9 60 pg_pass_primary 127.0.0.1 fuu databaseX_user welcome atabaseX
#echo

echo Database checks worked - running program
