package testing_in_go

import "testing"

// sub benchmark is like sub test you can use b.Run() and then write your benchmark
func TestSubTestTableMyFunc(t *testing.T) {
	test := [...]struct {
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
	for _, val := range test {
		t.Run(val.name, func(t *testing.T) {
			MyFunc()
		})
	}
}
