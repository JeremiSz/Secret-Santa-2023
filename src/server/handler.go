package server

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

type subSystem struct {
	http.Server
	wishlist [8]string
}

var (
	vALID_CODES = [8]uint16{2350, 1151, 2322, 2453, 1034, 1945, 1606, 1587}
)

const (
	rOOT_PATH       = "/"
	lOGIN_PATH      = "/login/"
	iD_MASK         = 7
	sNOW_CSS_PATH   = "./site/css/snow.css"
	iNDEX_HTML_PATH = "./site/html/index.html"
	lOGIN_HTML_PATH = "./site/html/info.html"
	eRROR_HTML_PATH = "./site/html/error.html"
)

func NewServer() *subSystem {
	router := http.NewServeMux()
	router.HandleFunc(rOOT_PATH, staticViewProvider)
	router.HandleFunc(lOGIN_PATH, loginProvider)

	return &subSystem{Server: http.Server{
		Addr:    ":8080",
		Handler: router,
	}, wishlist: [8]string{"", "", "", "", "", "", "", ""}}
}
func staticViewProvider(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len(rOOT_PATH):]
	log.Println("serving" + path)
	serveStatic(path, w, r)
}
func serveStatic(path string, w http.ResponseWriter, r *http.Request) {
	switch path {
	case "snow.css":
		http.ServeFile(w, r, "./site/css/snow.css")
	default:
		http.ServeFile(w, r, "./site/html/index.html")
	}
}
func loginProvider(w http.ResponseWriter, r *http.Request) {

	id, err := extractCode(r.FormValue("code"))
	if err != nil {
		http.ServeFile(w, r, eRROR_HTML_PATH)
		log.Println(err)
		return
	}

	//check if id is in valid group
	{
		i := 0
		isValid := false
		for i < len(vALID_CODES) && !isValid {
			isValid = uint16(id) == vALID_CODES[i]
			i++
		}

		if !isValid {
			err = errors.New("invalid id")
		}
	}

	if err != nil {
		http.ServeFile(w, r, eRROR_HTML_PATH)
		log.Println(err)
		return
	}

	//id = id & iD_MASK
	http.ServeFile(w, r, lOGIN_HTML_PATH)

}

func extractCode(message string) (uint16, error) {
	if len(message) < 5 {
		return 0, errors.New("empty message")
	}
	digits := message[0:2] + message[3:5]
	code, err := strconv.Atoi(digits)
	if err != nil {
		return 0, err
	}
	return uint16(code), nil
}
