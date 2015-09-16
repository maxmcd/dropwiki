package index

import "fmt"
import "sort"

type section struct {
	title    string
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

	renderedTitle := fmt.Sprintf("<li><h%d>%s</h%d></li>\n", startingDepth, s.title, startingDepth)
	return renderedTitle + renderedChildren
}
