package mustachetpl_test

import (
	"fmt"
	"testing"

	"github.com/cbroglie/mustache"
	"github.com/google/go-cmp/cmp"
)

func TestVars(t *testing.T) {
	tpl := `
* {{name}}
* {{age}}
* {{company}}
* {{{company}}}
`
	m := map[string]any{
		"name":    "Chris",
		"company": "<b>GitHub</b>",
	}

	out, err := mustache.Render(tpl, m)
	if err != nil {
		t.Fatal(err)
	}

	expect := `
* Chris
* 
* &lt;b&gt;GitHub&lt;/b&gt;
* <b>GitHub</b>
`

	if diff := cmp.Diff(expect, out); diff != "" {
		t.Errorf("unexpected output (-want +got):\n%s", diff)
	}
}

func TestDotted(t *testing.T) {
	tpl := `
* {{client.name}}
* {{age}}
* {{client.company.name}}
* {{{company.name}}}
`

	m := map[string]any{
		"client": map[string]any{
			"name": "Chris & Friends",
			"age":  50,
		},
		"company": map[string]any{
			"name": "<b>GitHub</b>",
		},
	}

	expect := `
* Chris &amp; Friends
* 
* 
* <b>GitHub</b>
`

	out, err := mustache.Render(tpl, m)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(expect, out); diff != "" {
		t.Errorf("unexpected output (-want +got):\n%s", diff)
	}
}

func TestLambdas(t *testing.T) {
	tpl := `
* {{time.hour}}
`

	m := map[string]any{
		"year":  1970,
		"month": 1,
		"day":   1,
		"time": func(text string, render mustache.RenderFunc) (string, error) {
			fmt.Printf("TEXT1 %q\n", text)
			return "wow", nil
		},
		"today": func(text string, render mustache.RenderFunc) (string, error) {
			fmt.Printf("TEXT2 %q\n", text)
			return "{{year}}-{{month}}-{{day}}", nil
		},
	}

	expect := `
* 0
`

	out, err := mustache.Render(tpl, m)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(expect, out); diff != "" {
		t.Errorf("unexpected output (-want +got):\n%s", diff)
	}
}

func TestLambda2(t *testing.T) {
	tmpl := `{{#lambda}}Hello {{name}}. {{#sub}}{{.}} {{/sub}}{{^negsub}}nothing{{/negsub}}{{/lambda}}`
	data := map[string]interface{}{
		"name": "world",
		"sub":  []string{"subv1", "subv2"},
		"lambda": func(text string, render mustache.RenderFunc) (string, error) {
			res, err := render(text)
			return res + "!", err
		},
	}

	output, err := mustache.Render(tmpl, data)
	if err != nil {
		t.Fatal(err)
	}
	expect := "Hello world. subv1 subv2 nothing!"
	if output != expect {
		t.Fatalf("TestLambda expected %q got %q", expect, output)
	}
}
