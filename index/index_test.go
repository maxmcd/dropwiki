package index

import "testing"
import "log"
import "github.com/stretchr/testify/assert"

func TestRenderIndexRendersIndexFromAPath(t *testing.T) {
	testPath := "test_fixtures/root"
	actual, err := RenderIndex(testPath)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	expected := "<ul>"
	expected += "<li><a href='test_fixtures/root'><h1>Root</h1></a></li>"
	expected += "<li><a href='test_fixtures/root/root_page.md'>Root Page</a></li>"
	expected += "<li><a href='test_fixtures/root/level1'><h2>Level1</h2></a></li>"
	expected += "<li><a href='test_fixtures/root/level1/level1_other_page.md'>Level1 Other Page</a></li>"
	expected += "<li><a href='test_fixtures/root/level1/level1_page.md'>Level1 Page</a></li>"
	expected += "<li><a href='test_fixtures/root/level1/level2'><h3>Level2</h3></a></li>"
	expected += "<li><a href='test_fixtures/root/level1/level2/level2_page.md'>Level2 Page</a></li>"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}

func Test_renderIndex(t *testing.T) {
	f := section{
		title: "root",
		url:   "root_URL",
		contents: []node{
			page{title: "root page", url: "root_page_URL"},
			section{
				title: "level1",
				url:   "level1_URL",
				contents: []node{
					page{title: "level1 page", url: "level1_page_URL"},
					section{
						title: "level2",
						url:   "level2_URL",
						contents: []node{
							page{title: "level2 page", url: "level2_page_URL"},
							page{title: "level2 other page", url: "level2_other_page_URL"},
						},
					},
				},
			},
		},
	}

	actual := renderIndex(f)

	expected := "<ul>"
	expected += "<li><a href='root_URL'><h1>root</h1></a></li>"
	expected += "<li><a href='root_page_URL'>root page</a></li>"
	expected += "<li><a href='level1_URL'><h2>level1</h2></a></li>"
	expected += "<li><a href='level1_page_URL'>level1 page</a></li>"
	expected += "<li><a href='level2_URL'><h3>level2</h3></a></li>"
	expected += "<li><a href='level2_page_URL'>level2 page</a></li>"
	expected += "<li><a href='level2_other_page_URL'>level2 other page</a></li>"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "should be the same")
}

func Test_renderIndexRendersPagesBeforeSections(t *testing.T) {
	f := section{
		title: "foo",
		url:   "URL",
		contents: []node{
			section{title: "folder", url: "URL"},
			page{title: "page", url: "URL"},
		},
	}

	actual := renderIndex(f)

	expected := "<ul>"
	expected += "<li><a href='URL'><h1>foo</h1></a></li>"
	expected += "<li><a href='URL'>page</a></li>"
	expected += "<li><a href='URL'><h2>folder</h2></a></li>"
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
