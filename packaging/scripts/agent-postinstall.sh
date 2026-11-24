#!/bin/sh
set -e
. /usr/lib/nextar/agent-lib.sh
after_install nextar-agent.service "$1" "$2"
