# domains
--
    import "github.com/remotejob/kaukotyoeu/domains"

Package domains used in project

## Usage

#### type Article

```go
type Article struct {
	Title string
	// Tags      string
	Contents  string
	Mcontents string
	Author    string
}
```

Article shot Articel type

#### type Articlefull

```go
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
```

Articlefull complite entity from DB

#### type Articletotempalte

```go
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
```

Articletotempalte limitede type used for templates

#### type Gphrase

```go
type Gphrase struct {
	Phrase string `bson:"Phrase"`
	Rating int    `bson:"Rating"`
}
```

Gphrase comment

#### type Job

```go
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
```


#### type JobOffer

```go
type JobOffer struct {
	Title       string
	Tags        []string
	Description string
}
```


#### type LogRecord

```go
type LogRecord struct {
	Date time.Time
	Log  string
}
```

LogRecord substitude Nginx log capacity

#### type Page

```go
type Page struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	Lastmod    string   `xml:"lastmod"`
	Changefreq string   `xml:"changefreq"`
}
```

Page sitemap Page

#### type Pages

```go
type Pages struct {
	//	Version string   `xml:"version,attr"`
	XMLName xml.Name `xml:"urlset"`
	XmlNS   string   `xml:"xmlns,attr"`
	//	XmlImageNS string   `xml:"xmlns:image,attr"`
	//	XmlNewsNS  string   `xml:"xmlns:news,attr"`
	Pages []*Page `xml:"url"`
}
```

Pages struct keep sitemap obj

#### type ServerConfig

```go
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
```

ServerConfig config struc for site

#### type SitemapObj

```go
type SitemapObj struct {
	Changefreq    string
	Hoursduration float64
	Loc           string
	Lastmod       string
}
```

SitemapObj obj for sitemap

#### type Sitemap_from_db

```go
type Sitemap_from_db struct {
	Stitle  string
	Site    string
	Updated time.Time
}
```

Sitemap_from_db only info for Sitemap format
