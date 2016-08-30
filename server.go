//go:generate  /home/juno/neonworkspace/gowork/bin/statik -src=./public

package main // import "github.com/remotejob/server"

import (
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

// func init() {

// 	var cfg domains.ServerConfig
// 	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
// 		log.Fatalln(err.Error())

// 	} else {

// 		themes = cfg.General.Themes
// 		locale = cfg.General.Locale

// 		addrs = cfg.Dbmgo.Addrs
// 		database = cfg.Dbmgo.Database
// 		username = cfg.Dbmgo.Username
// 		password = cfg.Dbmgo.Password
// 		mechanism = cfg.Dbmgo.Mechanism

// 		addrsext = cfg.Dbmgoext.Addrs
// 		databaseext = cfg.Dbmgoext.Database
// 		usernameext = cfg.Dbmgoext.Username
// 		passwordext = cfg.Dbmgoext.Password
// 		mechanismext = cfg.Dbmgoext.Mechanism

// 		sites = cfg.Sites.Site

// 	}

// }
func main() {

	fs := http.FileServer(http.Dir("assets"))

	r := mux.NewRouter()
	r.HandleFunc("/robots.txt", robots.Generate)
	r.HandleFunc("/sitemap.xml", handlers.CheckServeSitemap)
	r.HandleFunc("/job/{locale}/{themes}/{mtitle}.html", handlers.CreateArticelePage)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	r.HandleFunc("/", handlers.CreateIndexPage)
	log.Println("Listening at port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}
