package handlers

import (
	"net/http"
	"strings"

	"github.com/remotejob/kaukotyoeu/sitemaps_maker"
	//	"time"
)

func CheckServeSitemap(w http.ResponseWriter, r *http.Request) {

	sitefull := r.Host
	site := strings.Split(sitefull, ":")[0]

	if site == "localhost" {

		site ="www.kaukotyo.eu"

	}

	sitemap := sitemaps_maker.Create(site)

	
	//	site :=r.URL.String()
	// 	filestr := "sitemaps/sitemap_" + site + ".xml"
	// //	fmt.Println("site", filestr)
	// 	if _, err := os.Stat(filestr); os.IsNotExist(err) {

	// 		http.NotFound(w, r)

	// 	} else {
	// 		f, _ := os.Open(filestr)
	// 		buf := bytes.NewBuffer(nil)
	// 		io.Copy(buf,f)

	// 		w.Header().Add("Content-type", "application/xml")
	// 		w.Write(buf.Bytes())

	// 	}

	w.Header().Add("Content-type", "application/xml")
	w.Write(sitemap)

}
