## Configuration File

The following parameters can be configured in Ping CLI's static configuration file, usually located at $HOME/.pingcli/config.yaml. The following describes the properties that can be set, and an example can be found at [example-configuration.md](./example-configuration.md)


#### General Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| activeProfile | ENUM_STRING | | The name of the stored custom configuration profile to use by default. |
| color | ENUM_BOOL | --color | Show text output in color. |
| outputFormat | ENUM_OUTPUT_FORMAT | --output-format / -O | Specify the console output format.<br><br>Options are: json, text.<br><br>Example: `json` |

#### Ping Platform Service Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| service.pingfederate.adminAPIPath | ENUM_STRING | --pingfederate-admin-api-path | The PingFederate API URL path used to communicate with PingFederate's admin API.<br><br>Example: `/pf-admin-api/v1` |
| service.pingfederate.authentication.accessTokenAuth.accessToken | ENUM_STRING | --pingfederate-access-token | The PingFederate access token used to authenticate to the PingFederate admin API when using a custom OAuth 2.0 token method. |
| service.pingfederate.authentication.basicAuth.password | ENUM_STRING | --pingfederate-password | The PingFederate password used to authenticate to the PingFederate admin API when using basic authentication. |
| service.pingfederate.authentication.basicAuth.username | ENUM_STRING | --pingfederate-username | The PingFederate username used to authenticate to the PingFederate admin API when using basic authentication. Example: `administrator` |
| service.pingfederate.authentication.clientCredentialsAuth.clientID | ENUM_STRING | --pingfederate-client-id | The PingFederate OAuth client ID used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingfederate.authentication.clientCredentialsAuth.clientSecret | ENUM_STRING | --pingfederate-client-secret | The PingFederate OAuth client secret used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingfederate.authentication.clientCredentialsAuth.scopes | ENUM_STRING_SLICE | --pingfederate-scopes | The PingFederate OAuth scopes used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type.<br><br>Accepts a comma-separated string to delimit multiple scopes.<br><br>Example: `openid,profile` |
| service.pingfederate.authentication.clientCredentialsAuth.tokenURL | ENUM_STRING | --pingfederate-token-url | The PingFederate OAuth token URL used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingfederate.authentication.type | ENUM_PINGFEDERATE_AUTH_TYPE | --pingfederate-authentication-type | The authentication type to use when connecting to the PingFederate admin API.<br><br>Options are: accessTokenAuth, basicAuth, clientCredentialsAuth.<br><br>Example: `basicAuth` |
| service.pingfederate.caCertificatePemFiles | ENUM_STRING_SLICE | --pingfederate-ca-certificate-pem-files | Relative or full paths to PEM-encoded certificate files to be trusted as root CAs when connecting to the PingFederate server over HTTPS.<br><br>Accepts a comma-separated string to delimit multiple PEM files. |
| service.pingfederate.httpsHost | ENUM_STRING | --pingfederate-https-host | The PingFederate HTTPS host used to communicate with PingFederate's admin API.<br><br>Example: `https://pingfederate-admin.bxretail.org` |
| service.pingfederate.insecureTrustAllTLS | ENUM_BOOL | --pingfederate-insecure-trust-all-tls | Trust any certificate when connecting to the PingFederate server admin API.<br><br>This is insecure and should not be enabled outside of testing. |
| service.pingfederate.xBypassExternalValidationHeader | ENUM_BOOL | --pingfederate-x-bypass-external-validation-header | Bypass connection tests when configuring PingFederate (the X-BypassExternalValidation header when using PingFederate's admin API). |
| service.pingone.authentication.type | ENUM_PINGONE_AUTH_TYPE | --pingone-authentication-type | The authentication type to use to authenticate to the PingOne management API.<br><br>Options are: worker.<br><br>Example: `worker` |
| service.pingone.authentication.worker.clientID | ENUM_UUID | --pingone-worker-client-id | The worker client ID used to authenticate to the PingOne management API. |
| service.pingone.authentication.worker.clientSecret | ENUM_STRING | --pingone-worker-client-secret | The worker client secret used to authenticate to the PingOne management API. |
| service.pingone.authentication.worker.environmentID | ENUM_UUID | --pingone-worker-environment-id | The ID of the PingOne environment that contains the worker client used to authenticate to the PingOne management API. |
| service.pingone.regionCode | ENUM_PINGONE_REGION_CODE | --pingone-region-code | The region code of the PingOne tenant.<br><br>Options are: AP, AU, CA, EU, NA.<br><br>Example: `NA` |

#### Platform Export Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| export.format | ENUM_STRING | --format / -f | Specifies the export format.<br><br>Options are: HCL.<br><br>Example: `HCL` |
| export.outputDirectory | ENUM_STRING | --output-directory / -d | Specifies the output directory for export. Example: `$HOME/pingcli-export` |
| export.overwrite | ENUM_BOOL | --overwrite / -o | Overwrite the existing generated exports in output directory. |
| export.pingone.environmentID | ENUM_UUID | --pingone-export-environment-id | The ID of the PingOne environment to export. Must be a valid PingOne UUID. |
| export.services | ENUM_EXPORT_SERVICES | --services / -s | Specifies the service(s) to export. Accepts a comma-separated string to delimit multiple services.<br><br>Options are: pingfederate, pingone-mfa, pingone-platform, pingone-protect, pingone-sso.<br><br>Example: `pingone-sso,pingone-mfa,pingfederate` |

#### Custom Request Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| request.service | ENUM_REQUEST_SERVICE | --service / -s | The Ping service (configured in the active profile) to send the custom request to.<br><br>Options are: pingone.<br><br>Example: `pingone` |