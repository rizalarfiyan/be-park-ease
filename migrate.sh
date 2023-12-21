#!/bin/sh

# Check Command goose
command="./goose"
if [ -x "$(command -v goose)" ]; then
    command="goose"
fi

if [ -z "$command" ]; then
    echo "Command goose could not be found"
    exit 1
fi

# Action goose
DSN_DB="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
OPTIONS="-dir ./database/schema postgres $DSN_DB"

case "$1" in
    "new")
    "$command" $OPTIONS create $2 sql
    ;;
    "up")
    "$command" $OPTIONS up
    ;;
    "redo")
    "$command" $OPTIONS redo
    ;;
    "status")
    "$command" $OPTIONS status
    ;;
    "down")
    "$command" $OPTIONS down
    ;;
    *)
    echo "Usage: $(basename "$0") new {name}/up/status/down"
    exit 1
esac
