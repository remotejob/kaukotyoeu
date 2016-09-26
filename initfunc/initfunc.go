package initfunc

import (
	"log"
	"os"

	"github.com/remotejob/kaukotyoeu/domains"
	gcfg "gopkg.in/gcfg.v1"
)

var themes string
var locale string

var addrs []string
var database string
var username string
var password string
var mechanism string
var mainroute string

//GetPar get start parameters
func GetPar() (string, string, []string, string, string, string, string, string) {
	if os.Getenv("SECRET_USERNAME") != "" {

		username = os.Getenv("SECRET_USERNAME")
		password = os.Getenv("SECRET_PASSWORD")
		themes = os.Getenv("THEMES")
		locale = os.Getenv("LOCALE")
		database = os.Getenv("DBADMIN")
		mechanism = "SCRAM-SHA-1"
		addrs = []string{os.Getenv("ADDRS")}
		mainroute = os.Getenv("MAINROUTE")
		log.Println("mongodbpass", password, "database", database)
	} else {
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

			// sites = cfg.Sites.Site
			// commonwords = cfg.Files.Commonwords
			// sitemapsdir = cfg.Dirs.Sitemapsdir
			mainroute = cfg.Routes.Mainroute

		}
	}

	return themes, locale, addrs, database, username, password, mechanism, mainroute
}
