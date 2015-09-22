package index

import (
	"os"
	"path/filepath"
	"sort"
)

// indexTree is an intermediate representation of the index,
// It is generated from a directory and its contents.
// An indexTree is just the root node in the tree.
type indexTree node

// newIndexTree creates an indexTree from the directory tree at `startPath`
func newIndexTree(startPath string) (indexTree, error) {
	title := toTitle(filepath.Base(startPath))
	url := startPath

	isFile, err := isFile(startPath)
	if err != nil {
		return nil, err
	}
	if isFile {
		return page{title: title, url: url}, nil
	}

	contentsNames, _ := readDirContentsNames(startPath)
	contents := []node{}
	for _, name := range contentsNames {
		path := filepath.Join(startPath, name)
		node, err := newIndexTree(path)
		if err != nil {
			return nil, err
		}
		contents = append(contents, node)
	}

	return section{title: title, url: url, contents: contents}, nil
}

func isFile(path string) (bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return false, err
	}
	return !info.IsDir(), nil
}

// readDirContentsNames reads the contents of `dirname` and returns their names, sorted
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
