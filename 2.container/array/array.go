package main
import "fmt"

func defineArray() {
	var arr1 [5]int
	arr2 := [3]int{1,2,33}
	arr3 := [...]int {1,2,3,5,66}
	var grid[4][5] bool

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for index, value := range arr3 {
		fmt.Println(index, value)
	}
}

func printArray(arr [5]int) {
	//array is passed by value, not reference, not change to 100
	for i, v := range arr {
		fmt.Println(i,v)
	}
	arr[0] = 100
}

func printPointerArray(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i,v)
	}
	arr[0] = 100
}

func main() {
	// defineArray()
	newArray := [...]int {1,2,3,5,66}
	printArray(newArray)
	fmt.Println(newArray)
	printPointerArray(&newArray)
	fmt.Println(newArray)
}
	
