package main

// This file contains the entire options documentation generator tool.
// It was consolidated from two separate files (main.go and docgen.go) in September 2025
// when cleaning up references to the removed internal/options package.
// The tool generates documentation for all CLI configuration options in either
// Markdown or AsciiDoc format.

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
)

// A tiny standalone tool (invoked via `make generate-options-docs`) to output
// documentation for configuration options in either Markdown or AsciiDoc.
func main() {
	outFile := flag.String("o", "", "Write output to file (extension .md or .adoc determines format unless flags override)")
	asAsciiDoc := flag.Bool("asciidoc", false, "Force AsciiDoc output (default: Markdown unless output file has .adoc/.asciidoc)")
	flag.Parse()

	// Decide format
	useAscii := false
	if *outFile != "" {
		useAscii = shouldOutputAsciiDoc(*outFile, *asAsciiDoc)
	} else if *asAsciiDoc {
		useAscii = true
	}

	var content string
	if useAscii {
		content = asciiDoc()
	} else {
		content = markdown()
	}

	if *outFile == "" {
		fmt.Print(content)
		return
	}

	if err := os.WriteFile(*outFile, []byte(content), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file: %v\n", err)
		os.Exit(1)
	}
}

// ---- merged from former docgen.go ----

// markdown renders the options documentation markdown table sections.
func markdown() string {
	configuration.InitAllOptions()
	propertyCategoryInformation := make(map[string][]string)
	for _, option := range options.Options() {
		if option.KoanfKey == "" || option.Flag == nil {
			continue
		}
		var flagInfo string
		if option.Flag.Shorthand != "" {
			flagInfo = fmt.Sprintf("--%s / -%s", option.CobraParamName, option.Flag.Shorthand)
		} else {
			flagInfo = fmt.Sprintf("--%s", option.CobraParamName)
		}
		usageString := strings.ReplaceAll(option.Flag.Usage, "\n", "<br><br>")
		category := "general"
		if strings.Contains(option.KoanfKey, ".") {
			category = strings.Split(option.KoanfKey, ".")[0]
		}
		propertyCategoryInformation[category] = append(propertyCategoryInformation[category], fmt.Sprintf("| %s | %d | %s | %s |", option.KoanfKey, option.Type, flagInfo, usageString))
	}
	var outputBuilder strings.Builder
	cats := make([]string, 0, len(propertyCategoryInformation))
	for k := range propertyCategoryInformation {
		cats = append(cats, k)
	}
	slices.Sort(cats)
	for _, category := range cats {
		properties := propertyCategoryInformation[category]
		slices.Sort(properties)
		outputBuilder.WriteString(fmt.Sprintf("#### %s Properties\n\n", category))
		outputBuilder.WriteString("| Config File Property | Type | Equivalent Parameter | Purpose |\n")
		outputBuilder.WriteString("|---|---|---|---|\n")
		for _, property := range properties {
			outputBuilder.WriteString(property + "\n")
		}
		outputBuilder.WriteString("\n")
	}
	return outputBuilder.String()
}

