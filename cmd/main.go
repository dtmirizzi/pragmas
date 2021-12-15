package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/dtmirizzi/pragmas/cmd/schema"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
)

//go:generate gojson -input=schema/schema.json -o=schema/schema.go -pkg=schema

//go:embed schema/schema.json
var b []byte

func main() {
	fmt.Println("starting server")

	r := mux.NewRouter()
	r.HandleFunc("/root.json", rootSchema).Methods("GET")
	r.HandleFunc("/root", root).Methods("POST")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("failed to start server ", err)
	}
}

func rootSchema(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(b)
	if err != nil {
		log.Fatal("failed to write root schema")
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	foo := schema.Foo{}
	err := json.Unmarshal(b, &foo)
	if err != nil {
		log.Fatal("failed to write root schema")
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("failed to read body")
	}
	// ensure r has required fields
	for _, f := range foo.Required {
		if !gjson.Get(string(b), f).Exists() {
			_, err := w.Write([]byte("please include missing required field " + f))
			if err != nil {
				log.Fatal("failed to write root schema")
			}
		}
	}

}
