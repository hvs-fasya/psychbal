package front

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	StaticPath = ""
)

func InitFront(staticPath string) {
	StaticPath = staticPath
}

//IndexHandler http handler for index.html file
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(StaticPath + "index.html")
	http.ServeFile(w, r, StaticPath+"index.html")
}

//FaviconHandler http handler for favicon.ico file
func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["file"]
	w.Header().Set("Content-Type", "image/x-icon")
	http.ServeFile(w, r, StaticPath+f)
}

//JSHandler http handler for js files
func JSHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	vars := mux.Vars(r)
	file := vars["js"]
	http.ServeFile(w, r, StaticPath+file)
}

//FontsHandler http handler for font files
func FontsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["font"]
	http.ServeFile(w, r, StaticPath+"fonts/"+f)
}
