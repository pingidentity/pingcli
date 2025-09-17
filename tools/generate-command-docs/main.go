package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/pingidentity/pingcli/cmd"
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

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		fail("create out dir", err)
	}

	// One file per command path.
	walkVisible(root, func(c *cobra.Command) {
		base := strings.ReplaceAll(c.CommandPath(), " ", "_")
		file := filepath.Join(*outDir, base+".adoc")
		depth := depthOf(c)
		content := renderSingle(c, depth, *date, *resourcePrefix)
		if err := os.WriteFile(file, []byte(content), 0o644); err != nil {
			fail("write file "+file, err)
		}
	})

	// Always (re)generate navigation file for documentation portal ingestion.
	navPath := filepath.Join(*outDir, "nav.adoc")
	navContent := renderNav(root)
	if err := os.WriteFile(navPath, []byte(navContent), 0o644); err != nil {
		fail("write nav file", err)
	}
}

func renderSingle(c *cobra.Command, depth int, date, resourcePrefix string) string {
	// Manual style: always use top-level title '=' regardless of hierarchy.
	base := strings.ReplaceAll(c.CommandPath(), " ", "_")
	b := &strings.Builder{}
	fmt.Fprintf(b, "= %s\n", c.CommandPath())
	fmt.Fprintf(b, ":created-date: %s\n:revdate: %s\n:resourceid: %s%s\n\n", date, date, resourcePrefix, base)

	// Short description (first paragraph only)
	if s := strings.TrimSpace(firstLine(c.Short, c.Long)); s != "" {
		b.WriteString(s)
		b.WriteString("\n\n")
	}

	// Synopsis section: prefer full Long (without first line duplication) else Short.
	b.WriteString("== Synopsis\n\n")
	if long := strings.TrimSpace(c.Long); long != "" {
		// Keep full long description as-is.
		b.WriteString(long)
		b.WriteString("\n\n")
	} else if short := strings.TrimSpace(c.Short); short != "" {
		b.WriteString(short + "\n\n")
	}
	// Usage block
	b.WriteString("----\n")
	b.WriteString(strings.TrimSpace(c.UseLine()) + "\n")
	b.WriteString("----\n\n")

	// Examples section (if any) - preserve original indentation & spacing.
	if rawEx := c.Example; strings.TrimSpace(rawEx) != "" {
		b.WriteString("== Examples\n\n")
		b.WriteString("----\n")
		b.WriteString(rawEx)
		if !strings.HasSuffix(rawEx, "\n") {
			b.WriteString("\n")
		}
		b.WriteString("----\n\n")
	}

	// Options (non-inherited) including help flag.
	local := c.NonInheritedFlags()
	inherited := c.InheritedFlags()
	if local != nil && local.HasAvailableFlags() {
		b.WriteString("== Options\n\n")
		b.WriteString(formatFlagBlock(local, true, c))
		b.WriteString("\n")
	}
	if inherited != nil && inherited.HasAvailableFlags() {
		b.WriteString("== Options inherited from parent commands\n\n")
		b.WriteString(formatFlagBlock(inherited, false, c))
		b.WriteString("\n")
	}

	// More information (link back to parent) if there is a parent (omit for root).
	if p := c.Parent(); p != nil {
		parentFile := strings.ReplaceAll(p.CommandPath(), " ", "_") + ".adoc"
		b.WriteString("== More information\n\n")
		fmt.Fprintf(b, "* xref:%s[]\t - %s\n", parentFile, firstLine(p.Short, p.Long))
		b.WriteString("\n")
	}

	// Subcommands listing (retain for commands that have them; use manual style heading depth)
	subs := visibleSubcommands(c)
	if len(subs) > 0 {
		b.WriteString("== Subcommands\n\n")
		sort.Slice(subs, func(i, j int) bool { return subs[i].Name() < subs[j].Name() })
		for _, sc := range subs {
			name := strings.ReplaceAll(sc.CommandPath(), " ", "_") + ".adoc"
			fmt.Fprintf(b, "* xref:%s[] - %s\n", name, firstLine(sc.Short, sc.Long))
		}
		b.WriteString("\n")
	}

	return b.String()
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
	type line struct{ spec, desc string }
	var lines []line
	for _, f := range flags {
		spec := ""
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
			desc = fmt.Sprintf("%s (default %s)", desc, f.DefValue)
		} else if f.DefValue != "" && f.DefValue != "<nil>" && f.DefValue != "0" && !strings.Contains(desc, "(default") {
			desc = fmt.Sprintf("%s (default %s)", desc, f.DefValue)
		}
		desc = strings.ReplaceAll(desc, "\n", " ")
		lines = append(lines, line{spec: spec, desc: desc})
	}
	if includeHelp {
		found := false
		for _, l := range lines {
			if strings.Contains(l.spec, "--help") {
				found = true
				break
			}
		}
		if !found {
			helpLine := line{spec: "-h, --help", desc: fmt.Sprintf("help for %s", c.Name())}
			if len(lines) == 0 {
				lines = append(lines, helpLine)
			} else {
				lines = append(lines[:1], append([]line{helpLine}, lines[1:]...)...)
			}
		}
	}
	maxSpec := 0
	for _, l := range lines {
		if len(l.spec) > maxSpec {
			maxSpec = len(l.spec)
		}
	}
	var b strings.Builder
	b.WriteString("----\n")
	for _, l := range lines {
		pad := maxSpec - len(l.spec)
		if pad < 0 {
			pad = 0
		}
		fmt.Fprintf(&b, "  %s%s   %s\n", l.spec, strings.Repeat(" ", pad), l.desc)
	}
	b.WriteString("----\n")
	return b.String()
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
	var subs []*cobra.Command
	for _, sc := range c.Commands() {
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
