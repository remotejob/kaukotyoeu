//Package domains struct used in project
package domains

import (
	"encoding/xml"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//LogRecord substitude Nginx log capacity
type LogRecord struct {
	Date  time.Time
	Log   string
	Ltype string
}

// Gphrase comment
type Gphrase struct {
	Phrase string `bson:"Phrase"`
	Rating int    `bson:"Rating"`
}

//Articlefull complite entity from DB
type Articlefull struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Title  string
	Stitle string
	// Tags      string
	Contents  string
	Mcontents string
	Site      string
	Author    string
	Created   time.Time
	Updated   time.Time
}

//Articletotempalte limitede type used for templates
type Articletotempalte struct {
	Title  string
	Stitle string
	// Tags      string
	Contents  string
	Mcontents string
	Site      string
	Created   string
	Updated   string
	Jsonld    []byte
}

//Article shot Articel type
type Article struct {
	Title string
	// Tags      string
	Contents  string
	Mcontents string
	Author    string
}

//Sitemap_from_db only info for Sitemap format
type Sitemap_from_db struct {
	Stitle  string
	Site    string
	Updated time.Time
}

//ServerConfig config struc for site
type ServerConfig struct {
	General struct {
		Themes string
		Locale string
	}
	Dbmgo struct {
		Addrs     []string
		Database  string
		Username  string
		Password  string
		Mechanism string
	}

	Dbmgoext struct {
		Addrs     []string
		Database  string
		Username  string
		Password  string
		Mechanism string
	}
	Sites struct {
		Site []string
	}
	Dirs struct {
		Sitemapsdir string
		Webrootdir  string
	}

	Routes struct {
		Mainroute string
	}

	Files struct {
		Commonwords string
	}
}

type JobOffer struct {
	Title       string
	Tags        []string
	Description string
}

//SitemapObj obj for sitemap
type SitemapObj struct {
	Changefreq    string
	Hoursduration float64
	Loc           string
	Lastmod       string
}

//Pages struct keep sitemap obj
type Pages struct {
	//	Version string   `xml:"version,attr"`
	XMLName xml.Name `xml:"urlset"`
	XmlNS   string   `xml:"xmlns,attr"`
	//	XmlImageNS string   `xml:"xmlns:image,attr"`
	//	XmlNewsNS  string   `xml:"xmlns:news,attr"`
	Pages []*Page `xml:"url"`
}

//Page sitemap Page
type Page struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	Lastmod    string   `xml:"lastmod"`
	Changefreq string   `xml:"changefreq"`
	//	Name       string   `xml:"news:news>news:publication>news:name"`
	//	Language   string   `xml:"news:news>news:publication>news:language"`
	//	Title      string   `xml:"news:news>news:title"`
	//	Keywords   string   `xml:"news:news>news:keywords"`
	//	Image      string   `xml:"image:image>image:loc"`
}

type Job struct {
	Maintitle string
	Subtitle  string
	Jobs      []struct {
		Name string
		Path string
		Img  string
		Item []struct {
			Title    string
			Rank     int
			Duration string
			Position string
			Details  string
			Location string
			Country  string
		}
	}
}
