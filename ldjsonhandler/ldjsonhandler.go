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

	var doc map[string]interface{}

	if len(articles) > 0 {
		createdstr := articles[0].Created.Format("2006-01-02")
		updatedstr := articles[0].Updated.Format("2006-01-02")

		pagelink := "http://" + articles[0].Site + "/job/fi_FI/blogi/" + articles[0].Stitle + ".html"

		publisher := map[string]interface{}{"@type": "Organization", "name": "Remote Job Finland OY", "logo": map[string]interface{}{"@type": "ImageObject", "url": "http://mazurov.eu/img/free_for_job.png", "height": "256px", "width": "256px"}}
		image := map[string]interface{}{"@type": "ImageObject", "url": "http://" + articles[0].Site + "/assets/img/free_for_job.png", "height": "256px", "width": "256px"}
		mainEntityOfPage := map[string]interface{}{"@type": "WebPage", "@id": "http://" + articles[0].Site}

		doc = map[string]interface{}{
			"@context":         "http://schema.org",
			"@type":            "Article",
			"author":           articles[0].Author,
			"headline":         articles[0].Title,
			"publisher":        publisher,
			"image":            image,
			"datepublished":    createdstr,
			"datemodified":     updatedstr,
			"mainEntityOfPage": mainEntityOfPage,
			// "keywords":         articlefull.Tags,
			"url": pagelink,
			//		"description":         "We love to do stuff to help people and stuff",
			"articleSection": "realestate",
			"articleBody":    articles[0].Contents,
		}
	} else {
		image := map[string]interface{}{"@type": "ImageObject", "url": "http://mazurov.eu/img/mazurovopt.jpg", "height": "200px", "width": "300px"}

		doc = map[string]interface{}{
			"@context":  "http://schema.org/",
			"@type":     "Person",
			"name":      "Aleksander Mazurov",
			"jobTitle":  "Programmer",
			"telephone": "+358 451 202 801",
			"url":       "http://mazurov.eu",
			"image":     image,
		}
	}
	comp, err := proc.Compact(doc, nil, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)

	}

	b, _ := json.Marshal(comp)

	return b
}
