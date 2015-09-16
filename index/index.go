package index

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// RenderIndex generates an HTML index from the directory structure
// of `path`
func RenderIndex(path string) (string, error) {
	nodeRepresentation, err := newNodeFrom(path)
	if err != nil {
		return "", err
	}

	renderedIndex := renderIndex(nodeRepresentation)

	return renderedIndex, nil
}

// renderIndex generates an HTML index from the intermediate
// representation `start`
func renderIndex(start node) string {
	return fmt.Sprintf("<ul>\n%s</ul>", start.renderIndex(1))
}

// newNodeFrom creates an intermediate representation of the
// index from the directory tree at `startPath`
func newNodeFrom(startPath string) (node, error) {
	name := toTitle(filepath.Base(startPath))

	info, err := os.Lstat(startPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return page{title: name}, nil
	}

	contentsNames, _ := readDirContentsNames(startPath)
	contents := []node{}
	for _, name := range contentsNames {
		path := filepath.Join(startPath, name)
		node, err := newNodeFrom(path)
		if err != nil {
			return nil, err
		}
		contents = append(contents, node)
	}

	return section{title: name, contents: contents}, nil
}

// readDirContentsNames reads the contents of `dirname` and returns their names, soted
func readDirContentsNames(dirname string) ([]string, error) {
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
