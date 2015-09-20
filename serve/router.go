package serve

import (
	"log"
	"net/http"
	"text/template"
)

// node is a node in an indexTree
// stealing this from index to search available files

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
    <div class="menu">
        {{.Index}}
    </div>
    <div class="content">
        {{.Content}}
    </div>
</body>
</html>
`

var T *template.Template

func init() {
	T = template.Must(template.New("handler").Parse(htmlTemplate))
}

func ServeFiles() {

	log.Fatal(http.ListenAndServe(":8001", http.HandlerFunc(handler)))
}

type response struct {
	Title   string
	Index   string
	Content string
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	location := findFile(path)
	if location == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found :("))
		return
	}

	html, err := generateHtml(location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	rsp := response{
		Title:   "title",
		Index:   "index",
		Content: "html",
	}
	w.Header().Write("text/html")
	err = T.Execute(w, rsp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func findFile(path string) (location string) {
	// ?
}
