#!/bin/sh
set -e
. /usr/lib/nextar/server-lib.sh
before_remove nextar-server.service "$1"
