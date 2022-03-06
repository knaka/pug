package pug

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test_Doctype(t *testing.T) {
	res, err := run(`doctype html`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<!DOCTYPE html>`, t)
	}
}

func Test_Nesting(t *testing.T) {
	res, err := run(`html
						head
							title
						body`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<html><head><title></title></head><body></body></html>`, t)
	}
}

func Test_Id(t *testing.T) {
	res, err := run(`div#test`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div id="test"></div>`, t)
	}
}

func Test_Class(t *testing.T) {
	res, err := run(`div.test`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div class="test"></div>`, t)
	}
}

func Test_Arg_Quotes(t *testing.T) {
	res, err := run(`div(x-arg='test')`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div x-arg="test"></div>`, t)
	}
}

func Test_Arg_Quotes_With_Inner_Quote(t *testing.T) {
	res, err := run(`div(x-arg='test=\''+'\';')`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div x-arg="test='';"></div>`, t)
	}
}

func Test_MultiClass(t *testing.T) {
	res, err := run(`
div.test.foo.bar(class="baz")
	p.foo(class=["bar", "baz"]): a(href="#") foo
		| bar
`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div class="test foo bar baz"><p class="foo bar baz"><a href="#">foobar</a></p></div>`, t)
	}
}

func Test_ObjectClass(t *testing.T) {
	res, err := run(`
div.test.foo(class={bar: true, "baz": 5<4, buzz: 4<5})
`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div class="test foo bar buzz"></div>`, t)
	}
}

func Test_Attribute(t *testing.T) {
	res, err := run(`
div(name="Test" @foo.bar="baz", commasep=1 unescaped!="<foo>").testclass
	p(style="text-align: center; color: maroon" "quoted"= "foo")
		span.class-name#id-name
`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div name="Test" @foo.bar="baz" commasep="1" unescaped="<foo>" class="testclass"><p style="text-align: center; color: maroon" quoted="foo"><span class="class-name" id="id-name"></span></p></div>`, t)
	}
}

func Test_EmptyAttribute(t *testing.T) {
	res, err := run(`div(name)`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div name></div>`, t)
	}
}

func Test_MapAttribute(t *testing.T) {
	res, err := run(`div(attr={foo: "bar"})`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div attr="{&#34;foo&#34;:&#34;bar&#34;}"></div>`, t)
	}
}

func Test_SafeAttribute(t *testing.T) {
	res, err := run(`
- var Color2 = Color
div(style='background: '+Color2)
`, struct {
		Color string
	}{Color: "rgb(0, 0, 0)"})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div style="background: rgb(0, 0, 0)"></div>`, t)
	}
}

func Test_SelfClose(t *testing.T) {
	res, err := run(`div(name="foo")/`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div name="foo" />`, t)
	}
}

func Test_Empty(t *testing.T) {
	res, err := run(``, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, ``, t)
	}
}

func Test_ArithmeticExpression(t *testing.T) {
	res, err := run(`| #{A + B * C}`, map[string]int{"A": 2, "B": 3, "C": 4})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `14`, t)
	}
}

func Test_BooleanExpression(t *testing.T) {
	res, err := run(`| #{C - A < B}`, map[string]int{"A": 2, "B": 3, "C": 4})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `true`, t)
	}
}

func Test_Interpolation(t *testing.T) {
	res, err := run(`| #{Key} !{Key}`, testStruct{Key: "<hr />"})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `&lt;hr /&gt; <hr />`, t)
	}
}

func Test_Buffered(t *testing.T) {
	res, err := run(`
p
 = Key
 != Key
`, testStruct{Key: "<hr />"})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p>&lt;hr /&gt;<hr /></p>`, t)
	}
}

func Test_TerneryExpression(t *testing.T) {
	res, err := run(`| #{ B > A ? A > B ? "x" : "y" : "z" }`, map[string]int{"A": 2, "B": 3})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `y`, t)
	}
}

func Test_ArrayExpression(t *testing.T) {
	res, err := run(`| #{ [1,2,3] }`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `[1 2 3]`, t)
	}
}

func Test_If(t *testing.T) {
	res, err := run(`
if Key == "foo"
	| foo
else
 | bar
`, testStruct{Key: "foo"})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `foo`, t)
	}
}

func Test_Unless(t *testing.T) {
	res, err := run(`
unless Key != "bar"
	| foo
`, testStruct{Key: "foo"})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `foo`, t)
	}
}

func Test_TerneryClass(t *testing.T) {
	res, err := run(`
each item, i in Items
	p(class=i % 2 == 0 ? "even" : "odd") #{item}`, testStruct{Items: []string{"test1", "test2"}})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p class="even">test1</p><p class="odd">test2</p>`, t)
	}
}

func Test_Style(t *testing.T) {
	res, err := run(`p(style="color: red"): span(style={color: "green", "font-size": 20})`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p style="color: red"><span style="color:green;font-size:20"></span></p>`, t)
	}
}

func Test_Index(t *testing.T) {
	res, err := run(`p.index= Items[1]`, testStruct{Items: []string{"test1", "test2"}})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p class="index">test2</p>`, t)
	}
}

func Test_NilClass(t *testing.T) {
	res, err := run(`p(class=nil)`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p class=""></p>`, t)
	}
}

func Test_MapAccess(t *testing.T) {
	res, err := run(`p #{a.b().c}`, map[string]interface{}{
		"a": map[string]interface{}{
			"b": func() interface{} {
				return map[string]interface{}{
					"c": "d",
				}
			},
		},
	})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p>d</p>`, t)
	}
}

func Test_Dollar_In_TagAttributes(t *testing.T) {
	res, err := run(`input(placeholder="$ per "+kwh)`, map[string]interface{}{
		"kwh": "kWh",
	})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<input placeholder="$ per kWh" />`, t)
	}
}

