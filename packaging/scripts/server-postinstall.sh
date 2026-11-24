#!/bin/sh
set -e
. /usr/lib/nextar/server-lib.sh
ensure_user nextar /var/lib/nextar-server
readable_by nextar /etc/nextar/server.yaml
after_install nextar-server.service "$1" "$2"
