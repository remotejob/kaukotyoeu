//Kauko Työ
//
// kaukotyoeu
//
//Project demonstrate real application installed on Google Container Engine
//
// * Working link http://kaukotyo.eu/
//
// * Main difference from previous project https://github.com/remotejob/clusters_export/tree/master/docker-kaukotyo
// in use golang web server instead of Nginx.
//
//First component it's server.go (golang web server) file
//
// * static contents as well as templates (assets,templates dirs)  incorporated in docker container by command COPY in Dockerfile
//
//  COPY assets/ /assets/
//  COPY templates /templates/
//
// * We are need only one line in code to serve all assets contents r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
//
// * to make some visual effects used small JavaScript fragment in templates files (home_page.html, layout.html)
//
// * credential taken from config.gcfg but better use http://kubernetes.io/docs/user-guide/security-context/
//
//Second element Standard mongodb image
//
// * to make it more reliable and powerful used kubernetes controller mydb-controller.yml
// increase "replicas: 1" from 1 to more.
//
// * args: ["--auth"] line include authentication for mongodb so after it possible access from outside.
//
// * file mydb-service.yml expose DB for work.
//
//Start UP project
//
// * create Standard persistent disk in Google Cloud Platform (for DBase) name it "mymongo-disk"
//
// * modify mydb-controller.yml first without args: ["--auth"] create  authentication for mongodb after restart controller
//
// * kubectl create -f mydb-controller.yml
//
// * after included authentication expose service.
//
// * kubectl create -f mydb-service.yml
//
// * modify Makefile
//
// * last command "make"
//
//To fill contents take a look at other project
//
//https://github.com/remotejob/kaukotyoeu_utils
package main

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
