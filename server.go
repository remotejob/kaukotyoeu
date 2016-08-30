//go:generate  /home/juno/neonworkspace/gowork/bin/statik -src=./public

package main // import "github.com/remotejob/godocker"

import (
	"log"
	"net/http"

	gcfg "gopkg.in/gcfg.v1"

	"github.com/gorilla/mux"
	// _ "github.com/remotejob/godocker/statik"
	"github.com/remotejob/kaukotyoeu/domains"
	"github.com/remotejob/kaukotyoeu/handlers"
	"github.com/remotejob/kaukotyoeu/handlers/robots"
)

var themes string
var locale string

var addrs []string
var database string
var username string
var password string
var mechanism string

var addrsext []string
var databaseext string
var usernameext string
var passwordext string
var mechanismext string

var sites []string

//Employees title name
// type Employees struct {
// 	Title string
// 	Name  string
// }

// func testhandler(w http.ResponseWriter, r *http.Request) {

// 	mongoDBDialInfo := &mgo.DialInfo{
// 		Addrs:     []string{"mymongo"},
// 		Timeout:   60 * time.Second,
// 		Database:  "admin",
// 		Username:  "admin",
// 		Password:  "admin1Rel",
// 		Mechanism: "SCRAM-SHA-1",
// 	}

// 	session, err := mgo.DialWithInfo(mongoDBDialInfo)

// 	defer session.Close()

// 	session.SetMode(mgo.Monotonic, true)

// 	c := session.DB("node-mongo-employee").C("employees")

// 	result := []Employees{}
// 	err = c.Find(nil).All(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, empl := range result {
// 		fmt.Fprintf(w, "Hi  %s %s", empl.Name, empl.Title)
// 	}

// }
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

		addrsext = cfg.Dbmgoext.Addrs
		databaseext = cfg.Dbmgoext.Database
		usernameext = cfg.Dbmgoext.Username
		passwordext = cfg.Dbmgoext.Password
		mechanismext = cfg.Dbmgoext.Mechanism

		sites = cfg.Sites.Site

	}

}
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/robots.txt", robots.Generate)
	r.HandleFunc("/sitemap.xml", handlers.CheckServeSitemap)

	log.Println("Listening at port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))

	// statikFS, err := fs.New()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// http.HandleFunc("/test", testhandler)

	// // fs := http.FileServer(http.Dir("/home/juno/neonworkspace/gowork/src/github.com/remotejob/godocker/assets"))
	// fs := http.FileServer(http.Dir("assets"))

	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// // http.Handle("/assets", http.FileServer(http.Dir("/home/juno/neonworkspace/gowork/src/github.com/remotejob/godocker/assets")))
	// http.Handle("/", http.FileServer(statikFS))
	// http.ListenAndServe(":8080", nil)
}
