package main

import (
	"flag"
	"path/filepath"

	"github.com/tinyjson/tinyjson/generator"
	//frontend "gitlab.com/sud-ua/parser/models/frontend"
	//"gitlab.com/sud-ua/parser/models/frontend/structs"
)

//go:generate tinyjson ./example_generates

func main() {
	flag.Parse()
	extra := flag.Args()

	if len(extra) == 0 {
		return
	}

	for _, pkg := range extra {
		pkg, _ = filepath.Abs(pkg)
		if err := generator.ParsePackage(pkg); err != nil {
			panic(err)
		}
	}
	generator.WriteTypes()

	return
}
