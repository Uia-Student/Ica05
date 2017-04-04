package main
import (
"net/http"
"fmt"
"github.com/jzelinskie/geddit"
"log"
)

var Port string = ":8001" // lytte port

func main() {
fmt.Println("Litt oauth stuff...")
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
subOpts := geddit.ListingOptions{ //reddit settings
		Limit: 10,
	}
fpage, _ := o.Frontpage(geddit.DefaultPopularity, subOpts)
fmt.Println(fpage)
fmt.Println("Starting server on port", Port) // debug
http.HandleFunc("/", handler) 
http.HandleFunc("/testfunc",testfunc)
http.ListenAndServe(Port, nil) // starter serveren
}
func handler(w http.ResponseWriter, r *http.Request) {
fmt.Println("Client tried to request /.") //debug
fmt.Fprintf(w, "Hi there, I Love %s!", r.URL.Path[1:]) // Vi skriver dette til "w" som er http.Responewriter.
}

func testfunc(w http.ResponseWriter, r *http.Request) {
fmt.Println("Litt oauth stuff...")
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
subOpts := geddit.ListingOptions{ //reddit settings
		Limit: 10,
	}
fpage, _ := o.Frontpage(geddit.DefaultPopularity, subOpts)
for _, s := range fpage {
		fmt.Fprintf(w,"Title: %s\nAuthor: %s\n\n", s.Title, s.Author)
	}
}