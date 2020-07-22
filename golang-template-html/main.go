package main

import (
	"log"
	"net/http"
	"path"
	"text/template"
)

// HelloTemplate is ...
func HelloTemplate(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)

	var data = map[string]interface{}{
		"name": "Thnovice",
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", HelloTemplate)
	log.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
