#### auth Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| login.fileStorage | --file-storage | PINGCLI_AUTH_FILE_STORAGE | 0 | Store authentication tokens in local file storage only. Without this flag, keychain storage is attempted first with fallback to local file storage. |

#### export Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| export.format | --format / -f | PINGCLI_EXPORT_FORMAT | 1 | Specifies the export format. (default HCL)<br><br>Options are: HCL. |
| export.outputDirectory | --output-directory / -d | PINGCLI_EXPORT_OUTPUT_DIRECTORY | 14 | Specifies the output directory for export. Can be an absolute filepath or a relative filepath of the present working directory. <br><br>Example: '/Users/example/pingcli-export'<br><br>Example: 'pingcli-export' |
| export.overwrite | --overwrite / -o | PINGCLI_EXPORT_OVERWRITE | 0 | Overwrites the existing generated exports in output directory. (default false) |
| export.pingOne.environmentID | --pingone-export-environment-id | PINGCLI_PINGONE_EXPORT_ENVIRONMENT_ID | 16 | The ID of the PingOne environment to export. Must be a valid PingOne UUID. |
| export.serviceGroup | --service-group / -g | PINGCLI_EXPORT_SERVICE_GROUP | 2 | Specifies the service group to export. <br><br>Options are: pingone.<br><br>Example: 'pingone' |
| export.services | --services / -s | PINGCLI_EXPORT_SERVICES | 3 | Specifies the service(s) to export. Accepts a comma-separated string to delimit multiple services. <br><br>Options are: pingfederate, pingone-authorize, pingone-mfa, pingone-platform, pingone-protect, pingone-sso.<br><br>Example: 'pingone-sso,pingone-mfa,pingfederate' |

#### general Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| detailedExitCode | --detailed-exitcode / -D | PINGCLI_DETAILED_EXITCODE | 0 | Enable detailed exit code output. (default false)<br><br>0 - pingcli command succeeded with no errors or warnings.<br><br>1 - pingcli command failed with errors.<br><br>2 - pingcli command succeeded with warnings. |
| noColor | --no-color | PINGCLI_NO_COLOR | 0 | Disable text output in color. (default false) |
| outputFormat | --output-format / -O | PINGCLI_OUTPUT_FORMAT | 8 | Specify the console output format. (default text)<br><br>Options are: json, text. |

#### license Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| license.devopsKey | --devops-key / -k | PINGCLI_LICENSE_DEVOPS_KEY | 14 | The DevOps key for the license request. <br><br> See https://developer.pingidentity.com/devops/how-to/devopsRegistration.html on how to register a DevOps user. <br><br> You can save the DevOps user and key in your profile using the 'pingcli config' commands. |
| license.devopsUser | --devops-user / -u | PINGCLI_LICENSE_DEVOPS_USER | 14 | The DevOps user for the license request. <br><br> See https://developer.pingidentity.com/devops/how-to/devopsRegistration.html on how to register a DevOps user. <br><br> You can save the DevOps user and key in your profile using the 'pingcli config' commands. |

#### request Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| request.fail | --fail / -f |  | 0 | Return non-zero exit code when HTTP custom request returns a failure status code. |
| request.service | --service / -s | PINGCLI_REQUEST_SERVICE | 13 | The Ping service (configured in the active profile) to send the custom request to.<br><br>Options are: pingone.<br><br>Example: 'pingone' |

