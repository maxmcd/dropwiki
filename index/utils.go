package index

import "strings"
import "fmt"

func toTitle(filename string) string {
	r := strings.NewReplacer("_", " ", ".md", "")
	return strings.Title(r.Replace(filename))
}

func wrapWithTag(content, tagName string) string {
	return fmt.Sprintf("<%s>%s</%s>", tagName, content, tagName)
}

func wrapWithAnchorTag(content, href string) string {

	return fmt.Sprintf("<a href='%s'>%s</a>", href, content)
}
