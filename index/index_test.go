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

	expected := "<ul>"
	expected += "<li><h1>Root</h1></li>"
	expected += "<li>Root Page</li>"
	expected += "<li><h2>Level1</h2></li>"
	expected += "<li>Level1 Other Page</li>"
	expected += "<li>Level1 Page</li>"
	expected += "<li><h3>Level2</h3></li>"
	expected += "<li>Level2 Page</li>"
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

	expected := "<ul>"
	expected += "<li><h1>root</h1></li>"
	expected += "<li>root page.org</li>"
	expected += "<li><h2>level1</h2></li>"
	expected += "<li>level1 page</li>"
	expected += "<li><h3>level2</h3></li>"
	expected += "<li>level2 page</li>"
	expected += "<li>level2 other page</li>"
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

	expected := "<ul>"
	expected += "<li><h1>foo</h1></li>"
	expected += "<li>page</li>"
	expected += "<li><h2>folder</h2></li>"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}

func Test_newNodeFromTestDir(t *testing.T) {
	testDir := "test_fixtures/root"
	actual, err := newNodeFrom(testDir)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	expected := section{
		title: "Root",
		url:   "test_fixtures/root",
		contents: []node{
			section{
				title: "Level1",
				url:   "test_fixtures/root/level1",
				contents: []node{
					page{title: "Level1 Other Page", url: "test_fixtures/root/level1/level1_other_page.md"},
					page{title: "Level1 Page", url: "test_fixtures/root/level1/level1_page.md"},
					section{
						title: "Level2",
						url:   "test_fixtures/root/level1/level2",
						contents: []node{
							page{title: "Level2 Page", url: "test_fixtures/root/level1/level2/level2_page.md"},
						},
					},
				},
			},
			page{title: "Root Page", url: "test_fixtures/root/root_page.md"},
		},
	}

	assert.Equal(t, expected, actual, "should be the same")
}
