package fuzz

import (
	"testing"

	"github.com/TallantM/go-framework-demo/internal/utils"
)

func FuzzPostData(f *testing.F) {
	f.Add("https://jsonplaceholder.typicode.com/posts", "title", "body")
	f.Fuzz(func(t *testing.T, url, title, body string) {
		_, _ = utils.PostData(url, title, body) // Ignore errors; focus on crashes or panics
	})
}

func FuzzPatchData(f *testing.F) {
	f.Add("https://jsonplaceholder.typicode.com/posts/1", "key:value")
	f.Fuzz(func(t *testing.T, url, updateStr string) {
		updates := map[string]string{"update": updateStr}
		_, _ = utils.PatchData(url, updates)
	})
}
