package pug

import (
	"html/template"

	"github.com/knaka/pug/compiler"
	"github.com/knaka/pug/runtime"
)

type Options struct {
	// Setting if pretty printing is enabled.
	// Pretty printing ensures that the output html is properly indented and in human readable form.
	// If disabled, produced HTML is compact. This might be more suitable in production environments.
	// Default: false
	PrettyPrint bool

	// A Dir implements FileSystem using the native file system restricted to a specific directory tree.
	//
	// While the FileSystem.Open method takes '/'-separated paths, a Dir's string value is a filename on the native file system, not a URL, so it is separated by filepath.Separator, which isn't necessarily '/'.
	// By default, a os package is used but you can supply a different filesystem using this option
	Dir compiler.Dir

	// ExcludeImports is a list of imports, which will not be imported while compiling template.
	ExcludedImports []string

	Funcs template.FuncMap
}

func newContext(dir compiler.Dir, options ...Options) compiler.Context {
	opt := getOptions(options)

	indentString := ""

	if opt.PrettyPrint {
		indentString = "  "
	}

	if opt.Dir != nil {
		dir = opt.Dir
	}

	context := compiler.NewContext(dir, indentString, opt.ExcludedImports)

	return context
}

func compileTemplate(options Options, name string, tplstring string) (*template.Template, error) {
	return template.New(name).Funcs(options.Funcs).Funcs(runtime.FuncMap).Parse(tplstring)
}

// Parses and compiles the contents of supplied filename. Returns corresponding Go Template (html/templates) instance.
// Necessary runtime functions will be injected and the template will be ready to be executed
func CompileFile(filename string, options ...Options) (*template.Template, error) {
	ctx := newContext(compiler.FsDir("."), options...)
	if tplstring, err := ctx.CompileFile(filename); err != nil {
		return nil, err
	} else {
		return compileTemplate(getOptions(options), filename, tplstring)
	}
}

// Parses and compiles the supplied template string. Returns corresponding Go Template (html/templates) instance.
// Necessary runtime functions will be injected and the template will be ready to be executed
func CompileString(input string, options ...Options) (*template.Template, error) {
	ctx := newContext(compiler.StringInputDir(input), options...)
	if tplstring, err := ctx.CompileFile(""); err != nil {
		return nil, err
	} else {
		return compileTemplate(getOptions(options), "", tplstring)
	}
}

// Parses the contents of supplied filename template and return the Go Template source You would not be using this unless debugging / checking the output.
// Please use Compile method to obtain a template instance directly
func ParseFile(filename string, options ...Options) (string, error) {
	return newContext(compiler.FsDir("."), options...).CompileFile(filename)
}

// Parses the supplied template string and return the Go Template source You would not be using this unless debugging / checking the output.
// Please use Compile method to obtain a template instance directly
func ParseString(input string, options ...Options) (string, error) {
	return newContext(compiler.StringInputDir(input), options...).CompileFile("")
}

func getOptions(o []Options) Options {
	opt := Options{}

	if len(o) > 0 {
		opt = o[0]
	}

	return opt
}
