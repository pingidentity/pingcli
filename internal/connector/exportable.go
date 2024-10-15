package connector

import (
	_ "embed"
)

// Embed import block template needed for export generation
//
//go:embed templates/hcl_import_block.template
var HCLImportBlockTemplate string

//go:embed templates/hcl_import_header.template
var HCLImportHeaderTemplate string

// A connector that allows exporting configuration
type Exportable interface {
	Export(format, outputDir string, overwriteExport bool) error
	ConnectorServiceName() string
}
