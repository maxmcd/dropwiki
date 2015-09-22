package index

// GenerateIndex generates an HTML index from the directory structure
// of `path`
func GenerateIndex(path string) (string, error) {
	intermediateIndexRepresentation, err := newIndexTree(path)
	if err != nil {
		return "", err
	}

	return renderIndex(intermediateIndexRepresentation), nil
}

// renderIndex renders an HTML index from an indexTree
func renderIndex(it indexTree) string {
	return wrapWithTag(it.renderIndex(1), "ul")
}
