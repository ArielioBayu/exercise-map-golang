package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	result := [][]string{}
	for key, value := range mapData {

		elements := []string{key, value}
		// fmt.Println(elements)

		result = append(result, elements)
		// fmt.Println(result)
	}
	return result // TODO: replace this
}

func main() {
	mapData := map[string]string{
		"hello": "world",
		"john":  "doe",
		"age":   "14",
	}
	fmt.Println(MapToSlice(mapData))
}
