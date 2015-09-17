package index

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newIndexTree_generates_an_indexTree_from_a_path(t *testing.T) {
	testDir := "test_fixtures/root"
	actual, err := newIndexTree(testDir)
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

	assert.Equal(t, expected, actual, "")
}
