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
	expected += "<li><h1>Root</h1></li>\n"
	expected += "<li>Root Page</li>\n"
	expected += "<li><h2>Level1</h2></li>\n"
	expected += "<li>Level1 Other Page</li>\n"
	expected += "<li>Level1 Page</li>\n"
	expected += "<li><h3>Level2</h3></li>\n"
	expected += "<li>Level2 Page</li>\n"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}

func Test_renderIndex(t *testing.T) {
	f := section{
		title: "root",
		contents: []node{
			page{title: "root page.org"},
			section{
				title: "level1",
				contents: []node{
					page{title: "level1 page"},
					section{
						title: "level2",
						contents: []node{
							page{title: "level2 page"},
							page{title: "level2 other page"},
						},
					},
				},
			},
		},
	}

	actual := renderIndex(f)

	expected := "<ul>\n"
	expected += "<li><h1>root</h1></li>\n"
	expected += "<li>root page.org</li>\n"
	expected += "<li><h2>level1</h2></li>\n"
	expected += "<li>level1 page</li>\n"
	expected += "<li><h3>level2</h3></li>\n"
	expected += "<li>level2 page</li>\n"
	expected += "<li>level2 other page</li>\n"
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
		title: "Root",
		contents: []node{
			section{
				title: "Level1",
				contents: []node{
					page{"Level1 Other Page"},
					page{"Level1 Page"},
					section{
						title: "Level2",
						contents: []node{
							page{"Level2 Page"},
						},
					},
				},
			},
			page{"Root Page"},
		},
	}

	assert.Equal(t, expected, actual, "should be the same")
}
