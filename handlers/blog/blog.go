package blog

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/remotejob/kaukotyoeu/dbhandler"
	"github.com/remotejob/kaukotyoeu/domains"
	"github.com/remotejob/kaukotyoeu/handlers/insertlog"
	"github.com/remotejob/kaukotyoeu/initfunc"
	"github.com/remotejob/kaukotyoeu/ldjsonhandler"
	shuffle "github.com/shogo82148/go-shuffle"
	mgo "gopkg.in/mgo.v2"
)

var themes string
var locale string

var addrs []string
var database string
var username string
var password string
var mechanism string

func checkReq(r *http.Request) {

	if strings.Contains(r.Referer(), "www.google") {

		now := time.Now()
		log := r.Referer() + "," + r.RequestURI
		record := domains.LogRecord{Date: now,
			Log: log}
		go insertlog.InsertIntoDB(record)

	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	themes, locale, addrs, database, username, password, mechanism, _ = initfunc.GetPar()

}

//CreateArticelePage createPage
func CreateArticelePage(w http.ResponseWriter, r *http.Request) {

	checkReq(r)

	vars := mux.Vars(r)

	mtitle := vars["mtitle"]

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

	article := dbhandler.GetOneArticle(*dbsession, mtitle)

	if len(article) == 1 {

		lp := path.Join("templates", "layout.html")
		lphead := path.Join("templates", "header_common.html")

		funcMap := template.FuncMap{
			"marshal": func(a domains.Articlefull) template.JS {

				var articles []domains.Articlefull

				articles = append(articles, a)

				b := ldjsonhandler.Create(articles, "Selected Article")

				return template.JS(b)
			},
			"title": func(a domains.Articlefull) string {

				return a.Title
			},
		}
		t, err := template.New("layout.html").Funcs(funcMap).ParseFiles(lp, lphead)
		check(err)

		err = t.Execute(w, article[0])
		check(err)

	} else {
		http.NotFound(w, r)
		// http.Error(w, http.StatusText(404), 404)
	}

}

//CreateIndexPage create Index Page
func CreateIndexPage(w http.ResponseWriter, r *http.Request) {

	sitefull := r.Host
	site := strings.Split(sitefull, ":")[0]

	log.Println("site", site)

	if site == "localhost" {

		site = "www.kaukotyo.eu"

	}

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

	lp := path.Join("templates", "home_page.html")

	headercommon := path.Join("templates", "header_common.html")
	funcMap := template.FuncMap{
		"marshal": func(articles []domains.Articlefull) template.JS {
			b := ldjsonhandler.Create(articles, "Index Page")
			return template.JS(b)
		},
		"title": func(a []domains.Articlefull) string {

			return "Index Page"
		},
	}

	t, err := template.New("home_page.html").Funcs(funcMap).ParseFiles(lp, headercommon)
	check(err)

	allarticles := dbhandler.GetAllForStatic(*dbsession, site)

	if len(allarticles) > 0 {

		var numberstoshuffle []int
		for num := range allarticles {

			numberstoshuffle = append(numberstoshuffle, num)

		}
		rand.Seed(time.Now().UTC().UnixNano())

		shuffle.Ints(numberstoshuffle)

		var atricleToInject []domains.Articlefull

		for c, i := range numberstoshuffle {

			if c < 10 {

				atricleToInject = append(atricleToInject, allarticles[i])
			}

		}

		if len(atricleToInject) > 0 {

			log.Println(len(atricleToInject))

			err = t.Execute(w, atricleToInject)
			check(err)
		}

	} else {
		http.Error(w, http.StatusText(404), 404)
	}

}
