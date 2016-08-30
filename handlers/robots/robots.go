package robots

import (
	"bytes"
	"net/http"
	"strings"
)

func Generate(w http.ResponseWriter, r *http.Request) {

	var buffer bytes.Buffer

	sitefull := r.Host
	site := strings.Split(sitefull, ":")[0]

	buffer.WriteString("User-agent: *\nAllow: /\nSitemap: http://" + site + "/sitemap.xml\n")

	w.Header().Add("Content-type", "text/plain")
	w.Write(buffer.Bytes())

}
