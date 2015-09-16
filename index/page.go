package index

import "fmt"

type page struct {
	name string
}

func (p page) sortingWeight() nodeSortingWeight {
	return pageSortingWeight
}

func (p page) value() string {
	return p.name
}

func (p page) children() []node {
	// a Page has no Children
	return []node{}
}

func (p page) renderIndex(d int) string {
	return fmt.Sprintf("<li>%s</li>\n", p.name)
}
