package index

import "strings"

func toTitle(filename string) string {
	r := strings.NewReplacer("_", " ", ".md", "")
	return strings.Title(r.Replace(filename))
}
