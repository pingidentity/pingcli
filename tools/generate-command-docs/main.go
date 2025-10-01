package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/pingidentity/pingcli/cmd"
	"github.com/pingidentity/pingcli/tools/docutil"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func main() {
	outDir := flag.String("o", "./docs", "Output directory for AsciiDoc command pages")
	date := flag.String("date", time.Now().Format("January 2, 2006"), "Created/revision date used in headers (e.g., March 23, 2025)")
	resourcePrefix := flag.String("resource-prefix", "pingcli_command_reference_", "Prefix for :resourceid:")
	version := flag.String("version", "dev", "Version string for root command init")
	commit := flag.String("commit", "dev", "Commit SHA for root command init")
	flag.Parse()

	root := cmd.NewRootCommand(*version, *commit)
	root.DisableAutoGenTag = true

	// Use tighter directory permissions (group readable/executable only) to satisfy gosec G301.
	// Not huge since these are just docs, but still better to be consistent.
	if err := os.MkdirAll(*outDir, 0o750); err != nil {
		fail("create out dir", err)
	}

	// One file per command path with deterministic, content-based updates.
	walkVisible(root, func(c *cobra.Command) {
		base := strings.ReplaceAll(c.CommandPath(), " ", "_")
		file := filepath.Join(*outDir, base+".adoc")

		// If file exists, extract existing created-date so it is preserved.
		var existingCreated string
		if oldRaw, err := readFileIfWithin(file, *outDir); err == nil {
			existingCreated = docutil.ExtractDateLine(string(oldRaw), ":created-date:")
		}
		createdDate := *date
		if existingCreated != "" {
			createdDate = existingCreated
		}

		content := renderSingle(c, createdDate, *date, *resourcePrefix)

		// Determine if underlying (non-date) content actually changed; if not, skip rewrite.
		var prevBody string
		if oldRaw, err := readFileIfWithin(file, *outDir); err == nil {
			prevBody = docutil.NormalizeForCompare(string(oldRaw))
		}
		newBody := docutil.NormalizeForCompare(content)
		if prevBody == newBody && prevBody != "" {
			// Skip updating revision date to avoid needless churn.
			return
		}

		// Restrict file permissions (no world access) for consistency with directory perms.
		if err := os.WriteFile(file, []byte(content), 0o600); err != nil {
			fail("write file "+file, err)
		}
	})

	// Navigation file: only write if changed to keep diffs minimal.
	navPath := filepath.Join(*outDir, "nav.adoc")
	navContent := renderNav(root)
	if oldNav, err := readFileIfWithin(navPath, *outDir); err == nil {
		if string(oldNav) == navContent {
			// Unchanged
			return
		}
	}
	if err := os.WriteFile(navPath, []byte(navContent), 0o600); err != nil {
		fail("write nav file", err)
	}
}

