package index

// node is the intermediate representation of a directory/file node
type node interface {
	sortingWeight() nodeSortingWeight
	value() string
	children() []node
	renderIndex(int) string
}

// nodeSortingWeight determines the sorting order for elements
// in a []node slice
type nodeSortingWeight int

const (
	pageSortingWeight nodeSortingWeight = iota // pages come before sections
	sectionSortingWeight
)

// byWeight implements the sort.Interface for []node based
// on the sortingWeight of the elements
type byWeight []node

func (ns byWeight) Len() int           { return len(ns) }
func (ns byWeight) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }
func (ns byWeight) Less(i, j int) bool { return ns[i].sortingWeight() < ns[j].sortingWeight() }
