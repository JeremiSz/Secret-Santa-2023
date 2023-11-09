package server

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

type subSystem struct {
	http.Server
}

var (
	wishlists = Wishlist{sync.RWMutex{}, [8]string{"Best\ntest"}}
)

const (
	rOOT_PATH  = "/"
	lOGIN_PATH = "/login/"
	dATA_PATH  = "/data/"

	sNOW_CSS_PATH         = "./site/css/snow.css"
	iNDEX_HTML_PATH       = "./site/html/index.html"
	lOGIN_HTML_PATH       = "./site/html/info.html"
	eRROR_HTML_PATH       = "./site/html/error.html"
	iNFO_HTML_PATH        = "./site/html/info.html"
	eRROR_MINI_HTML_PATH  = "./site/html/error_mini.html"
	tARGET_LIST_HTML_PATH = "./site/html/target_list.html"
	cONFIRM_HTML_PATH     = "./site/html/confirm.html"
)

func NewServer() *subSystem {
	router := http.NewServeMux()
	router.HandleFunc(rOOT_PATH, staticViewProvider)
	router.HandleFunc(lOGIN_PATH, loginProvider)
	router.HandleFunc(dATA_PATH, wishlistProvider)

	return &subSystem{Server: http.Server{
		Addr:    ":8080",
		Handler: router,
	}}
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
	id, err := validateCode(r.FormValue("code"))
	if err != nil {
		http.ServeFile(w, r, eRROR_HTML_PATH)
		log.Println(err)
		return
	}
	data := wishlists.LoadWishlists(id)

	template, err := template.ParseFiles(iNFO_HTML_PATH)
	if err != nil {
		http.ServeFile(w, r, eRROR_HTML_PATH)
		log.Println(err)
		return
	}
	template.Execute(w, data)
}

func wishlistProvider(w http.ResponseWriter, r *http.Request) {
	log.Println("wishlists" + r.Method)
	switch r.Method {
	case "GET":
		message := r.FormValue("code")
		id, err := strconv.Atoi(message)
		if err != nil {
			http.ServeFile(w, r, eRROR_MINI_HTML_PATH)
			log.Println(err)
			return
		}
		data := wishlists.LoadWishlists(uint8(id))
		template, err := template.ParseFiles(tARGET_LIST_HTML_PATH)
		if err != nil {
			http.ServeFile(w, r, eRROR_MINI_HTML_PATH)
			log.Println(err)
			return
		}
		template.Execute(w, data)
	case "POST":
		message := r.FormValue("code")
		id, err := strconv.Atoi(message)
		if err != nil {
			http.ServeFile(w, r, eRROR_MINI_HTML_PATH)
			log.Println(err)
			return
		}
		log.Println("got here")
		wishlists.SaveWishlist(uint8(id), r.FormValue("list"))
		http.ServeFile(w, r, cONFIRM_HTML_PATH)
	default:
		http.ServeFile(w, r, eRROR_HTML_PATH)
	}
}
