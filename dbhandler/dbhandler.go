// Package dbhandler different DB functions
//
// List
//
// In the files of a library package, you can attach a comment to the package
// name, which will be shown in the overview for that package's documentation.
// Just like for command documentation, the convention is to use a file called
// doc.go for your package documentation, though it is not necessary.  If you
// have comments on the package name in multiple files, they'll be concatenated,
// but generally it's best to just have comments in one file.
package dbhandler

import (
	//	"fmt"
	"log"

	"github.com/remotejob/kaukotyoeu/domains"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//InsertLogRecord InsertLogRecord into database
func InsertLogRecord(session mgo.Session, record domains.LogRecord) {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("log").C("logrecords")

	err := c.Insert(record)

	if err != nil {
		panic(err)
	}

}

//GetAllForStatic coments
func GetAllForStatic(session mgo.Session, site string) []domains.Articlefull {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("blog").C("articles")
	var results []domains.Articlefull
	err := c.Find(bson.M{"site": site}).Limit(100).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	return results

}

//GetOneArticle get one artice by mtitle
func GetOneArticle(session mgo.Session, stitle string) domains.Articlefull {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("blog").C("articles")

	var result domains.Articlefull

	err := c.Find(bson.M{"stitle": stitle}).Select(nil).One(&result)
	if err != nil {

		log.Fatal(err)
		//		return
	}

	return result

}

//GetAllSitemaplinks get all articles for sitemap.xml
func GetAllSitemaplinks(session mgo.Session, site string) []domains.Sitemap_from_db {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("blog").C("articles")
	var results []domains.Sitemap_from_db
	err := c.Find(bson.M{"site": site}).Select(bson.M{"stitle": 1, "site": 1, "updated": 1}).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	return results
}

//GetAllUseful probably not used
func GetAllUseful(session mgo.Session, themes string, locale string) []domains.Gphrase {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("gkeywords").C("keywords")

	var results []domains.Gphrase

	err := c.Find(bson.M{"Themes": themes, "Locale": locale}).Select(bson.M{"Phrase": 1, "Rating": 1}).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	return results
}
