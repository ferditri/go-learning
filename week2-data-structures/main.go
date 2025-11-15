package main

import "fmt"

func main() {

	// day 1-2

	// array

	var arr [5]int
	arr[0] = 12
	arr[1] = 23
	arr[2] = 55
	arr[3] = 99
	arr[4] = 52

	fmt.Println("Array: ", arr)
	fmt.Println("Array pertama: ", arr[0])
	fmt.Println("Panjang array: ", len(arr))

	buah := [3]string{"Apple", "Banana", "Grapes"}

	for x := 0; x < len(buah); x++ {
		fmt.Printf("Index %d: %s\n", x, buah[x])
	}

	for index, value := range buah {
		fmt.Printf("%d. %s\n", index+1, value)
	}
}
