package main
import (
"net/http"
"fmt"
)

var Port string = ":8001" // lytte port
func main() {
fmt.Println("Starting server on port", Port) // debug
http.HandleFunc("/",handler) //Vær gang en client spør etter "/" på serveren, blir de tatt til funksjonen handler.
http.ListenAndServe(Port, nil) // starter serveren
}
func handler(w http.ResponseWriter, r *http.Request) {
fmt.Println("Client tried to request /.") //debug
fmt.Fprintf(w, "Hi there, I Love %s!", r.URL.Path[1:]) // Vi skriver dette til "w" som er http.Responewriter.
}