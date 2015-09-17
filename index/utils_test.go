package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toTitle(t *testing.T) {
	actual := toTitle("my_page.md")
	expected := "My Page"
	assert.Equal(t, expected, actual, "")
}

func Test_wrapWithTag(t *testing.T) {
	actual := wrapWithTag("foo", "div")
	expected := "<div>foo</div>"
	assert.Equal(t, expected, actual, "")
}

func Test_wrapWithAnchorTag(t *testing.T) {
	actual := wrapWithAnchorTag("foo", "http://foo.com")
	expected := "<a href='http://foo.com'>foo</a>"
	assert.Equal(t, expected, actual, "")
}
