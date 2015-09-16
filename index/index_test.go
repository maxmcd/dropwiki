package index

import "testing"
import "log"
import "github.com/stretchr/testify/assert"

func TestRenderIndexRendersIndexFromAPath(t *testing.T) {
	testDropwikiPath := "./test_fixtures/root"
	actual, err := RenderIndex(testDropwikiPath)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	expected := "<ul>\n"
	expected += "<li><h1>root</h1></li>\n"
	expected += "<li>root_file.md</li>\n"
	expected += "<li><h2>level1</h2></li>\n"
	expected += "<li>level1_file.md</li>\n"
	expected += "<li>level1_other_file.md</li>\n"
	expected += "<li><h3>level2</h3></li>\n"
	expected += "<li>level2_file.md</li>\n"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}

func Test_renderIndex(t *testing.T) {
	f := directory{
		name: "root",
		contents: []node{
			page{name: "root_file.org"},
			directory{
				name: "level1",
				contents: []node{
					page{name: "level1_file.md"},
					directory{
						name: "level2",
						contents: []node{
							page{name: "level2_file.md"},
							page{name: "level2_other_file.md"},
						},
					},
				},
			},
		},
	}

	actual := renderIndex(f)

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

func Test_renderIndexRendersPagesBeforeFolders(t *testing.T) {
	f := directory{
		name: "foo",
		contents: []node{
			directory{name: "folder"},
			page{name: "page"},
		},
	}

	actual := renderIndex(f)

	expected := "<ul>\n"
	expected += "<li><h1>foo</h1></li>\n"
	expected += "<li>page</li>\n"
	expected += "<li><h2>folder</h2></li>\n"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}

func Test_newNodeFromTestDir(t *testing.T) {
	testDir := "./test_fixtures/root"
	actual, err := newNodeFrom(testDir)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	expected := directory{
		name: "root",
		contents: []node{
			directory{
				name: "level1",
				contents: []node{
					page{"level1_file.md"},
					page{"level1_other_file.md"},
					directory{
						name: "level2",
						contents: []node{
							page{"level2_file.md"},
						},
					},
				},
			},
			page{"root_file.md"},
		},
	}

	assert.Equal(t, expected, actual, "should be the same")
}