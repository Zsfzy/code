package main

import "fmt"

func main()  {
	arr := []int{87,233,12,3,5432,99999,5,34,65,456,456,23}
	fmt.Println(bubbleSort(arr))
}
func bubbleSort(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++  {
		for l := 0; l < arrLen-1-i ; l++  {
			if (arr[l] > arr[l+1]) {
				tmp := arr[l]
				arr[l] = arr[l+1]
				arr[l+1] = tmp
			}
		}
	}

	return arr
}
