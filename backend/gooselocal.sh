#! /bin/sh

set -a
. ./.env.local
set +a

goose -dir ./sql/migration "$@"