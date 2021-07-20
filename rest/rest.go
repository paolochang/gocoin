package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/paolochang/gocoin/blockchain"
	"github.com/paolochang/gocoin/utils"
)

var port string

type url string

// by implementing method, type URL implement MarshalText interface 
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL url `json:"url"`
	Method string `json:"method"`
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription {
		{
			URL: url("/"),
			Method: "GET",
			Description: "See documentation",
		},
		{
			URL: url("/blocks"),
			Method: "GET",
			Description: "See all blocks",
		},
		{
			URL: url("/blocks"),
			Method: "POST",
			Description: "Add a block",
			Payload: "data:string",
		},
		{
			URL: url("/blocks/{height}"),
			Method: "GET",
			Description: "See a block",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
	json.NewEncoder(rw).Encode(data) /* this line is equivalent to above three lines */
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block := blockchain.GetBlockchain().GetBlock(id)
	json.NewEncoder(rw).Encode(block)
}

func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.HandleFunc("/",documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}