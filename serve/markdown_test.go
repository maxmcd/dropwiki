package serve

import (
	"io/ioutil"
	"testing"
)

func TestGenerateHtml(t *testing.T) {
	html, err := generateHtml("./test_fixtures/markdown_test.md")
	if err != nil {
		t.Error(err)
	}
	output, err := ioutil.ReadFile("./test_fixtures/output.html")
	if err != nil {
		t.Error(err)
	}
	if string(output) != string(html) {
		t.Errorf("incorrect html output")
	}
}

func BenchmarkGenerateHtml(b *testing.B) {
	// run the Fib function b.N times
	var t *testing.T
	for n := 0; n < b.N; n++ {
		_, err := generateHtml("./test_fixtures/markdown_test.md")
		if err != nil {
			t.Error(err)
		}
	}
}
