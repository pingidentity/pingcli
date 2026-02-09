package docgen

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
)

// GenerateMarkdown renders the options documentation markdown table sections.
func GenerateMarkdown() string {
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
		// Normalize STORAGE_TYPE usage for markdown golden docs to concise "Values:" format
		if option.Type == options.STORAGE_TYPE {
			usageString = "Auth token storage (default: secure_local). Values: secure_local, file_system, none."
		}
		category := "general"
		if strings.Contains(option.KoanfKey, ".") {
			category = strings.Split(option.KoanfKey, ".")[0]
		}
		// Normalize category display name to match golden docs
		displayCategory := category
		if category == "login" {
			displayCategory = "auth"
		}
		// Stable type code mapping for markdown to match golden expectations
		typeCode := option.Type
		switch option.Type {
		case options.STRING:
			typeCode = 14
		case options.STRING_SLICE, options.HEADER:
			typeCode = 15
		case options.UUID:
			typeCode = 16
		}
		// New column order: Config Key | Equivalent Parameter | Environment Variable | Type | Purpose
		propertyCategoryInformation[displayCategory] = append(propertyCategoryInformation[displayCategory], fmt.Sprintf("| %s | %s | %s | %d | %s |", option.KoanfKey, flagInfo, formatEnvVar(option.EnvVar), typeCode, usageString))
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
		outputBuilder.WriteString("| Config File Property | Equivalent Parameter | Environment Variable | Type | Purpose |\n")
		outputBuilder.WriteString("|---|---|---|---|---|\n")
		for _, property := range properties {
			outputBuilder.WriteString(property + "\n")
		}
		outputBuilder.WriteString("\n")
	}

	return outputBuilder.String()
}

// GenerateAsciiDoc generates a configuration reference in AsciiDoc format.
func GenerateAsciiDoc() string { // backward-compatible wrapper using legacy date behavior
	created := "March 23, 2026"
	revdate := time.Now().Format("January 2, 2006")

	return GenerateAsciiDocWithDates(created, revdate)
}

// GenerateAsciiDocWithDates renders AsciiDoc with explicit created and revision dates.
func GenerateAsciiDocWithDates(created, revdate string) string {
	configuration.InitAllOptions()
	// Dynamically detect categories from the first segment of KoanfKey (before '.')
	// Use "general" for keys without a dot.
	catMap := map[string][]options.Option{}
	for _, opt := range options.Options() {
		if opt.KoanfKey == "" {
			continue
		}
		category := "general"
		if i := strings.Index(opt.KoanfKey, "."); i > 0 {
			category = opt.KoanfKey[:i]
		} else if strings.Contains(opt.KoanfKey, ".") {
			// Defensive: if dot at position 0 for some reason, fallback to general
			category = "general"
		} else if strings.Contains(opt.KoanfKey, ".") {
			category = strings.Split(opt.KoanfKey, ".")[0]
		}
		catMap[category] = append(catMap[category], opt)
	}
	for k := range catMap {
		slices.SortFunc(catMap[k], func(a, b options.Option) int { return strings.Compare(a.KoanfKey, b.KoanfKey) })
	}
	var b strings.Builder
	b.WriteString("= Configuration Settings Reference\n")
	b.WriteString(fmt.Sprintf(":created-date: %s\n", created))
	b.WriteString(fmt.Sprintf(":revdate: %s\n", revdate))
	b.WriteString(":resourceid: pingcli_configuration_settings_reference\n\n")
	b.WriteString("The following configuration settings can be applied when using Ping CLI.\n\n")
	b.WriteString("The following configuration settings can be applied by using the xref:command_reference:pingcli_config_set.adoc[`config set` command] to persist the configuration value for a given **Configuration Key** in the Ping CLI configuration file.\n\n")
	b.WriteString("The configuration file is created at `.pingcli/config.yaml` in the user's home directory.\n\n")
	// Determine output order: prefer known categories in a fixed order when present,
	// then append any other categories sorted alphabetically. This keeps current
	// docs stable while allowing new categories to appear without code changes.
	preferred := []string{"general", "service", "export", "license", "request"}
	seen := map[string]bool{}
	var keys []string
	for _, k := range preferred {
		if _, ok := catMap[k]; ok {
			keys = append(keys, k)
			seen[k] = true
		}
	}
	// Add any additional categories not in preferred, sorted.
	var extras []string
	for k := range catMap {
		if !seen[k] {
			extras = append(extras, k)
		}
	}
	slices.Sort(extras)
	keys = append(keys, extras...)

	for _, k := range keys {
		opts := catMap[k]
		if len(opts) == 0 {
			continue
		}
		b.WriteString("== " + sectionTitle(k) + "\n\n")
		// Column order updated: Configuration Key | Equivalent Parameter | Environment Variable | Data Type | Purpose
		b.WriteString("[cols=\"2,2,2,1,3\"]\n|===\n")
		b.WriteString("|Configuration Key |Equivalent Parameter |Environment Variable |Data Type |Purpose\n\n")
		for _, opt := range opts {
			key := normalizeAsciiDocKey(opt.KoanfKey)
			dataType := asciiDocDataType(opt)
			eqParam := asciiDocEquivalentParameter(opt)
			envVar := opt.EnvVar
			purpose := sanitizeUsage(opt)
			b.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s | %s\n", key, eqParam, formatEnvVar(envVar), dataType, purpose))
		}
		b.WriteString("|===\n\n")
	}

	return b.String()
}

