# dbhandler
--
    import "github.com/remotejob/kaukotyoeu/dbhandler"

Package dbhandler different DB functions


### List

In the files of a library package, you can attach a comment to the package name,
which will be shown in the overview for that package's documentation. Just like
for command documentation, the convention is to use a file called doc.go for
your package documentation, though it is not necessary. If you have comments on
the package name in multiple files, they'll be concatenated, but generally it's
best to just have comments in one file.

## Usage

#### func  GetAllForStatic

```go
func GetAllForStatic(session mgo.Session, site string) []domains.Articlefull
```
GetAllForStatic coments

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
