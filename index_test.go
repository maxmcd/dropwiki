package index

import "testing"
import "github.com/stretchr/testify/assert"

func TestRenderIndexForEmptyFolder(t *testing.T) {
	f := Directory{
		Name:     "foo",
		Contents: []Node{},
	}

	actual := RenderIndex(f)

	expected := "<ul>\n<li><h1>foo</h1></li>\n</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}

func TestRenderIndex(t *testing.T) {
	f := Directory{
		Name: "root",
		Contents: []Node{
			Page{Name: "root_file.org"},
			Directory{
				Name: "level1",
				Contents: []Node{
					Page{Name: "level1_file.md"},
					Directory{
						Name: "level2",
						Contents: []Node{
							Page{Name: "level2_file.md"},
							Page{Name: "level2_other_file.md"},
						},
					},
				},
			},
		},
	}

	actual := RenderIndex(f)

	expected := "<ul>\n"
	expected += "<li><h1>root</h1></li>\n"
	expected += "<li>root_file.org</li>\n"
	expected += "<li><h2>level1</h2></li>\n"
	expected += "<li>level1_file.md</li>\n"
	expected += "<li><h3>level2</h3></li>\n"
	expected += "<li>level2_file.md</li>\n"
	expected += "<li>level2_other_file.md</li>\n"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}

func TestRenderIndexRendersPagesBeforeFolders(t *testing.T) {
	f := Directory{
		Name: "foo",
		Contents: []Node{
			Directory{Name: "folder"},
			Page{Name: "page"},
		},
	}

	actual := RenderIndex(f)

	expected := "<ul>\n"
	expected += "<li><h1>foo</h1></li>\n"
	expected += "<li>page</li>\n"
	expected += "<li><h2>folder</h2></li>\n"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}
