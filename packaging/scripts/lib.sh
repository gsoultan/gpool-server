# Shared by the install scripts of all three packages. nfpm inlines nothing, so
# each script sources this file from where the package installs it.
#
# deb and rpm call maintainer scripts with different arguments for the same
# event, and the difference matters: stopping the service on an upgrade would
# drop every client connection for no reason.
#
#   event            deb                         rpm
#   first install    postinst configure          post 1
#   upgrade          postinst configure <old>    post 2
#   removal          prerm remove                preun 0
#   upgrade          prerm upgrade <new>         preun 1

is_upgrade_after_install() {
  case "$1" in
    configure) [ -n "$2" ] ;;
    *) [ "$1" = "2" ] ;;
  esac
}

is_removal() {
  case "$1" in
    remove | purge) return 0 ;;
    upgrade | deconfigure | failed-upgrade) return 1 ;;
    *) [ "$1" = "0" ] ;;
  esac
}

has_systemd() {
  [ -d /run/systemd/system ]
}

ensure_user() {
  user="$1"
  home="$2"
  getent group "$user" >/dev/null 2>&1 || groupadd --system "$user"
  getent passwd "$user" >/dev/null 2>&1 ||
    useradd --system --gid "$user" --home-dir "$home" --no-create-home \
      --shell /usr/sbin/nologin "$user"
}

# The config file is installed root-owned and 0640 because it can hold a token.
# The service runs as its own user, so that user's group must be able to read it.
# Done here rather than in the package because the group does not exist until
# ensure_user has run.
readable_by() {
  chgrp "$1" "$2"
  chmod 0640 "$2"
}

# after_install <unit> <install-arg-1> <install-arg-2>
#
# An upgrade restarts the service only if it was already running. A first
# install starts nothing: the operator has not configured it yet, and a database
# host is not a place for a service to start by surprise.
after_install() {
  unit="$1"
  has_systemd || return 0
  systemctl daemon-reload
  if is_upgrade_after_install "$2" "$3"; then
    systemctl try-restart "$unit" || true
    return 0
  fi
  echo "$unit installed. Start it with: systemctl enable --now $unit"
}

# before_remove <unit> <remove-arg-1>
before_remove() {
  unit="$1"
  has_systemd || return 0
  is_removal "$2" || return 0
  systemctl disable --now "$unit" || true
}
