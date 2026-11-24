#!/bin/sh
set -e
. /usr/lib/nextar/proxy-lib.sh
before_remove nextar-proxy.service "$1"
