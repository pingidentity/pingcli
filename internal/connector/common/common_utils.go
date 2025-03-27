// Copyright © 2025 Ping Identity Corporation

package common

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/logger"
)

func WriteFiles(exportableResources []connector.ExportableResource, format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	// Parse the HCL import block template
	hclImportBlockTemplate, err := template.New("HCLImportBlock").Parse(connector.HCLImportBlockTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse HCL import block template. err: %s", err.Error())
	}

	for _, exportableResource := range exportableResources {
		importBlocks, err := exportableResource.ExportAll()
		if err != nil {
			return fmt.Errorf("failed to export resource %s. err: %s", exportableResource.ResourceType(), err.Error())
		}

		if len(*importBlocks) == 0 {
			// No resources exported. Avoid creating an empty import.tf file
			l.Debug().Msgf("Nothing exported for resource %s. Skipping import file generation...", exportableResource.ResourceType())
			continue
		}

		// Sort import blocks by ResourceName
		slices.SortFunc(*importBlocks, func(i, j connector.ImportBlock) int {
			return strings.Compare(i.ResourceName, j.ResourceName)
		})

		l.Debug().Msgf("Generating import file for %s resource...", exportableResource.ResourceType())

		outputFileName := fmt.Sprintf("%s.tf", exportableResource.ResourceType())
		outputFilePath := filepath.Join(outputDir, filepath.Base(outputFileName))

		// Check to see if outputFile already exists.
		// If so, default behavior is to exit and not overwrite.
		// This can be changed with the --overwrite export parameter
		_, err = os.Stat(outputFilePath)
		if err == nil && !overwriteExport {
			return fmt.Errorf("generated import file for %q already exists. Use --overwrite to overwrite existing export data", outputFileName)
		}

		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			return fmt.Errorf("failed to create export file %q. err: %s", outputFilePath, err.Error())
		}
		defer outputFile.Close()

		err = writeHeader(format, outputFilePath, outputFile)
		if err != nil {
			return err
		}

		for _, importBlock := range *importBlocks {
			// Sanitize import block "to". Add pingcli-- prefix, hexidecimal encode special chars and spaces
			importBlock.Sanitize()

			switch format {
			case customtypes.ENUM_EXPORT_FORMAT_HCL:
				err := hclImportBlockTemplate.Execute(outputFile, importBlock)
				if err != nil {
					return fmt.Errorf("failed to write import template to file %q. err: %s", outputFilePath, err.Error())
				}
			default:
				return fmt.Errorf("unrecognized export format %q. Must be one of: %s", format, customtypes.ExportFormatValidValues())
			}
		}
	}
	return nil
}

func writeHeader(format, outputFilePath string, outputFile *os.File) error {
	// Parse the HCL header
	hclImportHeaderTemplate, err := template.New("HCLImportHeader").Parse(connector.HCLImportHeaderTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse HCL import header template. err: %s", err.Error())
	}

	switch format {
	case customtypes.ENUM_EXPORT_FORMAT_HCL:
		err := hclImportHeaderTemplate.Execute(outputFile, nil)
		if err != nil {
			return fmt.Errorf("failed to write import template to file %q. err: %s", outputFilePath, err.Error())
		}
	default:
		return fmt.Errorf("unrecognized export format %q. Must be one of: %s", format, customtypes.ExportFormatValidValues())
	}

	return nil
}