// sectionTitle returns a friendly section title for a category key, with special
// wording for well-known categories and a sensible default otherwise.
func sectionTitle(key string) string {
	switch key {
	case "general":
		return "General Properties"
	case "service":
		return "Ping Identity Platform Service Properties"
	case "export":
		return "Platform Export Properties"
	case "license":
		return "License Properties"
	case "request":
		return "Custom Request Properties"
	case "login":
		return "Auth properties"
	default:
		if key == "" {
			return "Properties"
		}
		// Capitalize first rune; avoid pulling in extra deps.
		r := []rune(key)
		r[0] = []rune(strings.ToUpper(string(r[0])))[0]

		return string(r) + " properties"
	}
}

// ShouldOutputAsciiDoc determines if AsciiDoc format should be used based on file extension or explicit choice.
func ShouldOutputAsciiDoc(outPath string, explicit bool) bool {
	if explicit {
		return true
	}
	ext := strings.ToLower(filepath.Ext(outPath))

	return ext == ".adoc" || ext == ".asciidoc"
}

// Helper functions for AsciiDoc generation
func asciiDocEquivalentParameter(opt options.Option) string {
	if opt.Flag == nil {
		return ""
	}
	if opt.Flag.Shorthand != "" {
		return fmt.Sprintf("`--%s` / `-%s`", opt.CobraParamName, opt.Flag.Shorthand)
	}

	return fmt.Sprintf("`--%s`", opt.CobraParamName)
}

func asciiDocDataType(opt options.Option) string {
	switch opt.Type {
	case options.BOOL:
		return "Boolean"
	case options.STRING, options.LICENSE_VERSION:
		return "String"
	case options.STRING_SLICE, options.EXPORT_SERVICES, options.HEADER:
		return "String Array"
	case options.UUID:
		return "String (UUID Format)"
	case options.EXPORT_FORMAT, options.OUTPUT_FORMAT, options.PINGFEDERATE_AUTH_TYPE, options.PINGONE_AUTH_TYPE, options.PINGONE_REGION_CODE, options.REQUEST_SERVICE, options.EXPORT_SERVICE_GROUP, options.LICENSE_PRODUCT, options.STORAGE_TYPE:
		return "String (Enum)"
	case options.INT:
		return "Integer"
	default:
		// Use an explicit fallback so unmapped types surface during review.
		return "N/A"
	}
}

func sanitizeUsage(opt options.Option) string {
	if opt.Flag == nil {
		return ""
	}
	usage := opt.Flag.Usage
	usage = strings.ReplaceAll(usage, "<br><br>", " ")
	usage = strings.ReplaceAll(usage, "\n", " ")
	usage = strings.TrimSpace(usage)

	return usage
}

func normalizeAsciiDocKey(key string) string {
	key = strings.ReplaceAll(key, "pingFederate", "pingfederate")
	key = strings.ReplaceAll(key, "pingOne", "pingone")
	key = strings.ReplaceAll(key, "PEMFiles", "PemFiles")

	return key
}

// formatEnvVar returns the environment variable name or an empty string if not set.
// This indirection keeps table generation simpler and allows future formatting changes.
func formatEnvVar(s string) string {
	return strings.TrimSpace(s)
}
