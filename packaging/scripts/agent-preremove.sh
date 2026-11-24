#!/bin/sh
set -e
. /usr/lib/nextar/agent-lib.sh
before_remove nextar-agent.service "$1"
