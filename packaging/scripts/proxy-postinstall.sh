#!/bin/sh
set -e
. /usr/lib/nextar/proxy-lib.sh
ensure_user nextar-proxy /var/lib/nextar-proxy
readable_by nextar-proxy /etc/nextar/proxy.yaml
after_install nextar-proxy.service "$1" "$2"
