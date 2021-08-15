package testing_in_go

import "testing"

// sub benchmark is like sub test you can use b.Run() and then write your benchmark
func BenchmarkSubMyFunc(b *testing.B) {
	benchmark := [...]struct {
		id   int
		name string
	}{
		{
			id:   1,
			name: "dika",
		},
		{
			id:   2,
			name: "adi",
		},
		{
			id:   3,
			name: "setiawan",
		},
	}
	for _, val := range benchmark {
		b.Run(val.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MyFunc()
			}
		})
	}
}
