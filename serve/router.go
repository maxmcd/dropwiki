package serve

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found :("))
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	location := findFile(path)
	if location == "" {
		notFound(w)
		return
	}

	html, err := generateHtml(location)
	if err != nil {
		notFound(w)
		return
	}

	rsp := response{
		Title:   "title",
		Index:   "index",
		Content: string(html),
	}
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	err = T.Execute(w, rsp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func findFile(path string) string {
	path = "." + path

	if len(path) > 3 && path[len(path)-2:] == "md" {
		return path
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if file.IsDir() == false {
			if strings.EqualFold(file.Name(), "readme.md") {
				return path + "/readme.md"
			}
		}
	}
	return ""
}
