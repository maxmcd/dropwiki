package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestRenderIndexForEmptyFolder(t *testing.T) {
	f := Folder{
		Name:     "foo",
		Contents: []Node{},
	}

	actual := RenderIndex(f)

	expected := "<li><h1>foo</h1></li>\n"

	assert.Equal(t, expected, actual, "should be the same")
}

func TestRenderIndex(t *testing.T) {
	f := Folder{
		Name: "root",
		Contents: []Node{
			Page{
				Title: "root_file.org",
			},
			Folder{
				Name: "level1",
				Contents: []Node{
					Page{
						Title: "level1_file.md",
					},
					Folder{
						Name: "level2",
						Contents: []Node{
							Page{
								Title: "level2_file.md",
							},
							Page{
								Title: "level2_other_file.md",
							},
						},
					},
				},
			},
		},
	}

	actual := RenderIndex(f)

	expected := "<li><h1>root</h1></li>\n"
	expected += "<li>root_file.org</li>\n"
	expected += "<li><h2>level1</h2></li>\n"
	expected += "<li>level1_file.md</li>\n"
	expected += "<li><h3>level2</h3></li>\n"
	expected += "<li>level2_file.md</li>\n"
	expected += "<li>level2_other_file.md</li>\n"

	assert.Equal(t, expected, actual, "should be the same")
}