// asciiDoc generates a configuration reference in AsciiDoc format.
func asciiDoc() string {
	configuration.InitAllOptions()
	catMap := map[string][]options.Option{}
	for _, opt := range options.Options() {
		if opt.KoanfKey == "" {
			continue
		}
		root := opt.KoanfKey
		if strings.Contains(root, ".") {
			root = strings.Split(root, ".")[0]
		}
		switch root {
		case "service":
			catMap["service"] = append(catMap["service"], opt)
		case "export":
			catMap["export"] = append(catMap["export"], opt)
		case "license":
			catMap["license"] = append(catMap["license"], opt)
		case "request":
			catMap["request"] = append(catMap["request"], opt)
		default:
			if !strings.Contains(opt.KoanfKey, ".") {
				catMap["general"] = append(catMap["general"], opt)
			}
		}
	}
	for k := range catMap {
		slices.SortFunc(catMap[k], func(a, b options.Option) int { return strings.Compare(a.KoanfKey, b.KoanfKey) })
	}
	var b strings.Builder
	created := "March 23, 2025"
	revdate := time.Now().Format("January 2, 2006")
	b.WriteString("= Configuration Settings Reference\n")
	b.WriteString(fmt.Sprintf(":created-date: %s\n", created))
	b.WriteString(fmt.Sprintf(":revdate: %s\n", revdate))
	b.WriteString(":resourceid: pingcli_configuration_settings_reference\n\n")
	b.WriteString("The following configuration settings can be applied when using Ping CLI.\n\n")
	b.WriteString("The following configuration settings can be applied by using the xref:command_reference:pingcli_config_set.adoc[`config set` command] to persist the configuration value for a given **Configuration Key** in the Ping CLI configuration file.\n\n")
	b.WriteString("The configuration file is created at `.pingcli/config.yaml` in the user's home directory.\n\n")
	ordered := []struct{ key, title string }{{"general", "General Properties"}, {"service", "Ping Identity platform service properties"}, {"export", "Platform export properties"}, {"license", "License properties"}, {"request", "Custom request properties"}}
	for _, sec := range ordered {
		opts := catMap[sec.key]
		if len(opts) == 0 {
			continue
		}
		b.WriteString("== " + sec.title + "\n\n")
		b.WriteString("[cols=\"2,1,2,2\"]\n|===\n")
		b.WriteString("|Configuration Key |Data Type |Equivalent Parameter |Purpose\n\n")
		for _, opt := range opts {
			key := normalizeAsciiDocKeyLocal(opt.KoanfKey)
			dataType := asciiDocDataTypeLocal(opt)
			eqParam := asciiDocEquivalentParameterLocal(opt)
			purpose := sanitizeUsageLocal(opt)
			b.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s\n", key, dataType, eqParam, purpose))
		}
		b.WriteString("|===\n\n")
	}
	return b.String()
}

func shouldOutputAsciiDoc(outPath string, explicit bool) bool {
	if explicit {
		return true
	}
	ext := strings.ToLower(filepath.Ext(outPath))
	return ext == ".adoc" || ext == ".asciidoc"
}

func asciiDocEquivalentParameterLocal(opt options.Option) string {
	if opt.Flag == nil {
		return ""
	}
	if opt.Flag.Shorthand != "" {
		return fmt.Sprintf("`--%s` / `-%s`", opt.CobraParamName, opt.Flag.Shorthand)
	}
	return fmt.Sprintf("`--%s`", opt.CobraParamName)
}

func asciiDocDataTypeLocal(opt options.Option) string {
	switch opt.Type {
	case options.BOOL:
		return "Boolean"
	case options.STRING:
		return "String"
	case options.STRING_SLICE, options.EXPORT_SERVICES, options.HEADER:
		return "String Array"
	case options.UUID:
		return "String (UUID Format)"
	case options.EXPORT_FORMAT, options.OUTPUT_FORMAT, options.PINGFEDERATE_AUTH_TYPE, options.PINGONE_AUTH_TYPE, options.PINGONE_REGION_CODE, options.REQUEST_SERVICE, options.EXPORT_SERVICE_GROUP:
		return "String (Enum)"
	case options.INT:
		return "Integer"
	case options.LICENSE_PRODUCT, options.LICENSE_VERSION:
		return "String (Enum)"
	default:
		return "String"
	}
}

func sanitizeUsageLocal(opt options.Option) string {
	if opt.Flag == nil {
		return ""
	}
	usage := opt.Flag.Usage
	usage = strings.ReplaceAll(usage, "<br><br>", " ")
	usage = strings.ReplaceAll(usage, "\n", " ")
	usage = strings.TrimSpace(usage)

	// Word wrap at approximately 100 characters
	if len(usage) > 100 {
		words := strings.Fields(usage)
		var wrapped strings.Builder
		lineLength := 0

		for i, word := range words {
			// If adding this word exceeds our limit and it's not the first word in the line
			if lineLength+len(word) > 100 && lineLength > 0 {
				wrapped.WriteString("\n")
				lineLength = 0
			}

			// Add the word
			if i > 0 && lineLength > 0 {
				wrapped.WriteString(" ")
				lineLength++
			}
			wrapped.WriteString(word)
			lineLength += len(word)
		}

		return wrapped.String()
	}

	return usage
}

func normalizeAsciiDocKeyLocal(key string) string {
	key = strings.ReplaceAll(key, "pingFederate", "pingfederate")
	key = strings.ReplaceAll(key, "pingOne", "pingone")
	key = strings.ReplaceAll(key, "PEMFiles", "PemFiles")
	return key
}
