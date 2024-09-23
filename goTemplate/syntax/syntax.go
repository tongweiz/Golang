package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

type Data struct {
	Name string
}

func main() {

	tmpl, _ := template.New("comment").Parse(`
Before comment
{{/* This is a comment */}}
After comment
{{- /* Comment with white space trimming */ -}}
End
`)
	tmpl.Execute(os.Stdout, nil)
	/*	Before comment
		After commentEnd*/

	tmpl, _ = template.New("pipeline").Parse("Hello, {{.}}!")
	tmpl.Execute(os.Stdout, "world")
	//Hello, world!

	tmpl, _ = template.New("if").Parse(`
{{if .}}Value is true!{{end}}
`)
	tmpl.Execute(os.Stdout, true)
	//It checks whether the value is truthy. In Go, truthy values include true, non-zero numbers, non-empty strings, etc

	tmpl, _ = template.New("ifelse").Parse(`
{{if .}}True!{{else}}False!{{end}}
`)
	tmpl.Execute(os.Stdout, false)
	//False!

	tmpl, _ = template.New("elseif").Parse(`
{{if eq . 1}}One!{{else if eq . 2}}Two!{{end}}
`)
	tmpl.Execute(os.Stdout, 2)
	//2

	tmpl, _ = template.New("range").Parse(`
{{range .}}Item: {{.}}
{{end}}
`)
	tmpl.Execute(os.Stdout, []string{"apple", "banana", "cherry"})
	/*Item: apple
	  Item: banana
	  Item: cherry
	*/

	tmpl, _ = template.New("rangeelse").Parse(`
{{range .}}Item: {{.}}
{{else}}No items!
{{end}}
`)
	tmpl.Execute(os.Stdout, []string{})
	//No items!

	//Break and Continue

	tmpl, _ = template.New("template name").Parse(`
{{define "child"}}
    Hello, this is the child template.
{{end}}

This is the main template.
{{template "child"}}
`)
	tmpl.Execute(os.Stdout, nil)
	/*	This is the main template.

		Hello, this is the child template.*/

	tmpl, _ = template.New("template name pipeline").Parse(`
{{define "child"}}
    Hello, {{.Count}} items are made of {{.Material}}.
{{end}}

This is the main template.
{{template "child" .}}
`)
	sweaters := Inventory{"wool", 17}
	tmpl.Execute(os.Stdout, sweaters)
	/*	This is the main template.

		Hello, 17 items are made of wool.*/

	const rootTemplate = `
This is the root template.
{{block "header" .}}Default Header{{end}}

Main content goes here for {{.Name}}.

{{block "footer" .}}Default Footer{{end}}
`

	const customTemplate = `
{{define "header"}}Custom Header for Special Case{{end}}

{{define "footer"}}Custom Footer - See You Again!{{end}}
`
	data := Data{Name: "Tongwei"}

	// Parse the root template
	tmpl = template.Must(template.New("root").Parse(rootTemplate))

	// Redefine blocks by parsing customTemplate on top of rootTemplate
	customTmpl := template.Must(tmpl.Parse(customTemplate))

	// Execute the customized template
	customTmpl.Execute(os.Stdout, data)
	/*	This is the root template.
			Custom Header for Special Case

		Main content goes here for Tongwei.

			Custom Footer - See You Again!*/

	const tmpl1 = `{{with .Name}}Hello, {{.}}!{{end}}`
	data1 := struct{ Name string }{Name: "Tongwei"}
	data2 := struct{ Name string }{Name: ""}
	t := template.Must(template.New("example1").Parse(tmpl1))
	// Output when Name is set
	t.Execute(os.Stdout, data1) // Output: Hello, Tongwei!
	// Output when Name is empty
	t.Execute(os.Stdout, data2) // Output: (no output)

	const tmpl2 = `{{with .Name}}Hello, {{.}}!{{else}}Name is not provided.{{end}}`
	data1 = struct{ Name string }{Name: "Tongwei"}
	data2 = struct{ Name string }{Name: ""}
	t = template.Must(template.New("example2").Parse(tmpl2))
	// Output when Name is set
	t.Execute(os.Stdout, data1) // Output: Hello, Tongwei!
	// Output when Name is empty
	t.Execute(os.Stdout, data2) // Output: Name is not provided.

	const tmpl3 = `{{with .Name}}Hello, {{.}}!{{else with .FallbackName}}Using fallback: {{.}}.{{end}}`
	data3 := struct {
		Name         string
		FallbackName string
	}{Name: "", FallbackName: "Fallback User"}
	data4 := struct {
		Name         string
		FallbackName string
	}{Name: "Tongwei", FallbackName: "Fallback User"}
	t = template.Must(template.New("example3").Parse(tmpl3))
	// Output when Name is empty, FallbackName is used
	t.Execute(os.Stdout, data3) // Output: Using fallback: Fallback User.
	// Output when Name is set
	t.Execute(os.Stdout, data4) // Output: Hello, Tongwei!
}
