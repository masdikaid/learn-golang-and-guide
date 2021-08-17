package basic

import "fmt"

// map is most like array but with the custom index and unlimited length, so you don't need to set length of map when create map
func MapData() {
	// create map
	map1 := make(map[string]string)
	// create map and value immediately
	map2 := map[string]string{
		"key1": "key1",
		"key2": "key2",
		"key3": "key3",
		"key4": "key4",
	}

	// set or change value
	map1["keya"] = "keyA"
	map1["keyb"] = "keyB"
	map1["keyc"] = "keyC"

	// deleting map value
	delete(map2, "key4")

	fmt.Println(map1)
	fmt.Println(map2)

	// access map value
	fmt.Println(map2["key2"])

}
