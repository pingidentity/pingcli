## Configuration File

The following parameters can be configured in Ping CLI's static configuration file, usually located at $HOME/.pingcli/config.yaml. The following describes the properties that can be set, and an example can be found at [example-configuration.md](./example-configuration.md)


#### General Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| activeProfile | ENUM_STRING | | The name of the stored custom configuration profile to use by default. |
| noColor | ENUM_BOOL | --no-color | Disable text output in color. |
| outputFormat | ENUM_OUTPUT_FORMAT | --output-format / -O | Specify the console output format.<br><br>Options are: json, text.<br><br>Example: `json` |

#### Ping Platform Service Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| service.pingFederate.adminAPIPath | ENUM_STRING | --pingfederate-admin-api-path | The PingFederate API URL path used to communicate with PingFederate's admin API.<br><br>Example: `/pf-admin-api/v1` |
| service.pingFederate.authentication.accessTokenAuth.accessToken | ENUM_STRING | --pingfederate-access-token | The PingFederate access token used to authenticate to the PingFederate admin API when using a custom OAuth 2.0 token method. |
| service.pingFederate.authentication.basicAuth.password | ENUM_STRING | --pingfederate-password | The PingFederate password used to authenticate to the PingFederate admin API when using basic authentication. |
| service.pingFederate.authentication.basicAuth.username | ENUM_STRING | --pingfederate-username | The PingFederate username used to authenticate to the PingFederate admin API when using basic authentication. Example: `administrator` |
| service.pingFederate.authentication.clientCredentialsAuth.clientID | ENUM_STRING | --pingfederate-client-id | The PingFederate OAuth client ID used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.clientCredentialsAuth.clientSecret | ENUM_STRING | --pingfederate-client-secret | The PingFederate OAuth client secret used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.clientCredentialsAuth.scopes | ENUM_STRING_SLICE | --pingfederate-scopes | The PingFederate OAuth scopes used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type.<br><br>Accepts a comma-separated string to delimit multiple scopes.<br><br>Example: `openid,profile` |
| service.pingFederate.authentication.clientCredentialsAuth.tokenURL | ENUM_STRING | --pingfederate-token-url | The PingFederate OAuth token URL used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.type | ENUM_PINGFEDERATE_AUTH_TYPE | --pingfederate-authentication-type | The authentication type to use when connecting to the PingFederate admin API.<br><br>Options are: accessTokenAuth, basicAuth, clientCredentialsAuth.<br><br>Example: `basicAuth` |
| service.pingFederate.caCertificatePEMFiles | ENUM_STRING_SLICE | --pingfederate-ca-certificate-pem-files | Relative or full paths to PEM-encoded certificate files to be trusted as root CAs when connecting to the PingFederate server over HTTPS.<br><br>Accepts a comma-separated string to delimit multiple PEM files. |
| service.pingFederate.httpsHost | ENUM_STRING | --pingfederate-https-host | The PingFederate HTTPS host used to communicate with PingFederate's admin API.<br><br>Example: `https://pingfederate-admin.bxretail.org` |
| service.pingFederate.insecureTrustAllTLS | ENUM_BOOL | --pingfederate-insecure-trust-all-tls | Trust any certificate when connecting to the PingFederate server admin API.<br><br>This is insecure and shouldn't be enabled outside of testing. |
| service.pingFederate.xBypassExternalValidationHeader | ENUM_BOOL | --pingfederate-x-bypass-external-validation-header | Bypass connection tests when configuring PingFederate (the X-BypassExternalValidation header when using PingFederate's admin API). |
| service.pingOne.authentication.authCode.clientID | ENUM_STRING | | The authorization code client ID used to authenticate to the PingOne management API when using OAuth 2.0 authorization code flow. |
| service.pingOne.authentication.authCode.environmentID | ENUM_UUID | | The ID of the PingOne environment that contains the authorization code client used to authenticate to the PingOne management API. |
| service.pingOne.authentication.authCode.redirectURI | ENUM_STRING | | The redirect URI configured for the authorization code client application.<br><br>Example: `http://127.0.0.1:7464/callback` |
| service.pingOne.authentication.clientCredentials.clientID | ENUM_STRING | | The client credentials client ID used to authenticate to the PingOne management API when using OAuth 2.0 client credentials flow. |
| service.pingOne.authentication.clientCredentials.clientSecret | ENUM_STRING | | The client credentials client secret used to authenticate to the PingOne management API when using OAuth 2.0 client credentials flow. |
| service.pingOne.authentication.clientCredentials.environmentID | ENUM_UUID | | The ID of the PingOne environment that contains the client credentials application used to authenticate to the PingOne management API. |
| service.pingOne.authentication.deviceCode.clientID | ENUM_STRING | | The device code client ID used to authenticate to the PingOne management API when using OAuth 2.0 device code flow. |
| service.pingOne.authentication.deviceCode.environmentID | ENUM_UUID | | The ID of the PingOne environment that contains the device code client used to authenticate to the PingOne management API. |
| service.pingOne.regionCode | ENUM_PINGONE_REGION_CODE | --pingone-region-code | The region code of the PingOne tenant.<br><br>Options are: AP, AU, CA, EU, NA, SG.<br><br>Example: `NA` |

#### Platform Export Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| export.format | ENUM_EXPORT_FORMAT | --format / -f | Specifies the export format.<br><br>Options are: HCL.<br><br>Example: `HCL` |
| export.outputDirectory | ENUM_STRING | --output-directory / -d | Specifies the output directory for export. Example: `$HOME/pingcli-export` |
| export.overwrite | ENUM_BOOL | --overwrite / -o | Overwrites the existing generated exports in output directory. |
| export.pingone.environmentID | ENUM_UUID | --pingone-export-environment-id | The ID of the PingOne environment to export. Must be a valid PingOne UUID. |
| export.serviceGroup | ENUM_EXPORT_SERVICE_GROUP | --service-group / -g | Specifies the service group to export. <br><br>Options are: pingone.<br><br>Example: `pingone` |
| export.services | ENUM_EXPORT_SERVICES | --services / -s | Specifies the service(s) to export. Accepts a comma-separated string to delimit multiple services.<br><br>Options are: pingfederate, pingone-mfa, pingone-platform, pingone-protect, pingone-sso.<br><br>Example: `pingone-sso,pingone-mfa,pingfederate` |

#### Custom Request Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| request.fail | ENUM_BOOL | --fail / -f | Return non-zero exit code when HTTP custom request returns a failure status code. |
| request.service | ENUM_REQUEST_SERVICE | --service / -s | The Ping service (configured in the active profile) to send the custom request to.<br><br>Options are: pingone.<br><br>Example: `pingone` |