package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func main() {
	templateFile := os.Args[1]
	src := os.Args[2]
	output := os.Args[3]

	b, err := os.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}
	b2, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}

	templated, err := TemplatedString(fmt.Sprintf("%s\n%s", string(b), string(b2)), nil)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(output, []byte(templated), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TemplatedString(t string, i interface{}) (string, error) {
	b := bytes.NewBuffer([]byte{})
	tmpl, err := template.New("template").Funcs(sprig.TxtFuncMap()).Parse(t)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(b, i)

	return b.String(), err
}
