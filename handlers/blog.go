package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
	"github.com/remotejob/kaukotyoeu/dbhandler"
	"github.com/remotejob/kaukotyoeu/domains"
	gcfg "gopkg.in/gcfg.v1"
	mgo "gopkg.in/mgo.v2"
)

var themes string
var locale string

var addrs []string
var database string
var username string
var password string
var mechanism string
var sites []string
var commonwords string
var sitemapsdir string
var mainroute string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		themes = cfg.General.Themes
		locale = cfg.General.Locale

		addrs = cfg.Dbmgo.Addrs
		database = cfg.Dbmgo.Database
		username = cfg.Dbmgo.Username
		password = cfg.Dbmgo.Password
		mechanism = cfg.Dbmgo.Mechanism

		sites = cfg.Sites.Site
		commonwords = cfg.Files.Commonwords
		sitemapsdir = cfg.Dirs.Sitemapsdir
		mainroute = cfg.Routes.Mainroute

	}

}

//CreateArticelePage createPage
func CreateArticelePage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	mtitle := vars["mtitle"]

	fmt.Println(mtitle)

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Timeout:   60 * time.Second,
		Database:  database,
		Username:  username,
		Password:  password,
		Mechanism: mechanism,
	}

	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}
	defer dbsession.Close()
	lp := path.Join("templates", "layout.html")
	lphead := path.Join("templates", "header.html")

	funcMap := template.FuncMap{
		"marshal": func(a []byte) template.JS {

			return template.JS(a)
		},
	}
	t, err := template.New("layout.html").Funcs(funcMap).ParseFiles(lp, lphead)
	check(err)

	article := dbhandler.GetOneArticle(*dbsession, mtitle)

	err = t.Execute(w, article)
	check(err)

}
