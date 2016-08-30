package sitemaps_maker

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/remotejob/kaukotyoeu/dbhandler"
	"github.com/remotejob/kaukotyoeu/domains"
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
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
func Create(site string) []byte {
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

	for _, site := range sites {

		docList := new(domains.Pages)
		docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

		for _, sitemaplink := range allsitemaplinks {

			if sitemaplink.Site == site {

				doc := new(domains.Page)
				doc.Loc = "http://" + site + "/" + themes + "/" + locale + "/" + mainroute + "/" + sitemaplink.Stitle + ".html"
				doc.Lastmod = sitemaplink.Updated.Format(time.RFC3339)
				doc.Changefreq = "monthly"
				docList.Pages = append(docList.Pages, doc)
				fmt.Println(site, sitemaplink.Stitle)
			}

		}

		resultXML, err = xml.MarshalIndent(docList, "", "  ")
		if err != nil {

			//		golog.Crit(err.Error())
			log.Println(err.Error())
		}
		// ioutil.WriteFile(sitemapsdir+"/sitemap_"+site+".xml", resultXML, 0644)
		// if err != nil {

		// 	//		golog.Crit(err.Error())
		// 	log.Println(err.Error())
		// }

	}

	return resultXML

}
