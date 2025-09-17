#### export Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| export.format | 1 | --format / -f | Specifies the export format. (default HCL)<br><br>Options are: HCL. |
| export.outputDirectory | 14 | --output-directory / -d | Specifies the output directory for export. Can be an absolute filepath or a relative filepath of the present working directory. <br><br>Example: '/Users/example/pingcli-export'<br><br>Example: 'pingcli-export' |
| export.overwrite | 0 | --overwrite / -o | Overwrites the existing generated exports in output directory. (default false) |
| export.pingOne.environmentID | 16 | --pingone-export-environment-id | The ID of the PingOne environment to export. Must be a valid PingOne UUID. |
| export.serviceGroup | 2 | --service-group / -g | Specifies the service group to export. <br><br>Options are: pingone.<br><br>Example: 'pingone' |
| export.services | 3 | --services / -s | Specifies the service(s) to export. Accepts a comma-separated string to delimit multiple services. <br><br>Options are: pingfederate, pingone-authorize, pingone-mfa, pingone-platform, pingone-protect, pingone-sso.<br><br>Example: 'pingone-sso,pingone-mfa,pingfederate' |

#### general Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| detailedExitCode | 0 | --detailed-exitcode / -D | Enable detailed exit code output. (default false)<br><br>0 - pingcli command succeeded with no errors or warnings.<br><br>1 - pingcli command failed with errors.<br><br>2 - pingcli command succeeded with warnings. |
| noColor | 0 | --no-color | Disable text output in color. (default false) |
| outputFormat | 8 | --output-format / -O | Specify the console output format. (default text)<br><br>Options are: json, text. |

#### license Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| license.devopsKey | 14 | --devops-key / -k | The DevOps key for the license request. <br><br> See https://developer.pingidentity.com/devops/how-to/devopsRegistration.html on how to register a DevOps user. <br><br> You can save the DevOps user and key in your profile using the 'pingcli config' commands. |
| license.devopsUser | 14 | --devops-user / -u | The DevOps user for the license request. <br><br> See https://developer.pingidentity.com/devops/how-to/devopsRegistration.html on how to register a DevOps user. <br><br> You can save the DevOps user and key in your profile using the 'pingcli config' commands. |

#### request Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| request.fail | 0 | --fail / -f | Return non-zero exit code when HTTP custom request returns a failure status code. |
| request.service | 13 | --service / -s | The Ping service (configured in the active profile) to send the custom request to.<br><br>Options are: pingone.<br><br>Example: 'pingone' |

#### service Properties

| Config File Property | Type | Equivalent Parameter | Purpose |
|---|---|---|---|
| service.pingFederate.adminAPIPath | 14 | --pingfederate-admin-api-path | The PingFederate API URL path used to communicate with PingFederate's admin API. (default /pf-admin-api/v1) |
| service.pingFederate.authentication.accessTokenAuth.accessToken | 14 | --pingfederate-access-token | The PingFederate access token used to authenticate to the PingFederate admin API when using a custom OAuth 2.0 token method. |
| service.pingFederate.authentication.basicAuth.password | 14 | --pingfederate-password | The PingFederate password used to authenticate to the PingFederate admin API when using basic authentication. |
| service.pingFederate.authentication.basicAuth.username | 14 | --pingfederate-username | The PingFederate username used to authenticate to the PingFederate admin API when using basic authentication.<br><br>Example: 'administrator' |
| service.pingFederate.authentication.clientCredentialsAuth.clientID | 14 | --pingfederate-client-id | The PingFederate OAuth client ID used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.clientCredentialsAuth.clientSecret | 14 | --pingfederate-client-secret | The PingFederate OAuth client secret used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.clientCredentialsAuth.scopes | 15 | --pingfederate-scopes | The PingFederate OAuth scopes used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. (default [])<br><br>Accepts a comma-separated string to delimit multiple scopes.<br><br>Example: 'openid,profile' |
| service.pingFederate.authentication.clientCredentialsAuth.tokenURL | 14 | --pingfederate-token-url | The PingFederate OAuth token URL used to authenticate to the PingFederate admin API when using the OAuth 2.0 client credentials grant type. |
| service.pingFederate.authentication.type | 9 | --pingfederate-authentication-type | The authentication type to use when connecting to the PingFederate admin API.<br><br>Options are: accessTokenAuth, basicAuth, clientCredentialsAuth.<br><br>Example: 'basicAuth' |
| service.pingFederate.caCertificatePEMFiles | 15 | --pingfederate-ca-certificate-pem-files | Relative or full paths to PEM-encoded certificate files to be trusted as root CAs when connecting to the PingFederate server over HTTPS. (default [])<br><br>Accepts a comma-separated string to delimit multiple PEM files. |
| service.pingFederate.httpsHost | 14 | --pingfederate-https-host | The PingFederate HTTPS host used to communicate with PingFederate's admin API.<br><br>Example: 'https://pingfederate-admin.bxretail.org' |
| service.pingFederate.insecureTrustAllTLS | 0 | --pingfederate-insecure-trust-all-tls | Trust any certificate when connecting to the PingFederate server admin API. (default false)<br><br>This is insecure and shouldn't be enabled outside of testing. |
| service.pingFederate.xBypassExternalValidationHeader | 0 | --pingfederate-x-bypass-external-validation-header | Bypass connection tests when configuring PingFederate (the X-BypassExternalValidation header when using PingFederate's admin API). (default false) |
| service.pingOne.authentication.type | 10 | --pingone-authentication-type | The authentication type to use to authenticate to the PingOne management API. (default worker)<br><br>Options are: worker. |
| service.pingOne.authentication.worker.clientID | 16 | --pingone-worker-client-id | The worker client ID used to authenticate to the PingOne management API. |
| service.pingOne.authentication.worker.clientSecret | 14 | --pingone-worker-client-secret | The worker client secret used to authenticate to the PingOne management API. |
| service.pingOne.authentication.worker.environmentID | 16 | --pingone-worker-environment-id | The ID of the PingOne environment that contains the worker client used to authenticate to the PingOne management API. |
| service.pingOne.regionCode | 11 | --pingone-region-code | The region code of the PingOne tenant.<br><br>Options are: AP, AU, CA, EU, NA.<br><br>Example: 'NA' |