#### service Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| service.pingFederate.adminAPIPath | --pingfederate-admin-api-path | PINGCLI_PINGFEDERATE_ADMIN_API_PATH | 14 | The PingFederate API URL path used to communicate with PingFederate's admin API. (default /pf-admin-api/v1) |
| service.pingFederate.authentication.accessTokenAuth.accessToken | --pingfederate-access-token | PINGCLI_PINGFEDERATE_ACCESS_TOKEN | 14 | The PingFederate access token used to authenticate to the PingFederate admin API when using a custom OAuth 2.0 token method. |
| service.pingFederate.authentication.basicAuth.password | --pingfederate-password | PINGCLI_PINGFEDERATE_PASSWORD | 14 | The PingFederate password used to authenticate to the PingFederate admin API when using basic authentication. |
| service.pingFederate.authentication.basicAuth.username | --pingfederate-username | PINGCLI_PINGFEDERATE_USERNAME | 14 | The PingFederate username used to authenticate to the PingFederate admin API when using basic authentication.<br><br>Example: 'administrator' |
| service.pingFederate.authentication.clientCredentialsAuth.clientID | --pingfederate-client-id | PINGCLI_PINGFEDERATE_CLIENT_ID | 14 | The PingFederate OAuth client ID used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.clientCredentialsAuth.clientSecret | --pingfederate-client-secret | PINGCLI_PINGFEDERATE_CLIENT_SECRET | 14 | The PingFederate OAuth client secret used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.clientCredentialsAuth.scopes | --pingfederate-scopes | PINGCLI_PINGFEDERATE_SCOPES | 15 | The PingFederate OAuth scopes used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. (default [])<br><br>Accepts a comma-separated string to delimit multiple scopes.<br><br>Example: 'openid,profile' |
| service.pingFederate.authentication.clientCredentialsAuth.tokenURL | --pingfederate-token-url | PINGCLI_PINGFEDERATE_TOKEN_URL | 14 | The PingFederate OAuth token URL used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.type | --pingfederate-authentication-type | PINGCLI_PINGFEDERATE_AUTHENTICATION_TYPE | 9 | The authentication type to use when connecting to the PingFederate admin API.<br><br>Options are: accessTokenAuth, basicAuth, clientCredentialsAuth.<br><br>Example: 'basicAuth' |
| service.pingFederate.caCertificatePEMFiles | --pingfederate-ca-certificate-pem-files | PINGCLI_PINGFEDERATE_CA_CERTIFICATE_PEM_FILES | 15 | Relative or full paths to PEM-encoded certificate files to be trusted as root CAs when connecting to the PingFederate server over HTTPS. (default [])<br><br>Accepts a comma-separated string to delimit multiple PEM files. |
| service.pingFederate.httpsHost | --pingfederate-https-host | PINGCLI_PINGFEDERATE_HTTPS_HOST | 14 | The PingFederate HTTPS host used to communicate with PingFederate's admin API.<br><br>Example: 'https://pingfederate-admin.bxretail.org' |
| service.pingFederate.insecureTrustAllTLS | --pingfederate-insecure-trust-all-tls | PINGCLI_PINGFEDERATE_INSECURE_TRUST_ALL_TLS | 0 | Trust any certificate when connecting to the PingFederate server admin API. (default false)<br><br>This is insecure and shouldn't be enabled outside of testing. |
| service.pingFederate.xBypassExternalValidationHeader | --pingfederate-x-bypass-external-validation-header | PINGCLI_PINGFEDERATE_X_BYPASS_EXTERNAL_VALIDATION_HEADER | 0 | Bypass connection tests when configuring PingFederate (the X-BypassExternalValidation header when using PingFederate's admin API). (default false) |
| service.pingOne.authentication.authorizationCode.clientID | --pingone-oidc-authorization-code-client-id | PINGCLI_PINGONE_OIDC_AUTHORIZATION_CODE_CLIENT_ID | 16 | The authorization code client ID used to authenticate to the PingOne management API. |
| service.pingOne.authentication.authorizationCode.redirectURIPath | --pingone-oidc-authorization-code-redirect-uri-path | PINGCLI_PINGONE_OIDC_AUTHORIZATION_CODE_REDIRECT_URI_PATH | 14 | The redirect URI path to use when using the authorization code authorization grant type to authenticate to the PingOne management API. |
| service.pingOne.authentication.authorizationCode.redirectURIPort | --pingone-oidc-authorization-code-redirect-uri-port | PINGCLI_PINGONE_OIDC_AUTHORIZATION_CODE_REDIRECT_URI_PORT | 14 | The redirect URI port to use when using the authorization code authorization grant type to authenticate to the PingOne management API. |
| service.pingOne.authentication.clientCredentials.clientID | --pingone-client-credentials-client-id | PINGCLI_PINGONE_CLIENT_CREDENTIALS_CLIENT_ID | 16 | The client credentials client ID used to authenticate to the PingOne management API. |
| service.pingOne.authentication.clientCredentials.clientSecret | --pingone-client-credentials-client-secret | PINGCLI_PINGONE_CLIENT_CREDENTIALS_CLIENT_SECRET | 14 | The client credentials client secret used to authenticate to the PingOne management API. |
| service.pingOne.authentication.deviceCode.clientID | --pingone-device-code-client-id | PINGCLI_PINGONE_DEVICE_CODE_CLIENT_ID | 16 | The device code client ID used to authenticate to the PingOne management API. |
| service.pingOne.authentication.environmentID | --pingone-environment-id | PINGCLI_PINGONE_ENVIRONMENT_ID | 16 | The ID of the PingOne environment to use for authentication (used by all auth types). |
| service.pingOne.authentication.type | --pingone-authentication-type | PINGCLI_PINGONE_AUTHENTICATION_TYPE | 10 | The authorization grant type to use to authenticate to the PingOne management API. (default worker)<br><br>Options are: authorization_code, client_credentials, device_code, worker. |
| service.pingOne.authentication.worker.clientID | --pingone-worker-client-id | PINGCLI_PINGONE_WORKER_CLIENT_ID | 16 | DEPRECATED: Use --pingone-client-credentials-client-id instead. The worker client ID used to authenticate to the PingOne management API. |
| service.pingOne.authentication.worker.clientSecret | --pingone-worker-client-secret | PINGCLI_PINGONE_WORKER_CLIENT_SECRET | 14 | DEPRECATED: Use --pingone-client-credentials-client-secret instead. The worker client secret used to authenticate to the PingOne management API. |
| service.pingOne.authentication.worker.environmentID | --pingone-worker-environment-id | PINGCLI_PINGONE_WORKER_ENVIRONMENT_ID | 16 | DEPRECATED: Use --pingone-environment-id instead. The ID of the PingOne environment that contains the worker client used to authenticate to the PingOne management API. |
| service.pingOne.regionCode | --pingone-region-code | PINGCLI_PINGONE_REGION_CODE | 11 | The region code of the PingOne tenant.<br><br>Options are: AP, AU, CA, EU, NA, SG.<br><br>Example: 'NA' |

