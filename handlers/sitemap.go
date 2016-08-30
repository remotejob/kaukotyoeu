package handlers

import (
	// "bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/remotejob/kaukotyoeu/sitemaps_maker"
)

func CheckServeSitemap(w http.ResponseWriter, r *http.Request) {

	sitefull := r.Host
	site := strings.Split(sitefull, ":")[0]

	if site == "localhost" {

		site = "www.kaukotyo.eu"

	}

	sitemap := sitemaps_maker.Create(site)

	fmt.Println("CheckServeSitemap")

	// fmt.Println(bytes.NewBuffer(sitemap).String())

	w.Header().Add("Content-type", "application/xml")
	w.Write(sitemap)

}
