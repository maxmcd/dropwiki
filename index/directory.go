package index

import "fmt"
import "sort"

type directory struct {
	name     string
	contents []node
}

func (f directory) sortingWeight() nodeSortingWeight {
	return directorySortingWeight
}

func (f directory) value() string {
	return f.name
}

func (f directory) children() []node {
	sort.Sort(byWeight(f.contents))
	return f.contents
}

func (f directory) renderIndex(startingDepth int) string {
	renderedChildren := ""
	for _, c := range f.children() {
		renderedChildren += c.renderIndex(startingDepth + 1)
	}

	renderedTitle := fmt.Sprintf("<li><h%d>%s</h%d></li>\n", startingDepth, f.name, startingDepth)
	return renderedTitle + renderedChildren
}
