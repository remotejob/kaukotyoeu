//go:generate  /home/juno/neonworkspace/gowork/bin/statik -src=./public

package main // import "github.com/remotejob/kaukotyoeu"

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// _ "github.com/remotejob/godocker/statik"

	"github.com/remotejob/kaukotyoeu/handlers"
	"github.com/remotejob/kaukotyoeu/handlers/robots"
)

var themes string
var locale string

var addrs []string
var database string
var username string
var password string
var mechanism string

var addrsext []string
var databaseext string
var usernameext string
var passwordext string
var mechanismext string

var sites []string


func main() {

	// finalHandler := http.HandlerFunc(final)

	fs := http.FileServer(http.Dir("assets"))

	r := mux.NewRouter()
	r.HandleFunc("/robots.txt", robots.Generate)
	r.HandleFunc("/sitemap.xml", handlers.CheckServeSitemap)
	r.HandleFunc("/job/{locale}/{themes}/{mtitle}.html", handlers.CreateArticelePage)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	r.HandleFunc("/", handlers.CreateIndexPage)
	// loggedRouter := gorillahandlers.LoggingHandler(os.Stdout, r)
	log.Println("Listening at port 8080")
	// http.ListenAndServe(":8080", r)

	log.Fatal(http.ListenAndServe(":8080", r))

}
func final(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Header)

}
