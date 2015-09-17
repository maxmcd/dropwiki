package index

import "sort"
import "strconv"

type section struct {
	title    string
	url      string
	contents []node
}

func (s section) sortingWeight() nodeSortingWeight {
	return sectionSortingWeight
}

func (s section) value() string {
	return s.title
}

func (s section) children() []node {
	sort.Sort(byWeight(s.contents))
	return s.contents
}

func (s section) renderIndex(startingDepth int) string {
	renderedChildren := ""
	for _, c := range s.children() {
		renderedChildren += c.renderIndex(startingDepth + 1)
	}

	headerTagName := "h" + strconv.Itoa(startingDepth)
	header := wrapWithTag(s.title, headerTagName)
	linkedHeader := wrapWithAnchorTag(header, s.url)
	renderedTitle := wrapWithTag(linkedHeader, "li")
	return renderedTitle + renderedChildren
}
