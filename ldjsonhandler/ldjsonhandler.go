package ldjsonhandler

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/kazarena/json-gold/ld"
	"github.com/remotejob/kaukotyoeu/domains"
)

//Create Create
func Create(articles []domains.Articlefull, index bool) []byte {

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	var doc map[string]interface{}

	if !index {
		createdstr := articles[0].Created.Format("2006-01-02")
		updatedstr := articles[0].Updated.Format("2006-01-02")

		pagelink := "http://" + articles[0].Site + "/job/fi_FI/blogi/" + articles[0].Stitle + ".html"

		publisher := map[string]interface{}{"@type": "Organization", "name": "Remote Job Finland OY", "logo": map[string]interface{}{"@type": "ImageObject", "url": "http://mazurov.eu/img/mazurovopt.jpg", "height": "200px", "width": "300px"}}
		image := map[string]interface{}{"@type": "ImageObject", "url": "http://" + articles[0].Site + "/assets/img/free_for_job.png", "height": "256px", "width": "256px"}
		mainEntityOfPage := map[string]interface{}{"@type": "WebPage", "@id": "http://" + articles[0].Site}

		var headline string

		runes := bytes.Runes([]byte(articles[0].Title))
		if len(runes) > 110 {

			headline = string(runes[:109]) + "."

		} else {

			headline = articles[0].Title + "."

		}

		doc = map[string]interface{}{
			"@context":         "http://schema.org",
			"@type":            "Article",
			"author":           articles[0].Author,
			"headline":         headline,
			"publisher":        publisher,
			"image":            image,
			"datepublished":    createdstr,
			"datemodified":     updatedstr,
			"mainEntityOfPage": mainEntityOfPage,
			// "keywords":         articlefull.Tags,
			"url": pagelink,
			//		"description":         "We love to do stuff to help people and stuff",
			"articleSection": "job",
			"articleBody":    articles[0].Contents,
		}
	} else {
		// image := map[string]interface{}{"@type": "ImageObject", "url": "http://mazurov.eu/img/mazurovopt.jpg", "height": "200px", "width": "300px"}

		// doc = map[string]interface{}{
		// 	"@context":  "http://schema.org/",
		// 	"@type":     "Person",
		// 	"name":      "Aleksander Mazurov",
		// 	"jobTitle":  "Programmer",
		// 	"telephone": "+358 451 202 801",
		// 	"url":       "http://mazurov.eu",
		// 	"image":     image,
		// }

		var itemListElement []interface{}
		for pos, article := range articles {
			listItem := map[string]interface{}{"@type": "ListItem", "position": pos, "item": map[string]interface{}{"@id": "http://" + article.Site + "/job/fi_FI/blog/" + article.Stitle + ".html", "name": article.Title}}
			itemListElement = append(itemListElement, listItem)
		}

		doc = map[string]interface{}{
			"@context":        "http://schema.org/",
			"@type":           "BreadcrumbList",
			"itemListElement": itemListElement,
		}

	}

	comp, err := proc.Compact(doc, nil, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)

	}

	b, _ := json.Marshal(comp)

	return b
}
