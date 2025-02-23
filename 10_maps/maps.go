package main

import (
	"fmt"
	"maps"
)

//maps-dic, hash ,object

func main() {
	// m := make(map[string]string)
	// m["name"] = "golang"
	// m["area"] = "backend"
	// fmt.Println(m["name"], m["area"])

	//fmt.Println(m["name"], m["area"], m["phone"])
	//fmt.Println(len(m))

	// delete(m, "area")
	// fmt.Println(m)
	// clear(m)
	//fmt.Println(m)
	m := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}
	// fmt.Println(m)
	// v, ok := m["price"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }
	fmt.Println(maps.Equal(m, m2))
}
