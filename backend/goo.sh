#! /bin/sh

export GOOSE_DRIVER="postgres"
export GOOSE_DBSTRING="postgres://postgres:password@localhost:5555/anylb?sslmode=disable"

# set -a
# . ./.env.local
# set +a

goose -dir ./sql/migration "$@"
