package ldjsonhandler

import (
	"encoding/json"
	"log"

	"github.com/kazarena/json-gold/ld"
	"github.com/remotejob/kaukotyoeu/domains"
)

//Create Create
func Create(articles []domains.Articlefull) []byte {

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	image := map[string]interface{}{"@type": "ImageObject", "url": "http://mazurov.eu/img/mazurovopt.jpg", "height": "200px", "width": "300px"}

	doc := map[string]interface{}{
		"@context":  "http://schema.org/",
		"@type":     "Person",
		"name":      "Aleksander Mazurov",
		"jobTitle":  "Programmer",
		"telephone": "+358 451 202 801",
		"url":       "http://mazurov.eu",
		"image":     image,
	}

	comp, err := proc.Compact(doc, nil, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)

	}

	b, _ := json.Marshal(comp)

	return b
}
