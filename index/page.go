package index

// page represents a file/page in an indexTree
type page struct {
	title string
	url   string
}

func (p page) sortingWeight() nodeSortingWeight {
	return pageSortingWeight
}

func (p page) value() string {
	return p.title
}

func (p page) children() []node {
	// a Page has no Children
	return []node{}
}

func (p page) renderIndex(d int) string {
	return wrapWithTag(wrapWithAnchorTag(p.title, p.url), "li")
}
