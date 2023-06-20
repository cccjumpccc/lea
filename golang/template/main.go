package main

import (
	"html/template"
	"os"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	// queryBuf := &bytes.Buffer{}
	template.Must(template.New("").Parse("{{.query}} aaa {{.stime}} bbb {{.etime}}")).
		Execute(os.Stdout, map[string]string{
			// "query": strconv.Quote("qqq"),
			// "stime": strconv.Quote("sss"),
			// "etime": strconv.Quote("eee"),
			"query": "qqq",
			"stime": "sss",
			"etime": "eee",
		})
}
