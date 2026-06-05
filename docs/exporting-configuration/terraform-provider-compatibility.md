# Ping CLI Platform Export Terraform Provider Support

The platform export command exports Configuration as Code packages for the Ping Platform. The CLI can export Terraform HCL to use with released Terraform providers.

See [HCL Export Compatibility](./hcl-export-compatibility.md) to learn more about which Ping platform services are
supported in export.

The following describes the CLI's support of each of these Ping Terraform providers.

| Service             | Support | Versions |
| ------------------- | ---- | ------- |
| DaVinci | :x: | TBD |
| PingDirectory | :x: | TBD |
| PingFederate | :large_orange_diamond: | >= 1.0.0 |
| PingOne | :large_orange_diamond: | >= 1.1.1 |


Key:
* :white_check_mark: - Supported / Released
* :large_orange_diamond: - Partial support / In progress
* :x: - No current support / Roadmap