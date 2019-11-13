package main

import "awesomeProject/latex"

func main() {
	m := make(map[string]string)
	m["Test Section Title"] = "Section Content"
	latex.ParseFileWithTemplate(m)
	latex.CompilePdf("TestTemplate")
}
