# dbhandler
--
    import "github.com/remotejob/kaukotyoeu/dbhandler"

Package dbhandler ----> different DB functions

### List

## Usage

#### func  GetAllForStatic

```go
func GetAllForStatic(session mgo.Session, site string) []domains.Articlefull
```
GetAllForStatic get database records for static pages

#### func  GetAllSitemaplinks

```go
func GetAllSitemaplinks(session mgo.Session, site string) []domains.Sitemap_from_db
```
GetAllSitemaplinks get all articles for sitemap.xml

#### func  GetAllUseful

```go
func GetAllUseful(session mgo.Session, themes string, locale string) []domains.Gphrase
```
GetAllUseful probably not used

#### func  GetOneArticle

```go
func GetOneArticle(session mgo.Session, stitle string) domains.Articlefull
```
GetOneArticle get one artice by mtitle

#### func  InsertLogRecord

```go
func InsertLogRecord(session mgo.Session, record domains.LogRecord)
```
InsertLogRecord InsertLogRecord into database