func renderSingle(c *cobra.Command, createdDate, revDate, resourcePrefix string) string {
	type singlePageData struct {
		CommandPath      string
		CreatedDate      string
		RevDate          string
		ResourceID       string
		Short            string
		Synopsis         string
		Use              string
		ExampleBlock     string
		HasLocal         bool
		LocalOptions     string
		HasInherited     bool
		InheritedOptions string
		ParentBlock      string
		SubcommandsBlock string
	}

	// Precompute fields exactly matching previous output.
	base := strings.ReplaceAll(c.CommandPath(), " ", "_")
	short := strings.TrimSpace(firstLine(c.Short, c.Long))
	var synopsis string
	if long := strings.TrimSpace(c.Long); long != "" {
		synopsis = long + "\n\n"
	} else if s := strings.TrimSpace(c.Short); s != "" {
		synopsis = s + "\n\n"
	}
	use := strings.TrimSpace(c.UseLine())
	var exampleBlock string
	if rawEx := c.Example; strings.TrimSpace(rawEx) != "" {
		var eb strings.Builder
		eb.WriteString("== Examples\n\n")
		eb.WriteString("----\n")
		eb.WriteString(rawEx)
		if !strings.HasSuffix(rawEx, "\n") {
			eb.WriteString("\n")
		}
		eb.WriteString("----\n\n")
		exampleBlock = eb.String()
	}

	local := c.NonInheritedFlags()
	inherited := c.InheritedFlags()
	hasLocal := local != nil && local.HasAvailableFlags()
	hasInherited := inherited != nil && inherited.HasAvailableFlags()
	var localBlock, inheritedBlock string
	if hasLocal {
		localBlock = formatFlagBlock(local, true, c)
	}
	if hasInherited {
		inheritedBlock = formatFlagBlock(inherited, false, c)
	}

	var parentBlock string
	if p := c.Parent(); p != nil {
		parentFile := strings.ReplaceAll(p.CommandPath(), " ", "_") + ".adoc"
		var pb strings.Builder
		pb.WriteString("== More information\n\n")
		fmt.Fprintf(&pb, "* xref:%s[]\t - %s\n\n", parentFile, firstLine(p.Short, p.Long))
		parentBlock = pb.String()
	}

	var subcommandsBlock string
	subs := visibleSubcommands(c)
	if len(subs) > 0 {
		sort.Slice(subs, func(i, j int) bool { return subs[i].Name() < subs[j].Name() })
		var sb strings.Builder
		sb.WriteString("== Subcommands\n\n")
		for _, sc := range subs {
			name := strings.ReplaceAll(sc.CommandPath(), " ", "_") + ".adoc"
			fmt.Fprintf(&sb, "* xref:%s[] - %s\n", name, firstLine(sc.Short, sc.Long))
		}
		sb.WriteString("\n")
		subcommandsBlock = sb.String()
	}

	data := singlePageData{
		CommandPath:      c.CommandPath(),
		CreatedDate:      createdDate,
		RevDate:          revDate,
		ResourceID:       resourcePrefix + base,
		Short:            short,
		Synopsis:         synopsis,
		Use:              use,
		ExampleBlock:     exampleBlock,
		HasLocal:         hasLocal,
		LocalOptions:     localBlock,
		HasInherited:     hasInherited,
		InheritedOptions: inheritedBlock,
		ParentBlock:      parentBlock,
		SubcommandsBlock: subcommandsBlock,
	}

	var buf bytes.Buffer
	if err := singlePageTpl.Execute(&buf, data); err != nil {
		// Fallback should never happen; keep previous behavior if it does.
		return ""
	}

	return buf.String()
}

// formatFlagBlock renders a flag set into a code-fenced block similar to manual pages.
func formatFlagBlock(fs *pflag.FlagSet, includeHelp bool, c *cobra.Command) string {
	var flags []*pflag.Flag
	fs.VisitAll(func(f *pflag.Flag) { flags = append(flags, f) })
	sort.Slice(flags, func(i, j int) bool {
		si, sj := flags[i].Shorthand, flags[j].Shorthand
		if si == sj {
			return flags[i].Name < flags[j].Name
		}
		if si == "" {
			return false
		}
		if sj == "" {
			return true
		}

		return si < sj
	})
	type line struct {
		Spec string
		Pad  int
		Desc string
	}
	lines := make([]line, 0, len(flags))
	for _, f := range flags {
		var spec string
		if f.Shorthand != "" {
			spec = fmt.Sprintf("-%s, --%s", f.Shorthand, f.Name)
		} else {
			spec = fmt.Sprintf("    --%s", f.Name)
		}
		typeName := f.Value.Type()
		if typeName != "bool" {
			spec += " " + typeName
		}
		desc := f.Usage
		if typeName == "bool" {
			// Add only if usage text does not already contain a default and DefValue is meaningful.
			if !strings.Contains(desc, "(default") && f.DefValue != "" {
				desc = fmt.Sprintf("%s (default %s)", desc, f.DefValue)
			}
		} else if f.DefValue != "" && f.DefValue != "<nil>" && f.DefValue != "0" && !strings.Contains(desc, "(default") {
			desc = fmt.Sprintf("%s (default %s)", desc, f.DefValue)
		}

		// Collapse internal newlines but otherwise keep original spacing; no manual wrapping.
		desc = strings.ReplaceAll(desc, "\n", " ")

		lines = append(lines, line{Spec: spec, Desc: desc})
	}
	if includeHelp {
		found := false
		for _, l := range lines {
			if strings.Contains(l.Spec, "--help") {
				found = true

				break
			}
		}
		if !found {
			helpLine := line{Spec: "-h, --help", Desc: fmt.Sprintf("help for %s", c.Name())}
			if len(lines) == 0 {
				lines = append(lines, helpLine)
			} else {
				lines = append(lines[:1], append([]line{helpLine}, lines[1:]...)...)
			}
		}
	}
	maxSpec := 0
	for _, l := range lines {
		if len(l.Spec) > maxSpec {
			maxSpec = len(l.Spec)
		}
	}
	for i := range lines {
		pad := maxSpec - len(lines[i].Spec)
		if pad < 0 {
			pad = 0
		}
		lines[i].Pad = pad
	}
	var buf bytes.Buffer
	if err := flagBlockTpl.Execute(&buf, struct{ Lines []line }{Lines: lines}); err != nil {
		return ""
	}

	return buf.String()
}

