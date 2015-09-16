package index

import "testing"
import "github.com/stretchr/testify/assert"

func Test_toTitle(t *testing.T) {
	actual := toTitle("my_page.md")
	expected := "my page"
	assert.Equal(t, expected, actual, "")
}
