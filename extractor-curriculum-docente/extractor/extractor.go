package extractor

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	UrlBuscador = "http://www.hermes.unal.edu.co/pages/Docentes/Docente.jsf?u="
	UrlPdf      = "http://www.hermes.unal.edu.co/birt-viewer/run?__report=/data/servidores/apache3/webapps/ROOT/reportes/portafolio/hoja-vida-docente.rptdesign&__format=pdf&BDPruebas=false&inv="
)

func getIdDocente(username string) string {

	doc := getDocumentFromUrl(getUrl(username))
	if doc == nil {
		return ""
	}

	linkElement := doc.Find("#j_id_6\\:j_id_27")
	onclickText, _ := linkElement.Attr("onclick")

	id := extractIDPersona(onclickText)

	return id
}

func GetPdfLink(username string) string {

	id := getIdDocente(username)
	if id == "" {
		return ""
	}

	return UrlPdf + id + "&foto=0"
}

func extractIDPersona(rawText string) string {
	re := regexp.MustCompile(`'idPersona\\',\\'(\d+)\\'`)
	match := re.FindStringSubmatch(rawText)

	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}
	return ""
}

func SavePdfFile(username string, link string) {

	filename := username + ".pdf"

	// Create blank file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Put content on file
	resp, err := http.Get(link)
	if err != nil {
		log.Println("Error 500: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == 500 {
		log.Println("Error 500: ", username)
		return
	}

	bits, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Println("Error writting ", username, bits, err)
	}
}
