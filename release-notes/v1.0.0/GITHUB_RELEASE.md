# Ping CLI v1.0.0

Ping CLI is a unified command-line interface for configuring and managing Ping
Identity services. This is the first stable release. It represents a complete
redesign of the tool — the command structure, output format, authentication model,
and resource coverage are all substantially different from pre-1.0 versions.

## What's in 1.0

### CRUD Resource Management

Ping CLI now supports create, read, update, and delete operations against Ping Identity
APIs directly from the CLI. Resources are organized under product subcommands.

**PingOne** — 50+ resources including Applications, Users, Groups, Identity Providers,
Sign-On Policies, Password Policies, Agreements, Gateways, Schemas, Roles, Webhooks,
Notifications, Risk Predictors, and more.

**PingOne Universal Services** — dedicated subcommands for MFA, Verify, Credentials,
Authorize, and Protect.

**DaVinci** — Flows, Flow Policies, Applications, Connector Instances, and Connector
Catalog.

### Authentication

- Multiple grant types: authorization code, client credentials, device code
- Auto-login and silent token refresh for human auth flows
- Per-profile logout with `--all` and `--profile` flags
- Auth status via `auth status` with structured output

### Output Formats

- Human-readable text (default) with terminal-width-aware tables and Unicode symbols
- JSON (`-O json`)
- NDJSON wrapped (`-O ndjson-wrapped`) — newline-delimited JSON where each record
  is wrapped in a structured envelope with resource, action, effect, and duration
- JMESPath filtering via `--query` flag on all JSON output paths
- Structured error payloads on all output paths

### Observability

- `--debug` flag: stack traces, transaction IDs, raw HTTP response data on errors
- `--log-level`, `--log-file`, `--log-file-level` flags for structured file logging
- HTTP debug round-tripper for request/response inspection
- OpenTelemetry tracing throughout

### Plugin System

External plugins via HashiCorp go-plugin and gRPC, with `plugin` subcommand for
discovery and execution.

### Configuration

- Profile-based configuration with case-insensitive keys
- Environment variable support for all auth flags (no more `--pingone-*` flags)
- `config list-keys` command

## Platform Support

Linux (amd64, arm64), macOS (amd64, arm64), Windows (amd64).
Installation via direct download, brew, or Docker image.

## Breaking Changes from Pre-1.0

The entire command structure has been redesigned. Pre-1.0 commands, flags, config
keys, and output formats are not compatible with 1.0. Consult the
[documentation](https://developer.pingidentity.com/pingcli/1.0) for current usage.

Existing configuration files from pre-1.0 installs are migrated automatically
the first time 1.0 runs.

The `platform export` command from pre-1.0 is not included in 1.0 — it is being
redesigned. Pre-1.0 releases remain available and will be maintained until a
replacement is ready.
