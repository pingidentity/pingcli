// Copyright © 2025 Ping Identity Corporation

package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/logger"
)

var (
	utilsErrorPrefix = "connector common utils error"
)

func WriteFiles(exportableResources []connector.ExportableResource, format, outputDir string, overwriteExport bool) (err error) {
	l := logger.Get()

	// Parse the HCL import block template
	hclImportBlockTemplate, err := template.New("HCLImportBlock").Parse(connector.HCLImportBlockTemplate)
	if err != nil {
		return &errs.PingCLIError{Prefix: utilsErrorPrefix, Err: fmt.Errorf("%w: %w", ErrParseHCLTemplate, err)}
	}

	for _, exportableResource := range exportableResources {
		importBlocks, err := exportableResource.ExportAll()
		if err != nil {
			return &errs.PingCLIError{Prefix: utilsErrorPrefix, Err: fmt.Errorf("%w '%s': %w", ErrExportResource, exportableResource.ResourceType(), err)}
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
		outputFilePath = filepath.Clean(outputFilePath)

		// Check to see if outputFile already exists.
		// If so, default behavior is to exit and not overwrite.
		// This can be changed with the --overwrite export parameter
		_, err = os.Stat(outputFilePath)
		if err == nil && !overwriteExport {
			return &errs.PingCLIError{Prefix: utilsErrorPrefix, Err: fmt.Errorf("%w for %q", ErrFileAlreadyExists, outputFileName)}
		}

		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			return &errs.PingCLIError{Prefix: utilsErrorPrefix, Err: fmt.Errorf("%w for %q: %w", ErrCreateExportFile, outputFileName, err)}
		}
		defer func() {
			cErr := outputFile.Close()
			if cErr != nil {
				err = errors.Join(err, cErr)
			}
		}()

		err = writeHeader(format, outputFilePath, outputFile)
		if err != nil {
			return err
		}

		for _, importBlock := range *importBlocks {
			// Sanitize import block "to". Add pingcli__ prefix, hexadecimal encode special chars and spaces
			importBlock.Sanitize()

			switch format {
			case customtypes.ENUM_EXPORT_FORMAT_HCL:
				err := hclImportBlockTemplate.Execute(outputFile, importBlock)
				if err != nil {
					return &errs.PingCLIError{Prefix: utilsErrorPrefix, Err: fmt.Errorf("%w for %q: %w", ErrWriteTemplateToFile, outputFileName, err)}
				}
			default:
				return &errs.PingCLIError{Prefix: utilsErrorPrefix, Err: fmt.Errorf("%w '%s': must be one of '%s'", ErrUnrecognizedExportFormat, format, customtypes.ExportFormatValidValues())}
			}
		}
	}

	return nil
}

func writeHeader(format, outputFilePath string, outputFile *os.File) error {
	// Parse the HCL header
	hclImportHeaderTemplate, err := template.New("HCLImportHeader").Parse(connector.HCLImportHeaderTemplate)
	if err != nil {
		return &errs.PingCLIError{Prefix: utilsErrorPrefix, Err: fmt.Errorf("%w: %w", ErrParseHCLTemplate, err)}
	}

	switch format {
	case customtypes.ENUM_EXPORT_FORMAT_HCL:
		err := hclImportHeaderTemplate.Execute(outputFile, nil)
		if err != nil {
			return &errs.PingCLIError{Prefix: utilsErrorPrefix, Err: fmt.Errorf("%w for %q: %w", ErrWriteTemplateToFile, outputFilePath, err)}
		}
	default:
		return &errs.PingCLIError{Prefix: utilsErrorPrefix, Err: fmt.Errorf("%w '%s': must be one of '%s'", ErrUnrecognizedExportFormat, format, customtypes.ExportFormatValidValues())}
	}

	return nil
}
