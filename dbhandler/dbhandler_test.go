package dbhandler

import (
	//	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	"fmt"
	"testing"
	"time"

	mgo "gopkg.in/mgo.v2"
	//	"time"
	//	"fmt"
)

func TestGetAllUseful(t *testing.T) {

	addrs := []string{"104.131.38.162"}

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Timeout:   60 * time.Second,
		Database:  "admin",
		Username:  "admin",
		Password:  "admin1Rel",
		Mechanism: "SCRAM-SHA-1",
	}

	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	all := GetAllUseful(*dbsession, "realestate", "ru_RU")

	fmt.Println("all", all[10], len(all))

}
