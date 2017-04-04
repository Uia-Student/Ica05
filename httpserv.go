package main

import (
	"net/http"
	"fmt"
	"github.com/jzelinskie/geddit"
	"log"
	"html/template"
	"io/ioutil"
)
type Page struct {
    Title string
    Body  []byte
}
var Port string = ":8001" // lytte port

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func main() {

	fmt.Println("Prover aa logge inn...") //Testing login crendicals, to check if they are valid
	fmt.Println("Setter oauth info")
	o, err := geddit.NewOAuthSession( //Create a new auth session
		"JjJUDNL57-hisw", // secret ID
		"ZcD6YBUPJfh1ob1iu3WMbcp_H5g", // Unknown
		"http://none", //URL's
		"http://www.google.com", /// URL's
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logger inn")
	err = o.LoginAuth("uiaica05", "uiaica05123") // Username + Password
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Login var successfull.")
	//subOpts := geddit.ListingOptions{ //reddit settings
	//	Limit: 10,
	//}
	//fpage, _ := o.Frontpage(geddit.DefaultPopularity, subOpts)
	fmt.Println("Registring handlers") // debug
	http.HandleFunc("/",personalHandler) // Add a hook to the handler
	http.HandleFunc("/about",frontpageHandler) // About page
	
		fmt.Println("Starting server on port", Port) // debug
	http.ListenAndServe(Port, nil) // starter serveren // Ready to serve
}
func frontpageHandler(w http.ResponseWriter, r *http.Request) {          
	fmt.Fprintf(w, "Welcome to Nerds With Attitude's front page\nThis webpage is a shittier version of Reddit.com\n Write a subreddit at the end of the url to visit it.\n Example: localhost:8001/movies")
}
func personalHandler(w http.ResponseWriter, r *http.Request) {

    title := r.URL.Path[len("/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }

    t, _ := template.ParseFiles("view.html")
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
	subOpts := geddit.ListingOptions{ //reddit settings, 20 posts
		Limit: 20,
	}
	fpage, _ := o.SubredditSubmissions(p.Title, geddit.DefaultPopularity, subOpts) // Get subreddit from title, Default sorting, max 20 posts
	t.Execute(w, p) // Execute the template
	for _, s := range fpage { // Display all posts in order.
		fmt.Fprintf(w,"<div><br>Title: %s<br>Author: %s<br>Comments: %v<br>Points: %v<br>URL: <a href=\x22%s\x22>Link</a><br>Reddit URL: <a href=\x22https://www.reddit.com/%s\x22>Click here</a></div><br>", s.Title, s.Author, s.NumComments, s.Score, s.URL,s.Permalink)
	}
	
}
/* Code that is not in use, was used early on.
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
*/ 

