## Configuration File

The following parameters can be configured in Ping CLI's static configuration file, usually located at $HOME/.pingcli/config.yaml. The following describes the properties that can be set, and an example can be found at [example-configuration.md](./example-configuration.md)

#### General Properties

| Configuration Key | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| `detailedExitCode` | Boolean | `--detailed-exitcode` / `-D` | Enable detailed exit code output. (default false)<br><br>0 - pingcli command succeeded with no errors or warnings.<br><br>1 - pingcli command failed with errors.<br><br>2 - pingcli command succeeded with warnings. |
| `noColor` | Boolean | `--no-color` | Disable text output in color. (default false) |
| `outputFormat` | String (enum) | `--output-format` / `-O` | Specify the console output format. (default text)<br><br>Options are: json, text. |

#### Export Properties

| Configuration Key | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| `export.format` | String (enum) | `--format` / `-f` | Specifies the export format. (default HCL)<br><br>Options are: HCL. |
| `export.outputDirectory` | String | `--output-directory` / `-d` | Specifies the output directory for export. Can be an absolute filepath or a relative filepath of the present working directory. <br><br>Example: '/Users/example/pingcli-export'<br><br>Example: 'pingcli-export' |
| `export.overwrite` | Boolean | `--overwrite` / `-o` | Overwrites the existing generated exports in output directory. (default false) |
| `export.pingOne.environmentID` | String (UUID Format) | `--pingone-export-environment-id` | The ID of the PingOne environment to export. Must be a valid PingOne UUID. |
| `export.serviceGroup` | String (enum) | `--service-group` / `-g` | Specifies the service group to export. <br><br>Options are: pingone.<br><br>Example: 'pingone' |
| `export.services` | String Array (enum) | `--services` / `-s` | Specifies the service(s) to export. Accepts a comma-separated string to delimit multiple services. <br><br>Options are: pingfederate, pingone-authorize, pingone-mfa, pingone-platform, pingone-protect, pingone-sso.<br><br>Example: 'pingone-sso,pingone-mfa,pingfederate' |

#### Request Properties

| Configuration Key | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| `request.fail` | Boolean | `--fail` / `-f` | Return non-zero exit code when HTTP custom request returns a failure status code. |
| `request.service` | String (enum) | `--service` / `-s` | The Ping service (configured in the active profile) to send the custom request to.<br><br>Options are: pingone.<br><br>Example: 'pingone' |

#### Service Properties

| Configuration Key | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| `service.pingFederate.adminAPIPath` | String | `--pingfederate-admin-api-path` | The PingFederate API URL path used to communicate with PingFederate's admin API. (default /pf-admin-api/v1) |
| `service.pingFederate.authentication.accessTokenAuth.accessToken` | String | `--pingfederate-access-token` | The PingFederate access token used to authenticate to the PingFederate admin API when using a custom OAuth 2.0 token method. |
| `service.pingFederate.authentication.basicAuth.password` | String | `--pingfederate-password` | The PingFederate password used to authenticate to the PingFederate admin API when using basic authentication. |
| `service.pingFederate.authentication.basicAuth.username` | String | `--pingfederate-username` | The PingFederate username used to authenticate to the PingFederate admin API when using basic authentication.<br><br>Example: 'administrator' |
| `service.pingFederate.authentication.clientCredentialsAuth.clientID` | String | `--pingfederate-client-id` | The PingFederate OAuth client ID used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| `service.pingFederate.authentication.clientCredentialsAuth.clientSecret` | String | `--pingfederate-client-secret` | The PingFederate OAuth client secret used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| `service.pingFederate.authentication.clientCredentialsAuth.scopes` | String Array | `--pingfederate-scopes` | The PingFederate OAuth scopes used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. (default [])<br><br>Accepts a comma-separated string to delimit multiple scopes.<br><br>Example: 'openid,profile' |
| `service.pingFederate.authentication.clientCredentialsAuth.tokenURL` | String | `--pingfederate-token-url` | The PingFederate OAuth token URL used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| `service.pingFederate.authentication.type` | String (enum) | `--pingfederate-authentication-type` | The authentication type to use when connecting to the PingFederate admin API.<br><br>Options are: accessTokenAuth, basicAuth, clientCredentialsAuth.<br><br>Example: 'basicAuth' |
| `service.pingFederate.caCertificatePEMFiles` | String Array | `--pingfederate-ca-certificate-pem-files` | Relative or full paths to PEM-encoded certificate files to be trusted as root CAs when connecting to the PingFederate server over HTTPS. (default [])<br><br>Accepts a comma-separated string to delimit multiple PEM files. |
| `service.pingFederate.httpsHost` | String | `--pingfederate-https-host` | The PingFederate HTTPS host used to communicate with PingFederate's admin API.<br><br>Example: 'https://pingfederate-admin.bxretail.org' |
| `service.pingFederate.insecureTrustAllTLS` | Boolean | `--pingfederate-insecure-trust-all-tls` | Trust any certificate when connecting to the PingFederate server admin API. (default false)<br><br>This is insecure and shouldn't be enabled outside of testing. |
| `service.pingFederate.xBypassExternalValidationHeader` | Boolean | `--pingfederate-x-bypass-external-validation-header` | Bypass connection tests when configuring PingFederate (the X-BypassExternalValidation header when using PingFederate's admin API). (default false) |
| `service.pingOne.authentication.type` | String (enum) | `--pingone-authentication-type` | The authentication type to use to authenticate to the PingOne management API. (default worker)<br><br>Options are: worker. |
| `service.pingOne.authentication.worker.clientID` | String (UUID Format) | `--pingone-worker-client-id` | The worker client ID used to authenticate to the PingOne management API. |
| `service.pingOne.authentication.worker.clientSecret` | String | `--pingone-worker-client-secret` | The worker client secret used to authenticate to the PingOne management API. |
| `service.pingOne.authentication.worker.environmentID` | String (UUID Format) | `--pingone-worker-environment-id` | The ID of the PingOne environment that contains the worker client used to authenticate to the PingOne management API. |
| `service.pingOne.regionCode` | String (enum) | `--pingone-region-code` | The region code of the PingOne tenant.<br><br>Options are: AP, AU, CA, EU, NA.<br><br>Example: 'NA' |