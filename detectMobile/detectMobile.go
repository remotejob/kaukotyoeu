package detectMobile

import (
	"fmt"
	"net/http"

	"github.com/remotejob/goDevice"
)

//Detect return true if mobile
func Detect(w http.ResponseWriter, r *http.Request) {

	deviceType := goDevice.GetType(r)

	if deviceType == "Mobile" {
		fmt.Fprintf(w, "<h1>Mobile</h1>")
	} else if deviceType == "Web" {
		fmt.Fprintf(w, "<h1>Web</h1>")
	} else if deviceType == "Tab" {
		fmt.Fprintf(w, "<h1>Tablet</h1>")
	}

}
