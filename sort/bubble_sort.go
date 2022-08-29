package main

import "fmt"

func bubble_sort(datalist []int, ascending bool) []int {
	temp := 0

	for i := 0; i < len(datalist); i++ {

		for k := i + 1; k < len(datalist); k++ {
			if ascending == true {

				if datalist[i] > datalist[k] {
					temp = datalist[i]
					datalist[i] = datalist[k]
					datalist[k] = temp
				}
			}
			if ascending == false {

				if datalist[i] < datalist[k] {
					temp = datalist[i]
					datalist[i] = datalist[k]
					datalist[k] = temp
				}
			}

		}

	}
	return datalist
}

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(bubble_sort(items, false))
}
