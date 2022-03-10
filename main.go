package main

import (
	"fmt"
	"os"
)

type Books struct {
	name   string
	writer string
}

var data = []Books{}

//searching data
func search(key string) *Books {
	for i, v := range data {
		if v.name == key {
			return &data[i]
		}
	}
	return nil
}

//listing data
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {

	arguments := os.Args
	//appending data to struct
	data = append(data, Books{"book1", "writer1"})
	data = append(data, Books{"book2", "writer2"})
	data = append(data, Books{"book3", "writer3"})

	switch arguments[1] {

	case "search":

		result := search(arguments[2])
		if result == nil {
			fmt.Println("not found -> ", arguments[2])
			return
		}
		fmt.Println(*result)

	case "list":
		list()
	default:
		fmt.Println("not found")
	}
}
