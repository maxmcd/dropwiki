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
	expected += "<li>root_page.md</li>\n"
	expected += "<li><h2>level1</h2></li>\n"
	expected += "<li>level1_other_page.md</li>\n"
	expected += "<li>level1_page.md</li>\n"
	expected += "<li><h3>level2</h3></li>\n"
	expected += "<li>level2_page.md</li>\n"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}

func Test_renderIndex(t *testing.T) {
	f := section{
		title: "root",
		contents: []node{
			page{title: "root_page.org"},
			section{
				title: "level1",
				contents: []node{
					page{title: "level1_page.md"},
					section{
						title: "level2",
						contents: []node{
							page{title: "level2_page.md"},
							page{title: "level2_other_page.md"},
						},
					},
				},
			},
		},
	}

	actual := renderIndex(f)

	expected := "<ul>\n"
	expected += "<li><h1>root</h1></li>\n"
	expected += "<li>root_page.org</li>\n"
	expected += "<li><h2>level1</h2></li>\n"
	expected += "<li>level1_page.md</li>\n"
	expected += "<li><h3>level2</h3></li>\n"
	expected += "<li>level2_page.md</li>\n"
	expected += "<li>level2_other_page.md</li>\n"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}

func Test_renderIndexRendersPagesBeforeSections(t *testing.T) {
	f := section{
		title: "foo",
		contents: []node{
			section{title: "folder"},
			page{title: "page"},
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
	expected := section{
		title: "root",
		contents: []node{
			section{
				title: "level1",
				contents: []node{
					page{"level1_other_page.md"},
					page{"level1_page.md"},
					section{
						title: "level2",
						contents: []node{
							page{"level2_page.md"},
						},
					},
				},
			},
			page{"root_page.md"},
		},
	}

	assert.Equal(t, expected, actual, "should be the same")
}
