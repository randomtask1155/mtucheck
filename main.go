package main


import (
	"html/template"
	"net/http"
	"os"
	"fmt"
	"encoding/json"
)

var (
	startPageTemplate = template.Must(template.ParseFiles("index.tmpl"))
)

type DataRequest struct {
	Size int `json:"dataSize"`
}

type DataResponse struct {
	DataString string `json:"dataString"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	startPageTemplate.Execute(w, nil)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	var b []byte
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dreq := new(DataRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dreq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("<p>Error: %s</p>", err )))
		return
	}

	d := ""
	for i := 0; i < dreq.Size; i++ {
		d += "a"
	}
	resp := new(DataResponse)
	resp.DataString = d
	jdata, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("<p>Error: %s</p>", err )))
		return
	}
	w.Write(jdata)
}

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc( "/header", okHandler)
	http.HandleFunc("/data", dataHandler)

	// File serving handlers
	http.Handle("/assets/", http.FileServer(http.Dir("")))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		fmt.Printf("Failed to start http server: %s\n", err)
	}
	return
}

