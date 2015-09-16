package index

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// NodeSortingWeight determines the sorting weights for elements
// in a []Node slice
type NodeSortingWeight int

const (
	PageSortingWeight NodeSortingWeight = iota // Pages come before Directories
	DirectorySortingWeight
)

type Node interface {
	SortingWeight() NodeSortingWeight
	Value() string
	Children() []Node
	renderIndex(int) string
}

// ByWeight implements the sort.Interface for []Node based
// on the SortingWeight of the elements
type ByWeight []Node

func (ns ByWeight) Len() int           { return len(ns) }
func (ns ByWeight) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }
func (ns ByWeight) Less(i, j int) bool { return ns[i].SortingWeight() < ns[j].SortingWeight() }

type Page struct {
	Name string
}

func (p Page) SortingWeight() NodeSortingWeight {
	return PageSortingWeight
}

func (p Page) Value() string {
	return p.Name
}

func (p Page) Children() []Node {
	// a Page has no Children
	return []Node{}
}

func (p Page) renderIndex(d int) string {
	return fmt.Sprintf("<li>%s</li>\n", p.Name)
}

type Directory struct {
	Name     string
	Contents []Node
}

func (f Directory) SortingWeight() NodeSortingWeight {
	return DirectorySortingWeight
}

func (f Directory) Value() string {
	return f.Name
}

func (f Directory) Children() []Node {
	sort.Sort(ByWeight(f.Contents))
	return f.Contents
}

func (f Directory) renderIndex(startingDepth int) string {
	renderedChildren := ""
	for _, c := range f.Children() {
		renderedChildren += c.renderIndex(startingDepth + 1)
	}

	renderedTitle := fmt.Sprintf("<li><h%d>%s</h%d></li>\n", startingDepth, f.Name, startingDepth)
	return renderedTitle + renderedChildren
}

func RenderIndex(start Node) string {
	return fmt.Sprintf("<ul>\n%s</ul>", start.renderIndex(1))
}

// recursively make a Node from a path
func NewNodeFrom(startPath string) (Node, error) {
	name := filepath.Base(startPath)

	info, err := os.Lstat(startPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		// TODO: name needs to be properly formatted/capitalized
		return Page{Name: name}, nil
	}

	contentsNames, _ := readDirNames(startPath)
	contents := []Node{}
	for _, name := range contentsNames {
		path := filepath.Join(startPath, name)
		node, err := NewNodeFrom(path)
		if err != nil {
			return nil, err
		}
		contents = append(contents, node)
	}

	return Directory{Name: name, Contents: contents}, nil
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
