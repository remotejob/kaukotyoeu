package handlers

import (
	// "bytes"
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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
var resultXML []byte

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
func CheckServeSitemap(w http.ResponseWriter, r *http.Request) {

	sitefull := r.Host
	site := strings.Split(sitefull, ":")[0]

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

	allsitemaplinks := dbhandler.GetAllSitemaplinks(*dbsession, site)

	fmt.Println("CheckServeSitemap")

	docList := new(domains.Pages)
	docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

	for _, sitemaplink := range allsitemaplinks {

		if sitemaplink.Site == site {

			doc := new(domains.Page)
			doc.Loc = "http://" + site + "/" + themes + "/" + locale + "/" + mainroute + "/" + sitemaplink.Stitle + ".html"
			doc.Lastmod = sitemaplink.Updated.Format(time.RFC3339)
			doc.Changefreq = "monthly"
			docList.Pages = append(docList.Pages, doc)
			// fmt.Println(site, sitemaplink.Stitle)
		}

	}

	resultXML, err = xml.MarshalIndent(docList, "", "  ")
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(bytes.NewBuffer(resultXML).String())

	// fmt.Println(bytes.NewBuffer(sitemap).String())

	w.Header().Add("Content-type", "application/xml")
	w.Write(resultXML)

}
