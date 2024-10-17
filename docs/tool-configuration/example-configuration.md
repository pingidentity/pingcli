## Example Ping CLI configuration file

```
activeprofile: default
default:
    color: true
    description: Default profile created by pingcli
    export:
        format: HCL
        outputdirectory: /Users/me/pingcli/export
        overwrite: false
        pingone:
            environmentid: 12345678-1234-1234-1234-123456789012
        services:
            - pingfederate
            - pingone-mfa
            - pingone-platform
            - pingone-protect
            - pingone-sso
    outputformat: text
    request:
        service: pingone
    service:
        pingfederate:
            adminapipath: /pf-admin-api/v1
            authentication:
                accesstokenauth:
                    accesstoken: token
                basicauth:
                    password: password
                    username: administrator
                clientcredentialsauth:
                    clientid: clientID
                    clientsecret: secret
                    scopes:
                        - openid
                        - profile
                    tokenurl: https://pingfederate-admin.bxretail.org/as/token.oauth2
                type: clientcredentialsauth
            cacertificatepemfiles: []
            httpshost: https://pingfederate-admin.bxretail.org
            insecuretrustalltls: false
            xbypassexternalvalidationheader: false
        pingone:
            authentication:
                type: worker
                worker:
                    clientid: 12345678-1234-1234-1234-123456789012
                    clientsecret: secret
                    environmentid: 12345678-1234-1234-1234-123456789012
            regioncode: NA
```