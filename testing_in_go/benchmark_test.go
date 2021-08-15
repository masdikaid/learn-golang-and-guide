package testing_in_go

import "testing"

// benchmark test can help know you how fast your code
// basically benchmark will run your code many times that determine by b.N
// you can run with add -bench in your command and will be this "go test ./... -v -bench ."
// or you can chose your specific benchmark to run "go test ./... -v -bench BenchmarkMyFunc"
func BenchmarkMyFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyFunc()
	}
}
