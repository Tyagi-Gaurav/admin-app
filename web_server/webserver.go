package web_server

import (
	"log"
	"net/http"
	"path"
	"text/template"
)

const template_dir = "web_server/templates"

type WebServer struct {
}

func (ws *WebServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(template_dir, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("Error: Invalid HTTP method")
	}
}

func (ws *WebServer) Run() {
	http.HandleFunc("/", ws.Index)
}
