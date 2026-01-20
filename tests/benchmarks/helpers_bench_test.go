package benchmarks

import (
	"testing"

	"github.com/TallantM/go-framework-demo/internal/utils"
)

func BenchmarkPostData(b *testing.B) {
	b.ReportAllocs() // Report memory allocations
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := utils.PostData("https://jsonplaceholder.typicode.com/posts", "title", "body")
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	// Optional: Add CPU profiling if needed (e.g., pprof.StartCPUProfile)
}

func BenchmarkGetData(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := utils.GetData("https://jsonplaceholder.typicode.com/posts/1")
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPatchData(b *testing.B) {
	updates := map[string]string{"title": "updated"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := utils.PatchData("https://jsonplaceholder.typicode.com/posts/1", updates)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
