package index

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// nodeSortingWeight determines the sorting weights for elements
// in a []Node slice
type nodeSortingWeight int

const (
	pageSortingWeight nodeSortingWeight = iota // pages come before directories
	directorySortingWeight
)

// node is the intermediate representation of a directory/file node
type node interface {
	sortingWeight() nodeSortingWeight
	value() string
	children() []node
	renderIndex(int) string
}

// byWeight implements the sort.Interface for []node based
// on the sortingWeight of the elements
type byWeight []node

func (ns byWeight) Len() int           { return len(ns) }
func (ns byWeight) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }
func (ns byWeight) Less(i, j int) bool { return ns[i].sortingWeight() < ns[j].sortingWeight() }

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

// renderIndex generates an HTML index starting from the intermediate node
// representation `start`
func renderIndex(start node) string {
	return fmt.Sprintf("<ul>\n%s</ul>", start.renderIndex(1))
}

// newNodeFrom creates an intermediate representation of the
// direcory tree starting from `startPath`
func newNodeFrom(startPath string) (node, error) {
	name := filepath.Base(startPath)

	info, err := os.Lstat(startPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		// TODO: name needs to be properly formatted/capitalized
		return page{name: name}, nil
	}

	contentsNames, _ := readDirNames(startPath)
	contents := []node{}
	for _, name := range contentsNames {
		path := filepath.Join(startPath, name)
		node, err := newNodeFrom(path)
		if err != nil {
			return nil, err
		}
		contents = append(contents, node)
	}

	return directory{name: name, contents: contents}, nil
}

func readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1) // -1 => read all
	if err != nil {
		return nil, err

	}
	sort.Strings(names)
	return names, nil
}
