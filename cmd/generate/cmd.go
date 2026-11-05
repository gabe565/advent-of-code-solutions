package main

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

const module = "github.com/gabe565/advent-of-code-solutions"

//go:embed templates
var templates embed.FS

func main() {
	if err := newCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

func newCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "generate year day",
		Short: "Generate a new day",
		RunE:  run,
		Args:  cobra.ExactArgs(2),
	}
	return &cmd
}

type data struct {
	Year int
	Day  int
}

func run(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	year, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid year: %w", err)
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid day: %w", err)
	}

	d := data{Year: year, Day: day}
	yearPkg := "year" + strconv.Itoa(year)
	dayPkg := "day" + strconv.Itoa(day)
	yearDir := filepath.Join("cmd", yearPkg)
	dayDir := filepath.Join(yearDir, dayPkg)

	if err := renderDir(cmd, "templates/day", dayDir, d); err != nil {
		return err
	}

	yearCmd := filepath.Join(yearDir, "cmd.go")
	switch _, err := os.Stat(yearCmd); {
	case errors.Is(err, os.ErrNotExist):
		if err := renderDir(cmd, "templates/year", yearDir, d); err != nil {
			return err
		}

		rootCmd := filepath.Join("cmd", "cmd.go")
		if err := addCommand(rootCmd, module+"/cmd/"+yearPkg, yearPkg+".New()"); err != nil {
			return err
		}
		cmd.Println("Updated", rootCmd)
	case err != nil:
		return err
	default:
		if err := addCommand(
			yearCmd,
			module+"/cmd/"+yearPkg+"/"+dayPkg,
			dayPkg+".New().Cmd()",
		); err != nil {
			return err
		}
		cmd.Println("Updated", yearCmd)
	}

	if err := writeInputs(cmd, year, day); err != nil {
		return err
	}

	return nil
}

func renderDir(cmd *cobra.Command, dir, dst string, d data) error {
	tmpl, err := template.ParseFS(templates, dir+"/*.tmpl")
	if err != nil {
		return err
	}

	paths, err := fs.Glob(templates, dir+"/*.tmpl")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, 0o755); err != nil {
		return err
	}

	for _, path := range paths {
		name := filepath.Base(path)

		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, name, d); err != nil {
			return err
		}

		src, err := format.Source(buf.Bytes())
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}

		path := filepath.Join(dst, strings.TrimSuffix(name, ".tmpl"))

		if err := os.WriteFile(path, src, 0o644); err != nil {
			return err
		}

		cmd.Println("Wrote", path)
	}

	return nil
}

func addCommand(path, importPath, expr string) error {
	src, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, src, parser.ImportsOnly)
	if err != nil {
		return err
	}

	var importDecl *ast.GenDecl
	for _, decl := range f.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.IMPORT && gen.Rparen.IsValid() {
			importDecl = gen
			break
		}
	}
	if importDecl == nil {
		return fmt.Errorf("%s: no import block found", path)
	}

	for _, spec := range importDecl.Specs {
		if p, _ := strconv.Unquote(spec.(*ast.ImportSpec).Path.Value); p == importPath { //nolint:errcheck
			return fmt.Errorf("%s: %s is already imported", path, importPath)
		}
	}

	f, err = parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return err
	}

	var call *ast.CallExpr
	ast.Inspect(f, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if sel, ok := c.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "AddCommand" {
				call = c
				return false
			}
		}
		return call == nil
	})
	if call == nil {
		return fmt.Errorf("%s: no AddCommand call found", path)
	}

	inserts := []struct {
		pos  token.Pos
		text string
	}{
		{importDecl.Rparen, "\t" + strconv.Quote(importPath) + "\n"},
		{call.Rparen, "\t" + expr + ",\n"},
	}
	slices.SortFunc(inserts, func(a, b struct {
		pos  token.Pos
		text string
	}) int {
		return int(b.pos - a.pos)
	})

	for _, ins := range inserts {
		offset := fset.Position(ins.pos).Offset
		lineStart := bytes.LastIndexByte(src[:offset], '\n') + 1
		src = slices.Insert(src, lineStart, []byte(ins.text)...)
	}

	src, err = format.Source(src)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	return os.WriteFile(path, src, 0o644) //nolint:gosec
}

func writeInputs(cmd *cobra.Command, year, day int) error {
	inputDir := filepath.Join("inputs", strconv.Itoa(year), strconv.Itoa(day))

	if err := os.MkdirAll(inputDir, 0o755); err != nil {
		return err
	}

	examplePath := filepath.Join(inputDir, "example.txt")
	inputPath := filepath.Join(inputDir, "input.txt")

	if err := errors.Join(
		os.WriteFile(examplePath, nil, 0o644),
		os.WriteFile(inputPath, nil, 0o644),
	); err != nil {
		return err
	}

	cmd.Println("Wrote " + examplePath)
	cmd.Println("Wrote " + inputPath)

	return nil
}
