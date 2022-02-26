package main

import (
	"fmt"
	"strings"
)

type FilterCallback func(string) bool // alias scheme closure

func filter(data []string, callback FilterCallback) []string {
	var result []string

	for _, each := range data {
		// if filtered == true
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o") // bool result
	})

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5 // bool result
	})

	fmt.Println("Data asli\t\t: ", data)
	fmt.Println("Filter ada huruf \"o\"\t: ", dataContainsO)
	fmt.Println("Filter jumlah huruf \"5\"\t: ", dataLength5)

}
