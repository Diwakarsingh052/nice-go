package main

import "fmt"

func main() {
	//fetchData = 	/data
	//x := fetchData()
	//make([]int,100000)
	//orignalList = append(orignalList,x....)

	data := make([]int, 0, 1000000)
	lastCap := cap(data)
	var count int

	for r := 1; r <= 1000000; r++ {
		data = append(data, r)
		if lastCap != cap(data) {
			count++
			capCh := float64(cap(data)-lastCap) / float64(lastCap) * 100
			lastCap = cap(data)
			fmt.Printf("Add [%p] Cap[%d - %v]\n", data, cap(data), capCh)
		}
	}
	fmt.Println(count)
}
