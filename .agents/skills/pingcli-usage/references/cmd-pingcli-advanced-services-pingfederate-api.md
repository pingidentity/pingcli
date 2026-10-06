# `pingcli advanced-services pingfederate api`
Send a custom REST API request to the management API of PingFederate (Advanced Services).

## Synopsis

Send a custom REST API request to the management API of PingFederate (Advanced Services)

The custom REST API request is most powerful when product connection details have been configured in the CLI configuration file.

The command offers a cURL-like experience to interact with the Ping products and services, with authentication and environment details dynamically filled by the CLI.

```
pingcli advanced-services pingfederate api [flags] API_URI
```

## Examples

```
Send a custom API request to the configured PingFederate (Advanced Services) server, making a GET request against the /serverSettings endpoint.
    pingcli advanced-services pingfederate api serverSettings

  Send a custom API request to the configured PingFederate (Advanced Services) server, making a GET request to retrieve JSON configuration for authentication policies.
    pingcli advanced-services pingfederate api --http-method GET --output-format json authenticationPolicies/default

  Send a custom API request to the configured PingFederate (Advanced Services) server, making a POST request to create a new data store with JSON data sourced from a file.
    pingcli advanced-services pingfederate api --http-method POST --data ./my-data-store.json dataStores

  Send a custom API request to the configured PingFederate (Advanced Services) server, making a PUT request using raw JSON data.
    pingcli advanced-services pingfederate api --http-method PUT --data-raw '{"name": "My data store"}' dataStores/my-data-store-id

  Send a custom API request to the configured PingFederate (Advanced Services) server, making a DELETE request to remove a data store.
    pingcli advanced-services pingfederate api --http-method DELETE dataStores/my-data-store-id
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for api |
| `-f, --fail` | `` | Return non-zero exit code when HTTP custom API request returns a failure status code. |
| `-m, --http-method string` | `` | The HTTP method to use for the request. (default GET) Options are: DELETE, GET, PATCH, POST, PUT. Example: 'POST' |
| `-r, --header []string` | `` | A custom header to send in the request. Example: --header "Content-Type: application/vnd.pingidentity.user.import+json" |
| `--as-pingfederate-admin-api-path string` | `` | The PingFederate admin API URL path for the Advanced Services private tenant. (default /pf-admin-api/v1) |
| `--as-pingfederate-ca-certificate-pem-files []string` | `` | Relative or full paths to PEM-encoded certificate files for the Advanced Services PingFederate server. (default []) Accepts a comma-separated string to delimit multiple PEM files. |
| `--as-pingfederate-https-host string` | `` | The PingFederate HTTPS host for the Advanced Services private tenant. Example: 'https://pingfederate-admin.bxretail.org' |
| `--as-pingfederate-insecure-trust-all-tls` | `` | Trust any certificate when connecting to the PingFederate Advanced Services admin API. (default false) This is insecure and shouldn't be enabled outside of testing. |
| `--as-pingfederate-software-version string` | `` | The PingFederate software version for the Advanced Services private tenant. (default 13.0) |
| `--as-pingfederate-x-bypass-external-validation-header` | `` | Bypass connection tests when configuring PingFederate in Advanced Services (the X-BypassExternalValidation header). (default false) |
| `--data string` | `` | The file containing data to send in the request.  Example: './data.json' |
| `--data-raw string` | `` | The raw data to send in the request.  Example: '{"name": "My environment"}' |


## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-typed, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, ndjson-typed, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Parent Command

- [`pingcli advanced-services pingfederate`](cmd-pingcli-advanced-services-pingfederate.md) — Administration tools for PingFederate deployed as a private tenant service in PingOne Advanced Services
