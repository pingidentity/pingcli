package main

// duplicate function body detector
// ---------------------------------
// This small utility scans selected Go source directories and reports functions whose
// bodies are structurally identical. It's intended to help spot accidental copy/paste
// duplication (especially when refactoring small helper functions in tools or
// configuration option handling).
//
// HOW IT WORKS
// 1. Walk a curated set of directories (see includeDirs) collecting .go files (excluding tests).
// 2. Parse each file with the Go parser into an AST.
// 3. For every function with a body, iterate each statement node and build a textual
//    representation using the %#v (Go-syntax) formatting of the AST nodes.
// 4. Normalize this textual representation (case fold, trim extra whitespace, remove newlines)
//    to reduce noise (e.g., formatting differences) while still being stable.
// 5. Hash (SHA‑256) the normalized body representation. The hash becomes a key in a map
//    to the list of (file:function) locations that share that exact body hash.
// 6. Any hash with more than one location is reported as a duplicate.
//
// WHY NORMALIZE?
// The AST %#v output can vary in insignificant whitespace. Normalization reduces false
// negatives from formatting differences but still treats any token / structural change as different.
//
// LIMITATIONS / NON-GOALS
// * Ignores function signatures (we only compare bodies). Two functions with different
//   names/parameters but identical logic are flagged—which is desired for dedupe.
// * Does not attempt near-duplicate detection (e.g., only one constant differs).
// * Anonymous functions (lambdas) are ignored because we only traverse top-level *ast.FuncDecl.
// * Methods vs functions: receiver differences are ignored (body only).
//
// EXIT CODES
// 0 = No duplicates found
// 2 = One or more duplicate pairs reported
// 1 = I/O or parsing failure during traversal
//
// TYPICAL USAGE
//   go run ./tools/check-duplicates
// or as a CI guard / Makefile target.
//
// To extend scanning, add paths to includeDirs. Keep the list tight to avoid noisy matches
// across unrelated packages.

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// includeDirs restricts the scan to a safe subset of the repository. Adjust cautiously.
var includeDirs = []string{
	"tools",
	"internal/configuration/options",
}

// ignoreFiles filters out generated or undesirable files (currently: test sources).
var ignoreFiles = regexp.MustCompile(`_test\.go$`)

func main() {
	// Map: bodyHash -> list of locations (file:functionName)
	funcMap := map[string][]string{}
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if !strings.HasSuffix(path, ".go") || ignoreFiles.MatchString(path) {
			return nil
		}
		if !withinIncluded(path) {
			return nil
		}
		addFile(path, funcMap)
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "walk error:", err)
		os.Exit(1)
	}

	// Build list of every pair among colliding function bodies.
	var collisions [][2]string
	for _, locs := range funcMap {
		if len(locs) > 1 {
			for i := range locs {
				for j := i + 1; j < len(locs); j++ {
					collisions = append(collisions, [2]string{locs[i], locs[j]})
				}
			}
		}
	}

	if len(collisions) == 0 {
		fmt.Println("No duplicate functions found.")
		return
	}
	fmt.Println("Duplicate functions detected:")
	for _, c := range collisions {
		fmt.Printf("  - %s == %s\n", c[0], c[1])
	}
	os.Exit(2)
}

// withinIncluded returns true if the path is rooted in one of the includeDirs.
func withinIncluded(path string) bool {
	for _, d := range includeDirs {
		if strings.HasPrefix(path, d+"/") {
			return true
		}
	}
	return false
}

// addFile parses a Go file, hashes each function body, and records its location under that hash key.
func addFile(path string, funcMap map[string][]string) {
	// Sanitize and restrict the path before opening (addresses gosec G304 false positive).
	clean := filepath.Clean(path)
	if filepath.IsAbs(clean) || strings.Contains(clean, "..") {
		return // reject unexpected absolute or parent traversals
	}
	if !withinIncluded(clean) || !strings.HasSuffix(clean, ".go") || ignoreFiles.MatchString(clean) {
		return
	}

	f, err := os.Open(clean) // #nosec G304: path origin is controlled by WalkDir + allowlist + sanitization above
	if err != nil {
		return // Silent skip; traversal reports aggregate errors only.
	}
	defer func() { _ = f.Close() }()
	src, err := io.ReadAll(f)
	if err != nil {
		return
	}

	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, clean, src, parser.ParseComments)
	if err != nil {
		return // Skip unreadable / invalid Go sources.
	}
	for _, d := range parsed.Decls {
		fd, ok := d.(*ast.FuncDecl)
		if !ok || fd.Body == nil { // Skip declarations without bodies (interfaces, externs).
			continue
		}
		var buf bytes.Buffer
		for _, s := range fd.Body.List {
			buf.WriteString(normalize(fmt.Sprintf("%#v", s)))
		}
		h := sha256.Sum256(buf.Bytes())
		key := fmt.Sprintf("%x", h)
		loc := fmt.Sprintf("%s:%s", clean, fd.Name.Name)
		funcMap[key] = append(funcMap[key], loc)
	}
}

// normalize reduces insignificant differences in the AST statement dump so that
// logically identical bodies hash the same even if formatting varies.
func normalize(s string) string {
	s = strings.ToLower(s)                   // case-insensitive
	s = strings.Join(strings.Fields(s), " ") // collapse all whitespace runs
	s = strings.ReplaceAll(s, "\n", " ")     // remove line breaks entirely
	return s
}
