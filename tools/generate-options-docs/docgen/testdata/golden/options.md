#### auth Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| auth.services | --service / -s | PINGCLI_AUTH_SERVICE | 1 | Specifies the service(s) to authenticate. Accepts a comma-separated string to delimit multiple services. <br><br>Options are: pingfederate, pingone.<br><br>Example: 'pingone,pingfederate' |
| auth.useKeychain | --use-keychain | PINGCLI_AUTH_USE_KEYCHAIN | 0 | Use system keychain for storing authentication tokens. If false or keychain is unavailable, tokens will be stored in ~/.pingcli/credentials/. (default true) |

#### export Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| export.format | --format / -f | PINGCLI_EXPORT_FORMAT | 2 | Specifies the export format. (default HCL)<br><br>Options are: HCL. |
| export.outputDirectory | --output-directory / -d | PINGCLI_EXPORT_OUTPUT_DIRECTORY | 15 | Specifies the output directory for export. Can be an absolute filepath or a relative filepath of the present working directory. <br><br>Example: '/Users/example/pingcli-export'<br><br>Example: 'pingcli-export' |
| export.overwrite | --overwrite / -o | PINGCLI_EXPORT_OVERWRITE | 0 | Overwrites the existing generated exports in output directory. (default false) |
| export.pingOne.environmentID | --pingone-export-environment-id | PINGCLI_PINGONE_EXPORT_ENVIRONMENT_ID | 17 | The ID of the PingOne environment to export. Must be a valid PingOne UUID. |
| export.serviceGroup | --service-group / -g | PINGCLI_EXPORT_SERVICE_GROUP | 3 | Specifies the service group to export. <br><br>Options are: pingone.<br><br>Example: 'pingone' |
| export.services | --services / -s | PINGCLI_EXPORT_SERVICES | 4 | Specifies the service(s) to export. Accepts a comma-separated string to delimit multiple services. <br><br>Options are: pingfederate, pingone-authorize, pingone-mfa, pingone-platform, pingone-protect, pingone-sso.<br><br>Example: 'pingone-sso,pingone-mfa,pingfederate' |

#### general Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| detailedExitCode | --detailed-exitcode / -D | PINGCLI_DETAILED_EXITCODE | 0 | Enable detailed exit code output. (default false)<br><br>0 - pingcli command succeeded with no errors or warnings.<br><br>1 - pingcli command failed with errors.<br><br>2 - pingcli command succeeded with warnings. |
| noColor | --no-color | PINGCLI_NO_COLOR | 0 | Disable text output in color. (default false) |
| outputFormat | --output-format / -O | PINGCLI_OUTPUT_FORMAT | 9 | Specify the console output format. (default text)<br><br>Options are: json, text. |

#### license Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| license.devopsKey | --devops-key / -k | PINGCLI_LICENSE_DEVOPS_KEY | 15 | The DevOps key for the license request. <br><br> See https://developer.pingidentity.com/devops/how-to/devopsRegistration.html on how to register a DevOps user. <br><br> You can save the DevOps user and key in your profile using the 'pingcli config' commands. |
| license.devopsUser | --devops-user / -u | PINGCLI_LICENSE_DEVOPS_USER | 15 | The DevOps user for the license request. <br><br> See https://developer.pingidentity.com/devops/how-to/devopsRegistration.html on how to register a DevOps user. <br><br> You can save the DevOps user and key in your profile using the 'pingcli config' commands. |

#### request Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| request.fail | --fail / -f |  | 0 | Return non-zero exit code when HTTP custom request returns a failure status code. |
| request.service | --service / -s | PINGCLI_REQUEST_SERVICE | 14 | The Ping service (configured in the active profile) to send the custom request to.<br><br>Options are: pingone.<br><br>Example: 'pingone' |

#### service Properties

| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |
|---|---|---|---|---|
| service.pingFederate.adminAPIPath | --pingfederate-admin-api-path | PINGCLI_PINGFEDERATE_ADMIN_API_PATH | 15 | The PingFederate API URL path used to communicate with PingFederate's admin API. (default /pf-admin-api/v1) |
| service.pingFederate.authentication.accessTokenAuth.accessToken | --pingfederate-access-token | PINGCLI_PINGFEDERATE_ACCESS_TOKEN | 15 | The PingFederate access token used to authenticate to the PingFederate admin API when using a custom OAuth 2.0 token method. |
| service.pingFederate.authentication.basicAuth.password | --pingfederate-password | PINGCLI_PINGFEDERATE_PASSWORD | 15 | The PingFederate password used to authenticate to the PingFederate admin API when using basic authentication. |
| service.pingFederate.authentication.basicAuth.username | --pingfederate-username | PINGCLI_PINGFEDERATE_USERNAME | 15 | The PingFederate username used to authenticate to the PingFederate admin API when using basic authentication.<br><br>Example: 'administrator' |
| service.pingFederate.authentication.clientCredentialsAuth.clientID | --pingfederate-client-id | PINGCLI_PINGFEDERATE_CLIENT_ID | 15 | The PingFederate OAuth client ID used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.clientCredentialsAuth.clientSecret | --pingfederate-client-secret | PINGCLI_PINGFEDERATE_CLIENT_SECRET | 15 | The PingFederate OAuth client secret used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.clientCredentialsAuth.scopes | --pingfederate-scopes | PINGCLI_PINGFEDERATE_SCOPES | 16 | The PingFederate OAuth scopes used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. (default [])<br><br>Accepts a comma-separated string to delimit multiple scopes.<br><br>Example: 'openid,profile' |
| service.pingFederate.authentication.clientCredentialsAuth.tokenURL | --pingfederate-token-url | PINGCLI_PINGFEDERATE_TOKEN_URL | 15 | The PingFederate OAuth token URL used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.type | --pingfederate-authentication-type | PINGCLI_PINGFEDERATE_AUTHENTICATION_TYPE | 10 | The authentication type to use when connecting to the PingFederate admin API.<br><br>Options are: accessTokenAuth, basicAuth, clientCredentialsAuth.<br><br>Example: 'basicAuth' |
| service.pingFederate.caCertificatePEMFiles | --pingfederate-ca-certificate-pem-files | PINGCLI_PINGFEDERATE_CA_CERTIFICATE_PEM_FILES | 16 | Relative or full paths to PEM-encoded certificate files to be trusted as root CAs when connecting to the PingFederate server over HTTPS. (default [])<br><br>Accepts a comma-separated string to delimit multiple PEM files. |
| service.pingFederate.httpsHost | --pingfederate-https-host | PINGCLI_PINGFEDERATE_HTTPS_HOST | 15 | The PingFederate HTTPS host used to communicate with PingFederate's admin API.<br><br>Example: 'https://pingfederate-admin.bxretail.org' |
| service.pingFederate.insecureTrustAllTLS | --pingfederate-insecure-trust-all-tls | PINGCLI_PINGFEDERATE_INSECURE_TRUST_ALL_TLS | 0 | Trust any certificate when connecting to the PingFederate server admin API. (default false)<br><br>This is insecure and shouldn't be enabled outside of testing. |
| service.pingFederate.xBypassExternalValidationHeader | --pingfederate-x-bypass-external-validation-header | PINGCLI_PINGFEDERATE_X_BYPASS_EXTERNAL_VALIDATION_HEADER | 0 | Bypass connection tests when configuring PingFederate (the X-BypassExternalValidation header when using PingFederate's admin API). (default false) |
| service.pingOne.authentication.authCode.clientID | --pingone-oidc-auth-code-client-id | PINGCLI_PINGONE_OIDC_AUTH_CODE_CLIENT_ID | 17 | The auth code client ID used to authenticate to the PingOne management API. |
| service.pingOne.authentication.authCode.environmentID | --pingone-oidc-auth-code-environment-id | PINGCLI_PINGONE_OIDC_AUTH_CODE_ENVIRONMENT_ID | 17 | The ID of the PingOne environment that contains the auth code client used to authenticate to the PingOne management API. |
| service.pingOne.authentication.authCode.redirectURI | --pingone-oidc-auth-code-redirect-uri | PINGCLI_PINGONE_OIDC_AUTH_CODE_REDIRECT_URI | 15 | The redirect URI to use when using the auth code authentication type to authenticate to the PingOne management API. |
| service.pingOne.authentication.authCode.scopes | --pingone-oidc-auth-code-scopes | PINGCLI_PINGONE_OIDC_AUTH_CODE_SCOPES | 16 | The auth code scope(s) used to authenticate to the PingOne management API. |
| service.pingOne.authentication.clientCredentials.clientID | --pingone-client-credentials-client-id | PINGCLI_PINGONE_CLIENT_CREDENTIALS_CLIENT_ID | 17 | The client credentials client ID used to authenticate to the PingOne management API. |
| service.pingOne.authentication.clientCredentials.clientSecret | --pingone-client-credentials-client-secret | PINGCLI_PINGONE_CLIENT_CREDENTIALS_CLIENT_SECRET | 15 | The client credentials client secret used to authenticate to the PingOne management API. |
| service.pingOne.authentication.clientCredentials.environmentID | --pingone-client-credentials-environment-id | PINGCLI_PINGONE_CLIENT_CREDENTIALS_ENVIRONMENT_ID | 17 | The ID of the PingOne environment that contains the client credentials client used to authenticate to the PingOne management API. |
| service.pingOne.authentication.clientCredentials.scopes | --pingone-client-credentials-scopes | PINGCLI_PINGONE_CLIENT_CREDENTIALS_SCOPES | 16 | The scopes to request for the client credentials used to authenticate to the PingOne management API. |
| service.pingOne.authentication.deviceCode.clientID | --pingone-device-code-client-id | PINGCLI_PINGONE_DEVICE_CODE_CLIENT_ID | 17 | The device code client ID used to authenticate to the PingOne management API. |
| service.pingOne.authentication.deviceCode.environmentID | --pingone-device-code-environment-id | PINGCLI_PINGONE_DEVICE_CODE_ENVIRONMENT_ID | 17 | The ID of the PingOne environment that contains the device code client used to authenticate to the PingOne management API. |
| service.pingOne.authentication.deviceCode.scopes | --pingone-device-code-scopes | PINGCLI_PINGONE_DEVICE_CODE_SCOPES | 16 | The device code scope(s) used to authenticate to the PingOne management API. |
| service.pingOne.authentication.type | --pingone-authentication-type | PINGCLI_PINGONE_AUTHENTICATION_TYPE | 11 | The authentication type to use to authenticate to the PingOne management API. (default worker)<br><br>Options are: auth_code, client_credentials, device_code, worker. |
| service.pingOne.authentication.worker.clientID | --pingone-worker-client-id | PINGCLI_PINGONE_WORKER_CLIENT_ID | 17 | The worker client ID used to authenticate to the PingOne management API. |
| service.pingOne.authentication.worker.clientSecret | --pingone-worker-client-secret | PINGCLI_PINGONE_WORKER_CLIENT_SECRET | 15 | The worker client secret used to authenticate to the PingOne management API. |
| service.pingOne.authentication.worker.environmentID | --pingone-worker-environment-id | PINGCLI_PINGONE_WORKER_ENVIRONMENT_ID | 17 | The ID of the PingOne environment that contains the worker client used to authenticate to the PingOne management API. |
| service.pingOne.regionCode | --pingone-region-code | PINGCLI_PINGONE_REGION_CODE | 12 | The region code of the PingOne tenant.<br><br>Options are: AP, AU, CA, EU, NA, SG.<br><br>Example: 'NA' |

