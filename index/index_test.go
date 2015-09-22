package index

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateIndex_generates_an_HTML_index_from_a_path(t *testing.T) {
	testPath := "test_fixtures/root"
	actual, err := GenerateIndex(testPath)
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

	assert.Equal(t, expected, actual, "")
}

func Test_renderIndex_renders_an_HTML_index_from_an_indexTree(t *testing.T) {
	indexTree := section{
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

	actual := renderIndex(indexTree)

	expected := "<ul>"
	expected += "<li><a href='root_URL'><h1>root</h1></a></li>"
	expected += "<li><a href='root_page_URL'>root page</a></li>"
	expected += "<li><a href='level1_URL'><h2>level1</h2></a></li>"
	expected += "<li><a href='level1_page_URL'>level1 page</a></li>"
	expected += "<li><a href='level2_URL'><h3>level2</h3></a></li>"
	expected += "<li><a href='level2_page_URL'>level2 page</a></li>"
	expected += "<li><a href='level2_other_page_URL'>level2 other page</a></li>"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "")
}

func Test_renderIndex_puts_pages_before_sections(t *testing.T) {
	indexTree := section{
		title: "foo",
		url:   "URL",
		contents: []node{
			section{title: "folder", url: "URL"},
			page{title: "page", url: "URL"},
		},
	}

	actual := renderIndex(indexTree)

	expected := "<ul>"
	expected += "<li><a href='URL'><h1>foo</h1></a></li>"
	expected += "<li><a href='URL'>page</a></li>"
	expected += "<li><a href='URL'><h2>folder</h2></a></li>"
	expected += "</ul>"

	assert.Equal(t, expected, actual, "")
}
