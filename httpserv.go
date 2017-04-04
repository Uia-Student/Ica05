package main

import (
	"net/http"
	"fmt"
	"github.com/jzelinskie/geddit"
	"log"
)

var Port string = ":8001" // lytte port

func main() {
	fmt.Println("Prover aa logge inn...")
	fmt.Println("Setter oauth info")
	o, err := geddit.NewOAuthSession(
		"JjJUDNL57-hisw",
		"ZcD6YBUPJfh1ob1iu3WMbcp_H5g",
		"http://none",
		"http://www.google.com",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logger inn")
	err = o.LoginAuth("uiaica05", "uiaica05123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Login var successfull.")
	//subOpts := geddit.ListingOptions{ //reddit settings
	//	Limit: 10,
	//}
	//fpage, _ := o.Frontpage(geddit.DefaultPopularity, subOpts)
	fmt.Println("Registring handlers") // debug
	http.HandleFunc("/", frontpageHandler)
	http.HandleFunc("/reddit", redditHandler)
	http.HandleFunc("/news", worldnewsHandler)
	http.HandleFunc("/norge", norgeHandler)
		fmt.Println("Starting server on port", Port) // debug
	http.ListenAndServe(Port, nil) // starter serveren
}
func frontpageHandler(w http.ResponseWriter, r *http.Request) {           //debug
	fmt.Fprintf(w, "Welcome to Nerds With Attitude's front page. Avaible links:")
}

func redditHandler(w http.ResponseWriter, r *http.Request) {
	o, err := geddit.NewOAuthSession(
		"JjJUDNL57-hisw",
		"ZcD6YBUPJfh1ob1iu3WMbcp_H5g",
		"http://none",
		"http://www.google.com",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = o.LoginAuth("uiaica05", "uiaica05123")
	if err != nil {
		log.Fatal(err)
	}
	subOpts := geddit.ListingOptions{ //reddit settings
		Limit: 20,
	}
	fpage, _ := o.Frontpage(geddit.DefaultPopularity, subOpts)
	for _, s := range fpage {
		fmt.Fprintf(w, "Title: %s\nAuthor: %s\nComments: %v\nPoints: %v\nURL: <a href=\x22%s\x22>Link</a>\nReddit URL: reddit.com/%s\n\n", s.Title, s.Author, s.NumComments, s.Score, s.URL,s.Permalink)
	}	
	}
	
func worldnewsHandler(w http.ResponseWriter, r *http.Request) {
	o, err := geddit.NewOAuthSession(
		"JjJUDNL57-hisw",
		"ZcD6YBUPJfh1ob1iu3WMbcp_H5g",
		"http://none",
		"http://www.google.com",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = o.LoginAuth("uiaica05", "uiaica05123")
	if err != nil {
		log.Fatal(err)
	}
	subOpts := geddit.ListingOptions{ //reddit settings
		Limit: 20,
	}
	fpage, _ := o.SubredditSubmissions("worldnews",geddit.DefaultPopularity, subOpts)
	for _, s := range fpage {
		fmt.Fprintf(w, "Title: %s\nAuthor: %s\nComments: %v\nPoints: %v\n\n", s.Title, s.Author, s.NumComments, s.Score)
	}
}
func norgeHandler(w http.ResponseWriter, r *http.Request) {
	o, err := geddit.NewOAuthSession(
		"JjJUDNL57-hisw",
		"ZcD6YBUPJfh1ob1iu3WMbcp_H5g",
		"http://none",
		"http://www.google.com",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = o.LoginAuth("uiaica05", "uiaica05123")
	if err != nil {
		log.Fatal(err)
	}
	subOpts := geddit.ListingOptions{ //reddit settings
		Limit: 20,
	}
	fpage, _ := o.SubredditSubmissions("norge",geddit.DefaultPopularity, subOpts)
	for _, s := range fpage {
		fmt.Fprintf(w, "Title: %s\nAuthor: %s\nComments: %v\nPoints: %v\n\n", s.Title, s.Author, s.NumComments, s.Score)
	}
}


