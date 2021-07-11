package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

// by implementing method, type URL implement MarshalText interface 
func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL URL `json:"url"`
	Method string `json:"method"`
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{
			URL: URL("/"),
			Method: "GET",
			Description: "See documentation",
		},
		{
			URL: URL("/blocks"),
			Method: "POST",
			Description: "Add a block",
			Payload: "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
	json.NewEncoder(rw).Encode(data) /* this line is equivalent to above three lines */
}

func main() {
	http.HandleFunc("/",documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

