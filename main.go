package main

import (
	"math/rand"
	"net/http"
	"html/template"
	_ "github.com/lib/pq"
	"fmt"
	"github.com/gorilla/mux"
	"time"
	"regexp"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func safeLongUrl(url []byte) string {
	re := regexp.MustCompile("http(s?)://")
	if re.Match(url) {
		return string(url)
	}
	return "http://" + string(url)
}

func generateShort() string {
	return RandStringRunes(5)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, nil)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	long := r.FormValue("url")
	short := generateShort()
	exists, _ := UrlFromLong(long)
	if exists != nil {
		fmt.Fprintf(w, "<h1>Url <i>%s</i> already shorten</h1>"+
			"You can now access your website using this short link : %s",
			exists.LongUrl, exists.ShortUrl)
		return
	}
	url := Url{ShortUrl:short, LongUrl:long, Hits:0}
	NewUrl(url)
	fmt.Fprintf(w, "<h1>New short url save for <i>%s</i></h1>"+
		"You can now access your website using this short link : %s",
		url.LongUrl, url.ShortUrl)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url, _ := UrlFromShort(vars["shorturl"])
	if url != nil {
		http.Redirect(w, r, safeLongUrl([]byte(url.LongUrl)), 301)
	}
	fmt.Fprintf(w, "<h1>No match for url %s</h1>", vars["shorturl"])
}

func main() {
	InitDB("user=urlshortener dbname=urlshortener password=postgrespassword host=postgres sslmode=disable")
	r := mux.NewRouter()
	r.HandleFunc("/create", HomeHandler)
	r.HandleFunc("/shorten", ShortenHandler)
	r.HandleFunc("/{shorturl}", RedirectHandler)
	http.Handle("/", r)


	http.ListenAndServe(":8080", nil)
}

