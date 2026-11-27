# Nextar

Nextar manages databases on servers you register. It installs and tunes the
database and a proxy on each one, and gives a cluster load balancing, automatic
failover, automatic failback, backups with point-in-time recovery and safe
version upgrades.

**Status: packaging skeleton only.** The three binaries build and answer
`version`; nothing else is implemented yet. Any other command exits with an
error, so a service manager never reports a placeholder as running.

## The three binaries

| Binary | Runs on | Role |
| :--- | :--- | :--- |
| `nextar-server` | One machine | The manager: UI, API and state store. Never in the query path. |
| `nextar-agent` | Every registered worker | Carries out what `nextar-server` asks: install, tune, back up, upgrade. |
| `nextar-proxy` | Every worker that takes traffic | Pooling, routing and the failover vote. |

`nextar-server` keeps its own data (workers, clusters, users, audit log) in
PostgreSQL, through [storm](https://github.com/gsoultan/storm). That database
must not be one that this `nextar-server` manages: if that cluster failed,
`nextar-server` could not start to repair it.

A *worker* is a server you register. `nextar-agent` is the program that runs on it.

They are separate so each can be updated, restarted and resource-limited on its
own: restarting the agent never drops a client connection, and the proxy carries
only what a query needs.

## Build

```bash
gofmt -l . && go vet ./... && go test -race ./...
```

## Linux packages

[GoReleaser](https://goreleaser.com) builds static binaries for linux/amd64 and
linux/arm64, and a `.deb`, `.rpm` and `.tar.gz` for each.

```bash
goreleaser check
goreleaser release --snapshot --clean --skip=sign,sbom   # everything, into dist/
```

Each package installs:

| Path | Notes |
| :--- | :--- |
| `/usr/bin/nextar-<name>` | The binary. |
| `/usr/lib/systemd/system/nextar-<name>.service` | Not started on first install. An upgrade restarts it only if it was running. |
| `/etc/nextar/<name>.yaml` | Never overwritten by an upgrade. Mode 0640, readable by the service's group. |

`nextar-server` runs as the `nextar` user and `nextar-proxy` as `nextar-proxy`;
both are sandboxed by their units. `nextar-agent` runs as root because it
installs packages and manages the database.

The proxy shares a machine with the database, so its unit sets hard limits:
256 MB of memory, two cores, and an OOM score that makes the kernel pick the
proxy before the database.

Removing a package stops and disables its service. It does not delete
`/var/lib/nextar-*`.

## Releasing

Pushing a `v*` tag runs `.github/workflows/release.yml`, which publishes a
**draft** GitHub release for a person to review and publish.

### Release signing key

`checksums.txt` is signed with [cosign](https://github.com/sigstore/cosign). The
release workflow refuses to run without the key.

```bash
cosign generate-key-pair     # writes cosign.key and cosign.pub
```

Add two repository secrets: `COSIGN_PRIVATE_KEY` (the contents of `cosign.key`)
and `COSIGN_PASSWORD`. Keep `cosign.key` out of the repository. `cosign.pub` is
the key that will be built into the binaries for self-update, so losing the
private key means installed copies can no longer be updated automatically.
