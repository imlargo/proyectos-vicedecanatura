package main

import (
	"extractor-curriculum-docente/extractor"
)

func main() {

	errs := 0

	var noEncontrado []string = make([]string, 0)
	println(len(docentes))

	for i := 0; i < len(docentes); i++ {

		username := docentes[i]

		println(i, ". ", username)
		linkPdf := extractor.GetPdfLink(username)
		if linkPdf == "" {
			errs += 1
			noEncontrado = append(noEncontrado, username)
			println("--- Error: ", username)
			continue
		}
		filename := "datos/" + username
		go extractor.SavePdfFile(filename, linkPdf)
	}

	println("Errors: ", errs)

	for _, d := range noEncontrado {
		println(d)
	}
}