// firstLine returns the first non-empty line from short or long description.
func firstLine(short, long string) string {
	if strings.TrimSpace(short) != "" {
		return strings.SplitN(strings.TrimSpace(short), "\n", 2)[0]
	}
	if strings.TrimSpace(long) != "" {
		return strings.SplitN(strings.TrimSpace(long), "\n", 2)[0]
	}

	return ""
}

func visibleSubcommands(c *cobra.Command) []*cobra.Command {
	cmds := c.Commands()
	subs := make([]*cobra.Command, 0, len(cmds))
	for _, sc := range cmds {
		if sc.Hidden {
			continue
		}
		subs = append(subs, sc)
	}

	return subs
}

func walkVisible(c *cobra.Command, fn func(*cobra.Command)) {
	fn(c)
	children := visibleSubcommands(c)
	sort.Slice(children, func(i, j int) bool { return children[i].Name() < children[j].Name() })
	for _, sc := range children {
		walkVisible(sc, fn)
	}
}

func depthOf(c *cobra.Command) int { return len(strings.Split(c.CommandPath(), " ")) }

func fail(doing string, err error) {
	fmt.Fprintf(os.Stderr, "error while %s: %v\n", doing, err)
	os.Exit(1)
}

// renderNav builds nav.adoc content with hierarchical bullet list.
// Format mirrors manually created original nav.adoc but with synthetic top-level group.
func renderNav(root *cobra.Command) string {
	var b strings.Builder
	b.WriteString("* Command Reference\n")

	// Add root command first
	rootFile := strings.ReplaceAll(root.CommandPath(), " ", "_") + ".adoc"
	fmt.Fprintf(&b, "** xref:command_reference:%s[]\n", rootFile)

	// Add all other commands
	walkVisible(root, func(c *cobra.Command) {
		if c == root {
			return
		}
		stars := strings.Repeat("*", depthOf(c))
		file := strings.ReplaceAll(c.CommandPath(), " ", "_") + ".adoc"
		fmt.Fprintf(&b, "%s xref:command_reference:%s[]\n", stars, file)
	})
	b.WriteString("\n")

	return b.String()
}

// readFileIfWithin validates that path is within base before reading to satisfy gosec G304.
func readFileIfWithin(path, base string) ([]byte, error) {
	cleanBase := filepath.Clean(base)
	cleanPath := filepath.Clean(path)
	if !strings.HasPrefix(cleanPath+string(os.PathSeparator), cleanBase+string(os.PathSeparator)) {
		return nil, fmt.Errorf("refusing to read path outside base directory: %s", path)
	}
	data, err := os.ReadFile(cleanPath) // #nosec G304 path validated above
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Templates and helpers
var singlePageTpl = template.Must(template.New("single").Parse(`= {{.CommandPath}}
:created-date: {{.CreatedDate}}
:revdate: {{.RevDate}}
:resourceid: {{.ResourceID}}

{{if .Short}}{{.Short}}

{{end}}== Synopsis

{{.Synopsis}}----
{{.Use}}
----

{{if .ExampleBlock}}{{.ExampleBlock}}{{end}}{{if .HasLocal}}== Options

{{.LocalOptions}}
{{end}}{{if .HasInherited}}== Options inherited from parent commands

{{.InheritedOptions}}
{{end}}{{if .ParentBlock}}{{.ParentBlock}}{{end}}{{if .SubcommandsBlock}}{{.SubcommandsBlock}}{{end}}`))

func repeat(n int) string { return strings.Repeat(" ", n) }

var flagBlockTpl = template.Must(template.New("flag").Funcs(template.FuncMap{"repeat": repeat}).Parse(`----
{{range .Lines}}  {{.Spec}}{{repeat .Pad}}   {{.Desc}}
{{end}}----
`))