func Test_MixinBasic(t *testing.T) {
	res, err := run(`
mixin test()
	p #{Key}

+test()
`, testStruct{Key: "value"})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p>value</p>`, t)
	}
}

func Test_MixinWithArgs(t *testing.T) {
	res, err := run(`
mixin test(arg, arg2)
	p #{Key} #{arg} #{arg2}

+test(15, 1+1)
`, testStruct{Key: "value"})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p>value 15 2</p>`, t)
	}
}

func Test_Each(t *testing.T) {
	res, err := run(`
each v in Items
		p #{v}
		`, testStruct{
		Items: []string{"t1", "t2"},
	})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p>t1</p><p>t2</p>`, t)
	}
}

func Test_EachInEach(t *testing.T) {
	res, err := run(`
each v1 in Arr1
	each v2 in Arr2
		p #{v1}#{v2}
		`, struct {
		Arr1 []int
		Arr2 []int
	}{
		Arr1: []int{1, 2},
		Arr2: []int{3, 4},
	})

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p>13</p><p>14</p><p>23</p><p>24</p>`, t)
	}
}

func Test_Assignment(t *testing.T) {
	res, err := run(`
- var vrb = "test"
- var vrb = "test2"
p #{vrb}
`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p>test2</p>`, t)
	}
}

func Test_Reassignment(t *testing.T) {
	res, err := run(`
- var vRb = "test"
p #{vRb}
if 1==1
	- var vRb = "test2"
p #{vRb}
`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p>test</p><p>test2</p>`, t)
	}
}

func Test_Block(t *testing.T) {
	res, err := run(`
block deneme
		p Test
		`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p>Test</p>`, t)
	}
}

func Test_RawText(t *testing.T) {
	res, err := run(`
style.
  body{ color: red }
p a
`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, "<style>  body{ color: red }\n</style><p>a</p>", t)
	}
}

func Test_Import(t *testing.T) {
	tpl, err := CompileFile("examples/import/import.pug")

	if err != nil {
		t.Fatal(err)
	}

	buf := &bytes.Buffer{}

	if err := tpl.Execute(buf, nil); err != nil {
		t.Fatal(err)
	}

	expect(string(buf.Bytes()), "<style>body { color: red; }\n</style><p>Main<p>import1</p><p>import2</p></p><p>import2</p>", t)
}

func Test_ImportException(t *testing.T) {
	tpl, err := CompileFile("examples/import/import.pug", Options{
		ExcludedImports: []string{"sub/style.css"},
	})

	if err != nil {
		t.Fatal(err)
	}

	buf := &bytes.Buffer{}

	tpl.New("sub/style.css").Parse("body { color: green; }")

	if err := tpl.Execute(buf, nil); err != nil {
		t.Fatal(err)
	}

	expect(string(buf.Bytes()), "<style>body { color: green; }</style><p>Main<p>import1</p><p>import2</p></p><p>import2</p>", t)
}

func Test_Extend(t *testing.T) {
	tpl, err := CompileFile("examples/extend/extend.pug")

	if err != nil {
		t.Fatal(err)
	}

	buf := &bytes.Buffer{}

	if err := tpl.Execute(buf, nil); err != nil {
		t.Fatal(err)
	}

	expect(string(buf.Bytes()), "<body><p>extend-test1</p><p>mid-test2</p><p>base-test3</p><p>extend-test3-append</p></body>", t)
}

func Test_Issue_4(t *testing.T) {
	res, err := run(`
if true && true
	p #{true && true}
		`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<p>true</p>`, t)
	}
}

func Test_Issue_3(t *testing.T) {
	res, err := run(`
.cl1
		.cl2
			 .cl3 test
		`, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		expect(res, `<div class="cl1"><div class="cl2"><div class="cl3">test</div></div></div>`, t)
	}
}

func Benchmark_Parse(b *testing.B) {
	code := `
	!!! 5
	html
		head
			title Test Title
		body
			nav#mainNav[data-foo="bar"]
			div#content
				div.left
				div.center
					block center
						p Main Content
							.long ? somevar && someothervar
				div.right`

	for i := 0; i < b.N; i++ {
		CompileString(code)
	}
}

func expect(cur, expected string, t *testing.T) {
	if cur != expected {
		t.Fatalf("Expected {%s} got {%s}.", expected, cur)
	}
}

func run(tpl string, data interface{}) (string, error) {
	// fmt.Println(ParseString(tpl, Options{PrettyPrint: true}))

	t, err := CompileString(tpl)
	if err != nil {
		return "", fmt.Errorf("could not compile template: %v", err)
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("could not execute template: %v", err)
	}
	return strings.TrimSpace(buf.String()), nil
}

type testStruct struct {
	Key   string
	Items []string
}

func Test_TagInterpolation(t *testing.T) {
	tests := []struct {
		tpl  string
		data interface{}
		want string
	}{
		{
			"p Hello #[foo] World",
			nil,
			"<p>Hello <foo></foo> World</p>",
		},
		{
			"p Hello #[foo(bar='123')] World",
			nil,
			"<p>Hello <foo bar=\"123\"></foo> World</p>",
		},
		{
			"p Hello #[foo(bar='123') baz] World",
			nil,
			"<p>Hello <foo bar=\"123\">baz</foo> World</p>",
		},
	}
	for _, tt := range tests {
		if res, _ := run(tt.tpl, tt.data); res != tt.want {
			t.Errorf("res %v, want %v", res, tt.want)
		}
	}
}
