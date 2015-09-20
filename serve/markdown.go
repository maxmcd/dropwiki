package serve

import (
	"io/ioutil"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func generateHtml(path string) (html []byte, err error) {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	unsafe := blackfriday.MarkdownCommon(fileBytes)
	html = bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return
}
