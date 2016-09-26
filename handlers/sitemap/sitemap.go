package sitemap

import (
	"encoding/xml"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/remotejob/kaukotyoeu/dbhandler"
	"github.com/remotejob/kaukotyoeu/domains"
	"github.com/remotejob/kaukotyoeu/initfunc"
	mgo "gopkg.in/mgo.v2"
)

var themes string
var locale string

var addrs []string
var database string
var username string
var password string
var mechanism string

var mainroute string
var resultXML []byte

func init() {

	themes, locale, addrs, database, username, password, mechanism, mainroute = initfunc.GetPar()

}

//CheckServeSitemap create dinamic sitemap.xml file
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

	w.Header().Add("Content-type", "application/xml")
	_, err = w.Write(resultXML)
	if err != nil {
		log.Println(err.Error())
	}

}
