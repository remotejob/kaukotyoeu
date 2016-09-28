//Package insertlog insert log into DB
package insertlog

import (
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

func init() {

	themes, locale, addrs, database, username, password, mechanism, _ = initfunc.GetPar()

}

//InsertIntoDB Insert Log Into mongoDB
func InsertIntoDB(record domains.LogRecord) {
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

	dbhandler.InsertLogRecord(*dbsession, record)

}
