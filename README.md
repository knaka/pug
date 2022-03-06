

# pug
This package was originally written by eknkc. This package contains some fixes and modifications.

`import "github.com/knaka/pug"`

* [Overview](#pkg-overview)
* [Difference with original](#pkg-difference)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package pug.go is an elegant templating engine for Go Programming Language.
It is a port of Pug template engine, previously known as Jade.

Pug.go compiles .pug templates to standard go templates (<a href="https://golang.org/pkg/html/template/">https://golang.org/pkg/html/template/</a>) and returns a `*template.Template` instance.

While there is no JavaScript environment present, Pug.go provides basic expression support over go template syntax. Such as `a(href="/user/" + UserId)` would concatenate two strings. You can use arithmetic, logical and comparison operators as well as ternery if operator.

Please check *Pug Language Reference* for details: <a href="https://pugjs.org/api/getting-started.html">https://pugjs.org/api/getting-started.html</a>.

Differences between Pug and Pug.go (items with checkboxes are planned, just not present yet)

- [ ] Multiline attributes are not supported
- [ ] `&attributes` syntax is not supported
- [ ] `case` statement is not supported
- [ ] Filters are not supported
- [ ] Mixin rest arguments are not supported.
- Mixin blocks are not supported. Go templates do not allow variable template includes so this is tricky.
- `while` loops are not supported as Go templates do not provide it. We could use recursive templates or channel range loops etc but that would be unnecessary complexity.
- Unbuffered code blocks are not possible as we don't have a JS environment. However it is possible to define variables using `- var x = "foo"` syntax as an exception.

Apart from these missing features, everything in the language reference should be supported.


## <a name="pkg-difference">Difference with original</a>

- Fixed support for class names with '-'
- Fixed nested loops
- Added support for single quotes
- Added support for import exception list

## <a name="pkg-index">Index</a>
* [func CompileFile(filename string, options ...Options) (*template.Template, error)](#CompileFile)
* [func CompileString(input string, options ...Options) (*template.Template, error)](#CompileString)
* [func ParseFile(filename string, options ...Options) (string, error)](#ParseFile)
* [func ParseString(input string, options ...Options) (string, error)](#ParseString)
* [type Options](#Options)


#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/eknkc/pug/doc.go) [pug.go](/src/github.com/eknkc/pug/pug.go) 





## <a name="CompileFile">func</a> [CompileFile](/src/target/pug.go?s=1524:1605#L52)
``` go
func CompileFile(filename string, options ...Options) (*template.Template, error)
```
Parses and compiles the contents of supplied filename. Returns corresponding Go Template (html/templates) instance.
Necessary runtime functions will be injected and the template will be ready to be executed



## <a name="CompileString">func</a> [CompileString](/src/target/pug.go?s=2010:2090#L63)
``` go
func CompileString(input string, options ...Options) (*template.Template, error)
```
Parses and compiles the supplied template string. Returns corresponding Go Template (html/templates) instance.
Necessary runtime functions will be injected and the template will be ready to be executed



## <a name="ParseFile">func</a> [ParseFile](/src/target/pug.go?s=2509:2576#L74)
``` go
func ParseFile(filename string, options ...Options) (string, error)
```
Parses the contents of supplied filename template and return the Go Template source You would not be using this unless debugging / checking the output.
Please use Compile method to obtain a template instance directly



## <a name="ParseString">func</a> [ParseString](/src/target/pug.go?s=2865:2931#L80)
``` go
func ParseString(input string, options ...Options) (string, error)
```
Parses the supplied template string and return the Go Template source You would not be using this unless debugging / checking the output.
Please use Compile method to obtain a template instance directly




## <a name="Options">type</a> [Options](/src/target/pug.go?s=108:838#L10)
``` go
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
}
```













- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
